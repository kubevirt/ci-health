



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-05-13 (1x / 20.00%)


#### needs-investigation (1x / 100.00%)

<details>
<summary> no error snippets (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-05-13 12:02:10 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17794/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054532503019982848)
<details>
<summary>all...</summary>

* _2026-05-13 12:02:10 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17794/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054532503019982848)

</details>

<hr/>
</details>

### 2026-05-10 (2x / 40.00%)


#### external (2x / 100.00%)

<details>
<summary> download failure in context (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-05-10 14:21:04 &#43;0000 UTC_: <code>14:33:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17611/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2053480299794272256#1:build-log.txt%3A4762)
<details>
<summary>all...</summary>

* _2026-05-10 14:21:04 &#43;0000 UTC_: <code>14:33:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17611/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2053480299794272256#1:build-log.txt%3A4762)

* _2026-05-10 07:10:19 &#43;0000 UTC_: <code>07:11:16: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17746/pull-kubevirt-e2e-k8s-1.35-sig-compute/2053371905993347072#1:build-log.txt%3A324)

</details>

<hr/>
</details>

### 2026-05-07 (2x / 40.00%)


#### external (2x / 100.00%)

<details>
<summary> download failure in context (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-05-07 18:48:03 &#43;0000 UTC_: <code>19:09:59: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17738/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2052458735401439232#1:build-log.txt%3A3970)
<details>
<summary>all...</summary>

* _2026-05-07 18:48:03 &#43;0000 UTC_: <code>19:09:59: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17738/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2052458735401439232#1:build-log.txt%3A3970)

* _2026-05-07 18:43:42 &#43;0000 UTC_: <code>19:08:48: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17738/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2052458730666070016#1:build-log.txt%3A4079)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (4x / 80.00%)

<details>
<summary> download failure in context (4x / 80.00%) </summary>

<hr/>

**4x**: _2026-05-07 18:48:03 &#43;0000 UTC_: <code>19:09:59: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17738/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2052458735401439232#1:build-log.txt%3A3970)
<details>
<summary>all...</summary>

* _2026-05-10 14:21:04 &#43;0000 UTC_: <code>14:33:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17611/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2053480299794272256#1:build-log.txt%3A4762)
<details><summary>context</summary>
<pre>14:33:54: ERROR: Analysis of target &#39;//:buildifier&#39; failed; build aborted:
14:33:54: INFO: Elapsed time: 0.459s
14:33:54: INFO: 0 processes.
14:33:54: ERROR: Build failed. Not running target
make: *** [Makefile:28: bazel-build-functests] Error 1
&#43; ret=2
&#43; check_for_panics</pre>
</details>


* _2026-05-10 07:10:19 &#43;0000 UTC_: <code>07:11:16: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17746/pull-kubevirt-e2e-k8s-1.35-sig-compute/2053371905993347072#1:build-log.txt%3A324)
<details><summary>context</summary>
<pre>07:11:16: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted:
07:11:16: INFO: Elapsed time: 0.354s
07:11:16: INFO: 0 processes.
07:11:16: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>


* _2026-05-07 18:48:03 &#43;0000 UTC_: <code>19:09:59: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17738/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2052458735401439232#1:build-log.txt%3A3970)
<details><summary>context</summary>
<pre>19:09:59: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
19:09:59: INFO: Elapsed time: 1.710s
19:09:59: INFO: 0 processes.
19:09:59: ERROR: Build failed. Not running target
19:10:00: &#43; rm -f /tmp/kubevirt.deploy.WFEC
make: *** [Makefile:188: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


* _2026-05-07 18:43:42 &#43;0000 UTC_: <code>19:08:48: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17738/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2052458730666070016#1:build-log.txt%3A4079)
<details><summary>context</summary>
<pre>19:08:48: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
19:08:48: INFO: Elapsed time: 1.419s
19:08:48: INFO: 0 processes.
19:08:48: ERROR: Build failed. Not running target
19:08:50: &#43; rm -f /tmp/kubevirt.deploy.fD5m
make: *** [Makefile:188: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


</details>

<hr/>
</details>

### needs-investigation (1x / 20.00%)

<details>
<summary> no error snippets (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-05-13 12:02:10 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17794/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054532503019982848)
<details>
<summary>all...</summary>

* _2026-05-13 12:02:10 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17794/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054532503019982848)

</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (3x / 60.00%)


#### external (2x / 66.67%)

<details>
<summary> download failure in context (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-05-10 14:21:04 &#43;0000 UTC_: <code>14:33:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17611/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2053480299794272256#1:build-log.txt%3A4762)
<details>
<summary>all...</summary>

* _2026-05-10 14:21:04 &#43;0000 UTC_: <code>14:33:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17611/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2053480299794272256#1:build-log.txt%3A4762)

* _2026-05-10 07:10:19 &#43;0000 UTC_: <code>07:11:16: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17746/pull-kubevirt-e2e-k8s-1.35-sig-compute/2053371905993347072#1:build-log.txt%3A324)

</details>

<hr/>
</details>

#### needs-investigation (1x / 33.33%)

<details>
<summary> no error snippets (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-05-13 12:02:10 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17794/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054532503019982848)
<details>
<summary>all...</summary>

* _2026-05-13 12:02:10 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17794/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054532503019982848)

</details>

<hr/>
</details>

### release-1.8 (2x / 40.00%)


#### external (2x / 100.00%)

<details>
<summary> download failure in context (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-05-07 18:48:03 &#43;0000 UTC_: <code>19:09:59: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17738/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2052458735401439232#1:build-log.txt%3A3970)
<details>
<summary>all...</summary>

* _2026-05-07 18:48:03 &#43;0000 UTC_: <code>19:09:59: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17738/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2052458735401439232#1:build-log.txt%3A3970)

* _2026-05-07 18:43:42 &#43;0000 UTC_: <code>19:08:48: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17738/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2052458730666070016#1:build-log.txt%3A4079)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (4x / 80.00%)


#### external (3x / 75.00%)

<details>
<summary> download failure in context (3x / 75.00%) </summary>

<hr/>

**3x**: _2026-05-07 18:48:03 &#43;0000 UTC_: <code>19:09:59: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17738/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2052458735401439232#1:build-log.txt%3A3970)
<details>
<summary>all...</summary>

* _2026-05-10 14:21:04 &#43;0000 UTC_: <code>14:33:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17611/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2053480299794272256#1:build-log.txt%3A4762)

* _2026-05-10 07:10:19 &#43;0000 UTC_: <code>07:11:16: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17746/pull-kubevirt-e2e-k8s-1.35-sig-compute/2053371905993347072#1:build-log.txt%3A324)

* _2026-05-07 18:48:03 &#43;0000 UTC_: <code>19:09:59: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17738/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2052458735401439232#1:build-log.txt%3A3970)

</details>

<hr/>
</details>

#### needs-investigation (1x / 25.00%)

<details>
<summary> no error snippets (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-05-13 12:02:10 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17794/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054532503019982848)
<details>
<summary>all...</summary>

* _2026-05-13 12:02:10 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17794/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054532503019982848)

</details>

<hr/>
</details>

### sig-storage (1x / 20.00%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-05-07 18:43:42 &#43;0000 UTC_: <code>19:08:48: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17738/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2052458730666070016#1:build-log.txt%3A4079)
<details>
<summary>all...</summary>

* _2026-05-07 18:43:42 &#43;0000 UTC_: <code>19:08:48: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17738/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2052458730666070016#1:build-log.txt%3A4079)

</details>

<hr/>
</details>

Last updated: 2026-05-14 18:39:37
