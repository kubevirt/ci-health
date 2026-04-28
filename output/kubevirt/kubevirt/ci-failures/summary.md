



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-04-27 (1x / 9.09%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-04-27 07:18:06 &#43;0000 UTC_: <code>07:21:17: ERROR: Analysis of target &#39;//cmd/virt-operator:virt-operator-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17612/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2048662822543429632#1:build-log.txt%3A498)
<details>
<summary>all...</summary>

* _2026-04-27 07:18:06 &#43;0000 UTC_: <code>07:21:17: ERROR: Analysis of target &#39;//cmd/virt-operator:virt-operator-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17612/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2048662822543429632#1:build-log.txt%3A498)

</details>

<hr/>
</details>

### 2026-04-23 (2x / 18.18%)


#### needs-investigation (2x / 100.00%)

<details>
<summary> no matching pattern (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-04-23 12:45:40 &#43;0000 UTC_: <code>Error: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 root filesystem: deleting layer &#34;da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6 /var/lib/containers/storage/overlay/tempdirs/temp-dir-1832193173/1-da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17487/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2047295730619518976#1:build-log.txt%3A680)
<details>
<summary>all...</summary>

* _2026-04-23 12:47:46 &#43;0000 UTC_: <code>Error: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a root filesystem: deleting layer &#34;9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374 /var/lib/containers/storage/overlay/tempdirs/temp-dir-2010581549/1-9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17487/pull-kubevirt-e2e-k8s-1.34-sig-network/2047295730711793664#1:build-log.txt%3A682)

* _2026-04-23 12:45:40 &#43;0000 UTC_: <code>Error: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 root filesystem: deleting layer &#34;da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6 /var/lib/containers/storage/overlay/tempdirs/temp-dir-1832193173/1-da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17487/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2047295730619518976#1:build-log.txt%3A680)

</details>

<hr/>
</details>

### 2026-04-21 (8x / 72.73%)


#### external (8x / 100.00%)

<details>
<summary> container image pull failure in context (7x / 87.50%) </summary>

<hr/>

**7x**: _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2046711709439102976#1:build-log.txt%3A238)
<details>
<summary>all...</summary>

* _2026-04-21 22:10:16 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-kind-sriov/2046711709338439680#1:build-log.txt%3A196)

* _2026-04-21 22:05:04 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-kind-1.35-vgpu/2046711709095170048#1:build-log.txt%3A189)

* _2026-04-21 22:05:03 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-operator/2046711709980168192#1:build-log.txt%3A228)

* _2026-04-21 22:05:03 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-network/2046711709594292224#1:build-log.txt%3A238)

* _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2046711709439102976#1:build-log.txt%3A238)

* _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-compute/2046711709879504896#1:build-log.txt%3A205)

* _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-windows2016/2046711708981923840#1:build-log.txt%3A250)

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

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (9x / 81.82%)

<details>
<summary> container image pull failure in context (7x / 63.64%) </summary>

<hr/>

**7x**: _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2046711709439102976#1:build-log.txt%3A238)
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


</details>

<hr/>
</details>
<details>
<summary> download failure in context (2x / 18.18%) </summary>

<hr/>

**1x**: _2026-04-27 07:18:06 &#43;0000 UTC_: <code>07:21:17: ERROR: Analysis of target &#39;//cmd/virt-operator:virt-operator-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17612/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2048662822543429632#1:build-log.txt%3A498)
<details>
<summary>all...</summary>

* _2026-04-27 07:18:06 &#43;0000 UTC_: <code>07:21:17: ERROR: Analysis of target &#39;//cmd/virt-operator:virt-operator-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17612/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2048662822543429632#1:build-log.txt%3A498)
<details><summary>context</summary>
<pre>07:21:17: Repository rule oci_alias defined at:
07:21:17:   /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/rules_oci/oci/private/pull.bzl:417:28: in &lt;toplevel&gt;
07:21:17: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-operator/BUILD.bazel:47:10: //cmd/virt-operator:virt-operator-image depends on @zstd_linux_amd64//:zstd_toolchain in repository @zstd_linux_amd64 which failed to fetch. no such package &#39;@zstd_linux_amd64//&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/zstd-prebuilt/releases/download/v1.5.6/zstd_linux_amd64] to /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/zstd_linux_amd64/zstd: GET returned 502 Bad Gateway or Proxy Error
07:21:17: ERROR: Analysis of target &#39;//cmd/virt-operator:virt-operator-image&#39; failed; build aborted: Analysis failed
07:21:17: INFO: Elapsed time: 0.253s
07:21:17: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

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
</details>

### needs-investigation (2x / 18.18%)

<details>
<summary> no matching pattern (2x / 18.18%) </summary>

<hr/>

**2x**: _2026-04-23 12:45:40 &#43;0000 UTC_: <code>Error: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 root filesystem: deleting layer &#34;da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6 /var/lib/containers/storage/overlay/tempdirs/temp-dir-1832193173/1-da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17487/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2047295730619518976#1:build-log.txt%3A680)
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

<a id="per-branch"></a>

## per branch [⬆](#top)


### release-1.8 (1x / 9.09%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-04-27 07:18:06 &#43;0000 UTC_: <code>07:21:17: ERROR: Analysis of target &#39;//cmd/virt-operator:virt-operator-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17612/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2048662822543429632#1:build-log.txt%3A498)
<details>
<summary>all...</summary>

* _2026-04-27 07:18:06 &#43;0000 UTC_: <code>07:21:17: ERROR: Analysis of target &#39;//cmd/virt-operator:virt-operator-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17612/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2048662822543429632#1:build-log.txt%3A498)

</details>

<hr/>
</details>

### main (10x / 90.91%)


#### external (8x / 80.00%)

<details>
<summary> container image pull failure in context (7x / 70.00%) </summary>

<hr/>

**7x**: _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2046711709439102976#1:build-log.txt%3A238)
<details>
<summary>all...</summary>

* _2026-04-21 22:10:16 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-kind-sriov/2046711709338439680#1:build-log.txt%3A196)

* _2026-04-21 22:05:04 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-kind-1.35-vgpu/2046711709095170048#1:build-log.txt%3A189)

* _2026-04-21 22:05:03 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-operator/2046711709980168192#1:build-log.txt%3A228)

* _2026-04-21 22:05:03 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-network/2046711709594292224#1:build-log.txt%3A238)

* _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2046711709439102976#1:build-log.txt%3A238)

* _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-sig-compute/2046711709879504896#1:build-log.txt%3A205)

* _2026-04-21 22:05:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17555/pull-kubevirt-e2e-k8s-1.34-windows2016/2046711708981923840#1:build-log.txt%3A250)

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

**2x**: _2026-04-23 12:45:40 &#43;0000 UTC_: <code>Error: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 root filesystem: deleting layer &#34;da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6 /var/lib/containers/storage/overlay/tempdirs/temp-dir-1832193173/1-da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17487/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2047295730619518976#1:build-log.txt%3A680)
<details>
<summary>all...</summary>

* _2026-04-23 12:47:46 &#43;0000 UTC_: <code>Error: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a root filesystem: deleting layer &#34;9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374 /var/lib/containers/storage/overlay/tempdirs/temp-dir-2010581549/1-9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17487/pull-kubevirt-e2e-k8s-1.34-sig-network/2047295730711793664#1:build-log.txt%3A682)

* _2026-04-23 12:45:40 &#43;0000 UTC_: <code>Error: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 root filesystem: deleting layer &#34;da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6 /var/lib/containers/storage/overlay/tempdirs/temp-dir-1832193173/1-da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17487/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2047295730619518976#1:build-log.txt%3A680)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (5x / 45.45%)


#### external (5x / 100.00%)

<details>
<summary> container image pull failure in context (4x / 80.00%) </summary>

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
<summary> download failure in context (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-04-27 07:18:06 &#43;0000 UTC_: <code>07:21:17: ERROR: Analysis of target &#39;//cmd/virt-operator:virt-operator-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17612/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2048662822543429632#1:build-log.txt%3A498)
<details>
<summary>all...</summary>

* _2026-04-27 07:18:06 &#43;0000 UTC_: <code>07:21:17: ERROR: Analysis of target &#39;//cmd/virt-operator:virt-operator-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17612/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2048662822543429632#1:build-log.txt%3A498)

</details>

<hr/>
</details>

### sig-network (5x / 45.45%)


#### external (3x / 60.00%)

<details>
<summary> container image pull failure in context (3x / 60.00%) </summary>

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

#### needs-investigation (2x / 40.00%)

<details>
<summary> no matching pattern (2x / 40.00%) </summary>

<hr/>

**2x**: _2026-04-23 12:45:40 &#43;0000 UTC_: <code>Error: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 root filesystem: deleting layer &#34;da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6 /var/lib/containers/storage/overlay/tempdirs/temp-dir-1832193173/1-da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17487/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2047295730619518976#1:build-log.txt%3A680)
<details>
<summary>all...</summary>

* _2026-04-23 12:47:46 &#43;0000 UTC_: <code>Error: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: cleaning up container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a storage: unmounting container b28c55058e1707937637487dc58dfa2ec9d99b2c01efd3a96edffc155e28a92a root filesystem: deleting layer &#34;9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374 /var/lib/containers/storage/overlay/tempdirs/temp-dir-2010581549/1-9efad812e38cafd4b6973fe4c9c995cf9877578d2da6e132787fa52ebbde6374: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17487/pull-kubevirt-e2e-k8s-1.34-sig-network/2047295730711793664#1:build-log.txt%3A682)

* _2026-04-23 12:45:40 &#43;0000 UTC_: <code>Error: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: cleaning up container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 storage: unmounting container 8745183952e1b6c794d16f9a7a8f87b3bd22ddab239a8e62108c4b136234b604 root filesystem: deleting layer &#34;da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6 /var/lib/containers/storage/overlay/tempdirs/temp-dir-1832193173/1-da67bdf77bd503357e818379be44cece56761cec88d0031f8f55dbe87656a6a6: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17487/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2047295730619518976#1:build-log.txt%3A680)

</details>

<hr/>
</details>

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

Last updated: 2026-04-28 12:57:41
