



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-09-14 (2x / 66.67%)


#### external (2x / 100.00%)

<details>
<summary> download failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-09-14 14:19:35 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19052/pull-kubevirt-e2e-k8s-1.37-sig-network/2099500599073574912#1:build-log.txt%3A4018)
<details>
<summary>all...</summary>

* _2026-09-14 14:19:35 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19052/pull-kubevirt-e2e-k8s-1.37-sig-network/2099500599073574912#1:build-log.txt%3A4018)

</details>

<hr/>
</details>
<details>
<summary> podman container removal timeout (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-09-14 11:12:09 &#43;0000 UTC_: <code>Error: cannot remove container 8dda5b7c95ec4de1ce22ee25296269c9781d5bd8e5f18b79922c436d148852b6 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19041/pull-kubevirt-e2e-k8s-1.35-sig-storage/2099456004369420288#1:build-log.txt%3A3068)
<details>
<summary>all...</summary>

* _2026-09-14 11:12:09 &#43;0000 UTC_: <code>Error: cannot remove container 8dda5b7c95ec4de1ce22ee25296269c9781d5bd8e5f18b79922c436d148852b6 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19041/pull-kubevirt-e2e-k8s-1.35-sig-storage/2099456004369420288#1:build-log.txt%3A3068)

</details>

<hr/>
</details>

### 2026-09-09 (1x / 33.33%)


#### internal (1x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-09 15:19:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18939/pull-kubevirt-e2e-kind-1.36-sev/2097706364049559552#1:build-log.txt%3A670)
<details>
<summary>all...</summary>

* _2026-09-09 15:19:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18939/pull-kubevirt-e2e-kind-1.36-sev/2097706364049559552#1:build-log.txt%3A670)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (2x / 66.67%)

<details>
<summary> podman container removal timeout (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-09-14 11:12:09 &#43;0000 UTC_: <code>Error: cannot remove container 8dda5b7c95ec4de1ce22ee25296269c9781d5bd8e5f18b79922c436d148852b6 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19041/pull-kubevirt-e2e-k8s-1.35-sig-storage/2099456004369420288#1:build-log.txt%3A3068)
<details>
<summary>all...</summary>

* _2026-09-14 11:12:09 &#43;0000 UTC_: <code>Error: cannot remove container 8dda5b7c95ec4de1ce22ee25296269c9781d5bd8e5f18b79922c436d148852b6 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19041/pull-kubevirt-e2e-k8s-1.35-sig-storage/2099456004369420288#1:build-log.txt%3A3068)
<details><summary>context</summary>
<pre>ab442c0f419f7b9eed6a132f40f3c7ae8ffdfe8b032f096352a14062b314af70
8231b50111059f2c9588ee42058ad906a3528dbe61f80a3d69ea7c1720508460
Error: cannot remove container 8dda5b7c95ec4de1ce22ee25296269c9781d5bd8e5f18b79922c436d148852b6 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 8dda5b7c95ec4de1ce22ee25296269c9781d5bd8e5f18b79922c436d148852b6 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1020 is not a child of this shell
================================================================================
Done cleaning up after podman in container.</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-09-14 14:19:35 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19052/pull-kubevirt-e2e-k8s-1.37-sig-network/2099500599073574912#1:build-log.txt%3A4018)
<details>
<summary>all...</summary>

* _2026-09-14 14:19:35 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19052/pull-kubevirt-e2e-k8s-1.37-sig-network/2099500599073574912#1:build-log.txt%3A4018)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
14:36:52: selecting podman as container runtime
14:37:31: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.37-sig-network is being used by the following container(s): 73078e928d0f0cb69b0e44b1cec7cf0f69114147ddfc67bfeff47f47ba3b3268: volume is being used
make: *** [Makefile:180: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>

### internal (1x / 33.33%)

<details>
<summary> make cluster lifecycle target failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-09-09 15:19:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18939/pull-kubevirt-e2e-kind-1.36-sev/2097706364049559552#1:build-log.txt%3A670)
<details>
<summary>all...</summary>

* _2026-09-09 15:19:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18939/pull-kubevirt-e2e-kind-1.36-sev/2097706364049559552#1:build-log.txt%3A670)
<details><summary>context</summary>
<pre>15:25:09: 	sigs.k8s.io/kind/pkg/errors/concurrent.go:30
15:25:09: runtime.goexit
15:25:09: 	runtime/asm_amd64.s:1771
make: *** [Makefile:177: cluster-up] Error 1
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


#### external (2x / 66.67%)

<details>
<summary> podman container removal timeout (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-09-14 11:12:09 &#43;0000 UTC_: <code>Error: cannot remove container 8dda5b7c95ec4de1ce22ee25296269c9781d5bd8e5f18b79922c436d148852b6 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19041/pull-kubevirt-e2e-k8s-1.35-sig-storage/2099456004369420288#1:build-log.txt%3A3068)
<details>
<summary>all...</summary>

* _2026-09-14 11:12:09 &#43;0000 UTC_: <code>Error: cannot remove container 8dda5b7c95ec4de1ce22ee25296269c9781d5bd8e5f18b79922c436d148852b6 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19041/pull-kubevirt-e2e-k8s-1.35-sig-storage/2099456004369420288#1:build-log.txt%3A3068)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-09-14 14:19:35 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19052/pull-kubevirt-e2e-k8s-1.37-sig-network/2099500599073574912#1:build-log.txt%3A4018)
<details>
<summary>all...</summary>

* _2026-09-14 14:19:35 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19052/pull-kubevirt-e2e-k8s-1.37-sig-network/2099500599073574912#1:build-log.txt%3A4018)

</details>

<hr/>
</details>

#### internal (1x / 33.33%)

<details>
<summary> make cluster lifecycle target failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-09-09 15:19:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18939/pull-kubevirt-e2e-kind-1.36-sev/2097706364049559552#1:build-log.txt%3A670)
<details>
<summary>all...</summary>

* _2026-09-09 15:19:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18939/pull-kubevirt-e2e-kind-1.36-sev/2097706364049559552#1:build-log.txt%3A670)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-network (1x / 33.33%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-14 14:19:35 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19052/pull-kubevirt-e2e-k8s-1.37-sig-network/2099500599073574912#1:build-log.txt%3A4018)
<details>
<summary>all...</summary>

* _2026-09-14 14:19:35 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19052/pull-kubevirt-e2e-k8s-1.37-sig-network/2099500599073574912#1:build-log.txt%3A4018)

</details>

<hr/>
</details>

### sig-storage (1x / 33.33%)


#### external (1x / 100.00%)

<details>
<summary> podman container removal timeout (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-14 11:12:09 &#43;0000 UTC_: <code>Error: cannot remove container 8dda5b7c95ec4de1ce22ee25296269c9781d5bd8e5f18b79922c436d148852b6 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19041/pull-kubevirt-e2e-k8s-1.35-sig-storage/2099456004369420288#1:build-log.txt%3A3068)
<details>
<summary>all...</summary>

* _2026-09-14 11:12:09 &#43;0000 UTC_: <code>Error: cannot remove container 8dda5b7c95ec4de1ce22ee25296269c9781d5bd8e5f18b79922c436d148852b6 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19041/pull-kubevirt-e2e-k8s-1.35-sig-storage/2099456004369420288#1:build-log.txt%3A3068)

</details>

<hr/>
</details>

### sig-compute (1x / 33.33%)


#### internal (1x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-09 15:19:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18939/pull-kubevirt-e2e-kind-1.36-sev/2097706364049559552#1:build-log.txt%3A670)
<details>
<summary>all...</summary>

* _2026-09-09 15:19:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18939/pull-kubevirt-e2e-kind-1.36-sev/2097706364049559552#1:build-log.txt%3A670)

</details>

<hr/>
</details>

Last updated: 2026-09-16 12:15:49
