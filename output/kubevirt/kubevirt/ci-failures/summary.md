



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-06-11 (1x / 7.14%)


#### external (1x / 100.00%)

<details>
<summary> HTTP error in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-06-11 08:11:20 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16106/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2064983681084166144#1:build-log.txt%3A796)
<details>
<summary>all...</summary>

* _2026-06-11 08:11:20 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16106/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2064983681084166144#1:build-log.txt%3A796)

</details>

<hr/>
</details>

### 2026-06-09 (3x / 21.43%)


#### internal (2x / 66.67%)

<details>
<summary> make cluster lifecycle target failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-06-09 12:55:48 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2064329652029100032#1:build-log.txt%3A1196)
<details>
<summary>all...</summary>

* _2026-06-09 12:55:48 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2064329652029100032#1:build-log.txt%3A1196)

</details>

<hr/>
</details>
<details>
<summary> make bazel-build target failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-06-09 08:35:56 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2064264517830840320#1:build-log.txt%3A4515)
<details>
<summary>all...</summary>

* _2026-06-09 08:35:56 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2064264517830840320#1:build-log.txt%3A4515)

</details>

<hr/>
</details>

#### external (1x / 33.33%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-06-09 08:33:34 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2064264500244123648#1:build-log.txt%3A4384)
<details>
<summary>all...</summary>

* _2026-06-09 08:33:34 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2064264500244123648#1:build-log.txt%3A4384)

</details>

<hr/>
</details>

### 2026-06-08 (9x / 64.29%)


#### external (8x / 88.89%)

<details>
<summary> download failure in context (5x / 55.56%) </summary>

<hr/>

**4x**: _2026-06-08 09:21:42 &#43;0000 UTC_: <code>09:33:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-operator/2063914210630307840#1:build-log.txt%3A4332)
<details>
<summary>all...</summary>

* _2026-06-08 09:21:43 &#43;0000 UTC_: <code>09:32:04: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.35-sig-operator/2063914211544666112#1:build-log.txt%3A4139)

* _2026-06-08 09:21:42 &#43;0000 UTC_: <code>09:33:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-operator/2063914210630307840#1:build-log.txt%3A4332)

* _2026-06-08 09:21:42 &#43;0000 UTC_: <code>09:35:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.35-sig-compute/2063914211016183808#1:build-log.txt%3A4309)

* _2026-06-08 09:21:42 &#43;0000 UTC_: <code>09:34:43: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.35-sig-storage/2063914210865188864#1:build-log.txt%3A4144)

</details>

<hr/>

**1x**: _2026-06-08 09:21:39 &#43;0000 UTC_: <code>09:23:47: ERROR: Analysis of target &#39;//containerimages:fedora-realtime&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-storage/2063914210416398336#1:build-log.txt%3A845)
<details>
<summary>all...</summary>

* _2026-06-08 09:21:39 &#43;0000 UTC_: <code>09:23:47: ERROR: Analysis of target &#39;//containerimages:fedora-realtime&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-storage/2063914210416398336#1:build-log.txt%3A845)

</details>

<hr/>
</details>
<details>
<summary> failed external fetch in context (2x / 22.22%) </summary>

<hr/>

**1x**: _2026-06-08 09:21:41 &#43;0000 UTC_: <code>09:31:45: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-compute/2063914210517061632#1:build-log.txt%3A3995)
<details>
<summary>all...</summary>

* _2026-06-08 09:21:41 &#43;0000 UTC_: <code>09:31:45: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-compute/2063914210517061632#1:build-log.txt%3A3995)

</details>

<hr/>

**1x**: _2026-06-08 09:21:39 &#43;0000 UTC_: <code>09:24:33: ERROR: Analysis of target &#39;//containerimages:fedora-with-test-tooling&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-kind-1.35-vgpu/2063914209959219200#1:build-log.txt%3A500)
<details>
<summary>all...</summary>

* _2026-06-08 09:21:39 &#43;0000 UTC_: <code>09:24:33: ERROR: Analysis of target &#39;//containerimages:fedora-with-test-tooling&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-kind-1.35-vgpu/2063914209959219200#1:build-log.txt%3A500)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-06-08 13:09:50 &#43;0000 UTC_: <code>15:08:10:   {&#34;[namespace kubevirt-test-default1 name testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid d112888e-d01a-47a1-af22-ff1f17598cf2]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-06-08T15:07:46.018950Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-08T15:07:46.055406Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2063971602751557632#1:build-log.txt%3A6872)
<details>
<summary>all...</summary>

* _2026-06-08 13:09:50 &#43;0000 UTC_: <code>15:08:10:   {&#34;[namespace kubevirt-test-default1 name testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid d112888e-d01a-47a1-af22-ff1f17598cf2]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-06-08T15:07:46.018950Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-08T15:07:46.055406Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2063971602751557632#1:build-log.txt%3A6872)

</details>

<hr/>
</details>

#### needs-investigation (1x / 11.11%)

<details>
<summary> no matching pattern (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-06-08 13:09:41 &#43;0000 UTC_: <code>13:29:32: failed to create new CRI image service: validate service connection: validate CRI v1 image API for endpoint &#34;unix:///var/run/crio/crio.sock&#34;: rpc error: code = DeadlineExceeded desc = stream terminated by RST_STREAM with error code: CANCEL[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2063971600276918272#1:build-log.txt%3A674)
<details>
<summary>all...</summary>

* _2026-06-08 13:09:41 &#43;0000 UTC_: <code>13:29:32: failed to create new CRI image service: validate service connection: validate CRI v1 image API for endpoint &#34;unix:///var/run/crio/crio.sock&#34;: rpc error: code = DeadlineExceeded desc = stream terminated by RST_STREAM with error code: CANCEL[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2063971600276918272#1:build-log.txt%3A674)

</details>

<hr/>
</details>

### 2026-06-05 (1x / 7.14%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-06-05 00:01:46 &#43;0000 UTC_: <code>00:13:32: I0604 20:13:31.777281    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18028/pull-kubevirt-e2e-k8s-1.36-sig-network/2062686150438424576#1:build-log.txt%3A817)
<details>
<summary>all...</summary>

* _2026-06-05 00:01:46 &#43;0000 UTC_: <code>00:13:32: I0604 20:13:31.777281    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18028/pull-kubevirt-e2e-k8s-1.36-sig-network/2062686150438424576#1:build-log.txt%3A817)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (11x / 78.57%)

<details>
<summary> download failure in context (5x / 35.71%) </summary>

<hr/>

**4x**: _2026-06-08 09:21:42 &#43;0000 UTC_: <code>09:33:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-operator/2063914210630307840#1:build-log.txt%3A4332)
<details>
<summary>all...</summary>

* _2026-06-08 09:21:43 &#43;0000 UTC_: <code>09:32:04: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.35-sig-operator/2063914211544666112#1:build-log.txt%3A4139)
<details><summary>context</summary>
<pre>09:32:04: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted:
09:32:04: INFO: Elapsed time: 0.654s
09:32:04: INFO: 0 processes.
09:32:04: ERROR: Build failed. Not running target
09:32:07: &#43; rm -f /tmp/kubevirt.deploy.FBCM
make: *** [Makefile:189: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


* _2026-06-08 09:21:42 &#43;0000 UTC_: <code>09:33:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-operator/2063914210630307840#1:build-log.txt%3A4332)
<details><summary>context</summary>
<pre>09:33:03: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
09:33:03: INFO: Elapsed time: 1.043s
09:33:03: INFO: 0 processes.
09:33:03: ERROR: Build failed. Not running target
09:33:05: &#43; rm -f /tmp/kubevirt.deploy.EWc1
make: *** [Makefile:189: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


* _2026-06-08 09:21:42 &#43;0000 UTC_: <code>09:35:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.35-sig-compute/2063914211016183808#1:build-log.txt%3A4309)
<details><summary>context</summary>
<pre>09:35:03: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
09:35:03: INFO: Elapsed time: 0.973s
09:35:03: INFO: 0 processes.
09:35:03: ERROR: Build failed. Not running target
09:35:05: &#43; rm -f /tmp/kubevirt.deploy.8nFp
make: *** [Makefile:189: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


* _2026-06-08 09:21:42 &#43;0000 UTC_: <code>09:34:43: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.35-sig-storage/2063914210865188864#1:build-log.txt%3A4144)
<details><summary>context</summary>
<pre>09:34:43: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
09:34:43: INFO: Elapsed time: 1.616s
09:34:43: INFO: 0 processes.
09:34:43: ERROR: Build failed. Not running target
09:34:45: &#43; rm -f /tmp/kubevirt.deploy.UQxq
make: *** [Makefile:189: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


</details>

<hr/>

**1x**: _2026-06-08 09:21:39 &#43;0000 UTC_: <code>09:23:47: ERROR: Analysis of target &#39;//containerimages:fedora-realtime&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-storage/2063914210416398336#1:build-log.txt%3A845)
<details>
<summary>all...</summary>

* _2026-06-08 09:21:39 &#43;0000 UTC_: <code>09:23:47: ERROR: Analysis of target &#39;//containerimages:fedora-realtime&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-storage/2063914210416398336#1:build-log.txt%3A845)
<details><summary>context</summary>
<pre>09:23:47: Repository rule oci_pull defined at:
09:23:47:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_oci/oci/private/pull.bzl:342:27: in &lt;toplevel&gt;
09:23:47: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/fedora_realtime/BUILD.bazel:7:6: @fedora_realtime//:fedora_realtime depends on @fedora_realtime_single//:fedora_realtime_single in repository @fedora_realtime_single which failed to fetch. no such package &#39;@fedora_realtime_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/auth?scope=repository:kubevirt/fedora-realtime-container-disk:pull&amp;service=quay.io] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/fedora_realtime_single/www-authenticate.json: GET returned 502 Bad Gateway
09:23:47: ERROR: Analysis of target &#39;//containerimages:fedora-realtime&#39; failed; build aborted:
09:23:47: INFO: Elapsed time: 2.194s
09:23:47: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> failed external fetch in context (2x / 14.29%) </summary>

<hr/>

**1x**: _2026-06-08 09:21:41 &#43;0000 UTC_: <code>09:31:45: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-compute/2063914210517061632#1:build-log.txt%3A3995)
<details>
<summary>all...</summary>

* _2026-06-08 09:21:41 &#43;0000 UTC_: <code>09:31:45: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-compute/2063914210517061632#1:build-log.txt%3A3995)
<details><summary>context</summary>
<pre>09:31:45: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
09:31:45: INFO: Elapsed time: 0.667s
09:31:45: INFO: 0 processes.
09:31:45: ERROR: Build failed. Not running target
09:31:46: &#43; rm -f /tmp/kubevirt.deploy.1UAC
make: *** [Makefile:189: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


</details>

<hr/>

**1x**: _2026-06-08 09:21:39 &#43;0000 UTC_: <code>09:24:33: ERROR: Analysis of target &#39;//containerimages:fedora-with-test-tooling&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-kind-1.35-vgpu/2063914209959219200#1:build-log.txt%3A500)
<details>
<summary>all...</summary>

* _2026-06-08 09:21:39 &#43;0000 UTC_: <code>09:24:33: ERROR: Analysis of target &#39;//containerimages:fedora-with-test-tooling&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-kind-1.35-vgpu/2063914209959219200#1:build-log.txt%3A500)
<details><summary>context</summary>
<pre>09:24:33: *) If there is a configured URL Rewriter, check that it does not block the request.
09:24:33:
09:24:33: See for more: https://github.com/bazel-contrib/rules_oci/blob/main/docs/pull.md#authentication-using-credential-helpers
09:24:33: ERROR: Analysis of target &#39;//containerimages:fedora-with-test-tooling&#39; failed; build aborted: Analysis failed
09:24:33: INFO: Elapsed time: 0.533s
09:24:33: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 14.29%) </summary>

<hr/>

**1x**: _2026-06-09 08:33:34 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2064264500244123648#1:build-log.txt%3A4384)
<details>
<summary>all...</summary>

* _2026-06-09 08:33:34 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2064264500244123648#1:build-log.txt%3A4384)
<details><summary>context</summary>
<pre>09:08:13: INFO: Elapsed time: 0.398s, Critical Path: 0.13s
09:08:13: INFO: 1 process: 1 internal.
09:08:13: INFO: Running command line: bazel-bin/example-guest-agent-copier /root/go/src/kubevirt.io/kubevirt/_out/cmd/example-guest-agent/example-guest-agent
make: *** [Makefile:26: bazel-build-functests] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>


</details>

<hr/>

**1x**: _2026-06-08 13:09:50 &#43;0000 UTC_: <code>15:08:10:   {&#34;[namespace kubevirt-test-default1 name testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid d112888e-d01a-47a1-af22-ff1f17598cf2]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-06-08T15:07:46.018950Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-08T15:07:46.055406Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2063971602751557632#1:build-log.txt%3A6872)
<details>
<summary>all...</summary>

* _2026-06-08 13:09:50 &#43;0000 UTC_: <code>15:08:10:   {&#34;[namespace kubevirt-test-default1 name testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid d112888e-d01a-47a1-af22-ff1f17598cf2]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-06-08T15:07:46.018950Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-08T15:07:46.055406Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2063971602751557632#1:build-log.txt%3A6872)
<details><summary>context</summary>
<pre>15:08:10:
15:08:10:   Captured StdOut/StdErr Output &gt;&gt;
15:08:10:   {&#34;[namespace kubevirt-test-default1 name testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid d112888e-d01a-47a1-af22-ff1f17598cf2]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;Event(v1.ObjectReference{Kind:\&#34;VirtualMachineInstance\&#34;, Namespace:\&#34;kubevirt-test-default1\&#34;, Name:\&#34;testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\&#34;, UID:\&#34;d112888e-d01a-47a1-af22-ff1f17598cf2\&#34;, APIVersion:\&#34;kubevirt.io/v1\&#34;, ResourceVersion:\&#34;145126\&#34;, FieldPath:\&#34;\&#34;}): type: &#39;Normal&#39; reason: &#39;SuccessfulCreate&#39; Created virtual machine pod virt-launcher-testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxfrqsg&#34;,&#34;pos&#34;:&#34;watcher.go:150&#34;,&#34;timestamp&#34;:&#34;2026-06-08T15:07:42.269099Z&#34;}
15:08:10:   {&#34;[namespace kubevirt-test-default1 name testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid d112888e-d01a-47a1-af22-ff1f17598cf2]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-06-08T15:07:46.018950Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-08T15:07:46.055406Z&#34;}
15:08:10:   &lt;&lt; Captured StdOut/StdErr Output
15:08:10: ------------------------------
15:08:48: • [40.115 seconds]</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> HTTP error in context (1x / 7.14%) </summary>

<hr/>

**1x**: _2026-06-11 08:11:20 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16106/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2064983681084166144#1:build-log.txt%3A796)
<details>
<summary>all...</summary>

* _2026-06-11 08:11:20 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16106/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2064983681084166144#1:build-log.txt%3A796)
<details><summary>context</summary>
<pre>08:19:03: main.main()
08:19:03: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
08:19:03: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:174: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 7.14%) </summary>

<hr/>

**1x**: _2026-06-05 00:01:46 &#43;0000 UTC_: <code>00:13:32: I0604 20:13:31.777281    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18028/pull-kubevirt-e2e-k8s-1.36-sig-network/2062686150438424576#1:build-log.txt%3A817)
<details>
<summary>all...</summary>

* _2026-06-05 00:01:46 &#43;0000 UTC_: <code>00:13:32: I0604 20:13:31.777281    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18028/pull-kubevirt-e2e-k8s-1.36-sig-network/2062686150438424576#1:build-log.txt%3A817)
<details><summary>context</summary>
<pre>00:13:30: [control-plane-check] kube-scheduler is healthy after 3.060930196s
00:13:31: I0604 20:13:30.779805    1615 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;
00:13:31: I0604 20:13:31.277792    1615 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;
00:13:32: I0604 20:13:31.777281    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
00:13:32: I0604 20:13:31.777434    1615 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
00:13:32: [control-plane-check] kube-apiserver is healthy after 5.003768117s
00:13:32: I0604 20:13:32.282847    1615 kubeconfig.go:657] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists</pre>
</details>


</details>

<hr/>
</details>

### internal (2x / 14.29%)

<details>
<summary> make cluster lifecycle target failure (1x / 7.14%) </summary>

<hr/>

**1x**: _2026-06-09 12:55:48 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2064329652029100032#1:build-log.txt%3A1196)
<details>
<summary>all...</summary>

* _2026-06-09 12:55:48 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2064329652029100032#1:build-log.txt%3A1196)
<details><summary>context</summary>
<pre>13:15:45: INFO: 2 processes: 1 internal, 1 processwrapper-sandbox.
13:15:45: INFO: Running command line: bazel-bin/example-guest-agent-copier /root/go/src/kubevirt.io/kubevirt/_out/cmd/example-guest-agent/example-guest-agent
13:15:54: &#43; rm -f /tmp/kubevirt.deploy.r5Pz
make: *** [Makefile:174: cluster-sync] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> make bazel-build target failure (1x / 7.14%) </summary>

<hr/>

**1x**: _2026-06-09 08:35:56 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2064264517830840320#1:build-log.txt%3A4515)
<details>
<summary>all...</summary>

* _2026-06-09 08:35:56 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2064264517830840320#1:build-log.txt%3A4515)
<details><summary>context</summary>
<pre>09:11:37: INFO: Elapsed time: 0.318s, Critical Path: 0.10s
09:11:37: INFO: 1 process: 1 internal.
09:11:37: INFO: Running command line: bazel-bin/example-guest-agent-copier /root/go/src/kubevirt.io/kubevirt/_out/cmd/example-guest-agent/example-guest-agent
make: *** [Makefile:26: bazel-build-functests] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>


</details>

<hr/>
</details>

### needs-investigation (1x / 7.14%)

<details>
<summary> no matching pattern (1x / 7.14%) </summary>

<hr/>

**1x**: _2026-06-08 13:09:41 &#43;0000 UTC_: <code>13:29:32: failed to create new CRI image service: validate service connection: validate CRI v1 image API for endpoint &#34;unix:///var/run/crio/crio.sock&#34;: rpc error: code = DeadlineExceeded desc = stream terminated by RST_STREAM with error code: CANCEL[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2063971600276918272#1:build-log.txt%3A674)
<details>
<summary>all...</summary>

* _2026-06-08 13:09:41 &#43;0000 UTC_: <code>13:29:32: failed to create new CRI image service: validate service connection: validate CRI v1 image API for endpoint &#34;unix:///var/run/crio/crio.sock&#34;: rpc error: code = DeadlineExceeded desc = stream terminated by RST_STREAM with error code: CANCEL[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2063971600276918272#1:build-log.txt%3A674)
<details><summary>context</summary>
<pre>13:29:32: I0608 09:29:28.855786    1636 checks.go:201] validating availability of port 2380
13:29:32: I0608 09:29:28.855877    1636 checks.go:241] validating the existence and emptiness of directory /var/lib/etcd
13:29:32: [preflight] Some fatal errors occurred:
13:29:32: failed to create new CRI image service: validate service connection: validate CRI v1 image API for endpoint &#34;unix:///var/run/crio/crio.sock&#34;: rpc error: code = DeadlineExceeded desc = stream terminated by RST_STREAM with error code: CANCEL[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`
13:29:32: error execution phase preflight
13:29:32: k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow.(*Runner).Run.func1
13:29:32: 	k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow/runner.go:262</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (9x / 64.29%)


#### external (9x / 100.00%)

<details>
<summary> download failure in context (5x / 55.56%) </summary>

<hr/>

**4x**: _2026-06-08 09:21:42 &#43;0000 UTC_: <code>09:33:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-operator/2063914210630307840#1:build-log.txt%3A4332)
<details>
<summary>all...</summary>

* _2026-06-08 09:21:43 &#43;0000 UTC_: <code>09:32:04: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.35-sig-operator/2063914211544666112#1:build-log.txt%3A4139)

* _2026-06-08 09:21:42 &#43;0000 UTC_: <code>09:33:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-operator/2063914210630307840#1:build-log.txt%3A4332)

* _2026-06-08 09:21:42 &#43;0000 UTC_: <code>09:35:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.35-sig-compute/2063914211016183808#1:build-log.txt%3A4309)

* _2026-06-08 09:21:42 &#43;0000 UTC_: <code>09:34:43: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.35-sig-storage/2063914210865188864#1:build-log.txt%3A4144)

</details>

<hr/>

**1x**: _2026-06-08 09:21:39 &#43;0000 UTC_: <code>09:23:47: ERROR: Analysis of target &#39;//containerimages:fedora-realtime&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-storage/2063914210416398336#1:build-log.txt%3A845)
<details>
<summary>all...</summary>

* _2026-06-08 09:21:39 &#43;0000 UTC_: <code>09:23:47: ERROR: Analysis of target &#39;//containerimages:fedora-realtime&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-storage/2063914210416398336#1:build-log.txt%3A845)

</details>

<hr/>
</details>
<details>
<summary> failed external fetch in context (2x / 22.22%) </summary>

<hr/>

**1x**: _2026-06-08 09:21:41 &#43;0000 UTC_: <code>09:31:45: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-compute/2063914210517061632#1:build-log.txt%3A3995)
<details>
<summary>all...</summary>

* _2026-06-08 09:21:41 &#43;0000 UTC_: <code>09:31:45: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-compute/2063914210517061632#1:build-log.txt%3A3995)

</details>

<hr/>

**1x**: _2026-06-08 09:21:39 &#43;0000 UTC_: <code>09:24:33: ERROR: Analysis of target &#39;//containerimages:fedora-with-test-tooling&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-kind-1.35-vgpu/2063914209959219200#1:build-log.txt%3A500)
<details>
<summary>all...</summary>

* _2026-06-08 09:21:39 &#43;0000 UTC_: <code>09:24:33: ERROR: Analysis of target &#39;//containerimages:fedora-with-test-tooling&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-kind-1.35-vgpu/2063914209959219200#1:build-log.txt%3A500)

</details>

<hr/>
</details>
<details>
<summary> HTTP error in context (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-06-11 08:11:20 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16106/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2064983681084166144#1:build-log.txt%3A796)
<details>
<summary>all...</summary>

* _2026-06-11 08:11:20 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16106/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2064983681084166144#1:build-log.txt%3A796)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-06-05 00:01:46 &#43;0000 UTC_: <code>00:13:32: I0604 20:13:31.777281    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18028/pull-kubevirt-e2e-k8s-1.36-sig-network/2062686150438424576#1:build-log.txt%3A817)
<details>
<summary>all...</summary>

* _2026-06-05 00:01:46 &#43;0000 UTC_: <code>00:13:32: I0604 20:13:31.777281    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18028/pull-kubevirt-e2e-k8s-1.36-sig-network/2062686150438424576#1:build-log.txt%3A817)

</details>

<hr/>
</details>

### release-1.7 (3x / 21.43%)


#### internal (2x / 66.67%)

<details>
<summary> make cluster lifecycle target failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-06-09 12:55:48 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2064329652029100032#1:build-log.txt%3A1196)
<details>
<summary>all...</summary>

* _2026-06-09 12:55:48 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2064329652029100032#1:build-log.txt%3A1196)

</details>

<hr/>
</details>
<details>
<summary> make bazel-build target failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-06-09 08:35:56 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2064264517830840320#1:build-log.txt%3A4515)
<details>
<summary>all...</summary>

* _2026-06-09 08:35:56 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2064264517830840320#1:build-log.txt%3A4515)

</details>

<hr/>
</details>

#### external (1x / 33.33%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-06-09 08:33:34 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2064264500244123648#1:build-log.txt%3A4384)
<details>
<summary>all...</summary>

* _2026-06-09 08:33:34 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2064264500244123648#1:build-log.txt%3A4384)

</details>

<hr/>
</details>

### release-1.8 (2x / 14.29%)


#### needs-investigation (1x / 50.00%)

<details>
<summary> no matching pattern (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-06-08 13:09:41 &#43;0000 UTC_: <code>13:29:32: failed to create new CRI image service: validate service connection: validate CRI v1 image API for endpoint &#34;unix:///var/run/crio/crio.sock&#34;: rpc error: code = DeadlineExceeded desc = stream terminated by RST_STREAM with error code: CANCEL[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2063971600276918272#1:build-log.txt%3A674)
<details>
<summary>all...</summary>

* _2026-06-08 13:09:41 &#43;0000 UTC_: <code>13:29:32: failed to create new CRI image service: validate service connection: validate CRI v1 image API for endpoint &#34;unix:///var/run/crio/crio.sock&#34;: rpc error: code = DeadlineExceeded desc = stream terminated by RST_STREAM with error code: CANCEL[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2063971600276918272#1:build-log.txt%3A674)

</details>

<hr/>
</details>

#### external (1x / 50.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-06-08 13:09:50 &#43;0000 UTC_: <code>15:08:10:   {&#34;[namespace kubevirt-test-default1 name testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid d112888e-d01a-47a1-af22-ff1f17598cf2]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-06-08T15:07:46.018950Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-08T15:07:46.055406Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2063971602751557632#1:build-log.txt%3A6872)
<details>
<summary>all...</summary>

* _2026-06-08 13:09:50 &#43;0000 UTC_: <code>15:08:10:   {&#34;[namespace kubevirt-test-default1 name testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid d112888e-d01a-47a1-af22-ff1f17598cf2]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-06-08T15:07:46.018950Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-08T15:07:46.055406Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2063971602751557632#1:build-log.txt%3A6872)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (8x / 57.14%)


#### external (7x / 87.50%)

<details>
<summary> download failure in context (3x / 37.50%) </summary>

<hr/>

**3x**: _2026-06-08 09:21:42 &#43;0000 UTC_: <code>09:33:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-operator/2063914210630307840#1:build-log.txt%3A4332)
<details>
<summary>all...</summary>

* _2026-06-08 09:21:43 &#43;0000 UTC_: <code>09:32:04: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.35-sig-operator/2063914211544666112#1:build-log.txt%3A4139)

* _2026-06-08 09:21:42 &#43;0000 UTC_: <code>09:33:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-operator/2063914210630307840#1:build-log.txt%3A4332)

* _2026-06-08 09:21:42 &#43;0000 UTC_: <code>09:35:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.35-sig-compute/2063914211016183808#1:build-log.txt%3A4309)

</details>

<hr/>
</details>
<details>
<summary> failed external fetch in context (2x / 25.00%) </summary>

<hr/>

**1x**: _2026-06-08 09:21:41 &#43;0000 UTC_: <code>09:31:45: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-compute/2063914210517061632#1:build-log.txt%3A3995)
<details>
<summary>all...</summary>

* _2026-06-08 09:21:41 &#43;0000 UTC_: <code>09:31:45: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-compute/2063914210517061632#1:build-log.txt%3A3995)

</details>

<hr/>

**1x**: _2026-06-08 09:21:39 &#43;0000 UTC_: <code>09:24:33: ERROR: Analysis of target &#39;//containerimages:fedora-with-test-tooling&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-kind-1.35-vgpu/2063914209959219200#1:build-log.txt%3A500)
<details>
<summary>all...</summary>

* _2026-06-08 09:21:39 &#43;0000 UTC_: <code>09:24:33: ERROR: Analysis of target &#39;//containerimages:fedora-with-test-tooling&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-kind-1.35-vgpu/2063914209959219200#1:build-log.txt%3A500)

</details>

<hr/>
</details>
<details>
<summary> HTTP error in context (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-06-11 08:11:20 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16106/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2064983681084166144#1:build-log.txt%3A796)
<details>
<summary>all...</summary>

* _2026-06-11 08:11:20 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16106/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2064983681084166144#1:build-log.txt%3A796)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-06-09 08:33:34 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2064264500244123648#1:build-log.txt%3A4384)
<details>
<summary>all...</summary>

* _2026-06-09 08:33:34 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2064264500244123648#1:build-log.txt%3A4384)

</details>

<hr/>
</details>

#### needs-investigation (1x / 12.50%)

<details>
<summary> no matching pattern (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-06-08 13:09:41 &#43;0000 UTC_: <code>13:29:32: failed to create new CRI image service: validate service connection: validate CRI v1 image API for endpoint &#34;unix:///var/run/crio/crio.sock&#34;: rpc error: code = DeadlineExceeded desc = stream terminated by RST_STREAM with error code: CANCEL[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2063971600276918272#1:build-log.txt%3A674)
<details>
<summary>all...</summary>

* _2026-06-08 13:09:41 &#43;0000 UTC_: <code>13:29:32: failed to create new CRI image service: validate service connection: validate CRI v1 image API for endpoint &#34;unix:///var/run/crio/crio.sock&#34;: rpc error: code = DeadlineExceeded desc = stream terminated by RST_STREAM with error code: CANCEL[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2063971600276918272#1:build-log.txt%3A674)

</details>

<hr/>
</details>

### sig-network (3x / 21.43%)


#### internal (2x / 66.67%)

<details>
<summary> make cluster lifecycle target failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-06-09 12:55:48 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2064329652029100032#1:build-log.txt%3A1196)
<details>
<summary>all...</summary>

* _2026-06-09 12:55:48 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2064329652029100032#1:build-log.txt%3A1196)

</details>

<hr/>
</details>
<details>
<summary> make bazel-build target failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-06-09 08:35:56 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2064264517830840320#1:build-log.txt%3A4515)
<details>
<summary>all...</summary>

* _2026-06-09 08:35:56 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2064264517830840320#1:build-log.txt%3A4515)

</details>

<hr/>
</details>

#### external (1x / 33.33%)

<details>
<summary> transient kube-apiserver body decode noise (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-06-05 00:01:46 &#43;0000 UTC_: <code>00:13:32: I0604 20:13:31.777281    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18028/pull-kubevirt-e2e-k8s-1.36-sig-network/2062686150438424576#1:build-log.txt%3A817)
<details>
<summary>all...</summary>

* _2026-06-05 00:01:46 &#43;0000 UTC_: <code>00:13:32: I0604 20:13:31.777281    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18028/pull-kubevirt-e2e-k8s-1.36-sig-network/2062686150438424576#1:build-log.txt%3A817)

</details>

<hr/>
</details>

### sig-storage (3x / 21.43%)


#### external (3x / 100.00%)

<details>
<summary> download failure in context (2x / 66.67%) </summary>

<hr/>

**1x**: _2026-06-08 09:21:42 &#43;0000 UTC_: <code>09:34:43: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.35-sig-storage/2063914210865188864#1:build-log.txt%3A4144)
<details>
<summary>all...</summary>

* _2026-06-08 09:21:42 &#43;0000 UTC_: <code>09:34:43: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.35-sig-storage/2063914210865188864#1:build-log.txt%3A4144)

</details>

<hr/>

**1x**: _2026-06-08 09:21:39 &#43;0000 UTC_: <code>09:23:47: ERROR: Analysis of target &#39;//containerimages:fedora-realtime&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-storage/2063914210416398336#1:build-log.txt%3A845)
<details>
<summary>all...</summary>

* _2026-06-08 09:21:39 &#43;0000 UTC_: <code>09:23:47: ERROR: Analysis of target &#39;//containerimages:fedora-realtime&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17793/pull-kubevirt-e2e-k8s-1.34-sig-storage/2063914210416398336#1:build-log.txt%3A845)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-06-08 13:09:50 &#43;0000 UTC_: <code>15:08:10:   {&#34;[namespace kubevirt-test-default1 name testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid d112888e-d01a-47a1-af22-ff1f17598cf2]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-06-08T15:07:46.018950Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-08T15:07:46.055406Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2063971602751557632#1:build-log.txt%3A6872)
<details>
<summary>all...</summary>

* _2026-06-08 13:09:50 &#43;0000 UTC_: <code>15:08:10:   {&#34;[namespace kubevirt-test-default1 name testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid d112888e-d01a-47a1-af22-ff1f17598cf2]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-06-08T15:07:46.018950Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-08T15:07:46.055406Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2063971602751557632#1:build-log.txt%3A6872)

</details>

<hr/>
</details>

Last updated: 2026-06-11 19:04:29
