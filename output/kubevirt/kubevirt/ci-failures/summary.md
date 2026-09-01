



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-08-31 (2x / 100.00%)


#### internal (2x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-08-31 07:37:26 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18969/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations-1.7/2094328530190995456#1:build-log.txt%3A1350)
<details>
<summary>all...</summary>

* _2026-08-31 07:37:26 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18969/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations-1.7/2094328530190995456#1:build-log.txt%3A1350)

* _2026-08-31 07:37:21 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18969/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2094328540248936448#1:build-log.txt%3A1186)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### internal (2x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-08-31 07:37:26 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18969/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations-1.7/2094328530190995456#1:build-log.txt%3A1350)
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


### release-1.7 (2x / 100.00%)


#### internal (2x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-08-31 07:37:26 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18969/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations-1.7/2094328530190995456#1:build-log.txt%3A1350)
<details>
<summary>all...</summary>

* _2026-08-31 07:37:26 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18969/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations-1.7/2094328530190995456#1:build-log.txt%3A1350)

* _2026-08-31 07:37:21 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18969/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2094328540248936448#1:build-log.txt%3A1186)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (1x / 50.00%)


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

### sig-storage (1x / 50.00%)


#### internal (1x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-08-31 07:37:21 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18969/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2094328540248936448#1:build-log.txt%3A1186)
<details>
<summary>all...</summary>

* _2026-08-31 07:37:21 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18969/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2094328540248936448#1:build-log.txt%3A1186)

</details>

<hr/>
</details>

Last updated: 2026-09-01 09:10:59
