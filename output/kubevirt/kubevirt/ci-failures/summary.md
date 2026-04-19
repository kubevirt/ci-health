



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-04-17 (2x / 6.25%)


#### external (2x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-04-17 13:33:58 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-kind-sriov/2045131370006581248#1:build-log.txt%3A1394)
<details>
<summary>all...</summary>

* _2026-04-17 13:33:58 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-kind-sriov/2045131370006581248#1:build-log.txt%3A1394)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-04-17 13:25:22 &#43;0000 UTC_: <code>13:37:52: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-k8s-1.34-windows2016/2045131369809448960#1:build-log.txt%3A4706)
<details>
<summary>all...</summary>

* _2026-04-17 13:25:22 &#43;0000 UTC_: <code>13:37:52: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-k8s-1.34-windows2016/2045131369809448960#1:build-log.txt%3A4706)

</details>

<hr/>
</details>

### 2026-04-16 (1x / 3.12%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-04-16 20:05:44 &#43;0000 UTC_: <code>20:18:08: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17437/pull-kubevirt-e2e-k8s-1.35-sig-compute/2044869730921091072#1:build-log.txt%3A4004)
<details>
<summary>all...</summary>

* _2026-04-16 20:05:44 &#43;0000 UTC_: <code>20:18:08: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17437/pull-kubevirt-e2e-k8s-1.35-sig-compute/2044869730921091072#1:build-log.txt%3A4004)

</details>

<hr/>
</details>

### 2026-04-15 (1x / 3.12%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-04-15 14:15:15 &#43;0000 UTC_: <code>14:20:43: ERROR: Analysis of target &#39;//cmd/virt-launcher:virt-launcher-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17469/pull-kubevirt-e2e-kind-sriov/2044411425639632896#1:build-log.txt%3A751)
<details>
<summary>all...</summary>

* _2026-04-15 14:15:15 &#43;0000 UTC_: <code>14:20:43: ERROR: Analysis of target &#39;//cmd/virt-launcher:virt-launcher-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17469/pull-kubevirt-e2e-kind-sriov/2044411425639632896#1:build-log.txt%3A751)

</details>

<hr/>
</details>

### 2026-04-14 (21x / 65.62%)


#### external (17x / 80.95%)

<details>
<summary> container image pull failure in context (8x / 38.10%) </summary>

<hr/>

**5x**: _2026-04-14 22:34:45 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2044182454352220160#1:build-log.txt%3A193)
<details>
<summary>all...</summary>

* _2026-04-14 22:35:28 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.7/2044182619574243328#1:build-log.txt%3A203)

* _2026-04-14 22:35:21 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2044182611097554944#1:build-log.txt%3A210)

* _2026-04-14 22:35:21 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.7/2044182616965386240#1:build-log.txt%3A210)

* _2026-04-14 22:34:45 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2044182454352220160#1:build-log.txt%3A193)

* _2026-04-14 22:34:40 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2044182447226097664#1:build-log.txt%3A197)

</details>

<hr/>

**3x**: _2026-04-14 22:35:35 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.7/2044182620316635136#1:build-log.txt%3A559)
<details>
<summary>all...</summary>

* _2026-04-14 22:36:44 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2044182612783665152#1:build-log.txt%3A542)

* _2026-04-14 22:35:57 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-ipv6-sig-network-1.7/2044182607771471872#1:build-log.txt%3A582)

* _2026-04-14 22:35:35 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.7/2044182620316635136#1:build-log.txt%3A559)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (4x / 19.05%) </summary>

<hr/>

**4x**: _2026-04-14 22:36:02 &#43;0000 UTC_: <code>22:38:56: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17501/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.6/2044182784552996864#1:build-log.txt%3A341)
<details>
<summary>all...</summary>

* _2026-04-14 22:36:44 &#43;0000 UTC_: <code>22:40:07: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17501/pull-kubevirt-e2e-kind-1.30-vgpu-1.6/2044182775736569856#1:build-log.txt%3A256)

* _2026-04-14 22:36:02 &#43;0000 UTC_: <code>22:38:56: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17501/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.6/2044182784552996864#1:build-log.txt%3A341)

* _2026-04-14 22:34:41 &#43;0000 UTC_: <code>22:39:51: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2044182449767845888#1:build-log.txt%3A253)

* _2026-04-14 15:40:05 &#43;0000 UTC_: <code>15:43:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17067/pull-kubevirt-e2e-k8s-1.34-sig-storage/2044078011715686400#1:build-log.txt%3A353)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (3x / 14.29%) </summary>

<hr/>

**1x**: _2026-04-14 22:41:27 &#43;0000 UTC_: <code>22:55:06: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-kind-sriov-1.7/2044182599898763264#1:build-log.txt%3A1401)
<details>
<summary>all...</summary>

* _2026-04-14 22:41:27 &#43;0000 UTC_: <code>22:55:06: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-kind-sriov-1.7/2044182599898763264#1:build-log.txt%3A1401)

</details>

<hr/>

**1x**: _2026-04-14 22:36:18 &#43;0000 UTC_: <code>00:32:50: • [FAILED] [464.111 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations-1.7/2044182608597749760#1:build-log.txt%3A4664)
<details>
<summary>all...</summary>

* _2026-04-14 22:36:18 &#43;0000 UTC_: <code>00:32:50: • [FAILED] [464.111 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations-1.7/2044182608597749760#1:build-log.txt%3A4664)

</details>

<hr/>

**1x**: _2026-04-14 22:34:43 &#43;0000 UTC_: <code>00:33:03:   [FAILED] in [AfterEach] - tests/libkubevirt/kubevirt.go:49 @ 04/15/26 00:33:03.525</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2044182452733218816#1:build-log.txt%3A4923)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:43 &#43;0000 UTC_: <code>00:33:03:   [FAILED] in [AfterEach] - tests/libkubevirt/kubevirt.go:49 @ 04/15/26 00:33:03.525</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2044182452733218816#1:build-log.txt%3A4923)

</details>

<hr/>
</details>
<details>
<summary> download failure from external URL (1x / 4.76%) </summary>

<hr/>

**1x**: _2026-04-14 22:35:23 &#43;0000 UTC_: <code>22:38:35: ERROR: Error computing the main repository mapping: no such package &#39;@bazel_features//&#39;: java.io.IOException: Error downloading [https://github.com/bazel-contrib/bazel_features/releases/download/v1.10.0/bazel_features-v1.10.0.tar.gz] to /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/bazel_features/temp1058507301603734894/bazel_features-v1.10.0.tar.gz: Unknown host: release-assets.githubusercontent.com</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2044182611907055616#1:build-log.txt%3A297)
<details>
<summary>all...</summary>

* _2026-04-14 22:35:23 &#43;0000 UTC_: <code>22:38:35: ERROR: Error computing the main repository mapping: no such package &#39;@bazel_features//&#39;: java.io.IOException: Error downloading [https://github.com/bazel-contrib/bazel_features/releases/download/v1.10.0/bazel_features-v1.10.0.tar.gz] to /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/bazel_features/temp1058507301603734894/bazel_features-v1.10.0.tar.gz: Unknown host: release-assets.githubusercontent.com</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2044182611907055616#1:build-log.txt%3A297)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 4.76%) </summary>

<hr/>

**1x**: _2026-04-14 22:34:50 &#43;0000 UTC_: <code>23:54:10: I0414 19:54:10.637122    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2044182448891236352#1:build-log.txt%3A762)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:50 &#43;0000 UTC_: <code>23:54:10: I0414 19:54:10.637122    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2044182448891236352#1:build-log.txt%3A762)

</details>

<hr/>
</details>

#### needs-investigation (4x / 19.05%)

<details>
<summary> no matching pattern (4x / 19.05%) </summary>

<hr/>

**3x**: _2026-04-14 22:36:38 &#43;0000 UTC_: <code>00:59:09:   {&#34;[namespace kubevirt-test-default4 name testvmi-hz4mf-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 1296b6c2-024d-4270-8b10-f2b0ce4c230d]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T00:59:00.422621Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17501/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.6/2044182791259688960#1:build-log.txt%3A9926)
<details>
<summary>all...</summary>

* _2026-04-14 22:36:38 &#43;0000 UTC_: <code>00:59:09:   {&#34;[namespace kubevirt-test-default4 name testvmi-hz4mf-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 1296b6c2-024d-4270-8b10-f2b0ce4c230d]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T00:59:00.422621Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17501/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.6/2044182791259688960#1:build-log.txt%3A9926)

* _2026-04-14 22:36:13 &#43;0000 UTC_: <code>01:14:31:   {&#34;[namespace kubevirt-test-default1 name testvmi-gnklc-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 7f75faab-a51e-44b0-a355-dc176d472057]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T01:14:22.408215Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17501/pull-kubevirt-e2e-k8s-1.31-sig-storage-1.6/2044182787979743232#1:build-log.txt%3A11483)

* _2026-04-14 22:35:33 &#43;0000 UTC_: <code>00:56:47:   {&#34;[namespace kubevirt-test-default2 name testvmi-l826p-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid c2be602b-f84f-4cf3-8368-3243dbee43ef]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T00:56:38.578042Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2044182618622136320#1:build-log.txt%3A5169)

</details>

<hr/>

**1x**: _2026-04-14 22:34:46 &#43;0000 UTC_: <code>01:48:44:   {&#34;[namespace kubevirt-test-default2 name testvmi-rljs5-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 8017f1d0-8e11-49df-bdbd-fa1d76cff6d8]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-15T01:48:19.933126Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T01:48:19.973140Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2044182453085540352#1:build-log.txt%3A5398)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:46 &#43;0000 UTC_: <code>01:48:44:   {&#34;[namespace kubevirt-test-default2 name testvmi-rljs5-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 8017f1d0-8e11-49df-bdbd-fa1d76cff6d8]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-15T01:48:19.933126Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T01:48:19.973140Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2044182453085540352#1:build-log.txt%3A5398)

</details>

<hr/>
</details>

### 2026-04-12 (7x / 21.88%)


#### needs-investigation (3x / 42.86%)

<details>
<summary> no matching pattern (2x / 28.57%) </summary>

<hr/>

**2x**: _2026-04-12 06:07:02 &#43;0000 UTC_: <code>08:59:54:   {&#34;[namespace kubevirt-test-default3 name testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid dd335e64-0fa8-4b94-8995-6c7eff1ebb59]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:59:30.101369Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:59:30.131590Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2043209123444559872#1:build-log.txt%3A9526)
<details>
<summary>all...</summary>

* _2026-04-12 06:07:04 &#43;0000 UTC_: <code>08:44:24:   {&#34;[namespace kubevirt-test-default2 name testvmi-56m8w-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 28ac500b-21a1-49e4-999e-ac55e4a75146]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:43:59.412605Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:43:59.440925Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2043209130977529856#1:build-log.txt%3A7635)

* _2026-04-12 06:07:02 &#43;0000 UTC_: <code>08:59:54:   {&#34;[namespace kubevirt-test-default3 name testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid dd335e64-0fa8-4b94-8995-6c7eff1ebb59]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:59:30.101369Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:59:30.131590Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2043209123444559872#1:build-log.txt%3A9526)

</details>

<hr/>
</details>
<details>
<summary> no error snippets (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-04-12 06:06:33 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2043208957803106304)
<details>
<summary>all...</summary>

* _2026-04-12 06:06:33 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2043208957803106304)

</details>

<hr/>
</details>

#### external (3x / 42.86%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 28.57%) </summary>

<hr/>

**2x**: _2026-04-12 06:06:47 &#43;0000 UTC_: <code>08:24:22:   {&#34;[namespace kubevirt-test-default3 name testvmi-zzvjh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 6fd3fa08-b22f-457c-a6f2-fa4f38378618]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:24:15.430696Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:24:15.463892Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2043208964530769920#1:build-log.txt%3A5700)
<details>
<summary>all...</summary>

* _2026-04-12 06:07:17 &#43;0000 UTC_: <code>08:02:42:   {&#34;[namespace kubevirt-test-default4 name testvmi-nnfwm-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 64014b11-57a1-461e-b599-71f814ffcd45]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:02:28.034865Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:02:28.074924Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2043209126841946112#1:build-log.txt%3A4930)

* _2026-04-12 06:06:47 &#43;0000 UTC_: <code>08:24:22:   {&#34;[namespace kubevirt-test-default3 name testvmi-zzvjh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 6fd3fa08-b22f-457c-a6f2-fa4f38378618]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:24:15.430696Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:24:15.463892Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2043208964530769920#1:build-log.txt%3A5700)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-04-12 06:06:24 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043208958654550016#1:build-log.txt%3A539)
<details>
<summary>all...</summary>

* _2026-04-12 06:06:24 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043208958654550016#1:build-log.txt%3A539)

</details>

<hr/>
</details>

#### internal (1x / 14.29%)

<details>
<summary> make bazel-build target failure (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-04-12 08:21:13 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043242879446421504#1:build-log.txt%3A4241)
<details>
<summary>all...</summary>

* _2026-04-12 08:21:13 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043242879446421504#1:build-log.txt%3A4241)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (24x / 75.00%)

<details>
<summary> container image pull failure in context (9x / 28.12%) </summary>

<hr/>

**5x**: _2026-04-14 22:34:45 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2044182454352220160#1:build-log.txt%3A193)
<details>
<summary>all...</summary>

* _2026-04-14 22:35:28 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.7/2044182619574243328#1:build-log.txt%3A203)
<details><summary>context</summary>
<pre>hack/dockerized &#34;export BUILD_ARCH= &amp;&amp; DOCKER_PREFIX= DOCKER_TAG= DOCKER_TAG_ALT= IMAGE_PREFIX= IMAGE_PREFIX_ALT=kv- ./hack/multi-arch.sh build-images&#34;
22:39:29: Trying to pull quay.io/kubevirt/builder:2510281240-82d14f4b41...
22:39:47: Error: unable to copy from source docker://quay.io/kubevirt/builder:2510281240-82d14f4b41: initializing source docker://quay.io/kubevirt/builder:2510281240-82d14f4b41: pinging container registry quay.io: Get &#34;https://quay.io/v2/&#34;: dial tcp: lookup quay.io: Temporary failure in name resolution
make: *** [Makefile:37: bazel-build-images] Error 125
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-04-14 22:35:21 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2044182611097554944#1:build-log.txt%3A210)
<details><summary>context</summary>
<pre>hack/dockerized &#34;export BUILD_ARCH= &amp;&amp; DOCKER_PREFIX= DOCKER_TAG= DOCKER_TAG_ALT= IMAGE_PREFIX= IMAGE_PREFIX_ALT=kv- ./hack/multi-arch.sh build-images&#34;
22:38:25: Trying to pull quay.io/kubevirt/builder:2510281240-82d14f4b41...
22:39:00: Error: unable to copy from source docker://quay.io/kubevirt/builder:2510281240-82d14f4b41: copying system image from manifest list: parsing image configuration: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/23/23263cb33d82a2b13c0f0569de99085a960e10dd42950f2d4c67e53c9d26719d?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20260414%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20260414T223842Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=78d8100e5d759de57e5b5223b9989d14988b6d0d83cecee4fb5c2fc2dac2ef83&amp;region=us-east-1&amp;namespace=kubevirt&amp;repo_name=builder&amp;akamai_signature=exp=1776207222~hmac=7b248f119cbc0f0f2a5c76fd65c42e367699cd0a54259039182883d112b2361f&#34;: dial tcp: lookup cdn01.quay.io: Temporary failure in name resolution
make: *** [Makefile:37: bazel-build-images] Error 125
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-04-14 22:35:21 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.7/2044182616965386240#1:build-log.txt%3A210)
<details><summary>context</summary>
<pre>hack/dockerized &#34;export BUILD_ARCH= &amp;&amp; DOCKER_PREFIX= DOCKER_TAG= DOCKER_TAG_ALT= IMAGE_PREFIX= IMAGE_PREFIX_ALT=kv- ./hack/multi-arch.sh build-images&#34;
22:38:35: Trying to pull quay.io/kubevirt/builder:2510281240-82d14f4b41...
22:38:58: Error: unable to copy from source docker://quay.io/kubevirt/builder:2510281240-82d14f4b41: copying system image from manifest list: parsing image configuration: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/23/23263cb33d82a2b13c0f0569de99085a960e10dd42950f2d4c67e53c9d26719d?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20260414%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20260414T223840Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=86aacfca8536e6aeb09431f4bad96f957718ac85594941d14aae81a79f7a6c06&amp;region=us-east-1&amp;namespace=kubevirt&amp;repo_name=builder&amp;akamai_signature=exp=1776207220~hmac=d2dc216902a8f670738b37dc7092bfa9444c57dd3dac2de2c05193ebb508edca&#34;: dial tcp: lookup cdn01.quay.io: Temporary failure in name resolution
make: *** [Makefile:37: bazel-build-images] Error 125
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-04-14 22:34:45 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2044182454352220160#1:build-log.txt%3A193)
<details><summary>context</summary>
<pre>hack/dockerized &#34;export BUILD_ARCH= &amp;&amp; DOCKER_PREFIX= DOCKER_TAG= DOCKER_TAG_ALT= IMAGE_PREFIX= IMAGE_PREFIX_ALT=kv- ./hack/multi-arch.sh build-images&#34;
22:36:54: Trying to pull quay.io/kubevirt/builder:2602201229-4494d1587...
22:37:18: Error: unable to copy from source docker://quay.io/kubevirt/builder:2602201229-4494d1587: copying system image from manifest list: parsing image configuration: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/23/23779b65e496e294d12e915d6b24032b890f5b970a157a8ea5f99807c36da777?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20260414%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20260414T223700Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=a2745b76664f40c7d4fcc7a7231ae0c42fc4cf663aea78b564e3eb814c0942ec&amp;region=us-east-1&amp;namespace=kubevirt&amp;repo_name=builder&amp;akamai_signature=exp=1776207120~hmac=4bdce0f207fe57ec0390fc84e5516f5e4e44dac2136d6c35390f46b7c246d4b6&#34;: dial tcp: lookup cdn01.quay.io: Temporary failure in name resolution
make: *** [Makefile:39: bazel-build-images] Error 125
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-04-14 22:34:40 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2044182447226097664#1:build-log.txt%3A197)
<details><summary>context</summary>
<pre>hack/dockerized &#34;export BUILD_ARCH= &amp;&amp; DOCKER_PREFIX= DOCKER_TAG= DOCKER_TAG_ALT= IMAGE_PREFIX= IMAGE_PREFIX_ALT=kv- ./hack/multi-arch.sh build-images&#34;
22:36:54: Trying to pull quay.io/kubevirt/builder:2602201229-4494d1587...
22:37:13: Error: unable to copy from source docker://quay.io/kubevirt/builder:2602201229-4494d1587: copying system image from manifest list: parsing image configuration: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/23/23779b65e496e294d12e915d6b24032b890f5b970a157a8ea5f99807c36da777?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20260414%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20260414T223655Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=b25f5634b3d92ea6e6328a1ccd0f037f0f191feea75944f75a3c9106624ccf98&amp;region=us-east-1&amp;namespace=kubevirt&amp;repo_name=builder&amp;akamai_signature=exp=1776207115~hmac=7aca92c387892b00c3c08eea0e5c858183db88212366bbe4e601b0c5a45b0d63&#34;: dial tcp: lookup cdn01.quay.io: Temporary failure in name resolution
make: *** [Makefile:39: bazel-build-images] Error 125
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


</details>

<hr/>

**4x**: _2026-04-14 22:35:35 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.7/2044182620316635136#1:build-log.txt%3A559)
<details>
<summary>all...</summary>

* _2026-04-14 22:36:44 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2044182612783665152#1:build-log.txt%3A542)
<details><summary>context</summary>
<pre>22:52:05: Trying to pull quay.io/kubevirtci/gocli:2510241540-9fd40b1c...
22:53:15: Error: unable to copy from source docker://quay.io/kubevirtci/gocli:2510241540-9fd40b1c: initializing source docker://quay.io/kubevirtci/gocli:2510241540-9fd40b1c: pinging container registry quay.io: Get &#34;https://quay.io/v2/&#34;: net/http: TLS handshake timeout
22:53:15: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:159: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-04-14 22:35:57 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-ipv6-sig-network-1.7/2044182607771471872#1:build-log.txt%3A582)
<details><summary>context</summary>
<pre>23:02:06: Copying blob sha256:d256b949f856f7e42232670152385765dc793323d146a5728c5524041022e431
23:02:06: Copying blob sha256:18084e71cbfaebe8260f15f560912a624b8167f904c0fea7913b758991d61c7f
23:03:39: Error: unable to copy from source docker://quay.io/kubevirtci/gocli:2510241540-9fd40b1c: initializing source docker://quay.io/kubevirtci/gocli:2510241540-9fd40b1c: pinging container registry quay.io: Get &#34;https://quay.io/v2/&#34;: net/http: TLS handshake timeout
make: *** [Makefile:162: cluster-down] Error 125
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-04-14 22:35:35 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.7/2044182620316635136#1:build-log.txt%3A559)
<details><summary>context</summary>
<pre>23:16:21: Copying blob sha256:d7947dd0f4d10823605cb7ec555f03d25af0a9435628af6d1e95c8d16814fa46
23:17:47: Error: unable to copy from source docker://quay.io/kubevirtci/gocli:2510241540-9fd40b1c: initializing source docker://quay.io/kubevirtci/gocli:2510241540-9fd40b1c: pinging container registry quay.io: Get &#34;https://quay.io/v2/&#34;: net/http: TLS handshake timeout
23:17:47: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:159: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-04-12 06:06:24 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043208958654550016#1:build-log.txt%3A539)
<details><summary>context</summary>
<pre>06:28:07: Copying blob sha256:d7947dd0f4d10823605cb7ec555f03d25af0a9435628af6d1e95c8d16814fa46
06:29:13: Error: unable to copy from source docker://quay.io/kubevirtci/gocli:2510241540-9fd40b1c: reading blob sha256:18084e71cbfaebe8260f15f560912a624b8167f904c0fea7913b758991d61c7f: Get &#34;https://quay.io/v2/kubevirtci/gocli/blobs/sha256:18084e71cbfaebe8260f15f560912a624b8167f904c0fea7913b758991d61c7f&#34;: net/http: TLS handshake timeout
06:29:13: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:159: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> download failure in context (7x / 21.88%) </summary>

<hr/>

**6x**: _2026-04-16 20:05:44 &#43;0000 UTC_: <code>20:18:08: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17437/pull-kubevirt-e2e-k8s-1.35-sig-compute/2044869730921091072#1:build-log.txt%3A4004)
<details>
<summary>all...</summary>

* _2026-04-17 13:25:22 &#43;0000 UTC_: <code>13:37:52: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-k8s-1.34-windows2016/2045131369809448960#1:build-log.txt%3A4706)
<details><summary>context</summary>
<pre>13:37:52: ERROR: Analysis of target &#39;//:buildifier&#39; failed; build aborted:
13:37:52: INFO: Elapsed time: 18.277s
13:37:52: INFO: 0 processes.
13:37:52: ERROR: Build failed. Not running target
make: *** [Makefile:28: bazel-build-functests] Error 1
&#43; ret=2
&#43; check_for_panics</pre>
</details>


* _2026-04-16 20:05:44 &#43;0000 UTC_: <code>20:18:08: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17437/pull-kubevirt-e2e-k8s-1.35-sig-compute/2044869730921091072#1:build-log.txt%3A4004)
<details><summary>context</summary>
<pre>20:18:08: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
20:18:08: INFO: Elapsed time: 18.375s
20:18:08: INFO: 0 processes.
20:18:08: ERROR: Build failed. Not running target
20:18:10: &#43; rm -f /tmp/kubevirt.deploy.YmnF
make: *** [Makefile:188: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


* _2026-04-14 22:36:44 &#43;0000 UTC_: <code>22:40:07: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17501/pull-kubevirt-e2e-kind-1.30-vgpu-1.6/2044182775736569856#1:build-log.txt%3A256)
<details><summary>context</summary>
<pre>22:40:07: ERROR: no such package &#39;@bazel_skylib//lib&#39;: java.io.IOException: Error downloading [https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.1.1/bazel-skylib-1.1.1.tar.gz, https://github.com/bazelbuild/bazel-skylib/releases/download/1.1.1/bazel-skylib-1.1.1.tar.gz] to /tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/bazel_skylib/temp11324301706757646990/bazel-skylib-1.1.1.tar.gz: Unknown host: release-assets.githubusercontent.com
22:40:07: INFO: Elapsed time: 36.395s
22:40:07: INFO: 0 processes.
22:40:07: ERROR: Build failed. Not running target
make: *** [Makefile:40: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>


* _2026-04-14 22:36:02 &#43;0000 UTC_: <code>22:38:56: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17501/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.6/2044182784552996864#1:build-log.txt%3A341)
<details><summary>context</summary>
<pre>22:38:56: ERROR: no such package &#39;@bazel_skylib//lib&#39;: java.io.IOException: Error downloading [https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.1.1/bazel-skylib-1.1.1.tar.gz, https://github.com/bazelbuild/bazel-skylib/releases/download/1.1.1/bazel-skylib-1.1.1.tar.gz] to /tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/bazel_skylib/temp9287357796072782434/bazel-skylib-1.1.1.tar.gz: Unknown host: github.com
22:38:56: INFO: Elapsed time: 29.270s
22:38:56: INFO: 0 processes.
22:38:56: ERROR: Build failed. Not running target
make: *** [Makefile:40: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>


* _2026-04-14 22:34:41 &#43;0000 UTC_: <code>22:39:51: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2044182449767845888#1:build-log.txt%3A253)
<details><summary>context</summary>
<pre>22:39:51: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted:
22:39:51: INFO: Elapsed time: 155.381s
22:39:51: INFO: 0 processes.
22:39:51: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>


* _2026-04-14 15:40:05 &#43;0000 UTC_: <code>15:43:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17067/pull-kubevirt-e2e-k8s-1.34-sig-storage/2044078011715686400#1:build-log.txt%3A353)
<details><summary>context</summary>
<pre>15:43:18: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted:
15:43:18: INFO: Elapsed time: 35.522s
15:43:18: INFO: 0 processes.
15:43:18: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>


</details>

<hr/>

**1x**: _2026-04-15 14:15:15 &#43;0000 UTC_: <code>14:20:43: ERROR: Analysis of target &#39;//cmd/virt-launcher:virt-launcher-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17469/pull-kubevirt-e2e-kind-sriov/2044411425639632896#1:build-log.txt%3A751)
<details>
<summary>all...</summary>

* _2026-04-15 14:15:15 &#43;0000 UTC_: <code>14:20:43: ERROR: Analysis of target &#39;//cmd/virt-launcher:virt-launcher-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17469/pull-kubevirt-e2e-kind-sriov/2044411425639632896#1:build-log.txt%3A751)
<details><summary>context</summary>
<pre>14:20:43: 		rctx.download(
14:20:43: Error in download: java.io.IOException: Error downloading [https://github.com/regclient/regclient/releases/download/v0.8.0/regctl-linux-amd64] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/oci_regctl_linux_amd64/regctl: Unknown host: release-assets.githubusercontent.com
14:20:43: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-launcher/BUILD.bazel:254:10: //cmd/virt-launcher:virt-launcher-image depends on @oci_regctl_linux_amd64//:regctl_toolchain in repository @oci_regctl_linux_amd64 which failed to fetch. no such package &#39;@oci_regctl_linux_amd64//&#39;: java.io.IOException: Error downloading [https://github.com/regclient/regclient/releases/download/v0.8.0/regctl-linux-amd64] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/oci_regctl_linux_amd64/regctl: Unknown host: release-assets.githubusercontent.com
14:20:43: ERROR: Analysis of target &#39;//cmd/virt-launcher:virt-launcher-image&#39; failed; build aborted: Analysis failed
14:20:43: INFO: Elapsed time: 19.255s
14:20:43: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (6x / 18.75%) </summary>

<hr/>

**2x**: _2026-04-12 06:06:47 &#43;0000 UTC_: <code>08:24:22:   {&#34;[namespace kubevirt-test-default3 name testvmi-zzvjh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 6fd3fa08-b22f-457c-a6f2-fa4f38378618]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:24:15.430696Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:24:15.463892Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2043208964530769920#1:build-log.txt%3A5700)
<details>
<summary>all...</summary>

* _2026-04-12 06:07:17 &#43;0000 UTC_: <code>08:02:42:   {&#34;[namespace kubevirt-test-default4 name testvmi-nnfwm-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 64014b11-57a1-461e-b599-71f814ffcd45]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:02:28.034865Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:02:28.074924Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2043209126841946112#1:build-log.txt%3A4930)
<details><summary>context</summary>
<pre>08:02:42:
08:02:42:   Captured StdOut/StdErr Output &gt;&gt;
08:02:42:   {&#34;[namespace kubevirt-test-default4 name testvmi-nnfwm-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 64014b11-57a1-461e-b599-71f814ffcd45]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;Event(v1.ObjectReference{Kind:\&#34;VirtualMachineInstance\&#34;, Namespace:\&#34;kubevirt-test-default4\&#34;, Name:\&#34;testvmi-nnfwm-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\&#34;, UID:\&#34;64014b11-57a1-461e-b599-71f814ffcd45\&#34;, APIVersion:\&#34;kubevirt.io/v1\&#34;, ResourceVersion:\&#34;88940\&#34;, FieldPath:\&#34;\&#34;}): type: &#39;Normal&#39; reason: &#39;SuccessfulCreate&#39; Created virtual machine pod virt-launcher-testvmi-nnfwm-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxt5llq&#34;,&#34;pos&#34;:&#34;watcher.go:150&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:02:24.063043Z&#34;}
08:02:42:   {&#34;[namespace kubevirt-test-default4 name testvmi-nnfwm-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 64014b11-57a1-461e-b599-71f814ffcd45]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:02:28.034865Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:02:28.074924Z&#34;}
08:02:42:   &lt;&lt; Captured StdOut/StdErr Output
08:02:42: ------------------------------
08:03:21: • [39.841 seconds]</pre>
</details>


* _2026-04-12 06:06:47 &#43;0000 UTC_: <code>08:24:22:   {&#34;[namespace kubevirt-test-default3 name testvmi-zzvjh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 6fd3fa08-b22f-457c-a6f2-fa4f38378618]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:24:15.430696Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:24:15.463892Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2043208964530769920#1:build-log.txt%3A5700)
<details><summary>context</summary>
<pre>08:24:22:
08:24:22:   Captured StdOut/StdErr Output &gt;&gt;
08:24:22:   {&#34;[namespace kubevirt-test-default3 name testvmi-zzvjh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 6fd3fa08-b22f-457c-a6f2-fa4f38378618]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;Event(v1.ObjectReference{Kind:\&#34;VirtualMachineInstance\&#34;, Namespace:\&#34;kubevirt-test-default3\&#34;, Name:\&#34;testvmi-zzvjh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\&#34;, UID:\&#34;6fd3fa08-b22f-457c-a6f2-fa4f38378618\&#34;, APIVersion:\&#34;kubevirt.io/v1\&#34;, ResourceVersion:\&#34;129061\&#34;, FieldPath:\&#34;\&#34;}): type: &#39;Normal&#39; reason: &#39;SuccessfulCreate&#39; Created virtual machine pod virt-launcher-testvmi-zzvjh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxg4hjw&#34;,&#34;pos&#34;:&#34;watcher.go:150&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:24:11.189611Z&#34;}
08:24:22:   {&#34;[namespace kubevirt-test-default3 name testvmi-zzvjh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 6fd3fa08-b22f-457c-a6f2-fa4f38378618]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:24:15.430696Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:24:15.463892Z&#34;}
08:24:22:   &lt;&lt; Captured StdOut/StdErr Output
08:24:22: ------------------------------
08:24:24: • [64.167 seconds]</pre>
</details>


</details>

<hr/>

**1x**: _2026-04-17 13:33:58 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-kind-sriov/2045131370006581248#1:build-log.txt%3A1394)
<details>
<summary>all...</summary>

* _2026-04-17 13:33:58 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-kind-sriov/2045131370006581248#1:build-log.txt%3A1394)
<details><summary>context</summary>
<pre>13:42:41: Downloading cni-plugins-linux-amd64-v0.8.5.tgz
13:42:41: &#43;&#43; curl -sSL -o /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-sriov/cni-plugins-linux-amd64-v0.8.5.tgz https://github.com/containernetworking/plugins/releases/download/v0.8.5/cni-plugins-linux-amd64-v0.8.5.tgz
13:42:47: curl: (6) Could not resolve host: github.com
make: *** [Makefile:173: cluster-up] Error 6
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>

<hr/>

**1x**: _2026-04-14 22:41:27 &#43;0000 UTC_: <code>22:55:06: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-kind-sriov-1.7/2044182599898763264#1:build-log.txt%3A1401)
<details>
<summary>all...</summary>

* _2026-04-14 22:41:27 &#43;0000 UTC_: <code>22:55:06: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-kind-sriov-1.7/2044182599898763264#1:build-log.txt%3A1401)
<details><summary>context</summary>
<pre>22:55:06: &#43;&#43; for node in ${master_nodes[@]}
22:55:06: &#43;&#43; _kubectl taint nodes sriov-control-plane node-role.kubernetes.io/master:NoSchedule-
22:55:06: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-sriov/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-sriov/.kubeconfig taint nodes sriov-control-plane node-role.kubernetes.io/master:NoSchedule-
22:55:06: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found
22:55:06: &#43;&#43; _kubectl taint nodes sriov-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
22:55:06: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-sriov/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-sriov/.kubeconfig taint nodes sriov-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
22:55:07: node/sriov-control-plane untainted</pre>
</details>


</details>

<hr/>

**1x**: _2026-04-14 22:36:18 &#43;0000 UTC_: <code>00:32:50: • [FAILED] [464.111 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations-1.7/2044182608597749760#1:build-log.txt%3A4664)
<details>
<summary>all...</summary>

* _2026-04-14 22:36:18 &#43;0000 UTC_: <code>00:32:50: • [FAILED] [464.111 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations-1.7/2044182608597749760#1:build-log.txt%3A4664)
<details><summary>context</summary>
<pre>00:21:10: ------------------------------
00:32:50: •SSSSSSSSSSSSSSSSSSSSSSSS•SSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSS•
00:32:50: ------------------------------
00:32:50: • [FAILED] [464.111 seconds]
00:32:50: [rfe_id:393][crit:high][vendor:cnv-qe@redhat.com][level:system][sig-compute] Live Migration with a live-migrate eviction strategy set [ref_id:2293] with a VMI running with an eviction strategy set  with node tainted during node drain [It] [test_id:2221] should migrate a VMI under load to another node [sig-compute, sig-compute-migrations, requires-two-schedulable-nodes, Serial]
00:32:50: tests/migration/eviction_strategy.go:280
00:32:50:</pre>
</details>


</details>

<hr/>

**1x**: _2026-04-14 22:34:43 &#43;0000 UTC_: <code>00:33:03:   [FAILED] in [AfterEach] - tests/libkubevirt/kubevirt.go:49 @ 04/15/26 00:33:03.525</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2044182452733218816#1:build-log.txt%3A4923)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:43 &#43;0000 UTC_: <code>00:33:03:   [FAILED] in [AfterEach] - tests/libkubevirt/kubevirt.go:49 @ 04/15/26 00:33:03.525</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2044182452733218816#1:build-log.txt%3A4923)
<details><summary>context</summary>
<pre>00:33:03:   On failure, artifacts will be collected in /logs/artifacts/k8s-reporter/2/2_*
00:33:03:   STEP: Collecting Logs due to failure @ 04/15/26 00:32:29.733
00:33:03:   Global test cleanup started.
00:33:03:   [FAILED] in [AfterEach] - tests/libkubevirt/kubevirt.go:49 @ 04/15/26 00:33:03.525
00:33:03:   &lt;&lt; Timeline
00:33:03:
00:33:03:   [PANICKED] Test Panicked</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> download failure from external URL (1x / 3.12%) </summary>

<hr/>

**1x**: _2026-04-14 22:35:23 &#43;0000 UTC_: <code>22:38:35: ERROR: Error computing the main repository mapping: no such package &#39;@bazel_features//&#39;: java.io.IOException: Error downloading [https://github.com/bazel-contrib/bazel_features/releases/download/v1.10.0/bazel_features-v1.10.0.tar.gz] to /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/bazel_features/temp1058507301603734894/bazel_features-v1.10.0.tar.gz: Unknown host: release-assets.githubusercontent.com</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2044182611907055616#1:build-log.txt%3A297)
<details>
<summary>all...</summary>

* _2026-04-14 22:35:23 &#43;0000 UTC_: <code>22:38:35: ERROR: Error computing the main repository mapping: no such package &#39;@bazel_features//&#39;: java.io.IOException: Error downloading [https://github.com/bazel-contrib/bazel_features/releases/download/v1.10.0/bazel_features-v1.10.0.tar.gz] to /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/bazel_features/temp1058507301603734894/bazel_features-v1.10.0.tar.gz: Unknown host: release-assets.githubusercontent.com</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2044182611907055616#1:build-log.txt%3A297)
<details><summary>context</summary>
<pre>22:38:35:   /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/utils.bzl:233:18: in maybe
22:38:35: Repository rule http_archive defined at:
22:38:35:   /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:372:31: in &lt;toplevel&gt;
22:38:35: ERROR: Error computing the main repository mapping: no such package &#39;@bazel_features//&#39;: java.io.IOException: Error downloading [https://github.com/bazel-contrib/bazel_features/releases/download/v1.10.0/bazel_features-v1.10.0.tar.gz] to /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/bazel_features/temp1058507301603734894/bazel_features-v1.10.0.tar.gz: Unknown host: release-assets.githubusercontent.com
make: *** [Makefile:37: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 3.12%) </summary>

<hr/>

**1x**: _2026-04-14 22:34:50 &#43;0000 UTC_: <code>23:54:10: I0414 19:54:10.637122    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2044182448891236352#1:build-log.txt%3A762)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:50 &#43;0000 UTC_: <code>23:54:10: I0414 19:54:10.637122    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2044182448891236352#1:build-log.txt%3A762)
<details><summary>context</summary>
<pre>23:54:05: [control-plane-check] Checking kube-scheduler at https://127.0.0.1:10259/livez
23:54:07: [control-plane-check] kube-controller-manager is healthy after 2.757175666s
23:54:09: [control-plane-check] kube-scheduler is healthy after 4.198605757s
23:54:10: I0414 19:54:10.637122    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
23:54:11: [control-plane-check] kube-apiserver is healthy after 6.00277331s
23:54:11: I0414 19:54:11.140884    1615 kubeconfig.go:657] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists
23:54:11: I0414 19:54:11.143116    1615 kubeconfig.go:730] creating the ClusterRoleBinding for the kubeadm:cluster-admins Group by using super-admin.conf</pre>
</details>


</details>

<hr/>
</details>

### needs-investigation (7x / 21.88%)

<details>
<summary> no matching pattern (6x / 18.75%) </summary>

<hr/>

**3x**: _2026-04-14 22:36:38 &#43;0000 UTC_: <code>00:59:09:   {&#34;[namespace kubevirt-test-default4 name testvmi-hz4mf-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 1296b6c2-024d-4270-8b10-f2b0ce4c230d]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T00:59:00.422621Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17501/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.6/2044182791259688960#1:build-log.txt%3A9926)
<details>
<summary>all...</summary>

* _2026-04-14 22:36:38 &#43;0000 UTC_: <code>00:59:09:   {&#34;[namespace kubevirt-test-default4 name testvmi-hz4mf-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 1296b6c2-024d-4270-8b10-f2b0ce4c230d]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T00:59:00.422621Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17501/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.6/2044182791259688960#1:build-log.txt%3A9926)
<details><summary>context</summary>
<pre>00:59:09:   {&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;Sent: \&#34;mount /dev/vdc /mount_dir\\n\&#34;&#34;,&#34;pos&#34;:&#34;stdlib.go:105&#34;,&#34;timestamp&#34;:&#34;2026-04-15T00:58:58.303623Z&#34;,&#34;ts&#34;:&#34;2026/04/15 00:58:58&#34;}
00:59:09:   {&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;read closing down: EOF&#34;,&#34;pos&#34;:&#34;stdlib.go:105&#34;,&#34;timestamp&#34;:&#34;2026-04-15T00:59:00.304740Z&#34;,&#34;ts&#34;:&#34;2026/04/15 00:59:00&#34;}
00:59:09:   {&#34;[namespace kubevirt-test-default4 name testvmi-hz4mf-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 1296b6c2-024d-4270-8b10-f2b0ce4c230d]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;VirtualMachineInstance soft rebooted&#34;,&#34;pos&#34;:&#34;watcher.go:159&#34;,&#34;timestamp&#34;:&#34;2026-04-15T00:59:00.342085Z&#34;}
00:59:09:   {&#34;[namespace kubevirt-test-default4 name testvmi-hz4mf-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 1296b6c2-024d-4270-8b10-f2b0ce4c230d]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T00:59:00.422621Z&#34;}
00:59:09:   &lt;&lt; Captured StdOut/StdErr Output
00:59:09: ------------------------------
00:59:14: • [144.754 seconds]</pre>
</details>


* _2026-04-14 22:36:13 &#43;0000 UTC_: <code>01:14:31:   {&#34;[namespace kubevirt-test-default1 name testvmi-gnklc-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 7f75faab-a51e-44b0-a355-dc176d472057]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T01:14:22.408215Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17501/pull-kubevirt-e2e-k8s-1.31-sig-storage-1.6/2044182787979743232#1:build-log.txt%3A11483)
<details><summary>context</summary>
<pre>01:14:31:   {&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;Sent: \&#34;mount /dev/vdc /mount_dir\\n\&#34;&#34;,&#34;pos&#34;:&#34;stdlib.go:105&#34;,&#34;timestamp&#34;:&#34;2026-04-15T01:14:22.261235Z&#34;,&#34;ts&#34;:&#34;2026/04/15 01:14:22&#34;}
01:14:31:   {&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;read closing down: EOF&#34;,&#34;pos&#34;:&#34;stdlib.go:105&#34;,&#34;timestamp&#34;:&#34;2026-04-15T01:14:22.311607Z&#34;,&#34;ts&#34;:&#34;2026/04/15 01:14:22&#34;}
01:14:31:   {&#34;[namespace kubevirt-test-default1 name testvmi-gnklc-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 7f75faab-a51e-44b0-a355-dc176d472057]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;VirtualMachineInstance soft rebooted&#34;,&#34;pos&#34;:&#34;watcher.go:159&#34;,&#34;timestamp&#34;:&#34;2026-04-15T01:14:22.356404Z&#34;}
01:14:31:   {&#34;[namespace kubevirt-test-default1 name testvmi-gnklc-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 7f75faab-a51e-44b0-a355-dc176d472057]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T01:14:22.408215Z&#34;}
01:14:31:   &lt;&lt; Captured StdOut/StdErr Output
01:14:31: ------------------------------
01:14:49: • [144.586 seconds]</pre>
</details>


* _2026-04-14 22:35:33 &#43;0000 UTC_: <code>00:56:47:   {&#34;[namespace kubevirt-test-default2 name testvmi-l826p-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid c2be602b-f84f-4cf3-8368-3243dbee43ef]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T00:56:38.578042Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2044182618622136320#1:build-log.txt%3A5169)
<details><summary>context</summary>
<pre>00:56:47:   {&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;Sent: \&#34;mount /dev/vdc /mount_dir\\n\&#34;&#34;,&#34;pos&#34;:&#34;stdlib.go:105&#34;,&#34;timestamp&#34;:&#34;2026-04-15T00:56:38.306120Z&#34;,&#34;ts&#34;:&#34;2026/04/15 00:56:38&#34;}
00:56:47:   {&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;read closing down: EOF&#34;,&#34;pos&#34;:&#34;stdlib.go:105&#34;,&#34;timestamp&#34;:&#34;2026-04-15T00:56:38.429493Z&#34;,&#34;ts&#34;:&#34;2026/04/15 00:56:38&#34;}
00:56:47:   {&#34;[namespace kubevirt-test-default2 name testvmi-l826p-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid c2be602b-f84f-4cf3-8368-3243dbee43ef]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;VirtualMachineInstance soft rebooted&#34;,&#34;pos&#34;:&#34;watcher.go:159&#34;,&#34;timestamp&#34;:&#34;2026-04-15T00:56:38.478041Z&#34;}
00:56:47:   {&#34;[namespace kubevirt-test-default2 name testvmi-l826p-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid c2be602b-f84f-4cf3-8368-3243dbee43ef]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T00:56:38.578042Z&#34;}
00:56:47:   &lt;&lt; Captured StdOut/StdErr Output
00:56:47: ------------------------------
00:57:46: •••••</pre>
</details>


</details>

<hr/>

**3x**: _2026-04-12 06:07:02 &#43;0000 UTC_: <code>08:59:54:   {&#34;[namespace kubevirt-test-default3 name testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid dd335e64-0fa8-4b94-8995-6c7eff1ebb59]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:59:30.101369Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:59:30.131590Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2043209123444559872#1:build-log.txt%3A9526)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:46 &#43;0000 UTC_: <code>01:48:44:   {&#34;[namespace kubevirt-test-default2 name testvmi-rljs5-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 8017f1d0-8e11-49df-bdbd-fa1d76cff6d8]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-15T01:48:19.933126Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T01:48:19.973140Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2044182453085540352#1:build-log.txt%3A5398)
<details><summary>context</summary>
<pre>01:48:44:
01:48:44:   Captured StdOut/StdErr Output &gt;&gt;
01:48:44:   {&#34;[namespace kubevirt-test-default2 name testvmi-rljs5-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 8017f1d0-8e11-49df-bdbd-fa1d76cff6d8]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;Event(v1.ObjectReference{Kind:\&#34;VirtualMachineInstance\&#34;, Namespace:\&#34;kubevirt-test-default2\&#34;, Name:\&#34;testvmi-rljs5-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\&#34;, UID:\&#34;8017f1d0-8e11-49df-bdbd-fa1d76cff6d8\&#34;, APIVersion:\&#34;kubevirt.io/v1\&#34;, ResourceVersion:\&#34;157662\&#34;, FieldPath:\&#34;\&#34;}): type: &#39;Normal&#39; reason: &#39;SuccessfulCreate&#39; Created virtual machine pod virt-launcher-testvmi-rljs5-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxf842j&#34;,&#34;pos&#34;:&#34;watcher.go:150&#34;,&#34;timestamp&#34;:&#34;2026-04-15T01:48:16.385235Z&#34;}
01:48:44:   {&#34;[namespace kubevirt-test-default2 name testvmi-rljs5-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 8017f1d0-8e11-49df-bdbd-fa1d76cff6d8]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-15T01:48:19.933126Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T01:48:19.973140Z&#34;}
01:48:44:   &lt;&lt; Captured StdOut/StdErr Output
01:48:44: ------------------------------
01:49:21: • [110.511 seconds]</pre>
</details>


* _2026-04-12 06:07:04 &#43;0000 UTC_: <code>08:44:24:   {&#34;[namespace kubevirt-test-default2 name testvmi-56m8w-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 28ac500b-21a1-49e4-999e-ac55e4a75146]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:43:59.412605Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:43:59.440925Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2043209130977529856#1:build-log.txt%3A7635)
<details><summary>context</summary>
<pre>08:44:24:
08:44:24:   Captured StdOut/StdErr Output &gt;&gt;
08:44:24:   {&#34;[namespace kubevirt-test-default2 name testvmi-56m8w-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 28ac500b-21a1-49e4-999e-ac55e4a75146]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;Event(v1.ObjectReference{Kind:\&#34;VirtualMachineInstance\&#34;, Namespace:\&#34;kubevirt-test-default2\&#34;, Name:\&#34;testvmi-56m8w-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\&#34;, UID:\&#34;28ac500b-21a1-49e4-999e-ac55e4a75146\&#34;, APIVersion:\&#34;kubevirt.io/v1\&#34;, ResourceVersion:\&#34;140603\&#34;, FieldPath:\&#34;\&#34;}): type: &#39;Normal&#39; reason: &#39;SuccessfulCreate&#39; Created virtual machine pod virt-launcher-testvmi-56m8w-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx5nmm7&#34;,&#34;pos&#34;:&#34;watcher.go:150&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:43:55.709663Z&#34;}
08:44:24:   {&#34;[namespace kubevirt-test-default2 name testvmi-56m8w-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 28ac500b-21a1-49e4-999e-ac55e4a75146]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:43:59.412605Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:43:59.440925Z&#34;}
08:44:24:   &lt;&lt; Captured StdOut/StdErr Output
08:44:24: ------------------------------
08:44:56: • [35.870 seconds]</pre>
</details>


* _2026-04-12 06:07:02 &#43;0000 UTC_: <code>08:59:54:   {&#34;[namespace kubevirt-test-default3 name testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid dd335e64-0fa8-4b94-8995-6c7eff1ebb59]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:59:30.101369Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:59:30.131590Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2043209123444559872#1:build-log.txt%3A9526)
<details><summary>context</summary>
<pre>08:59:54:
08:59:54:   Captured StdOut/StdErr Output &gt;&gt;
08:59:54:   {&#34;[namespace kubevirt-test-default3 name testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid dd335e64-0fa8-4b94-8995-6c7eff1ebb59]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;Event(v1.ObjectReference{Kind:\&#34;VirtualMachineInstance\&#34;, Namespace:\&#34;kubevirt-test-default3\&#34;, Name:\&#34;testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\&#34;, UID:\&#34;dd335e64-0fa8-4b94-8995-6c7eff1ebb59\&#34;, APIVersion:\&#34;kubevirt.io/v1\&#34;, ResourceVersion:\&#34;157478\&#34;, FieldPath:\&#34;\&#34;}): type: &#39;Normal&#39; reason: &#39;SuccessfulCreate&#39; Created virtual machine pod virt-launcher-testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxg8l8k&#34;,&#34;pos&#34;:&#34;watcher.go:150&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:59:25.751280Z&#34;}
08:59:54:   {&#34;[namespace kubevirt-test-default3 name testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid dd335e64-0fa8-4b94-8995-6c7eff1ebb59]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:59:30.101369Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:59:30.131590Z&#34;}
08:59:54:   &lt;&lt; Captured StdOut/StdErr Output
08:59:54: ------------------------------
09:00:08: • [82.837 seconds]</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> no error snippets (1x / 3.12%) </summary>

<hr/>

**1x**: _2026-04-12 06:06:33 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2043208957803106304)
<details>
<summary>all...</summary>

* _2026-04-12 06:06:33 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2043208957803106304)

</details>

<hr/>
</details>

### internal (1x / 3.12%)

<details>
<summary> make bazel-build target failure (1x / 3.12%) </summary>

<hr/>

**1x**: _2026-04-12 08:21:13 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043242879446421504#1:build-log.txt%3A4241)
<details>
<summary>all...</summary>

* _2026-04-12 08:21:13 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043242879446421504#1:build-log.txt%3A4241)
<details><summary>context</summary>
<pre>08:45:57: INFO: Elapsed time: 0.440s, Critical Path: 0.15s
08:45:57: INFO: 1 process: 1 internal.
08:45:57: INFO: Running command line: bazel-bin/example-guest-agent-copier /root/go/src/kubevirt.io/kubevirt/_out/cmd/example-guest-agent/example-guest-agent
make: *** [Makefile:26: bazel-build-functests] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (5x / 15.62%)


#### external (5x / 100.00%)

<details>
<summary> download failure in context (4x / 80.00%) </summary>

<hr/>

**3x**: _2026-04-16 20:05:44 &#43;0000 UTC_: <code>20:18:08: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17437/pull-kubevirt-e2e-k8s-1.35-sig-compute/2044869730921091072#1:build-log.txt%3A4004)
<details>
<summary>all...</summary>

* _2026-04-17 13:25:22 &#43;0000 UTC_: <code>13:37:52: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-k8s-1.34-windows2016/2045131369809448960#1:build-log.txt%3A4706)

* _2026-04-16 20:05:44 &#43;0000 UTC_: <code>20:18:08: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17437/pull-kubevirt-e2e-k8s-1.35-sig-compute/2044869730921091072#1:build-log.txt%3A4004)

* _2026-04-14 15:40:05 &#43;0000 UTC_: <code>15:43:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17067/pull-kubevirt-e2e-k8s-1.34-sig-storage/2044078011715686400#1:build-log.txt%3A353)

</details>

<hr/>

**1x**: _2026-04-15 14:15:15 &#43;0000 UTC_: <code>14:20:43: ERROR: Analysis of target &#39;//cmd/virt-launcher:virt-launcher-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17469/pull-kubevirt-e2e-kind-sriov/2044411425639632896#1:build-log.txt%3A751)
<details>
<summary>all...</summary>

* _2026-04-15 14:15:15 &#43;0000 UTC_: <code>14:20:43: ERROR: Analysis of target &#39;//cmd/virt-launcher:virt-launcher-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17469/pull-kubevirt-e2e-kind-sriov/2044411425639632896#1:build-log.txt%3A751)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-04-17 13:33:58 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-kind-sriov/2045131370006581248#1:build-log.txt%3A1394)
<details>
<summary>all...</summary>

* _2026-04-17 13:33:58 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-kind-sriov/2045131370006581248#1:build-log.txt%3A1394)

</details>

<hr/>
</details>

### release-1.7 (14x / 43.75%)


#### external (11x / 78.57%)

<details>
<summary> container image pull failure in context (7x / 50.00%) </summary>

<hr/>

**4x**: _2026-04-14 22:35:35 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.7/2044182620316635136#1:build-log.txt%3A559)
<details>
<summary>all...</summary>

* _2026-04-14 22:36:44 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2044182612783665152#1:build-log.txt%3A542)

* _2026-04-14 22:35:57 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-ipv6-sig-network-1.7/2044182607771471872#1:build-log.txt%3A582)

* _2026-04-14 22:35:35 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.7/2044182620316635136#1:build-log.txt%3A559)

* _2026-04-12 06:06:24 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043208958654550016#1:build-log.txt%3A539)

</details>

<hr/>

**3x**: _2026-04-14 22:35:21 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2044182611097554944#1:build-log.txt%3A210)
<details>
<summary>all...</summary>

* _2026-04-14 22:35:28 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.7/2044182619574243328#1:build-log.txt%3A203)

* _2026-04-14 22:35:21 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2044182611097554944#1:build-log.txt%3A210)

* _2026-04-14 22:35:21 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.7/2044182616965386240#1:build-log.txt%3A210)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (3x / 21.43%) </summary>

<hr/>

**1x**: _2026-04-14 22:41:27 &#43;0000 UTC_: <code>22:55:06: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-kind-sriov-1.7/2044182599898763264#1:build-log.txt%3A1401)
<details>
<summary>all...</summary>

* _2026-04-14 22:41:27 &#43;0000 UTC_: <code>22:55:06: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-kind-sriov-1.7/2044182599898763264#1:build-log.txt%3A1401)

</details>

<hr/>

**1x**: _2026-04-14 22:36:18 &#43;0000 UTC_: <code>00:32:50: • [FAILED] [464.111 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations-1.7/2044182608597749760#1:build-log.txt%3A4664)
<details>
<summary>all...</summary>

* _2026-04-14 22:36:18 &#43;0000 UTC_: <code>00:32:50: • [FAILED] [464.111 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations-1.7/2044182608597749760#1:build-log.txt%3A4664)

</details>

<hr/>

**1x**: _2026-04-12 06:06:47 &#43;0000 UTC_: <code>08:24:22:   {&#34;[namespace kubevirt-test-default3 name testvmi-zzvjh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 6fd3fa08-b22f-457c-a6f2-fa4f38378618]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:24:15.430696Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:24:15.463892Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2043208964530769920#1:build-log.txt%3A5700)
<details>
<summary>all...</summary>

* _2026-04-12 06:06:47 &#43;0000 UTC_: <code>08:24:22:   {&#34;[namespace kubevirt-test-default3 name testvmi-zzvjh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 6fd3fa08-b22f-457c-a6f2-fa4f38378618]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:24:15.430696Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:24:15.463892Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2043208964530769920#1:build-log.txt%3A5700)

</details>

<hr/>
</details>
<details>
<summary> download failure from external URL (1x / 7.14%) </summary>

<hr/>

**1x**: _2026-04-14 22:35:23 &#43;0000 UTC_: <code>22:38:35: ERROR: Error computing the main repository mapping: no such package &#39;@bazel_features//&#39;: java.io.IOException: Error downloading [https://github.com/bazel-contrib/bazel_features/releases/download/v1.10.0/bazel_features-v1.10.0.tar.gz] to /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/bazel_features/temp1058507301603734894/bazel_features-v1.10.0.tar.gz: Unknown host: release-assets.githubusercontent.com</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2044182611907055616#1:build-log.txt%3A297)
<details>
<summary>all...</summary>

* _2026-04-14 22:35:23 &#43;0000 UTC_: <code>22:38:35: ERROR: Error computing the main repository mapping: no such package &#39;@bazel_features//&#39;: java.io.IOException: Error downloading [https://github.com/bazel-contrib/bazel_features/releases/download/v1.10.0/bazel_features-v1.10.0.tar.gz] to /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/bazel_features/temp1058507301603734894/bazel_features-v1.10.0.tar.gz: Unknown host: release-assets.githubusercontent.com</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2044182611907055616#1:build-log.txt%3A297)

</details>

<hr/>
</details>

#### needs-investigation (2x / 14.29%)

<details>
<summary> no error snippets (1x / 7.14%) </summary>

<hr/>

**1x**: _2026-04-12 06:06:33 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2043208957803106304)
<details>
<summary>all...</summary>

* _2026-04-12 06:06:33 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2043208957803106304)

</details>

<hr/>
</details>
<details>
<summary> no matching pattern (1x / 7.14%) </summary>

<hr/>

**1x**: _2026-04-14 22:35:33 &#43;0000 UTC_: <code>00:56:47:   {&#34;[namespace kubevirt-test-default2 name testvmi-l826p-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid c2be602b-f84f-4cf3-8368-3243dbee43ef]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T00:56:38.578042Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2044182618622136320#1:build-log.txt%3A5169)
<details>
<summary>all...</summary>

* _2026-04-14 22:35:33 &#43;0000 UTC_: <code>00:56:47:   {&#34;[namespace kubevirt-test-default2 name testvmi-l826p-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid c2be602b-f84f-4cf3-8368-3243dbee43ef]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T00:56:38.578042Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2044182618622136320#1:build-log.txt%3A5169)

</details>

<hr/>
</details>

#### internal (1x / 7.14%)

<details>
<summary> make bazel-build target failure (1x / 7.14%) </summary>

<hr/>

**1x**: _2026-04-12 08:21:13 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043242879446421504#1:build-log.txt%3A4241)
<details>
<summary>all...</summary>

* _2026-04-12 08:21:13 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043242879446421504#1:build-log.txt%3A4241)

</details>

<hr/>
</details>

### release-1.6 (4x / 12.50%)


#### needs-investigation (2x / 50.00%)

<details>
<summary> no matching pattern (2x / 50.00%) </summary>

<hr/>

**2x**: _2026-04-14 22:36:38 &#43;0000 UTC_: <code>00:59:09:   {&#34;[namespace kubevirt-test-default4 name testvmi-hz4mf-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 1296b6c2-024d-4270-8b10-f2b0ce4c230d]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T00:59:00.422621Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17501/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.6/2044182791259688960#1:build-log.txt%3A9926)
<details>
<summary>all...</summary>

* _2026-04-14 22:36:38 &#43;0000 UTC_: <code>00:59:09:   {&#34;[namespace kubevirt-test-default4 name testvmi-hz4mf-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 1296b6c2-024d-4270-8b10-f2b0ce4c230d]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T00:59:00.422621Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17501/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.6/2044182791259688960#1:build-log.txt%3A9926)

* _2026-04-14 22:36:13 &#43;0000 UTC_: <code>01:14:31:   {&#34;[namespace kubevirt-test-default1 name testvmi-gnklc-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 7f75faab-a51e-44b0-a355-dc176d472057]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T01:14:22.408215Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17501/pull-kubevirt-e2e-k8s-1.31-sig-storage-1.6/2044182787979743232#1:build-log.txt%3A11483)

</details>

<hr/>
</details>

#### external (2x / 50.00%)

<details>
<summary> download failure in context (2x / 50.00%) </summary>

<hr/>

**2x**: _2026-04-14 22:36:02 &#43;0000 UTC_: <code>22:38:56: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17501/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.6/2044182784552996864#1:build-log.txt%3A341)
<details>
<summary>all...</summary>

* _2026-04-14 22:36:44 &#43;0000 UTC_: <code>22:40:07: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17501/pull-kubevirt-e2e-kind-1.30-vgpu-1.6/2044182775736569856#1:build-log.txt%3A256)

* _2026-04-14 22:36:02 &#43;0000 UTC_: <code>22:38:56: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17501/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.6/2044182784552996864#1:build-log.txt%3A341)

</details>

<hr/>
</details>

### release-1.8 (9x / 28.12%)


#### external (6x / 66.67%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 22.22%) </summary>

<hr/>

**1x**: _2026-04-14 22:34:43 &#43;0000 UTC_: <code>00:33:03:   [FAILED] in [AfterEach] - tests/libkubevirt/kubevirt.go:49 @ 04/15/26 00:33:03.525</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2044182452733218816#1:build-log.txt%3A4923)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:43 &#43;0000 UTC_: <code>00:33:03:   [FAILED] in [AfterEach] - tests/libkubevirt/kubevirt.go:49 @ 04/15/26 00:33:03.525</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2044182452733218816#1:build-log.txt%3A4923)

</details>

<hr/>

**1x**: _2026-04-12 06:07:17 &#43;0000 UTC_: <code>08:02:42:   {&#34;[namespace kubevirt-test-default4 name testvmi-nnfwm-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 64014b11-57a1-461e-b599-71f814ffcd45]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:02:28.034865Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:02:28.074924Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2043209126841946112#1:build-log.txt%3A4930)
<details>
<summary>all...</summary>

* _2026-04-12 06:07:17 &#43;0000 UTC_: <code>08:02:42:   {&#34;[namespace kubevirt-test-default4 name testvmi-nnfwm-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 64014b11-57a1-461e-b599-71f814ffcd45]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:02:28.034865Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:02:28.074924Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2043209126841946112#1:build-log.txt%3A4930)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (2x / 22.22%) </summary>

<hr/>

**2x**: _2026-04-14 22:34:45 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2044182454352220160#1:build-log.txt%3A193)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:45 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2044182454352220160#1:build-log.txt%3A193)

* _2026-04-14 22:34:40 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2044182447226097664#1:build-log.txt%3A197)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-04-14 22:34:50 &#43;0000 UTC_: <code>23:54:10: I0414 19:54:10.637122    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2044182448891236352#1:build-log.txt%3A762)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:50 &#43;0000 UTC_: <code>23:54:10: I0414 19:54:10.637122    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2044182448891236352#1:build-log.txt%3A762)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-04-14 22:34:41 &#43;0000 UTC_: <code>22:39:51: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2044182449767845888#1:build-log.txt%3A253)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:41 &#43;0000 UTC_: <code>22:39:51: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2044182449767845888#1:build-log.txt%3A253)

</details>

<hr/>
</details>

#### needs-investigation (3x / 33.33%)

<details>
<summary> no matching pattern (3x / 33.33%) </summary>

<hr/>

**3x**: _2026-04-12 06:07:02 &#43;0000 UTC_: <code>08:59:54:   {&#34;[namespace kubevirt-test-default3 name testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid dd335e64-0fa8-4b94-8995-6c7eff1ebb59]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:59:30.101369Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:59:30.131590Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2043209123444559872#1:build-log.txt%3A9526)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:46 &#43;0000 UTC_: <code>01:48:44:   {&#34;[namespace kubevirt-test-default2 name testvmi-rljs5-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 8017f1d0-8e11-49df-bdbd-fa1d76cff6d8]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-15T01:48:19.933126Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T01:48:19.973140Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2044182453085540352#1:build-log.txt%3A5398)

* _2026-04-12 06:07:04 &#43;0000 UTC_: <code>08:44:24:   {&#34;[namespace kubevirt-test-default2 name testvmi-56m8w-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 28ac500b-21a1-49e4-999e-ac55e4a75146]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:43:59.412605Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:43:59.440925Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2043209130977529856#1:build-log.txt%3A7635)

* _2026-04-12 06:07:02 &#43;0000 UTC_: <code>08:59:54:   {&#34;[namespace kubevirt-test-default3 name testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid dd335e64-0fa8-4b94-8995-6c7eff1ebb59]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:59:30.101369Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:59:30.131590Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2043209123444559872#1:build-log.txt%3A9526)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-network (6x / 18.75%)


#### external (6x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (3x / 50.00%) </summary>

<hr/>

**1x**: _2026-04-17 13:33:58 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-kind-sriov/2045131370006581248#1:build-log.txt%3A1394)
<details>
<summary>all...</summary>

* _2026-04-17 13:33:58 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-kind-sriov/2045131370006581248#1:build-log.txt%3A1394)

</details>

<hr/>

**1x**: _2026-04-14 22:41:27 &#43;0000 UTC_: <code>22:55:06: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-kind-sriov-1.7/2044182599898763264#1:build-log.txt%3A1401)
<details>
<summary>all...</summary>

* _2026-04-14 22:41:27 &#43;0000 UTC_: <code>22:55:06: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-kind-sriov-1.7/2044182599898763264#1:build-log.txt%3A1401)

</details>

<hr/>

**1x**: _2026-04-14 22:34:43 &#43;0000 UTC_: <code>00:33:03:   [FAILED] in [AfterEach] - tests/libkubevirt/kubevirt.go:49 @ 04/15/26 00:33:03.525</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2044182452733218816#1:build-log.txt%3A4923)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:43 &#43;0000 UTC_: <code>00:33:03:   [FAILED] in [AfterEach] - tests/libkubevirt/kubevirt.go:49 @ 04/15/26 00:33:03.525</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2044182452733218816#1:build-log.txt%3A4923)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (2x / 33.33%) </summary>

<hr/>

**1x**: _2026-04-14 22:35:57 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-ipv6-sig-network-1.7/2044182607771471872#1:build-log.txt%3A582)
<details>
<summary>all...</summary>

* _2026-04-14 22:35:57 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-ipv6-sig-network-1.7/2044182607771471872#1:build-log.txt%3A582)

</details>

<hr/>

**1x**: _2026-04-14 22:35:21 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2044182611097554944#1:build-log.txt%3A210)
<details>
<summary>all...</summary>

* _2026-04-14 22:35:21 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2044182611097554944#1:build-log.txt%3A210)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-04-15 14:15:15 &#43;0000 UTC_: <code>14:20:43: ERROR: Analysis of target &#39;//cmd/virt-launcher:virt-launcher-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17469/pull-kubevirt-e2e-kind-sriov/2044411425639632896#1:build-log.txt%3A751)
<details>
<summary>all...</summary>

* _2026-04-15 14:15:15 &#43;0000 UTC_: <code>14:20:43: ERROR: Analysis of target &#39;//cmd/virt-launcher:virt-launcher-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17469/pull-kubevirt-e2e-kind-sriov/2044411425639632896#1:build-log.txt%3A751)

</details>

<hr/>
</details>

### sig-compute (14x / 43.75%)


#### external (13x / 92.86%)

<details>
<summary> container image pull failure in context (7x / 50.00%) </summary>

<hr/>

**4x**: _2026-04-14 22:34:45 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2044182454352220160#1:build-log.txt%3A193)
<details>
<summary>all...</summary>

* _2026-04-14 22:35:28 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.7/2044182619574243328#1:build-log.txt%3A203)

* _2026-04-14 22:35:21 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.7/2044182616965386240#1:build-log.txt%3A210)

* _2026-04-14 22:34:45 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2044182454352220160#1:build-log.txt%3A193)

* _2026-04-14 22:34:40 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2044182447226097664#1:build-log.txt%3A197)

</details>

<hr/>

**3x**: _2026-04-14 22:35:35 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.7/2044182620316635136#1:build-log.txt%3A559)
<details>
<summary>all...</summary>

* _2026-04-14 22:36:44 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2044182612783665152#1:build-log.txt%3A542)

* _2026-04-14 22:35:35 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.7/2044182620316635136#1:build-log.txt%3A559)

* _2026-04-12 06:06:24 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043208958654550016#1:build-log.txt%3A539)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (5x / 35.71%) </summary>

<hr/>

**5x**: _2026-04-16 20:05:44 &#43;0000 UTC_: <code>20:18:08: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17437/pull-kubevirt-e2e-k8s-1.35-sig-compute/2044869730921091072#1:build-log.txt%3A4004)
<details>
<summary>all...</summary>

* _2026-04-17 13:25:22 &#43;0000 UTC_: <code>13:37:52: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17391/pull-kubevirt-e2e-k8s-1.34-windows2016/2045131369809448960#1:build-log.txt%3A4706)

* _2026-04-16 20:05:44 &#43;0000 UTC_: <code>20:18:08: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17437/pull-kubevirt-e2e-k8s-1.35-sig-compute/2044869730921091072#1:build-log.txt%3A4004)

* _2026-04-14 22:36:44 &#43;0000 UTC_: <code>22:40:07: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17501/pull-kubevirt-e2e-kind-1.30-vgpu-1.6/2044182775736569856#1:build-log.txt%3A256)

* _2026-04-14 22:36:02 &#43;0000 UTC_: <code>22:38:56: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17501/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.6/2044182784552996864#1:build-log.txt%3A341)

* _2026-04-14 22:34:41 &#43;0000 UTC_: <code>22:39:51: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2044182449767845888#1:build-log.txt%3A253)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 7.14%) </summary>

<hr/>

**1x**: _2026-04-14 22:36:18 &#43;0000 UTC_: <code>00:32:50: • [FAILED] [464.111 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations-1.7/2044182608597749760#1:build-log.txt%3A4664)
<details>
<summary>all...</summary>

* _2026-04-14 22:36:18 &#43;0000 UTC_: <code>00:32:50: • [FAILED] [464.111 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations-1.7/2044182608597749760#1:build-log.txt%3A4664)

</details>

<hr/>
</details>

#### internal (1x / 7.14%)

<details>
<summary> make bazel-build target failure (1x / 7.14%) </summary>

<hr/>

**1x**: _2026-04-12 08:21:13 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043242879446421504#1:build-log.txt%3A4241)
<details>
<summary>all...</summary>

* _2026-04-12 08:21:13 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043242879446421504#1:build-log.txt%3A4241)

</details>

<hr/>
</details>

### sig-storage (12x / 37.50%)


#### external (5x / 41.67%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 16.67%) </summary>

<hr/>

**2x**: _2026-04-12 06:06:47 &#43;0000 UTC_: <code>08:24:22:   {&#34;[namespace kubevirt-test-default3 name testvmi-zzvjh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 6fd3fa08-b22f-457c-a6f2-fa4f38378618]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:24:15.430696Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:24:15.463892Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2043208964530769920#1:build-log.txt%3A5700)
<details>
<summary>all...</summary>

* _2026-04-12 06:07:17 &#43;0000 UTC_: <code>08:02:42:   {&#34;[namespace kubevirt-test-default4 name testvmi-nnfwm-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 64014b11-57a1-461e-b599-71f814ffcd45]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:02:28.034865Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:02:28.074924Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2043209126841946112#1:build-log.txt%3A4930)

* _2026-04-12 06:06:47 &#43;0000 UTC_: <code>08:24:22:   {&#34;[namespace kubevirt-test-default3 name testvmi-zzvjh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 6fd3fa08-b22f-457c-a6f2-fa4f38378618]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:24:15.430696Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:24:15.463892Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2043208964530769920#1:build-log.txt%3A5700)

</details>

<hr/>
</details>
<details>
<summary> download failure from external URL (1x / 8.33%) </summary>

<hr/>

**1x**: _2026-04-14 22:35:23 &#43;0000 UTC_: <code>22:38:35: ERROR: Error computing the main repository mapping: no such package &#39;@bazel_features//&#39;: java.io.IOException: Error downloading [https://github.com/bazel-contrib/bazel_features/releases/download/v1.10.0/bazel_features-v1.10.0.tar.gz] to /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/bazel_features/temp1058507301603734894/bazel_features-v1.10.0.tar.gz: Unknown host: release-assets.githubusercontent.com</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2044182611907055616#1:build-log.txt%3A297)
<details>
<summary>all...</summary>

* _2026-04-14 22:35:23 &#43;0000 UTC_: <code>22:38:35: ERROR: Error computing the main repository mapping: no such package &#39;@bazel_features//&#39;: java.io.IOException: Error downloading [https://github.com/bazel-contrib/bazel_features/releases/download/v1.10.0/bazel_features-v1.10.0.tar.gz] to /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/bazel_features/temp1058507301603734894/bazel_features-v1.10.0.tar.gz: Unknown host: release-assets.githubusercontent.com</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2044182611907055616#1:build-log.txt%3A297)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 8.33%) </summary>

<hr/>

**1x**: _2026-04-14 22:34:50 &#43;0000 UTC_: <code>23:54:10: I0414 19:54:10.637122    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2044182448891236352#1:build-log.txt%3A762)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:50 &#43;0000 UTC_: <code>23:54:10: I0414 19:54:10.637122    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2044182448891236352#1:build-log.txt%3A762)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 8.33%) </summary>

<hr/>

**1x**: _2026-04-14 15:40:05 &#43;0000 UTC_: <code>15:43:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17067/pull-kubevirt-e2e-k8s-1.34-sig-storage/2044078011715686400#1:build-log.txt%3A353)
<details>
<summary>all...</summary>

* _2026-04-14 15:40:05 &#43;0000 UTC_: <code>15:43:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17067/pull-kubevirt-e2e-k8s-1.34-sig-storage/2044078011715686400#1:build-log.txt%3A353)

</details>

<hr/>
</details>

#### needs-investigation (7x / 58.33%)

<details>
<summary> no matching pattern (6x / 50.00%) </summary>

<hr/>

**3x**: _2026-04-14 22:36:38 &#43;0000 UTC_: <code>00:59:09:   {&#34;[namespace kubevirt-test-default4 name testvmi-hz4mf-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 1296b6c2-024d-4270-8b10-f2b0ce4c230d]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T00:59:00.422621Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17501/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.6/2044182791259688960#1:build-log.txt%3A9926)
<details>
<summary>all...</summary>

* _2026-04-14 22:36:38 &#43;0000 UTC_: <code>00:59:09:   {&#34;[namespace kubevirt-test-default4 name testvmi-hz4mf-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 1296b6c2-024d-4270-8b10-f2b0ce4c230d]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T00:59:00.422621Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17501/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.6/2044182791259688960#1:build-log.txt%3A9926)

* _2026-04-14 22:36:13 &#43;0000 UTC_: <code>01:14:31:   {&#34;[namespace kubevirt-test-default1 name testvmi-gnklc-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 7f75faab-a51e-44b0-a355-dc176d472057]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T01:14:22.408215Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17501/pull-kubevirt-e2e-k8s-1.31-sig-storage-1.6/2044182787979743232#1:build-log.txt%3A11483)

* _2026-04-14 22:35:33 &#43;0000 UTC_: <code>00:56:47:   {&#34;[namespace kubevirt-test-default2 name testvmi-l826p-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid c2be602b-f84f-4cf3-8368-3243dbee43ef]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T00:56:38.578042Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17500/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2044182618622136320#1:build-log.txt%3A5169)

</details>

<hr/>

**3x**: _2026-04-12 06:07:02 &#43;0000 UTC_: <code>08:59:54:   {&#34;[namespace kubevirt-test-default3 name testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid dd335e64-0fa8-4b94-8995-6c7eff1ebb59]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:59:30.101369Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:59:30.131590Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2043209123444559872#1:build-log.txt%3A9526)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:46 &#43;0000 UTC_: <code>01:48:44:   {&#34;[namespace kubevirt-test-default2 name testvmi-rljs5-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 8017f1d0-8e11-49df-bdbd-fa1d76cff6d8]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-15T01:48:19.933126Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T01:48:19.973140Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2044182453085540352#1:build-log.txt%3A5398)

* _2026-04-12 06:07:04 &#43;0000 UTC_: <code>08:44:24:   {&#34;[namespace kubevirt-test-default2 name testvmi-56m8w-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 28ac500b-21a1-49e4-999e-ac55e4a75146]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:43:59.412605Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:43:59.440925Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2043209130977529856#1:build-log.txt%3A7635)

* _2026-04-12 06:07:02 &#43;0000 UTC_: <code>08:59:54:   {&#34;[namespace kubevirt-test-default3 name testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid dd335e64-0fa8-4b94-8995-6c7eff1ebb59]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:59:30.101369Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:59:30.131590Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2043209123444559872#1:build-log.txt%3A9526)

</details>

<hr/>
</details>
<details>
<summary> no error snippets (1x / 8.33%) </summary>

<hr/>

**1x**: _2026-04-12 06:06:33 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2043208957803106304)
<details>
<summary>all...</summary>

* _2026-04-12 06:06:33 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2043208957803106304)

</details>

<hr/>
</details>

Last updated: 2026-04-19 06:38:52
