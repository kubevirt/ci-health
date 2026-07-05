



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-07-02 (1x / 4.76%)


#### needs-investigation (1x / 100.00%)

<details>
<summary> no matching pattern (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-07-02 07:28:38 &#43;0000 UTC_: <code>Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072583064571088896#1:build-log.txt%3A640)
<details>
<summary>all...</summary>

* _2026-07-02 07:28:38 &#43;0000 UTC_: <code>Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072583064571088896#1:build-log.txt%3A640)

</details>

<hr/>
</details>

### 2026-07-01 (19x / 90.48%)


#### external (19x / 100.00%)

<details>
<summary> download failure in context (19x / 100.00%) </summary>

<hr/>

**18x**: _2026-07-01 15:43:17 &#43;0000 UTC_: <code>15:45:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.34-sig-network/2072291187099701248#1:build-log.txt%3A588)
<details>
<summary>all...</summary>

* _2026-07-01 23:00:19 &#43;0000 UTC_: <code>23:02:17: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072414264844357632#1:build-log.txt%3A557)

* _2026-07-01 16:16:17 &#43;0000 UTC_: <code>16:18:01: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.36-sig-storage/2072353016559702016#1:build-log.txt%3A561)

* _2026-07-01 15:43:17 &#43;0000 UTC_: <code>15:45:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.34-sig-network/2072291187099701248#1:build-log.txt%3A588)

* _2026-07-01 15:41:18 &#43;0000 UTC_: <code>15:43:41: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072291187003232256#1:build-log.txt%3A591)

* _2026-07-01 15:41:06 &#43;0000 UTC_: <code>15:42:53: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072291186906763264#1:build-log.txt%3A569)

* _2026-07-01 11:33:46 &#43;0000 UTC_: <code>11:36:21: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.35-sig-operator/2072238776180019200#1:build-log.txt%3A665)

* _2026-07-01 11:30:18 &#43;0000 UTC_: <code>11:31:57: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.34-sig-network/2072238774993031168#1:build-log.txt%3A671)

* _2026-07-01 11:29:20 &#43;0000 UTC_: <code>11:31:11: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072238774904950784#1:build-log.txt%3A651)

* _2026-07-01 09:28:19 &#43;0000 UTC_: <code>09:29:53: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18298/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072238855989235712#1:build-log.txt%3A573)

* _2026-07-01 09:23:08 &#43;0000 UTC_: <code>09:24:45: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072241915692060672#1:build-log.txt%3A547)

* _2026-07-01 08:56:09 &#43;0000 UTC_: <code>08:58:05: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.34-sig-network/2072241915125829632#1:build-log.txt%3A577)

* _2026-07-01 08:51:37 &#43;0000 UTC_: <code>08:53:13: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18298/pull-kubevirt-e2e-k8s-1.34-sig-operator/2072238855586582528#1:build-log.txt%3A567)

* _2026-07-01 08:42:43 &#43;0000 UTC_: <code>08:44:20: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-sig-compute/2072235893535543296#1:build-log.txt%3A577)

* _2026-07-01 08:39:41 &#43;0000 UTC_: <code>08:41:54: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-sig-storage/2072235893455851520#1:build-log.txt%3A545)

* _2026-07-01 08:39:26 &#43;0000 UTC_: <code>08:41:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072235893300662272#1:build-log.txt%3A576)

* _2026-07-01 08:36:56 &#43;0000 UTC_: <code>08:38:52: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072235893233553408#1:build-log.txt%3A592)

* _2026-07-01 08:31:01 &#43;0000 UTC_: <code>08:32:48: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-windows2016/2072235893103529984#1:build-log.txt%3A578)

* _2026-07-01 07:11:39 &#43;0000 UTC_: <code>07:13:55: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18275/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2072215977843494912#1:build-log.txt%3A696)

</details>

<hr/>

**1x**: _2026-07-01 13:22:49 &#43;0000 UTC_: <code>13:39:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17955/pull-kubevirt-e2e-k8s-1.34-sig-performance/2072309073948184576#1:build-log.txt%3A2181)
<details>
<summary>all...</summary>

* _2026-07-01 13:22:49 &#43;0000 UTC_: <code>13:39:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17955/pull-kubevirt-e2e-k8s-1.34-sig-performance/2072309073948184576#1:build-log.txt%3A2181)

</details>

<hr/>
</details>

### 2026-06-29 (1x / 4.76%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-06-29 15:49:57 &#43;0000 UTC_: <code>15:54:22: I0629 11:54:22.111079    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18269/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2071621787858243584#1:build-log.txt%3A817)
<details>
<summary>all...</summary>

* _2026-06-29 15:49:57 &#43;0000 UTC_: <code>15:54:22: I0629 11:54:22.111079    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18269/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2071621787858243584#1:build-log.txt%3A817)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (20x / 95.24%)

<details>
<summary> download failure in context (19x / 90.48%) </summary>

<hr/>

**18x**: _2026-07-01 15:43:17 &#43;0000 UTC_: <code>15:45:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.34-sig-network/2072291187099701248#1:build-log.txt%3A588)
<details>
<summary>all...</summary>

* _2026-07-01 23:00:19 &#43;0000 UTC_: <code>23:02:17: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072414264844357632#1:build-log.txt%3A557)
<details><summary>context</summary>
<pre>23:02:17:  * Hash &#39;7c12895b777bcaa8ccae0605b4de635b68fc32d60fa08f421dc3818bf55ee212&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:7c12895b777bcaa8ccae0605b4de635b68fc32d60fa08f421dc3818bf55ee212
23:02:17: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
23:02:17: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
23:02:17: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
23:02:17: INFO: Elapsed time: 0.829s
23:02:17: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 16:16:17 &#43;0000 UTC_: <code>16:18:01: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.36-sig-storage/2072353016559702016#1:build-log.txt%3A561)
<details><summary>context</summary>
<pre>16:18:01:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
16:18:01: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
16:18:01: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
16:18:01: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
16:18:01: INFO: Elapsed time: 0.933s
16:18:01: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 15:43:17 &#43;0000 UTC_: <code>15:45:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.34-sig-network/2072291187099701248#1:build-log.txt%3A588)
<details><summary>context</summary>
<pre>15:45:00:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
15:45:00: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
15:45:00: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
15:45:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
15:45:00: INFO: Elapsed time: 1.118s
15:45:00: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 15:41:18 &#43;0000 UTC_: <code>15:43:41: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072291187003232256#1:build-log.txt%3A591)
<details><summary>context</summary>
<pre>15:43:41:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
15:43:41: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
15:43:41: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
15:43:41: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
15:43:41: INFO: Elapsed time: 0.808s
15:43:41: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 15:41:06 &#43;0000 UTC_: <code>15:42:53: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072291186906763264#1:build-log.txt%3A569)
<details><summary>context</summary>
<pre>15:42:53:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
15:42:53: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
15:42:53: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
15:42:53: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
15:42:53: INFO: Elapsed time: 0.958s
15:42:53: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 11:33:46 &#43;0000 UTC_: <code>11:36:21: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.35-sig-operator/2072238776180019200#1:build-log.txt%3A665)
<details><summary>context</summary>
<pre>11:36:21:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
11:36:21: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
11:36:21: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
11:36:21: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
11:36:21: INFO: Elapsed time: 1.137s
11:36:21: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 11:30:18 &#43;0000 UTC_: <code>11:31:57: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.34-sig-network/2072238774993031168#1:build-log.txt%3A671)
<details><summary>context</summary>
<pre>11:31:57:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
11:31:57: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
11:31:57: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
11:31:57: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
11:31:57: INFO: Elapsed time: 1.086s
11:31:57: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 11:29:20 &#43;0000 UTC_: <code>11:31:11: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072238774904950784#1:build-log.txt%3A651)
<details><summary>context</summary>
<pre>11:31:11:  * Hash &#39;7c12895b777bcaa8ccae0605b4de635b68fc32d60fa08f421dc3818bf55ee212&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:7c12895b777bcaa8ccae0605b4de635b68fc32d60fa08f421dc3818bf55ee212
11:31:11: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
11:31:11: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
11:31:11: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
11:31:11: INFO: Elapsed time: 2.058s
11:31:11: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 09:28:19 &#43;0000 UTC_: <code>09:29:53: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18298/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072238855989235712#1:build-log.txt%3A573)
<details><summary>context</summary>
<pre>09:29:53:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
09:29:53: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
09:29:53: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
09:29:53: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
09:29:53: INFO: Elapsed time: 0.834s
09:29:53: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


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


* _2026-07-01 08:51:37 &#43;0000 UTC_: <code>08:53:13: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18298/pull-kubevirt-e2e-k8s-1.34-sig-operator/2072238855586582528#1:build-log.txt%3A567)
<details><summary>context</summary>
<pre>08:53:13:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
08:53:13: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
08:53:13: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
08:53:13: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
08:53:13: INFO: Elapsed time: 0.707s
08:53:13: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 08:42:43 &#43;0000 UTC_: <code>08:44:20: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-sig-compute/2072235893535543296#1:build-log.txt%3A577)
<details><summary>context</summary>
<pre>08:44:20:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
08:44:20: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
08:44:20: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
08:44:20: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
08:44:20: INFO: Elapsed time: 0.923s
08:44:20: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 08:39:41 &#43;0000 UTC_: <code>08:41:54: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-sig-storage/2072235893455851520#1:build-log.txt%3A545)
<details><summary>context</summary>
<pre>08:41:54:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
08:41:54: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
08:41:54: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
08:41:54: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
08:41:54: INFO: Elapsed time: 0.599s
08:41:54: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 08:39:26 &#43;0000 UTC_: <code>08:41:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072235893300662272#1:build-log.txt%3A576)
<details><summary>context</summary>
<pre>08:41:00:  * Hash &#39;a53fd982787799c2d8cfaa37a2b6fbac4f416437768a25d2eb246dff46bb9d79&#39; for https://quay.io/v2/kubevirtci/fedora-with-test-tooling/manifests/sha256:a53fd982787799c2d8cfaa37a2b6fbac4f416437768a25d2eb246dff46bb9d79
08:41:00: If the definition of &#39;repository @fedora_with_test_tooling_single&#39; was updated, verify that the hashes were also updated.
08:41:00: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
08:41:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
08:41:00: INFO: Elapsed time: 0.930s
08:41:00: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 08:36:56 &#43;0000 UTC_: <code>08:38:52: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072235893233553408#1:build-log.txt%3A592)
<details><summary>context</summary>
<pre>08:38:52:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
08:38:52: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
08:38:52: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
08:38:52: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
08:38:52: INFO: Elapsed time: 0.752s
08:38:52: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 08:31:01 &#43;0000 UTC_: <code>08:32:48: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-windows2016/2072235893103529984#1:build-log.txt%3A578)
<details><summary>context</summary>
<pre>08:32:48:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
08:32:48: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
08:32:48: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
08:32:48: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
08:32:48: INFO: Elapsed time: 0.637s
08:32:48: INFO: 0 processes.
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

**1x**: _2026-07-01 13:22:49 &#43;0000 UTC_: <code>13:39:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17955/pull-kubevirt-e2e-k8s-1.34-sig-performance/2072309073948184576#1:build-log.txt%3A2181)
<details>
<summary>all...</summary>

* _2026-07-01 13:22:49 &#43;0000 UTC_: <code>13:39:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17955/pull-kubevirt-e2e-k8s-1.34-sig-performance/2072309073948184576#1:build-log.txt%3A2181)
<details><summary>context</summary>
<pre>13:39:54: ERROR: Analysis of target &#39;//:push-alpine-ext-kernel-boot-demo&#39; failed; build aborted:
13:39:54: INFO: Elapsed time: 1.124s
13:39:54: INFO: 0 processes.
13:39:54: ERROR: Build failed. Not running target
13:39:58: &#43; rm -f /tmp/kubevirt.deploy.azEg
make: *** [Makefile:189: cluster-sync] Error 1
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 4.76%) </summary>

<hr/>

**1x**: _2026-06-29 15:49:57 &#43;0000 UTC_: <code>15:54:22: I0629 11:54:22.111079    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18269/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2071621787858243584#1:build-log.txt%3A817)
<details>
<summary>all...</summary>

* _2026-06-29 15:49:57 &#43;0000 UTC_: <code>15:54:22: I0629 11:54:22.111079    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18269/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2071621787858243584#1:build-log.txt%3A817)
<details><summary>context</summary>
<pre>15:54:20: I0629 11:54:20.760234    1615 wait.go:283] kube-apiserver check failed at https://[fd00::101]:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;
15:54:21: I0629 11:54:21.109447    1615 wait.go:283] kube-apiserver check failed at https://[fd00::101]:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;
15:54:21: I0629 11:54:21.610150    1615 wait.go:283] kube-apiserver check failed at https://[fd00::101]:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;
15:54:22: I0629 11:54:22.111079    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
15:54:22: I0629 11:54:22.111196    1615 wait.go:283] kube-apiserver check failed at https://[fd00::101]:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
15:54:22: [control-plane-check] kube-apiserver is healthy after 6.502838664s
15:54:22: I0629 11:54:22.617481    1615 kubeconfig.go:657] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists</pre>
</details>


</details>

<hr/>
</details>

### needs-investigation (1x / 4.76%)

<details>
<summary> no matching pattern (1x / 4.76%) </summary>

<hr/>

**1x**: _2026-07-02 07:28:38 &#43;0000 UTC_: <code>Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072583064571088896#1:build-log.txt%3A640)
<details>
<summary>all...</summary>

* _2026-07-02 07:28:38 &#43;0000 UTC_: <code>Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072583064571088896#1:build-log.txt%3A640)
<details><summary>context</summary>
<pre>time=&#34;2026-07-02T07:38:37Z&#34; level=warning msg=&#34;Found incomplete layer \&#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f\&#34;, deleting it&#34;
time=&#34;2026-07-02T07:38:39Z&#34; level=warning msg=&#34;Found incomplete layer \&#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f\&#34;, deleting it&#34;
time=&#34;2026-07-02T07:38:39Z&#34; level=error msg=&#34;cleaning up storage: removing container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer \&#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f\&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-183572116/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link&#34;
Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link
time=&#34;2026-07-02T07:38:39Z&#34; level=warning msg=&#34;Found incomplete layer \&#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f\&#34;, deleting it&#34;
/usr/local/bin/runner.sh: line 50: wait: pid 1220 is not a child of this shell
================================================================================</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (21x / 100.00%)


#### external (20x / 95.24%)

<details>
<summary> download failure in context (19x / 90.48%) </summary>

<hr/>

**18x**: _2026-07-01 15:43:17 &#43;0000 UTC_: <code>15:45:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.34-sig-network/2072291187099701248#1:build-log.txt%3A588)
<details>
<summary>all...</summary>

* _2026-07-01 23:00:19 &#43;0000 UTC_: <code>23:02:17: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072414264844357632#1:build-log.txt%3A557)

* _2026-07-01 16:16:17 &#43;0000 UTC_: <code>16:18:01: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.36-sig-storage/2072353016559702016#1:build-log.txt%3A561)

* _2026-07-01 15:43:17 &#43;0000 UTC_: <code>15:45:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.34-sig-network/2072291187099701248#1:build-log.txt%3A588)

* _2026-07-01 15:41:18 &#43;0000 UTC_: <code>15:43:41: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072291187003232256#1:build-log.txt%3A591)

* _2026-07-01 15:41:06 &#43;0000 UTC_: <code>15:42:53: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072291186906763264#1:build-log.txt%3A569)

* _2026-07-01 11:33:46 &#43;0000 UTC_: <code>11:36:21: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.35-sig-operator/2072238776180019200#1:build-log.txt%3A665)

* _2026-07-01 11:30:18 &#43;0000 UTC_: <code>11:31:57: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.34-sig-network/2072238774993031168#1:build-log.txt%3A671)

* _2026-07-01 11:29:20 &#43;0000 UTC_: <code>11:31:11: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072238774904950784#1:build-log.txt%3A651)

* _2026-07-01 09:28:19 &#43;0000 UTC_: <code>09:29:53: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18298/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072238855989235712#1:build-log.txt%3A573)

* _2026-07-01 09:23:08 &#43;0000 UTC_: <code>09:24:45: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072241915692060672#1:build-log.txt%3A547)

* _2026-07-01 08:56:09 &#43;0000 UTC_: <code>08:58:05: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.34-sig-network/2072241915125829632#1:build-log.txt%3A577)

* _2026-07-01 08:51:37 &#43;0000 UTC_: <code>08:53:13: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18298/pull-kubevirt-e2e-k8s-1.34-sig-operator/2072238855586582528#1:build-log.txt%3A567)

* _2026-07-01 08:42:43 &#43;0000 UTC_: <code>08:44:20: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-sig-compute/2072235893535543296#1:build-log.txt%3A577)

* _2026-07-01 08:39:41 &#43;0000 UTC_: <code>08:41:54: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-sig-storage/2072235893455851520#1:build-log.txt%3A545)

* _2026-07-01 08:39:26 &#43;0000 UTC_: <code>08:41:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072235893300662272#1:build-log.txt%3A576)

* _2026-07-01 08:36:56 &#43;0000 UTC_: <code>08:38:52: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072235893233553408#1:build-log.txt%3A592)

* _2026-07-01 08:31:01 &#43;0000 UTC_: <code>08:32:48: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-windows2016/2072235893103529984#1:build-log.txt%3A578)

* _2026-07-01 07:11:39 &#43;0000 UTC_: <code>07:13:55: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18275/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2072215977843494912#1:build-log.txt%3A696)

</details>

<hr/>

**1x**: _2026-07-01 13:22:49 &#43;0000 UTC_: <code>13:39:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17955/pull-kubevirt-e2e-k8s-1.34-sig-performance/2072309073948184576#1:build-log.txt%3A2181)
<details>
<summary>all...</summary>

* _2026-07-01 13:22:49 &#43;0000 UTC_: <code>13:39:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17955/pull-kubevirt-e2e-k8s-1.34-sig-performance/2072309073948184576#1:build-log.txt%3A2181)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 4.76%) </summary>

<hr/>

**1x**: _2026-06-29 15:49:57 &#43;0000 UTC_: <code>15:54:22: I0629 11:54:22.111079    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18269/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2071621787858243584#1:build-log.txt%3A817)
<details>
<summary>all...</summary>

* _2026-06-29 15:49:57 &#43;0000 UTC_: <code>15:54:22: I0629 11:54:22.111079    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18269/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2071621787858243584#1:build-log.txt%3A817)

</details>

<hr/>
</details>

#### needs-investigation (1x / 4.76%)

<details>
<summary> no matching pattern (1x / 4.76%) </summary>

<hr/>

**1x**: _2026-07-02 07:28:38 &#43;0000 UTC_: <code>Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072583064571088896#1:build-log.txt%3A640)
<details>
<summary>all...</summary>

* _2026-07-02 07:28:38 &#43;0000 UTC_: <code>Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072583064571088896#1:build-log.txt%3A640)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (8x / 38.10%)


#### external (7x / 87.50%)

<details>
<summary> download failure in context (7x / 87.50%) </summary>

<hr/>

**7x**: _2026-07-01 08:42:43 &#43;0000 UTC_: <code>08:44:20: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-sig-compute/2072235893535543296#1:build-log.txt%3A577)
<details>
<summary>all...</summary>

* _2026-07-01 11:33:46 &#43;0000 UTC_: <code>11:36:21: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.35-sig-operator/2072238776180019200#1:build-log.txt%3A665)

* _2026-07-01 09:28:19 &#43;0000 UTC_: <code>09:29:53: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18298/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072238855989235712#1:build-log.txt%3A573)

* _2026-07-01 09:23:08 &#43;0000 UTC_: <code>09:24:45: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072241915692060672#1:build-log.txt%3A547)

* _2026-07-01 08:51:37 &#43;0000 UTC_: <code>08:53:13: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18298/pull-kubevirt-e2e-k8s-1.34-sig-operator/2072238855586582528#1:build-log.txt%3A567)

* _2026-07-01 08:42:43 &#43;0000 UTC_: <code>08:44:20: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-sig-compute/2072235893535543296#1:build-log.txt%3A577)

* _2026-07-01 08:31:01 &#43;0000 UTC_: <code>08:32:48: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-windows2016/2072235893103529984#1:build-log.txt%3A578)

* _2026-07-01 07:11:39 &#43;0000 UTC_: <code>07:13:55: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18275/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2072215977843494912#1:build-log.txt%3A696)

</details>

<hr/>
</details>

#### needs-investigation (1x / 12.50%)

<details>
<summary> no matching pattern (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-07-02 07:28:38 &#43;0000 UTC_: <code>Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072583064571088896#1:build-log.txt%3A640)
<details>
<summary>all...</summary>

* _2026-07-02 07:28:38 &#43;0000 UTC_: <code>Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072583064571088896#1:build-log.txt%3A640)

</details>

<hr/>
</details>

### sig-network (10x / 47.62%)


#### external (10x / 100.00%)

<details>
<summary> download failure in context (9x / 90.00%) </summary>

<hr/>

**9x**: _2026-07-01 15:43:17 &#43;0000 UTC_: <code>15:45:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.34-sig-network/2072291187099701248#1:build-log.txt%3A588)
<details>
<summary>all...</summary>

* _2026-07-01 23:00:19 &#43;0000 UTC_: <code>23:02:17: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072414264844357632#1:build-log.txt%3A557)

* _2026-07-01 15:43:17 &#43;0000 UTC_: <code>15:45:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.34-sig-network/2072291187099701248#1:build-log.txt%3A588)

* _2026-07-01 15:41:18 &#43;0000 UTC_: <code>15:43:41: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072291187003232256#1:build-log.txt%3A591)

* _2026-07-01 15:41:06 &#43;0000 UTC_: <code>15:42:53: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072291186906763264#1:build-log.txt%3A569)

* _2026-07-01 11:30:18 &#43;0000 UTC_: <code>11:31:57: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.34-sig-network/2072238774993031168#1:build-log.txt%3A671)

* _2026-07-01 11:29:20 &#43;0000 UTC_: <code>11:31:11: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072238774904950784#1:build-log.txt%3A651)

* _2026-07-01 08:56:09 &#43;0000 UTC_: <code>08:58:05: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.34-sig-network/2072241915125829632#1:build-log.txt%3A577)

* _2026-07-01 08:39:26 &#43;0000 UTC_: <code>08:41:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072235893300662272#1:build-log.txt%3A576)

* _2026-07-01 08:36:56 &#43;0000 UTC_: <code>08:38:52: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072235893233553408#1:build-log.txt%3A592)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 10.00%) </summary>

<hr/>

**1x**: _2026-06-29 15:49:57 &#43;0000 UTC_: <code>15:54:22: I0629 11:54:22.111079    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18269/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2071621787858243584#1:build-log.txt%3A817)
<details>
<summary>all...</summary>

* _2026-06-29 15:49:57 &#43;0000 UTC_: <code>15:54:22: I0629 11:54:22.111079    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18269/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2071621787858243584#1:build-log.txt%3A817)

</details>

<hr/>
</details>

### sig-performance (1x / 4.76%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-07-01 13:22:49 &#43;0000 UTC_: <code>13:39:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17955/pull-kubevirt-e2e-k8s-1.34-sig-performance/2072309073948184576#1:build-log.txt%3A2181)
<details>
<summary>all...</summary>

* _2026-07-01 13:22:49 &#43;0000 UTC_: <code>13:39:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17955/pull-kubevirt-e2e-k8s-1.34-sig-performance/2072309073948184576#1:build-log.txt%3A2181)

</details>

<hr/>
</details>

### sig-storage (2x / 9.52%)


#### external (2x / 100.00%)

<details>
<summary> download failure in context (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-07-01 08:39:41 &#43;0000 UTC_: <code>08:41:54: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-sig-storage/2072235893455851520#1:build-log.txt%3A545)
<details>
<summary>all...</summary>

* _2026-07-01 16:16:17 &#43;0000 UTC_: <code>16:18:01: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.36-sig-storage/2072353016559702016#1:build-log.txt%3A561)

* _2026-07-01 08:39:41 &#43;0000 UTC_: <code>08:41:54: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-sig-storage/2072235893455851520#1:build-log.txt%3A545)

</details>

<hr/>
</details>

Last updated: 2026-07-05 21:31:00
