



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-03-31 (1x / 25.00%)


#### external (1x / 100.00%)

<details>
<summary> failed external fetch in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-03-31 21:04:34 &#43;0000 UTC_: <code>21:07:21: ERROR: Analysis of target &#39;//cmd/virt-handler:virt-handler-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17315/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2039070412440080384#1:build-log.txt%3A521)
<details>
<summary>all...</summary>

* _2026-03-31 21:04:34 &#43;0000 UTC_: <code>21:07:21: ERROR: Analysis of target &#39;//cmd/virt-handler:virt-handler-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17315/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2039070412440080384#1:build-log.txt%3A521)

</details>

<hr/>
</details>

### 2026-03-30 (3x / 75.00%)


#### external (3x / 100.00%)

<details>
<summary> container image pull failure in context (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-03-30 20:19:48 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17344/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2038608441399840768#1:build-log.txt%3A223)
<details>
<summary>all...</summary>

* _2026-03-30 20:19:50 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17344/pull-kubevirt-e2e-k8s-1.35-sig-network/2038608442087706624#1:build-log.txt%3A237)

* _2026-03-30 20:19:48 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17344/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2038608441399840768#1:build-log.txt%3A223)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-03-30 19:49:00 &#43;0000 UTC_: <code>20:02:40: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17344/pull-kubevirt-e2e-k8s-1.35-sig-storage/2038608442943344640#1:build-log.txt%3A4148)
<details>
<summary>all...</summary>

* _2026-03-30 19:49:00 &#43;0000 UTC_: <code>20:02:40: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17344/pull-kubevirt-e2e-k8s-1.35-sig-storage/2038608442943344640#1:build-log.txt%3A4148)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (4x / 100.00%)

<details>
<summary> container image pull failure in context (2x / 50.00%) </summary>

<hr/>

**2x**: _2026-03-30 20:19:48 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17344/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2038608441399840768#1:build-log.txt%3A223)
<details>
<summary>all...</summary>

* _2026-03-30 20:19:50 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17344/pull-kubevirt-e2e-k8s-1.35-sig-network/2038608442087706624#1:build-log.txt%3A237)
<details><summary>context</summary>
<pre>20:20:50: time=&#34;2026-03-30T20:20:50Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (2/3). Error: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: Requesting bearer token: received unexpected HTTP status: 502 Bad Gateway&#34;
20:20:59: time=&#34;2026-03-30T20:20:59Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (3/3). Error: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: Requesting bearer token: received unexpected HTTP status: 502 Bad Gateway&#34;
20:21:01: Error: unable to copy from source docker://quay.io/kubevirt/builder:2602201229-4494d1587: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway
make: *** [Makefile:39: bazel-build-images] Error 125
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-03-30 20:19:48 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17344/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2038608441399840768#1:build-log.txt%3A223)
<details><summary>context</summary>
<pre>20:21:20: time=&#34;2026-03-30T20:21:20Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (2/3). Error: copying system image from manifest list: determining manifest MIME type for docker://quay.io/kubevirt/builder:2602201229-4494d1587: reading manifest sha256:4cf5a164d0818ff8a3f05be1006bb2dd503e4f5d24c91df9cad845b3f9866102 in quay.io/kubevirt/builder: received unexpected HTTP status: 502 Bad Gateway&#34;
20:21:28: time=&#34;2026-03-30T20:21:28Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (3/3). Error: copying system image from manifest list: determining manifest MIME type for docker://quay.io/kubevirt/builder:2602201229-4494d1587: reading manifest sha256:4cf5a164d0818ff8a3f05be1006bb2dd503e4f5d24c91df9cad845b3f9866102 in quay.io/kubevirt/builder: received unexpected HTTP status: 502 Bad Gateway&#34;
20:21:29: Error: unable to copy from source docker://quay.io/kubevirt/builder:2602201229-4494d1587: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway
make: *** [Makefile:39: bazel-build-images] Error 125
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-03-30 19:49:00 &#43;0000 UTC_: <code>20:02:40: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17344/pull-kubevirt-e2e-k8s-1.35-sig-storage/2038608442943344640#1:build-log.txt%3A4148)
<details>
<summary>all...</summary>

* _2026-03-30 19:49:00 &#43;0000 UTC_: <code>20:02:40: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17344/pull-kubevirt-e2e-k8s-1.35-sig-storage/2038608442943344640#1:build-log.txt%3A4148)
<details><summary>context</summary>
<pre>20:02:40: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
20:02:40: INFO: Elapsed time: 0.474s
20:02:40: INFO: 0 processes.
20:02:40: ERROR: Build failed. Not running target
20:02:42: &#43; rm -f /tmp/kubevirt.deploy.IbSk
make: *** [Makefile:188: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> failed external fetch in context (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-03-31 21:04:34 &#43;0000 UTC_: <code>21:07:21: ERROR: Analysis of target &#39;//cmd/virt-handler:virt-handler-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17315/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2039070412440080384#1:build-log.txt%3A521)
<details>
<summary>all...</summary>

* _2026-03-31 21:04:34 &#43;0000 UTC_: <code>21:07:21: ERROR: Analysis of target &#39;//cmd/virt-handler:virt-handler-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17315/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2039070412440080384#1:build-log.txt%3A521)
<details><summary>context</summary>
<pre>21:07:21: *) If there is a configured URL Rewriter, check that it does not block the request.
21:07:21:
21:07:21: See for more: https://github.com/bazel-contrib/rules_oci/blob/main/docs/pull.md#authentication-using-credential-helpers
21:07:21: ERROR: Analysis of target &#39;//cmd/virt-handler:virt-handler-image&#39; failed; build aborted:
21:07:21: INFO: Elapsed time: 6.558s
21:07:21: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (4x / 100.00%)


#### external (4x / 100.00%)

<details>
<summary> container image pull failure in context (2x / 50.00%) </summary>

<hr/>

**2x**: _2026-03-30 20:19:48 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17344/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2038608441399840768#1:build-log.txt%3A223)
<details>
<summary>all...</summary>

* _2026-03-30 20:19:50 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17344/pull-kubevirt-e2e-k8s-1.35-sig-network/2038608442087706624#1:build-log.txt%3A237)

* _2026-03-30 20:19:48 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17344/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2038608441399840768#1:build-log.txt%3A223)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-03-30 19:49:00 &#43;0000 UTC_: <code>20:02:40: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17344/pull-kubevirt-e2e-k8s-1.35-sig-storage/2038608442943344640#1:build-log.txt%3A4148)
<details>
<summary>all...</summary>

* _2026-03-30 19:49:00 &#43;0000 UTC_: <code>20:02:40: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17344/pull-kubevirt-e2e-k8s-1.35-sig-storage/2038608442943344640#1:build-log.txt%3A4148)

</details>

<hr/>
</details>
<details>
<summary> failed external fetch in context (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-03-31 21:04:34 &#43;0000 UTC_: <code>21:07:21: ERROR: Analysis of target &#39;//cmd/virt-handler:virt-handler-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17315/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2039070412440080384#1:build-log.txt%3A521)
<details>
<summary>all...</summary>

* _2026-03-31 21:04:34 &#43;0000 UTC_: <code>21:07:21: ERROR: Analysis of target &#39;//cmd/virt-handler:virt-handler-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17315/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2039070412440080384#1:build-log.txt%3A521)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (2x / 50.00%)


#### external (2x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-03-30 20:19:48 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17344/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2038608441399840768#1:build-log.txt%3A223)
<details>
<summary>all...</summary>

* _2026-03-30 20:19:48 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17344/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2038608441399840768#1:build-log.txt%3A223)

</details>

<hr/>
</details>
<details>
<summary> failed external fetch in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-03-31 21:04:34 &#43;0000 UTC_: <code>21:07:21: ERROR: Analysis of target &#39;//cmd/virt-handler:virt-handler-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17315/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2039070412440080384#1:build-log.txt%3A521)
<details>
<summary>all...</summary>

* _2026-03-31 21:04:34 &#43;0000 UTC_: <code>21:07:21: ERROR: Analysis of target &#39;//cmd/virt-handler:virt-handler-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17315/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2039070412440080384#1:build-log.txt%3A521)

</details>

<hr/>
</details>

### sig-network (1x / 25.00%)


#### external (1x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-03-30 20:19:50 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17344/pull-kubevirt-e2e-k8s-1.35-sig-network/2038608442087706624#1:build-log.txt%3A237)
<details>
<summary>all...</summary>

* _2026-03-30 20:19:50 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17344/pull-kubevirt-e2e-k8s-1.35-sig-network/2038608442087706624#1:build-log.txt%3A237)

</details>

<hr/>
</details>

### sig-storage (1x / 25.00%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-03-30 19:49:00 &#43;0000 UTC_: <code>20:02:40: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17344/pull-kubevirt-e2e-k8s-1.35-sig-storage/2038608442943344640#1:build-log.txt%3A4148)
<details>
<summary>all...</summary>

* _2026-03-30 19:49:00 &#43;0000 UTC_: <code>20:02:40: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17344/pull-kubevirt-e2e-k8s-1.35-sig-storage/2038608442943344640#1:build-log.txt%3A4148)

</details>

<hr/>
</details>

Last updated: 2026-04-06 15:24:24
