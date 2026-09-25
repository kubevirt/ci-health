



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-09-24 (2x / 100.00%)


#### external (2x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-09-24 06:12:02 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17922/pull-kubevirt-e2e-kind-1.37-vgpu/2103004361624915968#1:build-log.txt%3A509)
<details>
<summary>all...</summary>

* _2026-09-24 06:12:02 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17922/pull-kubevirt-e2e-kind-1.37-vgpu/2103004361624915968#1:build-log.txt%3A509)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-09-24 08:55:19 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19214/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2103045438079766528#1:build-log.txt%3A1865)
<details>
<summary>all...</summary>

* _2026-09-24 08:55:19 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19214/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2103045438079766528#1:build-log.txt%3A1865)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (2x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-09-24 06:12:02 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17922/pull-kubevirt-e2e-kind-1.37-vgpu/2103004361624915968#1:build-log.txt%3A509)
<details>
<summary>all...</summary>

* _2026-09-24 06:12:02 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17922/pull-kubevirt-e2e-kind-1.37-vgpu/2103004361624915968#1:build-log.txt%3A509)
<details><summary>context</summary>
<pre>06:25:34: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
06:25:34: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
06:25:38: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1519539013/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/5c/5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20260924%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20260924T062537Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=2f35355f57c757563f50779ec88c0ed96d02ad9028d2be72e47e0d8aada528dc&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1790232037~hmac=df0af00d89f54e51b4a3357eb77394f301d234403168562cec1be3addd950e42&#34;: EOF)
make: *** [Makefile:177: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-09-24 08:55:19 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19214/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2103045438079766528#1:build-log.txt%3A1865)
<details>
<summary>all...</summary>

* _2026-09-24 08:55:19 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19214/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2103045438079766528#1:build-log.txt%3A1865)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
09:02:35: selecting podman as container runtime
09:03:08: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network is being used by the following container(s): 6e8772b9281a5e60a84b8770a1de0b028641b03275abaca545336f27093cc63e: volume is being used
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
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-09-24 08:55:19 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19214/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2103045438079766528#1:build-log.txt%3A1865)
<details>
<summary>all...</summary>

* _2026-09-24 08:55:19 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19214/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2103045438079766528#1:build-log.txt%3A1865)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-09-24 06:12:02 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17922/pull-kubevirt-e2e-kind-1.37-vgpu/2103004361624915968#1:build-log.txt%3A509)
<details>
<summary>all...</summary>

* _2026-09-24 06:12:02 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17922/pull-kubevirt-e2e-kind-1.37-vgpu/2103004361624915968#1:build-log.txt%3A509)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-network (1x / 50.00%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-24 08:55:19 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19214/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2103045438079766528#1:build-log.txt%3A1865)
<details>
<summary>all...</summary>

* _2026-09-24 08:55:19 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19214/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2103045438079766528#1:build-log.txt%3A1865)

</details>

<hr/>
</details>

### sig-compute (1x / 50.00%)


#### external (1x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-24 06:12:02 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17922/pull-kubevirt-e2e-kind-1.37-vgpu/2103004361624915968#1:build-log.txt%3A509)
<details>
<summary>all...</summary>

* _2026-09-24 06:12:02 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17922/pull-kubevirt-e2e-kind-1.37-vgpu/2103004361624915968#1:build-log.txt%3A509)

</details>

<hr/>
</details>

Last updated: 2026-09-25 09:11:21
