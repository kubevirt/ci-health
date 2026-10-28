



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-10-26 (1x / 50.00%)


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

### 2026-10-21 (1x / 50.00%)


#### external (1x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-10-21 12:56:21 &#43;0000 UTC_: <code>Error: cleaning up container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e: unmounting container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e storage: cleaning up container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e storage: unmounting container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e root filesystem: deleting layer &#34;fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a /var/lib/containers/storage/overlay/tempdirs/temp-dir-3166173718/1-fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19064/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2112890491417661440#1:build-log.txt%3A591)
<details>
<summary>all...</summary>

* _2026-10-21 12:56:21 &#43;0000 UTC_: <code>Error: cleaning up container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e: unmounting container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e storage: cleaning up container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e storage: unmounting container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e root filesystem: deleting layer &#34;fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a /var/lib/containers/storage/overlay/tempdirs/temp-dir-3166173718/1-fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19064/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2112890491417661440#1:build-log.txt%3A591)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (2x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-10-21 12:56:21 &#43;0000 UTC_: <code>Error: cleaning up container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e: unmounting container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e storage: cleaning up container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e storage: unmounting container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e root filesystem: deleting layer &#34;fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a /var/lib/containers/storage/overlay/tempdirs/temp-dir-3166173718/1-fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19064/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2112890491417661440#1:build-log.txt%3A591)
<details>
<summary>all...</summary>

* _2026-10-21 12:56:21 &#43;0000 UTC_: <code>Error: cleaning up container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e: unmounting container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e storage: cleaning up container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e storage: unmounting container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e root filesystem: deleting layer &#34;fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a /var/lib/containers/storage/overlay/tempdirs/temp-dir-3166173718/1-fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19064/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2112890491417661440#1:build-log.txt%3A591)
<details><summary>context</summary>
<pre>time=&#34;2026-10-21T13:06:01Z&#34; level=warning msg=&#34;Found incomplete layer \&#34;fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a\&#34;, deleting it&#34;
time=&#34;2026-10-21T13:06:01Z&#34; level=warning msg=&#34;Found incomplete layer \&#34;fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a\&#34;, deleting it&#34;
time=&#34;2026-10-21T13:06:01Z&#34; level=error msg=&#34;cleaning up storage: removing container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e root filesystem: deleting layer \&#34;fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a\&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a /var/lib/containers/storage/overlay/tempdirs/temp-dir-2592037309/1-fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a: invalid cross-device link&#34;
Error: cleaning up container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e: unmounting container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e storage: cleaning up container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e storage: unmounting container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e root filesystem: deleting layer &#34;fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a /var/lib/containers/storage/overlay/tempdirs/temp-dir-3166173718/1-fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a: invalid cross-device link
time=&#34;2026-10-21T13:06:01Z&#34; level=warning msg=&#34;Found incomplete layer \&#34;fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a\&#34;, deleting it&#34;
/usr/local/bin/runner.sh: line 50: wait: pid 1020 is not a child of this shell
================================================================================</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 50.00%) </summary>

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

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (2x / 100.00%)


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

**1x**: _2026-10-21 12:56:21 &#43;0000 UTC_: <code>Error: cleaning up container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e: unmounting container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e storage: cleaning up container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e storage: unmounting container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e root filesystem: deleting layer &#34;fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a /var/lib/containers/storage/overlay/tempdirs/temp-dir-3166173718/1-fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19064/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2112890491417661440#1:build-log.txt%3A591)
<details>
<summary>all...</summary>

* _2026-10-21 12:56:21 &#43;0000 UTC_: <code>Error: cleaning up container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e: unmounting container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e storage: cleaning up container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e storage: unmounting container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e root filesystem: deleting layer &#34;fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a /var/lib/containers/storage/overlay/tempdirs/temp-dir-3166173718/1-fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19064/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2112890491417661440#1:build-log.txt%3A591)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-network (2x / 100.00%)


#### external (2x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-10-21 12:56:21 &#43;0000 UTC_: <code>Error: cleaning up container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e: unmounting container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e storage: cleaning up container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e storage: unmounting container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e root filesystem: deleting layer &#34;fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a /var/lib/containers/storage/overlay/tempdirs/temp-dir-3166173718/1-fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19064/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2112890491417661440#1:build-log.txt%3A591)
<details>
<summary>all...</summary>

* _2026-10-21 12:56:21 &#43;0000 UTC_: <code>Error: cleaning up container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e: unmounting container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e storage: cleaning up container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e storage: unmounting container 5fd55014d234219b6b5c86e212763b12e253f6c8052319ab7175ef955350873e root filesystem: deleting layer &#34;fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a /var/lib/containers/storage/overlay/tempdirs/temp-dir-3166173718/1-fa8bf1d9b9284b1fee0a4d112c6650760bc1213c54ea2536108ad05f64f3950a: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19064/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2112890491417661440#1:build-log.txt%3A591)

</details>

<hr/>
</details>
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

Last updated: 2026-10-28 09:12:02
