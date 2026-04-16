



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-04-14 (1x / 7.69%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-04-14 15:40:05 &#43;0000 UTC_: <code>15:43:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17067/pull-kubevirt-e2e-k8s-1.34-sig-storage/2044078011715686400#1:build-log.txt%3A353)
<details>
<summary>all...</summary>

* _2026-04-14 15:40:05 &#43;0000 UTC_: <code>15:43:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17067/pull-kubevirt-e2e-k8s-1.34-sig-storage/2044078011715686400#1:build-log.txt%3A353)

</details>

<hr/>
</details>

### 2026-04-12 (7x / 53.85%)


#### needs-investigation (3x / 42.86%)

<details>
<summary> no matching pattern (2x / 28.57%) </summary>

<hr/>

**2x**: _2026-04-12 06:07:02 &#43;0000 UTC_: <code>08:59:54:   {&#34;[namespace kubevirt-test-default3 name testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid dd335e64-0fa8-4b94-8995-6c7eff1ebb59]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:59:30.101369Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:59:30.131590Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2043209123444559872#1:build-log.txt%3A9526)
<details>
<summary>all...</summary>

* _2026-04-12 06:07:04 &#43;0000 UTC_: <code>08:44:24:   {&#34;[namespace kubevirt-test-default2 name testvmi-56m8w-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 28ac500b-21a1-49e4-999e-ac55e4a75146]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:43:59.412605Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:43:59.440925Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2043209130977529856#1:build-log.txt%3A7635)

* _2026-04-12 06:07:02 &#43;0000 UTC_: <code>08:59:54:   {&#34;[namespace kubevirt-test-default3 name testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid dd335e64-0fa8-4b94-8995-6c7eff1ebb59]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:59:30.101369Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:59:30.131590Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2043209123444559872#1:build-log.txt%3A9526)

</details>

<hr/>
</details>
<details>
<summary> no error snippets (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-04-12 06:06:33 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2043208957803106304)
<details>
<summary>all...</summary>

* _2026-04-12 06:06:33 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2043208957803106304)

</details>

<hr/>
</details>

#### external (3x / 42.86%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 28.57%) </summary>

<hr/>

**2x**: _2026-04-12 06:06:47 &#43;0000 UTC_: <code>08:24:22:   {&#34;[namespace kubevirt-test-default3 name testvmi-zzvjh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 6fd3fa08-b22f-457c-a6f2-fa4f38378618]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:24:15.430696Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:24:15.463892Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2043208964530769920#1:build-log.txt%3A5700)
<details>
<summary>all...</summary>

* _2026-04-12 06:07:17 &#43;0000 UTC_: <code>08:02:42:   {&#34;[namespace kubevirt-test-default4 name testvmi-nnfwm-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 64014b11-57a1-461e-b599-71f814ffcd45]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:02:28.034865Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:02:28.074924Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2043209126841946112#1:build-log.txt%3A4930)

* _2026-04-12 06:06:47 &#43;0000 UTC_: <code>08:24:22:   {&#34;[namespace kubevirt-test-default3 name testvmi-zzvjh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 6fd3fa08-b22f-457c-a6f2-fa4f38378618]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:24:15.430696Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:24:15.463892Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2043208964530769920#1:build-log.txt%3A5700)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-04-12 06:06:24 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043208958654550016#1:build-log.txt%3A539)
<details>
<summary>all...</summary>

* _2026-04-12 06:06:24 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043208958654550016#1:build-log.txt%3A539)

</details>

<hr/>
</details>

#### internal (1x / 14.29%)

<details>
<summary> make bazel-build target failure (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-04-12 08:21:13 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043242879446421504#1:build-log.txt%3A4241)
<details>
<summary>all...</summary>

* _2026-04-12 08:21:13 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043242879446421504#1:build-log.txt%3A4241)

</details>

<hr/>
</details>

### 2026-04-10 (2x / 15.38%)


#### needs-investigation (1x / 50.00%)

<details>
<summary> no matching pattern (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-04-10 06:12:35 &#43;0000 UTC_: <code>08:07:41: • [FAILED] [182.207 seconds]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17445/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2042485694953164800#1:build-log.txt%3A5316)
<details>
<summary>all...</summary>

* _2026-04-10 06:12:35 &#43;0000 UTC_: <code>08:07:41: • [FAILED] [182.207 seconds]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17445/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2042485694953164800#1:build-log.txt%3A5316)

</details>

<hr/>
</details>

#### external (1x / 50.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-04-10 06:11:44 &#43;0000 UTC_: <code>08:07:33: • [FAILED] [175.103 seconds]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17444/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.6/2042485527759818752#1:build-log.txt%3A3381)
<details>
<summary>all...</summary>

* _2026-04-10 06:11:44 &#43;0000 UTC_: <code>08:07:33: • [FAILED] [175.103 seconds]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17444/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.6/2042485527759818752#1:build-log.txt%3A3381)

</details>

<hr/>
</details>

### 2026-04-09 (3x / 23.08%)


#### internal (2x / 66.67%)

<details>
<summary> make cluster lifecycle target failure (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-04-09 01:56:53 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2042058998902951936#1:build-log.txt%3A1040)
<details>
<summary>all...</summary>

* _2026-04-09 01:56:54 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2042059010437287936#1:build-log.txt%3A1142)

* _2026-04-09 01:56:53 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2042058998902951936#1:build-log.txt%3A1040)

</details>

<hr/>
</details>

#### external (1x / 33.33%)

<details>
<summary> download failure in context (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-04-09 14:20:27 &#43;0000 UTC_: <code>14:31:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17340/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2042246124231725056#1:build-log.txt%3A4058)
<details>
<summary>all...</summary>

* _2026-04-09 14:20:27 &#43;0000 UTC_: <code>14:31:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17340/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2042246124231725056#1:build-log.txt%3A4058)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (6x / 46.15%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (3x / 23.08%) </summary>

<hr/>

**2x**: _2026-04-12 06:06:47 &#43;0000 UTC_: <code>08:24:22:   {&#34;[namespace kubevirt-test-default3 name testvmi-zzvjh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 6fd3fa08-b22f-457c-a6f2-fa4f38378618]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:24:15.430696Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:24:15.463892Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2043208964530769920#1:build-log.txt%3A5700)
<details>
<summary>all...</summary>

* _2026-04-12 06:07:17 &#43;0000 UTC_: <code>08:02:42:   {&#34;[namespace kubevirt-test-default4 name testvmi-nnfwm-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 64014b11-57a1-461e-b599-71f814ffcd45]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:02:28.034865Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:02:28.074924Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2043209126841946112#1:build-log.txt%3A4930)
<details><summary>context</summary>
<pre>08:02:42:
08:02:42:   Captured StdOut/StdErr Output &gt;&gt;
08:02:42:   {&#34;[namespace kubevirt-test-default4 name testvmi-nnfwm-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 64014b11-57a1-461e-b599-71f814ffcd45]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;Event(v1.ObjectReference{Kind:\&#34;VirtualMachineInstance\&#34;, Namespace:\&#34;kubevirt-test-default4\&#34;, Name:\&#34;testvmi-nnfwm-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\&#34;, UID:\&#34;64014b11-57a1-461e-b599-71f814ffcd45\&#34;, APIVersion:\&#34;kubevirt.io/v1\&#34;, ResourceVersion:\&#34;88940\&#34;, FieldPath:\&#34;\&#34;}): type: &#39;Normal&#39; reason: &#39;SuccessfulCreate&#39; Created virtual machine pod virt-launcher-testvmi-nnfwm-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxt5llq&#34;,&#34;pos&#34;:&#34;watcher.go:150&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:02:24.063043Z&#34;}
08:02:42:   {&#34;[namespace kubevirt-test-default4 name testvmi-nnfwm-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 64014b11-57a1-461e-b599-71f814ffcd45]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:02:28.034865Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:02:28.074924Z&#34;}
08:02:42:   &lt;&lt; Captured StdOut/StdErr Output
08:02:42: ------------------------------
08:03:21: • [39.841 seconds]</pre>
</details>


* _2026-04-12 06:06:47 &#43;0000 UTC_: <code>08:24:22:   {&#34;[namespace kubevirt-test-default3 name testvmi-zzvjh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 6fd3fa08-b22f-457c-a6f2-fa4f38378618]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:24:15.430696Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:24:15.463892Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2043208964530769920#1:build-log.txt%3A5700)
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

**1x**: _2026-04-10 06:11:44 &#43;0000 UTC_: <code>08:07:33: • [FAILED] [175.103 seconds]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17444/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.6/2042485527759818752#1:build-log.txt%3A3381)
<details>
<summary>all...</summary>

* _2026-04-10 06:11:44 &#43;0000 UTC_: <code>08:07:33: • [FAILED] [175.103 seconds]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17444/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.6/2042485527759818752#1:build-log.txt%3A3381)
<details><summary>context</summary>
<pre>08:04:38:
08:04:38:   There were additional failures detected.  To view them in detail run ginkgo -vv
08:04:38: ------------------------------
08:07:33: • [FAILED] [175.103 seconds]
08:07:33: [sig-operator]Operator [BeforeEach] [rfe_id:2291][crit:high][vendor:cnv-qe@redhat.com][level:component]infrastructure management [test_id:3150]should be able to update kubevirt install with custom image tag [Serial, sig-operator, Upgrade]
08:07:33:   [BeforeEach] tests/operator/operator.go:169
08:07:33:   [It] tests/operator/operator.go:1224</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> download failure in context (2x / 15.38%) </summary>

<hr/>

**2x**: _2026-04-09 14:20:27 &#43;0000 UTC_: <code>14:31:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17340/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2042246124231725056#1:build-log.txt%3A4058)
<details>
<summary>all...</summary>

* _2026-04-14 15:40:05 &#43;0000 UTC_: <code>15:43:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17067/pull-kubevirt-e2e-k8s-1.34-sig-storage/2044078011715686400#1:build-log.txt%3A353)
<details><summary>context</summary>
<pre>15:43:18: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted:
15:43:18: INFO: Elapsed time: 35.522s
15:43:18: INFO: 0 processes.
15:43:18: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>


* _2026-04-09 14:20:27 &#43;0000 UTC_: <code>14:31:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17340/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2042246124231725056#1:build-log.txt%3A4058)
<details><summary>context</summary>
<pre>14:31:10: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted:
14:31:10: INFO: Elapsed time: 1.273s
14:31:10: INFO: 0 processes.
14:31:10: ERROR: Build failed. Not running target
14:31:12: &#43; rm -f /tmp/kubevirt.deploy.7s31
make: *** [Makefile:188: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 7.69%) </summary>

<hr/>

**1x**: _2026-04-12 06:06:24 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043208958654550016#1:build-log.txt%3A539)
<details>
<summary>all...</summary>

* _2026-04-12 06:06:24 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043208958654550016#1:build-log.txt%3A539)
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

### needs-investigation (4x / 30.77%)

<details>
<summary> no matching pattern (3x / 23.08%) </summary>

<hr/>

**2x**: _2026-04-12 06:07:02 &#43;0000 UTC_: <code>08:59:54:   {&#34;[namespace kubevirt-test-default3 name testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid dd335e64-0fa8-4b94-8995-6c7eff1ebb59]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:59:30.101369Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:59:30.131590Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2043209123444559872#1:build-log.txt%3A9526)
<details>
<summary>all...</summary>

* _2026-04-12 06:07:04 &#43;0000 UTC_: <code>08:44:24:   {&#34;[namespace kubevirt-test-default2 name testvmi-56m8w-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 28ac500b-21a1-49e4-999e-ac55e4a75146]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:43:59.412605Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:43:59.440925Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2043209130977529856#1:build-log.txt%3A7635)
<details><summary>context</summary>
<pre>08:44:24:
08:44:24:   Captured StdOut/StdErr Output &gt;&gt;
08:44:24:   {&#34;[namespace kubevirt-test-default2 name testvmi-56m8w-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 28ac500b-21a1-49e4-999e-ac55e4a75146]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;Event(v1.ObjectReference{Kind:\&#34;VirtualMachineInstance\&#34;, Namespace:\&#34;kubevirt-test-default2\&#34;, Name:\&#34;testvmi-56m8w-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\&#34;, UID:\&#34;28ac500b-21a1-49e4-999e-ac55e4a75146\&#34;, APIVersion:\&#34;kubevirt.io/v1\&#34;, ResourceVersion:\&#34;140603\&#34;, FieldPath:\&#34;\&#34;}): type: &#39;Normal&#39; reason: &#39;SuccessfulCreate&#39; Created virtual machine pod virt-launcher-testvmi-56m8w-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx5nmm7&#34;,&#34;pos&#34;:&#34;watcher.go:150&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:43:55.709663Z&#34;}
08:44:24:   {&#34;[namespace kubevirt-test-default2 name testvmi-56m8w-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 28ac500b-21a1-49e4-999e-ac55e4a75146]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:43:59.412605Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:43:59.440925Z&#34;}
08:44:24:   &lt;&lt; Captured StdOut/StdErr Output
08:44:24: ------------------------------
08:44:56: • [35.870 seconds]</pre>
</details>


* _2026-04-12 06:07:02 &#43;0000 UTC_: <code>08:59:54:   {&#34;[namespace kubevirt-test-default3 name testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid dd335e64-0fa8-4b94-8995-6c7eff1ebb59]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:59:30.101369Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:59:30.131590Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2043209123444559872#1:build-log.txt%3A9526)
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

**1x**: _2026-04-10 06:12:35 &#43;0000 UTC_: <code>08:07:41: • [FAILED] [182.207 seconds]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17445/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2042485694953164800#1:build-log.txt%3A5316)
<details>
<summary>all...</summary>

* _2026-04-10 06:12:35 &#43;0000 UTC_: <code>08:07:41: • [FAILED] [182.207 seconds]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17445/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2042485694953164800#1:build-log.txt%3A5316)
<details><summary>context</summary>
<pre>08:04:38:
08:04:38:   There were additional failures detected.  To view them in detail run ginkgo -vv
08:04:38: ------------------------------
08:07:41: • [FAILED] [182.207 seconds]
08:07:41: [sig-network]  Istio with passt binding [BeforeEach] Virtual Machine with istio supported interface Outbound traffic With Sidecar allowing only registered external services VMI with explicit ports Should not be able to reach http service outside of mesh [sig-network, Istio, netCustomBindingPlugins, Serial]
08:07:41:   [BeforeEach] tests/network/vmi_istio.go:456
08:07:41:   [It] tests/network/vmi_istio.go:428</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> no error snippets (1x / 7.69%) </summary>

<hr/>

**1x**: _2026-04-12 06:06:33 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2043208957803106304)
<details>
<summary>all...</summary>

* _2026-04-12 06:06:33 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2043208957803106304)

</details>

<hr/>
</details>

### internal (3x / 23.08%)

<details>
<summary> make cluster lifecycle target failure (2x / 15.38%) </summary>

<hr/>

**2x**: _2026-04-09 01:56:53 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2042058998902951936#1:build-log.txt%3A1040)
<details>
<summary>all...</summary>

* _2026-04-09 01:56:54 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2042059010437287936#1:build-log.txt%3A1142)
<details><summary>context</summary>
<pre>02:23:46: INFO: 2 processes: 1 internal, 1 processwrapper-sandbox.
02:23:46: INFO: Running command line: bazel-bin/example-guest-agent-copier /root/go/src/kubevirt.io/kubevirt/_out/cmd/example-guest-agent/example-guest-agent
02:23:58: &#43; rm -f /tmp/kubevirt.deploy.h4Rn
make: *** [Makefile:174: cluster-sync] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>


* _2026-04-09 01:56:53 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2042058998902951936#1:build-log.txt%3A1040)
<details><summary>context</summary>
<pre>02:40:33: INFO: 2 processes: 1 internal, 1 processwrapper-sandbox.
02:40:33: INFO: Running command line: bazel-bin/example-guest-agent-copier /root/go/src/kubevirt.io/kubevirt/_out/cmd/example-guest-agent/example-guest-agent
02:40:50: &#43; rm -f /tmp/kubevirt.deploy.RzZ0
make: *** [Makefile:174: cluster-sync] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> make bazel-build target failure (1x / 7.69%) </summary>

<hr/>

**1x**: _2026-04-12 08:21:13 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043242879446421504#1:build-log.txt%3A4241)
<details>
<summary>all...</summary>

* _2026-04-12 08:21:13 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043242879446421504#1:build-log.txt%3A4241)
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


### release-1.7 (7x / 53.85%)


#### internal (3x / 42.86%)

<details>
<summary> make cluster lifecycle target failure (2x / 28.57%) </summary>

<hr/>

**2x**: _2026-04-09 01:56:53 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2042058998902951936#1:build-log.txt%3A1040)
<details>
<summary>all...</summary>

* _2026-04-09 01:56:54 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2042059010437287936#1:build-log.txt%3A1142)

* _2026-04-09 01:56:53 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2042058998902951936#1:build-log.txt%3A1040)

</details>

<hr/>
</details>
<details>
<summary> make bazel-build target failure (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-04-12 08:21:13 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043242879446421504#1:build-log.txt%3A4241)
<details>
<summary>all...</summary>

* _2026-04-12 08:21:13 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043242879446421504#1:build-log.txt%3A4241)

</details>

<hr/>
</details>

#### needs-investigation (2x / 28.57%)

<details>
<summary> no matching pattern (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-04-10 06:12:35 &#43;0000 UTC_: <code>08:07:41: • [FAILED] [182.207 seconds]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17445/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2042485694953164800#1:build-log.txt%3A5316)
<details>
<summary>all...</summary>

* _2026-04-10 06:12:35 &#43;0000 UTC_: <code>08:07:41: • [FAILED] [182.207 seconds]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17445/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2042485694953164800#1:build-log.txt%3A5316)

</details>

<hr/>
</details>
<details>
<summary> no error snippets (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-04-12 06:06:33 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2043208957803106304)
<details>
<summary>all...</summary>

* _2026-04-12 06:06:33 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2043208957803106304)

</details>

<hr/>
</details>

#### external (2x / 28.57%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-04-12 06:06:47 &#43;0000 UTC_: <code>08:24:22:   {&#34;[namespace kubevirt-test-default3 name testvmi-zzvjh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 6fd3fa08-b22f-457c-a6f2-fa4f38378618]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:24:15.430696Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:24:15.463892Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2043208964530769920#1:build-log.txt%3A5700)
<details>
<summary>all...</summary>

* _2026-04-12 06:06:47 &#43;0000 UTC_: <code>08:24:22:   {&#34;[namespace kubevirt-test-default3 name testvmi-zzvjh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 6fd3fa08-b22f-457c-a6f2-fa4f38378618]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:24:15.430696Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:24:15.463892Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2043208964530769920#1:build-log.txt%3A5700)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-04-12 06:06:24 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043208958654550016#1:build-log.txt%3A539)
<details>
<summary>all...</summary>

* _2026-04-12 06:06:24 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043208958654550016#1:build-log.txt%3A539)

</details>

<hr/>
</details>

### release-1.8 (3x / 23.08%)


#### needs-investigation (2x / 66.67%)

<details>
<summary> no matching pattern (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-04-12 06:07:02 &#43;0000 UTC_: <code>08:59:54:   {&#34;[namespace kubevirt-test-default3 name testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid dd335e64-0fa8-4b94-8995-6c7eff1ebb59]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:59:30.101369Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:59:30.131590Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2043209123444559872#1:build-log.txt%3A9526)
<details>
<summary>all...</summary>

* _2026-04-12 06:07:04 &#43;0000 UTC_: <code>08:44:24:   {&#34;[namespace kubevirt-test-default2 name testvmi-56m8w-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 28ac500b-21a1-49e4-999e-ac55e4a75146]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:43:59.412605Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:43:59.440925Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2043209130977529856#1:build-log.txt%3A7635)

* _2026-04-12 06:07:02 &#43;0000 UTC_: <code>08:59:54:   {&#34;[namespace kubevirt-test-default3 name testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid dd335e64-0fa8-4b94-8995-6c7eff1ebb59]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:59:30.101369Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:59:30.131590Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2043209123444559872#1:build-log.txt%3A9526)

</details>

<hr/>
</details>

#### external (1x / 33.33%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-04-12 06:07:17 &#43;0000 UTC_: <code>08:02:42:   {&#34;[namespace kubevirt-test-default4 name testvmi-nnfwm-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 64014b11-57a1-461e-b599-71f814ffcd45]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:02:28.034865Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:02:28.074924Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2043209126841946112#1:build-log.txt%3A4930)
<details>
<summary>all...</summary>

* _2026-04-12 06:07:17 &#43;0000 UTC_: <code>08:02:42:   {&#34;[namespace kubevirt-test-default4 name testvmi-nnfwm-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 64014b11-57a1-461e-b599-71f814ffcd45]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:02:28.034865Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:02:28.074924Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2043209126841946112#1:build-log.txt%3A4930)

</details>

<hr/>
</details>

### release-1.6 (1x / 7.69%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-04-10 06:11:44 &#43;0000 UTC_: <code>08:07:33: • [FAILED] [175.103 seconds]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17444/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.6/2042485527759818752#1:build-log.txt%3A3381)
<details>
<summary>all...</summary>

* _2026-04-10 06:11:44 &#43;0000 UTC_: <code>08:07:33: • [FAILED] [175.103 seconds]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17444/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.6/2042485527759818752#1:build-log.txt%3A3381)

</details>

<hr/>
</details>

### main (2x / 15.38%)


#### external (2x / 100.00%)

<details>
<summary> download failure in context (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-04-09 14:20:27 &#43;0000 UTC_: <code>14:31:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17340/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2042246124231725056#1:build-log.txt%3A4058)
<details>
<summary>all...</summary>

* _2026-04-14 15:40:05 &#43;0000 UTC_: <code>15:43:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17067/pull-kubevirt-e2e-k8s-1.34-sig-storage/2044078011715686400#1:build-log.txt%3A353)

* _2026-04-09 14:20:27 &#43;0000 UTC_: <code>14:31:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17340/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2042246124231725056#1:build-log.txt%3A4058)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-storage (7x / 53.85%)


#### needs-investigation (3x / 42.86%)

<details>
<summary> no matching pattern (2x / 28.57%) </summary>

<hr/>

**2x**: _2026-04-12 06:07:02 &#43;0000 UTC_: <code>08:59:54:   {&#34;[namespace kubevirt-test-default3 name testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid dd335e64-0fa8-4b94-8995-6c7eff1ebb59]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:59:30.101369Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:59:30.131590Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2043209123444559872#1:build-log.txt%3A9526)
<details>
<summary>all...</summary>

* _2026-04-12 06:07:04 &#43;0000 UTC_: <code>08:44:24:   {&#34;[namespace kubevirt-test-default2 name testvmi-56m8w-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 28ac500b-21a1-49e4-999e-ac55e4a75146]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:43:59.412605Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:43:59.440925Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2043209130977529856#1:build-log.txt%3A7635)

* _2026-04-12 06:07:02 &#43;0000 UTC_: <code>08:59:54:   {&#34;[namespace kubevirt-test-default3 name testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid dd335e64-0fa8-4b94-8995-6c7eff1ebb59]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:59:30.101369Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:59:30.131590Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2043209123444559872#1:build-log.txt%3A9526)

</details>

<hr/>
</details>
<details>
<summary> no error snippets (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-04-12 06:06:33 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2043208957803106304)
<details>
<summary>all...</summary>

* _2026-04-12 06:06:33 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2043208957803106304)

</details>

<hr/>
</details>

#### external (3x / 42.86%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 28.57%) </summary>

<hr/>

**2x**: _2026-04-12 06:06:47 &#43;0000 UTC_: <code>08:24:22:   {&#34;[namespace kubevirt-test-default3 name testvmi-zzvjh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 6fd3fa08-b22f-457c-a6f2-fa4f38378618]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:24:15.430696Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:24:15.463892Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2043208964530769920#1:build-log.txt%3A5700)
<details>
<summary>all...</summary>

* _2026-04-12 06:07:17 &#43;0000 UTC_: <code>08:02:42:   {&#34;[namespace kubevirt-test-default4 name testvmi-nnfwm-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 64014b11-57a1-461e-b599-71f814ffcd45]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:02:28.034865Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:02:28.074924Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2043209126841946112#1:build-log.txt%3A4930)

* _2026-04-12 06:06:47 &#43;0000 UTC_: <code>08:24:22:   {&#34;[namespace kubevirt-test-default3 name testvmi-zzvjh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 6fd3fa08-b22f-457c-a6f2-fa4f38378618]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:24:15.430696Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:24:15.463892Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2043208964530769920#1:build-log.txt%3A5700)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-04-14 15:40:05 &#43;0000 UTC_: <code>15:43:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17067/pull-kubevirt-e2e-k8s-1.34-sig-storage/2044078011715686400#1:build-log.txt%3A353)
<details>
<summary>all...</summary>

* _2026-04-14 15:40:05 &#43;0000 UTC_: <code>15:43:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17067/pull-kubevirt-e2e-k8s-1.34-sig-storage/2044078011715686400#1:build-log.txt%3A353)

</details>

<hr/>
</details>

#### internal (1x / 14.29%)

<details>
<summary> make cluster lifecycle target failure (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-04-09 01:56:54 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2042059010437287936#1:build-log.txt%3A1142)
<details>
<summary>all...</summary>

* _2026-04-09 01:56:54 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2042059010437287936#1:build-log.txt%3A1142)

</details>

<hr/>
</details>

### sig-compute (5x / 38.46%)


#### external (3x / 60.00%)

<details>
<summary> download failure in context (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-04-09 14:20:27 &#43;0000 UTC_: <code>14:31:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17340/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2042246124231725056#1:build-log.txt%3A4058)
<details>
<summary>all...</summary>

* _2026-04-09 14:20:27 &#43;0000 UTC_: <code>14:31:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17340/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2042246124231725056#1:build-log.txt%3A4058)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-04-12 06:06:24 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043208958654550016#1:build-log.txt%3A539)
<details>
<summary>all...</summary>

* _2026-04-12 06:06:24 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043208958654550016#1:build-log.txt%3A539)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-04-10 06:11:44 &#43;0000 UTC_: <code>08:07:33: • [FAILED] [175.103 seconds]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17444/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.6/2042485527759818752#1:build-log.txt%3A3381)
<details>
<summary>all...</summary>

* _2026-04-10 06:11:44 &#43;0000 UTC_: <code>08:07:33: • [FAILED] [175.103 seconds]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17444/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.6/2042485527759818752#1:build-log.txt%3A3381)

</details>

<hr/>
</details>

#### internal (2x / 40.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-04-09 01:56:53 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2042058998902951936#1:build-log.txt%3A1040)
<details>
<summary>all...</summary>

* _2026-04-09 01:56:53 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17426/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2042058998902951936#1:build-log.txt%3A1040)

</details>

<hr/>
</details>
<details>
<summary> make bazel-build target failure (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-04-12 08:21:13 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043242879446421504#1:build-log.txt%3A4241)
<details>
<summary>all...</summary>

* _2026-04-12 08:21:13 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043242879446421504#1:build-log.txt%3A4241)

</details>

<hr/>
</details>

### sig-network (1x / 7.69%)


#### needs-investigation (1x / 100.00%)

<details>
<summary> no matching pattern (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-04-10 06:12:35 &#43;0000 UTC_: <code>08:07:41: • [FAILED] [182.207 seconds]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17445/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2042485694953164800#1:build-log.txt%3A5316)
<details>
<summary>all...</summary>

* _2026-04-10 06:12:35 &#43;0000 UTC_: <code>08:07:41: • [FAILED] [182.207 seconds]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17445/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2042485694953164800#1:build-log.txt%3A5316)

</details>

<hr/>
</details>

Last updated: 2026-04-16 00:34:28
