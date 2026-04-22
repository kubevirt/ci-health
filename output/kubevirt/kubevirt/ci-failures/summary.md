



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-04-17 (2x / 66.67%)


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

### 2026-04-16 (1x / 33.33%)


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


### external (3x / 100.00%)

<details>
<summary> download failure in context (2x / 66.67%) </summary>

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
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 33.33%) </summary>

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


### main (3x / 100.00%)


#### external (3x / 100.00%)

<details>
<summary> download failure in context (2x / 66.67%) </summary>

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
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 33.33%) </summary>

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


### sig-network (1x / 33.33%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-04-17 13:33:58 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-kind-sriov/2045131370006581248#1:build-log.txt%3A1394)
<details>
<summary>all...</summary>

* _2026-04-17 13:33:58 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-kind-sriov/2045131370006581248#1:build-log.txt%3A1394)

</details>

<hr/>
</details>

### sig-compute (2x / 66.67%)


#### external (2x / 100.00%)

<details>
<summary> download failure in context (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-04-17 13:25:22 &#43;0000 UTC_: <code>13:37:52: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-k8s-1.34-windows2016/2045131369809448960#1:build-log.txt%3A4706)
<details>
<summary>all...</summary>

* _2026-04-17 13:25:22 &#43;0000 UTC_: <code>13:37:52: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-k8s-1.34-windows2016/2045131369809448960#1:build-log.txt%3A4706)

* _2026-04-16 20:05:44 &#43;0000 UTC_: <code>20:18:08: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17437/pull-kubevirt-e2e-k8s-1.35-sig-compute/2044869730921091072#1:build-log.txt%3A4004)

</details>

<hr/>
</details>

Last updated: 2026-04-22 15:39:44
