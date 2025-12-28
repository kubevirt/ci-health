# CI Failure Analysis Summary

This report summarizes the analysis of CI job failures.

## Fixable Failures

### Taint Not Found (1)

-   **Quantity:** 1
-   **Branch:** main
-   **SIG:** compute
-   **K8s Version:** 1.33
-   **PR:** [16363](https://github.com/kubevirt/kubevirt/pull/16363)

<details>
<summary>Log Snippet</summary>

```
11:26:48: ++ for node in ${master_nodes[@]}
11:26:48: ++ _kubectl taint nodes vgpu-control-plane node-role.kubernetes.io/master:NoSchedule-
11:26:48: ++ /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubeconfig taint nodes vgpu-control-plane node-role.kubernetes.io/master:NoSchedule-
11:26:48: error: taint "node-role.kubernetes.io/master:NoSchedule" not found
11:26:48: ++ _kubectl taint nodes vgpu-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
11:26:48: ++ /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubeconfig taint nodes vgpu-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
11:26:48: error: taint "node-role.kubernetes.io/control-plane:NoSchedule" not found
```

</details>

<details>
<summary>Failed Job URLs</summary>

- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16363/pull-kubevirt-e2e-kind-1.33-vgpu/2001249228277420032

</details>

<details>
<summary>PR History URLs</summary>

- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16363

</details>

### Single Stack Cluster Expected (1)

-   **Quantity:** 1
-   **Branch:** release-1.5
-   **SIG:** network
-   **K8s Version:** 1.32
-   **PR:** [16357](https://github.com/kubevirt/kubevirt/pull/16357)

<details>
<summary>Log Snippet</summary>

```
15:20:29: ++ local primary_ip=fd00:10:244::c482
15:20:29: ++ [[ ! fd00:10:244::c482 =~ fd00 ]]
15:20:29: ++ _kubectl get pod -n kube-system calico-kube-controllers-7df54458b-m5n2l "-ojsonpath={ @.status.podIPs[1] }"
15:20:30: ++ echo 'error: single stack cluster expected'
15:20:30: error: single stack cluster expected
15:20:30: ++ exit 1
```

</details>

<details>
<summary>Failed Job URLs</summary>

- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16357/pull-kubevirt-e2e-k8s-1.32-ipv6-sig-network-1.5/2000922274189807616

</details>

<details>
<summary>PR History URLs</summary>

- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16357

</details>

## Non-Fixable Failures

### Image Pull Timeout (1)

-   **Quantity:** 1
-   **Branch:** release-1.5
-   **SIG:** compute
-   **K8s Version:** 1.30
-   **PR:** [15990](https://github.com/kubevirt/kubevirt/pull/15990)

<details>
<summary>Log Snippet</summary>

```
17:49:40: Error in fail: Pull command failed: Timed out (/tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/go_puller_linux_amd64/file/downloaded -directory /tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/fedora_realtime/image -os linux -os-version  -os-features  -architecture amd64 -variant  -features  -name quay.io/kubevirt/fedora-realtime-container-disk@sha256:437f4e02986daf0058239f4a282d32304dcac629d5d1b4c75a74025f1ce22811)
```
</details>

<details>
<summary>Failed Job URLs</summary>

- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15990/pull-kubevirt-e2e-kind-1.30-vgpu-1.5/1983948169213382656

</details>

<details>
<summary>PR History URLs</summary>

- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15990

</details>

### Body Not Decodable (1)

-   **Quantity:** 1
-   **Branch:** main
-   **SIG:** compute
-   **K8s Version:** 1.34
-   **PR:** [16343](https://github.com/kubevirt/kubevirt/pull/16343)

<details>
<summary>Log Snippet</summary>

```
12:21:38: I1216 07:21:38.802623    1606 request.go:1500] "Body was not decodable (unable to check for Status)" err="couldn't get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \"json:\"apiVersion,omitempty\""; Kind string "json:\"kind,omitempty\"" }"
```
</details>

<details>
<summary>Failed Job URLs</summary>

- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16343/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations/2000900961081495552

</details>

<details>
<summary>PR History URLs</summary>

- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16343

</details>

### Unexpected Build Process Termination (1)

-   **Quantity:** 1
-   **Branch:** main
-   **SIG:** network
-   **K8s Version:** 1.34
-   **PR:** [16340](https://github.com/kubevirt/kubevirt/pull/16340)

<details>
<summary>Log Snippet</summary>

```
10:13:53: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-ipv6-sig-network is being used by the following container(s): aeec17beda21f9ebad49b29efce2a1144ba0bd461ca4e9705cb998119fc5982f, 131c2d56383a38252963e2b32222069945f6a009f067c636c4f06f0c662c3a4c: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
...
/usr/local/bin/runner.sh: line 50: wait: pid 1225 is not a child of this shell
```
</details>

<details>
<summary>Failed Job URLs</summary>

- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16340/pull-kubevirt-e2e-k8s-1.34-ipv6-sig-network/2000504107663626240

</details>

<details>
<summary>PR History URLs</summary>

- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16340

</details>
