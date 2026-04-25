



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-04-23 (2x / 12.50%)


#### needs-investigation (2x / 100.00%)

<details>
<summary> no matching pattern (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-04-23 12:47:46 &#43;0000 UTC_: <code>Error: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a root filesystem: deleting layer &#34;9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374 /var/lib/containers/storage/overlay/tempdirs/temp-dir-2010581549/1-9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17487/pull-kubevirt-e2e-k8s-1.34-sig-network/2047295730711793664#1:build-log.txt%3A682)
<details>
<summary>all...</summary>

* _2026-04-23 12:47:46 &#43;0000 UTC_: <code>Error: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a root filesystem: deleting layer &#34;9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374 /var/lib/containers/storage/overlay/tempdirs/temp-dir-2010581549/1-9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17487/pull-kubevirt-e2e-k8s-1.34-sig-network/2047295730711793664#1:build-log.txt%3A682)

* _2026-04-23 12:45:40 &#43;0000 UTC_: <code>Error: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 root filesystem: deleting layer &#34;da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6 /var/lib/containers/storage/overlay/tempdirs/temp-dir-1832193173/1-da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17487/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2047295730619518976#1:build-log.txt%3A680)

</details>

<hr/>
</details>

### 2026-04-21 (8x / 50.00%)


#### external (8x / 100.00%)

<details>
<summary> container image pull failure in context (7x / 87.50%) </summary>

<hr/>

**7x**: _2026-04-21 22:05:03 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-network/2046711709594292224#1:build-log.txt%3A238)
<details>
<summary>all...</summary>

* _2026-04-21 22:10:16 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-kind-sriov/2046711709338439680#1:build-log.txt%3A196)

* _2026-04-21 22:05:04 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-kind-1.35-vgpu/2046711709095170048#1:build-log.txt%3A189)

* _2026-04-21 22:05:03 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-network/2046711709594292224#1:build-log.txt%3A238)

* _2026-04-21 22:05:03 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-operator/2046711709980168192#1:build-log.txt%3A228)

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

### 2026-04-18 (6x / 37.50%)


#### external (5x / 83.33%)

<details>
<summary> download failure in context (5x / 83.33%) </summary>

<hr/>

**5x**: _2026-04-18 17:50:18 &#43;0000 UTC_: <code>18:23:44: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-k8s-1.31-sig-network-1.6/2045560420638396416#1:build-log.txt%3A2734)
<details>
<summary>all...</summary>

* _2026-04-18 17:50:20 &#43;0000 UTC_: <code>18:24:04: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-k8s-1.32-sig-operator-1.6/2045560426552365056#1:build-log.txt%3A2670)

* _2026-04-18 17:50:19 &#43;0000 UTC_: <code>18:28:57: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.6/2045560430264324096#1:build-log.txt%3A2580)

* _2026-04-18 17:50:18 &#43;0000 UTC_: <code>18:23:44: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-k8s-1.31-sig-network-1.6/2045560420638396416#1:build-log.txt%3A2734)

* _2026-04-18 17:50:17 &#43;0000 UTC_: <code>18:15:31: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-k8s-1.31-sig-operator-1.6/2045560423167561728#1:build-log.txt%3A2611)

* _2026-04-18 17:50:13 &#43;0000 UTC_: <code>17:52:35: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-kind-1.30-vgpu-1.6/2045560409280221184#1:build-log.txt%3A450)

</details>

<hr/>
</details>

#### internal (1x / 16.67%)

<details>
<summary> make cluster lifecycle target failure (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-04-18 17:56:00 &#43;0000 UTC_: <code>make: *** [Makefile:166: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-kind-sriov-1.6/2045560409401856000#1:build-log.txt%3A1338)
<details>
<summary>all...</summary>

* _2026-04-18 17:56:00 &#43;0000 UTC_: <code>make: *** [Makefile:166: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-kind-sriov-1.6/2045560409401856000#1:build-log.txt%3A1338)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (13x / 81.25%)

<details>
<summary> container image pull failure in context (7x / 43.75%) </summary>

<hr/>

**7x**: _2026-04-21 22:05:03 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-network/2046711709594292224#1:build-log.txt%3A238)
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
<summary> download failure in context (6x / 37.50%) </summary>

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

**5x**: _2026-04-18 17:50:18 &#43;0000 UTC_: <code>18:23:44: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-k8s-1.31-sig-network-1.6/2045560420638396416#1:build-log.txt%3A2734)
<details>
<summary>all...</summary>

* _2026-04-18 17:50:20 &#43;0000 UTC_: <code>18:24:04: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-k8s-1.32-sig-operator-1.6/2045560426552365056#1:build-log.txt%3A2670)
<details><summary>context</summary>
<pre>18:24:04: ERROR: Analysis of target &#39;//:buildifier&#39; failed; build aborted:
18:24:04: INFO: Elapsed time: 24.850s
18:24:04: INFO: 0 processes.
18:24:04: ERROR: Build failed. Not running target
make: *** [Makefile:26: bazel-build-functests] Error 1
&#43; ret=2
&#43; make cluster-down</pre>
</details>


* _2026-04-18 17:50:19 &#43;0000 UTC_: <code>18:28:57: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.6/2045560430264324096#1:build-log.txt%3A2580)
<details><summary>context</summary>
<pre>18:28:57: ERROR: Analysis of target &#39;//:buildifier&#39; failed; build aborted:
18:28:57: INFO: Elapsed time: 25.243s
18:28:57: INFO: 0 processes.
18:28:57: ERROR: Build failed. Not running target
make: *** [Makefile:26: bazel-build-functests] Error 1
&#43; ret=2
&#43; make cluster-down</pre>
</details>


* _2026-04-18 17:50:18 &#43;0000 UTC_: <code>18:23:44: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-k8s-1.31-sig-network-1.6/2045560420638396416#1:build-log.txt%3A2734)
<details><summary>context</summary>
<pre>18:23:44: ERROR: Analysis of target &#39;//:buildifier&#39; failed; build aborted:
18:23:44: INFO: Elapsed time: 24.581s
18:23:44: INFO: 0 processes.
18:23:44: ERROR: Build failed. Not running target
make: *** [Makefile:26: bazel-build-functests] Error 1
&#43; ret=2
&#43; make cluster-down</pre>
</details>


* _2026-04-18 17:50:17 &#43;0000 UTC_: <code>18:15:31: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-k8s-1.31-sig-operator-1.6/2045560423167561728#1:build-log.txt%3A2611)
<details><summary>context</summary>
<pre>18:15:31: ERROR: Analysis of target &#39;//:buildifier&#39; failed; build aborted:
18:15:31: INFO: Elapsed time: 24.557s
18:15:31: INFO: 0 processes.
18:15:31: ERROR: Build failed. Not running target
make: *** [Makefile:26: bazel-build-functests] Error 1
&#43; ret=2
&#43; make cluster-down</pre>
</details>


* _2026-04-18 17:50:13 &#43;0000 UTC_: <code>17:52:35: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-kind-1.30-vgpu-1.6/2045560409280221184#1:build-log.txt%3A450)
<details><summary>context</summary>
<pre>17:52:35: ERROR: Analysis of target &#39;//rpm:sandbox_x86_64&#39; failed; build aborted:
17:52:35: INFO: Elapsed time: 13.524s
17:52:35: INFO: 0 processes.
17:52:35: ERROR: Build failed. Not running target
make: *** [Makefile:40: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>


</details>

<hr/>
</details>

### needs-investigation (2x / 12.50%)

<details>
<summary> no matching pattern (2x / 12.50%) </summary>

<hr/>

**2x**: _2026-04-23 12:47:46 &#43;0000 UTC_: <code>Error: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a root filesystem: deleting layer &#34;9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374 /var/lib/containers/storage/overlay/tempdirs/temp-dir-2010581549/1-9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17487/pull-kubevirt-e2e-k8s-1.34-sig-network/2047295730711793664#1:build-log.txt%3A682)
<details>
<summary>all...</summary>

* _2026-04-23 12:47:46 &#43;0000 UTC_: <code>Error: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a root filesystem: deleting layer &#34;9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374 /var/lib/containers/storage/overlay/tempdirs/temp-dir-2010581549/1-9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17487/pull-kubevirt-e2e-k8s-1.34-sig-network/2047295730711793664#1:build-log.txt%3A682)
<details><summary>context</summary>
<pre>time=&#34;2026-04-23T13:27:44Z&#34; level=warning msg=&#34;Found incomplete layer \&#34;9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374\&#34;, deleting it&#34;
time=&#34;2026-04-23T13:27:44Z&#34; level=warning msg=&#34;Found incomplete layer \&#34;9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374\&#34;, deleting it&#34;
time=&#34;2026-04-23T13:27:44Z&#34; level=error msg=&#34;cleaning up storage: removing container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a root filesystem: deleting layer \&#34;9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374\&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374 /var/lib/containers/storage/overlay/tempdirs/temp-dir-2421049645/1-9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374: invalid cross-device link&#34;
Error: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a root filesystem: deleting layer &#34;9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374 /var/lib/containers/storage/overlay/tempdirs/temp-dir-2010581549/1-9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374: invalid cross-device link
time=&#34;2026-04-23T13:27:44Z&#34; level=warning msg=&#34;Found incomplete layer \&#34;9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374\&#34;, deleting it&#34;
/usr/local/bin/runner.sh: line 50: wait: pid 1220 is not a child of this shell
================================================================================</pre>
</details>


* _2026-04-23 12:45:40 &#43;0000 UTC_: <code>Error: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 root filesystem: deleting layer &#34;da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6 /var/lib/containers/storage/overlay/tempdirs/temp-dir-1832193173/1-da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17487/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2047295730619518976#1:build-log.txt%3A680)
<details><summary>context</summary>
<pre>time=&#34;2026-04-23T13:29:20Z&#34; level=warning msg=&#34;Found incomplete layer \&#34;da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6\&#34;, deleting it&#34;
time=&#34;2026-04-23T13:29:21Z&#34; level=warning msg=&#34;Found incomplete layer \&#34;da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6\&#34;, deleting it&#34;
time=&#34;2026-04-23T13:29:21Z&#34; level=error msg=&#34;cleaning up storage: removing container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 root filesystem: deleting layer \&#34;da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6\&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6 /var/lib/containers/storage/overlay/tempdirs/temp-dir-3197161821/1-da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6: invalid cross-device link&#34;
Error: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 root filesystem: deleting layer &#34;da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6 /var/lib/containers/storage/overlay/tempdirs/temp-dir-1832193173/1-da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6: invalid cross-device link
time=&#34;2026-04-23T13:29:21Z&#34; level=warning msg=&#34;Found incomplete layer \&#34;da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6\&#34;, deleting it&#34;
/usr/local/bin/runner.sh: line 50: wait: pid 1220 is not a child of this shell
================================================================================</pre>
</details>


</details>

<hr/>
</details>

### internal (1x / 6.25%)

<details>
<summary> make cluster lifecycle target failure (1x / 6.25%) </summary>

<hr/>

**1x**: _2026-04-18 17:56:00 &#43;0000 UTC_: <code>make: *** [Makefile:166: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-kind-sriov-1.6/2045560409401856000#1:build-log.txt%3A1338)
<details>
<summary>all...</summary>

* _2026-04-18 17:56:00 &#43;0000 UTC_: <code>make: *** [Makefile:166: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-kind-sriov-1.6/2045560409401856000#1:build-log.txt%3A1338)
<details><summary>context</summary>
<pre>18:05:33: gzip: stdin: not in gzip format
18:05:33: tar: Child returned status 1
18:05:33: tar: Error is not recoverable: exiting now
make: *** [Makefile:166: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (10x / 62.50%)


#### external (8x / 80.00%)

<details>
<summary> container image pull failure in context (7x / 70.00%) </summary>

<hr/>

**7x**: _2026-04-21 22:05:03 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-network/2046711709594292224#1:build-log.txt%3A238)
<details>
<summary>all...</summary>

* _2026-04-21 22:10:16 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-kind-sriov/2046711709338439680#1:build-log.txt%3A196)

* _2026-04-21 22:05:04 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-kind-1.35-vgpu/2046711709095170048#1:build-log.txt%3A189)

* _2026-04-21 22:05:03 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-network/2046711709594292224#1:build-log.txt%3A238)

* _2026-04-21 22:05:03 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-operator/2046711709980168192#1:build-log.txt%3A228)

* _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-compute/2046711709879504896#1:build-log.txt%3A205)

* _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-windows2016/2046711708981923840#1:build-log.txt%3A250)

* _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2046711709439102976#1:build-log.txt%3A238)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 10.00%) </summary>

<hr/>

**1x**: _2026-04-21 22:05:03 &#43;0000 UTC_: <code>22:07:22: ERROR: Analysis of target &#39;//containerimages:alpine-with-test-tooling&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-storage/2046711709757870080#1:build-log.txt%3A569)
<details>
<summary>all...</summary>

* _2026-04-21 22:05:03 &#43;0000 UTC_: <code>22:07:22: ERROR: Analysis of target &#39;//containerimages:alpine-with-test-tooling&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-storage/2046711709757870080#1:build-log.txt%3A569)

</details>

<hr/>
</details>

#### needs-investigation (2x / 20.00%)

<details>
<summary> no matching pattern (2x / 20.00%) </summary>

<hr/>

**2x**: _2026-04-23 12:47:46 &#43;0000 UTC_: <code>Error: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a root filesystem: deleting layer &#34;9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374 /var/lib/containers/storage/overlay/tempdirs/temp-dir-2010581549/1-9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17487/pull-kubevirt-e2e-k8s-1.34-sig-network/2047295730711793664#1:build-log.txt%3A682)
<details>
<summary>all...</summary>

* _2026-04-23 12:47:46 &#43;0000 UTC_: <code>Error: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a root filesystem: deleting layer &#34;9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374 /var/lib/containers/storage/overlay/tempdirs/temp-dir-2010581549/1-9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17487/pull-kubevirt-e2e-k8s-1.34-sig-network/2047295730711793664#1:build-log.txt%3A682)

* _2026-04-23 12:45:40 &#43;0000 UTC_: <code>Error: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 root filesystem: deleting layer &#34;da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6 /var/lib/containers/storage/overlay/tempdirs/temp-dir-1832193173/1-da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17487/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2047295730619518976#1:build-log.txt%3A680)

</details>

<hr/>
</details>

### release-1.6 (6x / 37.50%)


#### internal (1x / 16.67%)

<details>
<summary> make cluster lifecycle target failure (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-04-18 17:56:00 &#43;0000 UTC_: <code>make: *** [Makefile:166: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-kind-sriov-1.6/2045560409401856000#1:build-log.txt%3A1338)
<details>
<summary>all...</summary>

* _2026-04-18 17:56:00 &#43;0000 UTC_: <code>make: *** [Makefile:166: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-kind-sriov-1.6/2045560409401856000#1:build-log.txt%3A1338)

</details>

<hr/>
</details>

#### external (5x / 83.33%)

<details>
<summary> download failure in context (5x / 83.33%) </summary>

<hr/>

**5x**: _2026-04-18 17:50:18 &#43;0000 UTC_: <code>18:23:44: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-k8s-1.31-sig-network-1.6/2045560420638396416#1:build-log.txt%3A2734)
<details>
<summary>all...</summary>

* _2026-04-18 17:50:20 &#43;0000 UTC_: <code>18:24:04: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-k8s-1.32-sig-operator-1.6/2045560426552365056#1:build-log.txt%3A2670)

* _2026-04-18 17:50:19 &#43;0000 UTC_: <code>18:28:57: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.6/2045560430264324096#1:build-log.txt%3A2580)

* _2026-04-18 17:50:18 &#43;0000 UTC_: <code>18:23:44: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-k8s-1.31-sig-network-1.6/2045560420638396416#1:build-log.txt%3A2734)

* _2026-04-18 17:50:17 &#43;0000 UTC_: <code>18:15:31: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-k8s-1.31-sig-operator-1.6/2045560423167561728#1:build-log.txt%3A2611)

* _2026-04-18 17:50:13 &#43;0000 UTC_: <code>17:52:35: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-kind-1.30-vgpu-1.6/2045560409280221184#1:build-log.txt%3A450)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-network (7x / 43.75%)


#### external (4x / 57.14%)

<details>
<summary> container image pull failure in context (3x / 42.86%) </summary>

<hr/>

**3x**: _2026-04-21 22:05:03 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-network/2046711709594292224#1:build-log.txt%3A238)
<details>
<summary>all...</summary>

* _2026-04-21 22:10:16 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-kind-sriov/2046711709338439680#1:build-log.txt%3A196)

* _2026-04-21 22:05:03 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-network/2046711709594292224#1:build-log.txt%3A238)

* _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2046711709439102976#1:build-log.txt%3A238)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-04-18 17:50:18 &#43;0000 UTC_: <code>18:23:44: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-k8s-1.31-sig-network-1.6/2045560420638396416#1:build-log.txt%3A2734)
<details>
<summary>all...</summary>

* _2026-04-18 17:50:18 &#43;0000 UTC_: <code>18:23:44: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-k8s-1.31-sig-network-1.6/2045560420638396416#1:build-log.txt%3A2734)

</details>

<hr/>
</details>

#### needs-investigation (2x / 28.57%)

<details>
<summary> no matching pattern (2x / 28.57%) </summary>

<hr/>

**2x**: _2026-04-23 12:47:46 &#43;0000 UTC_: <code>Error: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a root filesystem: deleting layer &#34;9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374 /var/lib/containers/storage/overlay/tempdirs/temp-dir-2010581549/1-9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17487/pull-kubevirt-e2e-k8s-1.34-sig-network/2047295730711793664#1:build-log.txt%3A682)
<details>
<summary>all...</summary>

* _2026-04-23 12:47:46 &#43;0000 UTC_: <code>Error: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a root filesystem: deleting layer &#34;9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374 /var/lib/containers/storage/overlay/tempdirs/temp-dir-2010581549/1-9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17487/pull-kubevirt-e2e-k8s-1.34-sig-network/2047295730711793664#1:build-log.txt%3A682)

* _2026-04-23 12:45:40 &#43;0000 UTC_: <code>Error: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 root filesystem: deleting layer &#34;da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6 /var/lib/containers/storage/overlay/tempdirs/temp-dir-1832193173/1-da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17487/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2047295730619518976#1:build-log.txt%3A680)

</details>

<hr/>
</details>

#### internal (1x / 14.29%)

<details>
<summary> make cluster lifecycle target failure (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-04-18 17:56:00 &#43;0000 UTC_: <code>make: *** [Makefile:166: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-kind-sriov-1.6/2045560409401856000#1:build-log.txt%3A1338)
<details>
<summary>all...</summary>

* _2026-04-18 17:56:00 &#43;0000 UTC_: <code>make: *** [Makefile:166: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-kind-sriov-1.6/2045560409401856000#1:build-log.txt%3A1338)

</details>

<hr/>
</details>

### sig-storage (1x / 6.25%)


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

### sig-compute (8x / 50.00%)


#### external (8x / 100.00%)

<details>
<summary> download failure in context (4x / 50.00%) </summary>

<hr/>

**4x**: _2026-04-18 17:50:17 &#43;0000 UTC_: <code>18:15:31: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-k8s-1.31-sig-operator-1.6/2045560423167561728#1:build-log.txt%3A2611)
<details>
<summary>all...</summary>

* _2026-04-18 17:50:20 &#43;0000 UTC_: <code>18:24:04: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-k8s-1.32-sig-operator-1.6/2045560426552365056#1:build-log.txt%3A2670)

* _2026-04-18 17:50:19 &#43;0000 UTC_: <code>18:28:57: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.6/2045560430264324096#1:build-log.txt%3A2580)

* _2026-04-18 17:50:17 &#43;0000 UTC_: <code>18:15:31: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-k8s-1.31-sig-operator-1.6/2045560423167561728#1:build-log.txt%3A2611)

* _2026-04-18 17:50:13 &#43;0000 UTC_: <code>17:52:35: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17529/pull-kubevirt-e2e-kind-1.30-vgpu-1.6/2045560409280221184#1:build-log.txt%3A450)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (4x / 50.00%) </summary>

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

Last updated: 2026-04-25 15:19:54
