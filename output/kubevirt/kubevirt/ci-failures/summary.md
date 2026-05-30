



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-05-28 (1x / 3.33%)


#### needs-investigation (1x / 100.00%)

<details>
<summary> no error snippets (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-05-28 09:40:36 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17942/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059932689309372416)
<details>
<summary>all...</summary>

* _2026-05-28 09:40:36 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17942/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059932689309372416)

</details>

<hr/>
</details>

### 2026-05-27 (1x / 3.33%)


#### needs-investigation (1x / 100.00%)

<details>
<summary> no error snippets (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-05-27 10:39:00 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17911/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059584804717858816)
<details>
<summary>all...</summary>

* _2026-05-27 10:39:00 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17911/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059584804717858816)

</details>

<hr/>
</details>

### 2026-05-26 (12x / 40.00%)


#### needs-investigation (6x / 50.00%)

<details>
<summary> no matching pattern (4x / 33.33%) </summary>

<hr/>

**1x**: _2026-05-26 18:24:41 &#43;0000 UTC_: <code>20:51:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-slpvj-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid c86a0bc5-07ae-4204-a46c-f89c66b191b5]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-05-26T20:50:37.865872Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-05-26T20:50:37.903851Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17919/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2059339837684584448#1:build-log.txt%3A8998)
<details>
<summary>all...</summary>

* _2026-05-26 18:24:41 &#43;0000 UTC_: <code>20:51:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-slpvj-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid c86a0bc5-07ae-4204-a46c-f89c66b191b5]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-05-26T20:50:37.865872Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-05-26T20:50:37.903851Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17919/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2059339837684584448#1:build-log.txt%3A8998)

</details>

<hr/>

**3x**: _2026-05-26 00:07:20 &#43;0000 UTC_: <code>00:23:57: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: dial tcp 127.0.0.1:10248: connect: connection refused</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059063686630215680#1:build-log.txt%3A6633)
<details>
<summary>all...</summary>

* _2026-05-26 04:07:33 &#43;0000 UTC_: <code>04:23:46: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059124119928049664#1:build-log.txt%3A6615)

* _2026-05-26 02:07:30 &#43;0000 UTC_: <code>02:23:47: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2059093907123212288#1:build-log.txt%3A6643)

* _2026-05-26 00:07:20 &#43;0000 UTC_: <code>00:23:57: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: dial tcp 127.0.0.1:10248: connect: connection refused</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059063686630215680#1:build-log.txt%3A6633)

</details>

<hr/>
</details>
<details>
<summary> no error snippets (2x / 16.67%) </summary>

<hr/>

**1x**: _2026-05-26 13:44:11 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17550/pull-kubevirt-e2e-k8s-1.35-sig-storage/2059267975239700480)
<details>
<summary>all...</summary>

* _2026-05-26 13:44:11 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17550/pull-kubevirt-e2e-k8s-1.35-sig-storage/2059267975239700480)

</details>

<hr/>

**1x**: _2026-05-26 08:00:04 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17907/pull-kubevirt-e2e-k8s-1.35-sig-operator/2059182628493332480)
<details>
<summary>all...</summary>

* _2026-05-26 08:00:04 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17907/pull-kubevirt-e2e-k8s-1.35-sig-operator/2059182628493332480)

</details>

<hr/>
</details>

#### internal (5x / 41.67%)

<details>
<summary> make bazel-build target failure (4x / 33.33%) </summary>

<hr/>

**4x**: _2026-05-26 08:00:01 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17907/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059182623372087296#1:build-log.txt%3A200)
<details>
<summary>all...</summary>

* _2026-05-26 08:00:01 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17907/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059182623372087296#1:build-log.txt%3A200)

* _2026-05-26 07:27:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059174345552433152#1:build-log.txt%3A203)

* _2026-05-26 06:09:09 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17759/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059154748795260928#1:build-log.txt%3A200)

* _2026-05-26 01:07:29 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059078799525416960#1:build-log.txt%3A203)

</details>

<hr/>
</details>
<details>
<summary> taint not found during cluster setup (1x / 8.33%) </summary>

<hr/>

**1x**: _2026-05-26 08:14:12 &#43;0000 UTC_: <code>08:22:41: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17830/pull-kubevirt-e2e-kind-sriov/2059184808495419392#1:build-log.txt%3A1431)
<details>
<summary>all...</summary>

* _2026-05-26 08:14:12 &#43;0000 UTC_: <code>08:22:41: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17830/pull-kubevirt-e2e-kind-sriov/2059184808495419392#1:build-log.txt%3A1431)

</details>

<hr/>
</details>

#### external (1x / 8.33%)

<details>
<summary> transient kube-apiserver body decode noise (1x / 8.33%) </summary>

<hr/>

**1x**: _2026-05-26 08:00:04 &#43;0000 UTC_: <code>08:05:25: I0526 04:05:25.182639    1602 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17907/pull-kubevirt-e2e-k8s-1.35-sig-network/2059182625913835520#1:build-log.txt%3A765)
<details>
<summary>all...</summary>

* _2026-05-26 08:00:04 &#43;0000 UTC_: <code>08:05:25: I0526 04:05:25.182639    1602 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17907/pull-kubevirt-e2e-k8s-1.35-sig-network/2059182625913835520#1:build-log.txt%3A765)

</details>

<hr/>
</details>

### 2026-05-25 (16x / 53.33%)


#### needs-investigation (7x / 43.75%)

<details>
<summary> no matching pattern (4x / 25.00%) </summary>

<hr/>

**2x**: _2026-05-25 18:09:04 &#43;0000 UTC_: <code>18:19:50: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-3f2e7223365c31e10f3aa8729b3bad4d3184cd6274f8ce715bd594be2df9e31f`: No space left on device</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058973438944481280#1:build-log.txt%3A641)
<details>
<summary>all...</summary>

* _2026-05-25 18:09:04 &#43;0000 UTC_: <code>18:19:50: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-3f2e7223365c31e10f3aa8729b3bad4d3184cd6274f8ce715bd594be2df9e31f`: No space left on device</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058973438944481280#1:build-log.txt%3A641)

* _2026-05-25 08:14:21 &#43;0000 UTC_: <code>08:25:53: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-e2bf97930c84c036cb4bc89deb6708e780df3ddcdb4f359553f73f210835653b`: No space left on device</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058823736756277248#1:build-log.txt%3A652)

</details>

<hr/>

**2x**: _2026-05-25 12:02:20 &#43;0000 UTC_: <code>12:20:46: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058881203599052800#1:build-log.txt%3A6641)
<details>
<summary>all...</summary>

* _2026-05-25 22:06:15 &#43;0000 UTC_: <code>22:22:52: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2059033219742830592#1:build-log.txt%3A6663)

* _2026-05-25 12:02:20 &#43;0000 UTC_: <code>12:20:46: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058881203599052800#1:build-log.txt%3A6641)

</details>

<hr/>
</details>
<details>
<summary> no error snippets (3x / 18.75%) </summary>

<hr/>

**1x**: _2026-05-25 13:32:59 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2058903178400763904)
<details>
<summary>all...</summary>

* _2026-05-25 13:32:59 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2058903178400763904)

</details>

<hr/>

**1x**: _2026-05-25 13:13:34 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17757/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2058899143815860224)
<details>
<summary>all...</summary>

* _2026-05-25 13:13:34 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17757/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2058899143815860224)

</details>

<hr/>

**1x**: _2026-05-25 12:02:41 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-vgpu/2058881219403190272)
<details>
<summary>all...</summary>

* _2026-05-25 12:02:41 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-vgpu/2058881219403190272)

</details>

<hr/>
</details>

#### external (3x / 18.75%)

<details>
<summary> transient kube-apiserver body decode noise (3x / 18.75%) </summary>

<hr/>

**3x**: _2026-05-25 13:38:47 &#43;0000 UTC_: <code>13:43:50: I0525 09:43:50.177279    1616 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-k8s-1.35-sig-operator/2058903184314732544#1:build-log.txt%3A764)
<details>
<summary>all...</summary>

* _2026-05-25 13:38:47 &#43;0000 UTC_: <code>13:43:50: I0525 09:43:50.177279    1616 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-k8s-1.35-sig-operator/2058903184314732544#1:build-log.txt%3A764)

* _2026-05-25 13:34:07 &#43;0000 UTC_: <code>13:38:07: I0525 09:38:07.664647    1626 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2058903180946706432#1:build-log.txt%3A755)

* _2026-05-25 09:12:21 &#43;0000 UTC_: <code>09:16:10: I0525 05:16:10.064112    1606 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-k8s-1.34-sig-storage/2058830816263278592#1:build-log.txt%3A762)

</details>

<hr/>
</details>

#### internal (6x / 37.50%)

<details>
<summary> make bazel-build target failure (6x / 37.50%) </summary>

<hr/>

**6x**: _2026-05-25 21:06:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059018181887922176#1:build-log.txt%3A205)
<details>
<summary>all...</summary>

* _2026-05-25 21:06:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059018181887922176#1:build-log.txt%3A205)

* _2026-05-25 19:52:10 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058999436138254336#1:build-log.txt%3A212)

* _2026-05-25 17:10:43 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058958831135756288#1:build-log.txt%3A204)

* _2026-05-25 16:57:40 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058955558613946368#1:build-log.txt%3A212)

* _2026-05-25 15:06:25 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058927566504333312#1:build-log.txt%3A205)

* _2026-05-25 13:29:41 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058903179210264576#1:build-log.txt%3A208)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### needs-investigation (15x / 50.00%)

<details>
<summary> no matching pattern (8x / 26.67%) </summary>

<hr/>

**1x**: _2026-05-26 18:24:41 &#43;0000 UTC_: <code>20:51:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-slpvj-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid c86a0bc5-07ae-4204-a46c-f89c66b191b5]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-05-26T20:50:37.865872Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-05-26T20:50:37.903851Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17919/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2059339837684584448#1:build-log.txt%3A8998)
<details>
<summary>all...</summary>

* _2026-05-26 18:24:41 &#43;0000 UTC_: <code>20:51:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-slpvj-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid c86a0bc5-07ae-4204-a46c-f89c66b191b5]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-05-26T20:50:37.865872Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-05-26T20:50:37.903851Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17919/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2059339837684584448#1:build-log.txt%3A8998)
<details><summary>context</summary>
<pre>20:51:02:
20:51:02:   Captured StdOut/StdErr Output &gt;&gt;
20:51:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-slpvj-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid c86a0bc5-07ae-4204-a46c-f89c66b191b5]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;Event(v1.ObjectReference{Kind:\&#34;VirtualMachineInstance\&#34;, Namespace:\&#34;kubevirt-test-default3\&#34;, Name:\&#34;testvmi-slpvj-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\&#34;, UID:\&#34;c86a0bc5-07ae-4204-a46c-f89c66b191b5\&#34;, APIVersion:\&#34;kubevirt.io/v1\&#34;, ResourceVersion:\&#34;133481\&#34;, FieldPath:\&#34;\&#34;}): type: &#39;Normal&#39; reason: &#39;SuccessfulCreate&#39; Created virtual machine pod virt-launcher-testvmi-slpvj-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxrqxxb&#34;,&#34;pos&#34;:&#34;watcher.go:150&#34;,&#34;timestamp&#34;:&#34;2026-05-26T20:50:34.380220Z&#34;}
20:51:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-slpvj-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid c86a0bc5-07ae-4204-a46c-f89c66b191b5]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-05-26T20:50:37.865872Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-05-26T20:50:37.903851Z&#34;}
20:51:02:   &lt;&lt; Captured StdOut/StdErr Output
20:51:02: ------------------------------
20:51:37: • [40.121 seconds]</pre>
</details>


</details>

<hr/>

**5x**: _2026-05-25 12:02:20 &#43;0000 UTC_: <code>12:20:46: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058881203599052800#1:build-log.txt%3A6641)
<details>
<summary>all...</summary>

* _2026-05-26 04:07:33 &#43;0000 UTC_: <code>04:23:46: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059124119928049664#1:build-log.txt%3A6615)
<details><summary>context</summary>
<pre>04:23:46: 	- &#39;systemctl status kubelet&#39;
04:23:46: 	- &#39;journalctl -xeu kubelet&#39;
04:23:46:
04:23:46: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded
04:23:46:
04:23:46: k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow.(*Runner).Run.func1
04:23:46: 	k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow/runner.go:262</pre>
</details>


* _2026-05-26 02:07:30 &#43;0000 UTC_: <code>02:23:47: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2059093907123212288#1:build-log.txt%3A6643)
<details><summary>context</summary>
<pre>02:23:47: 	- &#39;systemctl status kubelet&#39;
02:23:47: 	- &#39;journalctl -xeu kubelet&#39;
02:23:47:
02:23:47: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded
02:23:47:
02:23:47: k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow.(*Runner).Run.func1
02:23:47: 	k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow/runner.go:262</pre>
</details>


* _2026-05-26 00:07:20 &#43;0000 UTC_: <code>00:23:57: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: dial tcp 127.0.0.1:10248: connect: connection refused</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059063686630215680#1:build-log.txt%3A6633)
<details><summary>context</summary>
<pre>00:23:57: 	- &#39;systemctl status kubelet&#39;
00:23:57: 	- &#39;journalctl -xeu kubelet&#39;
00:23:57:
00:23:57: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: dial tcp 127.0.0.1:10248: connect: connection refused
00:23:57:
00:23:57: k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow.(*Runner).Run.func1
00:23:57: 	k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow/runner.go:262</pre>
</details>


* _2026-05-25 22:06:15 &#43;0000 UTC_: <code>22:22:52: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2059033219742830592#1:build-log.txt%3A6663)
<details><summary>context</summary>
<pre>22:22:52: 	- &#39;systemctl status kubelet&#39;
22:22:52: 	- &#39;journalctl -xeu kubelet&#39;
22:22:52:
22:22:52: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded
22:22:52:
22:22:52: k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow.(*Runner).Run.func1
22:22:52: 	k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow/runner.go:262</pre>
</details>


* _2026-05-25 12:02:20 &#43;0000 UTC_: <code>12:20:46: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058881203599052800#1:build-log.txt%3A6641)
<details><summary>context</summary>
<pre>12:20:46: 	- &#39;systemctl status kubelet&#39;
12:20:46: 	- &#39;journalctl -xeu kubelet&#39;
12:20:46:
12:20:46: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded
12:20:46:
12:20:46: k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow.(*Runner).Run.func1
12:20:46: 	k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow/runner.go:262</pre>
</details>


</details>

<hr/>

**2x**: _2026-05-25 18:09:04 &#43;0000 UTC_: <code>18:19:50: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-3f2e7223365c31e10f3aa8729b3bad4d3184cd6274f8ce715bd594be2df9e31f`: No space left on device</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058973438944481280#1:build-log.txt%3A641)
<details>
<summary>all...</summary>

* _2026-05-25 18:09:04 &#43;0000 UTC_: <code>18:19:50: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-3f2e7223365c31e10f3aa8729b3bad4d3184cd6274f8ce715bd594be2df9e31f`: No space left on device</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058973438944481280#1:build-log.txt%3A641)
<details><summary>context</summary>
<pre>18:19:50:  ✗ Preparing nodes 📦 📦
18:19:50: ERROR: failed to create cluster: command &#34;podman run --name kind-1.35-worker --hostname kind-1.35-worker --label io.x-k8s.kind.role=worker --privileged --tmpfs /tmp --tmpfs /run --volume 44a1f40a1bfe0ea43c5f2787872ff4101f1ffc9c4e211135498f08c6230f4566:/var:suid,exec,dev --volume /lib/modules:/lib/modules:ro -e KIND_EXPERIMENTAL_CONTAINERD_SNAPSHOTTER --detach --tty --net kind --label io.x-k8s.kind.cluster=kind-1.35 -e container=podman --cgroupns=private --volume /dev/mapper:/dev/mapper --volume=/var/log/audit:/var/log/audit:ro docker.io/kindest/node@sha256:4613778f3cfcd10e615029370f5786704559103cf27bef934597ba562b269661&#34; failed with error: exit status 126
18:19:50:
18:19:50: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-3f2e7223365c31e10f3aa8729b3bad4d3184cd6274f8ce715bd594be2df9e31f`: No space left on device
18:19:50:
18:19:50: Stack Trace:
18:19:50: sigs.k8s.io/kind/pkg/errors.WithStack</pre>
</details>


* _2026-05-25 08:14:21 &#43;0000 UTC_: <code>08:25:53: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-e2bf97930c84c036cb4bc89deb6708e780df3ddcdb4f359553f73f210835653b`: No space left on device</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058823736756277248#1:build-log.txt%3A652)
<details><summary>context</summary>
<pre>08:25:53:  ✗ Preparing nodes 📦 📦
08:25:53: ERROR: failed to create cluster: command &#34;podman run --name kind-1.34-control-plane --hostname kind-1.34-control-plane --label io.x-k8s.kind.role=control-plane --privileged --tmpfs /tmp --tmpfs /run --volume e972a0936de444675a134034826b6ab581cc5cb5273dd45e9252f096913c1e7b:/var:suid,exec,dev --volume /lib/modules:/lib/modules:ro -e KIND_EXPERIMENTAL_CONTAINERD_SNAPSHOTTER --detach --tty --net kind --label io.x-k8s.kind.cluster=kind-1.34 -e container=podman --cgroupns=private --volume /dev/mapper:/dev/mapper --volume=/var/log/audit:/var/log/audit:ro --publish=127.0.0.1:34443:6443/tcp -e KUBECONFIG=/etc/kubernetes/admin.conf docker.io/kindest/node@sha256:08497ee19eace7b4b5348db5c6a1591d7752b164530a36f855cb0f2bdcbadd48&#34; failed with error: exit status 126
08:25:53:
08:25:53: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-e2bf97930c84c036cb4bc89deb6708e780df3ddcdb4f359553f73f210835653b`: No space left on device
08:25:53:
08:25:53: Stack Trace:
08:25:53: sigs.k8s.io/kind/pkg/errors.WithStack</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> no error snippets (7x / 23.33%) </summary>

<hr/>

**1x**: _2026-05-28 09:40:36 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17942/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059932689309372416)
<details>
<summary>all...</summary>

* _2026-05-28 09:40:36 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17942/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059932689309372416)

</details>

<hr/>

**1x**: _2026-05-27 10:39:00 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17911/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059584804717858816)
<details>
<summary>all...</summary>

* _2026-05-27 10:39:00 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17911/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059584804717858816)

</details>

<hr/>

**1x**: _2026-05-26 13:44:11 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17550/pull-kubevirt-e2e-k8s-1.35-sig-storage/2059267975239700480)
<details>
<summary>all...</summary>

* _2026-05-26 13:44:11 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17550/pull-kubevirt-e2e-k8s-1.35-sig-storage/2059267975239700480)

</details>

<hr/>

**1x**: _2026-05-26 08:00:04 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17907/pull-kubevirt-e2e-k8s-1.35-sig-operator/2059182628493332480)
<details>
<summary>all...</summary>

* _2026-05-26 08:00:04 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17907/pull-kubevirt-e2e-k8s-1.35-sig-operator/2059182628493332480)

</details>

<hr/>

**1x**: _2026-05-25 13:32:59 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2058903178400763904)
<details>
<summary>all...</summary>

* _2026-05-25 13:32:59 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2058903178400763904)

</details>

<hr/>

**1x**: _2026-05-25 13:13:34 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17757/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2058899143815860224)
<details>
<summary>all...</summary>

* _2026-05-25 13:13:34 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17757/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2058899143815860224)

</details>

<hr/>

**1x**: _2026-05-25 12:02:41 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-vgpu/2058881219403190272)
<details>
<summary>all...</summary>

* _2026-05-25 12:02:41 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-vgpu/2058881219403190272)

</details>

<hr/>
</details>

### internal (11x / 36.67%)

<details>
<summary> make bazel-build target failure (10x / 33.33%) </summary>

<hr/>

**10x**: _2026-05-26 08:00:01 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17907/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059182623372087296#1:build-log.txt%3A200)
<details>
<summary>all...</summary>

* _2026-05-26 08:00:01 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17907/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059182623372087296#1:build-log.txt%3A200)
<details><summary>context</summary>
<pre>&#43; make bazel-build-images
hack/dockerized &#34;export BUILD_ARCH= &amp;&amp; DOCKER_PREFIX= DOCKER_TAG= DOCKER_TAG_ALT= IMAGE_PREFIX= IMAGE_PREFIX_ALT=kv- ./hack/multi-arch.sh build-images&#34;
08:00:34: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-3ade505c57e218819e3b2588636252484d69f1c39a5d0b601f611893f13dff83`: No space left on device
make: *** [Makefile:39: bazel-build-images] Error 126
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-05-26 07:27:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059174345552433152#1:build-log.txt%3A203)
<details><summary>context</summary>
<pre>&#43; make bazel-build-images
hack/dockerized &#34;export BUILD_ARCH= &amp;&amp; DOCKER_PREFIX= DOCKER_TAG= DOCKER_TAG_ALT= IMAGE_PREFIX= IMAGE_PREFIX_ALT=kv- ./hack/multi-arch.sh build-images&#34;
07:27:31: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-efe2b228ff2a1ce1d95aab1fb13854b583e8e74cebd8c3956352cd52d22d5414`: No space left on device
make: *** [Makefile:39: bazel-build-images] Error 126
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-05-26 06:09:09 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17759/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059154748795260928#1:build-log.txt%3A200)
<details><summary>context</summary>
<pre>&#43; make bazel-build-images
hack/dockerized &#34;export BUILD_ARCH= &amp;&amp; DOCKER_PREFIX= DOCKER_TAG= DOCKER_TAG_ALT= IMAGE_PREFIX= IMAGE_PREFIX_ALT=kv- ./hack/multi-arch.sh build-images&#34;
06:09:38: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-d8b008396422eaa941cfdb4eeaeab217298aa937b4d801bf1374a8fcde5537a3`: No space left on device
make: *** [Makefile:39: bazel-build-images] Error 126
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-05-26 01:07:29 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059078799525416960#1:build-log.txt%3A203)
<details><summary>context</summary>
<pre>&#43; make bazel-build-images
hack/dockerized &#34;export BUILD_ARCH= &amp;&amp; DOCKER_PREFIX= DOCKER_TAG= DOCKER_TAG_ALT= IMAGE_PREFIX= IMAGE_PREFIX_ALT=kv- ./hack/multi-arch.sh build-images&#34;
01:07:58: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-90a0783b2881336943496bef956125c62eaf5e7261df0077ef92209fc1315874`: No space left on device
make: *** [Makefile:39: bazel-build-images] Error 126
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-05-25 21:06:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059018181887922176#1:build-log.txt%3A205)
<details><summary>context</summary>
<pre>&#43; make bazel-build-images
hack/dockerized &#34;export BUILD_ARCH= &amp;&amp; DOCKER_PREFIX= DOCKER_TAG= DOCKER_TAG_ALT= IMAGE_PREFIX= IMAGE_PREFIX_ALT=kv- ./hack/multi-arch.sh build-images&#34;
21:07:01: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-5f5a9b56d6e99ccf443eb29568f3e995c544703b2d8074653773ef2f790040f8`: No space left on device
make: *** [Makefile:39: bazel-build-images] Error 126
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-05-25 19:52:10 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058999436138254336#1:build-log.txt%3A212)
<details><summary>context</summary>
<pre>&#43; make bazel-build-images
hack/dockerized &#34;export BUILD_ARCH= &amp;&amp; DOCKER_PREFIX= DOCKER_TAG= DOCKER_TAG_ALT= IMAGE_PREFIX= IMAGE_PREFIX_ALT=kv- ./hack/multi-arch.sh build-images&#34;
19:52:39: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-3b6963538642576dc5e085a242ae8cf52a73f2621d3b14cf4115284e2b9e3de1`: No space left on device
make: *** [Makefile:39: bazel-build-images] Error 126
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-05-25 17:10:43 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058958831135756288#1:build-log.txt%3A204)
<details><summary>context</summary>
<pre>&#43; make bazel-build-images
hack/dockerized &#34;export BUILD_ARCH= &amp;&amp; DOCKER_PREFIX= DOCKER_TAG= DOCKER_TAG_ALT= IMAGE_PREFIX= IMAGE_PREFIX_ALT=kv- ./hack/multi-arch.sh build-images&#34;
17:11:16: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-b0d249f272298390f6f04648e0394a2014389781537f15f35b3ee15c3c28c9e7`: No space left on device
make: *** [Makefile:39: bazel-build-images] Error 126
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-05-25 16:57:40 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058955558613946368#1:build-log.txt%3A212)
<details><summary>context</summary>
<pre>&#43; make bazel-build-images
hack/dockerized &#34;export BUILD_ARCH= &amp;&amp; DOCKER_PREFIX= DOCKER_TAG= DOCKER_TAG_ALT= IMAGE_PREFIX= IMAGE_PREFIX_ALT=kv- ./hack/multi-arch.sh build-images&#34;
16:58:09: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-50696fa3ed6b5067c2a55c35d8faaec48ad10fd9f74941d9ddbd6820ae8cf086`: No space left on device
make: *** [Makefile:39: bazel-build-images] Error 126
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-05-25 15:06:25 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058927566504333312#1:build-log.txt%3A205)
<details><summary>context</summary>
<pre>&#43; make bazel-build-images
hack/dockerized &#34;export BUILD_ARCH= &amp;&amp; DOCKER_PREFIX= DOCKER_TAG= DOCKER_TAG_ALT= IMAGE_PREFIX= IMAGE_PREFIX_ALT=kv- ./hack/multi-arch.sh build-images&#34;
15:06:53: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-db258cc59e9e8630f72efb71bd85054c350d77722b1efbc478f8f2f89f9f2555`: No space left on device
make: *** [Makefile:39: bazel-build-images] Error 126
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-05-25 13:29:41 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058903179210264576#1:build-log.txt%3A208)
<details><summary>context</summary>
<pre>&#43; make bazel-build-images
hack/dockerized &#34;export BUILD_ARCH= &amp;&amp; DOCKER_PREFIX= DOCKER_TAG= DOCKER_TAG_ALT= IMAGE_PREFIX= IMAGE_PREFIX_ALT=kv- ./hack/multi-arch.sh build-images&#34;
13:31:23: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-d049bac7d5d79e49ab46708c1555ab6646f6baafd2c46104cee6e1455b1e0476`: No space left on device
make: *** [Makefile:39: bazel-build-images] Error 126
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> taint not found during cluster setup (1x / 3.33%) </summary>

<hr/>

**1x**: _2026-05-26 08:14:12 &#43;0000 UTC_: <code>08:22:41: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17830/pull-kubevirt-e2e-kind-sriov/2059184808495419392#1:build-log.txt%3A1431)
<details>
<summary>all...</summary>

* _2026-05-26 08:14:12 &#43;0000 UTC_: <code>08:22:41: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17830/pull-kubevirt-e2e-kind-sriov/2059184808495419392#1:build-log.txt%3A1431)
<details><summary>context</summary>
<pre>08:22:41: &#43;&#43; for node in ${master_nodes[@]}
08:22:41: &#43;&#43; _kubectl taint nodes sriov-control-plane node-role.kubernetes.io/master:NoSchedule-
08:22:41: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-sriov/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-sriov/.kubeconfig taint nodes sriov-control-plane node-role.kubernetes.io/master:NoSchedule-
08:22:41: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found
08:22:41: &#43;&#43; _kubectl taint nodes sriov-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
08:22:41: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-sriov/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-sriov/.kubeconfig taint nodes sriov-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
08:22:41: node/sriov-control-plane untainted</pre>
</details>


</details>

<hr/>
</details>

### external (4x / 13.33%)

<details>
<summary> transient kube-apiserver body decode noise (4x / 13.33%) </summary>

<hr/>

**4x**: _2026-05-25 13:38:47 &#43;0000 UTC_: <code>13:43:50: I0525 09:43:50.177279    1616 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-k8s-1.35-sig-operator/2058903184314732544#1:build-log.txt%3A764)
<details>
<summary>all...</summary>

* _2026-05-26 08:00:04 &#43;0000 UTC_: <code>08:05:25: I0526 04:05:25.182639    1602 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17907/pull-kubevirt-e2e-k8s-1.35-sig-network/2059182625913835520#1:build-log.txt%3A765)
<details><summary>context</summary>
<pre>08:05:23: [control-plane-check] kube-scheduler is healthy after 2.125371382s
08:05:24: I0526 04:05:24.183164    1602 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;
08:05:24: I0526 04:05:24.682687    1602 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;
08:05:25: I0526 04:05:25.182639    1602 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
08:05:25: I0526 04:05:25.182693    1602 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
08:05:25: [control-plane-check] kube-apiserver is healthy after 4.002363218s
08:05:25: I0526 04:05:25.685619    1602 kubeconfig.go:657] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists</pre>
</details>


* _2026-05-25 13:38:47 &#43;0000 UTC_: <code>13:43:50: I0525 09:43:50.177279    1616 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-k8s-1.35-sig-operator/2058903184314732544#1:build-log.txt%3A764)
<details><summary>context</summary>
<pre>13:43:48: I0525 09:43:48.675817    1616 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;
13:43:49: I0525 09:43:49.175326    1616 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;
13:43:49: I0525 09:43:49.675057    1616 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;
13:43:50: I0525 09:43:50.177279    1616 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
13:43:50: I0525 09:43:50.177368    1616 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
13:43:50: [control-plane-check] kube-apiserver is healthy after 4.004390628s
13:43:50: I0525 09:43:50.681968    1616 kubeconfig.go:657] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists</pre>
</details>


* _2026-05-25 13:34:07 &#43;0000 UTC_: <code>13:38:07: I0525 09:38:07.664647    1626 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2058903180946706432#1:build-log.txt%3A755)
<details><summary>context</summary>
<pre>13:38:06: [control-plane-check] kube-scheduler is healthy after 2.654985185s
13:38:06: I0525 09:38:06.663665    1626 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;
13:38:07: I0525 09:38:07.165095    1626 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;
13:38:07: I0525 09:38:07.664647    1626 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
13:38:07: I0525 09:38:07.664706    1626 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
13:38:08: [control-plane-check] kube-apiserver is healthy after 4.501516165s
13:38:08: I0525 09:38:08.165936    1626 kubeconfig.go:657] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists</pre>
</details>


* _2026-05-25 09:12:21 &#43;0000 UTC_: <code>09:16:10: I0525 05:16:10.064112    1606 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-k8s-1.34-sig-storage/2058830816263278592#1:build-log.txt%3A762)
<details><summary>context</summary>
<pre>09:16:06: [control-plane-check] Checking kube-scheduler at https://127.0.0.1:10259/livez
09:16:07: [control-plane-check] kube-controller-manager is healthy after 1.092449818s
09:16:08: [control-plane-check] kube-scheduler is healthy after 1.978771755s
09:16:10: I0525 05:16:10.064112    1606 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
09:16:10: [control-plane-check] kube-apiserver is healthy after 4.001529214s
09:16:10: I0525 05:16:10.566861    1606 kubeconfig.go:657] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists
09:16:10: I0525 05:16:10.569311    1606 kubeconfig.go:730] creating the ClusterRoleBinding for the kubeadm:cluster-admins Group by using super-admin.conf</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (25x / 83.33%)


#### internal (10x / 40.00%)

<details>
<summary> make bazel-build target failure (9x / 36.00%) </summary>

<hr/>

**9x**: _2026-05-26 08:00:01 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17907/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059182623372087296#1:build-log.txt%3A200)
<details>
<summary>all...</summary>

* _2026-05-26 08:00:01 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17907/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059182623372087296#1:build-log.txt%3A200)

* _2026-05-26 07:27:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059174345552433152#1:build-log.txt%3A203)

* _2026-05-26 06:09:09 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17759/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059154748795260928#1:build-log.txt%3A200)

* _2026-05-26 01:07:29 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059078799525416960#1:build-log.txt%3A203)

* _2026-05-25 21:06:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059018181887922176#1:build-log.txt%3A205)

* _2026-05-25 19:52:10 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058999436138254336#1:build-log.txt%3A212)

* _2026-05-25 16:57:40 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058955558613946368#1:build-log.txt%3A212)

* _2026-05-25 15:06:25 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058927566504333312#1:build-log.txt%3A205)

* _2026-05-25 13:29:41 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058903179210264576#1:build-log.txt%3A208)

</details>

<hr/>
</details>
<details>
<summary> taint not found during cluster setup (1x / 4.00%) </summary>

<hr/>

**1x**: _2026-05-26 08:14:12 &#43;0000 UTC_: <code>08:22:41: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17830/pull-kubevirt-e2e-kind-sriov/2059184808495419392#1:build-log.txt%3A1431)
<details>
<summary>all...</summary>

* _2026-05-26 08:14:12 &#43;0000 UTC_: <code>08:22:41: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17830/pull-kubevirt-e2e-kind-sriov/2059184808495419392#1:build-log.txt%3A1431)

</details>

<hr/>
</details>

#### needs-investigation (11x / 44.00%)

<details>
<summary> no error snippets (7x / 28.00%) </summary>

<hr/>

**1x**: _2026-05-28 09:40:36 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17942/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059932689309372416)
<details>
<summary>all...</summary>

* _2026-05-28 09:40:36 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17942/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059932689309372416)

</details>

<hr/>

**1x**: _2026-05-27 10:39:00 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17911/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059584804717858816)
<details>
<summary>all...</summary>

* _2026-05-27 10:39:00 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17911/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059584804717858816)

</details>

<hr/>

**1x**: _2026-05-26 13:44:11 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17550/pull-kubevirt-e2e-k8s-1.35-sig-storage/2059267975239700480)
<details>
<summary>all...</summary>

* _2026-05-26 13:44:11 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17550/pull-kubevirt-e2e-k8s-1.35-sig-storage/2059267975239700480)

</details>

<hr/>

**1x**: _2026-05-26 08:00:04 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17907/pull-kubevirt-e2e-k8s-1.35-sig-operator/2059182628493332480)
<details>
<summary>all...</summary>

* _2026-05-26 08:00:04 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17907/pull-kubevirt-e2e-k8s-1.35-sig-operator/2059182628493332480)

</details>

<hr/>

**1x**: _2026-05-25 13:32:59 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2058903178400763904)
<details>
<summary>all...</summary>

* _2026-05-25 13:32:59 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2058903178400763904)

</details>

<hr/>

**1x**: _2026-05-25 13:13:34 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17757/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2058899143815860224)
<details>
<summary>all...</summary>

* _2026-05-25 13:13:34 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17757/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2058899143815860224)

</details>

<hr/>

**1x**: _2026-05-25 12:02:41 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-vgpu/2058881219403190272)
<details>
<summary>all...</summary>

* _2026-05-25 12:02:41 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-vgpu/2058881219403190272)

</details>

<hr/>
</details>
<details>
<summary> no matching pattern (4x / 16.00%) </summary>

<hr/>

**3x**: _2026-05-25 12:02:20 &#43;0000 UTC_: <code>12:20:46: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058881203599052800#1:build-log.txt%3A6641)
<details>
<summary>all...</summary>

* _2026-05-26 04:07:33 &#43;0000 UTC_: <code>04:23:46: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059124119928049664#1:build-log.txt%3A6615)

* _2026-05-26 00:07:20 &#43;0000 UTC_: <code>00:23:57: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: dial tcp 127.0.0.1:10248: connect: connection refused</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059063686630215680#1:build-log.txt%3A6633)

* _2026-05-25 12:02:20 &#43;0000 UTC_: <code>12:20:46: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058881203599052800#1:build-log.txt%3A6641)

</details>

<hr/>

**1x**: _2026-05-25 18:09:04 &#43;0000 UTC_: <code>18:19:50: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-3f2e7223365c31e10f3aa8729b3bad4d3184cd6274f8ce715bd594be2df9e31f`: No space left on device</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058973438944481280#1:build-log.txt%3A641)
<details>
<summary>all...</summary>

* _2026-05-25 18:09:04 &#43;0000 UTC_: <code>18:19:50: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-3f2e7223365c31e10f3aa8729b3bad4d3184cd6274f8ce715bd594be2df9e31f`: No space left on device</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058973438944481280#1:build-log.txt%3A641)

</details>

<hr/>
</details>

#### external (4x / 16.00%)

<details>
<summary> transient kube-apiserver body decode noise (4x / 16.00%) </summary>

<hr/>

**4x**: _2026-05-25 13:38:47 &#43;0000 UTC_: <code>13:43:50: I0525 09:43:50.177279    1616 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-k8s-1.35-sig-operator/2058903184314732544#1:build-log.txt%3A764)
<details>
<summary>all...</summary>

* _2026-05-26 08:00:04 &#43;0000 UTC_: <code>08:05:25: I0526 04:05:25.182639    1602 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17907/pull-kubevirt-e2e-k8s-1.35-sig-network/2059182625913835520#1:build-log.txt%3A765)

* _2026-05-25 13:38:47 &#43;0000 UTC_: <code>13:43:50: I0525 09:43:50.177279    1616 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-k8s-1.35-sig-operator/2058903184314732544#1:build-log.txt%3A764)

* _2026-05-25 13:34:07 &#43;0000 UTC_: <code>13:38:07: I0525 09:38:07.664647    1626 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2058903180946706432#1:build-log.txt%3A755)

* _2026-05-25 09:12:21 &#43;0000 UTC_: <code>09:16:10: I0525 05:16:10.064112    1606 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-k8s-1.34-sig-storage/2058830816263278592#1:build-log.txt%3A762)

</details>

<hr/>
</details>

### release-1.8 (5x / 16.67%)


#### internal (1x / 20.00%)

<details>
<summary> make bazel-build target failure (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-05-25 17:10:43 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058958831135756288#1:build-log.txt%3A204)
<details>
<summary>all...</summary>

* _2026-05-25 17:10:43 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058958831135756288#1:build-log.txt%3A204)

</details>

<hr/>
</details>

#### needs-investigation (4x / 80.00%)

<details>
<summary> no matching pattern (4x / 80.00%) </summary>

<hr/>

**2x**: _2026-05-26 02:07:30 &#43;0000 UTC_: <code>02:23:47: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2059093907123212288#1:build-log.txt%3A6643)
<details>
<summary>all...</summary>

* _2026-05-26 02:07:30 &#43;0000 UTC_: <code>02:23:47: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2059093907123212288#1:build-log.txt%3A6643)

* _2026-05-25 22:06:15 &#43;0000 UTC_: <code>22:22:52: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2059033219742830592#1:build-log.txt%3A6663)

</details>

<hr/>

**1x**: _2026-05-26 18:24:41 &#43;0000 UTC_: <code>20:51:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-slpvj-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid c86a0bc5-07ae-4204-a46c-f89c66b191b5]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-05-26T20:50:37.865872Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-05-26T20:50:37.903851Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17919/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2059339837684584448#1:build-log.txt%3A8998)
<details>
<summary>all...</summary>

* _2026-05-26 18:24:41 &#43;0000 UTC_: <code>20:51:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-slpvj-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid c86a0bc5-07ae-4204-a46c-f89c66b191b5]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-05-26T20:50:37.865872Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-05-26T20:50:37.903851Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17919/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2059339837684584448#1:build-log.txt%3A8998)

</details>

<hr/>

**1x**: _2026-05-25 08:14:21 &#43;0000 UTC_: <code>08:25:53: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-e2bf97930c84c036cb4bc89deb6708e780df3ddcdb4f359553f73f210835653b`: No space left on device</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058823736756277248#1:build-log.txt%3A652)
<details>
<summary>all...</summary>

* _2026-05-25 08:14:21 &#43;0000 UTC_: <code>08:25:53: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-e2bf97930c84c036cb4bc89deb6708e780df3ddcdb4f359553f73f210835653b`: No space left on device</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058823736756277248#1:build-log.txt%3A652)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-storage (5x / 16.67%)


#### needs-investigation (4x / 80.00%)

<details>
<summary> no error snippets (3x / 60.00%) </summary>

<hr/>

**1x**: _2026-05-28 09:40:36 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17942/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059932689309372416)
<details>
<summary>all...</summary>

* _2026-05-28 09:40:36 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17942/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059932689309372416)

</details>

<hr/>

**1x**: _2026-05-27 10:39:00 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17911/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059584804717858816)
<details>
<summary>all...</summary>

* _2026-05-27 10:39:00 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17911/pull-kubevirt-e2e-k8s-1.34-sig-storage/2059584804717858816)

</details>

<hr/>

**1x**: _2026-05-26 13:44:11 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17550/pull-kubevirt-e2e-k8s-1.35-sig-storage/2059267975239700480)
<details>
<summary>all...</summary>

* _2026-05-26 13:44:11 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17550/pull-kubevirt-e2e-k8s-1.35-sig-storage/2059267975239700480)

</details>

<hr/>
</details>
<details>
<summary> no matching pattern (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-05-26 18:24:41 &#43;0000 UTC_: <code>20:51:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-slpvj-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid c86a0bc5-07ae-4204-a46c-f89c66b191b5]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-05-26T20:50:37.865872Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-05-26T20:50:37.903851Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17919/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2059339837684584448#1:build-log.txt%3A8998)
<details>
<summary>all...</summary>

* _2026-05-26 18:24:41 &#43;0000 UTC_: <code>20:51:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-slpvj-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid c86a0bc5-07ae-4204-a46c-f89c66b191b5]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-05-26T20:50:37.865872Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-05-26T20:50:37.903851Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17919/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2059339837684584448#1:build-log.txt%3A8998)

</details>

<hr/>
</details>

#### external (1x / 20.00%)

<details>
<summary> transient kube-apiserver body decode noise (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-05-25 09:12:21 &#43;0000 UTC_: <code>09:16:10: I0525 05:16:10.064112    1606 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-k8s-1.34-sig-storage/2058830816263278592#1:build-log.txt%3A762)
<details>
<summary>all...</summary>

* _2026-05-25 09:12:21 &#43;0000 UTC_: <code>09:16:10: I0525 05:16:10.064112    1606 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-k8s-1.34-sig-storage/2058830816263278592#1:build-log.txt%3A762)

</details>

<hr/>
</details>

### sig-network (3x / 10.00%)


#### internal (1x / 33.33%)

<details>
<summary> taint not found during cluster setup (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-05-26 08:14:12 &#43;0000 UTC_: <code>08:22:41: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17830/pull-kubevirt-e2e-kind-sriov/2059184808495419392#1:build-log.txt%3A1431)
<details>
<summary>all...</summary>

* _2026-05-26 08:14:12 &#43;0000 UTC_: <code>08:22:41: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17830/pull-kubevirt-e2e-kind-sriov/2059184808495419392#1:build-log.txt%3A1431)

</details>

<hr/>
</details>

#### needs-investigation (1x / 33.33%)

<details>
<summary> no error snippets (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-05-25 13:13:34 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17757/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2058899143815860224)
<details>
<summary>all...</summary>

* _2026-05-25 13:13:34 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17757/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2058899143815860224)

</details>

<hr/>
</details>

#### external (1x / 33.33%)

<details>
<summary> transient kube-apiserver body decode noise (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-05-26 08:00:04 &#43;0000 UTC_: <code>08:05:25: I0526 04:05:25.182639    1602 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17907/pull-kubevirt-e2e-k8s-1.35-sig-network/2059182625913835520#1:build-log.txt%3A765)
<details>
<summary>all...</summary>

* _2026-05-26 08:00:04 &#43;0000 UTC_: <code>08:05:25: I0526 04:05:25.182639    1602 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17907/pull-kubevirt-e2e-k8s-1.35-sig-network/2059182625913835520#1:build-log.txt%3A765)

</details>

<hr/>
</details>

### sig-compute (22x / 73.33%)


#### needs-investigation (10x / 45.45%)

<details>
<summary> no matching pattern (7x / 31.82%) </summary>

<hr/>

**5x**: _2026-05-25 12:02:20 &#43;0000 UTC_: <code>12:20:46: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058881203599052800#1:build-log.txt%3A6641)
<details>
<summary>all...</summary>

* _2026-05-26 04:07:33 &#43;0000 UTC_: <code>04:23:46: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059124119928049664#1:build-log.txt%3A6615)

* _2026-05-26 02:07:30 &#43;0000 UTC_: <code>02:23:47: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2059093907123212288#1:build-log.txt%3A6643)

* _2026-05-26 00:07:20 &#43;0000 UTC_: <code>00:23:57: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: dial tcp 127.0.0.1:10248: connect: connection refused</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059063686630215680#1:build-log.txt%3A6633)

* _2026-05-25 22:06:15 &#43;0000 UTC_: <code>22:22:52: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2059033219742830592#1:build-log.txt%3A6663)

* _2026-05-25 12:02:20 &#43;0000 UTC_: <code>12:20:46: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058881203599052800#1:build-log.txt%3A6641)

</details>

<hr/>

**2x**: _2026-05-25 18:09:04 &#43;0000 UTC_: <code>18:19:50: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-3f2e7223365c31e10f3aa8729b3bad4d3184cd6274f8ce715bd594be2df9e31f`: No space left on device</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058973438944481280#1:build-log.txt%3A641)
<details>
<summary>all...</summary>

* _2026-05-25 18:09:04 &#43;0000 UTC_: <code>18:19:50: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-3f2e7223365c31e10f3aa8729b3bad4d3184cd6274f8ce715bd594be2df9e31f`: No space left on device</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058973438944481280#1:build-log.txt%3A641)

* _2026-05-25 08:14:21 &#43;0000 UTC_: <code>08:25:53: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-e2bf97930c84c036cb4bc89deb6708e780df3ddcdb4f359553f73f210835653b`: No space left on device</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058823736756277248#1:build-log.txt%3A652)

</details>

<hr/>
</details>
<details>
<summary> no error snippets (3x / 13.64%) </summary>

<hr/>

**1x**: _2026-05-26 08:00:04 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17907/pull-kubevirt-e2e-k8s-1.35-sig-operator/2059182628493332480)
<details>
<summary>all...</summary>

* _2026-05-26 08:00:04 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17907/pull-kubevirt-e2e-k8s-1.35-sig-operator/2059182628493332480)

</details>

<hr/>

**1x**: _2026-05-25 13:32:59 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2058903178400763904)
<details>
<summary>all...</summary>

* _2026-05-25 13:32:59 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2058903178400763904)

</details>

<hr/>

**1x**: _2026-05-25 12:02:41 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-vgpu/2058881219403190272)
<details>
<summary>all...</summary>

* _2026-05-25 12:02:41 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-vgpu/2058881219403190272)

</details>

<hr/>
</details>

#### external (2x / 9.09%)

<details>
<summary> transient kube-apiserver body decode noise (2x / 9.09%) </summary>

<hr/>

**2x**: _2026-05-25 13:38:47 &#43;0000 UTC_: <code>13:43:50: I0525 09:43:50.177279    1616 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-k8s-1.35-sig-operator/2058903184314732544#1:build-log.txt%3A764)
<details>
<summary>all...</summary>

* _2026-05-25 13:38:47 &#43;0000 UTC_: <code>13:43:50: I0525 09:43:50.177279    1616 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-k8s-1.35-sig-operator/2058903184314732544#1:build-log.txt%3A764)

* _2026-05-25 13:34:07 &#43;0000 UTC_: <code>13:38:07: I0525 09:38:07.664647    1626 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2058903180946706432#1:build-log.txt%3A755)

</details>

<hr/>
</details>

#### internal (10x / 45.45%)

<details>
<summary> make bazel-build target failure (10x / 45.45%) </summary>

<hr/>

**10x**: _2026-05-26 08:00:01 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17907/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059182623372087296#1:build-log.txt%3A200)
<details>
<summary>all...</summary>

* _2026-05-26 08:00:01 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17907/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059182623372087296#1:build-log.txt%3A200)

* _2026-05-26 07:27:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059174345552433152#1:build-log.txt%3A203)

* _2026-05-26 06:09:09 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17759/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059154748795260928#1:build-log.txt%3A200)

* _2026-05-26 01:07:29 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059078799525416960#1:build-log.txt%3A203)

* _2026-05-25 21:06:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059018181887922176#1:build-log.txt%3A205)

* _2026-05-25 19:52:10 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058999436138254336#1:build-log.txt%3A212)

* _2026-05-25 17:10:43 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058958831135756288#1:build-log.txt%3A204)

* _2026-05-25 16:57:40 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058955558613946368#1:build-log.txt%3A212)

* _2026-05-25 15:06:25 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17902/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058927566504333312#1:build-log.txt%3A205)

* _2026-05-25 13:29:41 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17900/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058903179210264576#1:build-log.txt%3A208)

</details>

<hr/>
</details>

Last updated: 2026-05-30 12:35:47
