# CI Failure Summary

## Non-Fixable Failures

### Image Pull Failure

- **Quantity:** 5
- **Branches:** main, release-1.7
- **SIGs:** compute, operator, network, storage
- **Kubernetes Versions:** 1.32, 1.33, 1.34
- **PRs:** 16168

<details>
<summary>Error Snippet</summary>

```
Error: unable to copy from source docker://quay.io/kubevirtci/gocli:2510241540-9fd40b1c: initializing source docker://quay.io/kubevirtci/gocli:2510241540-9fd40b1c: pinging container registry quay.io: Get "https://quay.io/v2/": context deadline exceeded
```

</details>

<details>
<summary>Job URLs</summary>

- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16168/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/1991558793346945024
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16168/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.7/1991558797499305984
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16168/pull-kubevirt-e2e-k8s-1.34-ipv6-sig-network-1.7/1991558788657713152
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16168/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/1991558795054026752
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16168/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/1991558792470335488

</details>

<details>
<summary>PR History</summary>

- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16168

</details>

### Body Not Decodable

- **Quantity:** 4
- **Branches:** main, release-1.7
- **SIGs:** compute, network
- **Kubernetes Versions:** 1.33, 1.34
- **PRs:** 16216, 15401, 16157, 16168

<details>
<summary>Error Snippet</summary>

```
Body was not decodable (unable to check for Status)" err="couldn't get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct
```

</details>

<details>
<summary>Job URLs</summary>

- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16216/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/1992959273319010304
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15401/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations/1991785583269122048
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16157/pull-kubevirt-e2e-k8s-1.34-sig-compute-serial/1991879176415612928
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16168/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/1991764230839209984

</details>

<details>
<summary>PR History</summary>

- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16216
- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15401
- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16157
- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16168

</details>

### Failed to Delete Cluster Nodes

- **Quantity:** 2
- **Branches:** release-1.7
- **SIGs:** compute
- **Kubernetes Versions:** 1.33, 1.34
- **PRs:** 16216

<details>
<summary>Error Snippet</summary>

```
ERROR: failed to delete cluster "vgpu": failed to delete nodes: command "podman rm -f -v vgpu-control-plane" failed with error: exit status 125
Error: cannot remove container ... as it could not be stopped: given PID did not die within timeout
```

</details>

<details>
<summary>Job URLs</summary>

- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16216/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/1992959273319010304
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16216/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/1992959283058184192

</details>

<details>
<summary>PR History</summary>

- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16216

</details>

### Unexpected Build Process Termination

- **Quantity:** 2
- **Branches:** main
- **SIGs:** compute, operator
- **Kubernetes Versions:** 1.34
- **PRs:** 16157, 16165

<details>
<summary>Error Snippet</summary>

```
make: *** [Makefile:174: cluster-sync] Error 125
/usr/local/bin/runner.sh: line 50: wait: pid 1225 is not a child of this shell
```

</details>

<details>
<summary>Job URLs</summary>

- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16157/pull-kubevirt-e2e-k8s-1.34-sig-compute-serial/1992170388762660864
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16165/pull-kubevirt-e2e-k8s-1.34-sig-operator/1991454282557165568

</details>

<details>
<summary>PR History</summary>

- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16157
- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16165

</details>

### Kind Binary Download Failure

- **Quantity:** 1
- **Branches:** main
- **SIGs:** compute
- **Kubernetes Versions:** 1.33
- **PRs:** 16067

<details>
<summary>Error Snippet</summary>

```
/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kind: line 1: !DOCTYPE: No such file or directory
```

</details>

<details>
<summary>Job URLs</summary>

- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16067/pull-kubevirt-e2e-kind-1.33-vgpu/1992966586486493184

</details>

<details>
<summary>PR History</summary>

- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16067

</details>

## Fixable Failures

### Taint Not Found

- **Quantity:** 1
- **Branches:** release-1.7
- **SIGs:** compute
- **Kubernetes Versions:** 1.34
- **PRs:** 16216

<details>
<summary>Error Snippet</summary>

```
error: taint "node-role.kubernetes.io/master:NoSchedule" not found
```

</details>

<details>
<summary>Job URLs</summary>

- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16216/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/1992959283058184192

</details>

<details>
<summary>PR History</summary>

- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16216

</details>
