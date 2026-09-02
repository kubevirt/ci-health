



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-09-01 (2x / 40.00%)


#### external (2x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-09-01 19:59:10 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18978/pull-kubevirt-e2e-k8s-1.37-sig-network/2094853882197839872#1:build-log.txt%3A1859)
<details>
<summary>all...</summary>

* _2026-09-01 19:59:10 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18978/pull-kubevirt-e2e-k8s-1.37-sig-network/2094853882197839872#1:build-log.txt%3A1859)

* _2026-09-01 13:58:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18986/pull-kubevirt-e2e-k8s-1.36-sig-network/2094785540393013248#1:build-log.txt%3A1946)

</details>

<hr/>
</details>

### 2026-08-31 (3x / 60.00%)


#### external (1x / 33.33%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-08-31 08:52:40 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18981/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2094347452374061056#1:build-log.txt%3A1201)
<details>
<summary>all...</summary>

* _2026-08-31 08:52:40 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18981/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2094347452374061056#1:build-log.txt%3A1201)

</details>

<hr/>
</details>

#### internal (2x / 66.67%)

<details>
<summary> make cluster lifecycle target failure (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-08-31 07:37:21 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18969/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2094328540248936448#1:build-log.txt%3A1186)
<details>
<summary>all...</summary>

* _2026-08-31 07:37:26 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18969/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations-1.7/2094328530190995456#1:build-log.txt%3A1350)

* _2026-08-31 07:37:21 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18969/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2094328540248936448#1:build-log.txt%3A1186)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (3x / 60.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (3x / 60.00%) </summary>

<hr/>

**3x**: _2026-09-01 19:59:10 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18978/pull-kubevirt-e2e-k8s-1.37-sig-network/2094853882197839872#1:build-log.txt%3A1859)
<details>
<summary>all...</summary>

* _2026-09-01 19:59:10 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18978/pull-kubevirt-e2e-k8s-1.37-sig-network/2094853882197839872#1:build-log.txt%3A1859)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
20:06:26: selecting podman as container runtime
20:07:03: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.37-sig-network is being used by the following container(s): 33aa40444a95e70e203c32469ac053454d8ea169f0722c201e0e2290b4122168: volume is being used
make: *** [Makefile:180: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-09-01 13:58:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18986/pull-kubevirt-e2e-k8s-1.36-sig-network/2094785540393013248#1:build-log.txt%3A1946)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
14:09:48: selecting podman as container runtime
14:10:21: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network is being used by the following container(s): 07060073a1c37b1bf8a2b94ff00688edee5137962acbb91a9fd8fbff44f321f5: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-08-31 08:52:40 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18981/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2094347452374061056#1:build-log.txt%3A1201)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
09:27:37: selecting podman as container runtime
09:28:10: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7 is being used by the following container(s): 10bcf8f86081cb1356856dbe32c4e0de9ac069d27daebe1f4c19f4d1c8eff037, fe1e9d6d079ad0daabd792e14946b531233d88f3848139e2104f65fb2508cbf8: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>

### internal (2x / 40.00%)

<details>
<summary> make cluster lifecycle target failure (2x / 40.00%) </summary>

<hr/>

**2x**: _2026-08-31 07:37:21 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18969/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2094328540248936448#1:build-log.txt%3A1186)
<details>
<summary>all...</summary>

* _2026-08-31 07:37:26 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18969/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations-1.7/2094328530190995456#1:build-log.txt%3A1350)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
07:59:51: selecting podman as container runtime
08:00:58: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations-1.7 is being used by the following container(s): 44af09be771b50a17a5a5812cc1d3cfdc134139d1598ac1020d06476bd8f5c69, cf133e6fe06ee3e441f1532c21b517dfbcdb493092d9c83edd80c8376d11a9ea: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-08-31 07:37:21 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18969/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2094328540248936448#1:build-log.txt%3A1186)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
08:08:30: selecting podman as container runtime
08:09:13: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7 is being used by the following container(s): 431727bab8076d35537eb9e7ea5f1a374c8d6776d2fc58016b9690a6cad47901, 9e4d4ff6924e1c3e6edc1a5af6794c1531c70ee08ecc50e6071519d5c8588c54: volume is being used
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


### main (2x / 40.00%)


#### external (2x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-09-01 19:59:10 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18978/pull-kubevirt-e2e-k8s-1.37-sig-network/2094853882197839872#1:build-log.txt%3A1859)
<details>
<summary>all...</summary>

* _2026-09-01 19:59:10 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18978/pull-kubevirt-e2e-k8s-1.37-sig-network/2094853882197839872#1:build-log.txt%3A1859)

* _2026-09-01 13:58:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18986/pull-kubevirt-e2e-k8s-1.36-sig-network/2094785540393013248#1:build-log.txt%3A1946)

</details>

<hr/>
</details>

### release-1.7 (3x / 60.00%)


#### external (1x / 33.33%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-08-31 08:52:40 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18981/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2094347452374061056#1:build-log.txt%3A1201)
<details>
<summary>all...</summary>

* _2026-08-31 08:52:40 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18981/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2094347452374061056#1:build-log.txt%3A1201)

</details>

<hr/>
</details>

#### internal (2x / 66.67%)

<details>
<summary> make cluster lifecycle target failure (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-08-31 07:37:21 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18969/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2094328540248936448#1:build-log.txt%3A1186)
<details>
<summary>all...</summary>

* _2026-08-31 07:37:26 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18969/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations-1.7/2094328530190995456#1:build-log.txt%3A1350)

* _2026-08-31 07:37:21 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18969/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2094328540248936448#1:build-log.txt%3A1186)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-network (2x / 40.00%)


#### external (2x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-09-01 19:59:10 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18978/pull-kubevirt-e2e-k8s-1.37-sig-network/2094853882197839872#1:build-log.txt%3A1859)
<details>
<summary>all...</summary>

* _2026-09-01 19:59:10 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18978/pull-kubevirt-e2e-k8s-1.37-sig-network/2094853882197839872#1:build-log.txt%3A1859)

* _2026-09-01 13:58:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18986/pull-kubevirt-e2e-k8s-1.36-sig-network/2094785540393013248#1:build-log.txt%3A1946)

</details>

<hr/>
</details>

### sig-storage (2x / 40.00%)


#### external (1x / 50.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-08-31 08:52:40 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18981/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2094347452374061056#1:build-log.txt%3A1201)
<details>
<summary>all...</summary>

* _2026-08-31 08:52:40 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18981/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2094347452374061056#1:build-log.txt%3A1201)

</details>

<hr/>
</details>

#### internal (1x / 50.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-08-31 07:37:21 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18969/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2094328540248936448#1:build-log.txt%3A1186)
<details>
<summary>all...</summary>

* _2026-08-31 07:37:21 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18969/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2094328540248936448#1:build-log.txt%3A1186)

</details>

<hr/>
</details>

### sig-compute (1x / 20.00%)


#### internal (1x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-08-31 07:37:26 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18969/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations-1.7/2094328530190995456#1:build-log.txt%3A1350)
<details>
<summary>all...</summary>

* _2026-08-31 07:37:26 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18969/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations-1.7/2094328530190995456#1:build-log.txt%3A1350)

</details>

<hr/>
</details>

Last updated: 2026-09-02 09:12:43
