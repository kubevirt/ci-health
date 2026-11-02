package gh

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	retry "github.com/avast/retry-go"
	"github.com/shurcooL/githubv4"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) { return f(req) }

type trackedBody struct {
	io.Reader
	closed bool
}

func (b *trackedBody) Close() error { b.closed = true; return nil }

type networkError struct {
	timeout, temporary bool
}

func (e networkError) Error() string   { return "network failure" }
func (e networkError) Timeout() bool   { return e.timeout }
func (e networkError) Temporary() bool { return e.temporary }

func graphQLResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Status:     fmt.Sprintf("%d %s", status, http.StatusText(status)),
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func graphQLClient(transport http.RoundTripper) *githubv4.Client {
	return githubv4.NewEnterpriseClient("https://github.example/graphql", &http.Client{Transport: transport})
}

type viewerQuery struct {
	Viewer struct {
		Login string
		Name  string
	}
}

func TestQueryWithRetry(t *testing.T) {
	const success = `{"data":{"viewer":{"login":"octocat"}}}`
	tests := []struct {
		name     string
		status   int
		body     string
		err      error
		attempts int
		wantErr  string
	}{
		{name: "immediate success", status: 200, body: success, attempts: 1},
		{name: "empty response", status: 200, attempts: 2},
		{name: "truncated response", status: 200, body: `{"data":{"viewer":`, attempts: 2},
		{name: "wrapped EOF", err: fmt.Errorf("transport: %w", io.EOF), attempts: 2},
		{name: "wrapped unexpected EOF", err: fmt.Errorf("transport: %w", io.ErrUnexpectedEOF), attempts: 2},
		{name: "timeout", err: networkError{timeout: true}, attempts: 2},
		{name: "wrapped timeout", err: fmt.Errorf("transport: %w", networkError{timeout: true}), attempts: 2},
		{name: "non-timeout temporary error", err: networkError{temporary: true}, attempts: 1, wantErr: "network failure"},
		{name: "bad gateway", status: 502, attempts: 2},
		{name: "unavailable", status: 503, attempts: 2},
		{name: "gateway timeout", status: 504, attempts: 2},
		{name: "unauthorized", status: 401, body: "bad credentials", attempts: 1, wantErr: "401"},
		{name: "forbidden", status: 403, attempts: 1, wantErr: "403"},
		{name: "GraphQL error with partial data", status: 200, body: `{"data":{"viewer":{"login":"partial"}},"errors":[{"message":"access denied"}]}`, attempts: 1, wantErr: "access denied"},
		{name: "schema error after decoded field", status: 200, body: `{"data":{"viewer":{"login":"partial","unknown":"bad"}}}`, attempts: 1, wantErr: "unknown"},
		{name: "permanent network error", err: networkError{}, attempts: 1, wantErr: "network failure"},
		{name: "canceled transport", err: fmt.Errorf("transport: %w", context.Canceled), attempts: 1, wantErr: "context canceled"},
		{name: "deadline transport", err: fmt.Errorf("transport: %w", context.DeadlineExceeded), attempts: 1, wantErr: "context deadline exceeded"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attempts := 0
			client := graphQLClient(&githubTransport{inner: roundTripFunc(func(*http.Request) (*http.Response, error) {
				attempts++
				if attempts > 1 {
					return graphQLResponse(200, success), nil
				}
				if tt.err != nil {
					return nil, tt.err
				}
				return graphQLResponse(tt.status, tt.body), nil
			})})
			var result viewerQuery
			result.Viewer.Login, result.Viewer.Name = "original", "must not survive success"
			original := result
			// RandomDelay panics with MaxJitter(0), so also select FixedDelay.
			err := queryWithRetry(context.Background(), client, &result, nil,
				retry.Delay(0), retry.MaxJitter(0), retry.DelayType(retry.FixedDelay))
			if tt.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("error = %v, want containing %q", err, tt.wantErr)
				}
				if result != original {
					t.Errorf("failed query changed destination: %+v", result)
				}
			} else if err != nil {
				t.Fatalf("query failed: %v", err)
			} else if result.Viewer.Login != "octocat" || result.Viewer.Name != "" {
				t.Errorf("result not decoded into fresh state: %+v", result)
			}
			if attempts != tt.attempts {
				t.Errorf("attempts = %d, want %d", attempts, tt.attempts)
			}
		})
	}
}

func TestQueryWithRetryBudget(t *testing.T) {
	for _, mode := range []string{"EOF", "truncated response", "mixed HTTP and EOF", "success on fifth"} {
		t.Run(mode, func(t *testing.T) {
			attempts := 0
			client := graphQLClient(&githubTransport{inner: roundTripFunc(func(*http.Request) (*http.Response, error) {
				attempts++
				if mode == "success on fifth" && attempts == 5 {
					return graphQLResponse(200, `{"data":{"viewer":{"login":"octocat"}}}`), nil
				}
				if mode == "mixed HTTP and EOF" && attempts < 5 {
					return graphQLResponse([]int{502, 503, 504, 502}[attempts-1], "unavailable"), nil
				}
				if mode == "truncated response" {
					return graphQLResponse(200, `{"data":{"viewer":{"login":"partial"}`), nil
				}
				return nil, fmt.Errorf("attempt %d: %w", attempts, io.EOF)
			})})
			var result viewerQuery
			result.Viewer.Login = "original"
			original := result
			err := queryWithRetry(context.Background(), client, &result, nil,
				retry.Delay(0), retry.MaxJitter(0), retry.DelayType(retry.FixedDelay))
			if attempts != 5 {
				t.Fatalf("attempts = %d, want exactly 5", attempts)
			}
			if mode == "success on fifth" {
				if err != nil || result.Viewer.Login != "octocat" {
					t.Fatalf("result = %+v, error = %v", result, err)
				}
			} else {
				if mode == "truncated response" && !errors.Is(err, io.ErrUnexpectedEOF) {
					t.Errorf("error = %v, want unexpected EOF", err)
				} else if mode != "truncated response" && (!errors.Is(err, io.EOF) || !strings.Contains(err.Error(), "attempt 5")) {
					t.Errorf("error = %v, want last attempt's EOF", err)
				}
				if result != original {
					t.Errorf("exhausted query changed destination: %+v", result)
				}
			}
		})
	}
}

func TestQueryWithRetryCancellation(t *testing.T) {
	for _, before := range []bool{true, false} {
		t.Run(fmt.Sprintf("before_call=%t", before), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			if before {
				cancel()
			}
			attempts := 0
			client := graphQLClient(roundTripFunc(func(*http.Request) (*http.Response, error) {
				attempts++
				return nil, io.EOF
			}))
			backoff := make(chan struct{})
			options := []retry.Option{retry.Delay(time.Hour), retry.MaxJitter(0), retry.DelayType(func(uint, error, *retry.Config) time.Duration {
				close(backoff)
				return time.Hour
			})}
			var result viewerQuery
			result.Viewer.Login = "original"
			original := result
			done := make(chan error, 1)
			go func() { done <- queryWithRetry(ctx, client, &result, nil, options...) }()
			if !before {
				select {
				case <-backoff:
					cancel()
				case <-time.After(5 * time.Second):
					t.Fatal("query did not reach backoff")
				}
			}
			select {
			case err := <-done:
				if !errors.Is(err, context.Canceled) {
					t.Errorf("error = %v, want context.Canceled", err)
				}
			case <-time.After(5 * time.Second):
				t.Fatal("cancellation did not interrupt backoff")
			}
			wantAttempts := 1
			if before {
				wantAttempts = 0
			}
			if attempts != wantAttempts || result != original {
				t.Errorf("attempts = %d, result = %+v", attempts, result)
			}
		})
	}
}

func TestGithubTransportRoundTrip(t *testing.T) {
	for _, status := range []int{200, 401, 403, 429, 500, 502, 503, 504} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			body := &trackedBody{Reader: strings.NewReader("response")}
			response := graphQLResponse(status, "")
			response.Body = body
			request, _ := http.NewRequest(http.MethodPost, "https://github.example/graphql", nil)
			calls := 0
			transport := &githubTransport{inner: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				calls++
				if req != request {
					t.Error("underlying transport received a different request")
				}
				return response, nil
			})}
			got, err := transport.RoundTrip(request)
			transient := status == 502 || status == 503 || status == 504
			if transient {
				if got != nil || !errors.Is(err, errTransientHTTP) || !body.closed {
					t.Errorf("response = %v, error = %v, closed = %t", got, err, body.closed)
				}
			} else if got != response || err != nil || body.closed {
				t.Errorf("passthrough response = %v, error = %v, closed = %t", got, err, body.closed)
			}
			if calls != 1 {
				t.Errorf("RoundTrip made %d attempts, want 1", calls)
			}
			if err := body.Close(); err != nil {
				t.Fatal(err)
			}
		})
	}
	t.Run("underlying failure", func(t *testing.T) {
		failure := errors.New("transport failed")
		calls := 0
		transport := &githubTransport{inner: roundTripFunc(func(*http.Request) (*http.Response, error) {
			calls++
			return nil, failure
		})}
		request, _ := http.NewRequest(http.MethodGet, "https://github.example/graphql", nil)
		response, err := transport.RoundTrip(request)
		if response != nil || err != failure || calls != 1 {
			t.Fatalf("response = %v, error = %v, calls = %d", response, err, calls)
		}
	})
}

func TestClientQueries(t *testing.T) {
	const search = "repo:owner/repo type:pr is:merged"
	tests := []struct {
		name      string
		query     []string
		variables map[string]interface{}
		body      string
		call      func(*testing.T, *Client)
	}{
		{
			name: "merge queue", query: []string{"search(query: $querystring, type: ISSUE, first:100)", "... on PullRequest", "LABELED_EVENT, UNLABELED_EVENT"},
			variables: map[string]interface{}{"querystring": search},
			body:      `{"data":{"search":{"nodes":[{"number":42}]}}}`,
			call: func(t *testing.T, c *Client) {
				got, err := c.mergeQueuePRQuery(search)
				if err != nil || len(got) != 1 || got[0].Number != 42 {
					t.Fatalf("result = %+v, error = %v", got, err)
				}
			},
		},
		{
			name: "chatops", query: []string{"search(query: $querystring, type: ISSUE, first:100)", "ISSUE_COMMENT", "PULL_REQUEST_COMMIT"},
			variables: map[string]interface{}{"querystring": search},
			body:      `{"data":{"search":{"nodes":[{"number":43,"timelineItems":{"nodes":[{"bodyText":"/retest"}]}}]}}}`,
			call: func(t *testing.T, c *Client) {
				got, err := c.chatopsPRQuery(search)
				if err != nil || len(got) != 1 || got[0].Number != 43 || len(got[0].Nodes) != 1 || got[0].Nodes[0].BodyText != "/retest" {
					t.Fatalf("result = %+v, error = %v", got, err)
				}
			},
		},
		{
			name: "timeline", query: []string{"repository(owner: $owner, name: $repo)", "pullRequest(number: $number)", "LABELED_EVENT, UNLABELED_EVENT"},
			variables: map[string]interface{}{"owner": "owner", "repo": "repo", "number": float64(42)},
			body:      `{"data":{"repository":{"pullRequest":{"number":42,"timelineItems":{"nodes":[{"addedLabel":{"name":"lgtm"}}]}}}}}`,
			call: func(t *testing.T, c *Client) {
				got, err := c.FetchPRTimelineItems(42)
				if err != nil || got == nil || got.Number != 42 || len(got.Nodes) != 1 || got.Nodes[0].AddedLabel.Name != "lgtm" {
					t.Fatalf("result = %+v, error = %v", got, err)
				}
			},
		},
		{
			name: "supported branches", query: []string{`repository(owner: "kubevirt", name: "kubevirt")`, `refs(refPrefix: "refs/heads/", first: 100)`},
			body: `{"data":{"repository":{"refs":{"nodes":[{"name":"release-1.5"},{"name":"main"},{"name":"release-1.7"},{"name":"feature"},{"name":"release-1.6"},{"name":"release-1.4"}]}}}}`,
			call: func(t *testing.T, c *Client) {
				got, err := c.GetSupportedBranches(context.Background())
				if err != nil || !reflect.DeepEqual(got, []string{"main", "release-1.7", "release-1.6", "release-1.5"}) {
					t.Fatalf("result = %v, error = %v", got, err)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calls := 0
			client := &Client{source: "owner/repo", inner: graphQLClient(roundTripFunc(func(req *http.Request) (*http.Response, error) {
				calls++
				var payload struct {
					Query     string
					Variables map[string]interface{}
				}
				if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
					t.Fatalf("decode request: %v", err)
				}
				if err := req.Body.Close(); err != nil {
					t.Fatal(err)
				}
				if req.Method != http.MethodPost || !reflect.DeepEqual(payload.Variables, tt.variables) {
					t.Errorf("method = %s, variables = %v, want %v", req.Method, payload.Variables, tt.variables)
				}
				for _, fragment := range tt.query {
					if !strings.Contains(payload.Query, fragment) {
						t.Errorf("query %q missing %q", payload.Query, fragment)
					}
				}
				return graphQLResponse(200, tt.body), nil
			}))}
			tt.call(t, client)
			if calls != 1 {
				t.Errorf("requests = %d, want 1", calls)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	path := filepath.Join(t.TempDir(), "token")
	if err := os.WriteFile(path, []byte("  secret-token\n"), 0600); err != nil {
		t.Fatal(err)
	}
	// Intercept the OAuth transport's base without reaching the public GitHub API.
	originalTransport := http.DefaultTransport
	t.Cleanup(func() { http.DefaultTransport = originalTransport })
	calls := 0
	body := &trackedBody{Reader: strings.NewReader("unavailable")}
	http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		calls++
		defer func() {
			if err := req.Body.Close(); err != nil {
				t.Error(err)
			}
		}()
		if got := req.Header.Get("Authorization"); got != "Bearer secret-token" {
			t.Errorf("Authorization = %q", got)
		}
		if req.URL.String() != "https://api.github.com/graphql" {
			t.Errorf("endpoint = %s", req.URL)
		}
		response := graphQLResponse(503, "")
		response.Body = body
		return response, nil
	})
	client, err := NewClient(path, "owner/repo")
	if err != nil {
		t.Fatal(err)
	}
	if client.source != "owner/repo" {
		t.Errorf("source = %q", client.source)
	}
	var query viewerQuery
	// A direct query verifies githubTransport is installed outside OAuth,
	// classifies the response, closes its body, and does not retry itself.
	err = client.inner.Query(context.Background(), &query, nil)
	if !errors.Is(err, errTransientHTTP) || calls != 1 || !body.closed {
		t.Errorf("error = %v, calls = %d, body closed = %t", err, calls, body.closed)
	}
	missing, err := NewClient(filepath.Join(t.TempDir(), "missing"), "owner/repo")
	if missing != nil || !errors.Is(err, os.ErrNotExist) {
		t.Errorf("missing token: client = %v, error = %v", missing, err)
	}
}
