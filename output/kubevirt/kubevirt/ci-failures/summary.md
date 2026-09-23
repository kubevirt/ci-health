



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-09-18 (1x / 20.00%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-18 07:56:21 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19175/pull-kubevirt-e2e-k8s-1.36-sig-network/2100856101317644288#1:build-log.txt%3A1879)
<details>
<summary>all...</summary>

* _2026-09-18 07:56:21 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19175/pull-kubevirt-e2e-k8s-1.36-sig-network/2100856101317644288#1:build-log.txt%3A1879)

</details>

<hr/>
</details>

### 2026-09-17 (3x / 60.00%)


#### external (2x / 66.67%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-09-17 13:43:51 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19164/pull-kubevirt-e2e-k8s-1.36-sig-network/2100562626076479488#1:build-log.txt%3A1882)
<details>
<summary>all...</summary>

* _2026-09-17 13:43:51 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19164/pull-kubevirt-e2e-k8s-1.36-sig-network/2100562626076479488#1:build-log.txt%3A1882)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-09-17 07:29:29 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19162/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2100487140398862336#1:build-log.txt%3A666)
<details>
<summary>all...</summary>

* _2026-09-17 07:29:29 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19162/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2100487140398862336#1:build-log.txt%3A666)

</details>

<hr/>
</details>

#### internal (1x / 33.33%)

<details>
<summary> make cluster lifecycle target failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-09-17 03:05:06 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19150/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2100420610688552960#1:build-log.txt%3A2408)
<details>
<summary>all...</summary>

* _2026-09-17 03:05:06 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19150/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2100420610688552960#1:build-log.txt%3A2408)

</details>

<hr/>
</details>

### 2026-09-16 (1x / 20.00%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-16 08:01:57 &#43;0000 UTC_: <code>08:37:20: I0916 04:37:20.839064    1602 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19139/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2100132912123875328#1:build-log.txt%3A770)
<details>
<summary>all...</summary>

* _2026-09-16 08:01:57 &#43;0000 UTC_: <code>08:37:20: I0916 04:37:20.839064    1602 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19139/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2100132912123875328#1:build-log.txt%3A770)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (4x / 80.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 40.00%) </summary>

<hr/>

**2x**: _2026-09-18 07:56:21 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19175/pull-kubevirt-e2e-k8s-1.36-sig-network/2100856101317644288#1:build-log.txt%3A1879)
<details>
<summary>all...</summary>

* _2026-09-18 07:56:21 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19175/pull-kubevirt-e2e-k8s-1.36-sig-network/2100856101317644288#1:build-log.txt%3A1879)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
08:01:45: selecting podman as container runtime
08:02:38: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network is being used by the following container(s): 2b5cd12f54fcc250811be6b95bda70c694cd5ce6f6bc567b49fc54e6c1af615a: volume is being used
make: *** [Makefile:180: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-09-17 13:43:51 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19164/pull-kubevirt-e2e-k8s-1.36-sig-network/2100562626076479488#1:build-log.txt%3A1882)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
13:49:21: selecting podman as container runtime
13:50:01: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network is being used by the following container(s): 4d689305039cfb27be7198af2779554127f7e4db8b2454a1d0ebff95d7c37200: volume is being used
make: *** [Makefile:180: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-09-16 08:01:57 &#43;0000 UTC_: <code>08:37:20: I0916 04:37:20.839064    1602 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19139/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2100132912123875328#1:build-log.txt%3A770)
<details>
<summary>all...</summary>

* _2026-09-16 08:01:57 &#43;0000 UTC_: <code>08:37:20: I0916 04:37:20.839064    1602 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19139/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2100132912123875328#1:build-log.txt%3A770)
<details><summary>context</summary>
<pre>08:37:15: [control-plane-check] Checking kube-scheduler at https://127.0.0.1:10259/livez
08:37:17: [control-plane-check] kube-controller-manager is healthy after 1.585120019s
08:37:19: [control-plane-check] kube-scheduler is healthy after 3.577089938s
08:37:20: I0916 04:37:20.839064    1602 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
08:37:21: [control-plane-check] kube-apiserver is healthy after 5.502248038s
08:37:21: I0916 04:37:21.347717    1602 kubeconfig.go:657] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists
08:37:21: I0916 04:37:21.349907    1602 kubeconfig.go:730] creating the ClusterRoleBinding for the kubeadm:cluster-admins Group by using super-admin.conf</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-09-17 07:29:29 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19162/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2100487140398862336#1:build-log.txt%3A666)
<details>
<summary>all...</summary>

* _2026-09-17 07:29:29 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19162/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2100487140398862336#1:build-log.txt%3A666)
<details><summary>context</summary>
<pre>07:37:18: Copying blob sha256:fead7c895b0c4ee4f33b3a53f800ad5f2e8868176dfb1841f3b68dfc2430d880
07:37:18: Copying blob sha256:03167ee8e16d8d0d85197f473bfb43c9d3c405ade222087024f8ab7128883a99
07:37:19: Error: unable to copy from source docker://quay.io/kubevirtci/gocli:2602190951-056f8f79: writing blob: storing blob to file &#34;/var/tmp/container_images_storage2181722017/1&#34;: happened during read: (heuristic tuning data: total 2097152 @1240.448 ms, last retry 2097152 @0.104 ms, last progress @ 1190.166 ms): unexpected EOF
make: *** [Makefile:176: cluster-down] Error 125
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>

### internal (1x / 20.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-09-17 03:05:06 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19150/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2100420610688552960#1:build-log.txt%3A2408)
<details>
<summary>all...</summary>

* _2026-09-17 03:05:06 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19150/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2100420610688552960#1:build-log.txt%3A2408)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
03:20:40: selecting podman as container runtime
03:21:14: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9 is being used by the following container(s): f49ae0abc2e801f01279bd8a77aa8af54e7dd93c1d566f85900c01ac6d3acba7: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (2x / 40.00%)


#### external (2x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-09-18 07:56:21 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19175/pull-kubevirt-e2e-k8s-1.36-sig-network/2100856101317644288#1:build-log.txt%3A1879)
<details>
<summary>all...</summary>

* _2026-09-18 07:56:21 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19175/pull-kubevirt-e2e-k8s-1.36-sig-network/2100856101317644288#1:build-log.txt%3A1879)

* _2026-09-17 13:43:51 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19164/pull-kubevirt-e2e-k8s-1.36-sig-network/2100562626076479488#1:build-log.txt%3A1882)

</details>

<hr/>
</details>

### release-1.8 (2x / 40.00%)


#### external (2x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-09-16 08:01:57 &#43;0000 UTC_: <code>08:37:20: I0916 04:37:20.839064    1602 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19139/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2100132912123875328#1:build-log.txt%3A770)
<details>
<summary>all...</summary>

* _2026-09-16 08:01:57 &#43;0000 UTC_: <code>08:37:20: I0916 04:37:20.839064    1602 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19139/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2100132912123875328#1:build-log.txt%3A770)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-09-17 07:29:29 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19162/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2100487140398862336#1:build-log.txt%3A666)
<details>
<summary>all...</summary>

* _2026-09-17 07:29:29 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19162/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2100487140398862336#1:build-log.txt%3A666)

</details>

<hr/>
</details>

### release-1.9 (1x / 20.00%)


#### internal (1x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-17 03:05:06 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19150/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2100420610688552960#1:build-log.txt%3A2408)
<details>
<summary>all...</summary>

* _2026-09-17 03:05:06 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19150/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2100420610688552960#1:build-log.txt%3A2408)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (1x / 20.00%)


#### external (1x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-17 07:29:29 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19162/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2100487140398862336#1:build-log.txt%3A666)
<details>
<summary>all...</summary>

* _2026-09-17 07:29:29 &#43;0000 UTC_: <code>make: *** [Makefile:176: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19162/pull-kubevirt-e2e-k8s-1.33-sig-operator-1.8/2100487140398862336#1:build-log.txt%3A666)

</details>

<hr/>
</details>

### sig-network (3x / 60.00%)


#### internal (1x / 33.33%)

<details>
<summary> make cluster lifecycle target failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-09-17 03:05:06 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19150/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2100420610688552960#1:build-log.txt%3A2408)
<details>
<summary>all...</summary>

* _2026-09-17 03:05:06 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19150/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2100420610688552960#1:build-log.txt%3A2408)

</details>

<hr/>
</details>

#### external (2x / 66.67%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-09-18 07:56:21 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19175/pull-kubevirt-e2e-k8s-1.36-sig-network/2100856101317644288#1:build-log.txt%3A1879)
<details>
<summary>all...</summary>

* _2026-09-18 07:56:21 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19175/pull-kubevirt-e2e-k8s-1.36-sig-network/2100856101317644288#1:build-log.txt%3A1879)

* _2026-09-17 13:43:51 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19164/pull-kubevirt-e2e-k8s-1.36-sig-network/2100562626076479488#1:build-log.txt%3A1882)

</details>

<hr/>
</details>

### sig-monitoring (1x / 20.00%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-09-16 08:01:57 &#43;0000 UTC_: <code>08:37:20: I0916 04:37:20.839064    1602 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19139/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2100132912123875328#1:build-log.txt%3A770)
<details>
<summary>all...</summary>

* _2026-09-16 08:01:57 &#43;0000 UTC_: <code>08:37:20: I0916 04:37:20.839064    1602 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19139/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2100132912123875328#1:build-log.txt%3A770)

</details>

<hr/>
</details>

Last updated: 2026-09-23 03:13:21
