



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-10-05 (1x / 33.33%)


#### external (1x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-10-05 06:50:11 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17974/pull-kubevirt-e2e-kind-1.37-vgpu/2107000242497916928#1:build-log.txt%3A467)
<details>
<summary>all...</summary>

* _2026-10-05 06:50:11 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17974/pull-kubevirt-e2e-kind-1.37-vgpu/2107000242497916928#1:build-log.txt%3A467)

</details>

<hr/>
</details>

### 2026-09-30 (2x / 66.67%)


#### external (1x / 50.00%)

<details>
<summary> container image pull failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-09-30 06:40:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19205/pull-kubevirt-e2e-kind-1.37-vgpu/2105185961917812736#1:build-log.txt%3A447)
<details>
<summary>all...</summary>

* _2026-09-30 06:40:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19205/pull-kubevirt-e2e-kind-1.37-vgpu/2105185961917812736#1:build-log.txt%3A447)

</details>

<hr/>
</details>

#### internal (1x / 50.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-09-30 18:11:18 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19276/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2105359692388634624#1:build-log.txt%3A4019)
<details>
<summary>all...</summary>

* _2026-09-30 18:11:18 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19276/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2105359692388634624#1:build-log.txt%3A4019)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (2x / 66.67%)

<details>
<summary> container image pull failure in context (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-10-05 06:50:11 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17974/pull-kubevirt-e2e-kind-1.37-vgpu/2107000242497916928#1:build-log.txt%3A467)
<details>
<summary>all...</summary>

* _2026-10-05 06:50:11 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17974/pull-kubevirt-e2e-kind-1.37-vgpu/2107000242497916928#1:build-log.txt%3A467)
<details><summary>context</summary>
<pre>06:56:50: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
06:56:50: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
06:56:54: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage3405997224/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261005%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261005T065653Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=de4de6f00ab9777bccfd1f6806eedc082d1711d221ec401decadde871c2de012&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1791184313~hmac=a2f9251824ce47e686de7b529ae2dd4ab2c1ebee839039e922e9f0e57b6d2b6b&#34;: EOF)
make: *** [Makefile:177: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-09-30 06:40:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19205/pull-kubevirt-e2e-kind-1.37-vgpu/2105185961917812736#1:build-log.txt%3A447)
<details><summary>context</summary>
<pre>06:47:27: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
06:47:27: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
06:47:33: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1374624727/2&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20260930%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20260930T064729Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=bece4d68dfe552cd5c1522b88171d63dfede14193b999ace37b04b86befc4490&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1790751749~hmac=c79e58b95b20823cf63d13f7ca0f01b31728757402ba1ad25e2412552646d16d&#34;: EOF)
make: *** [Makefile:177: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>

<hr/>
</details>

### internal (1x / 33.33%)

<details>
<summary> make cluster lifecycle target failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-09-30 18:11:18 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19276/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2105359692388634624#1:build-log.txt%3A4019)
<details>
<summary>all...</summary>

* _2026-09-30 18:11:18 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19276/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2105359692388634624#1:build-log.txt%3A4019)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
18:33:00: selecting podman as container runtime
18:33:39: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.32-sig-network-1.7 is being used by the following container(s): 209344e98ddcb5c9d39e17488818e81fa787e498e5920efb9e3d4d52f7e893ab, 721330580dfb78f833ab17d7f2571a4ec39df16de9fd483bfa68d88daf2116d2: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (2x / 66.67%)


#### external (2x / 100.00%)

<details>
<summary> container image pull failure in context (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-10-05 06:50:11 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17974/pull-kubevirt-e2e-kind-1.37-vgpu/2107000242497916928#1:build-log.txt%3A467)
<details>
<summary>all...</summary>

* _2026-10-05 06:50:11 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17974/pull-kubevirt-e2e-kind-1.37-vgpu/2107000242497916928#1:build-log.txt%3A467)

* _2026-09-30 06:40:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19205/pull-kubevirt-e2e-kind-1.37-vgpu/2105185961917812736#1:build-log.txt%3A447)

</details>

<hr/>
</details>

### release-1.7 (1x / 33.33%)


#### internal (1x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-30 18:11:18 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19276/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2105359692388634624#1:build-log.txt%3A4019)
<details>
<summary>all...</summary>

* _2026-09-30 18:11:18 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19276/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2105359692388634624#1:build-log.txt%3A4019)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (2x / 66.67%)


#### external (2x / 100.00%)

<details>
<summary> container image pull failure in context (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-10-05 06:50:11 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17974/pull-kubevirt-e2e-kind-1.37-vgpu/2107000242497916928#1:build-log.txt%3A467)
<details>
<summary>all...</summary>

* _2026-10-05 06:50:11 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17974/pull-kubevirt-e2e-kind-1.37-vgpu/2107000242497916928#1:build-log.txt%3A467)

* _2026-09-30 06:40:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19205/pull-kubevirt-e2e-kind-1.37-vgpu/2105185961917812736#1:build-log.txt%3A447)

</details>

<hr/>
</details>

### sig-network (1x / 33.33%)


#### internal (1x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-30 18:11:18 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19276/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2105359692388634624#1:build-log.txt%3A4019)
<details>
<summary>all...</summary>

* _2026-09-30 18:11:18 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19276/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2105359692388634624#1:build-log.txt%3A4019)

</details>

<hr/>
</details>

Last updated: 2026-10-07 03:15:41
