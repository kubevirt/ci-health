



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-09-18 (1x / 50.00%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-18 07:56:21 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19175/pull-kubevirt-e2e-k8s-1.36-sig-network/2100856101317644288#1:build-log.txt%3A1879)
<details>
<summary>all...</summary>

* _2026-09-18 07:56:21 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19175/pull-kubevirt-e2e-k8s-1.36-sig-network/2100856101317644288#1:build-log.txt%3A1879)

</details>

<hr/>
</details>

### 2026-09-17 (1x / 50.00%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-17 13:43:51 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19164/pull-kubevirt-e2e-k8s-1.36-sig-network/2100562626076479488#1:build-log.txt%3A1882)
<details>
<summary>all...</summary>

* _2026-09-17 13:43:51 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19164/pull-kubevirt-e2e-k8s-1.36-sig-network/2100562626076479488#1:build-log.txt%3A1882)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (2x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-09-18 07:56:21 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19175/pull-kubevirt-e2e-k8s-1.36-sig-network/2100856101317644288#1:build-log.txt%3A1879)
<details>
<summary>all...</summary>

* _2026-09-18 07:56:21 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19175/pull-kubevirt-e2e-k8s-1.36-sig-network/2100856101317644288#1:build-log.txt%3A1879)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
08:01:45: selecting podman as container runtime
08:02:38: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network is being used by the following container(s): 2b5cd12f54fcc250811be6b95bda70c694cd5ce6f6bc567b49fc54e6c1af615a: volume is being used
make: *** [Makefile:180: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-09-17 13:43:51 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19164/pull-kubevirt-e2e-k8s-1.36-sig-network/2100562626076479488#1:build-log.txt%3A1882)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
13:49:21: selecting podman as container runtime
13:50:01: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network is being used by the following container(s): 4d689305039cfb27be7198af2779554127f7e4db8b2454a1d0ebff95d7c37200: volume is being used
make: *** [Makefile:180: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (2x / 100.00%)


#### external (2x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-09-18 07:56:21 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19175/pull-kubevirt-e2e-k8s-1.36-sig-network/2100856101317644288#1:build-log.txt%3A1879)
<details>
<summary>all...</summary>

* _2026-09-18 07:56:21 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19175/pull-kubevirt-e2e-k8s-1.36-sig-network/2100856101317644288#1:build-log.txt%3A1879)

* _2026-09-17 13:43:51 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19164/pull-kubevirt-e2e-k8s-1.36-sig-network/2100562626076479488#1:build-log.txt%3A1882)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-network (2x / 100.00%)


#### external (2x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-09-18 07:56:21 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19175/pull-kubevirt-e2e-k8s-1.36-sig-network/2100856101317644288#1:build-log.txt%3A1879)
<details>
<summary>all...</summary>

* _2026-09-18 07:56:21 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19175/pull-kubevirt-e2e-k8s-1.36-sig-network/2100856101317644288#1:build-log.txt%3A1879)

* _2026-09-17 13:43:51 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19164/pull-kubevirt-e2e-k8s-1.36-sig-network/2100562626076479488#1:build-log.txt%3A1882)

</details>

<hr/>
</details>

Last updated: 2026-09-24 09:09:16
