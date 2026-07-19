



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-07-18 (2x / 20.00%)


#### external (1x / 50.00%)

<details>
<summary> transient kube-apiserver body decode noise (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-07-18 12:17:09 &#43;0000 UTC_: <code>12:21:01: I0718 08:21:01.640723    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18492/pull-kubevirt-e2e-k8s-1.34-sig-network/2078453882106679296#1:build-log.txt%3A691)
<details>
<summary>all...</summary>

* _2026-07-18 12:17:09 &#43;0000 UTC_: <code>12:21:01: I0718 08:21:01.640723    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18492/pull-kubevirt-e2e-k8s-1.34-sig-network/2078453882106679296#1:build-log.txt%3A691)

</details>

<hr/>
</details>

#### needs-investigation (1x / 50.00%)

<details>
<summary> no error snippets (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-07-18 07:31:20 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18494/pull-kubevirt-e2e-k8s-1.34-sig-storage/2078381951114285056)
<details>
<summary>all...</summary>

* _2026-07-18 07:31:20 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18494/pull-kubevirt-e2e-k8s-1.34-sig-storage/2078381951114285056)

</details>

<hr/>
</details>

### 2026-07-17 (1x / 10.00%)


#### needs-investigation (1x / 100.00%)

<details>
<summary> no error snippets (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-07-17 00:02:39 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18455/pull-kubevirt-e2e-k8s-1.35-sig-operator/2077906655319691264)
<details>
<summary>all...</summary>

* _2026-07-17 00:02:39 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18455/pull-kubevirt-e2e-k8s-1.35-sig-operator/2077906655319691264)

</details>

<hr/>
</details>

### 2026-07-15 (1x / 10.00%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-07-15 02:53:43 &#43;0000 UTC_: <code>03:14:41: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18406/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2077224550935826432#1:build-log.txt%3A4402)
<details>
<summary>all...</summary>

* _2026-07-15 02:53:43 &#43;0000 UTC_: <code>03:14:41: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18406/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2077224550935826432#1:build-log.txt%3A4402)

</details>

<hr/>
</details>

### 2026-07-14 (3x / 30.00%)


#### internal (3x / 100.00%)

<details>
<summary> kind cluster creation failure (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-07-14 13:34:56 &#43;0000 UTC_: <code>13:42:32: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18418/pull-kubevirt-e2e-kind-1.34-sev/2077023792403582976#1:build-log.txt%3A680)
<details>
<summary>all...</summary>

* _2026-07-14 13:34:56 &#43;0000 UTC_: <code>13:42:32: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18418/pull-kubevirt-e2e-kind-1.34-sev/2077023792403582976#1:build-log.txt%3A680)

* _2026-07-14 11:09:50 &#43;0000 UTC_: <code>11:17:03: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17117/pull-kubevirt-e2e-kind-1.34-sev/2076987423987863552#1:build-log.txt%3A964)

</details>

<hr/>
</details>
<details>
<summary> make cluster lifecycle target failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-07-14 11:54:26 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18454/pull-kubevirt-e2e-kind-1.34-sev/2076998450771136512#1:build-log.txt%3A1242)
<details>
<summary>all...</summary>

* _2026-07-14 11:54:26 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18454/pull-kubevirt-e2e-kind-1.34-sev/2076998450771136512#1:build-log.txt%3A1242)

</details>

<hr/>
</details>

### 2026-07-13 (3x / 30.00%)


#### external (2x / 66.67%)

<details>
<summary> download failure in context (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-07-13 10:57:10 &#43;0000 UTC_: <code>11:07:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18351/pull-kubevirt-e2e-k8s-1.35-sig-compute/2076613329182265344#1:build-log.txt%3A4201)
<details>
<summary>all...</summary>

* _2026-07-13 10:57:10 &#43;0000 UTC_: <code>11:07:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18351/pull-kubevirt-e2e-k8s-1.35-sig-compute/2076613329182265344#1:build-log.txt%3A4201)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-07-13 13:28:41 &#43;0000 UTC_: <code>14:12:04: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18445/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076659968248909824#1:build-log.txt%3A5189)
<details>
<summary>all...</summary>

* _2026-07-13 13:28:41 &#43;0000 UTC_: <code>14:12:04: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18445/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076659968248909824#1:build-log.txt%3A5189)

</details>

<hr/>
</details>

#### internal (1x / 33.33%)

<details>
<summary> KubeVirt deployment timeout (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-07-13 13:29:22 &#43;0000 UTC_: <code>14:11:20: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18446/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076660067666497536#1:build-log.txt%3A5137)
<details>
<summary>all...</summary>

* _2026-07-13 13:29:22 &#43;0000 UTC_: <code>14:11:20: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18446/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076660067666497536#1:build-log.txt%3A5137)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### internal (4x / 40.00%)

<details>
<summary> kind cluster creation failure (2x / 20.00%) </summary>

<hr/>

**2x**: _2026-07-14 13:34:56 &#43;0000 UTC_: <code>13:42:32: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18418/pull-kubevirt-e2e-kind-1.34-sev/2077023792403582976#1:build-log.txt%3A680)
<details>
<summary>all...</summary>

* _2026-07-14 13:34:56 &#43;0000 UTC_: <code>13:42:32: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18418/pull-kubevirt-e2e-kind-1.34-sev/2077023792403582976#1:build-log.txt%3A680)
<details><summary>context</summary>
<pre>13:42:30:  ✓ Ensuring node image (kindest/node:v1.34.3) 🖼
13:42:30:  • Preparing nodes 📦 📦   ...
13:42:32:  ✗ Preparing nodes 📦 📦
13:42:32: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;
13:42:32:
13:42:32: Stack Trace:
13:42:32: sigs.k8s.io/kind/pkg/errors.Errorf</pre>
</details>


* _2026-07-14 11:09:50 &#43;0000 UTC_: <code>11:17:03: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17117/pull-kubevirt-e2e-kind-1.34-sev/2076987423987863552#1:build-log.txt%3A964)
<details><summary>context</summary>
<pre>11:17:01:  ✓ Ensuring node image (kindest/node:v1.34.3) 🖼
11:17:01:  • Preparing nodes 📦 📦   ...
11:17:03:  ✗ Preparing nodes 📦 📦
11:17:03: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;
11:17:03:
11:17:03: Stack Trace:
11:17:03: sigs.k8s.io/kind/pkg/errors.Errorf</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> KubeVirt deployment timeout (1x / 10.00%) </summary>

<hr/>

**1x**: _2026-07-13 13:29:22 &#43;0000 UTC_: <code>14:11:20: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18446/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076660067666497536#1:build-log.txt%3A5137)
<details>
<summary>all...</summary>

* _2026-07-13 13:29:22 &#43;0000 UTC_: <code>14:11:20: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18446/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076660067666497536#1:build-log.txt%3A5137)
<details><summary>context</summary>
<pre>14:05:20: &#43; sleep 1m
14:06:20: &#43; _kubectl wait -n kubevirt kv kubevirt --for condition=Available --timeout 5m
14:06:20: &#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.34/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.34/.kubeconfig wait -n kubevirt kv kubevirt --for condition=Available --timeout 5m
14:11:20: error: timed out waiting for the condition on kubevirts/kubevirt
14:11:20: &#43; (( count&#43;&#43; ))
14:11:20: &#43; (( count == 5 ))
14:11:20: &#43; echo &#39;KubeVirt not ready in time&#39;</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> make cluster lifecycle target failure (1x / 10.00%) </summary>

<hr/>

**1x**: _2026-07-14 11:54:26 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18454/pull-kubevirt-e2e-kind-1.34-sev/2076998450771136512#1:build-log.txt%3A1242)
<details>
<summary>all...</summary>

* _2026-07-14 11:54:26 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18454/pull-kubevirt-e2e-kind-1.34-sev/2076998450771136512#1:build-log.txt%3A1242)
<details><summary>context</summary>
<pre>12:02:30: Downloading cni-plugins-linux-amd64-v0.8.5.tgz
12:02:30: &#43;&#43; curl -sSL -o /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.34/cni-plugins-linux-amd64-v0.8.5.tgz https://github.com/containernetworking/plugins/releases/download/v0.8.5/cni-plugins-linux-amd64-v0.8.5.tgz
12:02:45: curl: (6) Could not resolve host: github.com
make: *** [Makefile:174: cluster-up] Error 6
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>

<hr/>
</details>

### external (4x / 40.00%)

<details>
<summary> download failure in context (2x / 20.00%) </summary>

<hr/>

**2x**: _2026-07-13 10:57:10 &#43;0000 UTC_: <code>11:07:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18351/pull-kubevirt-e2e-k8s-1.35-sig-compute/2076613329182265344#1:build-log.txt%3A4201)
<details>
<summary>all...</summary>

* _2026-07-15 02:53:43 &#43;0000 UTC_: <code>03:14:41: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18406/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2077224550935826432#1:build-log.txt%3A4402)
<details><summary>context</summary>
<pre>03:14:41: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted:
03:14:41: INFO: Elapsed time: 1.508s
03:14:41: INFO: 0 processes.
03:14:41: ERROR: Build failed. Not running target
03:14:44: &#43; rm -f /tmp/kubevirt.deploy.JG0I
make: *** [Makefile:189: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


* _2026-07-13 10:57:10 &#43;0000 UTC_: <code>11:07:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18351/pull-kubevirt-e2e-k8s-1.35-sig-compute/2076613329182265344#1:build-log.txt%3A4201)
<details><summary>context</summary>
<pre>11:07:18: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
11:07:18: INFO: Elapsed time: 2.099s
11:07:18: INFO: 0 processes.
11:07:18: ERROR: Build failed. Not running target
11:07:21: &#43; rm -f /tmp/kubevirt.deploy.thAt
make: *** [Makefile:189: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 10.00%) </summary>

<hr/>

**1x**: _2026-07-13 13:28:41 &#43;0000 UTC_: <code>14:12:04: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18445/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076659968248909824#1:build-log.txt%3A5189)
<details>
<summary>all...</summary>

* _2026-07-13 13:28:41 &#43;0000 UTC_: <code>14:12:04: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18445/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076659968248909824#1:build-log.txt%3A5189)
<details><summary>context</summary>
<pre>14:06:04: &#43; sleep 1m
14:07:04: &#43; _kubectl wait -n kubevirt kv kubevirt --for condition=Available --timeout 5m
14:07:04: &#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.34/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.34/.kubeconfig wait -n kubevirt kv kubevirt --for condition=Available --timeout 5m
14:12:04: error: timed out waiting for the condition on kubevirts/kubevirt
14:12:04: &#43; (( count&#43;&#43; ))
14:12:04: &#43; (( count == 5 ))
14:12:04: &#43; echo &#39;KubeVirt not ready in time&#39;</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 10.00%) </summary>

<hr/>

**1x**: _2026-07-18 12:17:09 &#43;0000 UTC_: <code>12:21:01: I0718 08:21:01.640723    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18492/pull-kubevirt-e2e-k8s-1.34-sig-network/2078453882106679296#1:build-log.txt%3A691)
<details>
<summary>all...</summary>

* _2026-07-18 12:17:09 &#43;0000 UTC_: <code>12:21:01: I0718 08:21:01.640723    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18492/pull-kubevirt-e2e-k8s-1.34-sig-network/2078453882106679296#1:build-log.txt%3A691)
<details><summary>context</summary>
<pre>12:21:00: [control-plane-check] kube-controller-manager is healthy after 3.792332ms
12:21:00: [control-plane-check] kube-scheduler is healthy after 4.318949ms
12:21:01: I0718 08:21:01.142611    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:21:01: I0718 08:21:01.640723    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:21:02: [control-plane-check] kube-apiserver is healthy after 2.00700825s
12:21:02: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace
12:21:02: I0718 08:21:02.146240    1605 uploadconfig.go:111] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap</pre>
</details>


</details>

<hr/>
</details>

### needs-investigation (2x / 20.00%)

<details>
<summary> no error snippets (2x / 20.00%) </summary>

<hr/>

**1x**: _2026-07-18 07:31:20 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18494/pull-kubevirt-e2e-k8s-1.34-sig-storage/2078381951114285056)
<details>
<summary>all...</summary>

* _2026-07-18 07:31:20 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18494/pull-kubevirt-e2e-k8s-1.34-sig-storage/2078381951114285056)

</details>

<hr/>

**1x**: _2026-07-17 00:02:39 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18455/pull-kubevirt-e2e-k8s-1.35-sig-operator/2077906655319691264)
<details>
<summary>all...</summary>

* _2026-07-17 00:02:39 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18455/pull-kubevirt-e2e-k8s-1.35-sig-operator/2077906655319691264)

</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (7x / 70.00%)


#### external (2x / 28.57%)

<details>
<summary> download failure in context (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-07-13 10:57:10 &#43;0000 UTC_: <code>11:07:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18351/pull-kubevirt-e2e-k8s-1.35-sig-compute/2076613329182265344#1:build-log.txt%3A4201)
<details>
<summary>all...</summary>

* _2026-07-13 10:57:10 &#43;0000 UTC_: <code>11:07:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18351/pull-kubevirt-e2e-k8s-1.35-sig-compute/2076613329182265344#1:build-log.txt%3A4201)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-07-18 12:17:09 &#43;0000 UTC_: <code>12:21:01: I0718 08:21:01.640723    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18492/pull-kubevirt-e2e-k8s-1.34-sig-network/2078453882106679296#1:build-log.txt%3A691)
<details>
<summary>all...</summary>

* _2026-07-18 12:17:09 &#43;0000 UTC_: <code>12:21:01: I0718 08:21:01.640723    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18492/pull-kubevirt-e2e-k8s-1.34-sig-network/2078453882106679296#1:build-log.txt%3A691)

</details>

<hr/>
</details>

#### internal (3x / 42.86%)

<details>
<summary> kind cluster creation failure (2x / 28.57%) </summary>

<hr/>

**2x**: _2026-07-14 13:34:56 &#43;0000 UTC_: <code>13:42:32: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18418/pull-kubevirt-e2e-kind-1.34-sev/2077023792403582976#1:build-log.txt%3A680)
<details>
<summary>all...</summary>

* _2026-07-14 13:34:56 &#43;0000 UTC_: <code>13:42:32: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18418/pull-kubevirt-e2e-kind-1.34-sev/2077023792403582976#1:build-log.txt%3A680)

* _2026-07-14 11:09:50 &#43;0000 UTC_: <code>11:17:03: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17117/pull-kubevirt-e2e-kind-1.34-sev/2076987423987863552#1:build-log.txt%3A964)

</details>

<hr/>
</details>
<details>
<summary> make cluster lifecycle target failure (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-07-14 11:54:26 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18454/pull-kubevirt-e2e-kind-1.34-sev/2076998450771136512#1:build-log.txt%3A1242)
<details>
<summary>all...</summary>

* _2026-07-14 11:54:26 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18454/pull-kubevirt-e2e-kind-1.34-sev/2076998450771136512#1:build-log.txt%3A1242)

</details>

<hr/>
</details>

#### needs-investigation (2x / 28.57%)

<details>
<summary> no error snippets (2x / 28.57%) </summary>

<hr/>

**1x**: _2026-07-18 07:31:20 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18494/pull-kubevirt-e2e-k8s-1.34-sig-storage/2078381951114285056)
<details>
<summary>all...</summary>

* _2026-07-18 07:31:20 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18494/pull-kubevirt-e2e-k8s-1.34-sig-storage/2078381951114285056)

</details>

<hr/>

**1x**: _2026-07-17 00:02:39 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18455/pull-kubevirt-e2e-k8s-1.35-sig-operator/2077906655319691264)
<details>
<summary>all...</summary>

* _2026-07-17 00:02:39 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18455/pull-kubevirt-e2e-k8s-1.35-sig-operator/2077906655319691264)

</details>

<hr/>
</details>

### release-1.9 (3x / 30.00%)


#### external (2x / 66.67%)

<details>
<summary> download failure in context (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-07-15 02:53:43 &#43;0000 UTC_: <code>03:14:41: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18406/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2077224550935826432#1:build-log.txt%3A4402)
<details>
<summary>all...</summary>

* _2026-07-15 02:53:43 &#43;0000 UTC_: <code>03:14:41: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18406/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2077224550935826432#1:build-log.txt%3A4402)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-07-13 13:28:41 &#43;0000 UTC_: <code>14:12:04: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18445/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076659968248909824#1:build-log.txt%3A5189)
<details>
<summary>all...</summary>

* _2026-07-13 13:28:41 &#43;0000 UTC_: <code>14:12:04: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18445/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076659968248909824#1:build-log.txt%3A5189)

</details>

<hr/>
</details>

#### internal (1x / 33.33%)

<details>
<summary> KubeVirt deployment timeout (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-07-13 13:29:22 &#43;0000 UTC_: <code>14:11:20: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18446/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076660067666497536#1:build-log.txt%3A5137)
<details>
<summary>all...</summary>

* _2026-07-13 13:29:22 &#43;0000 UTC_: <code>14:11:20: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18446/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076660067666497536#1:build-log.txt%3A5137)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-network (2x / 20.00%)


#### external (2x / 100.00%)

<details>
<summary> download failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-07-15 02:53:43 &#43;0000 UTC_: <code>03:14:41: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18406/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2077224550935826432#1:build-log.txt%3A4402)
<details>
<summary>all...</summary>

* _2026-07-15 02:53:43 &#43;0000 UTC_: <code>03:14:41: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18406/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2077224550935826432#1:build-log.txt%3A4402)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-07-18 12:17:09 &#43;0000 UTC_: <code>12:21:01: I0718 08:21:01.640723    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18492/pull-kubevirt-e2e-k8s-1.34-sig-network/2078453882106679296#1:build-log.txt%3A691)
<details>
<summary>all...</summary>

* _2026-07-18 12:17:09 &#43;0000 UTC_: <code>12:21:01: I0718 08:21:01.640723    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18492/pull-kubevirt-e2e-k8s-1.34-sig-network/2078453882106679296#1:build-log.txt%3A691)

</details>

<hr/>
</details>

### sig-storage (1x / 10.00%)


#### needs-investigation (1x / 100.00%)

<details>
<summary> no error snippets (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-07-18 07:31:20 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18494/pull-kubevirt-e2e-k8s-1.34-sig-storage/2078381951114285056)
<details>
<summary>all...</summary>

* _2026-07-18 07:31:20 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18494/pull-kubevirt-e2e-k8s-1.34-sig-storage/2078381951114285056)

</details>

<hr/>
</details>

### sig-compute (7x / 70.00%)


#### internal (4x / 57.14%)

<details>
<summary> kind cluster creation failure (2x / 28.57%) </summary>

<hr/>

**2x**: _2026-07-14 13:34:56 &#43;0000 UTC_: <code>13:42:32: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18418/pull-kubevirt-e2e-kind-1.34-sev/2077023792403582976#1:build-log.txt%3A680)
<details>
<summary>all...</summary>

* _2026-07-14 13:34:56 &#43;0000 UTC_: <code>13:42:32: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18418/pull-kubevirt-e2e-kind-1.34-sev/2077023792403582976#1:build-log.txt%3A680)

* _2026-07-14 11:09:50 &#43;0000 UTC_: <code>11:17:03: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17117/pull-kubevirt-e2e-kind-1.34-sev/2076987423987863552#1:build-log.txt%3A964)

</details>

<hr/>
</details>
<details>
<summary> make cluster lifecycle target failure (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-07-14 11:54:26 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18454/pull-kubevirt-e2e-kind-1.34-sev/2076998450771136512#1:build-log.txt%3A1242)
<details>
<summary>all...</summary>

* _2026-07-14 11:54:26 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18454/pull-kubevirt-e2e-kind-1.34-sev/2076998450771136512#1:build-log.txt%3A1242)

</details>

<hr/>
</details>
<details>
<summary> KubeVirt deployment timeout (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-07-13 13:29:22 &#43;0000 UTC_: <code>14:11:20: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18446/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076660067666497536#1:build-log.txt%3A5137)
<details>
<summary>all...</summary>

* _2026-07-13 13:29:22 &#43;0000 UTC_: <code>14:11:20: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18446/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076660067666497536#1:build-log.txt%3A5137)

</details>

<hr/>
</details>

#### external (2x / 28.57%)

<details>
<summary> download failure in context (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-07-13 10:57:10 &#43;0000 UTC_: <code>11:07:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18351/pull-kubevirt-e2e-k8s-1.35-sig-compute/2076613329182265344#1:build-log.txt%3A4201)
<details>
<summary>all...</summary>

* _2026-07-13 10:57:10 &#43;0000 UTC_: <code>11:07:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18351/pull-kubevirt-e2e-k8s-1.35-sig-compute/2076613329182265344#1:build-log.txt%3A4201)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-07-13 13:28:41 &#43;0000 UTC_: <code>14:12:04: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18445/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076659968248909824#1:build-log.txt%3A5189)
<details>
<summary>all...</summary>

* _2026-07-13 13:28:41 &#43;0000 UTC_: <code>14:12:04: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18445/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076659968248909824#1:build-log.txt%3A5189)

</details>

<hr/>
</details>

#### needs-investigation (1x / 14.29%)

<details>
<summary> no error snippets (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-07-17 00:02:39 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18455/pull-kubevirt-e2e-k8s-1.35-sig-operator/2077906655319691264)
<details>
<summary>all...</summary>

* _2026-07-17 00:02:39 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18455/pull-kubevirt-e2e-k8s-1.35-sig-operator/2077906655319691264)

</details>

<hr/>
</details>

Last updated: 2026-07-19 18:24:34
