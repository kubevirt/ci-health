



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-04-21 (8x / 72.73%)


#### external (8x / 100.00%)

<details>
<summary> container image pull failure in context (7x / 87.50%) </summary>

<hr/>

**7x**: _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-compute/2046711709879504896#1:build-log.txt%3A205)
<details>
<summary>all...</summary>

* _2026-04-21 22:10:16 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-kind-sriov/2046711709338439680#1:build-log.txt%3A196)

* _2026-04-21 22:05:04 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-kind-1.35-vgpu/2046711709095170048#1:build-log.txt%3A189)

* _2026-04-21 22:05:03 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-operator/2046711709980168192#1:build-log.txt%3A228)

* _2026-04-21 22:05:03 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-network/2046711709594292224#1:build-log.txt%3A238)

* _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-compute/2046711709879504896#1:build-log.txt%3A205)

* _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-windows2016/2046711708981923840#1:build-log.txt%3A250)

* _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2046711709439102976#1:build-log.txt%3A238)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-04-21 22:05:03 &#43;0000 UTC_: <code>22:07:22: ERROR: Analysis of target &#39;//containerimages:alpine-with-test-tooling&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-storage/2046711709757870080#1:build-log.txt%3A569)
<details>
<summary>all...</summary>

* _2026-04-21 22:05:03 &#43;0000 UTC_: <code>22:07:22: ERROR: Analysis of target &#39;//containerimages:alpine-with-test-tooling&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-storage/2046711709757870080#1:build-log.txt%3A569)

</details>

<hr/>
</details>

### 2026-04-17 (2x / 18.18%)


#### external (2x / 100.00%)

<details>
<summary> download failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-04-17 13:25:22 &#43;0000 UTC_: <code>13:37:52: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-k8s-1.34-windows2016/2045131369809448960#1:build-log.txt%3A4706)
<details>
<summary>all...</summary>

* _2026-04-17 13:25:22 &#43;0000 UTC_: <code>13:37:52: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-k8s-1.34-windows2016/2045131369809448960#1:build-log.txt%3A4706)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-04-17 13:33:58 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-kind-sriov/2045131370006581248#1:build-log.txt%3A1394)
<details>
<summary>all...</summary>

* _2026-04-17 13:33:58 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-kind-sriov/2045131370006581248#1:build-log.txt%3A1394)

</details>

<hr/>
</details>

### 2026-04-16 (1x / 9.09%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-04-16 20:05:44 &#43;0000 UTC_: <code>20:18:08: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17437/pull-kubevirt-e2e-k8s-1.35-sig-compute/2044869730921091072#1:build-log.txt%3A4004)
<details>
<summary>all...</summary>

* _2026-04-16 20:05:44 &#43;0000 UTC_: <code>20:18:08: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17437/pull-kubevirt-e2e-k8s-1.35-sig-compute/2044869730921091072#1:build-log.txt%3A4004)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (11x / 100.00%)

<details>
<summary> container image pull failure in context (7x / 63.64%) </summary>

<hr/>

**7x**: _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-compute/2046711709879504896#1:build-log.txt%3A205)
<details>
<summary>all...</summary>

* _2026-04-21 22:10:16 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-kind-sriov/2046711709338439680#1:build-log.txt%3A196)
<details><summary>context</summary>
<pre>22:10:37: time=&#34;2026-04-21T22:10:37Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (2/3). Error: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway&#34;
22:10:38: time=&#34;2026-04-21T22:10:38Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (3/3). Error: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway&#34;
22:10:40: Error: unable to copy from source docker://quay.io/kubevirt/builder:2602201229-4494d1587: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway
make: *** [Makefile:39: bazel-build-images] Error 125
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-04-21 22:05:04 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-kind-1.35-vgpu/2046711709095170048#1:build-log.txt%3A189)
<details><summary>context</summary>
<pre>22:06:55: time=&#34;2026-04-21T22:06:55Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (2/3). Error: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway&#34;
22:06:56: time=&#34;2026-04-21T22:06:56Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (3/3). Error: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway&#34;
22:06:57: Error: unable to copy from source docker://quay.io/kubevirt/builder:2602201229-4494d1587: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway
make: *** [Makefile:39: bazel-build-images] Error 125
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-04-21 22:05:03 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-operator/2046711709980168192#1:build-log.txt%3A228)
<details><summary>context</summary>
<pre>22:07:10: time=&#34;2026-04-21T22:07:10Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (2/3). Error: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway&#34;
22:07:11: time=&#34;2026-04-21T22:07:11Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (3/3). Error: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway&#34;
22:07:15: Error: unable to copy from source docker://quay.io/kubevirt/builder:2602201229-4494d1587: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway
make: *** [Makefile:39: bazel-build-images] Error 125
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-04-21 22:05:03 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-network/2046711709594292224#1:build-log.txt%3A238)
<details><summary>context</summary>
<pre>22:06:54: time=&#34;2026-04-21T22:06:54Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (2/3). Error: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway&#34;
22:06:55: time=&#34;2026-04-21T22:06:55Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (3/3). Error: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway&#34;
22:07:01: Error: unable to copy from source docker://quay.io/kubevirt/builder:2602201229-4494d1587: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: Requesting bearer token: received unexpected HTTP status: 502 Bad Gateway
make: *** [Makefile:39: bazel-build-images] Error 125
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-compute/2046711709879504896#1:build-log.txt%3A205)
<details><summary>context</summary>
<pre>22:06:46: time=&#34;2026-04-21T22:06:46Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (2/3). Error: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway&#34;
22:06:47: time=&#34;2026-04-21T22:06:47Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (3/3). Error: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway&#34;
22:06:48: Error: unable to copy from source docker://quay.io/kubevirt/builder:2602201229-4494d1587: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway
make: *** [Makefile:39: bazel-build-images] Error 125
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-windows2016/2046711708981923840#1:build-log.txt%3A250)
<details><summary>context</summary>
<pre>22:07:02: time=&#34;2026-04-21T22:07:02Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (2/3). Error: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway&#34;
22:07:03: time=&#34;2026-04-21T22:07:03Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (3/3). Error: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway&#34;
22:07:05: Error: unable to copy from source docker://quay.io/kubevirt/builder:2602201229-4494d1587: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: Requesting bearer token: received unexpected HTTP status: 502 Bad Gateway
make: *** [Makefile:39: bazel-build-images] Error 125
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2046711709439102976#1:build-log.txt%3A238)
<details><summary>context</summary>
<pre>22:06:43: time=&#34;2026-04-21T22:06:43Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (2/3). Error: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway&#34;
22:06:45: time=&#34;2026-04-21T22:06:45Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (3/3). Error: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway&#34;
22:06:46: Error: unable to copy from source docker://quay.io/kubevirt/builder:2602201229-4494d1587: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway
make: *** [Makefile:39: bazel-build-images] Error 125
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> download failure in context (3x / 27.27%) </summary>

<hr/>

**1x**: _2026-04-21 22:05:03 &#43;0000 UTC_: <code>22:07:22: ERROR: Analysis of target &#39;//containerimages:alpine-with-test-tooling&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-storage/2046711709757870080#1:build-log.txt%3A569)
<details>
<summary>all...</summary>

* _2026-04-21 22:05:03 &#43;0000 UTC_: <code>22:07:22: ERROR: Analysis of target &#39;//containerimages:alpine-with-test-tooling&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-storage/2046711709757870080#1:build-log.txt%3A569)
<details><summary>context</summary>
<pre>22:07:22: Repository rule oci_alias defined at:
22:07:22:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_oci/oci/private/pull.bzl:472:28: in &lt;toplevel&gt;
22:07:22: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:103:10: //containerimages:alpine-with-test-tooling depends on @alpine_with_test_tooling//:alpine_with_test_tooling in repository @alpine_with_test_tooling which failed to fetch. no such package &#39;@alpine_with_test_tooling//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/auth?scope=repository:kubevirtci/alpine-with-test-tooling-container-disk:pull&amp;service=quay.io] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine_with_test_tooling/www-authenticate.json: GET returned 502 Bad Gateway
22:07:22: ERROR: Analysis of target &#39;//containerimages:alpine-with-test-tooling&#39; failed; build aborted: Analysis failed
22:07:22: INFO: Elapsed time: 4.444s
22:07:22: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>

**2x**: _2026-04-17 13:25:22 &#43;0000 UTC_: <code>13:37:52: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-k8s-1.34-windows2016/2045131369809448960#1:build-log.txt%3A4706)
<details>
<summary>all...</summary>

* _2026-04-17 13:25:22 &#43;0000 UTC_: <code>13:37:52: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-k8s-1.34-windows2016/2045131369809448960#1:build-log.txt%3A4706)
<details><summary>context</summary>
<pre>13:37:52: ERROR: Analysis of target &#39;//:buildifier&#39; failed; build aborted:
13:37:52: INFO: Elapsed time: 18.277s
13:37:52: INFO: 0 processes.
13:37:52: ERROR: Build failed. Not running target
make: *** [Makefile:28: bazel-build-functests] Error 1
&#43; ret=2
&#43; check_for_panics</pre>
</details>


* _2026-04-16 20:05:44 &#43;0000 UTC_: <code>20:18:08: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17437/pull-kubevirt-e2e-k8s-1.35-sig-compute/2044869730921091072#1:build-log.txt%3A4004)
<details><summary>context</summary>
<pre>20:18:08: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
20:18:08: INFO: Elapsed time: 18.375s
20:18:08: INFO: 0 processes.
20:18:08: ERROR: Build failed. Not running target
20:18:10: &#43; rm -f /tmp/kubevirt.deploy.YmnF
make: *** [Makefile:188: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 9.09%) </summary>

<hr/>

**1x**: _2026-04-17 13:33:58 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-kind-sriov/2045131370006581248#1:build-log.txt%3A1394)
<details>
<summary>all...</summary>

* _2026-04-17 13:33:58 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-kind-sriov/2045131370006581248#1:build-log.txt%3A1394)
<details><summary>context</summary>
<pre>13:42:41: Downloading cni-plugins-linux-amd64-v0.8.5.tgz
13:42:41: &#43;&#43; curl -sSL -o /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-sriov/cni-plugins-linux-amd64-v0.8.5.tgz https://github.com/containernetworking/plugins/releases/download/v0.8.5/cni-plugins-linux-amd64-v0.8.5.tgz
13:42:47: curl: (6) Could not resolve host: github.com
make: *** [Makefile:173: cluster-up] Error 6
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (11x / 100.00%)


#### external (11x / 100.00%)

<details>
<summary> container image pull failure in context (7x / 63.64%) </summary>

<hr/>

**7x**: _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-compute/2046711709879504896#1:build-log.txt%3A205)
<details>
<summary>all...</summary>

* _2026-04-21 22:10:16 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-kind-sriov/2046711709338439680#1:build-log.txt%3A196)

* _2026-04-21 22:05:04 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-kind-1.35-vgpu/2046711709095170048#1:build-log.txt%3A189)

* _2026-04-21 22:05:03 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-operator/2046711709980168192#1:build-log.txt%3A228)

* _2026-04-21 22:05:03 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-network/2046711709594292224#1:build-log.txt%3A238)

* _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-compute/2046711709879504896#1:build-log.txt%3A205)

* _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-windows2016/2046711708981923840#1:build-log.txt%3A250)

* _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2046711709439102976#1:build-log.txt%3A238)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (3x / 27.27%) </summary>

<hr/>

**1x**: _2026-04-21 22:05:03 &#43;0000 UTC_: <code>22:07:22: ERROR: Analysis of target &#39;//containerimages:alpine-with-test-tooling&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-storage/2046711709757870080#1:build-log.txt%3A569)
<details>
<summary>all...</summary>

* _2026-04-21 22:05:03 &#43;0000 UTC_: <code>22:07:22: ERROR: Analysis of target &#39;//containerimages:alpine-with-test-tooling&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-storage/2046711709757870080#1:build-log.txt%3A569)

</details>

<hr/>

**2x**: _2026-04-17 13:25:22 &#43;0000 UTC_: <code>13:37:52: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-k8s-1.34-windows2016/2045131369809448960#1:build-log.txt%3A4706)
<details>
<summary>all...</summary>

* _2026-04-17 13:25:22 &#43;0000 UTC_: <code>13:37:52: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-k8s-1.34-windows2016/2045131369809448960#1:build-log.txt%3A4706)

* _2026-04-16 20:05:44 &#43;0000 UTC_: <code>20:18:08: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17437/pull-kubevirt-e2e-k8s-1.35-sig-compute/2044869730921091072#1:build-log.txt%3A4004)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 9.09%) </summary>

<hr/>

**1x**: _2026-04-17 13:33:58 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-kind-sriov/2045131370006581248#1:build-log.txt%3A1394)
<details>
<summary>all...</summary>

* _2026-04-17 13:33:58 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-kind-sriov/2045131370006581248#1:build-log.txt%3A1394)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-storage (1x / 9.09%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-04-21 22:05:03 &#43;0000 UTC_: <code>22:07:22: ERROR: Analysis of target &#39;//containerimages:alpine-with-test-tooling&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-storage/2046711709757870080#1:build-log.txt%3A569)
<details>
<summary>all...</summary>

* _2026-04-21 22:05:03 &#43;0000 UTC_: <code>22:07:22: ERROR: Analysis of target &#39;//containerimages:alpine-with-test-tooling&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-storage/2046711709757870080#1:build-log.txt%3A569)

</details>

<hr/>
</details>

### sig-compute (6x / 54.55%)


#### external (6x / 100.00%)

<details>
<summary> container image pull failure in context (4x / 66.67%) </summary>

<hr/>

**4x**: _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-compute/2046711709879504896#1:build-log.txt%3A205)
<details>
<summary>all...</summary>

* _2026-04-21 22:05:04 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-kind-1.35-vgpu/2046711709095170048#1:build-log.txt%3A189)

* _2026-04-21 22:05:03 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-operator/2046711709980168192#1:build-log.txt%3A228)

* _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-compute/2046711709879504896#1:build-log.txt%3A205)

* _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-windows2016/2046711708981923840#1:build-log.txt%3A250)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (2x / 33.33%) </summary>

<hr/>

**2x**: _2026-04-17 13:25:22 &#43;0000 UTC_: <code>13:37:52: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-k8s-1.34-windows2016/2045131369809448960#1:build-log.txt%3A4706)
<details>
<summary>all...</summary>

* _2026-04-17 13:25:22 &#43;0000 UTC_: <code>13:37:52: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-k8s-1.34-windows2016/2045131369809448960#1:build-log.txt%3A4706)

* _2026-04-16 20:05:44 &#43;0000 UTC_: <code>20:18:08: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17437/pull-kubevirt-e2e-k8s-1.35-sig-compute/2044869730921091072#1:build-log.txt%3A4004)

</details>

<hr/>
</details>

### sig-network (4x / 36.36%)


#### external (4x / 100.00%)

<details>
<summary> container image pull failure in context (3x / 75.00%) </summary>

<hr/>

**3x**: _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2046711709439102976#1:build-log.txt%3A238)
<details>
<summary>all...</summary>

* _2026-04-21 22:10:16 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-kind-sriov/2046711709338439680#1:build-log.txt%3A196)

* _2026-04-21 22:05:03 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-network/2046711709594292224#1:build-log.txt%3A238)

* _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2046711709439102976#1:build-log.txt%3A238)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-04-17 13:33:58 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-kind-sriov/2045131370006581248#1:build-log.txt%3A1394)
<details>
<summary>all...</summary>

* _2026-04-17 13:33:58 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-kind-sriov/2045131370006581248#1:build-log.txt%3A1394)

</details>

<hr/>
</details>

Last updated: 2026-04-22 18:27:04
