# CI Failure Summary

This report summarizes the analysis of Prow job failures that occurred before the main testing stage.

## High Volume Fixable Errors

### Invalid go.mod version

- **Branches:** main
- **SIGs:** operator, storage
- **Kubernetes:** 1.32
- **Quantity:** 2
- **Error Type:** Fixable

This error is caused by an invalid version string in the `go.mod` file, which prevents `gazelle` from reading the module paths correctly and causes the build to fail.

<details>
<summary>Log Snippet</summary>

```
gazelle: reading module paths from /root/go/src/kubevirt.io/kubevirt/go.mod: /root/go/src/kubevirt.io/kubevirt/go.mod:3: invalid go version '1.23.0': must match format 1.23
make: *** [Makefile:26: bazel-build-functests] Error 125
```

</details>

<details>
<summary>Failed Job URLs</summary>

- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15798/pull-kubevirt-e2e-k8s-1.32-sig-operator/1974106605586747392
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15799/pull-kubevirt-e2e-k8s-1.32-sig-storage/1973739563100672000

</details>

<details>
<summary>Pull Request History</summary>

- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15798
- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15799

</details>

## Other Fixable Errors

### Podman Daemon Failed to Start

- **Branch:** release-1.3
- **SIG:** compute
- **Kubernetes:** 1.30
- **Quantity:** 1
- **Error Type:** Fixable

The Podman daemon within the test container failed to become ready within the timeout period, indicating an issue with the CI environment setup.

<details>
<summary>Log Snippet</summary>

```
Reached maximum attempts, not waiting any longer...
Podman daemon failed to start successfully
```

</details>

<details>
<summary>Failed Job URL</summary>

- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15161/pull-kubevirt-e2e-k8s-1.30-sig-compute-1.3/1943497072455979008

</details>

<details>
<summary>Pull Request History</summary>

- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15161

</details>

### Istio CNI Configuration Not Found

- **Branch:** release-1.5
- **SIG:** network
- **Kubernetes:** 1.30
- **Quantity:** 1
- **Error Type:** Fixable

The `cluster-up` script failed because it could not find the Istio CNI configuration files, pointing to a problem in the environment's network setup scripts.

<details>
<summary>Log Snippet</summary>

```
cp: cannot stat '/etc/cni/multus/net.d/*istio*.conf': No such file or directory
make: *** [Makefile:163: cluster-up] Error 1
```

</details>

<details>
<summary>Failed Job URL</summary>

- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15801/pull-kubevirt-e2e-k8s-1.30-sig-network-1.5/1973720651801825280

</details>

<details>
<summary>Pull Request History</summary>

- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15801

</details>

### Cluster Sync and Cleanup Failure

- **Branch:** main
- **SIG:** network
- **Kubernetes:** 1.33
- **Quantity:** 1
- **Error Type:** Fixable

The `make cluster-sync` command failed with a generic error, which then prevented the `cluster-down` cleanup step from running correctly because container volumes were still in use.

<details>
<summary>Log Snippet</summary>

```
make: *** [Makefile:177: cluster-sync] Error 125
...
Error response from daemon: volume pull-kubevirt-e2e-k8s-1.33-ipv6-sig-network is being used by the following container(s)
make: *** [Makefile:165: cluster-down] Error 1
```

</details>

<details>
<summary>Failed Job URL</summary>

- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15798/pull-kubevirt-e2e-k8s-1.33-ipv6-sig-network/1974106604953407488

</details>

<details>
<summary>Pull Request History</summary>

- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15798

</details>

## Non-Fixable Errors

### External Network Connectivity Issue

- **Branch:** main
- **SIG:** compute
- **Kubernetes:** 1.34
- **Quantity:** 1
- **Error Type:** Non-Fixable

The build failed because it could not resolve `github.com` to download a build dependency. This is an external network issue from within the Prow job environment.

<details>
<summary>Log Snippet</summary>

```
WARNING: Download from https://github.com/brianmcarey/bazeldnf/releases/download/v0.5.9-2/bazeldnf-v0.5.9-2-linux-amd64 failed: class com.google.devtools.build.lib.bazel.repository.downloader.UnrecoverableHttpException Unknown host: github.com
```

</details>

<details>
<summary>Failed Job URL</summary>

- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15828/pull-kubevirt-e2e-kind-1.34-sev/1975204012558913536

</details>

<details>
<summary>Pull Request History</summary>

- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15828

</details>