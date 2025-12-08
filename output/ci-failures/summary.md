# CI Failure Analysis Summary

This report summarizes the analysis of CI job failures that did not reach the testing stage.

## Fixable Errors

### Invalid Go Version in `go.mod`

*   **Quantity Total:** 1
*   **Branch Name(s):** main
*   **SIG(s):** compute
*   **Kubernetes Version(s):** 1.32
*   **Description:** The `go.mod` file specifies a Go version (e.g., `1.24.0`) that is incompatible with the version of `gazelle` used in the build environment, which expects a format like `1.23`. This causes the `gazelle` command to fail, which in turn causes the Bazel build to fail.

<details>
<summary>Example Log Snippet</summary>

```
13:21:29: gazelle: reading module paths from /root/go/src/kubevirt.io/kubevirt/go.mod: /root/go/src/kubevirt.io/kubevirt/go.mod:3: invalid go version '1.24.0': must match format 1.23
13:21:29: /root/go/src/kubevirt.io/kubevirt/go.mod:206: unknown block type: tool
```

</details>

<details>
<summary>Failed Build Job URLs</summary>

*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15821/pull-kubevirt-e2e-k8s-1.32-sig-compute/1995838972147798016

</details>

<details>
<summary>Pull Request History URLs</summary>

*   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15821

</details>

### Taint Not Found

*   **Quantity Total:** 1
*   **Branch Name(s):** main
*   **SIG(s):** compute
*   **Kubernetes Version(s):** 1.33
*   **Description:** The cluster setup script attempts to remove a taint from a node that does not exist. This is a bug in the CI scripts.

<details>
<summary>Example Log Snippet</summary>

```
06:59:17: error: taint "node-role.kubernetes.io/master:NoSchedule" not found
06:59:17: error: taint "node-role.kubernetes.io/control-plane:NoSchedule" not found
```

</details>

<details>
<summary>Failed Build Job URLs</summary>

*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16131/pull-kubevirt-e2e-kind-1.33-vgpu/1993935254829666304

</details>

<details>
<summary>Pull Request History URLs</summary>

*   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16131

</details>

## Non-Fixable Errors

### Body Not Decodable

*   **Quantity Total:** 2
*   **Branch Name(s):** main
*   **SIG(s):** compute
*   **Kubernetes Version(s):** 1.34
*   **Description:** The `kube-apiserver` is unable to decode a request body. This is likely a transient network issue or a bug in the client sending the request.

<details>
<summary>Example Log Snippet</summary>

```
10:26:38: I1203 05:26:38.147896    1601 request.go:1500] "Body was not decodable (unable to check for Status)" err="couldn't get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \"json:\\\"apiVersion,omitempty\\\""; Kind string \"json:\\\"kind,omitempty\\\" }"
```

</details>

<details>
<summary>Failed Build Job URLs</summary>

*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16253/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations/1996160133175971840
*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15868/pull-kubevirt-e2e-k8s-1.34-sig-compute-serial/1995532685656723456

</details>

<details>
<summary>Pull Request History URLs</summary>

*   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16253
*   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15868

</details>

### Unexpected Build Process Termination

*   **Quantity Total:** 2
*   **Branch Name(s):** main, release-1.5
*   **SIG(s):** compute, network
*   **Kubernetes Version(s):** 1.34, 1.31
*   **Description:** The build process was terminated unexpectedly, resulting in a generic `make Error 125` or other errors indicating a sudden stop. This is likely a Prow infrastructure issue.

<details>
<summary>Example Log Snippet</summary>

```
make: *** [Makefile:174: cluster-sync] Error 125
+ ret=2
+ check_for_panics
+ set +x
+ make cluster-down
./kubevirtci/cluster-up/down.sh
16:56:01: selecting podman as container runtime
16:56:57: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations is being used by the following container(s): 50b8a0948caa46015eeae689b925b42fae8b1d6e4e81145be63d08115590be4e, 7902b42406fd3a8b5791833ebd0dd71b3edd44e26e03c0917d1bdf7d26a7a684: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
```

</details>

<details>
<summary>Failed Build Job URLs</summary>

*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15868/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations/1995532679772114944
*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16237/pull-kubevirt-e2e-k8s-1.31-sig-network-1.5/1993747414808268800

</details>

<details>
<summary>Pull Request History URLs</summary>

*   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15868
*   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16237

</details>

### Content Length Mismatch

*   **Quantity Total:** 1
*   **Branch Name(s):** main
*   **SIG(s):** compute
*   **Kubernetes Version(s):** 1.33
*   **Description:** Error downloading a container image from `quay.io` due to a content length mismatch. This is likely a transient network or registry issue.

<details>
<summary>Example Log Snippet</summary>

```
06:52:39: WARNING: Download from https://quay.io/v2/kubevirtci/fedora-with-test-tooling/blobs/sha256:6b529a079d4da130025f4e7b899f7203685411e5cb722a9ea0debbb953208d0f failed: class com.google.devtools.build.lib.bazel.repository.downloader.ContentLengthMismatchException Bytes read 908066816 but wanted 972172954
```

</details>

<details>
<summary>Failed Build Job URLs</summary>

*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16131/pull-kubevirt-e2e-kind-1.33-vgpu/1993935254829666304

</details>

<details>
<summary>Pull Request History URLs</summary>

*   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16131

</details>

### Host Audio Device Issue

*   **Quantity Total:** 1
*   **Branch Name(s):** release-1.5
*   **SIG(s):** network
*   **Kubernetes Version(s):** 1.31
*   **Description:** The job fails due to errors related to the host's audio devices (SDL/ALSA). This is an infrastructure issue.

<details>
<summary>Example Log Snippet</summary>

```
sdl: Reason: ALSA: Couldn't open audio device: Host is down
audio: Could not create a backend for voice `ac97.pi'
```

</details>

<details>
<summary>Failed Build Job URLs</summary>

*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16237/pull-kubevirt-e2e-k8s-1.31-sig-network-1.5/1993747414808268800

</details>

<details>
<summary>Pull Request History URLs</summary>

*   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16237

</details>
