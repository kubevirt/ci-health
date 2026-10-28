



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-10-28 (1x / 50.00%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-10-28 16:02:01 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1185)
<details>
<summary>all...</summary>

* _2026-10-28 16:02:01 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1185)

</details>

<hr/>
</details>

### 2026-10-26 (1x / 50.00%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-10-26 17:44:21 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19360/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2114742369638158336#1:build-log.txt%3A3317)
<details>
<summary>all...</summary>

* _2026-10-26 17:44:21 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19360/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2114742369638158336#1:build-log.txt%3A3317)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (2x / 100.00%)

<details>
<summary> download failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-10-26 17:44:21 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19360/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2114742369638158336#1:build-log.txt%3A3317)
<details>
<summary>all...</summary>

* _2026-10-26 17:44:21 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19360/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2114742369638158336#1:build-log.txt%3A3317)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
17:59:56: selecting podman as container runtime
18:00:28: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network is being used by the following container(s): 0e91ad027402179897c4b6084ba48d5bd5348b003abd7919b48a44eb060f9ac1: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-10-28 16:02:01 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1185)
<details>
<summary>all...</summary>

* _2026-10-28 16:02:01 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1185)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
16:14:40: selecting podman as container runtime
16:14:40: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-storage is being used by the following container(s): ab5ac3b271a5edd5373cc1acd2a27a8884d4400a81f5eed93a634e9d5f377457: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
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
<summary> download failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-10-26 17:44:21 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19360/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2114742369638158336#1:build-log.txt%3A3317)
<details>
<summary>all...</summary>

* _2026-10-26 17:44:21 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19360/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2114742369638158336#1:build-log.txt%3A3317)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-10-28 16:02:01 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1185)
<details>
<summary>all...</summary>

* _2026-10-28 16:02:01 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1185)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-storage (1x / 50.00%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-10-28 16:02:01 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1185)
<details>
<summary>all...</summary>

* _2026-10-28 16:02:01 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1185)

</details>

<hr/>
</details>

### sig-network (1x / 50.00%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-10-26 17:44:21 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19360/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2114742369638158336#1:build-log.txt%3A3317)
<details>
<summary>all...</summary>

* _2026-10-26 17:44:21 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19360/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2114742369638158336#1:build-log.txt%3A3317)

</details>

<hr/>
</details>

Last updated: 2026-10-28 21:15:24
