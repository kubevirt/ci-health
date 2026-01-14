
# CI Failure Analysis Summary

This report summarizes the analysis of CI job failures that did not reach the testing stage.

## Non-Fixable Failures

### Bazel Download Failure

*   **Quantity:** 53
*   **Branches:** release-1.5, main
*   **SIGs:** compute, network, operator, storage, performance
*   **Kubernetes Versions:** 1.30, 1.31, 1.32, 1.33
*   **Pull Requests:** 15856, 16465, 16470

A large number of jobs are failing due to an external dependency issue. The Bazel build process is unable to download `puller-darwin-amd64` from `storage.googleapis.com`, resulting in a `403 Forbidden` error. This is not a failure that can be fixed within the KubeVirt CI.

<details>
<summary>Log Snippet</summary>

```
16:26:45: WARNING: Download from https://storage.googleapis.com/rules_docker/aad94363e63d31d574cf701df484b3e8b868a96a/puller-darwin-amd64 failed: class java.io.IOException GET returned 403 Forbidden
16:26:45: ERROR: An error occurred during the fetch of repository 'go_puller_darwin':
16:26:45:    Traceback (most recent call last):
16:26:45: 	File "/tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/bazel_tools/tools/build_defs/repo/http.bzl", line 168, column 33, in _http_file_impl
16:26:45: 		download_info = ctx.download(
16:26:45: Error in download: java.io.IOException: Error downloading [https://storage.googleapis.com/rules_docker/aad94363e63d31d574cf701df484b3e8b868a96a/puller-darwin-amd64] to /tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/go_puller_darwin/file/downloaded: GET returned 403 Forbidden
```

</details>

<details>
<summary>Failed Job URLs</summary>

https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-kind-1.30-vgpu-1.5/1999515795889065984
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-kind-1.30-vgpu-1.5/1999505649045606400
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-kind-1.30-vgpu-1.5/1999503544763289600
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-compute-1.5/1999515805900869632
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-compute-1.5/1999505651839012864
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-compute-1.5/1999503557987930112
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.5/1999515811789672448
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.5/1999505658533122048
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.5/1999503564707205120
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-compute-1.5/1999515808757190656
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-compute-1.5/1999505655148318720
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-compute-1.5/1999503561334984704
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
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-network-1.5/1999515807108829184
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-network-1.5/1999505653449625600
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-network-1.5/1999503559703400448
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-network-1.5/1999515800792207360
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-network-1.5/1999505650249371648
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-network-1.5/1999503556301819904
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-operator-1.5/1999515811944861696
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-operator-1.5/1999505659380371456
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-operator-1.5/1999503565537677312
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-operator-1.5/1999515806081224704
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-operator-1.5/1999505652694650880
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-operator-1.5/1999503558830985216
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-operator-1.5/1999515809705103360
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-operator-1.5/1999505655982985216
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-operator-1.5/1999503562178039808
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16465/pull-kubevirt-e2e-k8s-1.31-sig-performance-1.5/2008229041038954496
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-performance-1.5/1999515795620630528
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-performance-1.5/1999505648546484224
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-performance-1.5/1999503544612294656
https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16470/pull-kubevirt-e2e-k8s-1.31-sig-performance-1.6/2008490702794657792
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
https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16465
https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16470

</details>

### API Server Error

*   **Quantity:** 1
*   **Branches:** main
*   **SIGs:** compute
*   **Kubernetes Versions:** 1.33
*   **Pull Requests:** 16484

One job failed with an error message `Body was not decodable`. This could be a transient network issue or a bug in the client.

<details>
<summary>Log Snippet</summary>

```
17:19:17: I0108 12:19:17.274661    1618 request.go:1664] "Body was not decodable (unable to check for Status)" err="couldn't get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \"json:\\\"apiVersion,omitempty\\\"\"; Kind string \"json:\\\"kind,omitempty\\\"\" }"
```

</details>

<details>
<summary>Failed Job URLs</summary>

https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16484/pull-kubevirt-e2e-k8s-1.33-sig-compute/2009292795868614657

</details>

<details>
<summary>PR History URLs</summary>

https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16484

</details>

### Unknown Failures

*   **Quantity:** 2
*   **Branches:** main
*   **SIGs:** storage, monitoring
*   **Kubernetes Versions:** 1.32
*   **Pull Requests:** 16451

Two jobs failed without any specific error message in the log. Further investigation of the full build logs is required to determine the root cause of these failures.

<details>
<summary>Failed Job URLs</summary>

https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16451/pull-kubevirt-e2e-k8s-1.32-sig-storage/2009194142705389568

</details>

<details>
<summary>PR History URLs</summary>

https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16451

</details>
