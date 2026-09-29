



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-09-28 (1x / 25.00%)


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

### 2026-09-27 (1x / 25.00%)


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

### 2026-09-24 (2x / 50.00%)


#### external (2x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-09-24 06:12:02 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17922/pull-kubevirt-e2e-kind-1.37-vgpu/2103004361624915968#1:build-log.txt%3A509)
<details>
<summary>all...</summary>

* _2026-09-24 06:12:02 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17922/pull-kubevirt-e2e-kind-1.37-vgpu/2103004361624915968#1:build-log.txt%3A509)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-09-24 08:55:19 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19214/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2103045438079766528#1:build-log.txt%3A1865)
<details>
<summary>all...</summary>

* _2026-09-24 08:55:19 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19214/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2103045438079766528#1:build-log.txt%3A1865)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (4x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 50.00%) </summary>

<hr/>

**2x**: _2026-09-27 06:50:39 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19163/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2104101244820787200#1:build-log.txt%3A1856)
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


* _2026-09-24 08:55:19 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19214/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2103045438079766528#1:build-log.txt%3A1865)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
09:02:35: selecting podman as container runtime
09:03:08: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network is being used by the following container(s): 6e8772b9281a5e60a84b8770a1de0b028641b03275abaca545336f27093cc63e: volume is being used
make: *** [Makefile:180: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-09-24 06:12:02 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17922/pull-kubevirt-e2e-kind-1.37-vgpu/2103004361624915968#1:build-log.txt%3A509)
<details>
<summary>all...</summary>

* _2026-09-24 06:12:02 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17922/pull-kubevirt-e2e-kind-1.37-vgpu/2103004361624915968#1:build-log.txt%3A509)
<details><summary>context</summary>
<pre>06:25:34: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
06:25:34: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
06:25:38: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1519539013/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/5c/5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20260924%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20260924T062537Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=2f35355f57c757563f50779ec88c0ed96d02ad9028d2be72e47e0d8aada528dc&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1790232037~hmac=df0af00d89f54e51b4a3357eb77394f301d234403168562cec1be3addd950e42&#34;: EOF)
make: *** [Makefile:177: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 25.00%) </summary>

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

<a id="per-branch"></a>

## per branch [⬆](#top)


### release-1.8 (1x / 25.00%)


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

### main (3x / 75.00%)


#### external (3x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-09-27 06:50:39 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19163/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2104101244820787200#1:build-log.txt%3A1856)
<details>
<summary>all...</summary>

* _2026-09-27 06:50:39 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19163/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2104101244820787200#1:build-log.txt%3A1856)

* _2026-09-24 08:55:19 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19214/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2103045438079766528#1:build-log.txt%3A1865)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-09-24 06:12:02 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17922/pull-kubevirt-e2e-kind-1.37-vgpu/2103004361624915968#1:build-log.txt%3A509)
<details>
<summary>all...</summary>

* _2026-09-24 06:12:02 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17922/pull-kubevirt-e2e-kind-1.37-vgpu/2103004361624915968#1:build-log.txt%3A509)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-monitoring (1x / 25.00%)


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

### sig-network (2x / 50.00%)


#### external (2x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-09-27 06:50:39 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19163/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2104101244820787200#1:build-log.txt%3A1856)
<details>
<summary>all...</summary>

* _2026-09-27 06:50:39 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19163/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2104101244820787200#1:build-log.txt%3A1856)

* _2026-09-24 08:55:19 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19214/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2103045438079766528#1:build-log.txt%3A1865)

</details>

<hr/>
</details>

### sig-compute (1x / 25.00%)


#### external (1x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-24 06:12:02 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17922/pull-kubevirt-e2e-kind-1.37-vgpu/2103004361624915968#1:build-log.txt%3A509)
<details>
<summary>all...</summary>

* _2026-09-24 06:12:02 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17922/pull-kubevirt-e2e-kind-1.37-vgpu/2103004361624915968#1:build-log.txt%3A509)

</details>

<hr/>
</details>

Last updated: 2026-09-29 18:10:14
