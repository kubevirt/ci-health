



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-09-14 (1x / 33.33%)


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

### 2026-09-09 (2x / 66.67%)


#### internal (1x / 50.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-09-09 15:19:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18939/pull-kubevirt-e2e-kind-1.36-sev/2097706364049559552#1:build-log.txt%3A670)
<details>
<summary>all...</summary>

* _2026-09-09 15:19:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18939/pull-kubevirt-e2e-kind-1.36-sev/2097706364049559552#1:build-log.txt%3A670)

</details>

<hr/>
</details>

#### external (1x / 50.00%)

<details>
<summary> container image pull failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-09-09 06:36:04 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19062/pull-kubevirt-e2e-kind-sriov-1.7/2097574590971645952#1:build-log.txt%3A1691)
<details>
<summary>all...</summary>

* _2026-09-09 06:36:04 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19062/pull-kubevirt-e2e-kind-sriov-1.7/2097574590971645952#1:build-log.txt%3A1691)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (2x / 66.67%)

<details>
<summary> container image pull failure in context (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-09-09 06:36:04 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19062/pull-kubevirt-e2e-kind-sriov-1.7/2097574590971645952#1:build-log.txt%3A1691)
<details>
<summary>all...</summary>

* _2026-09-09 06:36:04 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19062/pull-kubevirt-e2e-kind-sriov-1.7/2097574590971645952#1:build-log.txt%3A1691)
<details><summary>context</summary>
<pre>06:44:45: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
06:44:45: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
06:44:49: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage80778091/2&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/5c/5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20260909%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20260909T064447Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=c4ab1a2e41e0290dd0848902948ba35a4a8fb14fce22b5045da41d01640ebb9f&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1788937187~hmac=c5508d8ab454aed739605fdabae46dfd904e2d2b79e4208292d12933782d6b55&#34;: EOF)
make: *** [Makefile:159: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>

<hr/>
</details>
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


### main (2x / 66.67%)


#### internal (1x / 50.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-09-09 15:19:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18939/pull-kubevirt-e2e-kind-1.36-sev/2097706364049559552#1:build-log.txt%3A670)
<details>
<summary>all...</summary>

* _2026-09-09 15:19:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18939/pull-kubevirt-e2e-kind-1.36-sev/2097706364049559552#1:build-log.txt%3A670)

</details>

<hr/>
</details>

#### external (1x / 50.00%)

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

### release-1.7 (1x / 33.33%)


#### external (1x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-09 06:36:04 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19062/pull-kubevirt-e2e-kind-sriov-1.7/2097574590971645952#1:build-log.txt%3A1691)
<details>
<summary>all...</summary>

* _2026-09-09 06:36:04 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19062/pull-kubevirt-e2e-kind-sriov-1.7/2097574590971645952#1:build-log.txt%3A1691)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


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

### sig-network (1x / 33.33%)


#### external (1x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-09 06:36:04 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19062/pull-kubevirt-e2e-kind-sriov-1.7/2097574590971645952#1:build-log.txt%3A1691)
<details>
<summary>all...</summary>

* _2026-09-09 06:36:04 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19062/pull-kubevirt-e2e-kind-sriov-1.7/2097574590971645952#1:build-log.txt%3A1691)

</details>

<hr/>
</details>

Last updated: 2026-09-15 03:13:02
