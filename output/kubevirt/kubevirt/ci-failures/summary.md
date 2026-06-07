



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-06-05 (1x / 2.86%)


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

### 2026-06-01 (34x / 97.14%)


#### external (28x / 82.35%)

<details>
<summary> download failure in context (13x / 38.24%) </summary>

<hr/>

**1x**: _2026-06-01 14:03:07 &#43;0000 UTC_: <code>14:05:31: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-storage/2061439596666818560#1:build-log.txt%3A590)
<details>
<summary>all...</summary>

* _2026-06-01 14:03:07 &#43;0000 UTC_: <code>14:05:31: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-storage/2061439596666818560#1:build-log.txt%3A590)

</details>

<hr/>

**12x**: _2026-06-01 07:46:15 &#43;0000 UTC_: <code>08:18:52: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2061353428428263424#1:build-log.txt%3A4109)
<details>
<summary>all...</summary>

* _2026-06-01 13:55:33 &#43;0000 UTC_: <code>14:07:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2061439596553572352#1:build-log.txt%3A4208)

* _2026-06-01 10:14:55 &#43;0000 UTC_: <code>10:35:12: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2061390371253915648#1:build-log.txt%3A3979)

* _2026-06-01 10:08:23 &#43;0000 UTC_: <code>10:29:53: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2061384109464227840#1:build-log.txt%3A3943)

* _2026-06-01 10:06:05 &#43;0000 UTC_: <code>10:32:21: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2061384108692475904#1:build-log.txt%3A4084)

* _2026-06-01 07:46:25 &#43;0000 UTC_: <code>08:13:45: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2061353425072820224#1:build-log.txt%3A3955)

* _2026-06-01 07:46:15 &#43;0000 UTC_: <code>08:18:52: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2061353428428263424#1:build-log.txt%3A4109)

* _2026-06-01 07:46:12 &#43;0000 UTC_: <code>08:18:27: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2061353420920459264#1:build-log.txt%3A4085)

* _2026-06-01 07:46:10 &#43;0000 UTC_: <code>08:14:59: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial-1.8/2061353426742153216#1:build-log.txt%3A3975)

* _2026-06-01 07:46:07 &#43;0000 UTC_: <code>08:14:29: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2061353421725765632#1:build-log.txt%3A3951)

* _2026-06-01 07:46:02 &#43;0000 UTC_: <code>08:14:34: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2061353429304872960#1:build-log.txt%3A3964)

* _2026-06-01 07:45:59 &#43;0000 UTC_: <code>08:13:28: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2061353408442404864#1:build-log.txt%3A3983)

* _2026-06-01 07:45:58 &#43;0000 UTC_: <code>08:03:49: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-kind-1.35-vgpu-1.8/2061353408597594112#1:build-log.txt%3A4388)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (10x / 29.41%) </summary>

<hr/>

**3x**: _2026-06-01 14:05:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-compute/2061439596796841984#1:build-log.txt%3A239)
<details>
<summary>all...</summary>

* _2026-06-01 14:05:38 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-operator/2061439596926865408#1:build-log.txt%3A301)

* _2026-06-01 14:05:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-compute/2061439596796841984#1:build-log.txt%3A239)

* _2026-06-01 10:34:38 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-kind-1.35-vgpu-1.8/2061391210236350464#1:build-log.txt%3A184)

</details>

<hr/>

**7x**: _2026-06-01 10:27:12 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2061384111733346304#1:build-log.txt%3A599)
<details>
<summary>all...</summary>

* _2026-06-01 10:33:51 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-operator-1.8/2061384114241540096#1:build-log.txt%3A205)

* _2026-06-01 10:32:14 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2061384113398484992#1:build-log.txt%3A202)

* _2026-06-01 10:31:03 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2061384112605761536#1:build-log.txt%3A207)

* _2026-06-01 10:27:12 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2061384111733346304#1:build-log.txt%3A599)

* _2026-06-01 10:25:52 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial-1.8/2061384111007731712#1:build-log.txt%3A581)

* _2026-06-01 07:51:22 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-kind-sriov-1.8/2061353408887001088#1:build-log.txt%3A1671)

* _2026-06-01 07:46:34 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2061353424254930944#1:build-log.txt%3A606)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (5x / 14.71%) </summary>

<hr/>

**5x**: _2026-06-01 07:46:02 &#43;0000 UTC_: <code>08:12:03: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-network-1.8/2061353420056432640#1:build-log.txt%3A977)
<details>
<summary>all...</summary>

* _2026-06-01 10:18:48 &#43;0000 UTC_: <code>10:35:03: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.8/2061384110051430400#1:build-log.txt%3A909)

* _2026-06-01 07:46:18 &#43;0000 UTC_: <code>08:12:17: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2061353427622957056#1:build-log.txt%3A998)

* _2026-06-01 07:46:05 &#43;0000 UTC_: <code>08:09:47: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-network-1.8/2061353423407681536#1:build-log.txt%3A984)

* _2026-06-01 07:46:02 &#43;0000 UTC_: <code>08:12:03: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-network-1.8/2061353420056432640#1:build-log.txt%3A977)

* _2026-06-01 07:45:59 &#43;0000 UTC_: <code>08:15:29: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations-1.8/2061353417565016064#1:build-log.txt%3A1188)

</details>

<hr/>
</details>

#### needs-investigation (1x / 2.94%)

<details>
<summary> no matching pattern (1x / 2.94%) </summary>

<hr/>

**1x**: _2026-06-01 20:05:06 &#43;0000 UTC_: <code>22:09:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-2qbf8-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 90e398de-9d0a-43a6-831b-6bfceadaf099]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-01T22:08:53.733535Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18006/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2061539398683463680#1:build-log.txt%3A12941)
<details>
<summary>all...</summary>

* _2026-06-01 20:05:06 &#43;0000 UTC_: <code>22:09:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-2qbf8-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 90e398de-9d0a-43a6-831b-6bfceadaf099]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-01T22:08:53.733535Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18006/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2061539398683463680#1:build-log.txt%3A12941)

</details>

<hr/>
</details>

#### internal (5x / 14.71%)

<details>
<summary> deployment timeout (5x / 14.71%) </summary>

<hr/>

**5x**: _2026-06-01 07:46:02 &#43;0000 UTC_: <code>08:10:04: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2061353422568820736#1:build-log.txt%3A903)
<details>
<summary>all...</summary>

* _2026-06-01 14:01:39 &#43;0000 UTC_: <code>14:09:35: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-network/2061439596561960960#1:build-log.txt%3A1002)

* _2026-06-01 07:46:18 &#43;0000 UTC_: <code>08:11:57: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network-1.8/2061353416696795136#1:build-log.txt%3A998)

* _2026-06-01 07:46:07 &#43;0000 UTC_: <code>08:11:22: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-operator-1.8/2061353430110179328#1:build-log.txt%3A927)

* _2026-06-01 07:46:06 &#43;0000 UTC_: <code>08:05:36: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.8/2061353425936846848#1:build-log.txt%3A911)

* _2026-06-01 07:46:02 &#43;0000 UTC_: <code>08:10:04: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2061353422568820736#1:build-log.txt%3A903)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (29x / 82.86%)

<details>
<summary> download failure in context (13x / 37.14%) </summary>

<hr/>

**1x**: _2026-06-01 14:03:07 &#43;0000 UTC_: <code>14:05:31: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-storage/2061439596666818560#1:build-log.txt%3A590)
<details>
<summary>all...</summary>

* _2026-06-01 14:03:07 &#43;0000 UTC_: <code>14:05:31: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-storage/2061439596666818560#1:build-log.txt%3A590)
<details><summary>context</summary>
<pre>14:05:31: Repository rule oci_pull defined at:
14:05:31:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_oci/oci/private/pull.bzl:342:27: in &lt;toplevel&gt;
14:05:31: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/auth?scope=repository:kubevirt/alpine-ext-kernel-boot-demo:pull&amp;service=quay.io] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/www-authenticate.json: GET returned 502 Bad Gateway
14:05:31: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
14:05:31: INFO: Elapsed time: 14.904s
14:05:31: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>

**12x**: _2026-06-01 07:46:15 &#43;0000 UTC_: <code>08:18:52: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2061353428428263424#1:build-log.txt%3A4109)
<details>
<summary>all...</summary>

* _2026-06-01 13:55:33 &#43;0000 UTC_: <code>14:07:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2061439596553572352#1:build-log.txt%3A4208)
<details><summary>context</summary>
<pre>14:07:03: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
14:07:03: INFO: Elapsed time: 0.412s
14:07:03: INFO: 0 processes.
14:07:03: ERROR: Build failed. Not running target
14:07:04: &#43; rm -f /tmp/kubevirt.deploy.ywrU
make: *** [Makefile:189: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


* _2026-06-01 10:14:55 &#43;0000 UTC_: <code>10:35:12: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2061390371253915648#1:build-log.txt%3A3979)
<details><summary>context</summary>
<pre>10:35:12: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
10:35:12: INFO: Elapsed time: 0.505s
10:35:12: INFO: 0 processes.
10:35:12: ERROR: Build failed. Not running target
10:35:14: &#43; rm -f /tmp/kubevirt.deploy.BfJP
make: *** [Makefile:188: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


* _2026-06-01 10:08:23 &#43;0000 UTC_: <code>10:29:53: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2061384109464227840#1:build-log.txt%3A3943)
<details><summary>context</summary>
<pre>10:29:53: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
10:29:53: INFO: Elapsed time: 1.930s
10:29:53: INFO: 0 processes.
10:29:53: ERROR: Build failed. Not running target
10:29:54: &#43; rm -f /tmp/kubevirt.deploy.349K
make: *** [Makefile:188: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


* _2026-06-01 10:06:05 &#43;0000 UTC_: <code>10:32:21: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2061384108692475904#1:build-log.txt%3A4084)
<details><summary>context</summary>
<pre>10:32:21: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
10:32:21: INFO: Elapsed time: 20.573s
10:32:21: INFO: 0 processes.
10:32:21: ERROR: Build failed. Not running target
10:32:28: &#43; rm -f /tmp/kubevirt.deploy.2mhK
make: *** [Makefile:188: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


* _2026-06-01 07:46:25 &#43;0000 UTC_: <code>08:13:45: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2061353425072820224#1:build-log.txt%3A3955)
<details><summary>context</summary>
<pre>08:13:45: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
08:13:45: INFO: Elapsed time: 85.100s
08:13:45: INFO: 0 processes.
08:13:45: ERROR: Build failed. Not running target
08:13:47: &#43; rm -f /tmp/kubevirt.deploy.EH05
make: *** [Makefile:188: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


* _2026-06-01 07:46:15 &#43;0000 UTC_: <code>08:18:52: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2061353428428263424#1:build-log.txt%3A4109)
<details><summary>context</summary>
<pre>08:18:52: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
08:18:52: INFO: Elapsed time: 45.700s
08:18:52: INFO: 0 processes.
08:18:52: ERROR: Build failed. Not running target
08:18:54: &#43; rm -f /tmp/kubevirt.deploy.bZdk
make: *** [Makefile:188: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


* _2026-06-01 07:46:12 &#43;0000 UTC_: <code>08:18:27: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2061353420920459264#1:build-log.txt%3A4085)
<details><summary>context</summary>
<pre>08:18:27: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
08:18:27: INFO: Elapsed time: 42.487s
08:18:27: INFO: 0 processes.
08:18:27: ERROR: Build failed. Not running target
08:18:29: &#43; rm -f /tmp/kubevirt.deploy.U99t
make: *** [Makefile:188: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


* _2026-06-01 07:46:10 &#43;0000 UTC_: <code>08:14:59: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial-1.8/2061353426742153216#1:build-log.txt%3A3975)
<details><summary>context</summary>
<pre>08:14:59: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
08:14:59: INFO: Elapsed time: 22.056s
08:14:59: INFO: 0 processes.
08:14:59: ERROR: Build failed. Not running target
08:15:01: &#43; rm -f /tmp/kubevirt.deploy.9CwZ
make: *** [Makefile:188: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


* _2026-06-01 07:46:07 &#43;0000 UTC_: <code>08:14:29: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2061353421725765632#1:build-log.txt%3A3951)
<details><summary>context</summary>
<pre>08:14:29: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
08:14:29: INFO: Elapsed time: 21.088s
08:14:29: INFO: 0 processes.
08:14:29: ERROR: Build failed. Not running target
08:14:31: &#43; rm -f /tmp/kubevirt.deploy.EepO
make: *** [Makefile:188: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


* _2026-06-01 07:46:02 &#43;0000 UTC_: <code>08:14:34: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2061353429304872960#1:build-log.txt%3A3964)
<details><summary>context</summary>
<pre>08:14:34: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
08:14:34: INFO: Elapsed time: 11.840s
08:14:34: INFO: 0 processes.
08:14:34: ERROR: Build failed. Not running target
08:14:36: &#43; rm -f /tmp/kubevirt.deploy.Qv1t
make: *** [Makefile:188: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


* _2026-06-01 07:45:59 &#43;0000 UTC_: <code>08:13:28: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2061353408442404864#1:build-log.txt%3A3983)
<details><summary>context</summary>
<pre>08:13:28: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
08:13:28: INFO: Elapsed time: 13.765s
08:13:28: INFO: 0 processes.
08:13:28: ERROR: Build failed. Not running target
08:13:29: &#43; rm -f /tmp/kubevirt.deploy.ixqz
make: *** [Makefile:188: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


* _2026-06-01 07:45:58 &#43;0000 UTC_: <code>08:03:49: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-kind-1.35-vgpu-1.8/2061353408597594112#1:build-log.txt%3A4388)
<details><summary>context</summary>
<pre>08:03:49: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
08:03:49: INFO: Elapsed time: 0.505s
08:03:49: INFO: 0 processes.
08:03:49: ERROR: Build failed. Not running target
08:03:53: &#43; rm -f /tmp/kubevirt.deploy.wtei
make: *** [Makefile:188: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (10x / 28.57%) </summary>

<hr/>

**3x**: _2026-06-01 14:05:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-compute/2061439596796841984#1:build-log.txt%3A239)
<details>
<summary>all...</summary>

* _2026-06-01 14:05:38 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-operator/2061439596926865408#1:build-log.txt%3A301)
<details><summary>context</summary>
<pre>14:07:08: Copying blob sha256:ea389a75c17fde43e7cd5c66516847c2c408b5b3b516bafaeb3aedd7fe03ccf7
14:07:08: Copying blob sha256:41d62e0ccf8c09688ae70cac2e6a437e05e4fb444268ccdf95995c226b7757e5
14:07:14: Error: unable to copy from source docker://quay.io/kubevirt/builder:2602201229-4494d1587: copying system image from manifest list: reading blob sha256:957776ed783da3cb80920be93eaa5961b047f73eaccbd6a6ac22c5df186de45a: fetching blob: received unexpected HTTP status: 502 Bad Gateway
make: *** [Makefile:39: bazel-build-images] Error 125
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-06-01 14:05:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-compute/2061439596796841984#1:build-log.txt%3A239)
<details><summary>context</summary>
<pre>14:23:41: time=&#34;2026-06-01T14:23:41Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (2/3). Error: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 504 Gateway Time-out&#34;
14:23:52: time=&#34;2026-06-01T14:23:52Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (3/3). Error: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 504 Gateway Time-out&#34;
14:23:53: Error: unable to copy from source docker://quay.io/kubevirt/builder:2602201229-4494d1587: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway
make: *** [Makefile:39: bazel-build-images] Error 125
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-06-01 10:34:38 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-kind-1.35-vgpu-1.8/2061391210236350464#1:build-log.txt%3A184)
<details><summary>context</summary>
<pre>10:34:53: time=&#34;2026-06-01T10:34:53Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (2/3). Error: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway&#34;
10:34:54: time=&#34;2026-06-01T10:34:54Z&#34; level=warning msg=&#34;Failed, retrying in 1s ... (3/3). Error: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway&#34;
10:34:55: Error: unable to copy from source docker://quay.io/kubevirt/builder:2602201229-4494d1587: initializing source docker://quay.io/kubevirt/builder:2602201229-4494d1587: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway
make: *** [Makefile:39: bazel-build-images] Error 125
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


</details>

<hr/>

**7x**: _2026-06-01 10:27:12 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2061384111733346304#1:build-log.txt%3A599)
<details>
<summary>all...</summary>

* _2026-06-01 10:33:51 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-operator-1.8/2061384114241540096#1:build-log.txt%3A205)
<details><summary>context</summary>
<pre>10:34:11: selecting podman as container runtime
10:34:11: Trying to pull quay.io/kubevirtci/gocli:2602190951-056f8f79...
10:34:15: Error: unable to copy from source docker://quay.io/kubevirtci/gocli:2602190951-056f8f79: initializing source docker://quay.io/kubevirtci/gocli:2602190951-056f8f79: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway
make: *** [Makefile:176: cluster-down] Error 125
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-06-01 10:32:14 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2061384113398484992#1:build-log.txt%3A202)
<details><summary>context</summary>
<pre>10:33:38: selecting podman as container runtime
10:33:38: Trying to pull quay.io/kubevirtci/gocli:2602190951-056f8f79...
10:33:42: Error: unable to copy from source docker://quay.io/kubevirtci/gocli:2602190951-056f8f79: initializing source docker://quay.io/kubevirtci/gocli:2602190951-056f8f79: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway
make: *** [Makefile:176: cluster-down] Error 125
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-06-01 10:31:03 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2061384112605761536#1:build-log.txt%3A207)
<details><summary>context</summary>
<pre>10:33:23: selecting podman as container runtime
10:33:23: Trying to pull quay.io/kubevirtci/gocli:2602190951-056f8f79...
10:33:26: Error: unable to copy from source docker://quay.io/kubevirtci/gocli:2602190951-056f8f79: initializing source docker://quay.io/kubevirtci/gocli:2602190951-056f8f79: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway
make: *** [Makefile:176: cluster-down] Error 125
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-06-01 10:27:12 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2061384111733346304#1:build-log.txt%3A599)
<details><summary>context</summary>
<pre>10:35:06: selecting podman as container runtime
10:35:06: Trying to pull quay.io/kubevirtci/gocli:2602190951-056f8f79...
10:35:09: Error: unable to copy from source docker://quay.io/kubevirtci/gocli:2602190951-056f8f79: initializing source docker://quay.io/kubevirtci/gocli:2602190951-056f8f79: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway
make: *** [Makefile:176: cluster-down] Error 125
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-06-01 10:25:52 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial-1.8/2061384111007731712#1:build-log.txt%3A581)
<details><summary>context</summary>
<pre>10:33:06: selecting podman as container runtime
10:33:06: Trying to pull quay.io/kubevirtci/gocli:2602190951-056f8f79...
10:33:09: Error: unable to copy from source docker://quay.io/kubevirtci/gocli:2602190951-056f8f79: initializing source docker://quay.io/kubevirtci/gocli:2602190951-056f8f79: pinging container registry quay.io: received unexpected HTTP status: 502 Bad Gateway
make: *** [Makefile:176: cluster-down] Error 125
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-06-01 07:51:22 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-kind-sriov-1.8/2061353408887001088#1:build-log.txt%3A1671)
<details><summary>context</summary>
<pre>08:00:57: Copying blob sha256:e2ead8259a04d39492c25c9548078200c5ec429f628dcf7b7535137954cc2df0
08:00:57: Copying blob sha256:79e9f2f55bf5465a02ee6a6170e66005b20c7aa6b115af6fcd04fad706ea651a
08:00:57: Error: unable to copy from source docker://quay.io/kubevirtci/library-registry:2.7.1: copying system image from manifest list: reading blob sha256:0d96da54f60b86a4d869d44b44cfca69d71c4776b81d361bc057d6666ec0d878: fetching blob: received unexpected HTTP status: 502 Bad Gateway
make: *** [Makefile:173: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-01 07:46:34 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2061353424254930944#1:build-log.txt%3A606)
<details><summary>context</summary>
<pre>08:03:00: Copying blob sha256:fead7c895b0c4ee4f33b3a53f800ad5f2e8868176dfb1841f3b68dfc2430d880
08:03:00: Copying blob sha256:27f42aae469e8348e36956a88db953c0ead214e04bbe76daee245dd3047ce1b8
08:09:35: Error: unable to copy from source docker://quay.io/kubevirtci/gocli:2602190951-056f8f79: parsing image configuration: fetching blob: received unexpected HTTP status: 502 Bad Gateway
make: *** [Makefile:176: cluster-down] Error 125
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (5x / 14.29%) </summary>

<hr/>

**5x**: _2026-06-01 07:46:02 &#43;0000 UTC_: <code>08:12:03: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-network-1.8/2061353420056432640#1:build-log.txt%3A977)
<details>
<summary>all...</summary>

* _2026-06-01 10:18:48 &#43;0000 UTC_: <code>10:35:03: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.8/2061384110051430400#1:build-log.txt%3A909)
<details><summary>context</summary>
<pre>10:31:42: time=&#34;2026-06-01T10:31:42Z&#34; level=info msg=&#34;Object whereabouts-config applied successfully&#34;
10:31:42: time=&#34;2026-06-01T10:31:42Z&#34; level=info msg=&#34;Object whereabouts applied successfully&#34;
10:31:42: time=&#34;2026-06-01T10:31:42Z&#34; level=info msg=&#34;[node 1]: kubectl --kubeconfig=/etc/kubernetes/admin.conf wait deployment -n cluster-network-addons cluster-network-addons-operator --for condition=Available --timeout=200s&#34;
10:35:03: error: timed out waiting for the condition on deployments/cluster-network-addons-operator
10:35:03:
10:35:03: ===== 1bcb53ebf80e72c1b14dc520066edd06e3ea25b07923c5804f4a79de9958c697 ====
10:35:03: &#43; NUM_NODES=2</pre>
</details>


* _2026-06-01 07:46:18 &#43;0000 UTC_: <code>08:12:17: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2061353427622957056#1:build-log.txt%3A998)
<details><summary>context</summary>
<pre>08:08:56: time=&#34;2026-06-01T08:08:56Z&#34; level=info msg=&#34;Object whereabouts-config applied successfully&#34;
08:08:57: time=&#34;2026-06-01T08:08:57Z&#34; level=info msg=&#34;Object whereabouts applied successfully&#34;
08:08:57: time=&#34;2026-06-01T08:08:57Z&#34; level=info msg=&#34;[node 1]: kubectl --kubeconfig=/etc/kubernetes/admin.conf wait deployment -n cluster-network-addons cluster-network-addons-operator --for condition=Available --timeout=200s&#34;
08:12:17: error: timed out waiting for the condition on deployments/cluster-network-addons-operator
08:12:17:
08:12:17: ===== 6ca61095133e120cae40937924519c8d86c7ea776c1c21beb1db6d4ce49b6f3f ====
08:12:17: &#43; NUM_NODES=3</pre>
</details>


* _2026-06-01 07:46:05 &#43;0000 UTC_: <code>08:09:47: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-network-1.8/2061353423407681536#1:build-log.txt%3A984)
<details><summary>context</summary>
<pre>08:06:27: time=&#34;2026-06-01T08:06:27Z&#34; level=info msg=&#34;Object whereabouts-config applied successfully&#34;
08:06:27: time=&#34;2026-06-01T08:06:27Z&#34; level=info msg=&#34;Object whereabouts applied successfully&#34;
08:06:27: time=&#34;2026-06-01T08:06:27Z&#34; level=info msg=&#34;[node 1]: kubectl --kubeconfig=/etc/kubernetes/admin.conf wait deployment -n cluster-network-addons cluster-network-addons-operator --for condition=Available --timeout=200s&#34;
08:09:47: error: timed out waiting for the condition on deployments/cluster-network-addons-operator
08:09:47:
08:09:47: ===== bea9c216e5f4a17351eca35447d2d8e1a8bfa49e34a928f52ec94951ec5759d0 ====
08:09:47: &#43; NUM_NODES=3</pre>
</details>


* _2026-06-01 07:46:02 &#43;0000 UTC_: <code>08:12:03: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-network-1.8/2061353420056432640#1:build-log.txt%3A977)
<details><summary>context</summary>
<pre>08:08:43: time=&#34;2026-06-01T08:08:43Z&#34; level=info msg=&#34;Object whereabouts-config applied successfully&#34;
08:08:43: time=&#34;2026-06-01T08:08:43Z&#34; level=info msg=&#34;Object whereabouts applied successfully&#34;
08:08:43: time=&#34;2026-06-01T08:08:43Z&#34; level=info msg=&#34;[node 1]: kubectl --kubeconfig=/etc/kubernetes/admin.conf wait deployment -n cluster-network-addons cluster-network-addons-operator --for condition=Available --timeout=200s&#34;
08:12:03: error: timed out waiting for the condition on deployments/cluster-network-addons-operator
08:12:03:
08:12:03: ===== 691a01b11757f849be066630d14065074d3cc6ab8a5dbe6bd66dea34b9274b81 ====
08:12:03: &#43; NUM_NODES=3</pre>
</details>


* _2026-06-01 07:45:59 &#43;0000 UTC_: <code>08:15:29: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations-1.8/2061353417565016064#1:build-log.txt%3A1188)
<details><summary>context</summary>
<pre>08:12:09: time=&#34;2026-06-01T08:12:09Z&#34; level=info msg=&#34;Object whereabouts-config applied successfully&#34;
08:12:09: time=&#34;2026-06-01T08:12:09Z&#34; level=info msg=&#34;Object whereabouts applied successfully&#34;
08:12:09: time=&#34;2026-06-01T08:12:09Z&#34; level=info msg=&#34;[node 1]: kubectl --kubeconfig=/etc/kubernetes/admin.conf wait deployment -n cluster-network-addons cluster-network-addons-operator --for condition=Available --timeout=200s&#34;
08:15:29: error: timed out waiting for the condition on deployments/cluster-network-addons-operator
08:15:29:
08:15:29: ===== 8e7daae5841940ed0dd1d6f9ec4f689552893a675e77c04f646bc832901718b4 ====
08:15:29: &#43; NUM_NODES=3</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 2.86%) </summary>

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

### internal (5x / 14.29%)

<details>
<summary> deployment timeout (5x / 14.29%) </summary>

<hr/>

**5x**: _2026-06-01 07:46:02 &#43;0000 UTC_: <code>08:10:04: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2061353422568820736#1:build-log.txt%3A903)
<details>
<summary>all...</summary>

* _2026-06-01 14:01:39 &#43;0000 UTC_: <code>14:09:35: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-network/2061439596561960960#1:build-log.txt%3A1002)
<details><summary>context</summary>
<pre>14:06:15: time=&#34;2026-06-01T14:06:15Z&#34; level=info msg=&#34;Object whereabouts-config applied successfully&#34;
14:06:15: time=&#34;2026-06-01T14:06:15Z&#34; level=info msg=&#34;Object whereabouts applied successfully&#34;
14:06:15: time=&#34;2026-06-01T14:06:15Z&#34; level=info msg=&#34;[node 1]: kubectl --kubeconfig=/etc/kubernetes/admin.conf wait deployment -n cluster-network-addons cluster-network-addons-operator --for condition=Available --timeout=200s&#34;
14:09:35: error: timed out waiting for the condition on deployments/cluster-network-addons-operator
14:09:35:
14:09:35: ===== 3aee673b550f462b438293ae65a86d9f03c54de2ada1f1c52a982ff5e1e7ba94 ====
14:09:35: &#43; NUM_NODES=3</pre>
</details>


* _2026-06-01 07:46:18 &#43;0000 UTC_: <code>08:11:57: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network-1.8/2061353416696795136#1:build-log.txt%3A998)
<details><summary>context</summary>
<pre>08:08:37: time=&#34;2026-06-01T08:08:37Z&#34; level=info msg=&#34;Object whereabouts-config applied successfully&#34;
08:08:37: time=&#34;2026-06-01T08:08:37Z&#34; level=info msg=&#34;Object whereabouts applied successfully&#34;
08:08:37: time=&#34;2026-06-01T08:08:37Z&#34; level=info msg=&#34;[node 1]: kubectl --kubeconfig=/etc/kubernetes/admin.conf wait deployment -n cluster-network-addons cluster-network-addons-operator --for condition=Available --timeout=200s&#34;
08:11:57: error: timed out waiting for the condition on deployments/cluster-network-addons-operator
08:11:57:
08:11:57: ===== 5a5c668532533d2ed09c14b133b9e06bb5f52288fbea4cd85441854cc7de2160 ====
08:11:57: &#43; NUM_NODES=3</pre>
</details>


* _2026-06-01 07:46:07 &#43;0000 UTC_: <code>08:11:22: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-operator-1.8/2061353430110179328#1:build-log.txt%3A927)
<details><summary>context</summary>
<pre>08:08:02: time=&#34;2026-06-01T08:08:02Z&#34; level=info msg=&#34;Object whereabouts-config applied successfully&#34;
08:08:02: time=&#34;2026-06-01T08:08:02Z&#34; level=info msg=&#34;Object whereabouts applied successfully&#34;
08:08:02: time=&#34;2026-06-01T08:08:02Z&#34; level=info msg=&#34;[node 1]: kubectl --kubeconfig=/etc/kubernetes/admin.conf wait deployment -n cluster-network-addons cluster-network-addons-operator --for condition=Available --timeout=200s&#34;
08:11:22: error: timed out waiting for the condition on deployments/cluster-network-addons-operator
08:11:22:
08:11:22: ===== 1470dc57422f9799cef5b63e5f41a31aebf681286f0dc52d7db0151f29b77e93 ====
08:11:22: &#43; NUM_NODES=2</pre>
</details>


* _2026-06-01 07:46:06 &#43;0000 UTC_: <code>08:05:36: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.8/2061353425936846848#1:build-log.txt%3A911)
<details><summary>context</summary>
<pre>08:02:15: time=&#34;2026-06-01T08:02:15Z&#34; level=info msg=&#34;Object whereabouts-config applied successfully&#34;
08:02:16: time=&#34;2026-06-01T08:02:16Z&#34; level=info msg=&#34;Object whereabouts applied successfully&#34;
08:02:16: time=&#34;2026-06-01T08:02:16Z&#34; level=info msg=&#34;[node 1]: kubectl --kubeconfig=/etc/kubernetes/admin.conf wait deployment -n cluster-network-addons cluster-network-addons-operator --for condition=Available --timeout=200s&#34;
08:05:36: error: timed out waiting for the condition on deployments/cluster-network-addons-operator
08:05:36:
08:05:36: ===== 9fa84d931d50d8d7e574f028291bf61cabb4476b5c058449e9cea338c4db1e6d ====
08:05:36: &#43; NUM_NODES=2</pre>
</details>


* _2026-06-01 07:46:02 &#43;0000 UTC_: <code>08:10:04: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2061353422568820736#1:build-log.txt%3A903)
<details><summary>context</summary>
<pre>08:06:44: time=&#34;2026-06-01T08:06:44Z&#34; level=info msg=&#34;Object whereabouts-config applied successfully&#34;
08:06:44: time=&#34;2026-06-01T08:06:44Z&#34; level=info msg=&#34;Object whereabouts applied successfully&#34;
08:06:44: time=&#34;2026-06-01T08:06:44Z&#34; level=info msg=&#34;[node 1]: kubectl --kubeconfig=/etc/kubernetes/admin.conf wait deployment -n cluster-network-addons cluster-network-addons-operator --for condition=Available --timeout=200s&#34;
08:10:04: error: timed out waiting for the condition on deployments/cluster-network-addons-operator
08:10:04:
08:10:04: ===== 1e3704557f97e623e1b7e333080527a58687a4893f7c980442074b8068a43946 ====
08:10:04: &#43; NUM_NODES=2</pre>
</details>


</details>

<hr/>
</details>

### needs-investigation (1x / 2.86%)

<details>
<summary> no matching pattern (1x / 2.86%) </summary>

<hr/>

**1x**: _2026-06-01 20:05:06 &#43;0000 UTC_: <code>22:09:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-2qbf8-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 90e398de-9d0a-43a6-831b-6bfceadaf099]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-01T22:08:53.733535Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18006/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2061539398683463680#1:build-log.txt%3A12941)
<details>
<summary>all...</summary>

* _2026-06-01 20:05:06 &#43;0000 UTC_: <code>22:09:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-2qbf8-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 90e398de-9d0a-43a6-831b-6bfceadaf099]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-01T22:08:53.733535Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18006/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2061539398683463680#1:build-log.txt%3A12941)
<details><summary>context</summary>
<pre>22:09:02:   {&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;Sent: \&#34;mount /dev/vdc /mount_dir\\n\&#34;&#34;,&#34;pos&#34;:&#34;stdlib.go:105&#34;,&#34;timestamp&#34;:&#34;2026-06-01T22:08:51.586114Z&#34;,&#34;ts&#34;:&#34;2026/06/01 22:08:51&#34;}
22:09:02:   {&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;read closing down: EOF&#34;,&#34;pos&#34;:&#34;stdlib.go:105&#34;,&#34;timestamp&#34;:&#34;2026-06-01T22:08:53.586499Z&#34;,&#34;ts&#34;:&#34;2026/06/01 22:08:53&#34;}
22:09:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-2qbf8-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 90e398de-9d0a-43a6-831b-6bfceadaf099]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;VirtualMachineInstance soft rebooted&#34;,&#34;pos&#34;:&#34;watcher.go:159&#34;,&#34;timestamp&#34;:&#34;2026-06-01T22:08:53.684622Z&#34;}
22:09:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-2qbf8-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 90e398de-9d0a-43a6-831b-6bfceadaf099]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-01T22:08:53.733535Z&#34;}
22:09:02:   &lt;&lt; Captured StdOut/StdErr Output
22:09:02: ------------------------------
22:09:33: •••</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (6x / 17.14%)


#### external (5x / 83.33%)

<details>
<summary> container image pull failure in context (2x / 33.33%) </summary>

<hr/>

**2x**: _2026-06-01 14:05:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-compute/2061439596796841984#1:build-log.txt%3A239)
<details>
<summary>all...</summary>

* _2026-06-01 14:05:38 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-operator/2061439596926865408#1:build-log.txt%3A301)

* _2026-06-01 14:05:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-compute/2061439596796841984#1:build-log.txt%3A239)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (2x / 33.33%) </summary>

<hr/>

**1x**: _2026-06-01 14:03:07 &#43;0000 UTC_: <code>14:05:31: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-storage/2061439596666818560#1:build-log.txt%3A590)
<details>
<summary>all...</summary>

* _2026-06-01 14:03:07 &#43;0000 UTC_: <code>14:05:31: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-storage/2061439596666818560#1:build-log.txt%3A590)

</details>

<hr/>

**1x**: _2026-06-01 13:55:33 &#43;0000 UTC_: <code>14:07:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2061439596553572352#1:build-log.txt%3A4208)
<details>
<summary>all...</summary>

* _2026-06-01 13:55:33 &#43;0000 UTC_: <code>14:07:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2061439596553572352#1:build-log.txt%3A4208)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-06-05 00:01:46 &#43;0000 UTC_: <code>00:13:32: I0604 20:13:31.777281    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18028/pull-kubevirt-e2e-k8s-1.36-sig-network/2062686150438424576#1:build-log.txt%3A817)
<details>
<summary>all...</summary>

* _2026-06-05 00:01:46 &#43;0000 UTC_: <code>00:13:32: I0604 20:13:31.777281    1615 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18028/pull-kubevirt-e2e-k8s-1.36-sig-network/2062686150438424576#1:build-log.txt%3A817)

</details>

<hr/>
</details>

#### internal (1x / 16.67%)

<details>
<summary> deployment timeout (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-06-01 14:01:39 &#43;0000 UTC_: <code>14:09:35: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-network/2061439596561960960#1:build-log.txt%3A1002)
<details>
<summary>all...</summary>

* _2026-06-01 14:01:39 &#43;0000 UTC_: <code>14:09:35: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-network/2061439596561960960#1:build-log.txt%3A1002)

</details>

<hr/>
</details>

### release-1.7 (1x / 2.86%)


#### needs-investigation (1x / 100.00%)

<details>
<summary> no matching pattern (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-06-01 20:05:06 &#43;0000 UTC_: <code>22:09:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-2qbf8-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 90e398de-9d0a-43a6-831b-6bfceadaf099]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-01T22:08:53.733535Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18006/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2061539398683463680#1:build-log.txt%3A12941)
<details>
<summary>all...</summary>

* _2026-06-01 20:05:06 &#43;0000 UTC_: <code>22:09:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-2qbf8-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 90e398de-9d0a-43a6-831b-6bfceadaf099]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-01T22:08:53.733535Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18006/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2061539398683463680#1:build-log.txt%3A12941)

</details>

<hr/>
</details>

### release-1.8 (28x / 80.00%)


#### external (24x / 85.71%)

<details>
<summary> download failure in context (11x / 39.29%) </summary>

<hr/>

**11x**: _2026-06-01 07:46:15 &#43;0000 UTC_: <code>08:18:52: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2061353428428263424#1:build-log.txt%3A4109)
<details>
<summary>all...</summary>

* _2026-06-01 10:14:55 &#43;0000 UTC_: <code>10:35:12: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2061390371253915648#1:build-log.txt%3A3979)

* _2026-06-01 10:08:23 &#43;0000 UTC_: <code>10:29:53: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2061384109464227840#1:build-log.txt%3A3943)

* _2026-06-01 10:06:05 &#43;0000 UTC_: <code>10:32:21: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2061384108692475904#1:build-log.txt%3A4084)

* _2026-06-01 07:46:25 &#43;0000 UTC_: <code>08:13:45: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2061353425072820224#1:build-log.txt%3A3955)

* _2026-06-01 07:46:15 &#43;0000 UTC_: <code>08:18:52: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2061353428428263424#1:build-log.txt%3A4109)

* _2026-06-01 07:46:12 &#43;0000 UTC_: <code>08:18:27: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2061353420920459264#1:build-log.txt%3A4085)

* _2026-06-01 07:46:10 &#43;0000 UTC_: <code>08:14:59: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial-1.8/2061353426742153216#1:build-log.txt%3A3975)

* _2026-06-01 07:46:07 &#43;0000 UTC_: <code>08:14:29: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2061353421725765632#1:build-log.txt%3A3951)

* _2026-06-01 07:46:02 &#43;0000 UTC_: <code>08:14:34: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2061353429304872960#1:build-log.txt%3A3964)

* _2026-06-01 07:45:59 &#43;0000 UTC_: <code>08:13:28: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2061353408442404864#1:build-log.txt%3A3983)

* _2026-06-01 07:45:58 &#43;0000 UTC_: <code>08:03:49: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-kind-1.35-vgpu-1.8/2061353408597594112#1:build-log.txt%3A4388)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (8x / 28.57%) </summary>

<hr/>

**1x**: _2026-06-01 10:34:38 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-kind-1.35-vgpu-1.8/2061391210236350464#1:build-log.txt%3A184)
<details>
<summary>all...</summary>

* _2026-06-01 10:34:38 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-kind-1.35-vgpu-1.8/2061391210236350464#1:build-log.txt%3A184)

</details>

<hr/>

**7x**: _2026-06-01 10:27:12 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2061384111733346304#1:build-log.txt%3A599)
<details>
<summary>all...</summary>

* _2026-06-01 10:33:51 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-operator-1.8/2061384114241540096#1:build-log.txt%3A205)

* _2026-06-01 10:32:14 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2061384113398484992#1:build-log.txt%3A202)

* _2026-06-01 10:31:03 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2061384112605761536#1:build-log.txt%3A207)

* _2026-06-01 10:27:12 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2061384111733346304#1:build-log.txt%3A599)

* _2026-06-01 10:25:52 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial-1.8/2061384111007731712#1:build-log.txt%3A581)

* _2026-06-01 07:51:22 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-kind-sriov-1.8/2061353408887001088#1:build-log.txt%3A1671)

* _2026-06-01 07:46:34 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2061353424254930944#1:build-log.txt%3A606)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (5x / 17.86%) </summary>

<hr/>

**5x**: _2026-06-01 07:46:02 &#43;0000 UTC_: <code>08:12:03: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-network-1.8/2061353420056432640#1:build-log.txt%3A977)
<details>
<summary>all...</summary>

* _2026-06-01 10:18:48 &#43;0000 UTC_: <code>10:35:03: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.8/2061384110051430400#1:build-log.txt%3A909)

* _2026-06-01 07:46:18 &#43;0000 UTC_: <code>08:12:17: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2061353427622957056#1:build-log.txt%3A998)

* _2026-06-01 07:46:05 &#43;0000 UTC_: <code>08:09:47: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-network-1.8/2061353423407681536#1:build-log.txt%3A984)

* _2026-06-01 07:46:02 &#43;0000 UTC_: <code>08:12:03: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-network-1.8/2061353420056432640#1:build-log.txt%3A977)

* _2026-06-01 07:45:59 &#43;0000 UTC_: <code>08:15:29: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations-1.8/2061353417565016064#1:build-log.txt%3A1188)

</details>

<hr/>
</details>

#### internal (4x / 14.29%)

<details>
<summary> deployment timeout (4x / 14.29%) </summary>

<hr/>

**4x**: _2026-06-01 07:46:02 &#43;0000 UTC_: <code>08:10:04: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2061353422568820736#1:build-log.txt%3A903)
<details>
<summary>all...</summary>

* _2026-06-01 07:46:18 &#43;0000 UTC_: <code>08:11:57: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network-1.8/2061353416696795136#1:build-log.txt%3A998)

* _2026-06-01 07:46:07 &#43;0000 UTC_: <code>08:11:22: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-operator-1.8/2061353430110179328#1:build-log.txt%3A927)

* _2026-06-01 07:46:06 &#43;0000 UTC_: <code>08:05:36: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.8/2061353425936846848#1:build-log.txt%3A911)

* _2026-06-01 07:46:02 &#43;0000 UTC_: <code>08:10:04: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2061353422568820736#1:build-log.txt%3A903)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-network (9x / 25.71%)


#### external (7x / 77.78%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (3x / 33.33%) </summary>

<hr/>

**3x**: _2026-06-01 07:46:02 &#43;0000 UTC_: <code>08:12:03: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-network-1.8/2061353420056432640#1:build-log.txt%3A977)
<details>
<summary>all...</summary>

* _2026-06-01 07:46:18 &#43;0000 UTC_: <code>08:12:17: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2061353427622957056#1:build-log.txt%3A998)

* _2026-06-01 07:46:05 &#43;0000 UTC_: <code>08:09:47: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-network-1.8/2061353423407681536#1:build-log.txt%3A984)

* _2026-06-01 07:46:02 &#43;0000 UTC_: <code>08:12:03: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-network-1.8/2061353420056432640#1:build-log.txt%3A977)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (2x / 22.22%) </summary>

<hr/>

**2x**: _2026-06-01 10:27:12 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2061384111733346304#1:build-log.txt%3A599)
<details>
<summary>all...</summary>

* _2026-06-01 10:27:12 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-network-1.8/2061384111733346304#1:build-log.txt%3A599)

* _2026-06-01 07:51:22 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-kind-sriov-1.8/2061353408887001088#1:build-log.txt%3A1671)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-06-01 13:55:33 &#43;0000 UTC_: <code>14:07:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2061439596553572352#1:build-log.txt%3A4208)
<details>
<summary>all...</summary>

* _2026-06-01 13:55:33 &#43;0000 UTC_: <code>14:07:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2061439596553572352#1:build-log.txt%3A4208)

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

#### internal (2x / 22.22%)

<details>
<summary> deployment timeout (2x / 22.22%) </summary>

<hr/>

**2x**: _2026-06-01 14:01:39 &#43;0000 UTC_: <code>14:09:35: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-network/2061439596561960960#1:build-log.txt%3A1002)
<details>
<summary>all...</summary>

* _2026-06-01 14:01:39 &#43;0000 UTC_: <code>14:09:35: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-network/2061439596561960960#1:build-log.txt%3A1002)

* _2026-06-01 07:46:18 &#43;0000 UTC_: <code>08:11:57: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network-1.8/2061353416696795136#1:build-log.txt%3A998)

</details>

<hr/>
</details>

### sig-compute (19x / 54.29%)


#### external (16x / 84.21%)

<details>
<summary> download failure in context (8x / 42.11%) </summary>

<hr/>

**8x**: _2026-06-01 10:08:23 &#43;0000 UTC_: <code>10:29:53: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2061384109464227840#1:build-log.txt%3A3943)
<details>
<summary>all...</summary>

* _2026-06-01 10:14:55 &#43;0000 UTC_: <code>10:35:12: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2061390371253915648#1:build-log.txt%3A3979)

* _2026-06-01 10:08:23 &#43;0000 UTC_: <code>10:29:53: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2061384109464227840#1:build-log.txt%3A3943)

* _2026-06-01 07:46:25 &#43;0000 UTC_: <code>08:13:45: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2061353425072820224#1:build-log.txt%3A3955)

* _2026-06-01 07:46:10 &#43;0000 UTC_: <code>08:14:59: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial-1.8/2061353426742153216#1:build-log.txt%3A3975)

* _2026-06-01 07:46:07 &#43;0000 UTC_: <code>08:14:29: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2061353421725765632#1:build-log.txt%3A3951)

* _2026-06-01 07:46:02 &#43;0000 UTC_: <code>08:14:34: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2061353429304872960#1:build-log.txt%3A3964)

* _2026-06-01 07:45:59 &#43;0000 UTC_: <code>08:13:28: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2061353408442404864#1:build-log.txt%3A3983)

* _2026-06-01 07:45:58 &#43;0000 UTC_: <code>08:03:49: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-kind-1.35-vgpu-1.8/2061353408597594112#1:build-log.txt%3A4388)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (6x / 31.58%) </summary>

<hr/>

**3x**: _2026-06-01 14:05:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-compute/2061439596796841984#1:build-log.txt%3A239)
<details>
<summary>all...</summary>

* _2026-06-01 14:05:38 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-operator/2061439596926865408#1:build-log.txt%3A301)

* _2026-06-01 14:05:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-compute/2061439596796841984#1:build-log.txt%3A239)

* _2026-06-01 10:34:38 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-kind-1.35-vgpu-1.8/2061391210236350464#1:build-log.txt%3A184)

</details>

<hr/>

**3x**: _2026-06-01 10:25:52 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial-1.8/2061384111007731712#1:build-log.txt%3A581)
<details>
<summary>all...</summary>

* _2026-06-01 10:33:51 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-operator-1.8/2061384114241540096#1:build-log.txt%3A205)

* _2026-06-01 10:32:14 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-compute-1.8/2061384113398484992#1:build-log.txt%3A202)

* _2026-06-01 10:25:52 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial-1.8/2061384111007731712#1:build-log.txt%3A581)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 10.53%) </summary>

<hr/>

**2x**: _2026-06-01 07:45:59 &#43;0000 UTC_: <code>08:15:29: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations-1.8/2061353417565016064#1:build-log.txt%3A1188)
<details>
<summary>all...</summary>

* _2026-06-01 10:18:48 &#43;0000 UTC_: <code>10:35:03: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.8/2061384110051430400#1:build-log.txt%3A909)

* _2026-06-01 07:45:59 &#43;0000 UTC_: <code>08:15:29: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations-1.8/2061353417565016064#1:build-log.txt%3A1188)

</details>

<hr/>
</details>

#### internal (3x / 15.79%)

<details>
<summary> deployment timeout (3x / 15.79%) </summary>

<hr/>

**3x**: _2026-06-01 07:46:02 &#43;0000 UTC_: <code>08:10:04: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2061353422568820736#1:build-log.txt%3A903)
<details>
<summary>all...</summary>

* _2026-06-01 07:46:07 &#43;0000 UTC_: <code>08:11:22: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-operator-1.8/2061353430110179328#1:build-log.txt%3A927)

* _2026-06-01 07:46:06 &#43;0000 UTC_: <code>08:05:36: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.8/2061353425936846848#1:build-log.txt%3A911)

* _2026-06-01 07:46:02 &#43;0000 UTC_: <code>08:10:04: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2061353422568820736#1:build-log.txt%3A903)

</details>

<hr/>
</details>

### sig-storage (7x / 20.00%)


#### external (6x / 85.71%)

<details>
<summary> download failure in context (4x / 57.14%) </summary>

<hr/>

**1x**: _2026-06-01 14:03:07 &#43;0000 UTC_: <code>14:05:31: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-storage/2061439596666818560#1:build-log.txt%3A590)
<details>
<summary>all...</summary>

* _2026-06-01 14:03:07 &#43;0000 UTC_: <code>14:05:31: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-storage/2061439596666818560#1:build-log.txt%3A590)

</details>

<hr/>

**3x**: _2026-06-01 07:46:15 &#43;0000 UTC_: <code>08:18:52: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2061353428428263424#1:build-log.txt%3A4109)
<details>
<summary>all...</summary>

* _2026-06-01 10:06:05 &#43;0000 UTC_: <code>10:32:21: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2061384108692475904#1:build-log.txt%3A4084)

* _2026-06-01 07:46:15 &#43;0000 UTC_: <code>08:18:52: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2061353428428263424#1:build-log.txt%3A4109)

* _2026-06-01 07:46:12 &#43;0000 UTC_: <code>08:18:27: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2061353420920459264#1:build-log.txt%3A4085)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (2x / 28.57%) </summary>

<hr/>

**2x**: _2026-06-01 10:31:03 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2061384112605761536#1:build-log.txt%3A207)
<details>
<summary>all...</summary>

* _2026-06-01 10:31:03 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2061384112605761536#1:build-log.txt%3A207)

* _2026-06-01 07:46:34 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17990/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2061353424254930944#1:build-log.txt%3A606)

</details>

<hr/>
</details>

#### needs-investigation (1x / 14.29%)

<details>
<summary> no matching pattern (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-06-01 20:05:06 &#43;0000 UTC_: <code>22:09:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-2qbf8-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 90e398de-9d0a-43a6-831b-6bfceadaf099]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-01T22:08:53.733535Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18006/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2061539398683463680#1:build-log.txt%3A12941)
<details>
<summary>all...</summary>

* _2026-06-01 20:05:06 &#43;0000 UTC_: <code>22:09:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-2qbf8-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 90e398de-9d0a-43a6-831b-6bfceadaf099]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-01T22:08:53.733535Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18006/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2061539398683463680#1:build-log.txt%3A12941)

</details>

<hr/>
</details>

Last updated: 2026-06-07 00:36:05
