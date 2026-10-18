



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-10-16 (4x / 11.11%)


#### external (4x / 100.00%)

<details>
<summary> download failure in context (4x / 100.00%) </summary>

<hr/>

**3x**: _2026-10-16 10:10:12 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations-1.9/2111036843460923392#1:build-log.txt%3A1305)
<details>
<summary>all...</summary>

* _2026-10-16 10:10:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2111036843687415808#1:build-log.txt%3A777)

* _2026-10-16 10:10:12 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations-1.9/2111036843460923392#1:build-log.txt%3A1305)

* _2026-10-16 00:02:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.9/2110883575975907328#1:build-log.txt%3A781)

</details>

<hr/>

**1x**: _2026-10-16 00:01:12 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110883575862661120#1:build-log.txt%3A619)
<details>
<summary>all...</summary>

* _2026-10-16 00:01:12 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110883575862661120#1:build-log.txt%3A619)

</details>

<hr/>
</details>

### 2026-10-15 (4x / 11.11%)


#### external (4x / 100.00%)

<details>
<summary> download failure in context (3x / 75.00%) </summary>

<hr/>

**2x**: _2026-10-15 22:10:37 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19476/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2110841617408593920#1:build-log.txt%3A631)
<details>
<summary>all...</summary>

* _2026-10-15 22:10:37 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19476/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2110841617408593920#1:build-log.txt%3A631)

* _2026-10-15 20:59:54 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.9/2110837556139528192#1:build-log.txt%3A931)

</details>

<hr/>

**1x**: _2026-10-15 20:58:13 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110837556055642112#1:build-log.txt%3A551)
<details>
<summary>all...</summary>

* _2026-10-15 20:58:13 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110837556055642112#1:build-log.txt%3A551)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-10-15 00:02:06 &#43;0000 UTC_: <code>make: *** [Makefile:194: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19454/pull-kubevirt-e2e-kind-1.37-sig-compute-arm64/2110521441312051200#1:build-log.txt%3A2485)
<details>
<summary>all...</summary>

* _2026-10-15 00:02:06 &#43;0000 UTC_: <code>make: *** [Makefile:194: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19454/pull-kubevirt-e2e-kind-1.37-sig-compute-arm64/2110521441312051200#1:build-log.txt%3A2485)

</details>

<hr/>
</details>

### 2026-10-14 (23x / 63.89%)


#### external (23x / 100.00%)

<details>
<summary> download failure in context (23x / 100.00%) </summary>

<hr/>

**20x**: _2026-10-14 10:00:15 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-windows2016-1.9/2110309509023928320#1:build-log.txt%3A963)
<details>
<summary>all...</summary>

* _2026-10-14 10:08:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations-1.9/2110309667371487232#1:build-log.txt%3A1435)

* _2026-10-14 10:08:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2110309667509899264#1:build-log.txt%3A922)

* _2026-10-14 10:08:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-operator-1.9/2110309525570457600#1:build-log.txt%3A1120)

* _2026-10-14 10:07:32 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2110309524777734144#1:build-log.txt%3A1016)

* _2026-10-14 10:07:08 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-compute-1.9/2110309523909513216#1:build-log.txt%3A1065)

* _2026-10-14 10:06:48 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-storage-1.9/2110309523062263808#1:build-log.txt%3A1299)

* _2026-10-14 10:03:08 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-network-1.9/2110309522235985920#1:build-log.txt%3A1143)

* _2026-10-14 10:02:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.35-sig-operator-1.9/2110309521405513728#1:build-log.txt%3A1373)

* _2026-10-14 10:02:11 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.9/2110309519731986432#1:build-log.txt%3A641)

* _2026-10-14 10:01:32 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.9/2110309520545681408#1:build-log.txt%3A782)

* _2026-10-14 10:00:54 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2110309512178044928#1:build-log.txt%3A785)

* _2026-10-14 10:00:30 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.9/2110309518083624960#1:build-log.txt%3A926)

* _2026-10-14 10:00:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.9/2110309516342988800#1:build-log.txt%3A674)

* _2026-10-14 10:00:15 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-windows2016-1.9/2110309509023928320#1:build-log.txt%3A963)

* _2026-10-14 10:00:13 &#43;0000 UTC_: <code>make: *** [Makefile:25: bazel-build] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-sig-performance-1.9/2110309508864544768#1:build-log.txt%3A372)

* _2026-10-14 10:00:13 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations-1.9/2110309513822212096#1:build-log.txt%3A971)

* _2026-10-14 10:00:07 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.35-sig-network-1.9/2110309518888931328#1:build-log.txt%3A665)

* _2026-10-14 10:00:07 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.9/2110309517177655296#1:build-log.txt%3A1310)

* _2026-10-14 10:00:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-sig-network-1.9/2110309515508322304#1:build-log.txt%3A1100)

* _2026-10-14 10:00:01 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated-1.9/2110309513012711424#1:build-log.txt%3A627)

</details>

<hr/>

**3x**: _2026-10-14 10:00:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110309514791096320#1:build-log.txt%3A523)
<details>
<summary>all...</summary>

* _2026-10-14 10:00:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110309514791096320#1:build-log.txt%3A523)

* _2026-10-14 09:59:58 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-kind-1.35-vgpu-1.9/2110309509128785920#1:build-log.txt%3A978)

* _2026-10-14 09:59:58 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-kind-1.34-sev-1.9/2110309509225254912#1:build-log.txt%3A1021)

</details>

<hr/>
</details>

### 2026-10-13 (2x / 5.56%)


#### external (2x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-10-13 18:43:26 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19312/pull-kubevirt-e2e-k8s-1.35-sig-operator/2109978050710999040#1:build-log.txt%3A1161)
<details>
<summary>all...</summary>

* _2026-10-13 18:43:26 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19312/pull-kubevirt-e2e-k8s-1.35-sig-operator/2109978050710999040#1:build-log.txt%3A1161)

* _2026-10-13 12:08:33 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19312/pull-kubevirt-e2e-k8s-1.37-sig-network/2109978051877015552#1:build-log.txt%3A1927)

</details>

<hr/>
</details>

### 2026-10-12 (2x / 5.56%)


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

### 2026-10-11 (1x / 2.78%)


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

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (36x / 100.00%)

<details>
<summary> download failure in context (31x / 86.11%) </summary>

<hr/>

**5x**: _2026-10-16 00:01:12 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110883575862661120#1:build-log.txt%3A619)
<details>
<summary>all...</summary>

* _2026-10-16 00:01:12 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110883575862661120#1:build-log.txt%3A619)
<details><summary>context</summary>
<pre>00:03:23: INFO: Elapsed time: 0.528s
00:03:23: INFO: 0 processes.
00:03:23: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-10-15 20:58:13 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110837556055642112#1:build-log.txt%3A551)
<details><summary>context</summary>
<pre>21:00:32: INFO: Elapsed time: 0.544s
21:00:32: INFO: 0 processes.
21:00:32: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-10-14 10:00:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110309514791096320#1:build-log.txt%3A523)
<details><summary>context</summary>
<pre>10:02:52: INFO: Elapsed time: 0.894s
10:02:52: INFO: 0 processes.
10:02:52: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-10-14 09:59:58 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-kind-1.35-vgpu-1.9/2110309509128785920#1:build-log.txt%3A978)
<details><summary>context</summary>
<pre>10:02:44: INFO: Elapsed time: 2.690s
10:02:44: INFO: 0 processes.
10:02:44: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-10-14 09:59:58 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-kind-1.34-sev-1.9/2110309509225254912#1:build-log.txt%3A1021)
<details><summary>context</summary>
<pre>10:01:40: INFO: Elapsed time: 0.382s
10:01:40: INFO: 0 processes.
10:01:40: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


</details>

<hr/>

**26x**: _2026-10-14 10:00:15 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-windows2016-1.9/2110309509023928320#1:build-log.txt%3A963)
<details>
<summary>all...</summary>

* _2026-10-16 10:10:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2111036843687415808#1:build-log.txt%3A777)
<details><summary>context</summary>
<pre>10:12:26: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
10:12:26: Writing manifest to image destination
10:12:26: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9 is being used by the following container(s): 467a6eb719dc3da741f97fab37b6e1b85e925f93c0146e6731165f84e21c5dfe: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-16 10:10:12 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations-1.9/2111036843460923392#1:build-log.txt%3A1305)
<details><summary>context</summary>
<pre>10:12:18: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
10:12:18: Writing manifest to image destination
10:12:18: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations-1.9 is being used by the following container(s): 3d27501f9bdd2134df87b47c6afd8ac2c650dd8e8a9bee9ce0bbe68833ba0684: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-16 00:02:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.9/2110883575975907328#1:build-log.txt%3A781)
<details><summary>context</summary>
<pre>00:04:55: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
00:04:55: Writing manifest to image destination
00:04:55: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-storage-1.9 is being used by the following container(s): 39765831b7aa01b5aa8d5a3a3c94ba6d5b219c91da2e46ac2b52794e7c8427c8: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-15 22:10:37 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19476/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2110841617408593920#1:build-log.txt%3A631)
<details><summary>context</summary>
<pre>22:17:30: Copying config sha256:4ae1b8fc6db781f8ee8a341cb31f613a7d4146ea63038e32aac66a1bdc104001
22:17:30: Writing manifest to image destination
22:17:31: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7 is being used by the following container(s): eeec1c4a69df24bf9418d3ca9f826a9fc4798c31cf049003c95981dadd43933e: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-15 20:59:54 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.9/2110837556139528192#1:build-log.txt%3A931)
<details><summary>context</summary>
<pre>21:03:09: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
21:03:09: Writing manifest to image destination
21:03:09: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-storage-1.9 is being used by the following container(s): 4f7410576852793ce0103630d46e0782c5a297c3ecf113bfb22501bb36bc4703: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-14 10:08:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations-1.9/2110309667371487232#1:build-log.txt%3A1435)
<details><summary>context</summary>
<pre>10:12:45: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
10:12:50: Writing manifest to image destination
10:12:54: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations-1.9 is being used by the following container(s): ad7210887362a1f03d07b24cb00ee87ec349044771de564a6eef2f967279c40a: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-14 10:08:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2110309667509899264#1:build-log.txt%3A922)
<details><summary>context</summary>
<pre>10:11:48: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
10:11:48: Writing manifest to image destination
10:11:49: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9 is being used by the following container(s): 60c3c13d7d54d1ce40a2e42b7beab16cd48831d77ffbf6b7764a23fbe810304f: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-14 10:08:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-operator-1.9/2110309525570457600#1:build-log.txt%3A1120)
<details><summary>context</summary>
<pre>10:11:45: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
10:11:45: Writing manifest to image destination
10:11:45: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-operator-1.9 is being used by the following container(s): 14a055f3c70c14ca9d7b7f79d64d953cbea8be262505a56118940aa8d8ff6796: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-14 10:07:32 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2110309524777734144#1:build-log.txt%3A1016)
<details><summary>context</summary>
<pre>10:10:58: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
10:10:58: Writing manifest to image destination
10:10:58: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9 is being used by the following container(s): 78f4bc66788886730538e122e9ac70dd4dc3d230759e46fadaf7f580a27a0d07: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-14 10:07:08 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-compute-1.9/2110309523909513216#1:build-log.txt%3A1065)
<details><summary>context</summary>
<pre>10:10:47: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
10:10:47: Writing manifest to image destination
10:10:47: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute-1.9 is being used by the following container(s): bafbe59133fec753eefbceee5a6af27ac5e4f71fa5c007cd31abd3a13f11a044: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-14 10:06:48 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-storage-1.9/2110309523062263808#1:build-log.txt%3A1299)
<details><summary>context</summary>
<pre>10:10:17: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
10:10:17: Writing manifest to image destination
10:10:17: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-storage-1.9 is being used by the following container(s): a19766fde2b9be5ed533d16c15961e2b308c36a5545cb418834b74f0f438e161: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-14 10:03:08 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-network-1.9/2110309522235985920#1:build-log.txt%3A1143)
<details><summary>context</summary>
<pre>10:07:07: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
10:07:07: Writing manifest to image destination
10:07:07: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network-1.9 is being used by the following container(s): 9ad13485988eae307cb853cf51ec2ec4a9f33a9a214cdd539a504d162abb362b: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-14 10:02:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.35-sig-operator-1.9/2110309521405513728#1:build-log.txt%3A1373)
<details><summary>context</summary>
<pre>10:06:54: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
10:06:54: Writing manifest to image destination
10:06:55: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-operator-1.9 is being used by the following container(s): 752876f52bed55ff4ccb643d58322c4148f16f077053c97031a7375a5145ba61: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-14 10:02:11 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.9/2110309519731986432#1:build-log.txt%3A641)
<details><summary>context</summary>
<pre>10:06:43: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
10:06:43: Writing manifest to image destination
10:06:43: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-storage-1.9 is being used by the following container(s): 607e0eb55aefd530ab4144a9390d4d22c89caddeba3a24c60e10f31ab98737e7: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-14 10:01:32 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.9/2110309520545681408#1:build-log.txt%3A782)
<details><summary>context</summary>
<pre>10:03:47: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
10:03:47: Writing manifest to image destination
10:03:47: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-compute-1.9 is being used by the following container(s): f3a4c71ae52ae3a709108133560623a01cfec9f648f4bc21f03a6349d586f65c: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-14 10:00:54 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2110309512178044928#1:build-log.txt%3A785)
<details><summary>context</summary>
<pre>10:05:23: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
10:05:23: Writing manifest to image destination
10:05:23: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9 is being used by the following container(s): 10032392a031d9120c8d8fbd389d6b9869a8ccc47de758d156a006164345392a: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-14 10:00:30 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.9/2110309518083624960#1:build-log.txt%3A926)
<details><summary>context</summary>
<pre>10:05:58: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
10:05:58: Writing manifest to image destination
10:05:59: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-operator-1.9 is being used by the following container(s): 835541b5d39ddc18f4520e32b95edf678f3363adc242021d3a6eb0c3f97ae156, 16572cbee2ffb45dbfd6aecd4def0d809fabc9e5ea22d60578fdf1705aa32397: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-14 10:00:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.9/2110309516342988800#1:build-log.txt%3A674)
<details><summary>context</summary>
<pre>10:05:39: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
10:05:39: Writing manifest to image destination
10:05:40: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-storage-1.9 is being used by the following container(s): b6f4c899fc3b71536a8eb1f58cdcf8734d2f9ff32ad00e021436c227acbd9cc4: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-14 10:00:15 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-windows2016-1.9/2110309509023928320#1:build-log.txt%3A963)
<details><summary>context</summary>
<pre>10:05:40: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
10:05:40: Writing manifest to image destination
10:05:41: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-windows2016-1.9 is being used by the following container(s): 08b0682d6cc334cea5ab2fac135f9f42bdbd45c871303478e6e0adab838c7462: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-14 10:00:13 &#43;0000 UTC_: <code>make: *** [Makefile:25: bazel-build] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-sig-performance-1.9/2110309508864544768#1:build-log.txt%3A372)
<details><summary>context</summary>
<pre>10:06:05: INFO: Elapsed time: 4.391s
10:06:05: INFO: 0 processes.
10:06:05: ERROR: Build failed. Not running target
make: *** [Makefile:25: bazel-build] Error 1
&#43; EXIT_VALUE=2
&#43; set &#43;o xtrace
Cleaning up after podman in container.</pre>
</details>


* _2026-10-14 10:00:13 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations-1.9/2110309513822212096#1:build-log.txt%3A971)
<details><summary>context</summary>
<pre>10:06:43: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
10:06:43: Writing manifest to image destination
10:06:43: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations-1.9 is being used by the following container(s): 783182f6a50304f6c9ec68fb4cba6bb50f9ec6a61b59c89363e0d5cac34465c1: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-14 10:00:07 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.35-sig-network-1.9/2110309518888931328#1:build-log.txt%3A665)
<details><summary>context</summary>
<pre>10:02:40: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
10:02:40: Writing manifest to image destination
10:02:41: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-network-1.9 is being used by the following container(s): b743faedf7167d571d3bf5d77c4c57c7765f947e40056206d5d2fdc6fe00984a: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-14 10:00:07 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.9/2110309517177655296#1:build-log.txt%3A1310)
<details><summary>context</summary>
<pre>10:02:55: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
10:02:55: Writing manifest to image destination
10:02:55: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-compute-1.9 is being used by the following container(s): a463d251c87d4a7b8c7bf74cbd7dbb32e58066ef48d0b27107e521c6a208f9a8: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-14 10:00:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-sig-network-1.9/2110309515508322304#1:build-log.txt%3A1100)
<details><summary>context</summary>
<pre>10:03:51: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
10:03:51: Writing manifest to image destination
10:03:51: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-network-1.9 is being used by the following container(s): 1591c1aa9ddaa77bfaa42c057996b5397024fa0503a057dc2dd140ee083ed2b8: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-10-14 10:00:01 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated-1.9/2110309513012711424#1:build-log.txt%3A627)
<details><summary>context</summary>
<pre>10:02:52: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
10:02:52: Writing manifest to image destination
10:02:52: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated-1.9 is being used by the following container(s): 486c84e92cdf60ef22e0fa6bca091eb841e225ee4c0415ea05c08826169ae769: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


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
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (5x / 13.89%) </summary>

<hr/>

**5x**: _2026-10-13 18:43:26 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19312/pull-kubevirt-e2e-k8s-1.35-sig-operator/2109978050710999040#1:build-log.txt%3A1161)
<details>
<summary>all...</summary>

* _2026-10-15 00:02:06 &#43;0000 UTC_: <code>make: *** [Makefile:194: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19454/pull-kubevirt-e2e-kind-1.37-sig-compute-arm64/2110521441312051200#1:build-log.txt%3A2485)
<details><summary>context</summary>
<pre>00:15:48: INFO: Running command line: bazel-bin/push-virt-template-controller.exe --registry localhost:5000/kubevirt --repository virt-template-controller --tag v0.3.0
00:15:55: Error during deploy: deploying images: layer sha256:7750c1a1cd8e356f717610d15552bd8a4af1801b72b5855c9e8494cc07636e77 not found in any of the original registries
00:15:57: &#43; rm -f /tmp/kubevirt.deploy.p4GY
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>


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


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### release-1.9 (29x / 80.56%)


#### external (29x / 100.00%)

<details>
<summary> download failure in context (29x / 100.00%) </summary>

<hr/>

**5x**: _2026-10-16 00:01:12 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110883575862661120#1:build-log.txt%3A619)
<details>
<summary>all...</summary>

* _2026-10-16 00:01:12 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110883575862661120#1:build-log.txt%3A619)

* _2026-10-15 20:58:13 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110837556055642112#1:build-log.txt%3A551)

* _2026-10-14 10:00:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110309514791096320#1:build-log.txt%3A523)

* _2026-10-14 09:59:58 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-kind-1.35-vgpu-1.9/2110309509128785920#1:build-log.txt%3A978)

* _2026-10-14 09:59:58 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-kind-1.34-sev-1.9/2110309509225254912#1:build-log.txt%3A1021)

</details>

<hr/>

**24x**: _2026-10-14 10:00:15 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-windows2016-1.9/2110309509023928320#1:build-log.txt%3A963)
<details>
<summary>all...</summary>

* _2026-10-16 10:10:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2111036843687415808#1:build-log.txt%3A777)

* _2026-10-16 10:10:12 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations-1.9/2111036843460923392#1:build-log.txt%3A1305)

* _2026-10-16 00:02:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.9/2110883575975907328#1:build-log.txt%3A781)

* _2026-10-15 20:59:54 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.9/2110837556139528192#1:build-log.txt%3A931)

* _2026-10-14 10:08:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations-1.9/2110309667371487232#1:build-log.txt%3A1435)

* _2026-10-14 10:08:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2110309667509899264#1:build-log.txt%3A922)

* _2026-10-14 10:08:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-operator-1.9/2110309525570457600#1:build-log.txt%3A1120)

* _2026-10-14 10:07:32 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2110309524777734144#1:build-log.txt%3A1016)

* _2026-10-14 10:07:08 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-compute-1.9/2110309523909513216#1:build-log.txt%3A1065)

* _2026-10-14 10:06:48 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-storage-1.9/2110309523062263808#1:build-log.txt%3A1299)

* _2026-10-14 10:03:08 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-network-1.9/2110309522235985920#1:build-log.txt%3A1143)

* _2026-10-14 10:02:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.35-sig-operator-1.9/2110309521405513728#1:build-log.txt%3A1373)

* _2026-10-14 10:02:11 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.9/2110309519731986432#1:build-log.txt%3A641)

* _2026-10-14 10:01:32 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.9/2110309520545681408#1:build-log.txt%3A782)

* _2026-10-14 10:00:54 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2110309512178044928#1:build-log.txt%3A785)

* _2026-10-14 10:00:30 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.9/2110309518083624960#1:build-log.txt%3A926)

* _2026-10-14 10:00:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.9/2110309516342988800#1:build-log.txt%3A674)

* _2026-10-14 10:00:15 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-windows2016-1.9/2110309509023928320#1:build-log.txt%3A963)

* _2026-10-14 10:00:13 &#43;0000 UTC_: <code>make: *** [Makefile:25: bazel-build] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-sig-performance-1.9/2110309508864544768#1:build-log.txt%3A372)

* _2026-10-14 10:00:13 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations-1.9/2110309513822212096#1:build-log.txt%3A971)

* _2026-10-14 10:00:07 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.35-sig-network-1.9/2110309518888931328#1:build-log.txt%3A665)

* _2026-10-14 10:00:07 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.9/2110309517177655296#1:build-log.txt%3A1310)

* _2026-10-14 10:00:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-sig-network-1.9/2110309515508322304#1:build-log.txt%3A1100)

* _2026-10-14 10:00:01 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated-1.9/2110309513012711424#1:build-log.txt%3A627)

</details>

<hr/>
</details>

### release-1.7 (1x / 2.78%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-10-15 22:10:37 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19476/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2110841617408593920#1:build-log.txt%3A631)
<details>
<summary>all...</summary>

* _2026-10-15 22:10:37 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19476/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2110841617408593920#1:build-log.txt%3A631)

</details>

<hr/>
</details>

### main (6x / 16.67%)


#### external (6x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (5x / 83.33%) </summary>

<hr/>

**5x**: _2026-10-13 18:43:26 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19312/pull-kubevirt-e2e-k8s-1.35-sig-operator/2109978050710999040#1:build-log.txt%3A1161)
<details>
<summary>all...</summary>

* _2026-10-15 00:02:06 &#43;0000 UTC_: <code>make: *** [Makefile:194: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19454/pull-kubevirt-e2e-kind-1.37-sig-compute-arm64/2110521441312051200#1:build-log.txt%3A2485)

* _2026-10-13 18:43:26 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19312/pull-kubevirt-e2e-k8s-1.35-sig-operator/2109978050710999040#1:build-log.txt%3A1161)

* _2026-10-13 12:08:33 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19312/pull-kubevirt-e2e-k8s-1.37-sig-network/2109978051877015552#1:build-log.txt%3A1927)

* _2026-10-12 12:01:09 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19099/pull-kubevirt-e2e-k8s-1.36-sig-network/2109609064953548800#1:build-log.txt%3A1178)

* _2026-10-11 14:33:33 &#43;0000 UTC_: <code>make: *** [Makefile:191: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19401/pull-kubevirt-e2e-kind-1.37-sev/2109291175335694336#1:build-log.txt%3A3114)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 16.67%) </summary>

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


### sig-compute (22x / 61.11%)


#### external (22x / 100.00%)

<details>
<summary> download failure in context (19x / 86.36%) </summary>

<hr/>

**5x**: _2026-10-16 00:01:12 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110883575862661120#1:build-log.txt%3A619)
<details>
<summary>all...</summary>

* _2026-10-16 00:01:12 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110883575862661120#1:build-log.txt%3A619)

* _2026-10-15 20:58:13 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110837556055642112#1:build-log.txt%3A551)

* _2026-10-14 10:00:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110309514791096320#1:build-log.txt%3A523)

* _2026-10-14 09:59:58 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-kind-1.35-vgpu-1.9/2110309509128785920#1:build-log.txt%3A978)

* _2026-10-14 09:59:58 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-kind-1.34-sev-1.9/2110309509225254912#1:build-log.txt%3A1021)

</details>

<hr/>

**14x**: _2026-10-14 10:00:15 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-windows2016-1.9/2110309509023928320#1:build-log.txt%3A963)
<details>
<summary>all...</summary>

* _2026-10-16 10:10:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2111036843687415808#1:build-log.txt%3A777)

* _2026-10-16 10:10:12 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations-1.9/2111036843460923392#1:build-log.txt%3A1305)

* _2026-10-14 10:08:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations-1.9/2110309667371487232#1:build-log.txt%3A1435)

* _2026-10-14 10:08:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2110309667509899264#1:build-log.txt%3A922)

* _2026-10-14 10:08:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-operator-1.9/2110309525570457600#1:build-log.txt%3A1120)

* _2026-10-14 10:07:32 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2110309524777734144#1:build-log.txt%3A1016)

* _2026-10-14 10:07:08 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-compute-1.9/2110309523909513216#1:build-log.txt%3A1065)

* _2026-10-14 10:02:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.35-sig-operator-1.9/2110309521405513728#1:build-log.txt%3A1373)

* _2026-10-14 10:01:32 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.9/2110309520545681408#1:build-log.txt%3A782)

* _2026-10-14 10:00:30 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.9/2110309518083624960#1:build-log.txt%3A926)

* _2026-10-14 10:00:15 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-windows2016-1.9/2110309509023928320#1:build-log.txt%3A963)

* _2026-10-14 10:00:13 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations-1.9/2110309513822212096#1:build-log.txt%3A971)

* _2026-10-14 10:00:07 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.9/2110309517177655296#1:build-log.txt%3A1310)

* _2026-10-12 12:33:22 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18853/pull-kubevirt-e2e-k8s-1.37-sig-operator/2109622445727027200#1:build-log.txt%3A380)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (3x / 13.64%) </summary>

<hr/>

**3x**: _2026-10-13 18:43:26 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19312/pull-kubevirt-e2e-k8s-1.35-sig-operator/2109978050710999040#1:build-log.txt%3A1161)
<details>
<summary>all...</summary>

* _2026-10-15 00:02:06 &#43;0000 UTC_: <code>make: *** [Makefile:194: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19454/pull-kubevirt-e2e-kind-1.37-sig-compute-arm64/2110521441312051200#1:build-log.txt%3A2485)

* _2026-10-13 18:43:26 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19312/pull-kubevirt-e2e-k8s-1.35-sig-operator/2109978050710999040#1:build-log.txt%3A1161)

* _2026-10-11 14:33:33 &#43;0000 UTC_: <code>make: *** [Makefile:191: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19401/pull-kubevirt-e2e-kind-1.37-sev/2109291175335694336#1:build-log.txt%3A3114)

</details>

<hr/>
</details>

### sig-storage (6x / 16.67%)


#### external (6x / 100.00%)

<details>
<summary> download failure in context (6x / 100.00%) </summary>

<hr/>

**6x**: _2026-10-15 22:10:37 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19476/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2110841617408593920#1:build-log.txt%3A631)
<details>
<summary>all...</summary>

* _2026-10-16 00:02:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.9/2110883575975907328#1:build-log.txt%3A781)

* _2026-10-15 22:10:37 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19476/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2110841617408593920#1:build-log.txt%3A631)

* _2026-10-15 20:59:54 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.9/2110837556139528192#1:build-log.txt%3A931)

* _2026-10-14 10:06:48 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-storage-1.9/2110309523062263808#1:build-log.txt%3A1299)

* _2026-10-14 10:02:11 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.9/2110309519731986432#1:build-log.txt%3A641)

* _2026-10-14 10:00:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.9/2110309516342988800#1:build-log.txt%3A674)

</details>

<hr/>
</details>

### sig-performance (1x / 2.78%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-10-14 10:00:13 &#43;0000 UTC_: <code>make: *** [Makefile:25: bazel-build] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-sig-performance-1.9/2110309508864544768#1:build-log.txt%3A372)
<details>
<summary>all...</summary>

* _2026-10-14 10:00:13 &#43;0000 UTC_: <code>make: *** [Makefile:25: bazel-build] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-sig-performance-1.9/2110309508864544768#1:build-log.txt%3A372)

</details>

<hr/>
</details>

### sig-network (7x / 19.44%)


#### external (7x / 100.00%)

<details>
<summary> download failure in context (5x / 71.43%) </summary>

<hr/>

**5x**: _2026-10-14 10:00:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-sig-network-1.9/2110309515508322304#1:build-log.txt%3A1100)
<details>
<summary>all...</summary>

* _2026-10-14 10:03:08 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-network-1.9/2110309522235985920#1:build-log.txt%3A1143)

* _2026-10-14 10:00:54 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2110309512178044928#1:build-log.txt%3A785)

* _2026-10-14 10:00:07 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.35-sig-network-1.9/2110309518888931328#1:build-log.txt%3A665)

* _2026-10-14 10:00:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.34-sig-network-1.9/2110309515508322304#1:build-log.txt%3A1100)

* _2026-10-14 10:00:01 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19375/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated-1.9/2110309513012711424#1:build-log.txt%3A627)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 28.57%) </summary>

<hr/>

**2x**: _2026-10-12 12:01:09 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19099/pull-kubevirt-e2e-k8s-1.36-sig-network/2109609064953548800#1:build-log.txt%3A1178)
<details>
<summary>all...</summary>

* _2026-10-13 12:08:33 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19312/pull-kubevirt-e2e-k8s-1.37-sig-network/2109978051877015552#1:build-log.txt%3A1927)

* _2026-10-12 12:01:09 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19099/pull-kubevirt-e2e-k8s-1.36-sig-network/2109609064953548800#1:build-log.txt%3A1178)

</details>

<hr/>
</details>

Last updated: 2026-10-18 03:16:07
