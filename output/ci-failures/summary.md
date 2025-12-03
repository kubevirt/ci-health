# CI Failure Summary

This report summarizes the analysis of CI job failures.

## Fixable Errors

### Invalid Go Version in go.mod

*   **Quantity:** 1
*   **Branches:** main
*   **SIGs:** compute
*   **Kubernetes Versions:** 1.32
*   **Description:** The `go.mod` file specifies a go version that is not compatible with the gazelle version used in the build.
*   <details>
    <summary>Log Snippet</summary>

    ```
    gazelle: reading module paths from /root/go/src/kubevirt.io/kubevirt/go.mod: /root/go/src/kubevirt.io/kubevirt/go.mod:3: invalid go version '1.24.0': must match format 1.23
    /root/go/src/kubevirt.io/kubevirt/go.mod:204: unknown block type: tool
    ```
    </details>
*   <details>
    <summary>Failed Job URLs</summary>

    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15863/pull-kubevirt-e2e-k8s-1.32-sig-compute/1995192045848760320
    </details>
*   <details>
    <summary>Pull Request History</summary>

    *   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15863
    </details>

### Taint Not Found

*   **Quantity:** 1
*   **Branches:** main
*   **SIGs:** compute
*   **Kubernetes Versions:** 1.33
*   **Description:** The cluster setup script attempts to remove a taint that does not exist.
*   <details>
    <summary>Log Snippet</summary>

    ```
    error: taint "node-role.kubernetes.io/master:NoSchedule" not found
    error: taint "node-role.kubernetes.io/control-plane:NoSchedule" not found
    ```
    </details>
*   <details>
    <summary>Failed Job URLs</summary>

    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16131/pull-kubevirt-e2e-kind-1.33-vgpu/1993935254829666304
    </details>
*   <details>
    <summary>Pull Request History</summary>

    *   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16131
    </details>

## Non-Fixable Errors

### ContentLengthMismatchException

*   **Quantity:** 1
*   **Branches:** main
*   **SIGs:** compute
*   **Kubernetes Versions:** 1.33
*   **Description:** Error downloading `fedora_with_test_tooling_single` from quay.io. This is likely a network issue or a problem with the registry.
*   <details>
    <summary>Log Snippet</summary>

    ```
    WARNING: Download from https://quay.io/v2/kubevirtci/fedora-with-test-tooling/blobs/sha256:6b529a079d4da130025f4e7b899f7203685411e5cb722a9ea0debbb953208d0f failed: class com.google.devtools.build.lib.bazel.repository.downloader.ContentLengthMismatchException Bytes read 908066816 but wanted 972172954
    ```
    </details>
*   <details>
    <summary>Failed Job URLs</summary>

    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16131/pull-kubevirt-e2e-kind-1.33-vgpu/1993935254829666304
    </details>
*   <details>
    <summary>Pull Request History</summary>

    *   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16131
    </details>

### Body Not Decodable

*   **Quantity:** 2
*   **Branches:** main
*   **SIGs:** compute, network
*   **Kubernetes Versions:** 1.34, 1.33
*   **Description:** The `kube-apiserver` is unable to decode a request body. This could be a transient network issue, or it could be a bug in the client sending the request.
*   <details>
    <summary>Log Snippet</summary>

    ```
    I1130 13:11:31.124825    1593 request.go:1664] "Body was not decodable (unable to check for Status)" err="couldn't get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \"json:\"apiVersion,omitempty\""; Kind string \"json:\"kind,omitempty\"" }
    ```
    </details>
*   <details>
    <summary>Failed Job URLs</summary>

    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15868/pull-kubevirt-e2e-k8s-1.34-sig-compute-serial/1995532685656723456
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15863/pull-kubevirt-e2e-k8s-1.33-sig-network/1995192046280773632
    </details>
*   <details>
    <summary>Pull Request History</summary>

    *   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15868
    *   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15863
    </details>
