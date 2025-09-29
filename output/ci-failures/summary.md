# CI Failure Summary

This report summarizes the analysis of Prow job failures that occurred before the testing stage.

## High-Volume Fixable Errors

### Bazel Build Failure: Missing Dependency

A recurring build failure was identified across multiple jobs, caused by a missing strict dependency in the Bazel configuration.

*   **Branch:** `release-1.0`
*   **SIG(s):** `sig-network`, `sig-compute`, `sig-operator`, `sig-storage`
*   **Kubernetes Version(s):** `1.25`, `1.26`, `1.27`
*   **Pull Request:** [#15339](https://github.com/kubevirt/kubevirt/pull/15339)

<details>
<summary>Error Log Snippet</summary>

```
ERROR: /tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/org_golang_google_grpc/internal/syscall/BUILD.bazel:3:11: GoCompilePkg external/org_golang_google_grpc/internal/syscall/syscall.a failed: (Exit 1): builder failed: error executing command bazel-out/k8-opt-exec-2B5CBBC6/bin/external/go_sdk/builder compilepkg -sdk external/go_sdk -installsuffix linux_amd64 -tags selinux,selinux -src ... (remaining 27 arguments skipped)

Use --sandbox_debug to see verbose messages from the sandbox and retain the sandbox build root for debugging
compilepkg: missing strict dependencies:
	/tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/sandbox/linux-sandbox/305/execroot/kubevirt/external/org_golang_google_grpc/internal/syscall/syscall_linux.go: import of "golang.org/x/sys/unix"
No dependencies were provided.
Check that imports in Go sources match importpath attributes in deps.
```

</details>

<details>
<summary>Affected Job URLs</summary>

*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15339/pull-kubevirt-e2e-kind-1.27-sriov-1.0/1950863204456337408
*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15339/pull-kubevirt-e2e-kind-1.25-vgpu-1.0/1950863204104015872
*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15339/pull-kubevirt-e2e-k8s-1.25-sig-compute-1.0/1950863214392643584
*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15339/pull-kubevirt-e2e-k8s-1.26-sig-compute-1.0/1950863217722920960
*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15339/pull-kubevirt-e2e-k8s-1.27-sig-compute-1.0/1950863221078364160
*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15339/pull-kubevirt-e2e-k8s-1.25-sig-compute-migrations-1.0/1950863210231894016
*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15339/pull-kubevirt-e2e-k8s-1.25-ipv6-sig-network-1.0/1950863209439170560
*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15339/pull-kubevirt-e2e-k8s-1.25-sig-network-1.0/1950863212786225152
*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15339/pull-kubevirt-e2e-k8s-1.27-sig-network-1.0/1950863219438391296
*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15339/pull-kubevirt-e2e-k8s-1.26-sig-network-1.0/1950863216133279744
*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15339/pull-kubevirt-e2e-k8s-1.27-sig-operator-1.0/1950863221975945216
*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15339/pull-kubevirt-e2e-k8s-1.26-sig-operator-1.0/1950863218649862144
*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15339/pull-kubevirt-e2e-k8s-1.25-sig-operator-1.0/1950863215311196160
*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15339/pull-kubevirt-e2e-k8s-1.25-sig-performance-1.0/1950863203684585472
*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15339/pull-kubevirt-e2e-k8s-1.25-sig-storage-1.0/1961311882581118976
*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15339/pull-kubevirt-e2e-k8s-1.25-sig-storage-1.0/1950863213591531520
*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15339/pull-kubevirt-e2e-k8s-1.26-sig-storage-1.0/1950863217005694976
*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15339/pull-kubevirt-e2e-k8s-1.27-sig-storage-1.0/1950863220302417920

</details>

<details>
<summary>Pull Request History</summary>

*   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15339

</details>

## Non-Fixable Errors

During the analysis, warnings related to failing downloads from `http://mirror.stream.centos.org` were observed. These appear to be external service issues and are not the primary cause of the job failures.
