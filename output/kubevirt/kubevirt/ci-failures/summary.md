



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-08-25 (1x / 25.00%)


#### external (1x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-08-25 07:47:09 &#43;0000 UTC_: <code>Error: cleaning up container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888: unmounting container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 storage: cleaning up container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 storage: unmounting container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 root filesystem: deleting layer &#34;90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9 /var/lib/containers/storage/overlay/tempdirs/temp-dir-2941215159/1-90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18875/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2092154790824906752#1:build-log.txt%3A589)
<details>
<summary>all...</summary>

* _2026-08-25 07:47:09 &#43;0000 UTC_: <code>Error: cleaning up container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888: unmounting container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 storage: cleaning up container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 storage: unmounting container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 root filesystem: deleting layer &#34;90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9 /var/lib/containers/storage/overlay/tempdirs/temp-dir-2941215159/1-90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18875/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2092154790824906752#1:build-log.txt%3A589)

</details>

<hr/>
</details>

### 2026-08-24 (1x / 25.00%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-08-24 19:05:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18867/pull-kubevirt-e2e-k8s-1.36-sig-network/2091964976599142400#1:build-log.txt%3A1875)
<details>
<summary>all...</summary>

* _2026-08-24 19:05:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18867/pull-kubevirt-e2e-k8s-1.36-sig-network/2091964976599142400#1:build-log.txt%3A1875)

</details>

<hr/>
</details>

### 2026-08-22 (2x / 50.00%)


#### external (2x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-08-22 03:29:33 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18869/pull-kubevirt-e2e-kind-1.35-vgpu-1.9/2091004683941318656#1:build-log.txt%3A1376)
<details>
<summary>all...</summary>

* _2026-08-22 07:24:09 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18884/pull-kubevirt-e2e-kind-1.35-vgpu/2091063713351077888#1:build-log.txt%3A1271)

* _2026-08-22 03:29:33 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18869/pull-kubevirt-e2e-kind-1.35-vgpu-1.9/2091004683941318656#1:build-log.txt%3A1376)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (4x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (3x / 75.00%) </summary>

<hr/>

**3x**: _2026-08-24 19:05:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18867/pull-kubevirt-e2e-k8s-1.36-sig-network/2091964976599142400#1:build-log.txt%3A1875)
<details>
<summary>all...</summary>

* _2026-08-24 19:05:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18867/pull-kubevirt-e2e-k8s-1.36-sig-network/2091964976599142400#1:build-log.txt%3A1875)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
19:15:28: selecting podman as container runtime
19:16:06: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network is being used by the following container(s): edbc20eab0ffa24b2ea7039833cb3d89d336f25ca216ea57fa57b8767cd0f141: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-08-22 07:24:09 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18884/pull-kubevirt-e2e-kind-1.35-vgpu/2091063713351077888#1:build-log.txt%3A1271)
<details><summary>context</summary>
<pre>07:31:56: &#43; echo &#39;FATAL: Could not find available GPUs on host&#39;
07:31:56: FATAL: Could not find available GPUs on host
07:31:56: &#43; return 1
make: *** [Makefile:174: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-08-22 03:29:33 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18869/pull-kubevirt-e2e-kind-1.35-vgpu-1.9/2091004683941318656#1:build-log.txt%3A1376)
<details><summary>context</summary>
<pre>03:39:16: &#43; echo &#39;FATAL: Could not find available GPUs on host&#39;
03:39:16: FATAL: Could not find available GPUs on host
03:39:16: &#43; return 1
make: *** [Makefile:174: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-08-25 07:47:09 &#43;0000 UTC_: <code>Error: cleaning up container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888: unmounting container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 storage: cleaning up container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 storage: unmounting container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 root filesystem: deleting layer &#34;90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9 /var/lib/containers/storage/overlay/tempdirs/temp-dir-2941215159/1-90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18875/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2092154790824906752#1:build-log.txt%3A589)
<details>
<summary>all...</summary>

* _2026-08-25 07:47:09 &#43;0000 UTC_: <code>Error: cleaning up container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888: unmounting container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 storage: cleaning up container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 storage: unmounting container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 root filesystem: deleting layer &#34;90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9 /var/lib/containers/storage/overlay/tempdirs/temp-dir-2941215159/1-90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18875/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2092154790824906752#1:build-log.txt%3A589)
<details><summary>context</summary>
<pre>time=&#34;2026-08-25T07:54:47Z&#34; level=warning msg=&#34;Found incomplete layer \&#34;90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9\&#34;, deleting it&#34;
time=&#34;2026-08-25T07:54:47Z&#34; level=warning msg=&#34;Found incomplete layer \&#34;90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9\&#34;, deleting it&#34;
time=&#34;2026-08-25T07:54:47Z&#34; level=error msg=&#34;cleaning up storage: removing container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 root filesystem: deleting layer \&#34;90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9\&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9 /var/lib/containers/storage/overlay/tempdirs/temp-dir-3351095486/1-90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9: invalid cross-device link&#34;
Error: cleaning up container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888: unmounting container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 storage: cleaning up container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 storage: unmounting container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 root filesystem: deleting layer &#34;90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9 /var/lib/containers/storage/overlay/tempdirs/temp-dir-2941215159/1-90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9: invalid cross-device link
time=&#34;2026-08-25T07:54:47Z&#34; level=warning msg=&#34;Found incomplete layer \&#34;90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9\&#34;, deleting it&#34;
/usr/local/bin/runner.sh: line 50: wait: pid 1220 is not a child of this shell
================================================================================</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (3x / 75.00%)


#### external (3x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-08-24 19:05:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18867/pull-kubevirt-e2e-k8s-1.36-sig-network/2091964976599142400#1:build-log.txt%3A1875)
<details>
<summary>all...</summary>

* _2026-08-24 19:05:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18867/pull-kubevirt-e2e-k8s-1.36-sig-network/2091964976599142400#1:build-log.txt%3A1875)

* _2026-08-22 07:24:09 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18884/pull-kubevirt-e2e-kind-1.35-vgpu/2091063713351077888#1:build-log.txt%3A1271)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-08-25 07:47:09 &#43;0000 UTC_: <code>Error: cleaning up container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888: unmounting container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 storage: cleaning up container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 storage: unmounting container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 root filesystem: deleting layer &#34;90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9 /var/lib/containers/storage/overlay/tempdirs/temp-dir-2941215159/1-90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18875/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2092154790824906752#1:build-log.txt%3A589)
<details>
<summary>all...</summary>

* _2026-08-25 07:47:09 &#43;0000 UTC_: <code>Error: cleaning up container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888: unmounting container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 storage: cleaning up container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 storage: unmounting container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 root filesystem: deleting layer &#34;90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9 /var/lib/containers/storage/overlay/tempdirs/temp-dir-2941215159/1-90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18875/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2092154790824906752#1:build-log.txt%3A589)

</details>

<hr/>
</details>

### release-1.9 (1x / 25.00%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-08-22 03:29:33 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18869/pull-kubevirt-e2e-kind-1.35-vgpu-1.9/2091004683941318656#1:build-log.txt%3A1376)
<details>
<summary>all...</summary>

* _2026-08-22 03:29:33 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18869/pull-kubevirt-e2e-kind-1.35-vgpu-1.9/2091004683941318656#1:build-log.txt%3A1376)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-network (2x / 50.00%)


#### external (2x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-08-24 19:05:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18867/pull-kubevirt-e2e-k8s-1.36-sig-network/2091964976599142400#1:build-log.txt%3A1875)
<details>
<summary>all...</summary>

* _2026-08-24 19:05:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18867/pull-kubevirt-e2e-k8s-1.36-sig-network/2091964976599142400#1:build-log.txt%3A1875)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-08-25 07:47:09 &#43;0000 UTC_: <code>Error: cleaning up container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888: unmounting container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 storage: cleaning up container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 storage: unmounting container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 root filesystem: deleting layer &#34;90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9 /var/lib/containers/storage/overlay/tempdirs/temp-dir-2941215159/1-90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18875/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2092154790824906752#1:build-log.txt%3A589)
<details>
<summary>all...</summary>

* _2026-08-25 07:47:09 &#43;0000 UTC_: <code>Error: cleaning up container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888: unmounting container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 storage: cleaning up container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 storage: unmounting container c135e19365487620125c87f756daf5536ff0568e6f7270a8b45284bb10e02888 root filesystem: deleting layer &#34;90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9 /var/lib/containers/storage/overlay/tempdirs/temp-dir-2941215159/1-90d6594afb8eb63c35856e7e7add60b78121cc657258e336db61104ef1a6a3b9: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18875/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2092154790824906752#1:build-log.txt%3A589)

</details>

<hr/>
</details>

### sig-compute (2x / 50.00%)


#### external (2x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-08-22 03:29:33 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18869/pull-kubevirt-e2e-kind-1.35-vgpu-1.9/2091004683941318656#1:build-log.txt%3A1376)
<details>
<summary>all...</summary>

* _2026-08-22 07:24:09 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18884/pull-kubevirt-e2e-kind-1.35-vgpu/2091063713351077888#1:build-log.txt%3A1271)

* _2026-08-22 03:29:33 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18869/pull-kubevirt-e2e-kind-1.35-vgpu-1.9/2091004683941318656#1:build-log.txt%3A1376)

</details>

<hr/>
</details>

Last updated: 2026-08-29 00:14:39
