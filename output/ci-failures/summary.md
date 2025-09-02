# CI Failure Analysis Summary

This report summarizes the analysis of Prow job failures that did not reach the testing stage.

## Fixable Errors

### High Volume

**Missing `gofuzz` Dependency**

*   **Description:** A recurring Bazel build failure indicates that the `github.com/google/gofuzz` dependency is missing from the `vendor` directory or is not correctly referenced in the Bazel build files. This is a fixable issue that requires updating the project's dependencies.
*   **SIG(s):** compute, operator, network, storage
*   **Branch(es):** main, release-1.7
*   **Kubernetes Version(s):** 1.32, 1.33, 1.34
*   <details>
    <summary>Example Log Snippet</summary>

    ```
    ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/k8s.io/kube-openapi/pkg/spec3/BUILD.bazel:3:11: no such package 'vendor/github.com/google/gofuzz': BUILD file not found in any of the following directories. Add a BUILD file to a directory to mark it as a package.
     - /root/go/src/kubevirt.io/kubevirt/vendor/github.com/google/gofuzz and referenced by '//vendor/k8s.io/kube-openapi/pkg/spec3:go_default_library'
    ```
    </details>
*   <details>
    <summary>Failed Job URLs</summary>

    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16034/pull-kubevirt-e2e-kind-1.34-sev-1.7/1985841454768984064
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16034/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/1985841454672515072
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16034/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/1985841454529908736
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16034/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.7/1985841471093215232
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16034/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/1985841467754549248
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16034/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.7/1985841474436075520
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16034/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm6arm64-1.7/1985841464365551616
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16034/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations-1.7/1985841463665102848
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16034/pull-kubevirt-e2e-k8s-1.34-sig-compute-serial-1.7/1985841476172517376
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16034/pull-kubevirt-e2e-k8s-1.34-ipv6-sig-network-1.7/1985841462717190144
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16034/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/1985841472758353920
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16034/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/1985841469423882240
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16034/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/1985841466043273216
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16034/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.7/1985841475325267968
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16034/pull-kubevirt-e2e-k8s-1.32-sig-operator-1.7/1985841468694073344
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16034/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.7/1985841471932076032
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16034/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/1985841473861455872
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16034/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/1985841466890522624
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16034/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/1985841470321463296
    </details>
*   <details>
    <summary>Pull Request History URLs</summary>

    *   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16034
    *   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15999
    *   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15945
    *   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15409
    *   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15166
    *   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16027
    *   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16042
    *   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16005
    </details>

### Low Volume

**Taint Not Found**

*   **Description:** The cluster setup script attempts to remove a taint that does not exist. This is a fixable issue in the CI scripts.
*   **SIG(s):** compute
*   **Branch(es):** release-1.7
*   **Kubernetes Version(s):** 1.34
*   <details>
    <summary>Example Log Snippet</summary>

    ```
    error: taint "node-role.kubernetes.io/master:NoSchedule" not found
    ```
    </details>
*   <details>
    <summary>Failed Job URLs</summary>

    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15999/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/1984259723330850816
    </details>
*   <details>
    <summary>Pull Request History URLs</summary>

    *   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15999
    </details>

## Non-Fixable Errors

### Low Volume

**Failed to Delete Cluster Nodes**

*   **Description:** The `podman rm` command fails to remove the cluster nodes during teardown. The error message `given PID did not die within timeout` suggests that the container processes are not terminating gracefully. This is likely a non-fixable infrastructure issue.
*   **SIG(s):** compute
*   **Branch(es):** release-1.7
*   **Kubernetes Version(s):** 1.34
*   <details>
    <summary>Example Log Snippet</summary>

    ```
    ERROR: failed to delete cluster "kind-1.34": failed to delete nodes: command "podman rm -f -v kind-1.34-control-plane" failed with error: exit status 125
    ```
    </details>
*   <details>
    <summary>Failed Job URLs</summary>

    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15999/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/1984259723330850816
    </details>
*   <details>
    <summary>Pull Request History URLs</summary>

    *   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15999
    </details>

## Undetermined Errors

### Low Volume

**Body Not Decodable**

*   **Description:** The `kube-apiserver` is unable to decode a request body. This could be a transient network issue, or it could be a bug in the client sending the request. Further investigation is needed to determine the root cause.
*   **SIG(s):** compute, network, storage
*   **Branch(es):** main, release-1.7
*   **Kubernetes Version(s):** 1.34
*   <details>
    <summary>Example Log Snippet</summary>

    ```
    I1105 07:28:06.751387    1590 request.go:1500] "Body was not decodable (unable to check for Status)" err="couldn't get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \"json:\\\"apiVersion,omitempty\\\""; Kind string \"json:\\\"kind,omitempty\\\" }"
    ```
    </details>
*   <details>
    <summary>Failed Job URLs</summary>

    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16034/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/1986044017342681088
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16005/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations-1.7/1984910182047551488
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15988/pull-kubevirt-e2e-k8s-1.34-ipv6-sig-network/1986086031421607936
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15166/pull-kubevirt-e2e-k8s-1.34-ipv6-sig-network/1986087349003489280
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16027/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/1985643090706173952
    *   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16027/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/1985643086948077568
    </details>
*   <details>
    <summary>Pull Request History URLs</summary>

    *   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16034
    *   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16005
    *   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15988
    *   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15166
    *   https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16027
    </details>

## Jobs with No "ERROR" Keyword in Logs

*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15945/pull-kubevirt-e2e-k8s-1.34-sig-compute/1985287286975107072
*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16042/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations-1.7/1986185236496519168
*   https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15409/pull-kubevirt-e2e-k8s-1.34-sig-operator/1985957067449438208