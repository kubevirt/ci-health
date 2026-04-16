



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-04-14 (7x / 41.18%)


#### external (6x / 85.71%)

<details>
<summary> download failure in context (2x / 28.57%) </summary>

<hr/>

**2x**: _2026-04-14 15:40:05 &#43;0000 UTC_: <code>15:43:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17067/pull-kubevirt-e2e-k8s-1.34-sig-storage/2044078011715686400#1:build-log.txt%3A353)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:41 &#43;0000 UTC_: <code>22:39:51: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2044182449767845888#1:build-log.txt%3A253)

* _2026-04-14 15:40:05 &#43;0000 UTC_: <code>15:43:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17067/pull-kubevirt-e2e-k8s-1.34-sig-storage/2044078011715686400#1:build-log.txt%3A353)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (2x / 28.57%) </summary>

<hr/>

**2x**: _2026-04-14 22:34:40 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2044182447226097664#1:build-log.txt%3A197)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:45 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2044182454352220160#1:build-log.txt%3A193)

* _2026-04-14 22:34:40 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2044182447226097664#1:build-log.txt%3A197)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-04-14 22:34:50 &#43;0000 UTC_: <code>23:54:10: I0414 19:54:10.637122    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2044182448891236352#1:build-log.txt%3A762)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:50 &#43;0000 UTC_: <code>23:54:10: I0414 19:54:10.637122    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2044182448891236352#1:build-log.txt%3A762)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-04-14 22:34:43 &#43;0000 UTC_: <code>00:33:03:   [FAILED] in [AfterEach] - tests/libkubevirt/kubevirt.go:49 @ 04/15/26 00:33:03.525</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2044182452733218816#1:build-log.txt%3A4923)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:43 &#43;0000 UTC_: <code>00:33:03:   [FAILED] in [AfterEach] - tests/libkubevirt/kubevirt.go:49 @ 04/15/26 00:33:03.525</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2044182452733218816#1:build-log.txt%3A4923)

</details>

<hr/>
</details>

#### needs-investigation (1x / 14.29%)

<details>
<summary> no matching pattern (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-04-14 22:34:46 &#43;0000 UTC_: <code>01:48:44:   {&#34;[namespace kubevirt-test-default2 name testvmi-rljs5-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 8017f1d0-8e11-49df-bdbd-fa1d76cff6d8]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-15T01:48:19.933126Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T01:48:19.973140Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2044182453085540352#1:build-log.txt%3A5398)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:46 &#43;0000 UTC_: <code>01:48:44:   {&#34;[namespace kubevirt-test-default2 name testvmi-rljs5-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 8017f1d0-8e11-49df-bdbd-fa1d76cff6d8]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-15T01:48:19.933126Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T01:48:19.973140Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2044182453085540352#1:build-log.txt%3A5398)

</details>

<hr/>
</details>

### 2026-04-12 (7x / 41.18%)


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

### 2026-04-10 (2x / 11.76%)


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

### 2026-04-09 (1x / 5.88%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

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


### external (11x / 64.71%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (4x / 23.53%) </summary>

<hr/>

**1x**: _2026-04-14 22:34:43 &#43;0000 UTC_: <code>00:33:03:   [FAILED] in [AfterEach] - tests/libkubevirt/kubevirt.go:49 @ 04/15/26 00:33:03.525</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2044182452733218816#1:build-log.txt%3A4923)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:43 &#43;0000 UTC_: <code>00:33:03:   [FAILED] in [AfterEach] - tests/libkubevirt/kubevirt.go:49 @ 04/15/26 00:33:03.525</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2044182452733218816#1:build-log.txt%3A4923)
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
<summary> download failure in context (3x / 17.65%) </summary>

<hr/>

**3x**: _2026-04-14 15:40:05 &#43;0000 UTC_: <code>15:43:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17067/pull-kubevirt-e2e-k8s-1.34-sig-storage/2044078011715686400#1:build-log.txt%3A353)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:41 &#43;0000 UTC_: <code>22:39:51: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2044182449767845888#1:build-log.txt%3A253)
<details><summary>context</summary>
<pre>22:39:51: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted:
22:39:51: INFO: Elapsed time: 155.381s
22:39:51: INFO: 0 processes.
22:39:51: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>


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
<summary> container image pull failure in context (3x / 17.65%) </summary>

<hr/>

**2x**: _2026-04-14 22:34:40 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2044182447226097664#1:build-log.txt%3A197)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:45 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2044182454352220160#1:build-log.txt%3A193)
<details><summary>context</summary>
<pre>hack/dockerized &#34;export BUILD_ARCH= &amp;&amp; DOCKER_PREFIX= DOCKER_TAG= DOCKER_TAG_ALT= IMAGE_PREFIX= IMAGE_PREFIX_ALT=kv- ./hack/multi-arch.sh build-images&#34;
22:36:54: Trying to pull quay.io/kubevirt/builder:2602201229-4494d1587...
22:37:18: Error: unable to copy from source docker://quay.io/kubevirt/builder:2602201229-4494d1587: copying system image from manifest list: parsing image configuration: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/23/23779b65e496e294d12e915d6b24032b890f5b970a157a8ea5f99807c36da777?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20260414%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20260414T223700Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=a2745b76664f40c7d4fcc7a7231ae0c42fc4cf663aea78b564e3eb814c0942ec&amp;region=us-east-1&amp;namespace=kubevirt&amp;repo_name=builder&amp;akamai_signature=exp=1776207120~hmac=4bdce0f207fe57ec0390fc84e5516f5e4e44dac2136d6c35390f46b7c246d4b6&#34;: dial tcp: lookup cdn01.quay.io: Temporary failure in name resolution
make: *** [Makefile:39: bazel-build-images] Error 125
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-04-14 22:34:40 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2044182447226097664#1:build-log.txt%3A197)
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
<details>
<summary> transient kube-apiserver body decode noise (1x / 5.88%) </summary>

<hr/>

**1x**: _2026-04-14 22:34:50 &#43;0000 UTC_: <code>23:54:10: I0414 19:54:10.637122    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2044182448891236352#1:build-log.txt%3A762)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:50 &#43;0000 UTC_: <code>23:54:10: I0414 19:54:10.637122    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2044182448891236352#1:build-log.txt%3A762)
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

### needs-investigation (5x / 29.41%)

<details>
<summary> no matching pattern (4x / 23.53%) </summary>

<hr/>

**3x**: _2026-04-12 06:07:02 &#43;0000 UTC_: <code>08:59:54:   {&#34;[namespace kubevirt-test-default3 name testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid dd335e64-0fa8-4b94-8995-6c7eff1ebb59]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:59:30.101369Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:59:30.131590Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2043209123444559872#1:build-log.txt%3A9526)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:46 &#43;0000 UTC_: <code>01:48:44:   {&#34;[namespace kubevirt-test-default2 name testvmi-rljs5-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 8017f1d0-8e11-49df-bdbd-fa1d76cff6d8]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-15T01:48:19.933126Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T01:48:19.973140Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2044182453085540352#1:build-log.txt%3A5398)
<details><summary>context</summary>
<pre>01:48:44:
01:48:44:   Captured StdOut/StdErr Output &gt;&gt;
01:48:44:   {&#34;[namespace kubevirt-test-default2 name testvmi-rljs5-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 8017f1d0-8e11-49df-bdbd-fa1d76cff6d8]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;Event(v1.ObjectReference{Kind:\&#34;VirtualMachineInstance\&#34;, Namespace:\&#34;kubevirt-test-default2\&#34;, Name:\&#34;testvmi-rljs5-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\&#34;, UID:\&#34;8017f1d0-8e11-49df-bdbd-fa1d76cff6d8\&#34;, APIVersion:\&#34;kubevirt.io/v1\&#34;, ResourceVersion:\&#34;157662\&#34;, FieldPath:\&#34;\&#34;}): type: &#39;Normal&#39; reason: &#39;SuccessfulCreate&#39; Created virtual machine pod virt-launcher-testvmi-rljs5-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxf842j&#34;,&#34;pos&#34;:&#34;watcher.go:150&#34;,&#34;timestamp&#34;:&#34;2026-04-15T01:48:16.385235Z&#34;}
01:48:44:   {&#34;[namespace kubevirt-test-default2 name testvmi-rljs5-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 8017f1d0-8e11-49df-bdbd-fa1d76cff6d8]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-15T01:48:19.933126Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T01:48:19.973140Z&#34;}
01:48:44:   &lt;&lt; Captured StdOut/StdErr Output
01:48:44: ------------------------------
01:49:21: • [110.511 seconds]</pre>
</details>


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
<summary> no error snippets (1x / 5.88%) </summary>

<hr/>

**1x**: _2026-04-12 06:06:33 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2043208957803106304)
<details>
<summary>all...</summary>

* _2026-04-12 06:06:33 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2043208957803106304)

</details>

<hr/>
</details>

### internal (1x / 5.88%)

<details>
<summary> make bazel-build target failure (1x / 5.88%) </summary>

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


### release-1.8 (9x / 52.94%)


#### external (6x / 66.67%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 22.22%) </summary>

<hr/>

**1x**: _2026-04-14 22:34:43 &#43;0000 UTC_: <code>00:33:03:   [FAILED] in [AfterEach] - tests/libkubevirt/kubevirt.go:49 @ 04/15/26 00:33:03.525</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2044182452733218816#1:build-log.txt%3A4923)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:43 &#43;0000 UTC_: <code>00:33:03:   [FAILED] in [AfterEach] - tests/libkubevirt/kubevirt.go:49 @ 04/15/26 00:33:03.525</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2044182452733218816#1:build-log.txt%3A4923)

</details>

<hr/>

**1x**: _2026-04-12 06:07:17 &#43;0000 UTC_: <code>08:02:42:   {&#34;[namespace kubevirt-test-default4 name testvmi-nnfwm-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 64014b11-57a1-461e-b599-71f814ffcd45]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:02:28.034865Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:02:28.074924Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2043209126841946112#1:build-log.txt%3A4930)
<details>
<summary>all...</summary>

* _2026-04-12 06:07:17 &#43;0000 UTC_: <code>08:02:42:   {&#34;[namespace kubevirt-test-default4 name testvmi-nnfwm-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 64014b11-57a1-461e-b599-71f814ffcd45]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:02:28.034865Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:02:28.074924Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2043209126841946112#1:build-log.txt%3A4930)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (2x / 22.22%) </summary>

<hr/>

**2x**: _2026-04-14 22:34:40 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2044182447226097664#1:build-log.txt%3A197)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:45 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2044182454352220160#1:build-log.txt%3A193)

* _2026-04-14 22:34:40 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2044182447226097664#1:build-log.txt%3A197)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-04-14 22:34:50 &#43;0000 UTC_: <code>23:54:10: I0414 19:54:10.637122    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2044182448891236352#1:build-log.txt%3A762)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:50 &#43;0000 UTC_: <code>23:54:10: I0414 19:54:10.637122    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2044182448891236352#1:build-log.txt%3A762)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-04-14 22:34:41 &#43;0000 UTC_: <code>22:39:51: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2044182449767845888#1:build-log.txt%3A253)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:41 &#43;0000 UTC_: <code>22:39:51: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2044182449767845888#1:build-log.txt%3A253)

</details>

<hr/>
</details>

#### needs-investigation (3x / 33.33%)

<details>
<summary> no matching pattern (3x / 33.33%) </summary>

<hr/>

**3x**: _2026-04-12 06:07:02 &#43;0000 UTC_: <code>08:59:54:   {&#34;[namespace kubevirt-test-default3 name testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid dd335e64-0fa8-4b94-8995-6c7eff1ebb59]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:59:30.101369Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:59:30.131590Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2043209123444559872#1:build-log.txt%3A9526)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:46 &#43;0000 UTC_: <code>01:48:44:   {&#34;[namespace kubevirt-test-default2 name testvmi-rljs5-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 8017f1d0-8e11-49df-bdbd-fa1d76cff6d8]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-15T01:48:19.933126Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T01:48:19.973140Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2044182453085540352#1:build-log.txt%3A5398)

* _2026-04-12 06:07:04 &#43;0000 UTC_: <code>08:44:24:   {&#34;[namespace kubevirt-test-default2 name testvmi-56m8w-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 28ac500b-21a1-49e4-999e-ac55e4a75146]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:43:59.412605Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:43:59.440925Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2043209130977529856#1:build-log.txt%3A7635)

* _2026-04-12 06:07:02 &#43;0000 UTC_: <code>08:59:54:   {&#34;[namespace kubevirt-test-default3 name testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid dd335e64-0fa8-4b94-8995-6c7eff1ebb59]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:59:30.101369Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:59:30.131590Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2043209123444559872#1:build-log.txt%3A9526)

</details>

<hr/>
</details>

### main (2x / 11.76%)


#### external (2x / 100.00%)

<details>
<summary> download failure in context (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-04-14 15:40:05 &#43;0000 UTC_: <code>15:43:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17067/pull-kubevirt-e2e-k8s-1.34-sig-storage/2044078011715686400#1:build-log.txt%3A353)
<details>
<summary>all...</summary>

* _2026-04-14 15:40:05 &#43;0000 UTC_: <code>15:43:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17067/pull-kubevirt-e2e-k8s-1.34-sig-storage/2044078011715686400#1:build-log.txt%3A353)

* _2026-04-09 14:20:27 &#43;0000 UTC_: <code>14:31:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17340/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2042246124231725056#1:build-log.txt%3A4058)

</details>

<hr/>
</details>

### release-1.7 (5x / 29.41%)


#### needs-investigation (2x / 40.00%)

<details>
<summary> no error snippets (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-04-12 06:06:33 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2043208957803106304)
<details>
<summary>all...</summary>

* _2026-04-12 06:06:33 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2043208957803106304)

</details>

<hr/>
</details>
<details>
<summary> no matching pattern (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-04-10 06:12:35 &#43;0000 UTC_: <code>08:07:41: • [FAILED] [182.207 seconds]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17445/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2042485694953164800#1:build-log.txt%3A5316)
<details>
<summary>all...</summary>

* _2026-04-10 06:12:35 &#43;0000 UTC_: <code>08:07:41: • [FAILED] [182.207 seconds]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17445/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2042485694953164800#1:build-log.txt%3A5316)

</details>

<hr/>
</details>

#### external (2x / 40.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-04-12 06:06:47 &#43;0000 UTC_: <code>08:24:22:   {&#34;[namespace kubevirt-test-default3 name testvmi-zzvjh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 6fd3fa08-b22f-457c-a6f2-fa4f38378618]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:24:15.430696Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:24:15.463892Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2043208964530769920#1:build-log.txt%3A5700)
<details>
<summary>all...</summary>

* _2026-04-12 06:06:47 &#43;0000 UTC_: <code>08:24:22:   {&#34;[namespace kubevirt-test-default3 name testvmi-zzvjh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 6fd3fa08-b22f-457c-a6f2-fa4f38378618]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:24:15.430696Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:24:15.463892Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.7/2043208964530769920#1:build-log.txt%3A5700)

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

#### internal (1x / 20.00%)

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

### release-1.6 (1x / 5.88%)


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

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-storage (8x / 47.06%)


#### external (4x / 50.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 25.00%) </summary>

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
<summary> download failure in context (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-04-14 15:40:05 &#43;0000 UTC_: <code>15:43:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17067/pull-kubevirt-e2e-k8s-1.34-sig-storage/2044078011715686400#1:build-log.txt%3A353)
<details>
<summary>all...</summary>

* _2026-04-14 15:40:05 &#43;0000 UTC_: <code>15:43:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17067/pull-kubevirt-e2e-k8s-1.34-sig-storage/2044078011715686400#1:build-log.txt%3A353)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-04-14 22:34:50 &#43;0000 UTC_: <code>23:54:10: I0414 19:54:10.637122    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2044182448891236352#1:build-log.txt%3A762)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:50 &#43;0000 UTC_: <code>23:54:10: I0414 19:54:10.637122    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2044182448891236352#1:build-log.txt%3A762)

</details>

<hr/>
</details>

#### needs-investigation (4x / 50.00%)

<details>
<summary> no matching pattern (3x / 37.50%) </summary>

<hr/>

**3x**: _2026-04-12 06:07:02 &#43;0000 UTC_: <code>08:59:54:   {&#34;[namespace kubevirt-test-default3 name testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid dd335e64-0fa8-4b94-8995-6c7eff1ebb59]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:59:30.101369Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:59:30.131590Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2043209123444559872#1:build-log.txt%3A9526)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:46 &#43;0000 UTC_: <code>01:48:44:   {&#34;[namespace kubevirt-test-default2 name testvmi-rljs5-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 8017f1d0-8e11-49df-bdbd-fa1d76cff6d8]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-15T01:48:19.933126Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-15T01:48:19.973140Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2044182453085540352#1:build-log.txt%3A5398)

* _2026-04-12 06:07:04 &#43;0000 UTC_: <code>08:44:24:   {&#34;[namespace kubevirt-test-default2 name testvmi-56m8w-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 28ac500b-21a1-49e4-999e-ac55e4a75146]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:43:59.412605Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:43:59.440925Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2043209130977529856#1:build-log.txt%3A7635)

* _2026-04-12 06:07:02 &#43;0000 UTC_: <code>08:59:54:   {&#34;[namespace kubevirt-test-default3 name testvmi-v2bmp-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid dd335e64-0fa8-4b94-8995-6c7eff1ebb59]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-04-12T08:59:30.101369Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-04-12T08:59:30.131590Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17465/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2043209123444559872#1:build-log.txt%3A9526)

</details>

<hr/>
</details>
<details>
<summary> no error snippets (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-04-12 06:06:33 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2043208957803106304)
<details>
<summary>all...</summary>

* _2026-04-12 06:06:33 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2043208957803106304)

</details>

<hr/>
</details>

### sig-network (2x / 11.76%)


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

**1x**: _2026-04-14 22:34:43 &#43;0000 UTC_: <code>00:33:03:   [FAILED] in [AfterEach] - tests/libkubevirt/kubevirt.go:49 @ 04/15/26 00:33:03.525</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2044182452733218816#1:build-log.txt%3A4923)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:43 &#43;0000 UTC_: <code>00:33:03:   [FAILED] in [AfterEach] - tests/libkubevirt/kubevirt.go:49 @ 04/15/26 00:33:03.525</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2044182452733218816#1:build-log.txt%3A4923)

</details>

<hr/>
</details>

### sig-compute (7x / 41.18%)


#### external (6x / 85.71%)

<details>
<summary> container image pull failure in context (3x / 42.86%) </summary>

<hr/>

**2x**: _2026-04-14 22:34:40 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2044182447226097664#1:build-log.txt%3A197)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:45 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2044182454352220160#1:build-log.txt%3A193)

* _2026-04-14 22:34:40 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2044182447226097664#1:build-log.txt%3A197)

</details>

<hr/>

**1x**: _2026-04-12 06:06:24 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043208958654550016#1:build-log.txt%3A539)
<details>
<summary>all...</summary>

* _2026-04-12 06:06:24 &#43;0000 UTC_: <code>make: *** [Makefile:159: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17464/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2043208958654550016#1:build-log.txt%3A539)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (2x / 28.57%) </summary>

<hr/>

**2x**: _2026-04-14 22:34:41 &#43;0000 UTC_: <code>22:39:51: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2044182449767845888#1:build-log.txt%3A253)
<details>
<summary>all...</summary>

* _2026-04-14 22:34:41 &#43;0000 UTC_: <code>22:39:51: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17499/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2044182449767845888#1:build-log.txt%3A253)

* _2026-04-09 14:20:27 &#43;0000 UTC_: <code>14:31:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17340/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2042246124231725056#1:build-log.txt%3A4058)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-04-10 06:11:44 &#43;0000 UTC_: <code>08:07:33: • [FAILED] [175.103 seconds]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17444/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.6/2042485527759818752#1:build-log.txt%3A3381)
<details>
<summary>all...</summary>

* _2026-04-10 06:11:44 &#43;0000 UTC_: <code>08:07:33: • [FAILED] [175.103 seconds]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17444/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.6/2042485527759818752#1:build-log.txt%3A3381)

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

Last updated: 2026-04-16 09:41:19
