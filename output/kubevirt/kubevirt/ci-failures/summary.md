



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-06-05 (1x / 14.29%)


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

### 2026-06-01 (6x / 85.71%)


#### external (4x / 66.67%)

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

#### needs-investigation (1x / 16.67%)

<details>
<summary> no matching pattern (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-06-01 20:05:06 &#43;0000 UTC_: <code>22:09:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-2qbf8-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 90e398de-9d0a-43a6-831b-6bfceadaf099]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-01T22:08:53.733535Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18006/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2061539398683463680#1:build-log.txt%3A12941)
<details>
<summary>all...</summary>

* _2026-06-01 20:05:06 &#43;0000 UTC_: <code>22:09:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-2qbf8-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 90e398de-9d0a-43a6-831b-6bfceadaf099]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-01T22:08:53.733535Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18006/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2061539398683463680#1:build-log.txt%3A12941)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (5x / 71.43%)

<details>
<summary> download failure in context (2x / 28.57%) </summary>

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

**1x**: _2026-06-01 13:55:33 &#43;0000 UTC_: <code>14:07:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2061439596553572352#1:build-log.txt%3A4208)
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


</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (2x / 28.57%) </summary>

<hr/>

**2x**: _2026-06-01 14:05:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-compute/2061439596796841984#1:build-log.txt%3A239)
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


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 14.29%) </summary>

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

### internal (1x / 14.29%)

<details>
<summary> deployment timeout (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-06-01 14:01:39 &#43;0000 UTC_: <code>14:09:35: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-network/2061439596561960960#1:build-log.txt%3A1002)
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


</details>

<hr/>
</details>

### needs-investigation (1x / 14.29%)

<details>
<summary> no matching pattern (1x / 14.29%) </summary>

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


### release-1.7 (1x / 14.29%)


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

### main (6x / 85.71%)


#### external (5x / 83.33%)

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

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-network (3x / 42.86%)


#### external (2x / 66.67%)

<details>
<summary> download failure in context (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-06-01 13:55:33 &#43;0000 UTC_: <code>14:07:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2061439596553572352#1:build-log.txt%3A4208)
<details>
<summary>all...</summary>

* _2026-06-01 13:55:33 &#43;0000 UTC_: <code>14:07:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2061439596553572352#1:build-log.txt%3A4208)

</details>

<hr/>
</details>
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

#### internal (1x / 33.33%)

<details>
<summary> deployment timeout (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-06-01 14:01:39 &#43;0000 UTC_: <code>14:09:35: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-network/2061439596561960960#1:build-log.txt%3A1002)
<details>
<summary>all...</summary>

* _2026-06-01 14:01:39 &#43;0000 UTC_: <code>14:09:35: error: timed out waiting for the condition on deployments/cluster-network-addons-operator</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-network/2061439596561960960#1:build-log.txt%3A1002)

</details>

<hr/>
</details>

### sig-storage (2x / 28.57%)


#### needs-investigation (1x / 50.00%)

<details>
<summary> no matching pattern (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-06-01 20:05:06 &#43;0000 UTC_: <code>22:09:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-2qbf8-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 90e398de-9d0a-43a6-831b-6bfceadaf099]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-01T22:08:53.733535Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18006/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2061539398683463680#1:build-log.txt%3A12941)
<details>
<summary>all...</summary>

* _2026-06-01 20:05:06 &#43;0000 UTC_: <code>22:09:02:   {&#34;[namespace kubevirt-test-default3 name testvmi-2qbf8-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 90e398de-9d0a-43a6-831b-6bfceadaf099]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Failed to freeze VMI: server error. command Freeze failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: unable to execute QEMU agent command &#39;guest-fsfreeze-freeze&#39;: failed to open /mount_dir: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-01T22:08:53.733535Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18006/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2061539398683463680#1:build-log.txt%3A12941)

</details>

<hr/>
</details>

#### external (1x / 50.00%)

<details>
<summary> download failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-06-01 14:03:07 &#43;0000 UTC_: <code>14:05:31: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-storage/2061439596666818560#1:build-log.txt%3A590)
<details>
<summary>all...</summary>

* _2026-06-01 14:03:07 &#43;0000 UTC_: <code>14:05:31: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-storage/2061439596666818560#1:build-log.txt%3A590)

</details>

<hr/>
</details>

### sig-compute (2x / 28.57%)


#### external (2x / 100.00%)

<details>
<summary> container image pull failure in context (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-06-01 14:05:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-compute/2061439596796841984#1:build-log.txt%3A239)
<details>
<summary>all...</summary>

* _2026-06-01 14:05:38 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-operator/2061439596926865408#1:build-log.txt%3A301)

* _2026-06-01 14:05:30 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16932/pull-kubevirt-e2e-k8s-1.34-sig-compute/2061439596796841984#1:build-log.txt%3A239)

</details>

<hr/>
</details>

Last updated: 2026-06-08 13:38:53
