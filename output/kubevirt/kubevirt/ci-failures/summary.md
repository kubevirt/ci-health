



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-09-18 (1x / 10.00%)


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

### 2026-09-17 (3x / 30.00%)


#### external (2x / 66.67%)

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

### 2026-09-16 (1x / 10.00%)


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

### 2026-09-14 (5x / 50.00%)


#### external (5x / 100.00%)

<details>
<summary> download failure in context (4x / 80.00%) </summary>

<hr/>

**2x**: _2026-09-14 14:28:33 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18841/pull-kubevirt-e2e-kind-1.36-sig-compute-arm64/2099505469235138560#1:build-log.txt%3A254)
<details>
<summary>all...</summary>

* _2026-09-14 14:29:09 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18841/pull-kubevirt-e2e-kind-1.36-sev/2099505474956169216#1:build-log.txt%3A245)

* _2026-09-14 14:28:33 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18841/pull-kubevirt-e2e-kind-1.36-sig-compute-arm64/2099505469235138560#1:build-log.txt%3A254)

</details>

<hr/>

**2x**: _2026-09-14 14:19:35 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19052/pull-kubevirt-e2e-k8s-1.37-sig-network/2099500599073574912#1:build-log.txt%3A4018)
<details>
<summary>all...</summary>

* _2026-09-14 14:29:11 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18841/pull-kubevirt-e2e-k8s-1.36-sig-storage/2099505470753476608#1:build-log.txt%3A3038)

* _2026-09-14 14:19:35 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19052/pull-kubevirt-e2e-k8s-1.37-sig-network/2099500599073574912#1:build-log.txt%3A4018)

</details>

<hr/>
</details>
<details>
<summary> podman container removal timeout (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-09-14 11:12:09 &#43;0000 UTC_: <code>Error: cannot remove container 8dda5b7c95ec4de1ce22ee25296269c9781d5bd8e5f18b79922c436d148852b6 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19041/pull-kubevirt-e2e-k8s-1.35-sig-storage/2099456004369420288#1:build-log.txt%3A3068)
<details>
<summary>all...</summary>

* _2026-09-14 11:12:09 &#43;0000 UTC_: <code>Error: cannot remove container 8dda5b7c95ec4de1ce22ee25296269c9781d5bd8e5f18b79922c436d148852b6 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19041/pull-kubevirt-e2e-k8s-1.35-sig-storage/2099456004369420288#1:build-log.txt%3A3068)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (9x / 90.00%)

<details>
<summary> download failure in context (4x / 40.00%) </summary>

<hr/>

**2x**: _2026-09-14 14:28:33 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18841/pull-kubevirt-e2e-kind-1.36-sig-compute-arm64/2099505469235138560#1:build-log.txt%3A254)
<details>
<summary>all...</summary>

* _2026-09-14 14:29:09 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18841/pull-kubevirt-e2e-kind-1.36-sev/2099505474956169216#1:build-log.txt%3A245)
<details><summary>context</summary>
<pre>14:31:38: INFO: 1 process: 1 internal.
14:31:38: ERROR: Build did NOT complete successfully
14:31:38: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-09-14 14:28:33 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18841/pull-kubevirt-e2e-kind-1.36-sig-compute-arm64/2099505469235138560#1:build-log.txt%3A254)
<details><summary>context</summary>
<pre>14:32:36: INFO: 0 processes.
14:32:36: ERROR: Build did NOT complete successfully
14:32:36: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


</details>

<hr/>

**2x**: _2026-09-14 14:19:35 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19052/pull-kubevirt-e2e-k8s-1.37-sig-network/2099500599073574912#1:build-log.txt%3A4018)
<details>
<summary>all...</summary>

* _2026-09-14 14:29:11 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18841/pull-kubevirt-e2e-k8s-1.36-sig-storage/2099505470753476608#1:build-log.txt%3A3038)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
14:52:59: selecting podman as container runtime
14:53:42: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-storage is being used by the following container(s): 7fe7d07cdbb2d5bc1890d013cea14b7f484021756bbe39e99495f2ee1967d11e: volume is being used
make: *** [Makefile:180: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-09-14 14:19:35 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19052/pull-kubevirt-e2e-k8s-1.37-sig-network/2099500599073574912#1:build-log.txt%3A4018)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
14:36:52: selecting podman as container runtime
14:37:31: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.37-sig-network is being used by the following container(s): 73078e928d0f0cb69b0e44b1cec7cf0f69114147ddfc67bfeff47f47ba3b3268: volume is being used
make: *** [Makefile:180: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 20.00%) </summary>

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
<summary> container image pull failure in context (1x / 10.00%) </summary>

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
<details>
<summary> transient kube-apiserver body decode noise (1x / 10.00%) </summary>

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
<summary> podman container removal timeout (1x / 10.00%) </summary>

<hr/>

**1x**: _2026-09-14 11:12:09 &#43;0000 UTC_: <code>Error: cannot remove container 8dda5b7c95ec4de1ce22ee25296269c9781d5bd8e5f18b79922c436d148852b6 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19041/pull-kubevirt-e2e-k8s-1.35-sig-storage/2099456004369420288#1:build-log.txt%3A3068)
<details>
<summary>all...</summary>

* _2026-09-14 11:12:09 &#43;0000 UTC_: <code>Error: cannot remove container 8dda5b7c95ec4de1ce22ee25296269c9781d5bd8e5f18b79922c436d148852b6 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19041/pull-kubevirt-e2e-k8s-1.35-sig-storage/2099456004369420288#1:build-log.txt%3A3068)
<details><summary>context</summary>
<pre>ab442c0f419f7b9eed6a132f40f3c7ae8ffdfe8b032f096352a14062b314af70
8231b50111059f2c9588ee42058ad906a3528dbe61f80a3d69ea7c1720508460
Error: cannot remove container 8dda5b7c95ec4de1ce22ee25296269c9781d5bd8e5f18b79922c436d148852b6 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 8dda5b7c95ec4de1ce22ee25296269c9781d5bd8e5f18b79922c436d148852b6 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1020 is not a child of this shell
================================================================================
Done cleaning up after podman in container.</pre>
</details>


</details>

<hr/>
</details>

### internal (1x / 10.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 10.00%) </summary>

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


### main (7x / 70.00%)


#### external (7x / 100.00%)

<details>
<summary> download failure in context (4x / 57.14%) </summary>

<hr/>

**2x**: _2026-09-14 14:28:33 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18841/pull-kubevirt-e2e-kind-1.36-sig-compute-arm64/2099505469235138560#1:build-log.txt%3A254)
<details>
<summary>all...</summary>

* _2026-09-14 14:29:09 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18841/pull-kubevirt-e2e-kind-1.36-sev/2099505474956169216#1:build-log.txt%3A245)

* _2026-09-14 14:28:33 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18841/pull-kubevirt-e2e-kind-1.36-sig-compute-arm64/2099505469235138560#1:build-log.txt%3A254)

</details>

<hr/>

**2x**: _2026-09-14 14:19:35 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19052/pull-kubevirt-e2e-k8s-1.37-sig-network/2099500599073574912#1:build-log.txt%3A4018)
<details>
<summary>all...</summary>

* _2026-09-14 14:29:11 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18841/pull-kubevirt-e2e-k8s-1.36-sig-storage/2099505470753476608#1:build-log.txt%3A3038)

* _2026-09-14 14:19:35 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19052/pull-kubevirt-e2e-k8s-1.37-sig-network/2099500599073574912#1:build-log.txt%3A4018)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 28.57%) </summary>

<hr/>

**2x**: _2026-09-18 07:56:21 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19175/pull-kubevirt-e2e-k8s-1.36-sig-network/2100856101317644288#1:build-log.txt%3A1879)
<details>
<summary>all...</summary>

* _2026-09-18 07:56:21 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19175/pull-kubevirt-e2e-k8s-1.36-sig-network/2100856101317644288#1:build-log.txt%3A1879)

* _2026-09-17 13:43:51 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19164/pull-kubevirt-e2e-k8s-1.36-sig-network/2100562626076479488#1:build-log.txt%3A1882)

</details>

<hr/>
</details>
<details>
<summary> podman container removal timeout (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-09-14 11:12:09 &#43;0000 UTC_: <code>Error: cannot remove container 8dda5b7c95ec4de1ce22ee25296269c9781d5bd8e5f18b79922c436d148852b6 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19041/pull-kubevirt-e2e-k8s-1.35-sig-storage/2099456004369420288#1:build-log.txt%3A3068)
<details>
<summary>all...</summary>

* _2026-09-14 11:12:09 &#43;0000 UTC_: <code>Error: cannot remove container 8dda5b7c95ec4de1ce22ee25296269c9781d5bd8e5f18b79922c436d148852b6 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19041/pull-kubevirt-e2e-k8s-1.35-sig-storage/2099456004369420288#1:build-log.txt%3A3068)

</details>

<hr/>
</details>

### release-1.8 (2x / 20.00%)


#### external (2x / 100.00%)

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

### release-1.9 (1x / 10.00%)


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


### sig-network (4x / 40.00%)


#### external (3x / 75.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 50.00%) </summary>

<hr/>

**2x**: _2026-09-18 07:56:21 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19175/pull-kubevirt-e2e-k8s-1.36-sig-network/2100856101317644288#1:build-log.txt%3A1879)
<details>
<summary>all...</summary>

* _2026-09-18 07:56:21 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19175/pull-kubevirt-e2e-k8s-1.36-sig-network/2100856101317644288#1:build-log.txt%3A1879)

* _2026-09-17 13:43:51 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19164/pull-kubevirt-e2e-k8s-1.36-sig-network/2100562626076479488#1:build-log.txt%3A1882)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-09-14 14:19:35 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19052/pull-kubevirt-e2e-k8s-1.37-sig-network/2099500599073574912#1:build-log.txt%3A4018)
<details>
<summary>all...</summary>

* _2026-09-14 14:19:35 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19052/pull-kubevirt-e2e-k8s-1.37-sig-network/2099500599073574912#1:build-log.txt%3A4018)

</details>

<hr/>
</details>

#### internal (1x / 25.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-09-17 03:05:06 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19150/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2100420610688552960#1:build-log.txt%3A2408)
<details>
<summary>all...</summary>

* _2026-09-17 03:05:06 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19150/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2100420610688552960#1:build-log.txt%3A2408)

</details>

<hr/>
</details>

### sig-monitoring (1x / 10.00%)


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

### sig-storage (2x / 20.00%)


#### external (2x / 100.00%)

<details>
<summary> download failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-09-14 14:29:11 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18841/pull-kubevirt-e2e-k8s-1.36-sig-storage/2099505470753476608#1:build-log.txt%3A3038)
<details>
<summary>all...</summary>

* _2026-09-14 14:29:11 &#43;0000 UTC_: <code>make: *** [Makefile:180: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18841/pull-kubevirt-e2e-k8s-1.36-sig-storage/2099505470753476608#1:build-log.txt%3A3038)

</details>

<hr/>
</details>
<details>
<summary> podman container removal timeout (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-09-14 11:12:09 &#43;0000 UTC_: <code>Error: cannot remove container 8dda5b7c95ec4de1ce22ee25296269c9781d5bd8e5f18b79922c436d148852b6 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19041/pull-kubevirt-e2e-k8s-1.35-sig-storage/2099456004369420288#1:build-log.txt%3A3068)
<details>
<summary>all...</summary>

* _2026-09-14 11:12:09 &#43;0000 UTC_: <code>Error: cannot remove container 8dda5b7c95ec4de1ce22ee25296269c9781d5bd8e5f18b79922c436d148852b6 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19041/pull-kubevirt-e2e-k8s-1.35-sig-storage/2099456004369420288#1:build-log.txt%3A3068)

</details>

<hr/>
</details>

### sig-compute (3x / 30.00%)


#### external (3x / 100.00%)

<details>
<summary> download failure in context (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-09-14 14:28:33 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18841/pull-kubevirt-e2e-kind-1.36-sig-compute-arm64/2099505469235138560#1:build-log.txt%3A254)
<details>
<summary>all...</summary>

* _2026-09-14 14:29:09 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18841/pull-kubevirt-e2e-kind-1.36-sev/2099505474956169216#1:build-log.txt%3A245)

* _2026-09-14 14:28:33 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18841/pull-kubevirt-e2e-kind-1.36-sig-compute-arm64/2099505469235138560#1:build-log.txt%3A254)

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

Last updated: 2026-09-19 12:14:40
