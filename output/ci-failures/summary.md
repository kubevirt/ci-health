# CI Failure Summary

This summary provides an overview of the current CI failures, categorized by type.

---

## Non-fixable Failures

### Unexpected Build Process Termination

**Quantity:** 2

**Branches:** main, release-1.5

**SIGs:** compute, storage

**Kubernetes Versions:** 1.33, 1.32

<details>
<summary>Log Snippet</summary>

```
make: *** [Makefile:174: cluster-sync] Error 125
+ ret=2
+ check_for_panics
+ set +x
+ make cluster-down
./kubevirtci/cluster-up/down.sh
17:25:18: selecting podman as container runtime
17:25:49: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.33-sig-compute is being used by the following container(s): 36ccf9183164d26dd1a7e241ecbfc1252b4a3fdfc35ec27851c57f2bd8adffb4, 60b6ddd9fd08f49f076eddd8b7161dd936f28eaadc33d82787ec726a6384d403: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
+ true
+ exit 2
+ EXIT_VALUE=2
+ set +o xtrace
Cleaning up after podman in container.
================================================================================
Cleaning up after podman
time="2026-01-08T17:25:59Z" level=warning msg="StopSignal SIGTERM failed to stop container pull-kubevirt-e2e-k8s-1.33-sig-compute-bazel-server in 10 seconds, resorting to SIGKILL"
36ccf9183164d26dd1a7e241ecbfc1252b4a3fdfc35ec27851c57f2bd8adffb4
60b6ddd9fd08f49f076eddd8b7161dd936f28eaadc33d82787ec726a6384d403
/usr/local/bin/runner.sh: line 50: wait: pid 1225 is not a child of this shell
================================================================================
Done cleaning up after podman in container.
```

</details>

<details>
<summary>Failed Job URLs</summary>

- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16484/pull-kubevirt-e2e-k8s-1.33-sig-compute/2009292795868614657
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16451/pull-kubevirt-e2e-k8s-1.32-sig-storage/2009194142705389568

</details>

<details>
<summary>Pull Request History URLs</summary>

- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16484
- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16451

</details>

### Bazel Download Failure

**Quantity:** 48

**Branches:** release-1.5

**SIGs:** compute, operator, sev, vgpu, windows, network, sriov, storage

**Kubernetes Versions:** 1.30, 1.31, 1.32

<details>
<summary>Log Snippet</summary>

```
WARNING: Download from https://storage.googleapis.com/rules_docker/aad94363e63d31d574cf701df484b3e8b868a96a/puller-darwin-amd64 failed: class java.io.IOException GET returned 403 Forbidden
ERROR: An error occurred during the fetch of repository 'go_puller_darwin':
   Traceback (most recent call last):
      File "/tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/bazel_tools/tools/build_defs/repo/http.bzl", line 168, column 33, in _http_file_impl
              download_info = ctx.download(
Error in download: java.io.IOException: Error downloading [https://storage.googleapis.com/rules_docker/aad94363e63d31d574cf701df484b3e8b868a96a/puller-darwin-amd64] to /tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/go_puller_darwin/file/downloaded: GET returned 403 Forbidden
```

</details>

<details>
<summary>Failed Job URLs</summary>

- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-kind-1.30-vgpu-1.5/1999515795889065984
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-kind-1.30-vgpu-1.5/1999505649045606400
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-kind-1.30-vgpu-1.5/1999503544763289600
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.5/1999515811789672448
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.5/1999505658533122048
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.5/1999503564707205120
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-compute-1.5/1999515805900869632
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-compute-1.5/1999505651839012864
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-compute-1.5/1999503557987930112
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-compute-1.5/1999515808757190656
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-compute-1.5/1999505655148318720
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-compute-1.5/1999503561334984704
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.5/1999515797956857856
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.5/1999505649972547584
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.5/1999503553797820416
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-serial-1.5/1999515812154576896
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-serial-1.5/1999505660189872128
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-serial-1.5/1999503566443646976
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-operator-1.5/1999515809705103360
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-operator-1.5/1999505655982985216
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-operator-1.5/1999503562178039808
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-operator-1.5/1999515811944861696
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-operator-1.5/1999505659380371456
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-operator-1.5/1999503565537677312
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-operator-1.5/1999515806081224704
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-operator-1.5/1999505652694650880
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-operator-1.5/1999503558830985216
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-ipv6-sig-network-1.5/1999515797474512896
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-ipv6-sig-network-1.5/1999505649859301376
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-ipv6-sig-network-1.5/1999503553013485568
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-network-1.5/1999515800792207360
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-network-1.5/1999505650249371648
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-network-1.5/1999503556301819904
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-network-1.5/1999515807108829184
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-network-1.5/1999505653449625600
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-network-1.5/1999503559703400448
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-network-1.5/1999515810380386304
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-network-1.5/1999505656826040320
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-network-1.5/1999503563016900608
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.5/1999515810653016064
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.5/1999505657664901120
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.5/1999503563855761408
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-storage-1.5/1999515803644334080
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-storage-1.5/1999505651033706496
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-storage-1.5/1999503557153263616
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-storage-1.5/1999515807553425408
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-storage-1.5/1999505654343012352
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-storage-1.5/1999503560496123904

</details>

<details>
<summary>Pull Request History URLs</summary>

- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15856

</details>