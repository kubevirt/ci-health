package cifailures

import "testing"

func TestCategorizeError(t *testing.T) {
	tests := []struct {
		name      string
		errorText string
		expected  ErrorCategory
	}{
		// External patterns
		{
			name:      "GitHub download 404",
			errorText: `13:34:47: ERROR: /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/buildifier_prebuilt_toolchains/BUILD.bazel:26:18: @buildifier_prebuilt_toolchains//:buildifier_linux_amd64 depends on @buildifier_linux_amd64//file:buildifier in repository @buildifier_linux_amd64 which failed to fetch. no such package '@buildifier_linux_amd64//file': java.io.IOException: Error downloading [https://github.com/bazelbuild/buildtools/releases/download/v7.3.1/buildifier-linux-amd64] to /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/buildifier_linux_amd64/file/buildifier: GET returned 404 Not Found`,
			expected:  CategoryExternal,
		},
		{
			name:      "GitHub download 502",
			errorText: `16:52:49: ERROR: Error computing the main repository mapping: no such package '@aspect_bazel_lib//lib': java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /root/.cache/bazel/temp/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error`,
			expected:  CategoryExternal,
		},
		{
			name:      "podman container removal timeout",
			errorText: `15:45:00: Error: cannot remove container fdb28089f2bfccabdf0a36a9d248f6e2fead25b90ed175dc5805f3821801cb67 as it could not be stopped: given PID did not die within timeout`,
			expected:  CategoryExternal,
		},
		{
			name:      "podman storage file timeout",
			errorText: `18:01:07: Error: timed out waiting for file /var/lib/containers/storage/overlay-containers/5bc99eb00b10da749245414b8f8e6618a36bc3cd00e5284797ff1aafe1d47fc7/userdata/29017a592eaba4813b285a2991ab2453f22c71d973f44d9efd9931b4bc2fd0a0/exit/5bc99eb00b10da749245414b8f8e6618a36bc3cd00e5284797ff1aafe1d47fc7`,
			expected:  CategoryExternal,
		},
		{
			name:      "kind cluster deletion failure",
			errorText: `15:45:00: ERROR: failed to delete cluster "vgpu": failed to delete nodes: command "podman rm -f -v vgpu-control-plane" failed with error: exit status 125`,
			expected:  CategoryExternal,
		},
		{
			name:      "kube-apiserver body decode noise",
			errorText: `00:43:10: I0318 00:43:08.163777     241 request.go:1500] "Body was not decodable (unable to check for Status)" err="couldn't get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \"json:\\\"apiVersion,omitempty\\\"\"; Kind string \"json:\\\"kind,omitempty\\\"\" }"`,
			expected:  CategoryExternal,
		},
		{
			name:      "API rate limiter timeout",
			errorText: `09:52:59:           msg: "client rate limiter Wait returned an error: context deadline exceeded",`,
			expected:  CategoryExternal,
		},

		// PR Build patterns
		{
			name:      "bazel analysis failure",
			errorText: `13:34:47: ERROR: Analysis of target '//:buildifier' failed; build aborted:`,
			expected:  CategoryPRBuild,
		},
		{
			name:      "bazel analysis failure with target path",
			errorText: `08:33:09: ERROR: Analysis of target '//cmd/virt-launcher:virt-launcher-image' failed; build aborted:`,
			expected:  CategoryPRBuild,
		},
		{
			name:      "bazel build failed",
			errorText: `07:18:28: ERROR: Build failed. Not running target`,
			expected:  CategoryPRBuild,
		},
		{
			name:      "make bazel-build-images failure",
			errorText: `make: *** [Makefile:37: bazel-build-images] Error 125`,
			expected:  CategoryPRBuild,
		},
		{
			name:      "make bazel-build-functests failure",
			errorText: `make: *** [Makefile:26: bazel-build-functests] Error 125`,
			expected:  CategoryPRBuild,
		},

		// Internal patterns
		{
			name:      "taint control-plane not found",
			errorText: `16:41:33: error: taint "node-role.kubernetes.io/control-plane:NoSchedule" not found`,
			expected:  CategoryInternal,
		},
		{
			name:      "taint master not found",
			errorText: `16:41:33: error: taint "node-role.kubernetes.io/master:NoSchedule" not found`,
			expected:  CategoryInternal,
		},
		{
			name:      "make cluster-up failure",
			errorText: `make: *** [Makefile:173: cluster-up] Error 2`,
			expected:  CategoryInternal,
		},
		{
			name:      "make cluster-sync failure",
			errorText: `make: *** [Makefile:174: cluster-sync] Error 125`,
			expected:  CategoryInternal,
		},
		{
			name:      "make cluster-down failure",
			errorText: `make: *** [Makefile:176: cluster-down] Error 6`,
			expected:  CategoryInternal,
		},
		{
			name:      "kubevirt deployment timeout",
			errorText: `10:41:43: error: timed out waiting for the condition on kubevirts/kubevirt`,
			expected:  CategoryInternal,
		},
		{
			name:      "JWS signature not found",
			errorText: `02:52:21: I0319 02:52:13.772217     238 token.go:249] [discovery] Retrying due to error: could not find a JWS signature in the cluster-info ConfigMap for token ID "abcdef"`,
			expected:  CategoryInternal,
		},

		// Needs investigation
		{
			name:      "unknown error",
			errorText: `some completely unknown error message`,
			expected:  CategoryNeedsInvestigation,
		},
		{
			name:      "empty string",
			errorText: ``,
			expected:  CategoryNeedsInvestigation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CategorizeError(tt.errorText)
			if got != tt.expected {
				t.Errorf("CategorizeError() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestCategorizeJobBuildError(t *testing.T) {
	t.Run("no snippets returns needs-investigation", func(t *testing.T) {
		err := &JobBuildError{}
		if got := CategorizeJobBuildError(err); got != CategoryNeedsInvestigation {
			t.Errorf("CategorizeJobBuildError() = %q, want %q", got, CategoryNeedsInvestigation)
		}
	})

	t.Run("uses first snippet for categorization", func(t *testing.T) {
		err := &JobBuildError{
			BuildLogErrorSnippets: []*BuildLogErrorSnippet{
				{ErrorText: `16:41:33: error: taint "node-role.kubernetes.io/control-plane:NoSchedule" not found`},
				{ErrorText: `some other error`},
			},
		}
		if got := CategorizeJobBuildError(err); got != CategoryInternal {
			t.Errorf("CategorizeJobBuildError() = %q, want %q", got, CategoryInternal)
		}
	})
}
