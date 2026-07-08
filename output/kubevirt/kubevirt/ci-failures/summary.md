



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-07-07 (4x / 26.67%)


#### external (3x / 75.00%)

<details>
<summary> git clone failure in context (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-07-07 15:07:25 &#43;0000 UTC_: <code>15:10:38: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18369/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.9/2074407868987084800#1:build-log.txt%3A720)
<details>
<summary>all...</summary>

* _2026-07-07 15:07:25 &#43;0000 UTC_: <code>15:10:38: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18369/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.9/2074407868987084800#1:build-log.txt%3A720)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-07-07 14:16:30 &#43;0000 UTC_: <code>14:33:48: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18113/pull-kubevirt-e2e-kind-1.34-sev-1.8/2074497679852834816#1:build-log.txt%3A5277)
<details>
<summary>all...</summary>

* _2026-07-07 14:16:30 &#43;0000 UTC_: <code>14:33:48: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18113/pull-kubevirt-e2e-kind-1.34-sev-1.8/2074497679852834816#1:build-log.txt%3A5277)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-07-07 07:15:34 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18272/pull-kubevirt-e2e-kind-1.35-vgpu/2074391719960383488#1:build-log.txt%3A508)
<details>
<summary>all...</summary>

* _2026-07-07 07:15:34 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18272/pull-kubevirt-e2e-kind-1.35-vgpu/2074391719960383488#1:build-log.txt%3A508)

</details>

<hr/>
</details>

#### needs-investigation (1x / 25.00%)

<details>
<summary> no matching pattern (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-07-07 10:48:09 &#43;0000 UTC_: <code>11:01:41: error: timed out waiting for the condition</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18373/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated-1.9/2074410231747579904#1:build-log.txt%3A1042)
<details>
<summary>all...</summary>

* _2026-07-07 10:48:09 &#43;0000 UTC_: <code>11:01:41: error: timed out waiting for the condition</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18373/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated-1.9/2074410231747579904#1:build-log.txt%3A1042)

</details>

<hr/>
</details>

### 2026-07-02 (3x / 20.00%)


#### internal (2x / 66.67%)

<details>
<summary> kind cluster creation failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-07-02 13:04:52 &#43;0000 UTC_: <code>13:14:13: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18323/pull-kubevirt-e2e-kind-1.34-sev-1.7/2072667534082969600#1:build-log.txt%3A635)
<details>
<summary>all...</summary>

* _2026-07-02 13:04:52 &#43;0000 UTC_: <code>13:14:13: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18323/pull-kubevirt-e2e-kind-1.34-sev-1.7/2072667534082969600#1:build-log.txt%3A635)

</details>

<hr/>
</details>
<details>
<summary> make bazel-build target failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-07-02 11:38:26 &#43;0000 UTC_: <code>make: *** [Makefile:28: bazel-build-functests] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18318/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2072645930082897920#1:build-log.txt%3A4921)
<details>
<summary>all...</summary>

* _2026-07-02 11:38:26 &#43;0000 UTC_: <code>make: *** [Makefile:28: bazel-build-functests] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18318/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2072645930082897920#1:build-log.txt%3A4921)

</details>

<hr/>
</details>

#### needs-investigation (1x / 33.33%)

<details>
<summary> no matching pattern (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-07-02 07:28:38 &#43;0000 UTC_: <code>Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072583064571088896#1:build-log.txt%3A640)
<details>
<summary>all...</summary>

* _2026-07-02 07:28:38 &#43;0000 UTC_: <code>Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072583064571088896#1:build-log.txt%3A640)

</details>

<hr/>
</details>

### 2026-07-01 (8x / 53.33%)


#### external (8x / 100.00%)

<details>
<summary> download failure in context (8x / 100.00%) </summary>

<hr/>

**7x**: _2026-07-01 15:43:17 &#43;0000 UTC_: <code>15:45:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.34-sig-network/2072291187099701248#1:build-log.txt%3A588)
<details>
<summary>all...</summary>

* _2026-07-01 23:00:19 &#43;0000 UTC_: <code>23:02:17: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072414264844357632#1:build-log.txt%3A557)

* _2026-07-01 16:16:17 &#43;0000 UTC_: <code>16:18:01: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.36-sig-storage/2072353016559702016#1:build-log.txt%3A561)

* _2026-07-01 15:43:17 &#43;0000 UTC_: <code>15:45:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.34-sig-network/2072291187099701248#1:build-log.txt%3A588)

* _2026-07-01 15:41:18 &#43;0000 UTC_: <code>15:43:41: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072291187003232256#1:build-log.txt%3A591)

* _2026-07-01 15:41:06 &#43;0000 UTC_: <code>15:42:53: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072291186906763264#1:build-log.txt%3A569)

* _2026-07-01 13:49:45 &#43;0000 UTC_: <code>13:51:35: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18308/pull-kubevirt-e2e-k8s-1.36-sig-operator/2072293741271453696#1:build-log.txt%3A562)

* _2026-07-01 12:34:52 &#43;0000 UTC_: <code>12:37:06: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18256/pull-kubevirt-e2e-k8s-1.34-windows2016/2072297118487285760#1:build-log.txt%3A613)

</details>

<hr/>

**1x**: _2026-07-01 13:22:49 &#43;0000 UTC_: <code>13:39:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17955/pull-kubevirt-e2e-k8s-1.34-sig-performance/2072309073948184576#1:build-log.txt%3A2181)
<details>
<summary>all...</summary>

* _2026-07-01 13:22:49 &#43;0000 UTC_: <code>13:39:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17955/pull-kubevirt-e2e-k8s-1.34-sig-performance/2072309073948184576#1:build-log.txt%3A2181)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (11x / 73.33%)

<details>
<summary> download failure in context (9x / 60.00%) </summary>

<hr/>

**7x**: _2026-07-01 15:43:17 &#43;0000 UTC_: <code>15:45:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.34-sig-network/2072291187099701248#1:build-log.txt%3A588)
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


* _2026-07-01 13:49:45 &#43;0000 UTC_: <code>13:51:35: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18308/pull-kubevirt-e2e-k8s-1.36-sig-operator/2072293741271453696#1:build-log.txt%3A562)
<details><summary>context</summary>
<pre>13:51:35:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
13:51:35: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
13:51:35: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
13:51:35: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
13:51:35: INFO: Elapsed time: 1.937s
13:51:35: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 12:34:52 &#43;0000 UTC_: <code>12:37:06: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18256/pull-kubevirt-e2e-k8s-1.34-windows2016/2072297118487285760#1:build-log.txt%3A613)
<details><summary>context</summary>
<pre>12:37:06:  * Hash &#39;7c12895b777bcaa8ccae0605b4de635b68fc32d60fa08f421dc3818bf55ee212&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:7c12895b777bcaa8ccae0605b4de635b68fc32d60fa08f421dc3818bf55ee212
12:37:06: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
12:37:06: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
12:37:06: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
12:37:06: INFO: Elapsed time: 1.173s
12:37:06: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>

**2x**: _2026-07-01 13:22:49 &#43;0000 UTC_: <code>13:39:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17955/pull-kubevirt-e2e-k8s-1.34-sig-performance/2072309073948184576#1:build-log.txt%3A2181)
<details>
<summary>all...</summary>

* _2026-07-07 14:16:30 &#43;0000 UTC_: <code>14:33:48: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18113/pull-kubevirt-e2e-kind-1.34-sev-1.8/2074497679852834816#1:build-log.txt%3A5277)
<details><summary>context</summary>
<pre>14:33:48: ERROR: Analysis of target &#39;//:buildifier&#39; failed; build aborted:
14:33:48: INFO: Elapsed time: 15.157s
14:33:48: INFO: 0 processes.
14:33:48: ERROR: Build failed. Not running target
make: *** [Makefile:28: bazel-build-functests] Error 1
&#43; ret=2
&#43; check_for_panics</pre>
</details>


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
<summary> git clone failure in context (1x / 6.67%) </summary>

<hr/>

**1x**: _2026-07-07 15:07:25 &#43;0000 UTC_: <code>15:10:38: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18369/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.9/2074407868987084800#1:build-log.txt%3A720)
<details>
<summary>all...</summary>

* _2026-07-07 15:07:25 &#43;0000 UTC_: <code>15:10:38: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18369/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.9/2074407868987084800#1:build-log.txt%3A720)
<details><summary>context</summary>
<pre>15:10:38: Cloning into &#39;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/com_github_masterzen_simplexml&#39;...
15:10:38: fatal: unable to access &#39;https://github.com/masterzen/simplexml/&#39;: The requested URL returned error: 503
15:10:38: fetch_repo: exit status 128
15:10:38: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:
15:10:38: INFO: Elapsed time: 3.999s
15:10:38: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 6.67%) </summary>

<hr/>

**1x**: _2026-07-07 07:15:34 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18272/pull-kubevirt-e2e-kind-1.35-vgpu/2074391719960383488#1:build-log.txt%3A508)
<details>
<summary>all...</summary>

* _2026-07-07 07:15:34 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18272/pull-kubevirt-e2e-kind-1.35-vgpu/2074391719960383488#1:build-log.txt%3A508)
<details><summary>context</summary>
<pre>07:23:53: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
07:23:53: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
07:23:58: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage22200702/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/5c/5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20260707%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20260707T072355Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=0bad1f801acf853d196811953f289644a7b3d6f81ef0dc2a0614051b5c9b66bb&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1783409935~hmac=dcf8c26f98e1fd8873e1518f03d0f5eb47d400065d50d069cc837b2ee2abfe0b&#34;: EOF)
make: *** [Makefile:174: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>

<hr/>
</details>

### internal (2x / 13.33%)

<details>
<summary> kind cluster creation failure (1x / 6.67%) </summary>

<hr/>

**1x**: _2026-07-02 13:04:52 &#43;0000 UTC_: <code>13:14:13: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18323/pull-kubevirt-e2e-kind-1.34-sev-1.7/2072667534082969600#1:build-log.txt%3A635)
<details>
<summary>all...</summary>

* _2026-07-02 13:04:52 &#43;0000 UTC_: <code>13:14:13: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18323/pull-kubevirt-e2e-kind-1.34-sev-1.7/2072667534082969600#1:build-log.txt%3A635)
<details><summary>context</summary>
<pre>13:14:10:  ✓ Ensuring node image (kindest/node:v1.34.0) 🖼
13:14:10:  • Preparing nodes 📦 📦   ...
13:14:13:  ✗ Preparing nodes 📦 📦
13:14:13: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;
13:14:13:
13:14:13: Stack Trace:
13:14:13: sigs.k8s.io/kind/pkg/errors.Errorf</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> make bazel-build target failure (1x / 6.67%) </summary>

<hr/>

**1x**: _2026-07-02 11:38:26 &#43;0000 UTC_: <code>make: *** [Makefile:28: bazel-build-functests] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18318/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2072645930082897920#1:build-log.txt%3A4921)
<details>
<summary>all...</summary>

* _2026-07-02 11:38:26 &#43;0000 UTC_: <code>make: *** [Makefile:28: bazel-build-functests] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18318/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2072645930082897920#1:build-log.txt%3A4921)
<details><summary>context</summary>
<pre>12:12:08: Waiting for rsyncd to be ready.............................failed
12:12:08: rsync: [sender] failed to connect to 127.0.0.1 (127.0.0.1): Connection refused (111)
12:12:08: rsync error: error in socket IO (code 10) at clientserver.c(139) [sender=3.4.1]
make: *** [Makefile:28: bazel-build-functests] Error 10
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>


</details>

<hr/>
</details>

### needs-investigation (2x / 13.33%)

<details>
<summary> no matching pattern (2x / 13.33%) </summary>

<hr/>

**1x**: _2026-07-07 10:48:09 &#43;0000 UTC_: <code>11:01:41: error: timed out waiting for the condition</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18373/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated-1.9/2074410231747579904#1:build-log.txt%3A1042)
<details>
<summary>all...</summary>

* _2026-07-07 10:48:09 &#43;0000 UTC_: <code>11:01:41: error: timed out waiting for the condition</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18373/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated-1.9/2074410231747579904#1:build-log.txt%3A1042)
<details><summary>context</summary>
<pre>10:58:21: time=&#34;2026-07-07T10:58:21Z&#34; level=info msg=&#34;Object network-resources-injector applied successfully&#34;
10:58:21: time=&#34;2026-07-07T10:58:21Z&#34; level=info msg=&#34;[node 1]: kubectl --kubeconfig=/etc/kubernetes/admin.conf rollout status -n kube-system deploy/network-resources-injector --timeout=200s&#34;
10:58:21: Waiting for deployment &#34;network-resources-injector&#34; rollout to finish: 0 of 2 updated replicas are available...
11:01:41: error: timed out waiting for the condition
11:01:41:
11:01:41: ===== 44d3e19fddc933f5c3de66835fdefea380289f4c14b5f1802b4370c1c108c5f0 ====
11:01:41: &#43; NUM_NODES=3</pre>
</details>


</details>

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


### release-1.9 (2x / 13.33%)


#### needs-investigation (1x / 50.00%)

<details>
<summary> no matching pattern (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-07-07 10:48:09 &#43;0000 UTC_: <code>11:01:41: error: timed out waiting for the condition</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18373/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated-1.9/2074410231747579904#1:build-log.txt%3A1042)
<details>
<summary>all...</summary>

* _2026-07-07 10:48:09 &#43;0000 UTC_: <code>11:01:41: error: timed out waiting for the condition</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18373/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated-1.9/2074410231747579904#1:build-log.txt%3A1042)

</details>

<hr/>
</details>

#### external (1x / 50.00%)

<details>
<summary> git clone failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-07-07 15:07:25 &#43;0000 UTC_: <code>15:10:38: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18369/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.9/2074407868987084800#1:build-log.txt%3A720)
<details>
<summary>all...</summary>

* _2026-07-07 15:07:25 &#43;0000 UTC_: <code>15:10:38: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18369/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.9/2074407868987084800#1:build-log.txt%3A720)

</details>

<hr/>
</details>

### release-1.8 (2x / 13.33%)


#### internal (1x / 50.00%)

<details>
<summary> make bazel-build target failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-07-02 11:38:26 &#43;0000 UTC_: <code>make: *** [Makefile:28: bazel-build-functests] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18318/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2072645930082897920#1:build-log.txt%3A4921)
<details>
<summary>all...</summary>

* _2026-07-02 11:38:26 &#43;0000 UTC_: <code>make: *** [Makefile:28: bazel-build-functests] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18318/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2072645930082897920#1:build-log.txt%3A4921)

</details>

<hr/>
</details>

#### external (1x / 50.00%)

<details>
<summary> download failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-07-07 14:16:30 &#43;0000 UTC_: <code>14:33:48: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18113/pull-kubevirt-e2e-kind-1.34-sev-1.8/2074497679852834816#1:build-log.txt%3A5277)
<details>
<summary>all...</summary>

* _2026-07-07 14:16:30 &#43;0000 UTC_: <code>14:33:48: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18113/pull-kubevirt-e2e-kind-1.34-sev-1.8/2074497679852834816#1:build-log.txt%3A5277)

</details>

<hr/>
</details>

### main (10x / 66.67%)


#### external (9x / 90.00%)

<details>
<summary> download failure in context (8x / 80.00%) </summary>

<hr/>

**7x**: _2026-07-01 15:43:17 &#43;0000 UTC_: <code>15:45:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.34-sig-network/2072291187099701248#1:build-log.txt%3A588)
<details>
<summary>all...</summary>

* _2026-07-01 23:00:19 &#43;0000 UTC_: <code>23:02:17: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072414264844357632#1:build-log.txt%3A557)

* _2026-07-01 16:16:17 &#43;0000 UTC_: <code>16:18:01: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.36-sig-storage/2072353016559702016#1:build-log.txt%3A561)

* _2026-07-01 15:43:17 &#43;0000 UTC_: <code>15:45:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.34-sig-network/2072291187099701248#1:build-log.txt%3A588)

* _2026-07-01 15:41:18 &#43;0000 UTC_: <code>15:43:41: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072291187003232256#1:build-log.txt%3A591)

* _2026-07-01 15:41:06 &#43;0000 UTC_: <code>15:42:53: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072291186906763264#1:build-log.txt%3A569)

* _2026-07-01 13:49:45 &#43;0000 UTC_: <code>13:51:35: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18308/pull-kubevirt-e2e-k8s-1.36-sig-operator/2072293741271453696#1:build-log.txt%3A562)

* _2026-07-01 12:34:52 &#43;0000 UTC_: <code>12:37:06: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18256/pull-kubevirt-e2e-k8s-1.34-windows2016/2072297118487285760#1:build-log.txt%3A613)

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
<summary> container image pull failure in context (1x / 10.00%) </summary>

<hr/>

**1x**: _2026-07-07 07:15:34 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18272/pull-kubevirt-e2e-kind-1.35-vgpu/2074391719960383488#1:build-log.txt%3A508)
<details>
<summary>all...</summary>

* _2026-07-07 07:15:34 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18272/pull-kubevirt-e2e-kind-1.35-vgpu/2074391719960383488#1:build-log.txt%3A508)

</details>

<hr/>
</details>

#### needs-investigation (1x / 10.00%)

<details>
<summary> no matching pattern (1x / 10.00%) </summary>

<hr/>

**1x**: _2026-07-02 07:28:38 &#43;0000 UTC_: <code>Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072583064571088896#1:build-log.txt%3A640)
<details>
<summary>all...</summary>

* _2026-07-02 07:28:38 &#43;0000 UTC_: <code>Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072583064571088896#1:build-log.txt%3A640)

</details>

<hr/>
</details>

### release-1.7 (1x / 6.67%)


#### internal (1x / 100.00%)

<details>
<summary> kind cluster creation failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-07-02 13:04:52 &#43;0000 UTC_: <code>13:14:13: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18323/pull-kubevirt-e2e-kind-1.34-sev-1.7/2072667534082969600#1:build-log.txt%3A635)
<details>
<summary>all...</summary>

* _2026-07-02 13:04:52 &#43;0000 UTC_: <code>13:14:13: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18323/pull-kubevirt-e2e-kind-1.34-sev-1.7/2072667534082969600#1:build-log.txt%3A635)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-storage (2x / 13.33%)


#### external (2x / 100.00%)

<details>
<summary> git clone failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-07-07 15:07:25 &#43;0000 UTC_: <code>15:10:38: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18369/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.9/2074407868987084800#1:build-log.txt%3A720)
<details>
<summary>all...</summary>

* _2026-07-07 15:07:25 &#43;0000 UTC_: <code>15:10:38: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18369/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.9/2074407868987084800#1:build-log.txt%3A720)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-07-01 16:16:17 &#43;0000 UTC_: <code>16:18:01: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.36-sig-storage/2072353016559702016#1:build-log.txt%3A561)
<details>
<summary>all...</summary>

* _2026-07-01 16:16:17 &#43;0000 UTC_: <code>16:18:01: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.36-sig-storage/2072353016559702016#1:build-log.txt%3A561)

</details>

<hr/>
</details>

### sig-compute (6x / 40.00%)


#### external (4x / 66.67%)

<details>
<summary> download failure in context (3x / 50.00%) </summary>

<hr/>

**1x**: _2026-07-07 14:16:30 &#43;0000 UTC_: <code>14:33:48: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18113/pull-kubevirt-e2e-kind-1.34-sev-1.8/2074497679852834816#1:build-log.txt%3A5277)
<details>
<summary>all...</summary>

* _2026-07-07 14:16:30 &#43;0000 UTC_: <code>14:33:48: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18113/pull-kubevirt-e2e-kind-1.34-sev-1.8/2074497679852834816#1:build-log.txt%3A5277)

</details>

<hr/>

**2x**: _2026-07-01 12:34:52 &#43;0000 UTC_: <code>12:37:06: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18256/pull-kubevirt-e2e-k8s-1.34-windows2016/2072297118487285760#1:build-log.txt%3A613)
<details>
<summary>all...</summary>

* _2026-07-01 13:49:45 &#43;0000 UTC_: <code>13:51:35: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18308/pull-kubevirt-e2e-k8s-1.36-sig-operator/2072293741271453696#1:build-log.txt%3A562)

* _2026-07-01 12:34:52 &#43;0000 UTC_: <code>12:37:06: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18256/pull-kubevirt-e2e-k8s-1.34-windows2016/2072297118487285760#1:build-log.txt%3A613)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-07-07 07:15:34 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18272/pull-kubevirt-e2e-kind-1.35-vgpu/2074391719960383488#1:build-log.txt%3A508)
<details>
<summary>all...</summary>

* _2026-07-07 07:15:34 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18272/pull-kubevirt-e2e-kind-1.35-vgpu/2074391719960383488#1:build-log.txt%3A508)

</details>

<hr/>
</details>

#### needs-investigation (1x / 16.67%)

<details>
<summary> no matching pattern (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-07-02 07:28:38 &#43;0000 UTC_: <code>Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072583064571088896#1:build-log.txt%3A640)
<details>
<summary>all...</summary>

* _2026-07-02 07:28:38 &#43;0000 UTC_: <code>Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072583064571088896#1:build-log.txt%3A640)

</details>

<hr/>
</details>

#### internal (1x / 16.67%)

<details>
<summary> kind cluster creation failure (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-07-02 13:04:52 &#43;0000 UTC_: <code>13:14:13: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18323/pull-kubevirt-e2e-kind-1.34-sev-1.7/2072667534082969600#1:build-log.txt%3A635)
<details>
<summary>all...</summary>

* _2026-07-02 13:04:52 &#43;0000 UTC_: <code>13:14:13: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18323/pull-kubevirt-e2e-kind-1.34-sev-1.7/2072667534082969600#1:build-log.txt%3A635)

</details>

<hr/>
</details>

### sig-network (6x / 40.00%)


#### external (4x / 66.67%)

<details>
<summary> download failure in context (4x / 66.67%) </summary>

<hr/>

**4x**: _2026-07-01 15:43:17 &#43;0000 UTC_: <code>15:45:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.34-sig-network/2072291187099701248#1:build-log.txt%3A588)
<details>
<summary>all...</summary>

* _2026-07-01 23:00:19 &#43;0000 UTC_: <code>23:02:17: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072414264844357632#1:build-log.txt%3A557)

* _2026-07-01 15:43:17 &#43;0000 UTC_: <code>15:45:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.34-sig-network/2072291187099701248#1:build-log.txt%3A588)

* _2026-07-01 15:41:18 &#43;0000 UTC_: <code>15:43:41: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072291187003232256#1:build-log.txt%3A591)

* _2026-07-01 15:41:06 &#43;0000 UTC_: <code>15:42:53: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072291186906763264#1:build-log.txt%3A569)

</details>

<hr/>
</details>

#### needs-investigation (1x / 16.67%)

<details>
<summary> no matching pattern (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-07-07 10:48:09 &#43;0000 UTC_: <code>11:01:41: error: timed out waiting for the condition</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18373/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated-1.9/2074410231747579904#1:build-log.txt%3A1042)
<details>
<summary>all...</summary>

* _2026-07-07 10:48:09 &#43;0000 UTC_: <code>11:01:41: error: timed out waiting for the condition</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18373/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated-1.9/2074410231747579904#1:build-log.txt%3A1042)

</details>

<hr/>
</details>

#### internal (1x / 16.67%)

<details>
<summary> make bazel-build target failure (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-07-02 11:38:26 &#43;0000 UTC_: <code>make: *** [Makefile:28: bazel-build-functests] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18318/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2072645930082897920#1:build-log.txt%3A4921)
<details>
<summary>all...</summary>

* _2026-07-02 11:38:26 &#43;0000 UTC_: <code>make: *** [Makefile:28: bazel-build-functests] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18318/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2072645930082897920#1:build-log.txt%3A4921)

</details>

<hr/>
</details>

### sig-performance (1x / 6.67%)


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

Last updated: 2026-07-08 12:42:30
