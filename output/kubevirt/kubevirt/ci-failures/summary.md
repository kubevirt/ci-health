



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-10-16 (4x / 26.67%)


#### external (4x / 100.00%)

<details>
<summary> download failure in context (4x / 100.00%) </summary>

<hr/>

**3x**: _2026-10-16 10:10:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2111036843687415808#1:build-log.txt%3A777)
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

### 2026-10-15 (4x / 26.67%)


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

### 2026-10-14 (2x / 13.33%)


#### external (2x / 100.00%)

<details>
<summary> download failure in context (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-10-14 10:08:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2110309667509899264#1:build-log.txt%3A922)
<details>
<summary>all...</summary>

* _2026-10-14 10:08:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations-1.9/2110309667371487232#1:build-log.txt%3A1435)

* _2026-10-14 10:08:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2110309667509899264#1:build-log.txt%3A922)

</details>

<hr/>
</details>

### 2026-10-13 (2x / 13.33%)


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

### 2026-10-12 (2x / 13.33%)


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

### 2026-10-11 (1x / 6.67%)


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


### external (15x / 100.00%)

<details>
<summary> download failure in context (10x / 66.67%) </summary>

<hr/>

**8x**: _2026-10-16 10:10:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2111036843687415808#1:build-log.txt%3A777)
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

**2x**: _2026-10-16 00:01:12 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110883575862661120#1:build-log.txt%3A619)
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


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (5x / 33.33%) </summary>

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


### release-1.9 (8x / 53.33%)


#### external (8x / 100.00%)

<details>
<summary> download failure in context (8x / 100.00%) </summary>

<hr/>

**6x**: _2026-10-16 10:10:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2111036843687415808#1:build-log.txt%3A777)
<details>
<summary>all...</summary>

* _2026-10-16 10:10:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2111036843687415808#1:build-log.txt%3A777)

* _2026-10-16 10:10:12 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations-1.9/2111036843460923392#1:build-log.txt%3A1305)

* _2026-10-16 00:02:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.9/2110883575975907328#1:build-log.txt%3A781)

* _2026-10-15 20:59:54 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.9/2110837556139528192#1:build-log.txt%3A931)

* _2026-10-14 10:08:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations-1.9/2110309667371487232#1:build-log.txt%3A1435)

* _2026-10-14 10:08:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2110309667509899264#1:build-log.txt%3A922)

</details>

<hr/>

**2x**: _2026-10-16 00:01:12 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110883575862661120#1:build-log.txt%3A619)
<details>
<summary>all...</summary>

* _2026-10-16 00:01:12 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110883575862661120#1:build-log.txt%3A619)

* _2026-10-15 20:58:13 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110837556055642112#1:build-log.txt%3A551)

</details>

<hr/>
</details>

### release-1.7 (1x / 6.67%)


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

### main (6x / 40.00%)


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


### sig-compute (10x / 66.67%)


#### external (10x / 100.00%)

<details>
<summary> download failure in context (7x / 70.00%) </summary>

<hr/>

**5x**: _2026-10-16 10:10:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2111036843687415808#1:build-log.txt%3A777)
<details>
<summary>all...</summary>

* _2026-10-16 10:10:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2111036843687415808#1:build-log.txt%3A777)

* _2026-10-16 10:10:12 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations-1.9/2111036843460923392#1:build-log.txt%3A1305)

* _2026-10-14 10:08:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations-1.9/2110309667371487232#1:build-log.txt%3A1435)

* _2026-10-14 10:08:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19433/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2110309667509899264#1:build-log.txt%3A922)

* _2026-10-12 12:33:22 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18853/pull-kubevirt-e2e-k8s-1.37-sig-operator/2109622445727027200#1:build-log.txt%3A380)

</details>

<hr/>

**2x**: _2026-10-16 00:01:12 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110883575862661120#1:build-log.txt%3A619)
<details>
<summary>all...</summary>

* _2026-10-16 00:01:12 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110883575862661120#1:build-log.txt%3A619)

* _2026-10-15 20:58:13 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64-1.9/2110837556055642112#1:build-log.txt%3A551)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (3x / 30.00%) </summary>

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

### sig-storage (3x / 20.00%)


#### external (3x / 100.00%)

<details>
<summary> download failure in context (3x / 100.00%) </summary>

<hr/>

**3x**: _2026-10-15 22:10:37 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19476/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2110841617408593920#1:build-log.txt%3A631)
<details>
<summary>all...</summary>

* _2026-10-16 00:02:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.9/2110883575975907328#1:build-log.txt%3A781)

* _2026-10-15 22:10:37 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19476/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2110841617408593920#1:build-log.txt%3A631)

* _2026-10-15 20:59:54 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19352/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.9/2110837556139528192#1:build-log.txt%3A931)

</details>

<hr/>
</details>

### sig-network (2x / 13.33%)


#### external (2x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-10-12 12:01:09 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19099/pull-kubevirt-e2e-k8s-1.36-sig-network/2109609064953548800#1:build-log.txt%3A1178)
<details>
<summary>all...</summary>

* _2026-10-13 12:08:33 &#43;0000 UTC_: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19312/pull-kubevirt-e2e-k8s-1.37-sig-network/2109978051877015552#1:build-log.txt%3A1927)

* _2026-10-12 12:01:09 &#43;0000 UTC_: <code>make: *** [Makefile:179: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19099/pull-kubevirt-e2e-k8s-1.36-sig-network/2109609064953548800#1:build-log.txt%3A1178)

</details>

<hr/>
</details>

Last updated: 2026-10-17 09:14:01
