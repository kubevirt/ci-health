



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-09-30 (1x / 33.33%)


#### external (1x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-30 06:40:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19205/pull-kubevirt-e2e-kind-1.37-vgpu/2105185961917812736#1:build-log.txt%3A447)
<details>
<summary>all...</summary>

* _2026-09-30 06:40:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19205/pull-kubevirt-e2e-kind-1.37-vgpu/2105185961917812736#1:build-log.txt%3A447)

</details>

<hr/>
</details>

### 2026-09-28 (1x / 33.33%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-28 07:04:26 &#43;0000 UTC_: <code>07:39:45: I0928 03:39:45.854640    1613 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19237/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2104467039153295360#1:build-log.txt%3A774)
<details>
<summary>all...</summary>

* _2026-09-28 07:04:26 &#43;0000 UTC_: <code>07:39:45: I0928 03:39:45.854640    1613 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19237/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2104467039153295360#1:build-log.txt%3A774)

</details>

<hr/>
</details>

### 2026-09-27 (1x / 33.33%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-27 06:50:39 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19163/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2104101244820787200#1:build-log.txt%3A1856)
<details>
<summary>all...</summary>

* _2026-09-27 06:50:39 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19163/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2104101244820787200#1:build-log.txt%3A1856)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (3x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-09-30 06:40:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19205/pull-kubevirt-e2e-kind-1.37-vgpu/2105185961917812736#1:build-log.txt%3A447)
<details>
<summary>all...</summary>

* _2026-09-30 06:40:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19205/pull-kubevirt-e2e-kind-1.37-vgpu/2105185961917812736#1:build-log.txt%3A447)
<details><summary>context</summary>
<pre>06:47:27: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
06:47:27: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
06:47:33: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1374624727/2&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20260930%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20260930T064729Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=bece4d68dfe552cd5c1522b88171d63dfede14193b999ace37b04b86befc4490&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1790751749~hmac=c79e58b95b20823cf63d13f7ca0f01b31728757402ba1ad25e2412552646d16d&#34;: EOF)
make: *** [Makefile:177: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-09-28 07:04:26 &#43;0000 UTC_: <code>07:39:45: I0928 03:39:45.854640    1613 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19237/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2104467039153295360#1:build-log.txt%3A774)
<details>
<summary>all...</summary>

* _2026-09-28 07:04:26 &#43;0000 UTC_: <code>07:39:45: I0928 03:39:45.854640    1613 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19237/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2104467039153295360#1:build-log.txt%3A774)
<details><summary>context</summary>
<pre>07:39:41: [control-plane-check] Checking kube-scheduler at https://127.0.0.1:10259/livez
07:39:43: [control-plane-check] kube-controller-manager is healthy after 2.357893091s
07:39:44: [control-plane-check] kube-scheduler is healthy after 3.627448341s
07:39:45: I0928 03:39:45.854640    1613 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
07:39:46: [control-plane-check] kube-apiserver is healthy after 5.503634325s
07:39:46: I0928 03:39:46.362671    1613 kubeconfig.go:657] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists
07:39:46: I0928 03:39:46.365956    1613 kubeconfig.go:730] creating the ClusterRoleBinding for the kubeadm:cluster-admins Group by using super-admin.conf</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-09-27 06:50:39 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19163/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2104101244820787200#1:build-log.txt%3A1856)
<details>
<summary>all...</summary>

* _2026-09-27 06:50:39 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19163/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2104101244820787200#1:build-log.txt%3A1856)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
06:55:45: selecting podman as container runtime
06:56:25: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network is being used by the following container(s): b13460666ba7f6b0bc2166dc3770d28f92ee63645cb66f490be56ec0d6ee5751: volume is being used
make: *** [Makefile:180: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (2x / 66.67%)


#### external (2x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-09-30 06:40:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19205/pull-kubevirt-e2e-kind-1.37-vgpu/2105185961917812736#1:build-log.txt%3A447)
<details>
<summary>all...</summary>

* _2026-09-30 06:40:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19205/pull-kubevirt-e2e-kind-1.37-vgpu/2105185961917812736#1:build-log.txt%3A447)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-09-27 06:50:39 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19163/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2104101244820787200#1:build-log.txt%3A1856)
<details>
<summary>all...</summary>

* _2026-09-27 06:50:39 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19163/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2104101244820787200#1:build-log.txt%3A1856)

</details>

<hr/>
</details>

### release-1.8 (1x / 33.33%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-28 07:04:26 &#43;0000 UTC_: <code>07:39:45: I0928 03:39:45.854640    1613 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19237/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2104467039153295360#1:build-log.txt%3A774)
<details>
<summary>all...</summary>

* _2026-09-28 07:04:26 &#43;0000 UTC_: <code>07:39:45: I0928 03:39:45.854640    1613 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19237/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2104467039153295360#1:build-log.txt%3A774)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (1x / 33.33%)


#### external (1x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-30 06:40:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19205/pull-kubevirt-e2e-kind-1.37-vgpu/2105185961917812736#1:build-log.txt%3A447)
<details>
<summary>all...</summary>

* _2026-09-30 06:40:56 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19205/pull-kubevirt-e2e-kind-1.37-vgpu/2105185961917812736#1:build-log.txt%3A447)

</details>

<hr/>
</details>

### sig-monitoring (1x / 33.33%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-28 07:04:26 &#43;0000 UTC_: <code>07:39:45: I0928 03:39:45.854640    1613 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19237/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2104467039153295360#1:build-log.txt%3A774)
<details>
<summary>all...</summary>

* _2026-09-28 07:04:26 &#43;0000 UTC_: <code>07:39:45: I0928 03:39:45.854640    1613 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19237/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2104467039153295360#1:build-log.txt%3A774)

</details>

<hr/>
</details>

### sig-network (1x / 33.33%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-27 06:50:39 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19163/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2104101244820787200#1:build-log.txt%3A1856)
<details>
<summary>all...</summary>

* _2026-09-27 06:50:39 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19163/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2104101244820787200#1:build-log.txt%3A1856)

</details>

<hr/>
</details>

Last updated: 2026-10-01 09:12:44
