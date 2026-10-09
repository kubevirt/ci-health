



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-10-07 (3x / 60.00%)


#### external (2x / 66.67%)

<details>
<summary> container image pull failure in context (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-10-07 13:17:20 &#43;0000 UTC_: <code>Error: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab root filesystem: deleting layer &#34;927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8 /var/lib/containers/storage/overlay/tempdirs/temp-dir-992205427/1-927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19348/pull-kubevirt-e2e-k8s-1.36-sig-compute/2107740041999552512#1:build-log.txt%3A572)
<details>
<summary>all...</summary>

* _2026-10-07 13:17:20 &#43;0000 UTC_: <code>Error: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab root filesystem: deleting layer &#34;927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8 /var/lib/containers/storage/overlay/tempdirs/temp-dir-992205427/1-927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19348/pull-kubevirt-e2e-k8s-1.36-sig-compute/2107740041999552512#1:build-log.txt%3A572)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-10-07 19:29:15 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19217/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2107894287529152512#1:build-log.txt%3A1855)
<details>
<summary>all...</summary>

* _2026-10-07 19:29:15 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19217/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2107894287529152512#1:build-log.txt%3A1855)

</details>

<hr/>
</details>

#### internal (1x / 33.33%)

<details>
<summary> make cluster lifecycle target failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-10-07 10:56:37 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19348/pull-kubevirt-e2e-k8s-1.37-sig-network/2107740042326708224#1:build-log.txt%3A997)
<details>
<summary>all...</summary>

* _2026-10-07 10:56:37 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19348/pull-kubevirt-e2e-k8s-1.37-sig-network/2107740042326708224#1:build-log.txt%3A997)

</details>

<hr/>
</details>

### 2026-10-06 (1x / 20.00%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-10-06 09:03:34 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19325/pull-kubevirt-e2e-k8s-1.35-sig-operator/2107388856490790912#1:build-log.txt%3A1102)
<details>
<summary>all...</summary>

* _2026-10-06 09:03:34 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19325/pull-kubevirt-e2e-k8s-1.35-sig-operator/2107388856490790912#1:build-log.txt%3A1102)

</details>

<hr/>
</details>

### 2026-10-05 (1x / 20.00%)


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

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (4x / 80.00%)

<details>
<summary> container image pull failure in context (2x / 40.00%) </summary>

<hr/>

**1x**: _2026-10-07 13:17:20 &#43;0000 UTC_: <code>Error: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab root filesystem: deleting layer &#34;927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8 /var/lib/containers/storage/overlay/tempdirs/temp-dir-992205427/1-927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19348/pull-kubevirt-e2e-k8s-1.36-sig-compute/2107740041999552512#1:build-log.txt%3A572)
<details>
<summary>all...</summary>

* _2026-10-07 13:17:20 &#43;0000 UTC_: <code>Error: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab root filesystem: deleting layer &#34;927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8 /var/lib/containers/storage/overlay/tempdirs/temp-dir-992205427/1-927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19348/pull-kubevirt-e2e-k8s-1.36-sig-compute/2107740041999552512#1:build-log.txt%3A572)
<details><summary>context</summary>
<pre>time=&#34;2026-10-07T13:25:53Z&#34; level=warning msg=&#34;Found incomplete layer \&#34;927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8\&#34;, deleting it&#34;
time=&#34;2026-10-07T13:25:53Z&#34; level=warning msg=&#34;Found incomplete layer \&#34;927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8\&#34;, deleting it&#34;
time=&#34;2026-10-07T13:25:53Z&#34; level=error msg=&#34;cleaning up storage: removing container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab root filesystem: deleting layer \&#34;927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8\&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8 /var/lib/containers/storage/overlay/tempdirs/temp-dir-2711204353/1-927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8: invalid cross-device link&#34;
Error: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab root filesystem: deleting layer &#34;927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8 /var/lib/containers/storage/overlay/tempdirs/temp-dir-992205427/1-927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8: invalid cross-device link
time=&#34;2026-10-07T13:25:53Z&#34; level=warning msg=&#34;Found incomplete layer \&#34;927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8\&#34;, deleting it&#34;
/usr/local/bin/runner.sh: line 50: wait: pid 1021 is not a child of this shell
================================================================================</pre>
</details>


</details>

<hr/>

**1x**: _2026-10-05 06:50:11 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17974/pull-kubevirt-e2e-kind-1.37-vgpu/2107000242497916928#1:build-log.txt%3A467)
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


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 40.00%) </summary>

<hr/>

**2x**: _2026-10-07 19:29:15 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19217/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2107894287529152512#1:build-log.txt%3A1855)
<details>
<summary>all...</summary>

* _2026-10-07 19:29:15 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19217/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2107894287529152512#1:build-log.txt%3A1855)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
19:34:54: selecting podman as container runtime
19:35:26: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network is being used by the following container(s): 45098d90b12c716eff4f957d92068f5eec2ea717a04e155a6f14f31e83e8f363: volume is being used
make: *** [Makefile:179: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-06 09:03:34 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19325/pull-kubevirt-e2e-k8s-1.35-sig-operator/2107388856490790912#1:build-log.txt%3A1102)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
09:08:45: selecting podman as container runtime
09:09:06: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-operator is being used by the following container(s): 9554fa5958a18856196e7ff53c99b2266182ea388c28da183f1a18cbc634387d: volume is being used
make: *** [Makefile:179: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>

### internal (1x / 20.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-10-07 10:56:37 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19348/pull-kubevirt-e2e-k8s-1.37-sig-network/2107740042326708224#1:build-log.txt%3A997)
<details>
<summary>all...</summary>

* _2026-10-07 10:56:37 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19348/pull-kubevirt-e2e-k8s-1.37-sig-network/2107740042326708224#1:build-log.txt%3A997)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
11:19:59: selecting podman as container runtime
11:19:59: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.37-sig-network is being used by the following container(s): 86408ef1fcadc70e3e18b132b50f1f00d7b08ebe90aa0bcf35a6479c3226f7c2: volume is being used
make: *** [Makefile:179: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (5x / 100.00%)


#### external (4x / 80.00%)

<details>
<summary> container image pull failure in context (2x / 40.00%) </summary>

<hr/>

**1x**: _2026-10-07 13:17:20 &#43;0000 UTC_: <code>Error: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab root filesystem: deleting layer &#34;927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8 /var/lib/containers/storage/overlay/tempdirs/temp-dir-992205427/1-927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19348/pull-kubevirt-e2e-k8s-1.36-sig-compute/2107740041999552512#1:build-log.txt%3A572)
<details>
<summary>all...</summary>

* _2026-10-07 13:17:20 &#43;0000 UTC_: <code>Error: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab root filesystem: deleting layer &#34;927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8 /var/lib/containers/storage/overlay/tempdirs/temp-dir-992205427/1-927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19348/pull-kubevirt-e2e-k8s-1.36-sig-compute/2107740041999552512#1:build-log.txt%3A572)

</details>

<hr/>

**1x**: _2026-10-05 06:50:11 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17974/pull-kubevirt-e2e-kind-1.37-vgpu/2107000242497916928#1:build-log.txt%3A467)
<details>
<summary>all...</summary>

* _2026-10-05 06:50:11 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17974/pull-kubevirt-e2e-kind-1.37-vgpu/2107000242497916928#1:build-log.txt%3A467)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 40.00%) </summary>

<hr/>

**2x**: _2026-10-07 19:29:15 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19217/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2107894287529152512#1:build-log.txt%3A1855)
<details>
<summary>all...</summary>

* _2026-10-07 19:29:15 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19217/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2107894287529152512#1:build-log.txt%3A1855)

* _2026-10-06 09:03:34 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19325/pull-kubevirt-e2e-k8s-1.35-sig-operator/2107388856490790912#1:build-log.txt%3A1102)

</details>

<hr/>
</details>

#### internal (1x / 20.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-10-07 10:56:37 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19348/pull-kubevirt-e2e-k8s-1.37-sig-network/2107740042326708224#1:build-log.txt%3A997)
<details>
<summary>all...</summary>

* _2026-10-07 10:56:37 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19348/pull-kubevirt-e2e-k8s-1.37-sig-network/2107740042326708224#1:build-log.txt%3A997)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-network (2x / 40.00%)


#### external (1x / 50.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-10-07 19:29:15 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19217/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2107894287529152512#1:build-log.txt%3A1855)
<details>
<summary>all...</summary>

* _2026-10-07 19:29:15 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19217/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2107894287529152512#1:build-log.txt%3A1855)

</details>

<hr/>
</details>

#### internal (1x / 50.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-10-07 10:56:37 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19348/pull-kubevirt-e2e-k8s-1.37-sig-network/2107740042326708224#1:build-log.txt%3A997)
<details>
<summary>all...</summary>

* _2026-10-07 10:56:37 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19348/pull-kubevirt-e2e-k8s-1.37-sig-network/2107740042326708224#1:build-log.txt%3A997)

</details>

<hr/>
</details>

### sig-compute (3x / 60.00%)


#### external (3x / 100.00%)

<details>
<summary> container image pull failure in context (2x / 66.67%) </summary>

<hr/>

**1x**: _2026-10-07 13:17:20 &#43;0000 UTC_: <code>Error: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab root filesystem: deleting layer &#34;927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8 /var/lib/containers/storage/overlay/tempdirs/temp-dir-992205427/1-927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19348/pull-kubevirt-e2e-k8s-1.36-sig-compute/2107740041999552512#1:build-log.txt%3A572)
<details>
<summary>all...</summary>

* _2026-10-07 13:17:20 &#43;0000 UTC_: <code>Error: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab root filesystem: deleting layer &#34;927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8 /var/lib/containers/storage/overlay/tempdirs/temp-dir-992205427/1-927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19348/pull-kubevirt-e2e-k8s-1.36-sig-compute/2107740041999552512#1:build-log.txt%3A572)

</details>

<hr/>

**1x**: _2026-10-05 06:50:11 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17974/pull-kubevirt-e2e-kind-1.37-vgpu/2107000242497916928#1:build-log.txt%3A467)
<details>
<summary>all...</summary>

* _2026-10-05 06:50:11 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17974/pull-kubevirt-e2e-kind-1.37-vgpu/2107000242497916928#1:build-log.txt%3A467)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-10-06 09:03:34 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19325/pull-kubevirt-e2e-k8s-1.35-sig-operator/2107388856490790912#1:build-log.txt%3A1102)
<details>
<summary>all...</summary>

* _2026-10-06 09:03:34 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19325/pull-kubevirt-e2e-k8s-1.35-sig-operator/2107388856490790912#1:build-log.txt%3A1102)

</details>

<hr/>
</details>

Last updated: 2026-10-09 15:16:33
