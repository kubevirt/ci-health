

# CI failures for kubevirt/kubevirt

## per branch (10x)

<details>
<summary> release-1.7 (2x / 20.00%) </summary>

<hr/>

**1x**: _2026-03-24 13:09:20 &#43;0000 UTC_: <code>13:25:22: Error: cannot remove container 2f1cf993a4a9c84d32a41e048f8e0eb6c0008584f22cfe56f9d1133657d2cc05 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17271/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2036429922599178240#1:build-log.txt%3A1559)
<details>
<summary>all...</summary>

* _2026-03-24 13:09:20 &#43;0000 UTC_: <code>13:25:22: Error: cannot remove container 2f1cf993a4a9c84d32a41e048f8e0eb6c0008584f22cfe56f9d1133657d2cc05 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17271/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2036429922599178240#1:build-log.txt%3A1559)

</details>

<hr/>

**1x**: _2026-03-18 23:32:43 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17219/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.7/2034400005564928000#1:build-log.txt%3A1017)
<details>
<summary>all...</summary>

* _2026-03-18 23:32:43 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17219/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.7/2034400005564928000#1:build-log.txt%3A1017)

</details>

<hr/>
</details>
<details>
<summary> main (8x / 80.00%) </summary>

<hr/>

**5x**: _2026-03-19 02:44:17 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15895/pull-kubevirt-e2e-kind-sriov/2034459465838759936#1:build-log.txt%3A1390)
<details>
<summary>all...</summary>

* _2026-03-23 11:28:30 &#43;0000 UTC_: <code>make: *** [Makefile:188: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17250/pull-kubevirt-e2e-kind-1.35-vgpu/2036042272675467264#1:build-log.txt%3A4433)

* _2026-03-22 08:15:00 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 28</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17109/pull-kubevirt-e2e-kind-1.34-sev/2035631211845521408#1:build-log.txt%3A532)

* _2026-03-20 13:04:15 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16885/pull-kubevirt-e2e-k8s-1.33-sig-operator/2034979183943225344#1:build-log.txt%3A601)

* _2026-03-20 13:04:14 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16885/pull-kubevirt-e2e-k8s-1.34-windows2016/2034979182789791744#1:build-log.txt%3A622)

* _2026-03-19 02:44:17 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15895/pull-kubevirt-e2e-kind-sriov/2034459465838759936#1:build-log.txt%3A1390)

</details>

<hr/>

**2x**: _2026-03-24 12:39:39 &#43;0000 UTC_: <code>12:48:17: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17071/pull-kubevirt-e2e-kind-sriov/2036406364904886272#1:build-log.txt%3A1432)
<details>
<summary>all...</summary>

* _2026-03-24 12:39:39 &#43;0000 UTC_: <code>12:48:17: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17071/pull-kubevirt-e2e-kind-sriov/2036406364904886272#1:build-log.txt%3A1432)

* _2026-03-20 13:09:39 &#43;0000 UTC_: <code>13:19:40: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16885/pull-kubevirt-e2e-kind-sriov/2034979182982729728#1:build-log.txt%3A1434)

</details>

<hr/>

**1x**: _2026-03-20 13:04:14 &#43;0000 UTC_: <code>13:15:38: error: taint &#34;node-role.kubernetes.io/control-plane:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16885/pull-kubevirt-e2e-kind-1.35-vgpu/2034979182898843648#1:build-log.txt%3A1124)
<details>
<summary>all...</summary>

* _2026-03-20 13:04:14 &#43;0000 UTC_: <code>13:15:38: error: taint &#34;node-role.kubernetes.io/control-plane:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16885/pull-kubevirt-e2e-kind-1.35-vgpu/2034979182898843648#1:build-log.txt%3A1124)

</details>

<hr/>
</details>

## per day (10x)

<details>
<summary> 2026-03-24 (2x / 20.00%) </summary>

<hr/>

**1x**: _2026-03-24 13:09:20 &#43;0000 UTC_: <code>13:25:22: Error: cannot remove container 2f1cf993a4a9c84d32a41e048f8e0eb6c0008584f22cfe56f9d1133657d2cc05 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17271/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2036429922599178240#1:build-log.txt%3A1559)
<details>
<summary>all...</summary>

* _2026-03-24 13:09:20 &#43;0000 UTC_: <code>13:25:22: Error: cannot remove container 2f1cf993a4a9c84d32a41e048f8e0eb6c0008584f22cfe56f9d1133657d2cc05 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17271/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2036429922599178240#1:build-log.txt%3A1559)

</details>

<hr/>

**1x**: _2026-03-24 12:39:39 &#43;0000 UTC_: <code>12:48:17: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17071/pull-kubevirt-e2e-kind-sriov/2036406364904886272#1:build-log.txt%3A1432)
<details>
<summary>all...</summary>

* _2026-03-24 12:39:39 &#43;0000 UTC_: <code>12:48:17: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17071/pull-kubevirt-e2e-kind-sriov/2036406364904886272#1:build-log.txt%3A1432)

</details>

<hr/>
</details>
<details>
<summary> 2026-03-23 (1x / 10.00%) </summary>

<hr/>

**1x**: _2026-03-23 11:28:30 &#43;0000 UTC_: <code>make: *** [Makefile:188: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17250/pull-kubevirt-e2e-kind-1.35-vgpu/2036042272675467264#1:build-log.txt%3A4433)
<details>
<summary>all...</summary>

* _2026-03-23 11:28:30 &#43;0000 UTC_: <code>make: *** [Makefile:188: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17250/pull-kubevirt-e2e-kind-1.35-vgpu/2036042272675467264#1:build-log.txt%3A4433)

</details>

<hr/>
</details>
<details>
<summary> 2026-03-22 (1x / 10.00%) </summary>

<hr/>

**1x**: _2026-03-22 08:15:00 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 28</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17109/pull-kubevirt-e2e-kind-1.34-sev/2035631211845521408#1:build-log.txt%3A532)
<details>
<summary>all...</summary>

* _2026-03-22 08:15:00 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 28</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17109/pull-kubevirt-e2e-kind-1.34-sev/2035631211845521408#1:build-log.txt%3A532)

</details>

<hr/>
</details>
<details>
<summary> 2026-03-20 (4x / 40.00%) </summary>

<hr/>

**2x**: _2026-03-20 13:09:39 &#43;0000 UTC_: <code>13:19:40: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16885/pull-kubevirt-e2e-kind-sriov/2034979182982729728#1:build-log.txt%3A1434)
<details>
<summary>all...</summary>

* _2026-03-20 13:09:39 &#43;0000 UTC_: <code>13:19:40: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16885/pull-kubevirt-e2e-kind-sriov/2034979182982729728#1:build-log.txt%3A1434)

* _2026-03-20 13:04:14 &#43;0000 UTC_: <code>13:15:38: error: taint &#34;node-role.kubernetes.io/control-plane:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16885/pull-kubevirt-e2e-kind-1.35-vgpu/2034979182898843648#1:build-log.txt%3A1124)

</details>

<hr/>

**2x**: _2026-03-20 13:04:15 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16885/pull-kubevirt-e2e-k8s-1.33-sig-operator/2034979183943225344#1:build-log.txt%3A601)
<details>
<summary>all...</summary>

* _2026-03-20 13:04:15 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16885/pull-kubevirt-e2e-k8s-1.33-sig-operator/2034979183943225344#1:build-log.txt%3A601)

* _2026-03-20 13:04:14 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16885/pull-kubevirt-e2e-k8s-1.34-windows2016/2034979182789791744#1:build-log.txt%3A622)

</details>

<hr/>
</details>
<details>
<summary> 2026-03-19 (1x / 10.00%) </summary>

<hr/>

**1x**: _2026-03-19 02:44:17 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15895/pull-kubevirt-e2e-kind-sriov/2034459465838759936#1:build-log.txt%3A1390)
<details>
<summary>all...</summary>

* _2026-03-19 02:44:17 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15895/pull-kubevirt-e2e-kind-sriov/2034459465838759936#1:build-log.txt%3A1390)

</details>

<hr/>
</details>
<details>
<summary> 2026-03-18 (1x / 10.00%) </summary>

<hr/>

**1x**: _2026-03-18 23:32:43 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17219/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.7/2034400005564928000#1:build-log.txt%3A1017)
<details>
<summary>all...</summary>

* _2026-03-18 23:32:43 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17219/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.7/2034400005564928000#1:build-log.txt%3A1017)

</details>

<hr/>
</details>

## per SIG (10x)

<details>
<summary> sig-compute (7x / 70.00%) </summary>

<hr/>

**1x**: _2026-03-24 13:09:20 &#43;0000 UTC_: <code>13:25:22: Error: cannot remove container 2f1cf993a4a9c84d32a41e048f8e0eb6c0008584f22cfe56f9d1133657d2cc05 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17271/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2036429922599178240#1:build-log.txt%3A1559)
<details>
<summary>all...</summary>

* _2026-03-24 13:09:20 &#43;0000 UTC_: <code>13:25:22: Error: cannot remove container 2f1cf993a4a9c84d32a41e048f8e0eb6c0008584f22cfe56f9d1133657d2cc05 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17271/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2036429922599178240#1:build-log.txt%3A1559)

</details>

<hr/>

**5x**: _2026-03-20 13:04:15 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16885/pull-kubevirt-e2e-k8s-1.33-sig-operator/2034979183943225344#1:build-log.txt%3A601)
<details>
<summary>all...</summary>

* _2026-03-23 11:28:30 &#43;0000 UTC_: <code>make: *** [Makefile:188: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17250/pull-kubevirt-e2e-kind-1.35-vgpu/2036042272675467264#1:build-log.txt%3A4433)

* _2026-03-22 08:15:00 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 28</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17109/pull-kubevirt-e2e-kind-1.34-sev/2035631211845521408#1:build-log.txt%3A532)

* _2026-03-20 13:04:15 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16885/pull-kubevirt-e2e-k8s-1.33-sig-operator/2034979183943225344#1:build-log.txt%3A601)

* _2026-03-20 13:04:14 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16885/pull-kubevirt-e2e-k8s-1.34-windows2016/2034979182789791744#1:build-log.txt%3A622)

* _2026-03-18 23:32:43 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17219/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.7/2034400005564928000#1:build-log.txt%3A1017)

</details>

<hr/>

**1x**: _2026-03-20 13:04:14 &#43;0000 UTC_: <code>13:15:38: error: taint &#34;node-role.kubernetes.io/control-plane:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16885/pull-kubevirt-e2e-kind-1.35-vgpu/2034979182898843648#1:build-log.txt%3A1124)
<details>
<summary>all...</summary>

* _2026-03-20 13:04:14 &#43;0000 UTC_: <code>13:15:38: error: taint &#34;node-role.kubernetes.io/control-plane:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16885/pull-kubevirt-e2e-kind-1.35-vgpu/2034979182898843648#1:build-log.txt%3A1124)

</details>

<hr/>
</details>
<details>
<summary> sig-network (3x / 30.00%) </summary>

<hr/>

**2x**: _2026-03-24 12:39:39 &#43;0000 UTC_: <code>12:48:17: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17071/pull-kubevirt-e2e-kind-sriov/2036406364904886272#1:build-log.txt%3A1432)
<details>
<summary>all...</summary>

* _2026-03-24 12:39:39 &#43;0000 UTC_: <code>12:48:17: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17071/pull-kubevirt-e2e-kind-sriov/2036406364904886272#1:build-log.txt%3A1432)

* _2026-03-20 13:09:39 &#43;0000 UTC_: <code>13:19:40: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16885/pull-kubevirt-e2e-kind-sriov/2034979182982729728#1:build-log.txt%3A1434)

</details>

<hr/>

**1x**: _2026-03-19 02:44:17 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15895/pull-kubevirt-e2e-kind-sriov/2034459465838759936#1:build-log.txt%3A1390)
<details>
<summary>all...</summary>

* _2026-03-19 02:44:17 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15895/pull-kubevirt-e2e-kind-sriov/2034459465838759936#1:build-log.txt%3A1390)

</details>

<hr/>
</details>

Last updated: 2026-03-25 12:17:51
