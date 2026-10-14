



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-10-13 (2x / 28.57%)


#### external (2x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-10-13 12:08:33 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19312/pull-kubevirt-e2e-k8s-1.37-sig-network/2109978051877015552#1:build-log.txt%3A1927)
<details>
<summary>all...</summary>

* _2026-10-13 18:43:26 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19312/pull-kubevirt-e2e-k8s-1.35-sig-operator/2109978050710999040#1:build-log.txt%3A1161)

* _2026-10-13 12:08:33 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19312/pull-kubevirt-e2e-k8s-1.37-sig-network/2109978051877015552#1:build-log.txt%3A1927)

</details>

<hr/>
</details>

### 2026-10-12 (2x / 28.57%)


#### external (2x / 100.00%)

<details>
<summary> download failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-10-12 12:33:22 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18853/pull-kubevirt-e2e-k8s-1.37-sig-operator/2109622445727027200#1:build-log.txt%3A380)
<details>
<summary>all...</summary>

* _2026-10-12 12:33:22 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18853/pull-kubevirt-e2e-k8s-1.37-sig-operator/2109622445727027200#1:build-log.txt%3A380)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-10-12 12:01:09 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19099/pull-kubevirt-e2e-k8s-1.36-sig-network/2109609064953548800#1:build-log.txt%3A1178)
<details>
<summary>all...</summary>

* _2026-10-12 12:01:09 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19099/pull-kubevirt-e2e-k8s-1.36-sig-network/2109609064953548800#1:build-log.txt%3A1178)

</details>

<hr/>
</details>

### 2026-10-11 (1x / 14.29%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-10-11 14:33:33 &#43;0000 UTC_: <code>make: *** [Makefile:191: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19401/pull-kubevirt-e2e-kind-1.37-sev/2109291175335694336#1:build-log.txt%3A3114)
<details>
<summary>all...</summary>

* _2026-10-11 14:33:33 &#43;0000 UTC_: <code>make: *** [Makefile:191: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19401/pull-kubevirt-e2e-kind-1.37-sev/2109291175335694336#1:build-log.txt%3A3114)

</details>

<hr/>
</details>

### 2026-10-07 (2x / 28.57%)


#### external (2x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-10-07 13:17:20 &#43;0000 UTC_: <code>Error: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab root filesystem: deleting layer &#34;927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8 /var/lib/containers/storage/overlay/tempdirs/temp-dir-992205427/1-927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19348/pull-kubevirt-e2e-k8s-1.36-sig-compute/2107740041999552512#1:build-log.txt%3A572)
<details>
<summary>all...</summary>

* _2026-10-07 13:17:20 &#43;0000 UTC_: <code>Error: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab root filesystem: deleting layer &#34;927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8 /var/lib/containers/storage/overlay/tempdirs/temp-dir-992205427/1-927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19348/pull-kubevirt-e2e-k8s-1.36-sig-compute/2107740041999552512#1:build-log.txt%3A572)

</details>

<hr/>
</details>
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

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (7x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (5x / 71.43%) </summary>

<hr/>

**5x**: _2026-10-11 14:33:33 &#43;0000 UTC_: <code>make: *** [Makefile:191: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19401/pull-kubevirt-e2e-kind-1.37-sev/2109291175335694336#1:build-log.txt%3A3114)
<details>
<summary>all...</summary>

* _2026-10-13 18:43:26 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19312/pull-kubevirt-e2e-k8s-1.35-sig-operator/2109978050710999040#1:build-log.txt%3A1161)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
18:50:32: selecting podman as container runtime
18:50:33: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-operator is being used by the following container(s): ca56445905b9a936d77261cd720ab5f81b785d348178263943b4da3e4f16237d: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-13 12:08:33 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19312/pull-kubevirt-e2e-k8s-1.37-sig-network/2109978051877015552#1:build-log.txt%3A1927)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:13:48: selecting podman as container runtime
12:14:38: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.37-sig-network is being used by the following container(s): efef4d1a42c9473b3967245117e2b6719e67862ea6efb9b46793b403b08c0aba: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-12 12:01:09 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19099/pull-kubevirt-e2e-k8s-1.36-sig-network/2109609064953548800#1:build-log.txt%3A1178)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:05:44: selecting podman as container runtime
12:05:44: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network is being used by the following container(s): dbf9941fae8fe6b9656b6a66e0c41d9bc89b6a2853850486a488fcb11e5920ba: volume is being used
make: *** [Makefile:179: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-11 14:33:33 &#43;0000 UTC_: <code>make: *** [Makefile:191: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19401/pull-kubevirt-e2e-kind-1.37-sev/2109291175335694336#1:build-log.txt%3A3114)
<details><summary>context</summary>
<pre>15:13:18: failed to fetch vm exports: the server could not find the requested resource (get virtualmachinepools.pool.kubevirt.io)
15:13:18: failed to fetch migration policies: the server could not find the requested resource (get migrationpolicies.migrations.kubevirt.io)
15:13:18: &#43; rm -f /tmp/kubevirt.deploy.Bkos
make: *** [Makefile:191: cluster-sync] Error 1
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>


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


</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 14.29%) </summary>

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
</details>
<details>
<summary> download failure in context (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-10-12 12:33:22 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18853/pull-kubevirt-e2e-k8s-1.37-sig-operator/2109622445727027200#1:build-log.txt%3A380)
<details>
<summary>all...</summary>

* _2026-10-12 12:33:22 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18853/pull-kubevirt-e2e-k8s-1.37-sig-operator/2109622445727027200#1:build-log.txt%3A380)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:35:59: selecting podman as container runtime
12:35:59: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.37-sig-operator is being used by the following container(s): 86a58ba354853d642b7d7eb237c837c2a814da8ecb879bf4cb427403a7972422: volume is being used
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


### main (7x / 100.00%)


#### external (7x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (5x / 71.43%) </summary>

<hr/>

**5x**: _2026-10-11 14:33:33 &#43;0000 UTC_: <code>make: *** [Makefile:191: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19401/pull-kubevirt-e2e-kind-1.37-sev/2109291175335694336#1:build-log.txt%3A3114)
<details>
<summary>all...</summary>

* _2026-10-13 18:43:26 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19312/pull-kubevirt-e2e-k8s-1.35-sig-operator/2109978050710999040#1:build-log.txt%3A1161)

* _2026-10-13 12:08:33 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19312/pull-kubevirt-e2e-k8s-1.37-sig-network/2109978051877015552#1:build-log.txt%3A1927)

* _2026-10-12 12:01:09 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19099/pull-kubevirt-e2e-k8s-1.36-sig-network/2109609064953548800#1:build-log.txt%3A1178)

* _2026-10-11 14:33:33 &#43;0000 UTC_: <code>make: *** [Makefile:191: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19401/pull-kubevirt-e2e-kind-1.37-sev/2109291175335694336#1:build-log.txt%3A3114)

* _2026-10-07 19:29:15 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19217/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2107894287529152512#1:build-log.txt%3A1855)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-10-07 13:17:20 &#43;0000 UTC_: <code>Error: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab root filesystem: deleting layer &#34;927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8 /var/lib/containers/storage/overlay/tempdirs/temp-dir-992205427/1-927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19348/pull-kubevirt-e2e-k8s-1.36-sig-compute/2107740041999552512#1:build-log.txt%3A572)
<details>
<summary>all...</summary>

* _2026-10-07 13:17:20 &#43;0000 UTC_: <code>Error: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab root filesystem: deleting layer &#34;927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8 /var/lib/containers/storage/overlay/tempdirs/temp-dir-992205427/1-927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19348/pull-kubevirt-e2e-k8s-1.36-sig-compute/2107740041999552512#1:build-log.txt%3A572)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-10-12 12:33:22 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18853/pull-kubevirt-e2e-k8s-1.37-sig-operator/2109622445727027200#1:build-log.txt%3A380)
<details>
<summary>all...</summary>

* _2026-10-12 12:33:22 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18853/pull-kubevirt-e2e-k8s-1.37-sig-operator/2109622445727027200#1:build-log.txt%3A380)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (4x / 57.14%)


#### external (4x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 50.00%) </summary>

<hr/>

**2x**: _2026-10-11 14:33:33 &#43;0000 UTC_: <code>make: *** [Makefile:191: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19401/pull-kubevirt-e2e-kind-1.37-sev/2109291175335694336#1:build-log.txt%3A3114)
<details>
<summary>all...</summary>

* _2026-10-13 18:43:26 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19312/pull-kubevirt-e2e-k8s-1.35-sig-operator/2109978050710999040#1:build-log.txt%3A1161)

* _2026-10-11 14:33:33 &#43;0000 UTC_: <code>make: *** [Makefile:191: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19401/pull-kubevirt-e2e-kind-1.37-sev/2109291175335694336#1:build-log.txt%3A3114)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-10-12 12:33:22 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18853/pull-kubevirt-e2e-k8s-1.37-sig-operator/2109622445727027200#1:build-log.txt%3A380)
<details>
<summary>all...</summary>

* _2026-10-12 12:33:22 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18853/pull-kubevirt-e2e-k8s-1.37-sig-operator/2109622445727027200#1:build-log.txt%3A380)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-10-07 13:17:20 &#43;0000 UTC_: <code>Error: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab root filesystem: deleting layer &#34;927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8 /var/lib/containers/storage/overlay/tempdirs/temp-dir-992205427/1-927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19348/pull-kubevirt-e2e-k8s-1.36-sig-compute/2107740041999552512#1:build-log.txt%3A572)
<details>
<summary>all...</summary>

* _2026-10-07 13:17:20 &#43;0000 UTC_: <code>Error: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: cleaning up container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab storage: unmounting container 85be7e8e91579ec0a353caaf502fb9cb791271e721b834f611200ca2d394afab root filesystem: deleting layer &#34;927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8 /var/lib/containers/storage/overlay/tempdirs/temp-dir-992205427/1-927cc1fa2e1a1991b21727aa224b5e020a668c269a0d400ba405efc0001a43e8: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19348/pull-kubevirt-e2e-k8s-1.36-sig-compute/2107740041999552512#1:build-log.txt%3A572)

</details>

<hr/>
</details>

### sig-network (3x / 42.86%)


#### external (3x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (3x / 100.00%) </summary>

<hr/>

**3x**: _2026-10-12 12:01:09 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19099/pull-kubevirt-e2e-k8s-1.36-sig-network/2109609064953548800#1:build-log.txt%3A1178)
<details>
<summary>all...</summary>

* _2026-10-13 12:08:33 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19312/pull-kubevirt-e2e-k8s-1.37-sig-network/2109978051877015552#1:build-log.txt%3A1927)

* _2026-10-12 12:01:09 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19099/pull-kubevirt-e2e-k8s-1.36-sig-network/2109609064953548800#1:build-log.txt%3A1178)

* _2026-10-07 19:29:15 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19217/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2107894287529152512#1:build-log.txt%3A1855)

</details>

<hr/>
</details>

Last updated: 2026-10-14 12:16:43
