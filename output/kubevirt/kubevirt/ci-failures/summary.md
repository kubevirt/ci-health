



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-07-26 (9x / 6.82%)


#### external (7x / 77.78%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (6x / 66.67%) </summary>

<hr/>

**6x**: _2026-07-26 06:44:01 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18569/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2081269171055759360#1:build-log.txt%3A4441)
<details>
<summary>all...</summary>

* _2026-07-26 19:16:47 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18469/pull-kubevirt-e2e-kind-1.36-sev/2081458617176821760#1:build-log.txt%3A1171)

* _2026-07-26 17:10:22 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18469/pull-kubevirt-e2e-k8s-1.34-sig-compute/2081426767049920512#1:build-log.txt%3A2934)

* _2026-07-26 10:45:42 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18474/pull-kubevirt-e2e-k8s-1.35-sig-network/2081329966217170944#1:build-log.txt%3A4396)

* _2026-07-26 10:45:39 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18474/pull-kubevirt-e2e-k8s-1.34-sig-storage/2081329965713854464#1:build-log.txt%3A2866)

* _2026-07-26 06:44:05 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18569/pull-kubevirt-e2e-k8s-1.34-sig-compute/2081269171529715712#1:build-log.txt%3A2957)

* _2026-07-26 06:44:01 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18569/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2081269171055759360#1:build-log.txt%3A4441)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-07-26 17:10:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18469/pull-kubevirt-e2e-k8s-1.35-sig-network/2081426767381270528#1:build-log.txt%3A2180)
<details>
<summary>all...</summary>

* _2026-07-26 17:10:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18469/pull-kubevirt-e2e-k8s-1.35-sig-network/2081426767381270528#1:build-log.txt%3A2180)

</details>

<hr/>
</details>

#### internal (2x / 22.22%)

<details>
<summary> make cluster lifecycle target failure (2x / 22.22%) </summary>

<hr/>

**2x**: _2026-07-26 10:45:41 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18474/pull-kubevirt-e2e-k8s-1.34-windows2016/2081329964916936704#1:build-log.txt%3A2120)
<details>
<summary>all...</summary>

* _2026-07-26 17:10:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18469/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2081426766697598976#1:build-log.txt%3A1131)

* _2026-07-26 10:45:41 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18474/pull-kubevirt-e2e-k8s-1.34-windows2016/2081329964916936704#1:build-log.txt%3A2120)

</details>

<hr/>
</details>

### 2026-07-25 (29x / 21.97%)


#### external (27x / 93.10%)

<details>
<summary> bazel remote cache blob fetch failure (from secondary snippet) (22x / 75.86%) </summary>

<hr/>

**22x**: _2026-07-25 07:26:15 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-compute/2080916717827002368#1:build-log.txt%3A434)
<details>
<summary>all...</summary>

* _2026-07-25 07:51:26 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-windows2016/2080923721920417792#1:build-log.txt%3A415)

* _2026-07-25 07:28:43 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080916717994774528#1:build-log.txt%3A384)

* _2026-07-25 07:26:15 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-compute/2080916717827002368#1:build-log.txt%3A434)

* _2026-07-25 07:23:37 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-storage/2080916717680201728#1:build-log.txt%3A438)

* _2026-07-25 07:23:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-operator/2080916717386600448#1:build-log.txt%3A449)

* _2026-07-25 07:23:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080916717092999168#1:build-log.txt%3A449)

* _2026-07-25 07:23:35 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080916716325441536#1:build-log.txt%3A460)

* _2026-07-25 07:23:34 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-network/2080916716866506752#1:build-log.txt%3A367)

* _2026-07-25 07:23:34 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080916716623237120#1:build-log.txt%3A457)

* _2026-07-25 07:23:34 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-network/2080916717504040960#1:build-log.txt%3A432)

* _2026-07-25 07:23:30 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080916717239799808#1:build-log.txt%3A433)

* _2026-07-25 03:48:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-windows2016/2080862637561745408#1:build-log.txt%3A447)

* _2026-07-25 02:20:43 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-network/2080840488520257536#1:build-log.txt%3A444)

* _2026-07-25 02:20:42 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080840488079855616#1:build-log.txt%3A438)

* _2026-07-25 02:20:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-network/2080840487949832192#1:build-log.txt%3A425)

* _2026-07-25 02:20:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080840487677202432#1:build-log.txt%3A429)

* _2026-07-25 02:20:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-storage/2080840488658669568#1:build-log.txt%3A388)

* _2026-07-25 02:20:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-compute/2080840488834830336#1:build-log.txt%3A402)

* _2026-07-25 02:20:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-operator/2080840488360873984#1:build-log.txt%3A389)

* _2026-07-25 02:20:40 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080840488197296128#1:build-log.txt%3A372)

* _2026-07-25 02:20:40 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080840487815614464#1:build-log.txt%3A364)

* _2026-07-25 02:20:39 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080840488973242368#1:build-log.txt%3A392)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (5x / 17.24%) </summary>

<hr/>

**5x**: _2026-07-25 00:25:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18533/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080805657153376256#1:build-log.txt%3A2974)
<details>
<summary>all...</summary>

* _2026-07-25 07:23:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18585/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080916635056607232#1:build-log.txt%3A5259)

* _2026-07-25 07:23:10 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18585/pull-kubevirt-e2e-k8s-1.36-sig-network/2080916624981889024#1:build-log.txt%3A3651)

* _2026-07-25 00:35:31 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18533/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080805660005502976#1:build-log.txt%3A4006)

* _2026-07-25 00:28:33 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18533/pull-kubevirt-e2e-k8s-1.35-sig-network/2080805657438588928#1:build-log.txt%3A4385)

* _2026-07-25 00:25:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18533/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080805657153376256#1:build-log.txt%3A2974)

</details>

<hr/>
</details>

#### internal (2x / 6.90%)

<details>
<summary> make cluster lifecycle target failure (2x / 6.90%) </summary>

<hr/>

**2x**: _2026-07-25 00:13:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18533/pull-kubevirt-e2e-k8s-1.36-sig-network/2080805647405813760#1:build-log.txt%3A2212)
<details>
<summary>all...</summary>

* _2026-07-25 07:23:09 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18585/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080916614278025216#1:build-log.txt%3A789)

* _2026-07-25 00:13:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18533/pull-kubevirt-e2e-k8s-1.36-sig-network/2080805647405813760#1:build-log.txt%3A2212)

</details>

<hr/>
</details>

### 2026-07-24 (38x / 28.79%)


#### external (34x / 89.47%)

<details>
<summary> bazel remote cache blob fetch failure (from secondary snippet) (22x / 57.89%) </summary>

<hr/>

**22x**: _2026-07-24 22:45:44 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-compute/2080764930268598272#1:build-log.txt%3A388)
<details>
<summary>all...</summary>

* _2026-07-24 23:32:17 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-windows2016/2080797902535397376#1:build-log.txt%3A476)

* _2026-07-24 22:46:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080764931073904640#1:build-log.txt%3A373)

* _2026-07-24 22:45:46 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-storage/2080764930146963456#1:build-log.txt%3A406)

* _2026-07-24 22:45:44 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-compute/2080764930268598272#1:build-log.txt%3A388)

* _2026-07-24 22:44:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-network/2080764930012745728#1:build-log.txt%3A406)

* _2026-07-24 22:43:29 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-operator/2080764929849167872#1:build-log.txt%3A375)

* _2026-07-24 22:42:00 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080764929639452672#1:build-log.txt%3A386)

* _2026-07-24 22:41:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080764929391988736#1:build-log.txt%3A394)

* _2026-07-24 22:39:57 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-network/2080764929308102656#1:build-log.txt%3A405)

* _2026-07-24 22:38:34 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080764929115164672#1:build-log.txt%3A399)

* _2026-07-24 22:35:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080764928976752640#1:build-log.txt%3A382)

* _2026-07-24 18:13:39 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080711464485654528#1:build-log.txt%3A455)

* _2026-07-24 18:12:50 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-compute/2080711464292716544#1:build-log.txt%3A434)

* _2026-07-24 17:59:31 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-storage/2080711464133332992#1:build-log.txt%3A375)

* _2026-07-24 17:59:18 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-network/2080711463994920960#1:build-log.txt%3A397)

* _2026-07-24 17:59:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-operator/2080711463919423488#1:build-log.txt%3A453)

* _2026-07-24 17:56:22 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080711463818760192#1:build-log.txt%3A385)

* _2026-07-24 17:56:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080711463592267776#1:build-log.txt%3A441)

* _2026-07-24 17:55:57 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-network/2080711463470632960#1:build-log.txt%3A451)

* _2026-07-24 17:52:43 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080711463294472192#1:build-log.txt%3A397)

* _2026-07-24 17:52:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080711463156060160#1:build-log.txt%3A447)

* _2026-07-24 17:48:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-windows2016/2080711462820515840#1:build-log.txt%3A395)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (11x / 28.95%) </summary>

<hr/>

**11x**: _2026-07-24 13:16:21 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18443/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080642898696605696#1:build-log.txt%3A3556)
<details>
<summary>all...</summary>

* _2026-07-24 16:23:44 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18428/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080570620118044672#1:build-log.txt%3A3565)

* _2026-07-24 14:17:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18400/pull-kubevirt-e2e-k8s-1.35-sig-storage/2080658114465501184#1:build-log.txt%3A2570)

* _2026-07-24 14:16:53 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18168/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080651063739813888#1:build-log.txt%3A2651)

* _2026-07-24 14:15:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18168/pull-kubevirt-e2e-k8s-1.34-sig-network/2080651063618179072#1:build-log.txt%3A4014)

* _2026-07-24 13:21:54 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18443/pull-kubevirt-e2e-k8s-1.35-sig-network/2080642899699044352#1:build-log.txt%3A5078)

* _2026-07-24 13:16:21 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18443/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080642898696605696#1:build-log.txt%3A3556)

* _2026-07-24 13:11:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18428/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080570620541669376#1:build-log.txt%3A3677)

* _2026-07-24 13:09:32 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18428/pull-kubevirt-e2e-k8s-1.34-sig-network/2080570620331954176#1:build-log.txt%3A4014)

* _2026-07-24 07:51:34 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18542/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080561357886853120#1:build-log.txt%3A2659)

* _2026-07-24 07:51:32 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18542/pull-kubevirt-e2e-k8s-1.36-sig-network/2080561377302286336#1:build-log.txt%3A5074)

* _2026-07-24 07:51:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18542/pull-kubevirt-e2e-k8s-1.34-sig-operator/2080561358025265152#1:build-log.txt%3A2572)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 2.63%) </summary>

<hr/>

**1x**: _2026-07-24 16:07:20 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18400/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080685210604670976#1:build-log.txt%3A1193)
<details>
<summary>all...</summary>

* _2026-07-24 16:07:20 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18400/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080685210604670976#1:build-log.txt%3A1193)

</details>

<hr/>
</details>

#### internal (4x / 10.53%)

<details>
<summary> make cluster lifecycle target failure (4x / 10.53%) </summary>

<hr/>

**4x**: _2026-07-24 13:09:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18428/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080570620432617472#1:build-log.txt%3A1966)
<details>
<summary>all...</summary>

* _2026-07-24 13:16:39 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18443/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080642898814046208#1:build-log.txt%3A1929)

* _2026-07-24 13:09:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18428/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080570620432617472#1:build-log.txt%3A1966)

* _2026-07-24 07:51:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18542/pull-kubevirt-e2e-k8s-1.35-sig-network/2080561358188843008#1:build-log.txt%3A841)

* _2026-07-24 07:51:32 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18542/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080561368284532736#1:build-log.txt%3A1855)

</details>

<hr/>
</details>

### 2026-07-23 (13x / 9.85%)


#### external (8x / 61.54%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (8x / 61.54%) </summary>

<hr/>

**8x**: _2026-07-23 06:25:58 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18568/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2080177439815766016#1:build-log.txt%3A4360)
<details>
<summary>all...</summary>

* _2026-07-23 18:32:12 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18400/pull-kubevirt-e2e-k8s-1.36-sig-network/2080359448429203456#1:build-log.txt%3A5050)

* _2026-07-23 15:40:52 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18428/pull-kubevirt-e2e-k8s-1.36-sig-storage/2080316875354411008#1:build-log.txt%3A3593)

* _2026-07-23 14:26:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18296/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080297545166229504#1:build-log.txt%3A4045)

* _2026-07-23 13:25:38 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18296/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080267705973215232#1:build-log.txt%3A2570)

* _2026-07-23 12:38:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18491/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2080239156256051200#1:build-log.txt%3A2662)

* _2026-07-23 11:51:19 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18410/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080253331493621760#1:build-log.txt%3A3558)

* _2026-07-23 06:25:58 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18568/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2080177439815766016#1:build-log.txt%3A4360)

* _2026-07-23 03:35:15 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18453/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080107377016705024#1:build-log.txt%3A5485)

</details>

<hr/>
</details>

#### internal (5x / 38.46%)

<details>
<summary> make cluster lifecycle target failure (5x / 38.46%) </summary>

<hr/>

**5x**: _2026-07-23 10:35:26 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18296/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080239905669124096#1:build-log.txt%3A836)
<details>
<summary>all...</summary>

* _2026-07-23 21:52:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18400/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080410663418597376#1:build-log.txt%3A1825)

* _2026-07-23 16:11:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18568/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080233043007639552#1:build-log.txt%3A1943)

* _2026-07-23 11:52:45 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18414/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080250895450247168#1:build-log.txt%3A832)

* _2026-07-23 10:43:51 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18296/pull-kubevirt-e2e-k8s-1.34-sig-network/2080239905874644992#1:build-log.txt%3A834)

* _2026-07-23 10:35:26 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18296/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080239905669124096#1:build-log.txt%3A836)

</details>

<hr/>
</details>

### 2026-07-22 (16x / 12.12%)


#### internal (10x / 62.50%)

<details>
<summary> make cluster lifecycle target failure (10x / 62.50%) </summary>

<hr/>

**10x**: _2026-07-22 19:08:17 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17946/pull-kubevirt-e2e-k8s-1.35-sig-compute/2079925529972051968#1:build-log.txt%3A836)
<details>
<summary>all...</summary>

* _2026-07-22 19:08:17 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17946/pull-kubevirt-e2e-k8s-1.35-sig-compute/2079925529972051968#1:build-log.txt%3A836)

* _2026-07-22 19:07:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17946/pull-kubevirt-e2e-k8s-1.35-sig-storage/2079925529808474112#1:build-log.txt%3A835)

* _2026-07-22 16:07:50 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18530/pull-kubevirt-e2e-k8s-1.34-windows2016/2079961254075568128#1:build-log.txt%3A832)

* _2026-07-22 15:51:21 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18530/pull-kubevirt-e2e-k8s-1.34-sig-storage/2079956261884923904#1:build-log.txt%3A1902)

* _2026-07-22 13:28:23 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18565/pull-kubevirt-e2e-kind-1.34-sev-1.9/2079921377246187520#1:build-log.txt%3A812)

* _2026-07-22 12:33:47 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18530/pull-kubevirt-e2e-kind-1.36-sev/2079907633388589056#1:build-log.txt%3A648)

* _2026-07-22 12:27:27 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18562/pull-kubevirt-e2e-kind-1.34-sev-1.7/2079905794379223040#1:build-log.txt%3A655)

* _2026-07-22 11:52:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18453/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2079896117457195008#1:build-log.txt%3A866)

* _2026-07-22 11:24:21 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18550/pull-kubevirt-e2e-k8s-1.36-sig-compute/2079884525898305536#1:build-log.txt%3A2062)

* _2026-07-22 10:36:23 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18436/pull-kubevirt-e2e-k8s-1.36-sig-network/2079877999724859392#1:build-log.txt%3A1911)

</details>

<hr/>
</details>

#### external (5x / 31.25%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (5x / 31.25%) </summary>

<hr/>

**5x**: _2026-07-22 18:11:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18550/pull-kubevirt-e2e-k8s-1.35-sig-compute/2079887455724507136#1:build-log.txt%3A2629)
<details>
<summary>all...</summary>

* _2026-07-22 19:01:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17946/pull-kubevirt-e2e-k8s-1.34-sig-compute/2079925529477124096#1:build-log.txt%3A3703)

* _2026-07-22 18:11:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18550/pull-kubevirt-e2e-k8s-1.35-sig-compute/2079887455724507136#1:build-log.txt%3A2629)

* _2026-07-22 15:50:33 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18530/pull-kubevirt-e2e-k8s-1.34-sig-network/2079956261721346048#1:build-log.txt%3A3523)

* _2026-07-22 12:18:49 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18552/pull-kubevirt-e2e-kind-1.34-sev-1.9/2079903851376283648#1:build-log.txt%3A5353)

* _2026-07-22 12:05:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18550/pull-kubevirt-e2e-k8s-1.34-sig-network/2079887454780788736#1:build-log.txt%3A5057)

</details>

<hr/>
</details>

#### needs-investigation (1x / 6.25%)

<details>
<summary> no error snippets (1x / 6.25%) </summary>

<hr/>

**1x**: _2026-07-22 13:32:04 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18168/pull-kubevirt-e2e-kind-1.36-sev/2079922111794647040)
<details>
<summary>all...</summary>

* _2026-07-22 13:32:04 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18168/pull-kubevirt-e2e-kind-1.36-sev/2079922111794647040)

</details>

<hr/>
</details>

### 2026-07-21 (20x / 15.15%)


#### internal (7x / 35.00%)

<details>
<summary> make cluster lifecycle target failure (7x / 35.00%) </summary>

<hr/>

**7x**: _2026-07-21 12:59:49 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18538/pull-kubevirt-e2e-k8s-1.34-sig-network/2079550395427852288#1:build-log.txt%3A835)
<details>
<summary>all...</summary>

* _2026-07-21 17:45:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18436/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2079623486266937344#1:build-log.txt%3A1895)

* _2026-07-21 15:53:11 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.34-sig-operator/2079593296690155520#1:build-log.txt%3A834)

* _2026-07-21 14:26:15 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.34-sig-operator/2079548939396190208#1:build-log.txt%3A835)

* _2026-07-21 12:59:49 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18538/pull-kubevirt-e2e-k8s-1.34-sig-network/2079550395427852288#1:build-log.txt%3A835)

* _2026-07-21 11:37:21 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18402/pull-kubevirt-e2e-k8s-1.36-sig-network/2079528728521609216#1:build-log.txt%3A1836)

* _2026-07-21 07:32:57 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18499/pull-kubevirt-e2e-k8s-1.34-sig-network/2079469517045501952#1:build-log.txt%3A839)

* _2026-07-21 07:32:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18499/pull-kubevirt-e2e-k8s-1.35-sig-network/2079469518014386176#1:build-log.txt%3A841)

</details>

<hr/>
</details>

#### external (13x / 65.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (13x / 65.00%) </summary>

<hr/>

**13x**: _2026-07-21 15:43:51 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18301/pull-kubevirt-e2e-k8s-1.35-sig-compute/2079588132977643520#1:build-log.txt%3A2628)
<details>
<summary>all...</summary>

* _2026-07-21 17:49:35 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18436/pull-kubevirt-e2e-k8s-1.35-sig-network/2079623487349067776#1:build-log.txt%3A2546)

* _2026-07-21 15:43:51 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18301/pull-kubevirt-e2e-k8s-1.35-sig-compute/2079588132977643520#1:build-log.txt%3A2628)

* _2026-07-21 15:28:48 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18301/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2079588132075868160#1:build-log.txt%3A5343)

* _2026-07-21 15:25:37 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18402/pull-kubevirt-e2e-k8s-1.35-sig-network/2079554687773708288#1:build-log.txt%3A2550)

* _2026-07-21 14:23:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.34-sig-storage/2079548939106783232#1:build-log.txt%3A3656)

* _2026-07-21 14:17:49 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2079548938674769920#1:build-log.txt%3A5093)

* _2026-07-21 13:10:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18538/pull-kubevirt-e2e-k8s-1.35-sig-storage/2079550396535148544#1:build-log.txt%3A2572)

* _2026-07-21 11:38:35 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18402/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079528731050774528#1:build-log.txt%3A3794)

* _2026-07-21 10:34:47 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.35-sig-network/2079503746533953536#1:build-log.txt%3A4030)

* _2026-07-21 10:33:48 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.34-sig-compute/2079503746223575040#1:build-log.txt%3A2631)

* _2026-07-21 10:31:31 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.34-sig-network/2079503745925779456#1:build-log.txt%3A2541)

* _2026-07-21 07:32:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18499/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2079469516609294336#1:build-log.txt%3A4059)

* _2026-07-21 06:14:47 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18504/pull-kubevirt-e2e-k8s-1.34-sig-storage/2079449848859332608#1:build-log.txt%3A3651)

</details>

<hr/>
</details>

### 2026-07-20 (7x / 5.30%)


#### internal (3x / 42.86%)

<details>
<summary> make cluster lifecycle target failure (3x / 42.86%) </summary>

<hr/>

**3x**: _2026-07-20 22:21:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18499/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079305097958920192#1:build-log.txt%3A2044)
<details>
<summary>all...</summary>

* _2026-07-20 22:21:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18499/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079305097958920192#1:build-log.txt%3A2044)

* _2026-07-20 19:42:17 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18499/pull-kubevirt-e2e-k8s-1.36-sig-storage/2079289433311416320#1:build-log.txt%3A910)

* _2026-07-20 19:18:06 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.36-sig-storage/2079283980238065664#1:build-log.txt%3A849)

</details>

<hr/>
</details>

#### external (4x / 57.14%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (4x / 57.14%) </summary>

<hr/>

**4x**: _2026-07-20 22:17:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079296912455897088#1:build-log.txt%3A2656)
<details>
<summary>all...</summary>

* _2026-07-20 22:17:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079296912455897088#1:build-log.txt%3A2656)

* _2026-07-20 19:53:48 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079250202144804864#1:build-log.txt%3A2654)

* _2026-07-20 19:21:15 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.36-sig-performance/2079283962835898368#1:build-log.txt%3A5498)

* _2026-07-20 19:18:10 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.36-sig-compute/2079283981123063808#1:build-log.txt%3A2665)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (98x / 74.24%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (52x / 39.39%) </summary>

<hr/>

**52x**: _2026-07-22 18:11:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18550/pull-kubevirt-e2e-k8s-1.35-sig-compute/2079887455724507136#1:build-log.txt%3A2629)
<details>
<summary>all...</summary>

* _2026-07-26 19:16:47 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18469/pull-kubevirt-e2e-kind-1.36-sev/2081458617176821760#1:build-log.txt%3A1171)
<details><summary>context</summary>
<pre>19:24:19: Downloading cni-plugins-linux-amd64-v0.8.5.tgz
19:24:19: &#43;&#43; curl -sSL -o /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.36/cni-plugins-linux-amd64-v0.8.5.tgz https://github.com/containernetworking/plugins/releases/download/v0.8.5/cni-plugins-linux-amd64-v0.8.5.tgz
19:25:01: curl: (6) Could not resolve host: release-assets.githubusercontent.com
make: *** [Makefile:174: cluster-up] Error 6
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-07-26 17:10:22 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18469/pull-kubevirt-e2e-k8s-1.34-sig-compute/2081426767049920512#1:build-log.txt%3A2934)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
18:05:04: selecting podman as container runtime
18:05:04: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-compute is being used by the following container(s): 998c25c493b92b7e46700317645179b5e7356a93860fe61b3b6df38fcf2e153e: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-26 10:45:42 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18474/pull-kubevirt-e2e-k8s-1.35-sig-network/2081329966217170944#1:build-log.txt%3A4396)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
11:40:49: selecting podman as container runtime
11:40:49: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-network is being used by the following container(s): ccb5e4d339dc8c19d7e2fb93e9ba97615aa8a66688615658db41ae8106af3157: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-26 10:45:39 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18474/pull-kubevirt-e2e-k8s-1.34-sig-storage/2081329965713854464#1:build-log.txt%3A2866)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
11:40:23: selecting podman as container runtime
11:40:24: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-storage is being used by the following container(s): 20f6dcc7a58d0aa55d3b9a1ed6bb8ccf9f42192d7075aed2d409b98c36c42de2: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-26 06:44:05 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18569/pull-kubevirt-e2e-k8s-1.34-sig-compute/2081269171529715712#1:build-log.txt%3A2957)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
07:38:53: selecting podman as container runtime
07:38:53: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-compute is being used by the following container(s): e95f3b98250f5fcbe6c3c026fe4c2385e5a68678146317e7af5b6883d5c6d9c9: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-26 06:44:01 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18569/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2081269171055759360#1:build-log.txt%3A4441)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
07:40:26: selecting podman as container runtime
07:40:26: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network is being used by the following container(s): a7a916fa94d6d7d509dd5eadf83b874d3ffe7bce8aac56b150500fee8e1cb8c7: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 07:23:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18585/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080916635056607232#1:build-log.txt%3A5259)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
08:21:12: selecting podman as container runtime
08:21:12: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated is being used by the following container(s): 9a90442906657101c3cc4fcb270471addb466d55b4133169560e85bb387e8cba: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 07:23:10 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18585/pull-kubevirt-e2e-k8s-1.36-sig-network/2080916624981889024#1:build-log.txt%3A3651)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
08:27:33: selecting podman as container runtime
08:27:33: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network is being used by the following container(s): 9dff41e400929f3bca4005510c2c43e3537279d7cc670e69b8c6c391ac074eac: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 00:35:31 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18533/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080805660005502976#1:build-log.txt%3A4006)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
01:33:58: selecting podman as container runtime
01:33:58: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-operator is being used by the following container(s): b30f9afeaa118fb50ff871fb990c3722a6b734ada0aa7fb0ad71f11efb48d126: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 00:28:33 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18533/pull-kubevirt-e2e-k8s-1.35-sig-network/2080805657438588928#1:build-log.txt%3A4385)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
01:27:41: selecting podman as container runtime
01:27:41: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-network is being used by the following container(s): 688a52119a62e37c3dffcf75a89ad826c90ac7ab9b6cc2246015c2a934e12506: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 00:25:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18533/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080805657153376256#1:build-log.txt%3A2974)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
01:25:52: selecting podman as container runtime
01:25:53: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-compute is being used by the following container(s): a3739d37498d72231dd14dc3667825e5afd7b17eb198596e46f96915369f4038: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 16:23:44 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18428/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080570620118044672#1:build-log.txt%3A3565)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
17:20:10: selecting podman as container runtime
17:20:10: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network is being used by the following container(s): 413ce7d84b99869803ac0245f7877db03688d79dc76657686b67f226cf7ae9b2: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 14:17:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18400/pull-kubevirt-e2e-k8s-1.35-sig-storage/2080658114465501184#1:build-log.txt%3A2570)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
15:12:13: selecting podman as container runtime
15:12:13: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-storage is being used by the following container(s): 7993a63029b53c34fd4b056e3edc06a052b890e7f4caabd5e412534c42b4c411: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 14:16:53 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18168/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080651063739813888#1:build-log.txt%3A2651)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
15:12:05: selecting podman as container runtime
15:12:06: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-storage is being used by the following container(s): 2939b8a81980d5862523d4b4445488864ce1447aae71b93631c756e622406cb4: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 14:15:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18168/pull-kubevirt-e2e-k8s-1.34-sig-network/2080651063618179072#1:build-log.txt%3A4014)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
15:12:08: selecting podman as container runtime
15:12:08: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-network is being used by the following container(s): 2e843c0d892d1a14a5d27b654bbeca33caf1acbf966f702c7b2c3afcc2ee4903: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 13:21:54 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18443/pull-kubevirt-e2e-k8s-1.35-sig-network/2080642899699044352#1:build-log.txt%3A5078)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
14:17:28: selecting podman as container runtime
14:17:28: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-network is being used by the following container(s): b7cf44ceb82ba41a2eea54c2b53b950a2b1578216e5c2fd43ad58461041ae8c4: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 13:16:21 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18443/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080642898696605696#1:build-log.txt%3A3556)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
14:12:26: selecting podman as container runtime
14:12:26: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network is being used by the following container(s): 6641fed9bcef0139f3a82a901d678c973206417d50c154cd6f4c8e53b475116a: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 13:11:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18428/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080570620541669376#1:build-log.txt%3A3677)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
14:06:14: selecting podman as container runtime
14:06:14: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-compute is being used by the following container(s): a158fa941503181cf46708cf7ab287a84634040af1f2c7c1eb6d0e14a1a53313: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 13:09:32 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18428/pull-kubevirt-e2e-k8s-1.34-sig-network/2080570620331954176#1:build-log.txt%3A4014)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
14:04:49: selecting podman as container runtime
14:04:49: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-network is being used by the following container(s): 6052440d345f77b67f80bffd9eff86a718fed9159ccbd5b9f11bc8a1c7e118b1: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 07:51:34 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18542/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080561357886853120#1:build-log.txt%3A2659)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
08:53:07: selecting podman as container runtime
08:53:07: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-compute is being used by the following container(s): d77dbdbcc4d335a5e8f9bf6dfc9c5e2acf9915d29bc0fdbf211540987bc574cb: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 07:51:32 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18542/pull-kubevirt-e2e-k8s-1.36-sig-network/2080561377302286336#1:build-log.txt%3A5074)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
08:52:00: selecting podman as container runtime
08:52:00: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network is being used by the following container(s): 30a304e65dbc6341aaecb5527653241c691082ddb9cfcd9554fa6933bac79219: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 07:51:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18542/pull-kubevirt-e2e-k8s-1.34-sig-operator/2080561358025265152#1:build-log.txt%3A2572)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
08:50:13: selecting podman as container runtime
08:50:13: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-operator is being used by the following container(s): 49c0efac0faf91dda3098d84c2d06a63e0b9fddc63e0c16224a6baabcb49df85: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-23 18:32:12 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18400/pull-kubevirt-e2e-k8s-1.36-sig-network/2080359448429203456#1:build-log.txt%3A5050)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
19:28:15: selecting podman as container runtime
19:28:16: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network is being used by the following container(s): e71cf42884baa427a0c9e46d8cf8427177af11e5bbd7491ca1ff3f29f5ca3de7: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-23 15:40:52 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18428/pull-kubevirt-e2e-k8s-1.36-sig-storage/2080316875354411008#1:build-log.txt%3A3593)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
16:36:24: selecting podman as container runtime
16:36:25: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-storage is being used by the following container(s): d8cd49650664ee0d041298b030551e0a0ff7768da3e76acf3a9c17790b7fb516: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-23 14:26:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18296/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080297545166229504#1:build-log.txt%3A4045)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
15:22:41: selecting podman as container runtime
15:22:41: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network is being used by the following container(s): 39c3099b5e5ee0ade787337831598fa4fce185367538f2213fa8e09ae11b848f: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-23 13:25:38 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18296/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080267705973215232#1:build-log.txt%3A2570)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
14:20:30: selecting podman as container runtime
14:20:30: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network is being used by the following container(s): 6ab9b87e88094dfc234b561ce94298a9f04728f9c0a114097463ae8bf3430ff6: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-23 12:38:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18491/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2080239156256051200#1:build-log.txt%3A2662)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
13:32:52: selecting podman as container runtime
13:32:52: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute-serial is being used by the following container(s): 6f444d7ad4c811312655054d4a63fb4c031cdd43fe95d899dccde35b6921e15f: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-23 11:51:19 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18410/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080253331493621760#1:build-log.txt%3A3558)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:50:28: selecting podman as container runtime
12:50:28: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-operator is being used by the following container(s): 08747f4aedd0398d0829f79d6c6e97d5b14cf53c1b840b25106bdb8aadc10f3d: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-23 06:25:58 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18568/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2080177439815766016#1:build-log.txt%3A4360)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
07:23:01: selecting podman as container runtime
07:23:02: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations is being used by the following container(s): 7bc7881ad18f0ffb22060da6d530861a82063277443fab2c1e85676bed9550b0: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-23 03:35:15 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18453/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080107377016705024#1:build-log.txt%3A5485)
<details><summary>context</summary>
<pre>04:41:01: 
04:41:01: &#43; wait 32
04:41:22: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed
make: *** [Makefile:174: cluster-up] Error 1
&#43; EXIT_VALUE=2
&#43; set &#43;o xtrace
Cleaning up after podman in container.</pre>
</details>


* _2026-07-22 19:01:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17946/pull-kubevirt-e2e-k8s-1.34-sig-compute/2079925529477124096#1:build-log.txt%3A3703)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
19:56:13: selecting podman as container runtime
19:56:13: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-compute is being used by the following container(s): 5837b65f5b06323a284c3d546d9e61103cb50c13a0536fba79d6e9c1235249c6: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-22 18:11:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18550/pull-kubevirt-e2e-k8s-1.35-sig-compute/2079887455724507136#1:build-log.txt%3A2629)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
19:07:28: selecting podman as container runtime
19:07:28: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-compute is being used by the following container(s): 47ef8a246d9a7ac1f39cede4d61695ef451aeee229490d6442b29370985ce066: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-22 15:50:33 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18530/pull-kubevirt-e2e-k8s-1.34-sig-network/2079956261721346048#1:build-log.txt%3A3523)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
16:45:38: selecting podman as container runtime
16:45:39: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-network is being used by the following container(s): 96543283dc5a57508443e4b78fa003140212f2e553d86cd66f3b1e8fbc2f1887: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-22 12:18:49 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18552/pull-kubevirt-e2e-kind-1.34-sev-1.9/2079903851376283648#1:build-log.txt%3A5353)
<details><summary>context</summary>
<pre>13:02:36: failed to fetch vm exports: the server could not find the requested resource (get virtualmachinepools.pool.kubevirt.io)
13:02:36: failed to fetch migration policies: the server could not find the requested resource (get migrationpolicies.migrations.kubevirt.io)
13:02:36: &#43; rm -f /tmp/kubevirt.deploy.I0OY
make: *** [Makefile:189: cluster-sync] Error 1
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>


* _2026-07-22 12:05:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18550/pull-kubevirt-e2e-k8s-1.34-sig-network/2079887454780788736#1:build-log.txt%3A5057)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
13:03:13: selecting podman as container runtime
13:03:13: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-network is being used by the following container(s): 0e41d3493a5f22676a1ca970b810cf93ccfaa1b16810a291923512b8386cd6be: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-21 17:49:35 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18436/pull-kubevirt-e2e-k8s-1.35-sig-network/2079623487349067776#1:build-log.txt%3A2546)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
18:45:10: selecting podman as container runtime
18:45:10: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-network is being used by the following container(s): b6ff7c9515dc057908d0c03978abcd0b5869a444b18481f53f513688eadbf90b: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-21 15:43:51 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18301/pull-kubevirt-e2e-k8s-1.35-sig-compute/2079588132977643520#1:build-log.txt%3A2628)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
16:38:55: selecting podman as container runtime
16:38:55: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-compute is being used by the following container(s): 98c8326caed83f3c212d882bb43ef33ae36e3954a72cb1b6eceedbb4f902a3ee: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-21 15:28:48 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18301/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2079588132075868160#1:build-log.txt%3A5343)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
16:25:01: selecting podman as container runtime
16:25:01: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated is being used by the following container(s): 2a92063b43fa87668b7e4bea986759ec43b6c26c684f724ce4832e988e8550ac: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-21 15:25:37 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18402/pull-kubevirt-e2e-k8s-1.35-sig-network/2079554687773708288#1:build-log.txt%3A2550)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
16:20:18: selecting podman as container runtime
16:20:19: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-network is being used by the following container(s): 91eeb034947d380ffe509ebb5e25ee6b9ba187d214d5eee1c79e6a03a3be75ef: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-21 14:23:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.34-sig-storage/2079548939106783232#1:build-log.txt%3A3656)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
15:21:21: selecting podman as container runtime
15:21:21: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-storage is being used by the following container(s): 37d78d65fe53c0a96854c31e4f9e3d979f26807390689d46c7deab55eefb8cec: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-21 14:17:49 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2079548938674769920#1:build-log.txt%3A5093)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
15:15:01: selecting podman as container runtime
15:15:01: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network is being used by the following container(s): 5b05397e2829de9dadd658da93825bd080814824e0fa87b7a38fcc155eacbcc7: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-21 13:10:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18538/pull-kubevirt-e2e-k8s-1.35-sig-storage/2079550396535148544#1:build-log.txt%3A2572)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
14:05:56: selecting podman as container runtime
14:05:56: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-storage is being used by the following container(s): 3b3c3108be88fc85b6333934540d1c98f83e807525d7d91e9fa13faf4f5ff548: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-21 11:38:35 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18402/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079528731050774528#1:build-log.txt%3A3794)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:33:47: selecting podman as container runtime
12:33:47: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute-serial is being used by the following container(s): bb6511c3464076c818ef63a9304ab43c556ebf08a555fd940c423a6f9ec52f1c: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-21 10:34:47 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.35-sig-network/2079503746533953536#1:build-log.txt%3A4030)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
11:30:30: selecting podman as container runtime
11:30:30: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-network is being used by the following container(s): 0f98ecad74a8a95fda821525b3f1eb9758e8d48da04951585f9c8572eda7fb5f: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-21 10:33:48 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.34-sig-compute/2079503746223575040#1:build-log.txt%3A2631)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
11:29:42: selecting podman as container runtime
11:29:42: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-compute is being used by the following container(s): 7cf59713a9329c6ca8ecfd9bc2ecdabd79fa908df81a859c19e45b5e1b300984: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-21 10:31:31 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.34-sig-network/2079503745925779456#1:build-log.txt%3A2541)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
11:26:06: selecting podman as container runtime
11:26:06: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-network is being used by the following container(s): e12861dba6ab006a176a859c1dbf618a06508a9998e83400f3ec0d26aaf4fd34: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-21 07:32:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18499/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2079469516609294336#1:build-log.txt%3A4059)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
08:28:45: selecting podman as container runtime
08:28:45: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network is being used by the following container(s): 59cbcc6eba1b9b5e680007f337ce5e8c54e06130952f9592d63e60951d74ee44: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-21 06:14:47 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18504/pull-kubevirt-e2e-k8s-1.34-sig-storage/2079449848859332608#1:build-log.txt%3A3651)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
07:11:46: selecting podman as container runtime
07:11:46: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-storage is being used by the following container(s): e86ba5bbe8cf186d10ac8159730d1f85078c70037def1e0703a2c540858f3775: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-20 22:17:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079296912455897088#1:build-log.txt%3A2656)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
23:11:57: selecting podman as container runtime
23:11:57: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute-serial is being used by the following container(s): 847d6b62242399f340d1c2d2f87305e3b1a46dc3a093064131b3b875d24ff0a4: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-20 19:53:48 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079250202144804864#1:build-log.txt%3A2654)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
20:48:26: selecting podman as container runtime
20:48:26: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute-serial is being used by the following container(s): ebcc44748e08a020c58824b2f9fee3f6e2f70a203064765a2be40829df0602a2: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-20 19:21:15 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.36-sig-performance/2079283962835898368#1:build-log.txt%3A5498)
<details><summary>context</summary>
<pre>20:24:49: 
20:24:49: &#43; wait 38
20:25:00: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed
make: *** [Makefile:174: cluster-up] Error 1
&#43; EXIT_VALUE=2
&#43; set &#43;o xtrace
Cleaning up after podman in container.</pre>
</details>


* _2026-07-20 19:18:10 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.36-sig-compute/2079283981123063808#1:build-log.txt%3A2665)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
20:15:20: selecting podman as container runtime
20:15:20: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute is being used by the following container(s): 3554a944fd4cf7e05ac3dedbf56d9c2085157280eee12537018af9b4d75da1c3: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> bazel remote cache blob fetch failure (from secondary snippet) (44x / 33.33%) </summary>

<hr/>

**44x**: _2026-07-25 07:26:15 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-compute/2080916717827002368#1:build-log.txt%3A434)
<details>
<summary>all...</summary>

* _2026-07-25 07:51:26 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-windows2016/2080923721920417792#1:build-log.txt%3A415)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
07:55:01: selecting podman as container runtime
07:55:01: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-windows2016 is being used by the following container(s): 66aad5a6f8cc2e01d73964624180726499dd176335ec49e5526a7070c215ddf4: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 07:28:43 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080916717994774528#1:build-log.txt%3A384)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
07:33:42: selecting podman as container runtime
07:33:42: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-operator is being used by the following container(s): adffe9fba1d0cb87be5c3efb7e33a7c0e5806148b7fc7eb88fd0b8ae4da092e9: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 07:26:15 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-compute/2080916717827002368#1:build-log.txt%3A434)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
07:33:29: selecting podman as container runtime
07:33:29: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-compute is being used by the following container(s): 62812e473643e1b6e36fe86a06201439c27a48499109cc907bffb246458722fe: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 07:23:37 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-storage/2080916717680201728#1:build-log.txt%3A438)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
07:31:12: selecting podman as container runtime
07:31:12: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-storage is being used by the following container(s): e3b5938d714aeefb828b6800564037ffa67cf5522a821b247864390ba70cf9e2: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 07:23:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-operator/2080916717386600448#1:build-log.txt%3A449)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
07:31:31: selecting podman as container runtime
07:31:31: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-operator is being used by the following container(s): f06b1b99b3f1255b80ced29c51ea6e166bf77b75152da977e5e33fcb41ce09f1: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 07:23:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080916717092999168#1:build-log.txt%3A449)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
07:29:26: selecting podman as container runtime
07:29:26: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-storage is being used by the following container(s): 1b7e817ca7d20316bcdc0b7e031686188b7ead698af5f876e1d6c493d730855d: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 07:23:35 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080916716325441536#1:build-log.txt%3A460)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
07:34:29: selecting podman as container runtime
07:34:30: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network is being used by the following container(s): 8fcc89e7a913ab998c9840f50c27517bd031a186aa83418f810165275db3e6eb: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 07:23:34 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080916716623237120#1:build-log.txt%3A457)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
07:32:43: selecting podman as container runtime
07:32:43: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated is being used by the following container(s): 92c74fbfb884800f0c8e229285363ecf71174896824ff5464213f864342a1873: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 07:23:34 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-network/2080916716866506752#1:build-log.txt%3A367)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
07:27:50: selecting podman as container runtime
07:27:50: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-network is being used by the following container(s): 0cbeda65ae77665ae7985b5e7a28d9f1639c71f55736b35694552d60bc0b9dd2: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 07:23:34 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-network/2080916717504040960#1:build-log.txt%3A432)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
07:31:26: selecting podman as container runtime
07:31:26: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-network is being used by the following container(s): 5743f0398a9e8309738e4d1148b503a36ef0d0698bea2db96152d14df3d4543b: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 07:23:30 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080916717239799808#1:build-log.txt%3A433)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
07:31:05: selecting podman as container runtime
07:31:06: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-compute is being used by the following container(s): f9ba93c98fcd637913710d5a784f90924857ed5f52b3f9764bd552be95c5d8cb: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 03:48:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-windows2016/2080862637561745408#1:build-log.txt%3A447)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
03:52:40: selecting podman as container runtime
03:52:41: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-windows2016 is being used by the following container(s): 5c24271aaa8133562026d7f583451b2feb6a0e89341ccbe1212b02a5b5e72cf8: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 02:20:43 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-network/2080840488520257536#1:build-log.txt%3A444)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
02:26:07: selecting podman as container runtime
02:26:07: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-network is being used by the following container(s): 297c79641c689bff680e96554047c2277d3b8ce8a46f0fdcc26c187aeeca353f: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 02:20:42 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080840488079855616#1:build-log.txt%3A438)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
02:26:54: selecting podman as container runtime
02:26:54: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-storage is being used by the following container(s): cdedd3805e03fd29c4aa343c22c534da15fc246dbe8aa97251367f4efd913e1d: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 02:20:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080840487677202432#1:build-log.txt%3A429)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
02:25:56: selecting podman as container runtime
02:25:56: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network is being used by the following container(s): e97058d0a972415d2f6c7c141192f43d70ddd861430f39bdca99c37ac6a81715: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 02:20:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-storage/2080840488658669568#1:build-log.txt%3A388)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
02:26:08: selecting podman as container runtime
02:26:08: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-storage is being used by the following container(s): 8b2b6d8e57839c0b861549317ffa4717515cf2789ee541c91c7595ebf6878432: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 02:20:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-network/2080840487949832192#1:build-log.txt%3A425)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
02:26:08: selecting podman as container runtime
02:26:08: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-network is being used by the following container(s): 2c2ca5208dc037aaa3d3c47bb98fa97a877d0a204a12d861b48325738027b195: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 02:20:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-operator/2080840488360873984#1:build-log.txt%3A389)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
02:25:44: selecting podman as container runtime
02:25:45: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-operator is being used by the following container(s): 02ec8c67a48a1aa51dc138bd32b606a0ecbd48e8c45902bd357d4b9f4808c377: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 02:20:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-compute/2080840488834830336#1:build-log.txt%3A402)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
02:25:51: selecting podman as container runtime
02:25:52: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-compute is being used by the following container(s): 4f95de0ebf2bdeb2be49cf7c82e7f66f1c77d76307bfb0b87dfab0729f81f892: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 02:20:40 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080840488197296128#1:build-log.txt%3A372)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
02:26:09: selecting podman as container runtime
02:26:09: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-compute is being used by the following container(s): 8db7aa9ac18ba1d047bd0d2aff0265c6b89c7dc6aec0b4e8a072fc7f666a4227: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 02:20:40 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080840487815614464#1:build-log.txt%3A364)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
02:25:25: selecting podman as container runtime
02:25:25: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated is being used by the following container(s): ec1acf5954122f665b709574aab40c5327847def174a54ef1b9d0788e57c8e13: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-25 02:20:39 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080840488973242368#1:build-log.txt%3A392)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
02:25:43: selecting podman as container runtime
02:25:43: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-operator is being used by the following container(s): 4a03427aa8fba08a03c490a3f4fc8c6bb4fca18e71935aa2451b4b10a66456e9: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 23:32:17 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-windows2016/2080797902535397376#1:build-log.txt%3A476)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
23:37:40: selecting podman as container runtime
23:37:40: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-windows2016 is being used by the following container(s): 32c3d27848134cd37b52ff1fe5b61fc6182d0def588bddab843f01b692d27b99: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 22:46:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080764931073904640#1:build-log.txt%3A373)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
22:49:23: selecting podman as container runtime
22:49:24: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-operator is being used by the following container(s): b8c0ded224783312285404f117b145611da05ab9424fdd7dbaeb583d3248ca70: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 22:45:46 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-storage/2080764930146963456#1:build-log.txt%3A406)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
22:51:20: selecting podman as container runtime
22:51:21: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-storage is being used by the following container(s): a2a84a757deb139981d2ed6c7bad4c9d1cb519277d3750cceac4dcfa822dc544: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 22:45:44 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-compute/2080764930268598272#1:build-log.txt%3A388)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
22:49:26: selecting podman as container runtime
22:49:26: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-compute is being used by the following container(s): 74588e0b91831e93adbdba4ed3ca45b8204639337a2b9fddd1830b9b76577cc9: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 22:44:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-network/2080764930012745728#1:build-log.txt%3A406)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
22:49:24: selecting podman as container runtime
22:49:24: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-network is being used by the following container(s): c408f6d44416175166f7850dde3c7e35cdce272f48801358dd3b7f5e07708f40: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 22:43:29 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-operator/2080764929849167872#1:build-log.txt%3A375)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
22:47:53: selecting podman as container runtime
22:47:53: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-operator is being used by the following container(s): c73839f652412d0455afef398506456ec011945ead15f8ee2bc7ade1db775649: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 22:42:00 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080764929639452672#1:build-log.txt%3A386)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
22:45:51: selecting podman as container runtime
22:45:51: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-compute is being used by the following container(s): 82354622cc267aa0291fb34732ab4c0faaf4d4a920aa84bce786835c8b0f8682: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 22:41:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080764929391988736#1:build-log.txt%3A394)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
22:44:52: selecting podman as container runtime
22:44:52: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-storage is being used by the following container(s): c5a6d09cfd19a03b093f83dafdf6eeefd11678ca6962e9cd02248dfe1e2aff67: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 22:39:57 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-network/2080764929308102656#1:build-log.txt%3A405)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
22:44:31: selecting podman as container runtime
22:44:31: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-network is being used by the following container(s): 4b4f31b1a85f6c2260167ad3437f6732fc89d284fab2be374e0e7a840b9ce31a: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 22:38:34 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080764929115164672#1:build-log.txt%3A399)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
22:42:34: selecting podman as container runtime
22:42:35: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated is being used by the following container(s): 43402cf36265b14c9d8e68a0a75524573c537888ae3cf6aa640ceb6792540595: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 22:35:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080764928976752640#1:build-log.txt%3A382)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
22:39:08: selecting podman as container runtime
22:39:08: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network is being used by the following container(s): 2f60959626232796976f265d23c212bb406bb29e28313fd41513431e35ca3047: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 18:13:39 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080711464485654528#1:build-log.txt%3A455)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
18:20:22: selecting podman as container runtime
18:20:22: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-operator is being used by the following container(s): 21c7c1213d98214d7987425cb153b19cd7a16d9545a9a8f7e9212ed2cefd10ea: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 18:12:50 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-compute/2080711464292716544#1:build-log.txt%3A434)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
18:17:07: selecting podman as container runtime
18:17:07: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-compute is being used by the following container(s): 71513927761f2a89fd7512a532a9692e2888aa8da268a046cecda259e4774b3f: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 17:59:31 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-storage/2080711464133332992#1:build-log.txt%3A375)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
18:03:12: selecting podman as container runtime
18:03:12: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-storage is being used by the following container(s): b041be59a4c750fa7a369370a27aca9ea47b79d13e878800c9889b377f951895: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 17:59:18 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-network/2080711463994920960#1:build-log.txt%3A397)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
18:04:16: selecting podman as container runtime
18:04:16: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-network is being used by the following container(s): 041e308b03b62801c978fa92dc8bdcc0d01e9bbf7a71d78c478b00ed481e1a7b: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 17:59:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-operator/2080711463919423488#1:build-log.txt%3A453)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
18:03:53: selecting podman as container runtime
18:03:53: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-operator is being used by the following container(s): bff13bb461f170f15c1ecf5c8c6dd73ea65813a9fd1597468d0dc0e8ea2d934a: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 17:56:22 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080711463818760192#1:build-log.txt%3A385)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
18:00:09: selecting podman as container runtime
18:00:09: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-compute is being used by the following container(s): be352ad5a735af50289941e89bb445a772d2a68b5284e15f644b1ed543c218a9: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 17:56:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080711463592267776#1:build-log.txt%3A441)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
18:02:14: selecting podman as container runtime
18:02:14: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-storage is being used by the following container(s): c82b18bb2035ecacfed0df178e4f350cc421f2e01f811ecd8c28d41b2359ed75: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 17:55:57 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-network/2080711463470632960#1:build-log.txt%3A451)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
18:02:14: selecting podman as container runtime
18:02:14: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-network is being used by the following container(s): 54c9345691b3c3286971025d46ce3af6f70e04df269f2e1cc94f9d88d73e3820: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 17:52:43 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080711463294472192#1:build-log.txt%3A397)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
17:58:11: selecting podman as container runtime
17:58:12: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated is being used by the following container(s): aaa4da36c2753528c91d5464b4d8158892b79243217755ca7a037a3976af5c6b: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 17:52:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080711463156060160#1:build-log.txt%3A447)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
17:58:27: selecting podman as container runtime
17:58:28: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network is being used by the following container(s): fab5bb7f9d6b2e7058e2b92b23e1d556fba60e56d4475a20b469b8be582efdff: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 17:48:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-windows2016/2080711462820515840#1:build-log.txt%3A395)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
17:51:56: selecting podman as container runtime
17:51:57: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-windows2016 is being used by the following container(s): 85a4aea61d2ee39d8092d99b8066df16503972bb12cbd33f33f72d6ffd7bc1e8: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> download failure in context (2x / 1.52%) </summary>

<hr/>

**2x**: _2026-07-26 17:10:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18469/pull-kubevirt-e2e-k8s-1.35-sig-network/2081426767381270528#1:build-log.txt%3A2180)
<details>
<summary>all...</summary>

* _2026-07-26 17:10:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18469/pull-kubevirt-e2e-k8s-1.35-sig-network/2081426767381270528#1:build-log.txt%3A2180)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
18:05:50: selecting podman as container runtime
18:05:50: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-network is being used by the following container(s): 97078fde3f249bbec8815d0c042e183a4757be86aa9faabe41e383861b5e51ab: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 16:07:20 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18400/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080685210604670976#1:build-log.txt%3A1193)
<details><summary>context</summary>
<pre>16:19:21: INFO: 0 processes.
16:19:21: ERROR: Build failed. Not running target
16:19:23: &#43; rm -f /tmp/kubevirt.deploy.OSZM
make: *** [Makefile:189: cluster-sync] Error 1
&#43; EXIT_VALUE=2
&#43; set &#43;o xtrace
Cleaning up after podman in container.</pre>
</details>


</details>

<hr/>
</details>

### internal (33x / 25.00%)

<details>
<summary> make cluster lifecycle target failure (33x / 25.00%) </summary>

<hr/>

**33x**: _2026-07-22 19:08:17 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17946/pull-kubevirt-e2e-k8s-1.35-sig-compute/2079925529972051968#1:build-log.txt%3A836)
<details>
<summary>all...</summary>

* _2026-07-26 17:10:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18469/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2081426766697598976#1:build-log.txt%3A1131)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
18:05:48: selecting podman as container runtime
18:05:48: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated is being used by the following container(s): 1f0f8da9e299c1e489777444f7938832dc5283c11683e891d78f0f8f30efc4c8: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-26 10:45:41 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18474/pull-kubevirt-e2e-k8s-1.34-windows2016/2081329964916936704#1:build-log.txt%3A2120)
<details><summary>context</summary>
<pre>11:41:06: 6audio: Could not create a backend for voice `ac97.po&#39;
11:41:16: Error response from daemon: container 227570d07db44e2e14cc7ff5271ceb7c8644181c5274f64ad10aa16301ae4489 has dependent containers which must be removed before it: be1624e31dafcd41b3cbd219a36ecce4e36010a10beb1c3b37e18ab886cea6c1: container already exists
11:41:31: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed
make: *** [Makefile:174: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-07-25 07:23:09 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18585/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080916614278025216#1:build-log.txt%3A789)
<details><summary>context</summary>
<pre>08:23:27: 
08:23:27: &#43; wait 47
08:23:37: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed
make: *** [Makefile:174: cluster-up] Error 1
&#43; EXIT_VALUE=2
&#43; set &#43;o xtrace
Cleaning up after podman in container.</pre>
</details>


* _2026-07-25 00:13:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18533/pull-kubevirt-e2e-k8s-1.36-sig-network/2080805647405813760#1:build-log.txt%3A2212)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
01:11:20: selecting podman as container runtime
01:11:20: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network is being used by the following container(s): 4f51a70115c40a5fb1dfc8fce6be7b564b4187f80bb2a8de5ca2e06ca238f3a8: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 13:16:39 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18443/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080642898814046208#1:build-log.txt%3A1929)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
14:11:51: selecting podman as container runtime
14:11:52: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated is being used by the following container(s): 3174ab87ac991e3db2ac07ff4c77e75401943651303a25b317a652975364d305: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 13:09:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18428/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080570620432617472#1:build-log.txt%3A1966)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
14:04:44: selecting podman as container runtime
14:04:44: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-storage is being used by the following container(s): a838177b8d4cc9ac6a0b609cdb09160117c5df72d8c99d67a2140284e4b285d1: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 07:51:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18542/pull-kubevirt-e2e-k8s-1.35-sig-network/2080561358188843008#1:build-log.txt%3A841)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
08:50:03: selecting podman as container runtime
08:50:03: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-network is being used by the following container(s): 501092901c3d01aaf3f430827d9082e2d766a14df53deae88ace8dc5b8621b34: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-24 07:51:32 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18542/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080561368284532736#1:build-log.txt%3A1855)
<details><summary>context</summary>
<pre>08:51:46: &lt;sdl: Reason: ALSA: Couldn&#39;t open audio device: Host is down
08:51:46: 6audio: Could not create a backend for voice `ac97.pi&#39;
08:51:59: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed
make: *** [Makefile:174: cluster-up] Error 1
&#43; EXIT_VALUE=2
&#43; set &#43;o xtrace
Cleaning up after podman in container.</pre>
</details>


* _2026-07-23 21:52:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18400/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080410663418597376#1:build-log.txt%3A1825)
<details><summary>context</summary>
<pre>22:53:25: &lt;sdl: Reason: ALSA: Couldn&#39;t open audio device: Host is down
22:53:25: 6audio: Could not create a backend for voice `ac97.po&#39;
22:53:35: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed
make: *** [Makefile:174: cluster-up] Error 1
&#43; EXIT_VALUE=2
&#43; set &#43;o xtrace
Cleaning up after podman in container.</pre>
</details>


* _2026-07-23 16:11:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18568/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080233043007639552#1:build-log.txt%3A1943)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
17:06:04: selecting podman as container runtime
17:06:05: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-compute is being used by the following container(s): c425c06031068332055c6bde05770a24d32120b07d6a239be2eeb7aed8f0a314: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-23 11:52:45 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18414/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080250895450247168#1:build-log.txt%3A832)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:46:55: selecting podman as container runtime
12:46:55: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-storage is being used by the following container(s): 6d86e3250305f62d1d8881942332fdfbcccd98199cdfcf2a74f6206d02fa21f6: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-23 10:43:51 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18296/pull-kubevirt-e2e-k8s-1.34-sig-network/2080239905874644992#1:build-log.txt%3A834)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
11:37:49: selecting podman as container runtime
11:37:49: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-network is being used by the following container(s): a6522021541a28785bb561f976043a0448e1c96f334485245ee7fde1f42cf621: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-23 10:35:26 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18296/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080239905669124096#1:build-log.txt%3A836)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
11:29:40: selecting podman as container runtime
11:29:40: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network is being used by the following container(s): a622b7c4efc28d380ca8bd14aa94c53d40bb2440b1e5922ec1d0c108920479cc: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-22 19:08:17 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17946/pull-kubevirt-e2e-k8s-1.35-sig-compute/2079925529972051968#1:build-log.txt%3A836)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
20:02:47: selecting podman as container runtime
20:02:48: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-compute is being used by the following container(s): 8cc3664bcd6c65dad34aea9b69c608edb2a7d2b45698dd1e0f377308e3e4596a: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-22 19:07:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17946/pull-kubevirt-e2e-k8s-1.35-sig-storage/2079925529808474112#1:build-log.txt%3A835)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
20:02:22: selecting podman as container runtime
20:02:22: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-storage is being used by the following container(s): 3bf717a5078de15ec6f988bda935c8f3daedbf1bac03fb37cb72670036db4652: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-22 16:07:50 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18530/pull-kubevirt-e2e-k8s-1.34-windows2016/2079961254075568128#1:build-log.txt%3A832)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
17:02:26: selecting podman as container runtime
17:02:26: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-windows2016 is being used by the following container(s): bf4aa395cdc7544cb86da4302a89c33834e2fce938b19af51e158a98d3ab43a9: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-22 15:51:21 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18530/pull-kubevirt-e2e-k8s-1.34-sig-storage/2079956261884923904#1:build-log.txt%3A1902)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
16:46:36: selecting podman as container runtime
16:46:36: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-storage is being used by the following container(s): 9db7ce239289b49e492384c8b7645c91bd8ff4bbe57a0810f9541dd3b88c4e0e: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-22 13:28:23 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18565/pull-kubevirt-e2e-kind-1.34-sev-1.9/2079921377246187520#1:build-log.txt%3A812)
<details><summary>context</summary>
<pre>13:37:40: 	sigs.k8s.io/kind/pkg/errors/concurrent.go:30
13:37:40: runtime.goexit
13:37:40: 	runtime/asm_amd64.s:1693
make: *** [Makefile:174: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-07-22 12:33:47 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18530/pull-kubevirt-e2e-kind-1.36-sev/2079907633388589056#1:build-log.txt%3A648)
<details><summary>context</summary>
<pre>12:41:13: 	sigs.k8s.io/kind/pkg/errors/concurrent.go:30
12:41:13: runtime.goexit
12:41:13: 	runtime/asm_amd64.s:1771
make: *** [Makefile:174: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-07-22 12:27:27 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18562/pull-kubevirt-e2e-kind-1.34-sev-1.7/2079905794379223040#1:build-log.txt%3A655)
<details><summary>context</summary>
<pre>12:35:11: 	sigs.k8s.io/kind/pkg/errors/concurrent.go:30
12:35:11: runtime.goexit
12:35:11: 	runtime/asm_amd64.s:1700
make: *** [Makefile:159: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-07-22 11:52:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18453/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2079896117457195008#1:build-log.txt%3A866)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:48:08: selecting podman as container runtime
12:48:09: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated is being used by the following container(s): 37f176b379ecdf784d3552e032c06b795e15ffbd020118f483849962985b5b26: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-22 11:24:21 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18550/pull-kubevirt-e2e-k8s-1.36-sig-compute/2079884525898305536#1:build-log.txt%3A2062)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:19:16: selecting podman as container runtime
12:19:16: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute is being used by the following container(s): 81445f564a6c040610126061be2d9727af3c99048692ed8872fa5f9cf2beb2d5: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-22 10:36:23 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18436/pull-kubevirt-e2e-k8s-1.36-sig-network/2079877999724859392#1:build-log.txt%3A1911)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
11:31:34: selecting podman as container runtime
11:31:34: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network is being used by the following container(s): 33b6c5f9db67897cfd1061d340ce7ed2e11153e6a9855a63fe63f20f4fa30e3e: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-21 17:45:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18436/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2079623486266937344#1:build-log.txt%3A1895)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
18:40:08: selecting podman as container runtime
18:40:08: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated is being used by the following container(s): 847092e24f366edc35856845c3ad1c962fd3adae92fb80be6dc99156696c769e: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-21 15:53:11 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.34-sig-operator/2079593296690155520#1:build-log.txt%3A834)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
16:47:01: selecting podman as container runtime
16:47:01: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-operator is being used by the following container(s): de6db24e9ed14e7f5137f7ee2045739b3f80907c459577cb67e585c7e158102e: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-21 14:26:15 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.34-sig-operator/2079548939396190208#1:build-log.txt%3A835)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
15:21:16: selecting podman as container runtime
15:21:17: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-operator is being used by the following container(s): 8a586a10a0c3e2f4fbd340379fb33c9244bccd6beed715c20f0dbe29064b23a2: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-21 12:59:49 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18538/pull-kubevirt-e2e-k8s-1.34-sig-network/2079550395427852288#1:build-log.txt%3A835)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
13:54:02: selecting podman as container runtime
13:54:03: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-network is being used by the following container(s): acd450aa783a54b5836b0671471250d02590970911494ad333ec395adff094e0: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-21 11:37:21 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18402/pull-kubevirt-e2e-k8s-1.36-sig-network/2079528728521609216#1:build-log.txt%3A1836)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:34:56: selecting podman as container runtime
12:34:57: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network is being used by the following container(s): 0715d109c4db428e41c657bfe2294c0056cd5efa78141c70cb0cf12b020b2d36: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-21 07:32:57 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18499/pull-kubevirt-e2e-k8s-1.34-sig-network/2079469517045501952#1:build-log.txt%3A839)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
08:27:40: selecting podman as container runtime
08:27:40: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-network is being used by the following container(s): 907482bd7bbb60f56e2b2f6980f9e64dfa715b7ee38769a0ed763b9c37f01966: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-21 07:32:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18499/pull-kubevirt-e2e-k8s-1.35-sig-network/2079469518014386176#1:build-log.txt%3A841)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
08:27:44: selecting podman as container runtime
08:27:44: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.35-sig-network is being used by the following container(s): d652845f450829b63df0996cef889205b3bc434050aa48543ed621ac8c938382: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-20 22:21:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18499/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079305097958920192#1:build-log.txt%3A2044)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
23:15:36: selecting podman as container runtime
23:15:36: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute-serial is being used by the following container(s): 94eb6c5733aae29c981a0f1000da25860c9d4adae82ec3c4ec0790a16c039f5e: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-20 19:42:17 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18499/pull-kubevirt-e2e-k8s-1.36-sig-storage/2079289433311416320#1:build-log.txt%3A910)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
20:36:44: selecting podman as container runtime
20:36:44: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-storage is being used by the following container(s): bd8e3ab7c2afe926125b0705151c4ffc412f903c69906d543a3d1c197a7013c7: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-20 19:18:06 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.36-sig-storage/2079283980238065664#1:build-log.txt%3A849)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
20:12:43: selecting podman as container runtime
20:12:43: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-storage is being used by the following container(s): 0227cd7fbef60f70dc5203b04b3517e05ca6a3d95ea1eb9fe8f39d79746eb2bd: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>

### needs-investigation (1x / 0.76%)

<details>
<summary> no error snippets (1x / 0.76%) </summary>

<hr/>

**1x**: _2026-07-22 13:32:04 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18168/pull-kubevirt-e2e-kind-1.36-sev/2079922111794647040)
<details>
<summary>all...</summary>

* _2026-07-22 13:32:04 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18168/pull-kubevirt-e2e-kind-1.36-sev/2079922111794647040)

</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (129x / 97.73%)


#### external (97x / 75.19%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (51x / 39.53%) </summary>

<hr/>

**51x**: _2026-07-22 18:11:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18550/pull-kubevirt-e2e-k8s-1.35-sig-compute/2079887455724507136#1:build-log.txt%3A2629)
<details>
<summary>all...</summary>

* _2026-07-26 19:16:47 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18469/pull-kubevirt-e2e-kind-1.36-sev/2081458617176821760#1:build-log.txt%3A1171)

* _2026-07-26 17:10:22 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18469/pull-kubevirt-e2e-k8s-1.34-sig-compute/2081426767049920512#1:build-log.txt%3A2934)

* _2026-07-26 10:45:42 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18474/pull-kubevirt-e2e-k8s-1.35-sig-network/2081329966217170944#1:build-log.txt%3A4396)

* _2026-07-26 10:45:39 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18474/pull-kubevirt-e2e-k8s-1.34-sig-storage/2081329965713854464#1:build-log.txt%3A2866)

* _2026-07-26 06:44:05 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18569/pull-kubevirt-e2e-k8s-1.34-sig-compute/2081269171529715712#1:build-log.txt%3A2957)

* _2026-07-26 06:44:01 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18569/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2081269171055759360#1:build-log.txt%3A4441)

* _2026-07-25 07:23:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18585/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080916635056607232#1:build-log.txt%3A5259)

* _2026-07-25 07:23:10 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18585/pull-kubevirt-e2e-k8s-1.36-sig-network/2080916624981889024#1:build-log.txt%3A3651)

* _2026-07-25 00:35:31 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18533/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080805660005502976#1:build-log.txt%3A4006)

* _2026-07-25 00:28:33 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18533/pull-kubevirt-e2e-k8s-1.35-sig-network/2080805657438588928#1:build-log.txt%3A4385)

* _2026-07-25 00:25:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18533/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080805657153376256#1:build-log.txt%3A2974)

* _2026-07-24 16:23:44 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18428/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080570620118044672#1:build-log.txt%3A3565)

* _2026-07-24 14:17:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18400/pull-kubevirt-e2e-k8s-1.35-sig-storage/2080658114465501184#1:build-log.txt%3A2570)

* _2026-07-24 14:16:53 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18168/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080651063739813888#1:build-log.txt%3A2651)

* _2026-07-24 14:15:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18168/pull-kubevirt-e2e-k8s-1.34-sig-network/2080651063618179072#1:build-log.txt%3A4014)

* _2026-07-24 13:21:54 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18443/pull-kubevirt-e2e-k8s-1.35-sig-network/2080642899699044352#1:build-log.txt%3A5078)

* _2026-07-24 13:16:21 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18443/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080642898696605696#1:build-log.txt%3A3556)

* _2026-07-24 13:11:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18428/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080570620541669376#1:build-log.txt%3A3677)

* _2026-07-24 13:09:32 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18428/pull-kubevirt-e2e-k8s-1.34-sig-network/2080570620331954176#1:build-log.txt%3A4014)

* _2026-07-24 07:51:34 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18542/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080561357886853120#1:build-log.txt%3A2659)

* _2026-07-24 07:51:32 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18542/pull-kubevirt-e2e-k8s-1.36-sig-network/2080561377302286336#1:build-log.txt%3A5074)

* _2026-07-24 07:51:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18542/pull-kubevirt-e2e-k8s-1.34-sig-operator/2080561358025265152#1:build-log.txt%3A2572)

* _2026-07-23 18:32:12 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18400/pull-kubevirt-e2e-k8s-1.36-sig-network/2080359448429203456#1:build-log.txt%3A5050)

* _2026-07-23 15:40:52 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18428/pull-kubevirt-e2e-k8s-1.36-sig-storage/2080316875354411008#1:build-log.txt%3A3593)

* _2026-07-23 14:26:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18296/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080297545166229504#1:build-log.txt%3A4045)

* _2026-07-23 13:25:38 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18296/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080267705973215232#1:build-log.txt%3A2570)

* _2026-07-23 12:38:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18491/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2080239156256051200#1:build-log.txt%3A2662)

* _2026-07-23 11:51:19 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18410/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080253331493621760#1:build-log.txt%3A3558)

* _2026-07-23 06:25:58 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18568/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2080177439815766016#1:build-log.txt%3A4360)

* _2026-07-23 03:35:15 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18453/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080107377016705024#1:build-log.txt%3A5485)

* _2026-07-22 19:01:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17946/pull-kubevirt-e2e-k8s-1.34-sig-compute/2079925529477124096#1:build-log.txt%3A3703)

* _2026-07-22 18:11:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18550/pull-kubevirt-e2e-k8s-1.35-sig-compute/2079887455724507136#1:build-log.txt%3A2629)

* _2026-07-22 15:50:33 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18530/pull-kubevirt-e2e-k8s-1.34-sig-network/2079956261721346048#1:build-log.txt%3A3523)

* _2026-07-22 12:05:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18550/pull-kubevirt-e2e-k8s-1.34-sig-network/2079887454780788736#1:build-log.txt%3A5057)

* _2026-07-21 17:49:35 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18436/pull-kubevirt-e2e-k8s-1.35-sig-network/2079623487349067776#1:build-log.txt%3A2546)

* _2026-07-21 15:43:51 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18301/pull-kubevirt-e2e-k8s-1.35-sig-compute/2079588132977643520#1:build-log.txt%3A2628)

* _2026-07-21 15:28:48 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18301/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2079588132075868160#1:build-log.txt%3A5343)

* _2026-07-21 15:25:37 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18402/pull-kubevirt-e2e-k8s-1.35-sig-network/2079554687773708288#1:build-log.txt%3A2550)

* _2026-07-21 14:23:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.34-sig-storage/2079548939106783232#1:build-log.txt%3A3656)

* _2026-07-21 14:17:49 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2079548938674769920#1:build-log.txt%3A5093)

* _2026-07-21 13:10:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18538/pull-kubevirt-e2e-k8s-1.35-sig-storage/2079550396535148544#1:build-log.txt%3A2572)

* _2026-07-21 11:38:35 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18402/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079528731050774528#1:build-log.txt%3A3794)

* _2026-07-21 10:34:47 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.35-sig-network/2079503746533953536#1:build-log.txt%3A4030)

* _2026-07-21 10:33:48 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.34-sig-compute/2079503746223575040#1:build-log.txt%3A2631)

* _2026-07-21 10:31:31 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.34-sig-network/2079503745925779456#1:build-log.txt%3A2541)

* _2026-07-21 07:32:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18499/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2079469516609294336#1:build-log.txt%3A4059)

* _2026-07-21 06:14:47 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18504/pull-kubevirt-e2e-k8s-1.34-sig-storage/2079449848859332608#1:build-log.txt%3A3651)

* _2026-07-20 22:17:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079296912455897088#1:build-log.txt%3A2656)

* _2026-07-20 19:53:48 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079250202144804864#1:build-log.txt%3A2654)

* _2026-07-20 19:21:15 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.36-sig-performance/2079283962835898368#1:build-log.txt%3A5498)

* _2026-07-20 19:18:10 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.36-sig-compute/2079283981123063808#1:build-log.txt%3A2665)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache blob fetch failure (from secondary snippet) (44x / 34.11%) </summary>

<hr/>

**44x**: _2026-07-25 07:26:15 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-compute/2080916717827002368#1:build-log.txt%3A434)
<details>
<summary>all...</summary>

* _2026-07-25 07:51:26 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-windows2016/2080923721920417792#1:build-log.txt%3A415)

* _2026-07-25 07:28:43 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080916717994774528#1:build-log.txt%3A384)

* _2026-07-25 07:26:15 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-compute/2080916717827002368#1:build-log.txt%3A434)

* _2026-07-25 07:23:37 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-storage/2080916717680201728#1:build-log.txt%3A438)

* _2026-07-25 07:23:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-operator/2080916717386600448#1:build-log.txt%3A449)

* _2026-07-25 07:23:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080916717092999168#1:build-log.txt%3A449)

* _2026-07-25 07:23:35 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080916716325441536#1:build-log.txt%3A460)

* _2026-07-25 07:23:34 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080916716623237120#1:build-log.txt%3A457)

* _2026-07-25 07:23:34 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-network/2080916716866506752#1:build-log.txt%3A367)

* _2026-07-25 07:23:34 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-network/2080916717504040960#1:build-log.txt%3A432)

* _2026-07-25 07:23:30 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080916717239799808#1:build-log.txt%3A433)

* _2026-07-25 03:48:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-windows2016/2080862637561745408#1:build-log.txt%3A447)

* _2026-07-25 02:20:43 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-network/2080840488520257536#1:build-log.txt%3A444)

* _2026-07-25 02:20:42 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080840488079855616#1:build-log.txt%3A438)

* _2026-07-25 02:20:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080840487677202432#1:build-log.txt%3A429)

* _2026-07-25 02:20:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-storage/2080840488658669568#1:build-log.txt%3A388)

* _2026-07-25 02:20:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-network/2080840487949832192#1:build-log.txt%3A425)

* _2026-07-25 02:20:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-operator/2080840488360873984#1:build-log.txt%3A389)

* _2026-07-25 02:20:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-compute/2080840488834830336#1:build-log.txt%3A402)

* _2026-07-25 02:20:40 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080840488197296128#1:build-log.txt%3A372)

* _2026-07-25 02:20:40 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080840487815614464#1:build-log.txt%3A364)

* _2026-07-25 02:20:39 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080840488973242368#1:build-log.txt%3A392)

* _2026-07-24 23:32:17 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-windows2016/2080797902535397376#1:build-log.txt%3A476)

* _2026-07-24 22:46:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080764931073904640#1:build-log.txt%3A373)

* _2026-07-24 22:45:46 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-storage/2080764930146963456#1:build-log.txt%3A406)

* _2026-07-24 22:45:44 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-compute/2080764930268598272#1:build-log.txt%3A388)

* _2026-07-24 22:44:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-network/2080764930012745728#1:build-log.txt%3A406)

* _2026-07-24 22:43:29 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-operator/2080764929849167872#1:build-log.txt%3A375)

* _2026-07-24 22:42:00 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080764929639452672#1:build-log.txt%3A386)

* _2026-07-24 22:41:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080764929391988736#1:build-log.txt%3A394)

* _2026-07-24 22:39:57 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-network/2080764929308102656#1:build-log.txt%3A405)

* _2026-07-24 22:38:34 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080764929115164672#1:build-log.txt%3A399)

* _2026-07-24 22:35:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080764928976752640#1:build-log.txt%3A382)

* _2026-07-24 18:13:39 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080711464485654528#1:build-log.txt%3A455)

* _2026-07-24 18:12:50 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-compute/2080711464292716544#1:build-log.txt%3A434)

* _2026-07-24 17:59:31 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-storage/2080711464133332992#1:build-log.txt%3A375)

* _2026-07-24 17:59:18 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-network/2080711463994920960#1:build-log.txt%3A397)

* _2026-07-24 17:59:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-operator/2080711463919423488#1:build-log.txt%3A453)

* _2026-07-24 17:56:22 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080711463818760192#1:build-log.txt%3A385)

* _2026-07-24 17:56:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080711463592267776#1:build-log.txt%3A441)

* _2026-07-24 17:55:57 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-network/2080711463470632960#1:build-log.txt%3A451)

* _2026-07-24 17:52:43 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080711463294472192#1:build-log.txt%3A397)

* _2026-07-24 17:52:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080711463156060160#1:build-log.txt%3A447)

* _2026-07-24 17:48:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-windows2016/2080711462820515840#1:build-log.txt%3A395)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (2x / 1.55%) </summary>

<hr/>

**2x**: _2026-07-26 17:10:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18469/pull-kubevirt-e2e-k8s-1.35-sig-network/2081426767381270528#1:build-log.txt%3A2180)
<details>
<summary>all...</summary>

* _2026-07-26 17:10:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18469/pull-kubevirt-e2e-k8s-1.35-sig-network/2081426767381270528#1:build-log.txt%3A2180)

* _2026-07-24 16:07:20 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18400/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080685210604670976#1:build-log.txt%3A1193)

</details>

<hr/>
</details>

#### needs-investigation (1x / 0.78%)

<details>
<summary> no error snippets (1x / 0.78%) </summary>

<hr/>

**1x**: _2026-07-22 13:32:04 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18168/pull-kubevirt-e2e-kind-1.36-sev/2079922111794647040)
<details>
<summary>all...</summary>

* _2026-07-22 13:32:04 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18168/pull-kubevirt-e2e-kind-1.36-sev/2079922111794647040)

</details>

<hr/>
</details>

#### internal (31x / 24.03%)

<details>
<summary> make cluster lifecycle target failure (31x / 24.03%) </summary>

<hr/>

**31x**: _2026-07-22 19:08:17 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17946/pull-kubevirt-e2e-k8s-1.35-sig-compute/2079925529972051968#1:build-log.txt%3A836)
<details>
<summary>all...</summary>

* _2026-07-26 17:10:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18469/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2081426766697598976#1:build-log.txt%3A1131)

* _2026-07-26 10:45:41 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18474/pull-kubevirt-e2e-k8s-1.34-windows2016/2081329964916936704#1:build-log.txt%3A2120)

* _2026-07-25 07:23:09 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18585/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080916614278025216#1:build-log.txt%3A789)

* _2026-07-25 00:13:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18533/pull-kubevirt-e2e-k8s-1.36-sig-network/2080805647405813760#1:build-log.txt%3A2212)

* _2026-07-24 13:16:39 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18443/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080642898814046208#1:build-log.txt%3A1929)

* _2026-07-24 13:09:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18428/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080570620432617472#1:build-log.txt%3A1966)

* _2026-07-24 07:51:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18542/pull-kubevirt-e2e-k8s-1.35-sig-network/2080561358188843008#1:build-log.txt%3A841)

* _2026-07-24 07:51:32 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18542/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080561368284532736#1:build-log.txt%3A1855)

* _2026-07-23 21:52:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18400/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080410663418597376#1:build-log.txt%3A1825)

* _2026-07-23 16:11:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18568/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080233043007639552#1:build-log.txt%3A1943)

* _2026-07-23 11:52:45 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18414/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080250895450247168#1:build-log.txt%3A832)

* _2026-07-23 10:43:51 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18296/pull-kubevirt-e2e-k8s-1.34-sig-network/2080239905874644992#1:build-log.txt%3A834)

* _2026-07-23 10:35:26 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18296/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080239905669124096#1:build-log.txt%3A836)

* _2026-07-22 19:08:17 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17946/pull-kubevirt-e2e-k8s-1.35-sig-compute/2079925529972051968#1:build-log.txt%3A836)

* _2026-07-22 19:07:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17946/pull-kubevirt-e2e-k8s-1.35-sig-storage/2079925529808474112#1:build-log.txt%3A835)

* _2026-07-22 16:07:50 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18530/pull-kubevirt-e2e-k8s-1.34-windows2016/2079961254075568128#1:build-log.txt%3A832)

* _2026-07-22 15:51:21 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18530/pull-kubevirt-e2e-k8s-1.34-sig-storage/2079956261884923904#1:build-log.txt%3A1902)

* _2026-07-22 12:33:47 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18530/pull-kubevirt-e2e-kind-1.36-sev/2079907633388589056#1:build-log.txt%3A648)

* _2026-07-22 11:52:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18453/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2079896117457195008#1:build-log.txt%3A866)

* _2026-07-22 11:24:21 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18550/pull-kubevirt-e2e-k8s-1.36-sig-compute/2079884525898305536#1:build-log.txt%3A2062)

* _2026-07-22 10:36:23 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18436/pull-kubevirt-e2e-k8s-1.36-sig-network/2079877999724859392#1:build-log.txt%3A1911)

* _2026-07-21 17:45:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18436/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2079623486266937344#1:build-log.txt%3A1895)

* _2026-07-21 15:53:11 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.34-sig-operator/2079593296690155520#1:build-log.txt%3A834)

* _2026-07-21 14:26:15 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.34-sig-operator/2079548939396190208#1:build-log.txt%3A835)

* _2026-07-21 12:59:49 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18538/pull-kubevirt-e2e-k8s-1.34-sig-network/2079550395427852288#1:build-log.txt%3A835)

* _2026-07-21 11:37:21 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18402/pull-kubevirt-e2e-k8s-1.36-sig-network/2079528728521609216#1:build-log.txt%3A1836)

* _2026-07-21 07:32:57 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18499/pull-kubevirt-e2e-k8s-1.34-sig-network/2079469517045501952#1:build-log.txt%3A839)

* _2026-07-21 07:32:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18499/pull-kubevirt-e2e-k8s-1.35-sig-network/2079469518014386176#1:build-log.txt%3A841)

* _2026-07-20 22:21:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18499/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079305097958920192#1:build-log.txt%3A2044)

* _2026-07-20 19:42:17 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18499/pull-kubevirt-e2e-k8s-1.36-sig-storage/2079289433311416320#1:build-log.txt%3A910)

* _2026-07-20 19:18:06 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.36-sig-storage/2079283980238065664#1:build-log.txt%3A849)

</details>

<hr/>
</details>

### release-1.9 (2x / 1.52%)


#### internal (1x / 50.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-07-22 13:28:23 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18565/pull-kubevirt-e2e-kind-1.34-sev-1.9/2079921377246187520#1:build-log.txt%3A812)
<details>
<summary>all...</summary>

* _2026-07-22 13:28:23 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18565/pull-kubevirt-e2e-kind-1.34-sev-1.9/2079921377246187520#1:build-log.txt%3A812)

</details>

<hr/>
</details>

#### external (1x / 50.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-07-22 12:18:49 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18552/pull-kubevirt-e2e-kind-1.34-sev-1.9/2079903851376283648#1:build-log.txt%3A5353)
<details>
<summary>all...</summary>

* _2026-07-22 12:18:49 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18552/pull-kubevirt-e2e-kind-1.34-sev-1.9/2079903851376283648#1:build-log.txt%3A5353)

</details>

<hr/>
</details>

### release-1.7 (1x / 0.76%)


#### internal (1x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-07-22 12:27:27 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18562/pull-kubevirt-e2e-kind-1.34-sev-1.7/2079905794379223040#1:build-log.txt%3A655)
<details>
<summary>all...</summary>

* _2026-07-22 12:27:27 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18562/pull-kubevirt-e2e-kind-1.34-sev-1.7/2079905794379223040#1:build-log.txt%3A655)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-storage (21x / 15.91%)


#### external (15x / 71.43%)

<details>
<summary> bazel remote cache blob fetch failure (from secondary snippet) (8x / 38.10%) </summary>

<hr/>

**8x**: _2026-07-25 07:23:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080916717092999168#1:build-log.txt%3A449)
<details>
<summary>all...</summary>

* _2026-07-25 07:23:37 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-storage/2080916717680201728#1:build-log.txt%3A438)

* _2026-07-25 07:23:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080916717092999168#1:build-log.txt%3A449)

* _2026-07-25 02:20:42 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080840488079855616#1:build-log.txt%3A438)

* _2026-07-25 02:20:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-storage/2080840488658669568#1:build-log.txt%3A388)

* _2026-07-24 22:45:46 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-storage/2080764930146963456#1:build-log.txt%3A406)

* _2026-07-24 22:41:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080764929391988736#1:build-log.txt%3A394)

* _2026-07-24 17:59:31 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-storage/2080711464133332992#1:build-log.txt%3A375)

* _2026-07-24 17:56:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080711463592267776#1:build-log.txt%3A441)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (7x / 33.33%) </summary>

<hr/>

**7x**: _2026-07-26 10:45:39 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18474/pull-kubevirt-e2e-k8s-1.34-sig-storage/2081329965713854464#1:build-log.txt%3A2866)
<details>
<summary>all...</summary>

* _2026-07-26 10:45:39 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18474/pull-kubevirt-e2e-k8s-1.34-sig-storage/2081329965713854464#1:build-log.txt%3A2866)

* _2026-07-24 14:17:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18400/pull-kubevirt-e2e-k8s-1.35-sig-storage/2080658114465501184#1:build-log.txt%3A2570)

* _2026-07-24 14:16:53 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18168/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080651063739813888#1:build-log.txt%3A2651)

* _2026-07-23 15:40:52 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18428/pull-kubevirt-e2e-k8s-1.36-sig-storage/2080316875354411008#1:build-log.txt%3A3593)

* _2026-07-21 14:23:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.34-sig-storage/2079548939106783232#1:build-log.txt%3A3656)

* _2026-07-21 13:10:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18538/pull-kubevirt-e2e-k8s-1.35-sig-storage/2079550396535148544#1:build-log.txt%3A2572)

* _2026-07-21 06:14:47 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18504/pull-kubevirt-e2e-k8s-1.34-sig-storage/2079449848859332608#1:build-log.txt%3A3651)

</details>

<hr/>
</details>

#### internal (6x / 28.57%)

<details>
<summary> make cluster lifecycle target failure (6x / 28.57%) </summary>

<hr/>

**6x**: _2026-07-24 13:09:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18428/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080570620432617472#1:build-log.txt%3A1966)
<details>
<summary>all...</summary>

* _2026-07-24 13:09:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18428/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080570620432617472#1:build-log.txt%3A1966)

* _2026-07-23 11:52:45 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18414/pull-kubevirt-e2e-k8s-1.34-sig-storage/2080250895450247168#1:build-log.txt%3A832)

* _2026-07-22 19:07:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17946/pull-kubevirt-e2e-k8s-1.35-sig-storage/2079925529808474112#1:build-log.txt%3A835)

* _2026-07-22 15:51:21 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18530/pull-kubevirt-e2e-k8s-1.34-sig-storage/2079956261884923904#1:build-log.txt%3A1902)

* _2026-07-20 19:42:17 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18499/pull-kubevirt-e2e-k8s-1.36-sig-storage/2079289433311416320#1:build-log.txt%3A910)

* _2026-07-20 19:18:06 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.36-sig-storage/2079283980238065664#1:build-log.txt%3A849)

</details>

<hr/>
</details>

### sig-network (53x / 40.15%)


#### external (40x / 75.47%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (23x / 43.40%) </summary>

<hr/>

**23x**: _2026-07-26 06:44:01 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18569/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2081269171055759360#1:build-log.txt%3A4441)
<details>
<summary>all...</summary>

* _2026-07-26 10:45:42 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18474/pull-kubevirt-e2e-k8s-1.35-sig-network/2081329966217170944#1:build-log.txt%3A4396)

* _2026-07-26 06:44:01 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18569/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2081269171055759360#1:build-log.txt%3A4441)

* _2026-07-25 07:23:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18585/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080916635056607232#1:build-log.txt%3A5259)

* _2026-07-25 07:23:10 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18585/pull-kubevirt-e2e-k8s-1.36-sig-network/2080916624981889024#1:build-log.txt%3A3651)

* _2026-07-25 00:28:33 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18533/pull-kubevirt-e2e-k8s-1.35-sig-network/2080805657438588928#1:build-log.txt%3A4385)

* _2026-07-24 16:23:44 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18428/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080570620118044672#1:build-log.txt%3A3565)

* _2026-07-24 14:15:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18168/pull-kubevirt-e2e-k8s-1.34-sig-network/2080651063618179072#1:build-log.txt%3A4014)

* _2026-07-24 13:21:54 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18443/pull-kubevirt-e2e-k8s-1.35-sig-network/2080642899699044352#1:build-log.txt%3A5078)

* _2026-07-24 13:16:21 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18443/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080642898696605696#1:build-log.txt%3A3556)

* _2026-07-24 13:09:32 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18428/pull-kubevirt-e2e-k8s-1.34-sig-network/2080570620331954176#1:build-log.txt%3A4014)

* _2026-07-24 07:51:32 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18542/pull-kubevirt-e2e-k8s-1.36-sig-network/2080561377302286336#1:build-log.txt%3A5074)

* _2026-07-23 18:32:12 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18400/pull-kubevirt-e2e-k8s-1.36-sig-network/2080359448429203456#1:build-log.txt%3A5050)

* _2026-07-23 14:26:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18296/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080297545166229504#1:build-log.txt%3A4045)

* _2026-07-23 13:25:38 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18296/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080267705973215232#1:build-log.txt%3A2570)

* _2026-07-22 15:50:33 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18530/pull-kubevirt-e2e-k8s-1.34-sig-network/2079956261721346048#1:build-log.txt%3A3523)

* _2026-07-22 12:05:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18550/pull-kubevirt-e2e-k8s-1.34-sig-network/2079887454780788736#1:build-log.txt%3A5057)

* _2026-07-21 17:49:35 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18436/pull-kubevirt-e2e-k8s-1.35-sig-network/2079623487349067776#1:build-log.txt%3A2546)

* _2026-07-21 15:28:48 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18301/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2079588132075868160#1:build-log.txt%3A5343)

* _2026-07-21 15:25:37 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18402/pull-kubevirt-e2e-k8s-1.35-sig-network/2079554687773708288#1:build-log.txt%3A2550)

* _2026-07-21 14:17:49 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2079548938674769920#1:build-log.txt%3A5093)

* _2026-07-21 10:34:47 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.35-sig-network/2079503746533953536#1:build-log.txt%3A4030)

* _2026-07-21 10:31:31 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.34-sig-network/2079503745925779456#1:build-log.txt%3A2541)

* _2026-07-21 07:32:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18499/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2079469516609294336#1:build-log.txt%3A4059)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache blob fetch failure (from secondary snippet) (16x / 30.19%) </summary>

<hr/>

**16x**: _2026-07-25 07:23:35 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080916716325441536#1:build-log.txt%3A460)
<details>
<summary>all...</summary>

* _2026-07-25 07:23:35 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080916716325441536#1:build-log.txt%3A460)

* _2026-07-25 07:23:34 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-network/2080916717504040960#1:build-log.txt%3A432)

* _2026-07-25 07:23:34 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080916716623237120#1:build-log.txt%3A457)

* _2026-07-25 07:23:34 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-network/2080916716866506752#1:build-log.txt%3A367)

* _2026-07-25 02:20:43 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-network/2080840488520257536#1:build-log.txt%3A444)

* _2026-07-25 02:20:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080840487677202432#1:build-log.txt%3A429)

* _2026-07-25 02:20:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-network/2080840487949832192#1:build-log.txt%3A425)

* _2026-07-25 02:20:40 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080840487815614464#1:build-log.txt%3A364)

* _2026-07-24 22:44:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-network/2080764930012745728#1:build-log.txt%3A406)

* _2026-07-24 22:39:57 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-network/2080764929308102656#1:build-log.txt%3A405)

* _2026-07-24 22:38:34 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080764929115164672#1:build-log.txt%3A399)

* _2026-07-24 22:35:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080764928976752640#1:build-log.txt%3A382)

* _2026-07-24 17:59:18 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-network/2080711463994920960#1:build-log.txt%3A397)

* _2026-07-24 17:55:57 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-network/2080711463470632960#1:build-log.txt%3A451)

* _2026-07-24 17:52:43 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080711463294472192#1:build-log.txt%3A397)

* _2026-07-24 17:52:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080711463156060160#1:build-log.txt%3A447)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 1.89%) </summary>

<hr/>

**1x**: _2026-07-26 17:10:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18469/pull-kubevirt-e2e-k8s-1.35-sig-network/2081426767381270528#1:build-log.txt%3A2180)
<details>
<summary>all...</summary>

* _2026-07-26 17:10:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18469/pull-kubevirt-e2e-k8s-1.35-sig-network/2081426767381270528#1:build-log.txt%3A2180)

</details>

<hr/>
</details>

#### internal (13x / 24.53%)

<details>
<summary> make cluster lifecycle target failure (13x / 24.53%) </summary>

<hr/>

**13x**: _2026-07-23 10:35:26 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18296/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080239905669124096#1:build-log.txt%3A836)
<details>
<summary>all...</summary>

* _2026-07-26 17:10:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18469/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2081426766697598976#1:build-log.txt%3A1131)

* _2026-07-25 00:13:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18533/pull-kubevirt-e2e-k8s-1.36-sig-network/2080805647405813760#1:build-log.txt%3A2212)

* _2026-07-24 13:16:39 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18443/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2080642898814046208#1:build-log.txt%3A1929)

* _2026-07-24 07:51:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18542/pull-kubevirt-e2e-k8s-1.35-sig-network/2080561358188843008#1:build-log.txt%3A841)

* _2026-07-23 10:43:51 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18296/pull-kubevirt-e2e-k8s-1.34-sig-network/2080239905874644992#1:build-log.txt%3A834)

* _2026-07-23 10:35:26 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18296/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2080239905669124096#1:build-log.txt%3A836)

* _2026-07-22 11:52:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18453/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2079896117457195008#1:build-log.txt%3A866)

* _2026-07-22 10:36:23 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18436/pull-kubevirt-e2e-k8s-1.36-sig-network/2079877999724859392#1:build-log.txt%3A1911)

* _2026-07-21 17:45:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18436/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2079623486266937344#1:build-log.txt%3A1895)

* _2026-07-21 12:59:49 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18538/pull-kubevirt-e2e-k8s-1.34-sig-network/2079550395427852288#1:build-log.txt%3A835)

* _2026-07-21 11:37:21 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18402/pull-kubevirt-e2e-k8s-1.36-sig-network/2079528728521609216#1:build-log.txt%3A1836)

* _2026-07-21 07:32:57 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18499/pull-kubevirt-e2e-k8s-1.34-sig-network/2079469517045501952#1:build-log.txt%3A839)

* _2026-07-21 07:32:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18499/pull-kubevirt-e2e-k8s-1.35-sig-network/2079469518014386176#1:build-log.txt%3A841)

</details>

<hr/>
</details>

### sig-compute (52x / 39.39%)


#### external (40x / 76.92%)

<details>
<summary> bazel remote cache blob fetch failure (from secondary snippet) (20x / 38.46%) </summary>

<hr/>

**20x**: _2026-07-25 07:26:15 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-compute/2080916717827002368#1:build-log.txt%3A434)
<details>
<summary>all...</summary>

* _2026-07-25 07:51:26 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-windows2016/2080923721920417792#1:build-log.txt%3A415)

* _2026-07-25 07:28:43 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080916717994774528#1:build-log.txt%3A384)

* _2026-07-25 07:26:15 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-compute/2080916717827002368#1:build-log.txt%3A434)

* _2026-07-25 07:23:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-operator/2080916717386600448#1:build-log.txt%3A449)

* _2026-07-25 07:23:30 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080916717239799808#1:build-log.txt%3A433)

* _2026-07-25 03:48:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-windows2016/2080862637561745408#1:build-log.txt%3A447)

* _2026-07-25 02:20:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-compute/2080840488834830336#1:build-log.txt%3A402)

* _2026-07-25 02:20:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-operator/2080840488360873984#1:build-log.txt%3A389)

* _2026-07-25 02:20:40 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080840488197296128#1:build-log.txt%3A372)

* _2026-07-25 02:20:39 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080840488973242368#1:build-log.txt%3A392)

* _2026-07-24 23:32:17 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-windows2016/2080797902535397376#1:build-log.txt%3A476)

* _2026-07-24 22:46:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080764931073904640#1:build-log.txt%3A373)

* _2026-07-24 22:45:44 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-compute/2080764930268598272#1:build-log.txt%3A388)

* _2026-07-24 22:43:29 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-operator/2080764929849167872#1:build-log.txt%3A375)

* _2026-07-24 22:42:00 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080764929639452672#1:build-log.txt%3A386)

* _2026-07-24 18:13:39 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080711464485654528#1:build-log.txt%3A455)

* _2026-07-24 18:12:50 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.35-sig-compute/2080711464292716544#1:build-log.txt%3A434)

* _2026-07-24 17:59:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-operator/2080711463919423488#1:build-log.txt%3A453)

* _2026-07-24 17:56:22 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080711463818760192#1:build-log.txt%3A385)

* _2026-07-24 17:48:14 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18223/pull-kubevirt-e2e-k8s-1.34-windows2016/2080711462820515840#1:build-log.txt%3A395)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (20x / 38.46%) </summary>

<hr/>

**20x**: _2026-07-22 18:11:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18550/pull-kubevirt-e2e-k8s-1.35-sig-compute/2079887455724507136#1:build-log.txt%3A2629)
<details>
<summary>all...</summary>

* _2026-07-26 19:16:47 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18469/pull-kubevirt-e2e-kind-1.36-sev/2081458617176821760#1:build-log.txt%3A1171)

* _2026-07-26 17:10:22 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18469/pull-kubevirt-e2e-k8s-1.34-sig-compute/2081426767049920512#1:build-log.txt%3A2934)

* _2026-07-26 06:44:05 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18569/pull-kubevirt-e2e-k8s-1.34-sig-compute/2081269171529715712#1:build-log.txt%3A2957)

* _2026-07-25 00:35:31 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18533/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080805660005502976#1:build-log.txt%3A4006)

* _2026-07-25 00:25:55 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18533/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080805657153376256#1:build-log.txt%3A2974)

* _2026-07-24 13:11:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18428/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080570620541669376#1:build-log.txt%3A3677)

* _2026-07-24 07:51:34 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18542/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080561357886853120#1:build-log.txt%3A2659)

* _2026-07-24 07:51:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18542/pull-kubevirt-e2e-k8s-1.34-sig-operator/2080561358025265152#1:build-log.txt%3A2572)

* _2026-07-23 12:38:04 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18491/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2080239156256051200#1:build-log.txt%3A2662)

* _2026-07-23 11:51:19 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18410/pull-kubevirt-e2e-k8s-1.35-sig-operator/2080253331493621760#1:build-log.txt%3A3558)

* _2026-07-23 06:25:58 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18568/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2080177439815766016#1:build-log.txt%3A4360)

* _2026-07-22 19:01:16 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17946/pull-kubevirt-e2e-k8s-1.34-sig-compute/2079925529477124096#1:build-log.txt%3A3703)

* _2026-07-22 18:11:27 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18550/pull-kubevirt-e2e-k8s-1.35-sig-compute/2079887455724507136#1:build-log.txt%3A2629)

* _2026-07-22 12:18:49 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18552/pull-kubevirt-e2e-kind-1.34-sev-1.9/2079903851376283648#1:build-log.txt%3A5353)

* _2026-07-21 15:43:51 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18301/pull-kubevirt-e2e-k8s-1.35-sig-compute/2079588132977643520#1:build-log.txt%3A2628)

* _2026-07-21 11:38:35 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18402/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079528731050774528#1:build-log.txt%3A3794)

* _2026-07-21 10:33:48 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.34-sig-compute/2079503746223575040#1:build-log.txt%3A2631)

* _2026-07-20 22:17:20 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079296912455897088#1:build-log.txt%3A2656)

* _2026-07-20 19:53:48 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079250202144804864#1:build-log.txt%3A2654)

* _2026-07-20 19:18:10 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.36-sig-compute/2079283981123063808#1:build-log.txt%3A2665)

</details>

<hr/>
</details>

#### internal (11x / 21.15%)

<details>
<summary> make cluster lifecycle target failure (11x / 21.15%) </summary>

<hr/>

**11x**: _2026-07-22 19:08:17 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17946/pull-kubevirt-e2e-k8s-1.35-sig-compute/2079925529972051968#1:build-log.txt%3A836)
<details>
<summary>all...</summary>

* _2026-07-26 10:45:41 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18474/pull-kubevirt-e2e-k8s-1.34-windows2016/2081329964916936704#1:build-log.txt%3A2120)

* _2026-07-23 16:11:41 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18568/pull-kubevirt-e2e-k8s-1.34-sig-compute/2080233043007639552#1:build-log.txt%3A1943)

* _2026-07-22 19:08:17 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17946/pull-kubevirt-e2e-k8s-1.35-sig-compute/2079925529972051968#1:build-log.txt%3A836)

* _2026-07-22 16:07:50 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18530/pull-kubevirt-e2e-k8s-1.34-windows2016/2079961254075568128#1:build-log.txt%3A832)

* _2026-07-22 13:28:23 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18565/pull-kubevirt-e2e-kind-1.34-sev-1.9/2079921377246187520#1:build-log.txt%3A812)

* _2026-07-22 12:33:47 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18530/pull-kubevirt-e2e-kind-1.36-sev/2079907633388589056#1:build-log.txt%3A648)

* _2026-07-22 12:27:27 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18562/pull-kubevirt-e2e-kind-1.34-sev-1.7/2079905794379223040#1:build-log.txt%3A655)

* _2026-07-22 11:24:21 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18550/pull-kubevirt-e2e-k8s-1.36-sig-compute/2079884525898305536#1:build-log.txt%3A2062)

* _2026-07-21 15:53:11 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.34-sig-operator/2079593296690155520#1:build-log.txt%3A834)

* _2026-07-21 14:26:15 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.34-sig-operator/2079548939396190208#1:build-log.txt%3A835)

* _2026-07-20 22:21:28 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18499/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079305097958920192#1:build-log.txt%3A2044)

</details>

<hr/>
</details>

#### needs-investigation (1x / 1.92%)

<details>
<summary> no error snippets (1x / 1.92%) </summary>

<hr/>

**1x**: _2026-07-22 13:32:04 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18168/pull-kubevirt-e2e-kind-1.36-sev/2079922111794647040)
<details>
<summary>all...</summary>

* _2026-07-22 13:32:04 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18168/pull-kubevirt-e2e-kind-1.36-sev/2079922111794647040)

</details>

<hr/>
</details>

### sig-performance (6x / 4.55%)


#### external (3x / 50.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 33.33%) </summary>

<hr/>

**2x**: _2026-07-23 03:35:15 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18453/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080107377016705024#1:build-log.txt%3A5485)
<details>
<summary>all...</summary>

* _2026-07-23 03:35:15 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18453/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080107377016705024#1:build-log.txt%3A5485)

* _2026-07-20 19:21:15 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18529/pull-kubevirt-e2e-k8s-1.36-sig-performance/2079283962835898368#1:build-log.txt%3A5498)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-07-24 16:07:20 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18400/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080685210604670976#1:build-log.txt%3A1193)
<details>
<summary>all...</summary>

* _2026-07-24 16:07:20 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18400/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080685210604670976#1:build-log.txt%3A1193)

</details>

<hr/>
</details>

#### internal (3x / 50.00%)

<details>
<summary> make cluster lifecycle target failure (3x / 50.00%) </summary>

<hr/>

**3x**: _2026-07-25 07:23:09 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18585/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080916614278025216#1:build-log.txt%3A789)
<details>
<summary>all...</summary>

* _2026-07-25 07:23:09 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18585/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080916614278025216#1:build-log.txt%3A789)

* _2026-07-24 07:51:32 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18542/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080561368284532736#1:build-log.txt%3A1855)

* _2026-07-23 21:52:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18400/pull-kubevirt-e2e-k8s-1.36-sig-performance/2080410663418597376#1:build-log.txt%3A1825)

</details>

<hr/>
</details>

Last updated: 2026-07-27 19:01:30
