



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-10-28 (3x / 50.00%)


#### external (3x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-10-28 16:02:01 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1185)
<details>
<summary>all...</summary>

* _2026-10-28 16:02:01 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1185)

* _2026-10-28 15:34:15 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A1350)

</details>

<hr/>
</details>
<details>
<summary> podman container removal timeout (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-10-28 15:32:58 &#43;0000 UTC_: <code>Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2237)
<details>
<summary>all...</summary>

* _2026-10-28 15:32:58 &#43;0000 UTC_: <code>Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2237)

</details>

<hr/>
</details>

### 2026-10-27 (1x / 16.67%)


#### external (1x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-10-27 06:08:03 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19639/pull-kubevirt-e2e-kind-1.35-vgpu-1.9/2114962157354356736#1:build-log.txt%3A531)
<details>
<summary>all...</summary>

* _2026-10-27 06:08:03 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19639/pull-kubevirt-e2e-kind-1.35-vgpu-1.9/2114962157354356736#1:build-log.txt%3A531)

</details>

<hr/>
</details>

### 2026-10-26 (2x / 33.33%)


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
<summary> container image pull failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-10-26 07:00:55 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19624/pull-kubevirt-e2e-kind-1.37-vgpu/2114613076098551808#1:build-log.txt%3A439)
<details>
<summary>all...</summary>

* _2026-10-26 07:00:55 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19624/pull-kubevirt-e2e-kind-1.37-vgpu/2114613076098551808#1:build-log.txt%3A439)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (6x / 100.00%)

<details>
<summary> container image pull failure in context (2x / 33.33%) </summary>

<hr/>

**2x**: _2026-10-26 07:00:55 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19624/pull-kubevirt-e2e-kind-1.37-vgpu/2114613076098551808#1:build-log.txt%3A439)
<details>
<summary>all...</summary>

* _2026-10-27 06:08:03 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19639/pull-kubevirt-e2e-kind-1.35-vgpu-1.9/2114962157354356736#1:build-log.txt%3A531)
<details><summary>context</summary>
<pre>06:17:58: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
06:17:58: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
06:18:04: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage3981209837/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/5c/5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261027%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261027T061800Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=438c01702c70cb3131eca04505c8f91aadf038c2107670a3889b13bbd9613d78&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793082780~hmac=cb4aaea5b71afadd00e517f2d40de6f7d1474fe4f0f21342be930e05d43a3fba&#34;: EOF)
make: *** [Makefile:174: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-10-26 07:00:55 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19624/pull-kubevirt-e2e-kind-1.37-vgpu/2114613076098551808#1:build-log.txt%3A439)
<details><summary>context</summary>
<pre>07:07:14: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
07:07:14: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
07:07:16: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage2871619320/2&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/5c/5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261026%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261026T070716Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=b31d3d7e3c6bd5b4b899d81bb5c88ac8d43a95198c028530946491f75e924f77&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1792999336~hmac=92b3edd9f6efd0d46292fabef37128e78b880ee7f865a24f5390df899399911e&#34;: EOF)
make: *** [Makefile:179: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 33.33%) </summary>

<hr/>

**2x**: _2026-10-28 16:02:01 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1185)
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


* _2026-10-28 15:34:15 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A1350)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
15:40:38: selecting podman as container runtime
15:41:00: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-operator is being used by the following container(s): 072ddc794764862cead4ae6b6c5479ed96468412c055fb6419931fe946ae6cc7: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 16.67%) </summary>

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
<summary> podman container removal timeout (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-10-28 15:32:58 &#43;0000 UTC_: <code>Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2237)
<details>
<summary>all...</summary>

* _2026-10-28 15:32:58 &#43;0000 UTC_: <code>Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2237)
<details><summary>context</summary>
<pre>24814c8573d55105ee1958eebc65d6bbadec60390bf3034e207bf414b25d324a
065a95016f2c9afd305c46ed8fdf873e16712de179534c5dccf2dfb28fd45dd4
Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1020 is not a child of this shell
================================================================================
Done cleaning up after podman in container.</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (5x / 83.33%)


#### external (5x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 40.00%) </summary>

<hr/>

**2x**: _2026-10-28 16:02:01 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1185)
<details>
<summary>all...</summary>

* _2026-10-28 16:02:01 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1185)

* _2026-10-28 15:34:15 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A1350)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-10-26 07:00:55 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19624/pull-kubevirt-e2e-kind-1.37-vgpu/2114613076098551808#1:build-log.txt%3A439)
<details>
<summary>all...</summary>

* _2026-10-26 07:00:55 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19624/pull-kubevirt-e2e-kind-1.37-vgpu/2114613076098551808#1:build-log.txt%3A439)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-10-26 17:44:21 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19360/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2114742369638158336#1:build-log.txt%3A3317)
<details>
<summary>all...</summary>

* _2026-10-26 17:44:21 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19360/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2114742369638158336#1:build-log.txt%3A3317)

</details>

<hr/>
</details>
<details>
<summary> podman container removal timeout (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-10-28 15:32:58 &#43;0000 UTC_: <code>Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2237)
<details>
<summary>all...</summary>

* _2026-10-28 15:32:58 &#43;0000 UTC_: <code>Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2237)

</details>

<hr/>
</details>

### release-1.9 (1x / 16.67%)


#### external (1x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-10-27 06:08:03 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19639/pull-kubevirt-e2e-kind-1.35-vgpu-1.9/2114962157354356736#1:build-log.txt%3A531)
<details>
<summary>all...</summary>

* _2026-10-27 06:08:03 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19639/pull-kubevirt-e2e-kind-1.35-vgpu-1.9/2114962157354356736#1:build-log.txt%3A531)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-storage (2x / 33.33%)


#### external (2x / 100.00%)

<details>
<summary> podman container removal timeout (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-10-28 15:32:58 &#43;0000 UTC_: <code>Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2237)
<details>
<summary>all...</summary>

* _2026-10-28 15:32:58 &#43;0000 UTC_: <code>Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2237)

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

### sig-network (1x / 16.67%)


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

### sig-compute (3x / 50.00%)


#### external (3x / 100.00%)

<details>
<summary> container image pull failure in context (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-10-26 07:00:55 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19624/pull-kubevirt-e2e-kind-1.37-vgpu/2114613076098551808#1:build-log.txt%3A439)
<details>
<summary>all...</summary>

* _2026-10-27 06:08:03 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19639/pull-kubevirt-e2e-kind-1.35-vgpu-1.9/2114962157354356736#1:build-log.txt%3A531)

* _2026-10-26 07:00:55 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19624/pull-kubevirt-e2e-kind-1.37-vgpu/2114613076098551808#1:build-log.txt%3A439)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-10-28 15:34:15 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A1350)
<details>
<summary>all...</summary>

* _2026-10-28 15:34:15 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A1350)

</details>

<hr/>
</details>

Last updated: 2026-10-31 12:10:28
