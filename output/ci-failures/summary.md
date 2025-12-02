# CI Failure Summary

## Fixable Errors

### Kubectl taint not found (2 occurrences)
- **Branches**: main, release-1.7
- **SIGs**: sig-compute
- **Kubernetes Versions**: 1.33, 1.34
<details><summary>Example Log Snippet</summary>

```
error: taint "node-role.kubernetes.io/master:NoSchedule" not found
```
</details>
<details><summary>Failed Build Job URLs</summary>

- [pull-kubevirt-e2e-kind-1.33-vgpu/1993935254829666304](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16131/pull-kubevirt-e2e-kind-1.33-vgpu/1993935254829666304)
- [pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/1992959283058184192](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16216/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/1992959283058184192)
</details>
<details><summary>Pull Request History URLs</summary>

- [PR 16131 History](https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16131)
- [PR 16216 History](https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16216)
</details>

## Non-Fixable Errors

### Kube-apiserver Body not decodable / Internal Server Error (2 occurrences)
- **Branches**: main
- **SIGs**: sig-network
- **Kubernetes Versions**: 1.33
<details><summary>Example Log Snippet</summary>

```
"Body was not decodable (unable to check for Status)" err="couldn't get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \"json:\"apiVersion,omitempty\""; Kind string \"json:\"kind,omitempty\" }"
```
</details>
<details><summary>Failed Build Job URLs</summary>

- [pull-kubevirt-e2e-kind-1.33-vgpu/1993935254829666304](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16131/pull-kubevirt-e2e-kind-1.33-vgpu/1993935254829666304)
- [pull-kubevirt-e2e-k8s-1.33-sig-network/1995192046280773632](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15863/pull-kubevirt-e2e-k8s-1.33-sig-network/1995192046280773632)
</details>
<details><summary>Pull Request History URLs</summary>

- [PR 15863 History](https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15863)
- [PR 16131 History](https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16131)
</details>

### Podman failed to delete cluster nodes (PID did not die) (2 occurrences)
- **Branches**: release-1.7
- **SIGs**: sig-compute
- **Kubernetes Versions**: 1.33, 1.34
<details><summary>Example Log Snippet</summary>

```
ERROR: failed to delete cluster ...: failed to delete nodes: command "podman rm -f -v ..." failed with error: exit status 125 ... given PID did not die within timeout
```
</details>
<details><summary>Failed Build Job URLs</summary>

- [pull-kubevirt-e2e-kind-1.33-vgpu-1.7/1992959273319010304](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16216/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/1992959273319010304)
- [pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/1992959283058184192](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16216/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/1992959283058184192)
</details>
<details><summary>Pull Request History URLs</summary>

- [PR 16216 History](https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16216)
</details>

### Bazel ContentLengthMismatchException during image download (1 occurrences)
- **Branches**: main
- **SIGs**: sig-compute
- **Kubernetes Versions**: 1.33
<details><summary>Example Log Snippet</summary>

```
WARNING: Download from ... failed: ...ContentLengthMismatchException Bytes read ... but wanted ...
```
</details>
<details><summary>Failed Build Job URLs</summary>

- [pull-kubevirt-e2e-kind-1.33-vgpu/1993935254829666304](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16131/pull-kubevirt-e2e-kind-1.33-vgpu/1993935254829666304)
</details>
<details><summary>Pull Request History URLs</summary>

- [PR 16131 History](https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16131)
</details>