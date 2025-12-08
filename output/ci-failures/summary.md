# CI Failure Analysis Summary

## Non-Fixable: Failed to download bazel dependency

**Quantity:** 48

**Branches:** release-1.5

**SIGs:** sig-compute, sig-network, sig-operator, sig-storage, vgpu

**Kubernetes Versions:** 1.30, 1.31, 1.32, unknown

<details>
<summary>Example Log Snippet</summary>

```
16:26:45: Repository rule http_file defined at:
16:26:45:   /tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/bazel_tools/tools/build_defs/repo/http.bzl:441:28: in <toplevel>
16:26:45: WARNING: Download from https://storage.googleapis.com/rules_docker/aad94363e63d31d574cf701df484b3e8b868a96a/puller-darwin-amd64 failed: class java.io.IOException GET returned 403 Forbidden
16:26:45: ERROR: An error occurred during the fetch of repository 'go_puller_darwin':
16:26:45:    Traceback (most recent call last):
16:26:45: 	File "/tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/bazel_tools/tools/build_defs/repo/http.bzl", line 168, column 33, in _http_file_impl
16:26:45: 		download_info = ctx.download(
16:26:45: Error in download: java.io.IOException: Error downloading [https://storage.googleapis.com/rules_docker/aad94363e63d31d574cf701df484b3e8b868a96a/puller-darwin-amd64] to /tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/go_puller_darwin/file/downloaded: GET returned 403 Forbidden
16:26:45: ERROR: /home/prow/go/src/github.com/kubevirt/kubevirt/WORKSPACE:300:23: fetching http_file rule //external:go_puller_darwin: Traceback (most recent call last):
16:26:45: 	File "/tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/bazel_tools/tools/build_defs/repo/http.bzl", line 168, column 33, in _http_file_impl
--
16:26:45:   /home/prow/go/src/github.com/kubevirt/kubevirt/WORKSPACE:146:10: in <toplevel>
16:26:45: Repository rule http_file defined at:
16:26:45:   /tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/bazel_tools/tools/build_defs/repo/http.bzl:441:28: in <toplevel>
16:26:45: ERROR: /home/prow/go/src/github.com/kubevirt/kubevirt/containerimages/BUILD.bazel:115:16: no such package '@alpine-ext-kernel-boot-demo-container-base//image': no such package '@go_puller_darwin//file': java.io.IOException: Error downloading [https://storage.googleapis.com/rules_docker/aad94363e63d31d574cf701df484b3e8b868a96a/puller-darwin-amd64] to /tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/go_puller_darwin/file/downloaded: GET returned 403 Forbidden and referenced by '//containerimages:alpine-ext-kernel-boot-demo-container'
16:26:45: ERROR: Analysis of target '//:build-other-images_x86_64' failed; build aborted: 
16:26:45: INFO: Elapsed time: 14.576s
--
16:26:59: Repository rule http_file defined at:
16:26:59:   /tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/bazel_tools/tools/build_defs/repo/http.bzl:441:28: in <toplevel>
16:26:59: WARNING: Download from https://storage.googleapis.com/rules_docker/aad94363e63d31d574cf701df484b3e8b868a96a/puller-darwin-amd64 failed: class java.io.IOException GET returned 403 Forbidden
16:26:59: ERROR: An error occurred during the fetch of repository 'go_puller_darwin':
16:26:59:    Traceback (most recent call last):
16:26:59: 	File "/tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/bazel_tools/tools/build_defs/repo/http.bzl", line 168, column 33, in _http_file_impl
16:26:59: 		download_info = ctx.download(
16:26:59: Error in download: java.io.IOException: Error downloading [https://storage.googleapis.com/rules_docker/aad94363e63d31d574cf701df484b3e8b868a96a/puller-darwin-amd64] to /tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/go_puller_darwin/file/downloaded: GET returned 403 Forbidden
16:26:59: ERROR: /home/prow/go/src/github.com/kubevirt/kubevirt/WORKSPACE:300:23: fetching http_file rule //external:go_puller_darwin: Traceback (most recent call last):
16:26:59: 	File "/tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/bazel_tools/tools/build_defs/repo/http.bzl", line 168, column 33, in _http_file_impl
16:26:59: 		download_info = ctx.download(
16:26:59: Error in download: java.io.IOException: Error downloading [https://storage.googleapis.com/rules_docker/aad94363e63d31d574cf701df484b3e8b868a96a/puller-darwin-amd64] to /tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/go_puller_darwin/file/downloaded: GET returned 403 Forbidden
16:26:59: ERROR: /home/prow/go/src/github.com/kubevirt/kubevirt/BUILD.bazel:173:16: no such package '@go_image_base//image': no such package '@go_puller_darwin//file': java.io.IOException: Error downloading [https://storage.googleapis.com/rules_docker/aad94363e63d31d574cf701df484b3e8b868a96a/puller-darwin-amd64] to /tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/go_puller_darwin/file/downloaded: GET returned 403 Forbidden and referenced by '//:passwd-image'
16:26:59: ERROR: Analysis of target '//cmd/virt-controller:virt-controller-image' failed; build aborted: 
16:26:59: INFO: Elapsed time: 13.634s
--
16:27:13: Repository rule http_file defined at:
16:27:13:   /tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/bazel_tools/tools/build_defs/repo/http.bzl:441:28: in <toplevel>
16:27:13: WARNING: Download from https://storage.googleapis.com/rules_docker/aad94363e63d31d574cf701df484b3e8b868a96a/puller-darwin-amd64 failed: class java.io.IOException GET returned 403 Forbidden
16:27:13: ERROR: An error occurred during the fetch of repository 'go_puller_darwin':
16:27:13:    Traceback (most recent call last):
16:27:13: 	File "/tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/bazel_tools/tools/build_defs/repo/http.bzl", line 168, column 33, in _http_file_impl
16:27:13: 		download_info = ctx.download(
16:27:13: Error in download: java.io.IOException: Error downloading [https://storage.googleapis.com/rules_docker/aad94363e63d31d574cf701df484b3e8b868a96a/puller-darwin-amd64] to /tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/go_puller_darwin/file/downloaded: GET returned 403 Forbidden
16:27:13: ERROR: /home/prow/go/src/github.com/kubevirt/kubevirt/WORKSPACE:300:23: fetching http_file rule //external:go_puller_darwin: Traceback (most recent call last):
16:27:13: 	File "/tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/bazel_tools/tools/build_defs/repo/http.bzl", line 168, column 33, in _http_file_impl
16:27:13: 		download_info = ctx.download(
16:27:13: Error in download: java.io.IOException: Error downloading [https://storage.googleapis.com/rules_docker/aad94363e63d31d574cf701df484b3e8b868a96a/puller-darwin-amd64] to /tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/go_puller_darwin/file/downloaded: GET returned 403 Forbidden
16:27:13: ERROR: /home/prow/go/src/github.com/kubevirt/kubevirt/BUILD.bazel:173:16: no such package '@go_image_base//image': no such package '@go_puller_darwin//file': java.io.IOException: Error downloading [https://storage.googleapis.com/rules_docker/aad94363e63d31d574cf701df484b3e8b868a96a/puller-darwin-amd64] to /tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/go_puller_darwin/file/downloaded: GET returned 403 Forbidden and referenced by '//:passwd-image'
16:27:13: ERROR: Analysis of target '//cmd/virt-operator:virt-operator-image' failed; build aborted: 
16:27:13: INFO: Elapsed time: 13.594s
-----------------------------------------
```
</details>

<details>
<summary>Failed Job URLs</summary>

- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-kind-1.30-vgpu-1.5/1999515795889065984
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-kind-1.30-vgpu-1.5/1999505649045606400
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-kind-1.30-vgpu-1.5/1999503544763289600
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-compute-1.5/1999515805900869632
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-compute-1.5/1999505651839012864
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-compute-1.5/1999503557987930112
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.5/1999515811789672448
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.5/1999505658533122048
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.5/1999503564707205120
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-compute-1.5/1999515808757190656
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-compute-1.5/1999505655148318720
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-compute-1.5/1999503561334984704
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.5/1999515797956857856
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.5/1999505649972547584
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.5/1999503553797820416
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-serial-1.5/1999515812154576896
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-serial-1.5/1999505660189872128
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-compute-serial-1.5/1999503566443646976
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-operator-1.5/1999515806081224704
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-operator-1.5/1999505652694650880
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-operator-1.5/1999503558830985216
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-operator-1.5/1999515809705103360
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-operator-1.5/1999505655982985216
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-operator-1.5/1999503562178039808
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-operator-1.5/1999515811944861696
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-operator-1.5/1999505659380371456
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-operator-1.5/1999503565537677312
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-ipv6-sig-network-1.5/1999515797474512896
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-ipv6-sig-network-1.5/1999505649859301376
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-ipv6-sig-network-1.5/1999503553013485568
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-network-1.5/1999515810380386304
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-network-1.5/1999505656826040320
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-network-1.5/1999503563016900608
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-network-1.5/1999515807108829184
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-network-1.5/1999505653449625600
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-network-1.5/1999503559703400448
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-network-1.5/1999515800792207360
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-network-1.5/1999505650249371648
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-network-1.5/1999503556301819904
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.5/1999515810653016064
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.5/1999505657664901120
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.5/1999503563855761408
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-storage-1.5/1999515807553425408
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-storage-1.5/1999505654343012352
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.31-sig-storage-1.5/1999503560496123904
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-storage-1.5/1999515803644334080
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-storage-1.5/1999505651033706496
- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15856/pull-kubevirt-e2e-k8s-1.30-sig-storage-1.5/1999503557153263616
</details>

<details>
<summary>Pull Request History URLs</summary>

- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15856
</details>

## Non-Fixable: Unexpected Build Process Termination

**Quantity:** 1

**Branches:** main

**SIGs:** sig-storage

**Kubernetes Versions:** 1.32

<details>
<summary>Example Log Snippet</summary>

```
no matches found for ERROR in https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16451/pull-kubevirt-e2e-k8s-1.32-sig-storage/2009194142705389568
-----------------------------------------
```
</details>

<details>
<summary>Failed Job URLs</summary>

- https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16451/pull-kubevirt-e2e-k8s-1.32-sig-storage/2009194142705389568
</details>

<details>
<summary>Pull Request History URLs</summary>

- https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16451
</details>

