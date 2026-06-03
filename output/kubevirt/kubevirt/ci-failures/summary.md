



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-06-01 (5x / 55.56%)


#### external (4x / 80.00%)

<details>
<summary> download failure in context (2x / 40.00%) </summary>

<hr/>

**1x**: _2026-06-01 14:03:07 &#43;0000 UTC_: <code>14:05:31: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-storage/2061439596666818560#1:build-log.txt%3A590)
<details>
<summary>all...</summary>

* _2026-06-01 14:03:07 &#43;0000 UTC_: <code>14:05:31: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-storage/2061439596666818560#1:build-log.txt%3A590)

</details>

<hr/>

**1x**: _2026-06-01 13:55:33 &#43;0000 UTC_: <code>14:07:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2061439596553572352#1:build-log.txt%3A4208)
<details>
<summary>all...</summary>

* _2026-06-01 13:55:33 &#43;0000 UTC_: <code>14:07:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2061439596553572352#1:build-log.txt%3A4208)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (2x / 40.00%) </summary>

<hr/>

**2x**: _2026-06-01 14:05:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-compute/2061439596796841984#1:build-log.txt%3A239)
<details>
<summary>all...</summary>

* _2026-06-01 14:05:38 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-operator/2061439596926865408#1:build-log.txt%3A301)

* _2026-06-01 14:05:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-compute/2061439596796841984#1:build-log.txt%3A239)

</details>

<hr/>
</details>

#### internal (1x / 20.00%)

<details>
<summary> deployment timeout (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-06-01 14:01:39 &#43;0000 UTC_: <code>14:09:35: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-network/2061439596561960960#1:build-log.txt%3A1002)
<details>
<summary>all...</summary>

* _2026-06-01 14:01:39 &#43;0000 UTC_: <code>14:09:35: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-network/2061439596561960960#1:build-log.txt%3A1002)

</details>

<hr/>
</details>

### 2026-05-28 (3x / 33.33%)


#### needs-investigation (3x / 100.00%)

<details>
<summary> no error snippets (3x / 100.00%) </summary>

<hr/>

**1x**: _2026-05-28 11:19:39 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17822/pull-kubevirt-e2e-k8s-1.35-sig-storage/2059957565743173632)
<details>
<summary>all...</summary>

* _2026-05-28 11:19:39 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17822/pull-kubevirt-e2e-k8s-1.35-sig-storage/2059957565743173632)

</details>

<hr/>

**1x**: _2026-05-28 11:19:26 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17822/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2059957561557258240)
<details>
<summary>all...</summary>

* _2026-05-28 11:19:26 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17822/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2059957561557258240)

</details>

<hr/>

**1x**: _2026-05-28 09:40:36 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17942/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059932689309372416)
<details>
<summary>all...</summary>

* _2026-05-28 09:40:36 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17942/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059932689309372416)

</details>

<hr/>
</details>

### 2026-05-27 (1x / 11.11%)


#### needs-investigation (1x / 100.00%)

<details>
<summary> no error snippets (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-05-27 10:39:00 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17911/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059584804717858816)
<details>
<summary>all...</summary>

* _2026-05-27 10:39:00 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17911/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059584804717858816)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (4x / 44.44%)

<details>
<summary> download failure in context (2x / 22.22%) </summary>

<hr/>

**1x**: _2026-06-01 14:03:07 &#43;0000 UTC_: <code>14:05:31: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-storage/2061439596666818560#1:build-log.txt%3A590)
<details>
<summary>all...</summary>

* _2026-06-01 14:03:07 &#43;0000 UTC_: <code>14:05:31: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-storage/2061439596666818560#1:build-log.txt%3A590)
<details><summary>context</summary>
<pre>14:05:31: Repository rule oci_pull defined at:
14:05:31:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_oci/oci/private/pull.bzl:342:27: in &lt;toplevel&gt;
14:05:31: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/auth?scope=repository:kubevirt/alpine-ext-kernel-boot-demo:pull&amp;service=quay.io] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/www-authenticate.json: GET returned 502 Bad Gateway
14:05:31: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
14:05:31: INFO: Elapsed time: 14.904s
14:05:31: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>

**1x**: _2026-06-01 13:55:33 &#43;0000 UTC_: <code>14:07:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2061439596553572352#1:build-log.txt%3A4208)
<details>
<summary>all...</summary>

* _2026-06-01 13:55:33 &#43;0000 UTC_: <code>14:07:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2061439596553572352#1:build-log.txt%3A4208)
<details><summary>context</summary>
<pre>14:07:03: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
14:07:03: INFO: Elapsed time: 0.412s
14:07:03: INFO: 0 processes.
14:07:03: ERROR: Build failed. Not running target
14:07:04: &#43; rm -f /tmp/kubevirt.deploy.ywrU
make: *** [Makefile:189: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (2x / 22.22%) </summary>

<hr/>

**2x**: _2026-06-01 14:05:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-compute/2061439596796841984#1:build-log.txt%3A239)
<details>
<summary>all...</summary>

* _2026-06-01 14:05:38 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-operator/2061439596926865408#1:build-log.txt%3A301)
<details><summary>context</summary>
<pre>14:07:08: Copying blob sha256:ea389a75c17fde43e7cd5c66516847c2c408b5b3b516bafaeb3aedd7fe03ccf7
14:07:08: Copying blob sha256:41d62e0ccf8c09688ae70cac2e6a437e05e4fb444268ccdf95995c226b7757e5
14:07:14: Error: unable to copy from source docker://quay.io/kubevirt/builder:2602201229-4494d1587: copying system image from manifest list: reading blob sha256:957776ed783da3cb80920be93eaa5961b047f73eaccbd6a6ac22c5df186de45a: fetching blob: received unexpected HTTP status: 502 Bad Gateway
make: *** [Makefile:39: bazel-build-images] Error 125
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-06-01 14:05:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-compute/2061439596796841984#1:build-log.txt%3A239)
<details><summary>context</summary>
<pre>14:23:41: time=&#34;2026-06-01T14:23:41Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (2/3). Error: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 504 Gateway Time-out&#34;
14:23:52: time=&#34;2026-06-01T14:23:52Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (3/3). Error: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 504 Gateway Time-out&#34;
14:23:53: Error: unable to copy from source docker://quay.io/kubevirt/builder:2602201229-4494d1587: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway
make: *** [Makefile:39: bazel-build-images] Error 125
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


</details>

<hr/>
</details>

### internal (1x / 11.11%)

<details>
<summary> deployment timeout (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-06-01 14:01:39 &#43;0000 UTC_: <code>14:09:35: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-network/2061439596561960960#1:build-log.txt%3A1002)
<details>
<summary>all...</summary>

* _2026-06-01 14:01:39 &#43;0000 UTC_: <code>14:09:35: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-network/2061439596561960960#1:build-log.txt%3A1002)
<details><summary>context</summary>
<pre>14:06:15: time=&#34;2026-06-01T14:06:15Z&#34; level=info msg=&#34;Object whereabouts-config applied successfully&#34;
14:06:15: time=&#34;2026-06-01T14:06:15Z&#34; level=info msg=&#34;Object whereabouts applied successfully&#34;
14:06:15: time=&#34;2026-06-01T14:06:15Z&#34; level=info msg=&#34;[node 1]: kubectl --kubeconfig=/etc/kubernetes/admin.conf wait deployment -n cluster-network-addons cluster-network-addons-operator --for condition=Available --timeout=200s&#34;
14:09:35: error: timed out waiting for the condition on deployments/cluster-network-addons-operator
14:09:35:
14:09:35: ===== 3aee673b550f462b438293ae65a86d9f03c54de2ada1f1c52a982ff5e1e7ba94 ====
14:09:35: &#43; NUM_NODES=3</pre>
</details>


</details>

<hr/>
</details>

### needs-investigation (4x / 44.44%)

<details>
<summary> no error snippets (4x / 44.44%) </summary>

<hr/>

**1x**: _2026-05-28 11:19:39 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17822/pull-kubevirt-e2e-k8s-1.35-sig-storage/2059957565743173632)
<details>
<summary>all...</summary>

* _2026-05-28 11:19:39 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17822/pull-kubevirt-e2e-k8s-1.35-sig-storage/2059957565743173632)

</details>

<hr/>

**1x**: _2026-05-28 11:19:26 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17822/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2059957561557258240)
<details>
<summary>all...</summary>

* _2026-05-28 11:19:26 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17822/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2059957561557258240)

</details>

<hr/>

**1x**: _2026-05-28 09:40:36 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17942/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059932689309372416)
<details>
<summary>all...</summary>

* _2026-05-28 09:40:36 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17942/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059932689309372416)

</details>

<hr/>

**1x**: _2026-05-27 10:39:00 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17911/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059584804717858816)
<details>
<summary>all...</summary>

* _2026-05-27 10:39:00 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17911/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059584804717858816)

</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (9x / 100.00%)


#### external (4x / 44.44%)

<details>
<summary> download failure in context (2x / 22.22%) </summary>

<hr/>

**1x**: _2026-06-01 14:03:07 &#43;0000 UTC_: <code>14:05:31: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-storage/2061439596666818560#1:build-log.txt%3A590)
<details>
<summary>all...</summary>

* _2026-06-01 14:03:07 &#43;0000 UTC_: <code>14:05:31: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-storage/2061439596666818560#1:build-log.txt%3A590)

</details>

<hr/>

**1x**: _2026-06-01 13:55:33 &#43;0000 UTC_: <code>14:07:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2061439596553572352#1:build-log.txt%3A4208)
<details>
<summary>all...</summary>

* _2026-06-01 13:55:33 &#43;0000 UTC_: <code>14:07:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2061439596553572352#1:build-log.txt%3A4208)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (2x / 22.22%) </summary>

<hr/>

**2x**: _2026-06-01 14:05:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-compute/2061439596796841984#1:build-log.txt%3A239)
<details>
<summary>all...</summary>

* _2026-06-01 14:05:38 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-operator/2061439596926865408#1:build-log.txt%3A301)

* _2026-06-01 14:05:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-compute/2061439596796841984#1:build-log.txt%3A239)

</details>

<hr/>
</details>

#### needs-investigation (4x / 44.44%)

<details>
<summary> no error snippets (4x / 44.44%) </summary>

<hr/>

**1x**: _2026-05-28 11:19:39 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17822/pull-kubevirt-e2e-k8s-1.35-sig-storage/2059957565743173632)
<details>
<summary>all...</summary>

* _2026-05-28 11:19:39 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17822/pull-kubevirt-e2e-k8s-1.35-sig-storage/2059957565743173632)

</details>

<hr/>

**1x**: _2026-05-28 11:19:26 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17822/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2059957561557258240)
<details>
<summary>all...</summary>

* _2026-05-28 11:19:26 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17822/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2059957561557258240)

</details>

<hr/>

**1x**: _2026-05-28 09:40:36 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17942/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059932689309372416)
<details>
<summary>all...</summary>

* _2026-05-28 09:40:36 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17942/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059932689309372416)

</details>

<hr/>

**1x**: _2026-05-27 10:39:00 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17911/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059584804717858816)
<details>
<summary>all...</summary>

* _2026-05-27 10:39:00 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17911/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059584804717858816)

</details>

<hr/>
</details>

#### internal (1x / 11.11%)

<details>
<summary> deployment timeout (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-06-01 14:01:39 &#43;0000 UTC_: <code>14:09:35: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-network/2061439596561960960#1:build-log.txt%3A1002)
<details>
<summary>all...</summary>

* _2026-06-01 14:01:39 &#43;0000 UTC_: <code>14:09:35: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-network/2061439596561960960#1:build-log.txt%3A1002)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (3x / 33.33%)


#### external (2x / 66.67%)

<details>
<summary> container image pull failure in context (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-06-01 14:05:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-compute/2061439596796841984#1:build-log.txt%3A239)
<details>
<summary>all...</summary>

* _2026-06-01 14:05:38 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-operator/2061439596926865408#1:build-log.txt%3A301)

* _2026-06-01 14:05:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-compute/2061439596796841984#1:build-log.txt%3A239)

</details>

<hr/>
</details>

#### needs-investigation (1x / 33.33%)

<details>
<summary> no error snippets (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-05-28 11:19:26 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17822/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2059957561557258240)
<details>
<summary>all...</summary>

* _2026-05-28 11:19:26 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17822/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2059957561557258240)

</details>

<hr/>
</details>

### sig-storage (4x / 44.44%)


#### needs-investigation (3x / 75.00%)

<details>
<summary> no error snippets (3x / 75.00%) </summary>

<hr/>

**1x**: _2026-05-28 11:19:39 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17822/pull-kubevirt-e2e-k8s-1.35-sig-storage/2059957565743173632)
<details>
<summary>all...</summary>

* _2026-05-28 11:19:39 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17822/pull-kubevirt-e2e-k8s-1.35-sig-storage/2059957565743173632)

</details>

<hr/>

**1x**: _2026-05-28 09:40:36 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17942/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059932689309372416)
<details>
<summary>all...</summary>

* _2026-05-28 09:40:36 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17942/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059932689309372416)

</details>

<hr/>

**1x**: _2026-05-27 10:39:00 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17911/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059584804717858816)
<details>
<summary>all...</summary>

* _2026-05-27 10:39:00 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17911/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059584804717858816)

</details>

<hr/>
</details>

#### external (1x / 25.00%)

<details>
<summary> download failure in context (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-06-01 14:03:07 &#43;0000 UTC_: <code>14:05:31: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-storage/2061439596666818560#1:build-log.txt%3A590)
<details>
<summary>all...</summary>

* _2026-06-01 14:03:07 &#43;0000 UTC_: <code>14:05:31: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-storage/2061439596666818560#1:build-log.txt%3A590)

</details>

<hr/>
</details>

### sig-network (2x / 22.22%)


#### external (1x / 50.00%)

<details>
<summary> download failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-06-01 13:55:33 &#43;0000 UTC_: <code>14:07:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2061439596553572352#1:build-log.txt%3A4208)
<details>
<summary>all...</summary>

* _2026-06-01 13:55:33 &#43;0000 UTC_: <code>14:07:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2061439596553572352#1:build-log.txt%3A4208)

</details>

<hr/>
</details>

#### internal (1x / 50.00%)

<details>
<summary> deployment timeout (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-06-01 14:01:39 &#43;0000 UTC_: <code>14:09:35: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-network/2061439596561960960#1:build-log.txt%3A1002)
<details>
<summary>all...</summary>

* _2026-06-01 14:01:39 &#43;0000 UTC_: <code>14:09:35: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-network/2061439596561960960#1:build-log.txt%3A1002)

</details>

<hr/>
</details>

Last updated: 2026-06-03 04:43:37
