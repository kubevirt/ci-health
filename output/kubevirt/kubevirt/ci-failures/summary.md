



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-08-10 (1x / 50.00%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-08-10 13:52:33 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18737/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2086812749257510912#1:build-log.txt%3A1835)
<details>
<summary>all...</summary>

* _2026-08-10 13:52:33 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18737/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2086812749257510912#1:build-log.txt%3A1835)

</details>

<hr/>
</details>

### 2026-08-09 (1x / 50.00%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-08-09 12:42:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18752/pull-kubevirt-e2e-k8s-1.36-sig-network/2086432842375499776#1:build-log.txt%3A1841)
<details>
<summary>all...</summary>

* _2026-08-09 12:42:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18752/pull-kubevirt-e2e-k8s-1.36-sig-network/2086432842375499776#1:build-log.txt%3A1841)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (2x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-08-10 13:52:33 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18737/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2086812749257510912#1:build-log.txt%3A1835)
<details>
<summary>all...</summary>

* _2026-08-10 13:52:33 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18737/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2086812749257510912#1:build-log.txt%3A1835)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
13:57:31: selecting podman as container runtime
13:58:10: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network is being used by the following container(s): 351c3978b3a9cd92d61296e59786505a976965c14ff8c849f2283eff04823cd8: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-08-09 12:42:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18752/pull-kubevirt-e2e-k8s-1.36-sig-network/2086432842375499776#1:build-log.txt%3A1841)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:58:10: selecting podman as container runtime
12:58:54: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network is being used by the following container(s): 4c38961c723bc3d7d0f3c8598644852fcd67bb809a9ed9e7cf2b5ebb6c2607b8: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
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

**2x**: _2026-08-10 13:52:33 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18737/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2086812749257510912#1:build-log.txt%3A1835)
<details>
<summary>all...</summary>

* _2026-08-10 13:52:33 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18737/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2086812749257510912#1:build-log.txt%3A1835)

* _2026-08-09 12:42:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18752/pull-kubevirt-e2e-k8s-1.36-sig-network/2086432842375499776#1:build-log.txt%3A1841)

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

**2x**: _2026-08-10 13:52:33 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18737/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2086812749257510912#1:build-log.txt%3A1835)
<details>
<summary>all...</summary>

* _2026-08-10 13:52:33 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18737/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2086812749257510912#1:build-log.txt%3A1835)

* _2026-08-09 12:42:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18752/pull-kubevirt-e2e-k8s-1.36-sig-network/2086432842375499776#1:build-log.txt%3A1841)

</details>

<hr/>
</details>

Last updated: 2026-08-11 06:24:57
