



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-07-01 (3x / 60.00%)


#### external (3x / 100.00%)

<details>
<summary> download failure in context (3x / 100.00%) </summary>

<hr/>

**3x**: _2026-07-01 09:23:08 &#43;0000 UTC_: <code>09:24:45: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072241915692060672#1:build-log.txt%3A547)
<details>
<summary>all...</summary>

* _2026-07-01 09:23:08 &#43;0000 UTC_: <code>09:24:45: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072241915692060672#1:build-log.txt%3A547)

* _2026-07-01 08:56:09 &#43;0000 UTC_: <code>08:58:05: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.34-sig-network/2072241915125829632#1:build-log.txt%3A577)

* _2026-07-01 07:11:39 &#43;0000 UTC_: <code>07:13:55: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18275/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2072215977843494912#1:build-log.txt%3A696)

</details>

<hr/>
</details>

### 2026-06-25 (2x / 40.00%)


#### pr-build (1x / 50.00%)

<details>
<summary> panic detected in test output (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-06-25 06:41:50 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.34-sev/2070034457527586816#1:build-log.txt%3A5211)
<details>
<summary>all...</summary>

* _2026-06-25 06:41:50 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.34-sev/2070034457527586816#1:build-log.txt%3A5211)

</details>

<hr/>
</details>

#### external (1x / 50.00%)

<details>
<summary> podman container removal timeout (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-06-25 06:41:41 &#43;0000 UTC_: <code>07:06:26: Error: cannot remove container 0e0afb9af71a75d9b026ee76ea2ecb8cce792a6dec62cf50a08fee527c34ab5c as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2070034428989542400#1:build-log.txt%3A739)
<details>
<summary>all...</summary>

* _2026-06-25 06:41:41 &#43;0000 UTC_: <code>07:06:26: Error: cannot remove container 0e0afb9af71a75d9b026ee76ea2ecb8cce792a6dec62cf50a08fee527c34ab5c as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2070034428989542400#1:build-log.txt%3A739)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (4x / 80.00%)

<details>
<summary> download failure in context (3x / 60.00%) </summary>

<hr/>

**3x**: _2026-07-01 09:23:08 &#43;0000 UTC_: <code>09:24:45: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072241915692060672#1:build-log.txt%3A547)
<details>
<summary>all...</summary>

* _2026-07-01 09:23:08 &#43;0000 UTC_: <code>09:24:45: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072241915692060672#1:build-log.txt%3A547)
<details><summary>context</summary>
<pre>09:24:45:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
09:24:45: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
09:24:45: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
09:24:45: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
09:24:45: INFO: Elapsed time: 0.768s
09:24:45: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 08:56:09 &#43;0000 UTC_: <code>08:58:05: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.34-sig-network/2072241915125829632#1:build-log.txt%3A577)
<details><summary>context</summary>
<pre>08:58:05:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
08:58:05: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
08:58:05: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
08:58:05: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
08:58:05: INFO: Elapsed time: 0.863s
08:58:05: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 07:11:39 &#43;0000 UTC_: <code>07:13:55: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18275/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2072215977843494912#1:build-log.txt%3A696)
<details><summary>context</summary>
<pre>07:13:55:  * Hash &#39;5aed89bce5d6bb3ece39298fe92e2900d15dcecf58ce152e7c7d9f97c6dd43fe&#39; for https://quay.io/v2/kubevirt/fedora-realtime-container-disk/blobs/sha256:5aed89bce5d6bb3ece39298fe92e2900d15dcecf58ce152e7c7d9f97c6dd43fe
07:13:55: If the definition of &#39;repository @fedora_realtime_single&#39; was updated, verify that the hashes were also updated.
07:13:55: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
07:13:55: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
07:13:55: INFO: Elapsed time: 0.826s
07:13:55: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> podman container removal timeout (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-06-25 06:41:41 &#43;0000 UTC_: <code>07:06:26: Error: cannot remove container 0e0afb9af71a75d9b026ee76ea2ecb8cce792a6dec62cf50a08fee527c34ab5c as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2070034428989542400#1:build-log.txt%3A739)
<details>
<summary>all...</summary>

* _2026-06-25 06:41:41 &#43;0000 UTC_: <code>07:06:26: Error: cannot remove container 0e0afb9af71a75d9b026ee76ea2ecb8cce792a6dec62cf50a08fee527c34ab5c as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2070034428989542400#1:build-log.txt%3A739)
<details><summary>context</summary>
<pre>07:06:26: Command Output: time=&#34;2026-06-25T07:06:19Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container kind-1.35-control-plane in 10 seconds, resorting to SIGKILL&#34;
07:06:26: time=&#34;2026-06-25T07:06:21Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container kind-1.35-worker in 10 seconds, resorting to SIGKILL&#34;
07:06:26: Error: cannot remove container bef9fdc1edf782f21c2ad7d5d7c703d5bd4e05e110d7c0928819fcb4babcb18d as it could not be stopped: given PID did not die within timeout
07:06:26: Error: cannot remove container 0e0afb9af71a75d9b026ee76ea2ecb8cce792a6dec62cf50a08fee527c34ab5c as it could not be stopped: given PID did not die within timeout
07:06:26:
07:06:26: Stack Trace:
07:06:26: sigs.k8s.io/kind/pkg/errors.WithStack</pre>
</details>


</details>

<hr/>
</details>

### pr-build (1x / 20.00%)

<details>
<summary> panic detected in test output (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-06-25 06:41:50 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.34-sev/2070034457527586816#1:build-log.txt%3A5211)
<details>
<summary>all...</summary>

* _2026-06-25 06:41:50 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.34-sev/2070034457527586816#1:build-log.txt%3A5211)
<details><summary>context</summary>
<pre>&#43; set &#43;x

================================
ERROR: Found panic in test output
Files:
https://storage.googleapis.com/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.34-sev/2070034457527586816/artifacts/pods/0_kubevirt_virt-handler-dqq6m-virt-handler.log
https://storage.googleapis.com/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.34-sev/2070034457527586816/artifacts/pods/0_kubevirt_virt-handler-dqq6m-virt-handler_previous.log</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (5x / 100.00%)


#### external (4x / 80.00%)

<details>
<summary> download failure in context (3x / 60.00%) </summary>

<hr/>

**3x**: _2026-07-01 09:23:08 &#43;0000 UTC_: <code>09:24:45: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072241915692060672#1:build-log.txt%3A547)
<details>
<summary>all...</summary>

* _2026-07-01 09:23:08 &#43;0000 UTC_: <code>09:24:45: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072241915692060672#1:build-log.txt%3A547)

* _2026-07-01 08:56:09 &#43;0000 UTC_: <code>08:58:05: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.34-sig-network/2072241915125829632#1:build-log.txt%3A577)

* _2026-07-01 07:11:39 &#43;0000 UTC_: <code>07:13:55: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18275/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2072215977843494912#1:build-log.txt%3A696)

</details>

<hr/>
</details>
<details>
<summary> podman container removal timeout (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-06-25 06:41:41 &#43;0000 UTC_: <code>07:06:26: Error: cannot remove container 0e0afb9af71a75d9b026ee76ea2ecb8cce792a6dec62cf50a08fee527c34ab5c as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2070034428989542400#1:build-log.txt%3A739)
<details>
<summary>all...</summary>

* _2026-06-25 06:41:41 &#43;0000 UTC_: <code>07:06:26: Error: cannot remove container 0e0afb9af71a75d9b026ee76ea2ecb8cce792a6dec62cf50a08fee527c34ab5c as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2070034428989542400#1:build-log.txt%3A739)

</details>

<hr/>
</details>

#### pr-build (1x / 20.00%)

<details>
<summary> panic detected in test output (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-06-25 06:41:50 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.34-sev/2070034457527586816#1:build-log.txt%3A5211)
<details>
<summary>all...</summary>

* _2026-06-25 06:41:50 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.34-sev/2070034457527586816#1:build-log.txt%3A5211)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (4x / 80.00%)


#### external (3x / 75.00%)

<details>
<summary> download failure in context (2x / 50.00%) </summary>

<hr/>

**2x**: _2026-07-01 09:23:08 &#43;0000 UTC_: <code>09:24:45: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072241915692060672#1:build-log.txt%3A547)
<details>
<summary>all...</summary>

* _2026-07-01 09:23:08 &#43;0000 UTC_: <code>09:24:45: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072241915692060672#1:build-log.txt%3A547)

* _2026-07-01 07:11:39 &#43;0000 UTC_: <code>07:13:55: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18275/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2072215977843494912#1:build-log.txt%3A696)

</details>

<hr/>
</details>
<details>
<summary> podman container removal timeout (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-06-25 06:41:41 &#43;0000 UTC_: <code>07:06:26: Error: cannot remove container 0e0afb9af71a75d9b026ee76ea2ecb8cce792a6dec62cf50a08fee527c34ab5c as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2070034428989542400#1:build-log.txt%3A739)
<details>
<summary>all...</summary>

* _2026-06-25 06:41:41 &#43;0000 UTC_: <code>07:06:26: Error: cannot remove container 0e0afb9af71a75d9b026ee76ea2ecb8cce792a6dec62cf50a08fee527c34ab5c as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2070034428989542400#1:build-log.txt%3A739)

</details>

<hr/>
</details>

#### pr-build (1x / 25.00%)

<details>
<summary> panic detected in test output (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-06-25 06:41:50 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.34-sev/2070034457527586816#1:build-log.txt%3A5211)
<details>
<summary>all...</summary>

* _2026-06-25 06:41:50 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.34-sev/2070034457527586816#1:build-log.txt%3A5211)

</details>

<hr/>
</details>

### sig-network (1x / 20.00%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-07-01 08:56:09 &#43;0000 UTC_: <code>08:58:05: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.34-sig-network/2072241915125829632#1:build-log.txt%3A577)
<details>
<summary>all...</summary>

* _2026-07-01 08:56:09 &#43;0000 UTC_: <code>08:58:05: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.34-sig-network/2072241915125829632#1:build-log.txt%3A577)

</details>

<hr/>
</details>

Last updated: 2026-07-02 07:17:14
