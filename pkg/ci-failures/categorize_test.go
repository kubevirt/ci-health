package cifailures

import "testing"

func TestCategorizeError(t *testing.T) {
	tests := []struct {
		name      string
		errorText string
		expected  ErrorCategory
	}{
		// === External patterns ===
		{
			name:      "GitHub download 404",
			errorText: `13:34:47: ERROR: java.io.IOException: Error downloading [https://github.com/bazelbuild/buildtools/releases/download/v7.3.1/buildifier-linux-amd64] to /tmp/buildifier: GET returned 404 Not Found`,
			expected:  CategoryExternal,
		},
		{
			name:      "GitHub download 502",
			errorText: `16:52:49: ERROR: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/bazel-lib: GET returned 502 Bad Gateway or Proxy Error`,
			expected:  CategoryExternal,
		},
		{
			name:      "mirror download 404",
			errorText: `ERROR: Error downloading [http://mirror.stream.centos.org/9-stream/AppStream/x86_64/os/Packages/libvirt-client-10.10.0-13.el9.x86_64.rpm] to /tmp/downloaded: GET returned 404 Not Found`,
			expected:  CategoryExternal,
		},
		{
			name:      "bazel fetch repository error",
			errorText: `15:38:45: ERROR: An error occurred during the fetch of repository 'oci_regctl_linux_amd64':`,
			expected:  CategoryExternal,
		},
		{
			name:      "bazel fetch repository buildifier",
			errorText: `09:21:03: ERROR: An error occurred during the fetch of repository 'buildifier_linux_amd64':`,
			expected:  CategoryExternal,
		},
		{
			name:      "bazel fetch repository bazel_features",
			errorText: `13:59:25: ERROR: An error occurred during the fetch of repository 'bazel_features':`,
			expected:  CategoryExternal,
		},
		{
			name:      "bazel fetch repository rpm",
			errorText: `ERROR: An error occurred during the fetch of repository 'libvirt-client-0__10.10.0-13.el9.x86_64':`,
			expected:  CategoryExternal,
		},
		{
			name:      "WORKSPACE fetch http_file rule",
			errorText: `09:21:03: ERROR: /root/go/src/kubevirt.io/kubevirt/WORKSPACE:84:40: fetching http_file rule //external:buildifier_linux_amd64: Traceback (most recent call last):`,
			expected:  CategoryExternal,
		},
		{
			name:      "WORKSPACE fetch http_archive rule",
			errorText: `13:59:25: ERROR: /root/go/src/kubevirt.io/kubevirt/WORKSPACE:36:23: fetching http_archive rule //external:bazel_features: Traceback (most recent call last):`,
			expected:  CategoryExternal,
		},
		{
			name:      "WORKSPACE fetch go_repository rule",
			errorText: `ERROR: /root/go/src/kubevirt.io/kubevirt/WORKSPACE:267:14: fetching go_repository rule //external:com_github_masterzen_winrmcli: Traceback (most recent call last):`,
			expected:  CategoryExternal,
		},
		{
			name:      "WORKSPACE fetch rpm rule",
			errorText: `ERROR: /home/prow/go/src/github.com/kubevirt/kubevirt/WORKSPACE:4823:4: fetching rpm rule //external:libvirt-client-0__10.10.0-13.el9.x86_64: Traceback (most recent call last):`,
			expected:  CategoryExternal,
		},
		{
			name:      "WORKSPACE fetch regctl_repositories rule",
			errorText: `20:48:14: ERROR: /root/go/src/kubevirt.io/kubevirt/WORKSPACE:40:24: fetching regctl_repositories rule //external:oci_regctl_linux_amd64: Traceback (most recent call last):`,
			expected:  CategoryExternal,
		},
		{
			name:      "failed to fetch no such package",
			errorText: `20:48:14: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-operator/BUILD.bazel:47:10: //cmd/virt-operator:virt-operator-image depends on @oci_regctl_linux_amd64//:regctl_toolchain in repository @oci_regctl_linux_amd64 which failed to fetch. no such package '@oci_regctl_linux_amd64//'`,
			expected:  CategoryExternal,
		},
		{
			name:      "RPM depends on failed fetch",
			errorText: `ERROR: /home/prow/go/src/github.com/kubevirt/kubevirt/rpm/BUILD.bazel:1031:8: //rpm:launcherbase_x86_64 depends on @libvirt-client-0__10.10.0-13.el9.x86_64//rpm:rpm in repository @libvirt-client-0__10.10.0-13.el9.x86_64 which failed to fetch. no such package '@libvirt-client-0__10.10.0-13.el9.x86_64//rpm'`,
			expected:  CategoryExternal,
		},
		{
			name:      "quay.io image pull 500",
			errorText: `20:44:21: Error: unable to copy from source docker://quay.io/kubevirt/builder:2510281240-82d14f4b41: copying system image from manifest list: parsing image configuration: fetching blob: received unexpected HTTP status: 500 Internal Server Error`,
			expected:  CategoryExternal,
		},
		{
			name:      "quay.io image pull 502",
			errorText: `Error: unable to copy from source docker://quay.io/kubevirt/builder:2510281240-82d14f4b41: initializing source docker://quay.io/kubevirt/builder:2510281240-82d14f4b41: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway`,
			expected:  CategoryExternal,
		},
		{
			name:      "container registry bearer token failure",
			errorText: `20:40:43: time="2026-02-18T20:40:43Z" level=warning msg="Failed, retrying in 1s ... (1/3). Error: initializing source docker://quay.io/kubevirt/builder:2510281240-82d14f4b41: Requesting bearer token: received unexpected HTTP status: 502 Bad Gateway"`,
			expected:  CategoryExternal,
		},
		{
			name:      "DNS resolution failure",
			errorText: `Error: unable to copy from source docker://quay.io/kubevirt/builder:2602201229-4494d1587: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: Get "https://quay.io/v2/auth": dial tcp: lookup quay.io: Temporary failure in name resolution`,
			expected:  CategoryExternal,
		},
		{
			name:      "git clone HTTP2 framing error",
			errorText: `17:12:38: fatal: unable to access 'https://github.com/masterzen/winrm-cli/': Error in the HTTP2 framing layer`,
			expected:  CategoryExternal,
		},
		{
			name:      "git clone 502",
			errorText: `17:08:39: fatal: unable to access 'https://github.com/masterzen/winrm-cli/': The requested URL returned error: 502`,
			expected:  CategoryExternal,
		},
		{
			name:      "git clone 500",
			errorText: `17:10:46: fatal: unable to access 'https://github.com/packer-community/winrmcp/': The requested URL returned error: 500`,
			expected:  CategoryExternal,
		},
		{
			name:      "podman container removal timeout",
			errorText: `15:45:00: Error: cannot remove container fdb28089f2bfccabdf0a36a9d248f6e2fead25b90ed175dc5805f3821801cb67 as it could not be stopped: given PID did not die within timeout`,
			expected:  CategoryExternal,
		},
		{
			name:      "podman storage file timeout",
			errorText: `18:01:07: Error: timed out waiting for file /var/lib/containers/storage/overlay-containers/5bc99eb00b10/exit/5bc99eb00b10`,
			expected:  CategoryExternal,
		},
		{
			name:      "kind cluster deletion failure",
			errorText: `15:45:00: ERROR: failed to delete cluster "vgpu": failed to delete nodes: command "podman rm -f -v vgpu-control-plane" failed with error: exit status 125`,
			expected:  CategoryExternal,
		},
		{
			name:      "kube-apiserver body decode noise",
			errorText: `00:43:10: I0318 00:43:08.163777     241 request.go:1500] "Body was not decodable (unable to check for Status)"`,
			expected:  CategoryExternal,
		},
		{
			name:      "API rate limiter timeout",
			errorText: `09:52:59:           msg: "client rate limiter Wait returned an error: context deadline exceeded",`,
			expected:  CategoryExternal,
		},
		{
			name:      "skopeo retry warning",
			errorText: `time="2026-02-18T21:10:38Z" level=warning msg="Failed, retrying in 1s ... (1/3). Error: initializing source docker://quay.io/kubevirt/builder:2510281240-82d14f4b41: pinging container registry quay.io: received unexpected HTTP status: 504 Gateway Time-out"`,
			expected:  CategoryExternal,
		},
		{
			name:      "winrmcli build depends on failed git clone",
			errorText: `ERROR: /root/go/src/kubevirt.io/kubevirt/images/winrmcli/BUILD.bazel:7:8: //images/winrmcli:winrm-cli-tar depends on @com_github_masterzen_winrmcli//:winrm-cli in repository @com_github_masterzen_winrmcli which failed to fetch. no such package '@com_github_masterzen_winrmcli//'`,
			expected:  CategoryExternal,
		},

		// === PR Build patterns ===
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

		// === Internal patterns ===
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
			name:      "make bazel-build-images failure",
			errorText: `make: *** [Makefile:37: bazel-build-images] Error 125`,
			expected:  CategoryInternal,
		},
		{
			name:      "make bazel-build-functests failure",
			errorText: `make: *** [Makefile:26: bazel-build-functests] Error 125`,
			expected:  CategoryInternal,
		},
		{
			name:      "kubevirt deployment timeout",
			errorText: `10:41:43: error: timed out waiting for the condition on kubevirts/kubevirt`,
			expected:  CategoryInternal,
		},
		{
			name:      "CNAO deployment timeout",
			errorText: `error: timed out waiting for the condition on deployments/cluster-network-addons-operator`,
			expected:  CategoryInternal,
		},
		{
			name:      "JWS signature not found",
			errorText: `02:52:21: I0319 02:52:13.772217     238 token.go:249] [discovery] Retrying due to error: could not find a JWS signature in the cluster-info ConfigMap for token ID "abcdef"`,
			expected:  CategoryInternal,
		},

		// === External: bazel remote cache ===
		{
			name:      "bazel remote cache blob fetch failure",
			errorText: `07:53:32: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666`,
			expected:  CategoryExternal,
		},
		{
			name:      "bazel remote cache blob fetch bulk transfer",
			errorText: `18:37:58: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:`,
			expected:  CategoryExternal,
		},
		{
			name:      "bazel remote cache connection timeout",
			errorText: `09:00:52: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080`,
			expected:  CategoryExternal,
		},
		{
			name:      "bazel remote cache download timeout",
			errorText: `08:54:51: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-exportproxy/BUILD.bazel:27:10: GoLink cmd/virt-exportproxy/virt-exportproxy_/virt-exportproxy failed: Exec failed due to IOException: Download of '/kubevirt/kubevirt,...' timed out. Received 21508688 of 24245460 bytes.`,
			expected:  CategoryExternal,
		},
		{
			name:      "bazel remote cache DNS failure",
			errorText: `08:58:56: ERROR: /root/go/src/kubevirt.io/kubevirt/pkg/virt-handler/vsock/system/BUILD.bazel:3:11 Validating nogo output failed: Exec failed due to IOException: java.net.UnknownHostException: bazel-cache.kubevirt-prow`,
			expected:  CategoryExternal,
		},

		// === PR Build: panic ===
		{
			name:      "panic in test output",
			errorText: `ERROR: Found panic in test output`,
			expected:  CategoryPRBuild,
		},

		// === Internal: cluster creation ===
		{
			name:      "control plane startup failure",
			errorText: `04:03:20: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-apiserver check failed]`,
			expected:  CategoryInternal,
		},
		{
			name:      "kind cluster creation log line not found",
			errorText: `08:54:22: ERROR: failed to create cluster: could not find a log line that matches "Reached target .*Multi-User System.*|detected cgroup v1"`,
			expected:  CategoryInternal,
		},

		// === Needs investigation ===
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
				t.Errorf("CategorizeError() = %q, want %q\nerror_text: %s", got, tt.expected, tt.errorText)
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

	t.Run("context-based reclassification: make error with download failure in context", func(t *testing.T) {
		err := &JobBuildError{
			BuildLogErrorSnippets: []*BuildLogErrorSnippet{
				{
					ErrorText: `make: *** [Makefile:37: bazel-build-images] Error 125`,
					Context:   "20:52:12: Trying to pull quay.io/kubevirt/builder:2510281240-82d14f4b41...\n20:52:18: Error: unable to copy from source docker://quay.io/kubevirt/builder:2510281240-82d14f4b41\nmake: *** [Makefile:37: bazel-build-images] Error 125",
				},
			},
		}
		if got := CategorizeJobBuildError(err); got != CategoryExternal {
			t.Errorf("CategorizeJobBuildError() = %q, want %q", got, CategoryExternal)
		}
	})

	t.Run("context-based reclassification: build failed with download error in context", func(t *testing.T) {
		err := &JobBuildError{
			BuildLogErrorSnippets: []*BuildLogErrorSnippet{
				{
					ErrorText: `07:18:28: ERROR: Build failed. Not running target`,
					Context:   "07:18:28: Error in download: java.io.IOException: Error downloading [https://github.com/bazelbuild/buildtools/releases/download/v7.3.1/buildifier-linux-amd64]\n07:18:28: ERROR: Build failed. Not running target",
				},
			},
		}
		if got := CategorizeJobBuildError(err); got != CategoryExternal {
			t.Errorf("CategorizeJobBuildError() = %q, want %q", got, CategoryExternal)
		}
	})

	t.Run("context-based reclassification: analysis failed with download error in another snippet", func(t *testing.T) {
		err := &JobBuildError{
			BuildLogErrorSnippets: []*BuildLogErrorSnippet{
				{
					ErrorText: `15:38:49: ERROR: Analysis of target '//cmd/virt-operator:virt-operator-image' failed; build aborted: Analysis failed`,
					Context:   "15:38:49: ERROR: Analysis of target '//cmd/virt-operator:virt-operator-image' failed; build aborted: Analysis failed",
				},
				{
					ErrorText: `15:38:49: ERROR: An error occurred during the fetch of repository 'oci_regctl_linux_amd64':`,
					Context:   "WARNING: Download from https://github.com/regclient/regclient/releases/download/v0.7.0/regctl-linux-amd64 failed",
				},
			},
		}
		if got := CategorizeJobBuildError(err); got != CategoryExternal {
			t.Errorf("CategorizeJobBuildError() = %q, want %q", got, CategoryExternal)
		}
	})

	t.Run("context-based reclassification: git clone failure in context", func(t *testing.T) {
		err := &JobBuildError{
			BuildLogErrorSnippets: []*BuildLogErrorSnippet{
				{
					ErrorText: `ERROR: Analysis of target '//images/winrmcli:winrmcli-image' failed; build aborted:`,
					Context:   "fatal: unable to access 'https://github.com/masterzen/winrm-cli/': Error in the HTTP2 framing layer\nERROR: Analysis of target '//images/winrmcli:winrmcli-image' failed; build aborted:",
				},
			},
		}
		if got := CategorizeJobBuildError(err); got != CategoryExternal {
			t.Errorf("CategorizeJobBuildError() = %q, want %q", got, CategoryExternal)
		}
	})

	t.Run("genuine pr-build failure stays pr-build", func(t *testing.T) {
		err := &JobBuildError{
			BuildLogErrorSnippets: []*BuildLogErrorSnippet{
				{
					ErrorText: `ERROR: Build failed. Not running target`,
					Context:   "INFO: Elapsed time: 446.112s\nINFO: 7435 processes: 269 internal, 6 local, 7160 processwrapper-sandbox.\nERROR: Build failed. Not running target",
				},
			},
		}
		if got := CategorizeJobBuildError(err); got != CategoryPRBuild {
			t.Errorf("CategorizeJobBuildError() = %q, want %q", got, CategoryPRBuild)
		}
	})

	t.Run("external error stays external", func(t *testing.T) {
		err := &JobBuildError{
			BuildLogErrorSnippets: []*BuildLogErrorSnippet{
				{
					ErrorText: `Error: unable to copy from source docker://quay.io/kubevirt/builder:2510281240-82d14f4b41: received unexpected HTTP status: 500`,
				},
			},
		}
		if got := CategorizeJobBuildError(err); got != CategoryExternal {
			t.Errorf("CategorizeJobBuildError() = %q, want %q", got, CategoryExternal)
		}
	})

	t.Run("CategoryReason is populated for direct match", func(t *testing.T) {
		err := &JobBuildError{
			BuildLogErrorSnippets: []*BuildLogErrorSnippet{
				{ErrorText: `Error: unable to copy from source docker://quay.io/kubevirt/builder:latest`},
			},
		}
		CategorizeJobBuildError(err)
		if err.CategoryReason != "container image pull failure" {
			t.Errorf("CategoryReason = %q, want %q", err.CategoryReason, "container image pull failure")
		}
	})

	t.Run("CategoryReason is populated for context reclassification", func(t *testing.T) {
		err := &JobBuildError{
			BuildLogErrorSnippets: []*BuildLogErrorSnippet{
				{
					ErrorText: `make: *** [Makefile:37: bazel-build-images] Error 125`,
					Context:   "Error: unable to copy from source docker://quay.io/kubevirt/builder:latest\nmake: *** [Makefile:37: bazel-build-images] Error 125",
				},
			},
		}
		CategorizeJobBuildError(err)
		if err.CategoryReason != "container image pull failure in context" {
			t.Errorf("CategoryReason = %q, want %q", err.CategoryReason, "container image pull failure in context")
		}
	})

	t.Run("CategoryReason is populated for secondary snippet", func(t *testing.T) {
		err := &JobBuildError{
			BuildLogErrorSnippets: []*BuildLogErrorSnippet{
				{ErrorText: `ERROR: Analysis of target '//cmd/virt-operator:virt-operator-image' failed; build aborted: Analysis failed`},
				{ErrorText: `ERROR: An error occurred during the fetch of repository 'oci_regctl_linux_amd64':`},
			},
		}
		CategorizeJobBuildError(err)
		if err.CategoryReason != "bazel repository fetch failure (from secondary snippet)" {
			t.Errorf("CategoryReason = %q, want %q", err.CategoryReason, "bazel repository fetch failure (from secondary snippet)")
		}
	})
}
