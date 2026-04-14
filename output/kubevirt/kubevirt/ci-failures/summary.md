



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-04-09 (3x / 100.00%)


#### internal (2x / 66.67%)

<details>
<summary> make cluster lifecycle target failure (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-04-09 01:56:53 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2042058998902951936#1:build-log.txt%3A1040)
<details>
<summary>all...</summary>

* _2026-04-09 01:56:54 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2042059010437287936#1:build-log.txt%3A1142)

* _2026-04-09 01:56:53 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2042058998902951936#1:build-log.txt%3A1040)

</details>

<hr/>
</details>

#### external (1x / 33.33%)

<details>
<summary> download failure in context (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-04-09 14:20:27 &#43;0000 UTC_: <code>14:31:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17340/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2042246124231725056#1:build-log.txt%3A4058)
<details>
<summary>all...</summary>

* _2026-04-09 14:20:27 &#43;0000 UTC_: <code>14:31:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17340/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2042246124231725056#1:build-log.txt%3A4058)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### internal (2x / 66.67%)

<details>
<summary> make cluster lifecycle target failure (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-04-09 01:56:53 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2042058998902951936#1:build-log.txt%3A1040)
<details>
<summary>all...</summary>

* _2026-04-09 01:56:54 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2042059010437287936#1:build-log.txt%3A1142)
<details><summary>context</summary>
<pre>02:23:46: INFO: 2 processes: 1 internal, 1 processwrapper-sandbox.
02:23:46: INFO: Running command line: bazel-bin/example-guest-agent-copier /root/go/src/kubevirt.io/kubevirt/_out/cmd/example-guest-agent/example-guest-agent
02:23:58: &#43; rm -f /tmp/kubevirt.deploy.h4Rn
make: *** [Makefile:174: cluster-sync] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>


* _2026-04-09 01:56:53 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2042058998902951936#1:build-log.txt%3A1040)
<details><summary>context</summary>
<pre>02:40:33: INFO: 2 processes: 1 internal, 1 processwrapper-sandbox.
02:40:33: INFO: Running command line: bazel-bin/example-guest-agent-copier /root/go/src/kubevirt.io/kubevirt/_out/cmd/example-guest-agent/example-guest-agent
02:40:50: &#43; rm -f /tmp/kubevirt.deploy.RzZ0
make: *** [Makefile:174: cluster-sync] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>


</details>

<hr/>
</details>

### external (1x / 33.33%)

<details>
<summary> download failure in context (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-04-09 14:20:27 &#43;0000 UTC_: <code>14:31:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17340/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2042246124231725056#1:build-log.txt%3A4058)
<details>
<summary>all...</summary>

* _2026-04-09 14:20:27 &#43;0000 UTC_: <code>14:31:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17340/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2042246124231725056#1:build-log.txt%3A4058)
<details><summary>context</summary>
<pre>14:31:10: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted:
14:31:10: INFO: Elapsed time: 1.273s
14:31:10: INFO: 0 processes.
14:31:10: ERROR: Build failed. Not running target
14:31:12: &#43; rm -f /tmp/kubevirt.deploy.7s31
make: *** [Makefile:188: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (1x / 33.33%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-04-09 14:20:27 &#43;0000 UTC_: <code>14:31:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17340/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2042246124231725056#1:build-log.txt%3A4058)
<details>
<summary>all...</summary>

* _2026-04-09 14:20:27 &#43;0000 UTC_: <code>14:31:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17340/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2042246124231725056#1:build-log.txt%3A4058)

</details>

<hr/>
</details>

### release-1.7 (2x / 66.67%)


#### internal (2x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-04-09 01:56:53 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2042058998902951936#1:build-log.txt%3A1040)
<details>
<summary>all...</summary>

* _2026-04-09 01:56:54 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2042059010437287936#1:build-log.txt%3A1142)

* _2026-04-09 01:56:53 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2042058998902951936#1:build-log.txt%3A1040)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (2x / 66.67%)


#### internal (1x / 50.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-04-09 01:56:53 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2042058998902951936#1:build-log.txt%3A1040)
<details>
<summary>all...</summary>

* _2026-04-09 01:56:53 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2042058998902951936#1:build-log.txt%3A1040)

</details>

<hr/>
</details>

#### external (1x / 50.00%)

<details>
<summary> download failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-04-09 14:20:27 &#43;0000 UTC_: <code>14:31:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17340/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2042246124231725056#1:build-log.txt%3A4058)
<details>
<summary>all...</summary>

* _2026-04-09 14:20:27 &#43;0000 UTC_: <code>14:31:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17340/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2042246124231725056#1:build-log.txt%3A4058)

</details>

<hr/>
</details>

### sig-storage (1x / 33.33%)


#### internal (1x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-04-09 01:56:54 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2042059010437287936#1:build-log.txt%3A1142)
<details>
<summary>all...</summary>

* _2026-04-09 01:56:54 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2042059010437287936#1:build-log.txt%3A1142)

</details>

<hr/>
</details>

Last updated: 2026-04-14 09:43:16
