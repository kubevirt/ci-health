



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-06-05 (1x / 50.00%)


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

### 2026-06-01 (1x / 50.00%)


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

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (1x / 50.00%)

<details>
<summary> transient kube-apiserver body decode noise (1x / 50.00%) </summary>

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

### needs-investigation (1x / 50.00%)

<details>
<summary> no matching pattern (1x / 50.00%) </summary>

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


### main (1x / 50.00%)


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

### release-1.7 (1x / 50.00%)


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

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-network (1x / 50.00%)


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

### sig-storage (1x / 50.00%)


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

Last updated: 2026-06-08 16:33:15
