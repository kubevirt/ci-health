



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-11-02 (2x / 40.00%)


#### external (2x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-11-02 12:10:13 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A1184)
<details>
<summary>all...</summary>

* _2026-11-02 12:10:13 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A1184)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-11-02 09:09:03 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A434)
<details>
<summary>all...</summary>

* _2026-11-02 09:09:03 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A434)

</details>

<hr/>
</details>

### 2026-10-28 (3x / 60.00%)


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

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (5x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (3x / 60.00%) </summary>

<hr/>

**3x**: _2026-10-28 16:02:01 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1185)
<details>
<summary>all...</summary>

* _2026-11-02 12:10:13 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A1184)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:20:53: selecting podman as container runtime
12:20:54: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-windows2016 is being used by the following container(s): 6a6a2973d141a1448dc31509ecfd824ede38d3435e163c24d8b9cba7915f160b: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


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
<summary> podman container removal timeout (1x / 20.00%) </summary>

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
<details>
<summary> container image pull failure in context (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-11-02 09:09:03 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A434)
<details>
<summary>all...</summary>

* _2026-11-02 09:09:03 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A434)
<details><summary>context</summary>
<pre>09:16:19: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
09:16:19: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)
make: *** [Makefile:179: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (5x / 100.00%)


#### external (5x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (3x / 60.00%) </summary>

<hr/>

**3x**: _2026-10-28 16:02:01 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1185)
<details>
<summary>all...</summary>

* _2026-11-02 12:10:13 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A1184)

* _2026-10-28 16:02:01 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1185)

* _2026-10-28 15:34:15 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A1350)

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
<details>
<summary> container image pull failure in context (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-11-02 09:09:03 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A434)
<details>
<summary>all...</summary>

* _2026-11-02 09:09:03 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A434)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-storage (2x / 40.00%)


#### external (2x / 100.00%)

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

### sig-compute (3x / 60.00%)


#### external (3x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-10-28 15:34:15 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A1350)
<details>
<summary>all...</summary>

* _2026-11-02 12:10:13 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A1184)

* _2026-10-28 15:34:15 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A1350)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-11-02 09:09:03 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A434)
<details>
<summary>all...</summary>

* _2026-11-02 09:09:03 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A434)

</details>

<hr/>
</details>

Last updated: 2026-11-04 09:10:32
