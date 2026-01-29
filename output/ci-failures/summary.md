# CI Failure Summary

This report summarizes the root causes of CI job failures.

## Non-Fixable Failures

### High Volume: Dependency Download Failures

- **Quantity:** 57
- **Branches:** release-1.5, release-1.6, release-1.7, main
- **SIGs:** compute, network, operator, storage, performance
- **Kubernetes Versions:** 1.30, 1.31, 1.32, 1.33, 1.34

<details>
<summary>Error Snippet</summary>

```
16:26:45: WARNING: Download from https://storage.googleapis.com/rules_docker/aad94363e63d31d574cf701df484b3e8b868a96a/puller-darwin-amd64 failed: class java.io.IOException GET returned 403 Forbidden
16:26:45: ERROR: An error occurred during the fetch of repository 'go_puller_darwin':
16:26:45:    Traceback (most recent call last):
16:26:45:       File "/tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/bazel_tools/tools/build_defs/repo/http.bzl", line 168, column 33, in _http_file_impl
16:26:45:               download_info = ctx.download(
16:26:45: Error in download: java.io.IOException: Error downloading [https://storage.googleapis.com/rules_docker/aad94363e63d31d574cf701df484b3e8b868a96a/puller-darwin-amd64] to /tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/go_puller_darwin/file/downloaded: GET returned 403 Forbidden
```
</details>

<details>
<summary>Job URLs</summary>

https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-kind-1.30-vgpu-1.5/1999515795889065984
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-kind-1.30-vgpu-1.5/1999505649045606400
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-kind-1.30-vgpu-1.5/1999503544763289600
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16484/pull-kubevirt-e2e-k8s-1.33-sig-compute/2009292795868614657
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-compute-1.5/1999515808757190656
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-compute-1.5/1999505655148318720
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-compute-1.5/1999503561334984704
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.5/1999515811789672448
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.5/1999505658533122048
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.5/1999503564707205120
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-compute-1.5/1999515805900869632
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-compute-1.5/1999505651839012864
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-compute-1.5/1999503557987930112
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.5/1999515797956857856
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.5/1999505649972547584
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.5/1999503553797820416
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-serial-1.5/1999515812154576896
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-serial-1.5/1999505660189872128
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-serial-1.5/1999503566443646976
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-ipv6-sig-network-1.5/1999515797474512896
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-ipv6-sig-network-1.5/1999505649859301376
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-ipv6-sig-network-1.5/1999503553013485568
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-network-1.5/1999515810380386304
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-network-1.5/1999505656826040320
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-network-1.5/1999503563016900608
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-network-1.5/1999515800792207360
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-network-1.5/1999505650249371648
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-network-1.5/1999503556301819904
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-network-1.5/1999515807108829184
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-network-1.5/1999505653449625600
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-network-1.5/1999503559703400448
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-operator-1.5/1999515809705103360
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-operator-1.5/1999505655982985216
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-operator-1.5/1999503562178039808
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-operator-1.5/1999515811944861696
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-operator-1.5/1999505659380371456
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-operator-1.5/1999503565537677312
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-operator-1.5/1999515806081224704
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-operator-1.5/1999505652694650880
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-operator-1.5/1999503558830985216
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16504/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.7/2010328086058373120
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16465/pull-kubevirt-e2e-k8s-1.31-sig-performance-1.5/2008229041038954496
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-performance-1.5/1999515795620630528
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-performance-1.5/1999505648546484224
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-performance-1.5/1999503544612294656
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16470/pull-kubevirt-e2e-k8s-1.31-sig-performance-1.6/2008490702794657792
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16451/pull-kubevirt-e2e-k8s-1.32-sig-storage/2009194142705389568
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-storage-1.5/1999515807553425408
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-storage-1.5/1999505654343012352
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-storage-1.5/1999503560496123904
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.5/1999515810653016064
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.5/1999505657664901120
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.5/1999503563855761408
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-storage-1.5/1999515803644334080
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-storage-1.5/1999505651033706496
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-storage-1.5/1999503557153263616

</details>

<details>
<summary>PR History URLs</summary>

https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15856
https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16484
https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16504
https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16465
https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16470
</details>

### Unexpected Build Process Termination

- **Quantity:** 1
- **Branches:** main
- **SIGs:** storage
- **Kubernetes Versions:** 1.32

<details>
<summary>Error Snippet</summary>

```
09:40:21: + rm -f /tmp/kubevirt.deploy.Umse
make: *** [Makefile:174: cluster-sync] Error 125
+ ret=2
+ check_for_panics
+ set +x
+ make cluster-down
./kubevirtci/cluster-up/down.sh
09:40:25: selecting podman as container runtime
09:43:25: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.32-sig-storage is being used by the following container(s): b30aadb7fbb8ad36b07ab60df7737c6fe8e9aa201ba76a8a54021ab7321ef9c4, 2b2c44c5e5d7fc07c3701d170bcc379a9eb35cae768b9649c7f9eef95bbead3e: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
+ true
+ exit 2
+ EXIT_VALUE=2
+ set +o xtrace
Cleaning up after podman in container.
================================================================================
Cleaning up after podman
time="2026-01-08T09:43:35Z" level=warning msg="StopSignal SIGTERM failed to stop container pull-kubevirt-e2e-k8s-1.32-sig-storage-bazel-server in 10 seconds, resorting to SIGKILL"
2b2c44c5e5d7fc07c3701d170bcc379a9eb35cae768b9649c7f9eef95bbead3e
b30aadb7fbb8ad36b07ab60df7737c6fe8e9aa201ba76a8a54021ab7321ef9c4
/usr/local/bin/runner.sh: line 50: wait: pid 1225 is not a child of this shell
================================================================================
Done cleaning up after podman in container.
```
</details>

<details>
<summary>Job URLs</summary>

https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16451/pull-kubevirt-e2e-k8s-1.32-sig-storage/2009194142705389568

</details>

<details>
<summary>PR History URLs</summary>

https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16451

</details>