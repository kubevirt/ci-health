



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-08-12 (1x / 25.00%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-08-12 07:12:13 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18796/pull-kubevirt-e2e-kind-1.36-sev/2087436754020732928#1:build-log.txt%3A3065)
<details>
<summary>all...</summary>

* _2026-08-12 07:12:13 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18796/pull-kubevirt-e2e-kind-1.36-sev/2087436754020732928#1:build-log.txt%3A3065)

</details>

<hr/>
</details>

### 2026-08-11 (1x / 25.00%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-08-11 09:59:06 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18696/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2087116365491081216#1:build-log.txt%3A1809)
<details>
<summary>all...</summary>

* _2026-08-11 09:59:06 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18696/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2087116365491081216#1:build-log.txt%3A1809)

</details>

<hr/>
</details>

### 2026-08-10 (1x / 25.00%)


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

### 2026-08-09 (1x / 25.00%)


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


### external (4x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (4x / 100.00%) </summary>

<hr/>

**3x**: _2026-08-11 09:59:06 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18696/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2087116365491081216#1:build-log.txt%3A1809)
<details>
<summary>all...</summary>

* _2026-08-11 09:59:06 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18696/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2087116365491081216#1:build-log.txt%3A1809)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
10:03:29: selecting podman as container runtime
10:04:06: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network is being used by the following container(s): 1b3b76e9551ad92365b6d9f13cd729ce578e50cc7f611864149ccf86f1a052ce: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


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

**1x**: _2026-08-12 07:12:13 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18796/pull-kubevirt-e2e-kind-1.36-sev/2087436754020732928#1:build-log.txt%3A3065)
<details>
<summary>all...</summary>

* _2026-08-12 07:12:13 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18796/pull-kubevirt-e2e-kind-1.36-sev/2087436754020732928#1:build-log.txt%3A3065)
<details><summary>context</summary>
<pre>&#43; set &#43;x

================================
ERROR: Found panic in test output
Files:
https://storage.googleapis.com/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18796/pull-kubevirt-e2e-kind-1.36-sev/2087436754020732928/artifacts/pods/0_kubevirt_virt-handler-bhk8x-virt-handler.log
https://storage.googleapis.com/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18796/pull-kubevirt-e2e-kind-1.36-sev/2087436754020732928/artifacts/pods/0_kubevirt_virt-handler-bhk8x-virt-handler_previous.log</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (4x / 100.00%)


#### external (4x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (4x / 100.00%) </summary>

<hr/>

**3x**: _2026-08-11 09:59:06 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18696/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2087116365491081216#1:build-log.txt%3A1809)
<details>
<summary>all...</summary>

* _2026-08-11 09:59:06 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18696/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2087116365491081216#1:build-log.txt%3A1809)

* _2026-08-10 13:52:33 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18737/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2086812749257510912#1:build-log.txt%3A1835)

* _2026-08-09 12:42:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18752/pull-kubevirt-e2e-k8s-1.36-sig-network/2086432842375499776#1:build-log.txt%3A1841)

</details>

<hr/>

**1x**: _2026-08-12 07:12:13 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18796/pull-kubevirt-e2e-kind-1.36-sev/2087436754020732928#1:build-log.txt%3A3065)
<details>
<summary>all...</summary>

* _2026-08-12 07:12:13 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18796/pull-kubevirt-e2e-kind-1.36-sev/2087436754020732928#1:build-log.txt%3A3065)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (1x / 25.00%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-08-12 07:12:13 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18796/pull-kubevirt-e2e-kind-1.36-sev/2087436754020732928#1:build-log.txt%3A3065)
<details>
<summary>all...</summary>

* _2026-08-12 07:12:13 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18796/pull-kubevirt-e2e-kind-1.36-sev/2087436754020732928#1:build-log.txt%3A3065)

</details>

<hr/>
</details>

### sig-network (3x / 75.00%)


#### external (3x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (3x / 100.00%) </summary>

<hr/>

**3x**: _2026-08-11 09:59:06 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18696/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2087116365491081216#1:build-log.txt%3A1809)
<details>
<summary>all...</summary>

* _2026-08-11 09:59:06 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18696/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2087116365491081216#1:build-log.txt%3A1809)

* _2026-08-10 13:52:33 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18737/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2086812749257510912#1:build-log.txt%3A1835)

* _2026-08-09 12:42:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18752/pull-kubevirt-e2e-k8s-1.36-sig-network/2086432842375499776#1:build-log.txt%3A1841)

</details>

<hr/>
</details>

Last updated: 2026-08-13 15:48:19
