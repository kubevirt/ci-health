



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-09-09 (2x / 66.67%)


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

### 2026-09-07 (1x / 33.33%)


#### internal (1x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-07 12:22:06 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18889/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2096932826673844224#1:build-log.txt%3A4015)
<details>
<summary>all...</summary>

* _2026-09-07 12:22:06 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18889/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2096932826673844224#1:build-log.txt%3A4015)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### internal (2x / 66.67%)

<details>
<summary> make cluster lifecycle target failure (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-09-09 15:19:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18939/pull-kubevirt-e2e-kind-1.36-sev/2097706364049559552#1:build-log.txt%3A670)
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


* _2026-09-07 12:22:06 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18889/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2096932826673844224#1:build-log.txt%3A4015)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:45:36: selecting podman as container runtime
12:46:14: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7 is being used by the following container(s): eaec627e2dc367068811abba11ce563e65ad0ef7104d0a553a4aeca33760b667, 970ef41c06bae9bb159f873cb853d49834ae7442b110c0d1b13e83fa02bcf3c6: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>

### external (1x / 33.33%)

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

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (1x / 33.33%)


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

### release-1.7 (2x / 66.67%)


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

#### internal (1x / 50.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-09-07 12:22:06 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18889/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2096932826673844224#1:build-log.txt%3A4015)
<details>
<summary>all...</summary>

* _2026-09-07 12:22:06 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18889/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2096932826673844224#1:build-log.txt%3A4015)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


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

### sig-storage (1x / 33.33%)


#### internal (1x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-07 12:22:06 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18889/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2096932826673844224#1:build-log.txt%3A4015)
<details>
<summary>all...</summary>

* _2026-09-07 12:22:06 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18889/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2096932826673844224#1:build-log.txt%3A4015)

</details>

<hr/>
</details>

Last updated: 2026-09-14 12:14:05
