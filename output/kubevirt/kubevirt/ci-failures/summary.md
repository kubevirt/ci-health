



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-07-20 (4x / 30.77%)


#### external (3x / 75.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 50.00%) </summary>

<hr/>

**2x**: _2026-07-20 07:53:50 &#43;0000 UTC_: <code>09:40:52:   {&#34;[namespace kubevirt-test-default4 name testvmi-z9dd6-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid f28e9809-0ef3-4a6d-a97c-2613435f1a6f]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-07-20T09:40:25.101823Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-07-20T09:40:25.119943Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.36-sig-storage-1.9/2079112335531708416#1:build-log.txt%3A5631)
<details>
<summary>all...</summary>

* _2026-07-20 07:54:08 &#43;0000 UTC_: <code>10:27:52:   {&#34;[namespace kubevirt-test-default3 name testvmi-2lwrh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 33f00b10-c4c9-4651-a61c-85b3e8669791]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-07-20T10:27:24.832263Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-07-20T10:27:24.877758Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.9/2079112332193042432#1:build-log.txt%3A9862)

* _2026-07-20 07:53:50 &#43;0000 UTC_: <code>09:40:52:   {&#34;[namespace kubevirt-test-default4 name testvmi-z9dd6-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid f28e9809-0ef3-4a6d-a97c-2613435f1a6f]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-07-20T09:40:25.101823Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-07-20T09:40:25.119943Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.36-sig-storage-1.9/2079112335531708416#1:build-log.txt%3A5631)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-07-20 19:53:48 &#43;0000 UTC_: <code>19:57:33: I0720 15:57:33.518428    1640 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079250202144804864#1:build-log.txt%3A706)
<details>
<summary>all...</summary>

* _2026-07-20 19:53:48 &#43;0000 UTC_: <code>19:57:33: I0720 15:57:33.518428    1640 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079250202144804864#1:build-log.txt%3A706)

</details>

<hr/>
</details>

#### internal (1x / 25.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-07-20 07:53:33 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.9/2079112330968305664#1:build-log.txt%3A1029)
<details>
<summary>all...</summary>

* _2026-07-20 07:53:33 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.9/2079112330968305664#1:build-log.txt%3A1029)

</details>

<hr/>
</details>

### 2026-07-19 (1x / 7.69%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-07-19 13:53:33 &#43;0000 UTC_: <code>13:57:42: I0719 09:57:42.070577    1613 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18374/pull-kubevirt-e2e-k8s-1.35-sig-storage/2078840537326030848#1:build-log.txt%3A694)
<details>
<summary>all...</summary>

* _2026-07-19 13:53:33 &#43;0000 UTC_: <code>13:57:42: I0719 09:57:42.070577    1613 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18374/pull-kubevirt-e2e-k8s-1.35-sig-storage/2078840537326030848#1:build-log.txt%3A694)

</details>

<hr/>
</details>

### 2026-07-18 (2x / 15.38%)


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

### 2026-07-17 (1x / 7.69%)


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

### 2026-07-16 (1x / 7.69%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-07-16 07:36:08 &#43;0000 UTC_: <code>07:46:34: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2077658380150771712#1:build-log.txt%3A4236)
<details>
<summary>all...</summary>

* _2026-07-16 07:36:08 &#43;0000 UTC_: <code>07:46:34: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2077658380150771712#1:build-log.txt%3A4236)

</details>

<hr/>
</details>

### 2026-07-15 (1x / 7.69%)


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

### 2026-07-14 (3x / 23.08%)


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

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (7x / 53.85%)

<details>
<summary> transient kube-apiserver body decode noise (3x / 23.08%) </summary>

<hr/>

**3x**: _2026-07-20 19:53:48 &#43;0000 UTC_: <code>19:57:33: I0720 15:57:33.518428    1640 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079250202144804864#1:build-log.txt%3A706)
<details>
<summary>all...</summary>

* _2026-07-20 19:53:48 &#43;0000 UTC_: <code>19:57:33: I0720 15:57:33.518428    1640 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079250202144804864#1:build-log.txt%3A706)
<details><summary>context</summary>
<pre>19:57:32: [control-plane-check] kube-controller-manager is healthy after 2.7894ms
19:57:32: [control-plane-check] kube-scheduler is healthy after 3.205464ms
19:57:33: I0720 15:57:33.017598    1640 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
19:57:33: I0720 15:57:33.518428    1640 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
19:57:33: I0720 15:57:33.518474    1640 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
19:57:34: I0720 15:57:34.018096    1640 uploadconfig.go:109] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap
19:57:34: [control-plane-check] kube-apiserver is healthy after 1.502804162s</pre>
</details>


* _2026-07-19 13:53:33 &#43;0000 UTC_: <code>13:57:42: I0719 09:57:42.070577    1613 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18374/pull-kubevirt-e2e-k8s-1.35-sig-storage/2078840537326030848#1:build-log.txt%3A694)
<details><summary>context</summary>
<pre>13:57:41: [control-plane-check] kube-scheduler is healthy after 6.129231ms
13:57:41: [control-plane-check] kube-controller-manager is healthy after 6.239012ms
13:57:41: I0719 09:57:41.569690    1613 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
13:57:42: I0719 09:57:42.070577    1613 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
13:57:42: I0719 09:57:42.070623    1613 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
13:57:42: [control-plane-check] kube-apiserver is healthy after 1.503034392s
13:57:42: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>


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
<details>
<summary> download failure in context (2x / 15.38%) </summary>

<hr/>

**2x**: _2026-07-16 07:36:08 &#43;0000 UTC_: <code>07:46:34: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2077658380150771712#1:build-log.txt%3A4236)
<details>
<summary>all...</summary>

* _2026-07-16 07:36:08 &#43;0000 UTC_: <code>07:46:34: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2077658380150771712#1:build-log.txt%3A4236)
<details><summary>context</summary>
<pre>07:46:34: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted:
07:46:34: INFO: Elapsed time: 0.854s
07:46:34: INFO: 0 processes.
07:46:34: ERROR: Build failed. Not running target
07:46:36: &#43; rm -f /tmp/kubevirt.deploy.rldn
make: *** [Makefile:189: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


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


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 15.38%) </summary>

<hr/>

**2x**: _2026-07-20 07:53:50 &#43;0000 UTC_: <code>09:40:52:   {&#34;[namespace kubevirt-test-default4 name testvmi-z9dd6-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid f28e9809-0ef3-4a6d-a97c-2613435f1a6f]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-07-20T09:40:25.101823Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-07-20T09:40:25.119943Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.36-sig-storage-1.9/2079112335531708416#1:build-log.txt%3A5631)
<details>
<summary>all...</summary>

* _2026-07-20 07:54:08 &#43;0000 UTC_: <code>10:27:52:   {&#34;[namespace kubevirt-test-default3 name testvmi-2lwrh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 33f00b10-c4c9-4651-a61c-85b3e8669791]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-07-20T10:27:24.832263Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-07-20T10:27:24.877758Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.9/2079112332193042432#1:build-log.txt%3A9862)
<details><summary>context</summary>
<pre>10:27:52:
10:27:52:   Captured StdOut/StdErr Output &gt;&gt;
10:27:52:   {&#34;[namespace kubevirt-test-default3 name testvmi-2lwrh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 33f00b10-c4c9-4651-a61c-85b3e8669791]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;Event(v1.ObjectReference{Kind:\&#34;VirtualMachineInstance\&#34;, Namespace:\&#34;kubevirt-test-default3\&#34;, Name:\&#34;testvmi-2lwrh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\&#34;, UID:\&#34;33f00b10-c4c9-4651-a61c-85b3e8669791\&#34;, APIVersion:\&#34;kubevirt.io/v1\&#34;, ResourceVersion:\&#34;113368\&#34;, FieldPath:\&#34;\&#34;}): type: &#39;Normal&#39; reason: &#39;SuccessfulCreate&#39; Created virtual machine pod virt-launcher-testvmi-2lwrh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx5m8wj&#34;,&#34;pos&#34;:&#34;watcher.go:150&#34;,&#34;timestamp&#34;:&#34;2026-07-20T10:27:20.937664Z&#34;}
10:27:52:   {&#34;[namespace kubevirt-test-default3 name testvmi-2lwrh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 33f00b10-c4c9-4651-a61c-85b3e8669791]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-07-20T10:27:24.832263Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-07-20T10:27:24.877758Z&#34;}
10:27:52:   I0720 10:27:45.775444  513283 warnings.go:110] &#34;Warning: template.kubevirt.io/v1alpha1 VirtualMachineTemplateRequest is deprecated; use template.kubevirt.io/v1beta1 VirtualMachineTemplateRequest&#34;
10:27:52:   I0720 10:27:45.778482  513283 warnings.go:110] &#34;Warning: template.kubevirt.io/v1alpha1 VirtualMachineTemplateRequest is deprecated; use template.kubevirt.io/v1beta1 VirtualMachineTemplateRequest&#34;
10:27:52:   I0720 10:27:45.790538  513283 warnings.go:110] &#34;Warning: template.kubevirt.io/v1alpha1 VirtualMachineTemplate is deprecated; use template.kubevirt.io/v1beta1 VirtualMachineTemplate&#34;</pre>
</details>


* _2026-07-20 07:53:50 &#43;0000 UTC_: <code>09:40:52:   {&#34;[namespace kubevirt-test-default4 name testvmi-z9dd6-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid f28e9809-0ef3-4a6d-a97c-2613435f1a6f]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-07-20T09:40:25.101823Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-07-20T09:40:25.119943Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.36-sig-storage-1.9/2079112335531708416#1:build-log.txt%3A5631)
<details><summary>context</summary>
<pre>09:40:52:
09:40:52:   Captured StdOut/StdErr Output &gt;&gt;
09:40:52:   {&#34;[namespace kubevirt-test-default4 name testvmi-z9dd6-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid f28e9809-0ef3-4a6d-a97c-2613435f1a6f]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;Event(v1.ObjectReference{Kind:\&#34;VirtualMachineInstance\&#34;, Namespace:\&#34;kubevirt-test-default4\&#34;, Name:\&#34;testvmi-z9dd6-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\&#34;, UID:\&#34;f28e9809-0ef3-4a6d-a97c-2613435f1a6f\&#34;, APIVersion:\&#34;kubevirt.io/v1\&#34;, ResourceVersion:\&#34;15897\&#34;, FieldPath:\&#34;\&#34;}): type: &#39;Normal&#39; reason: &#39;SuccessfulCreate&#39; Created virtual machine pod virt-launcher-testvmi-z9dd6-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxswhnf&#34;,&#34;pos&#34;:&#34;watcher.go:150&#34;,&#34;timestamp&#34;:&#34;2026-07-20T09:40:21.657635Z&#34;}
09:40:52:   {&#34;[namespace kubevirt-test-default4 name testvmi-z9dd6-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid f28e9809-0ef3-4a6d-a97c-2613435f1a6f]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-07-20T09:40:25.101823Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-07-20T09:40:25.119943Z&#34;}
09:40:52:   I0720 09:40:45.850503  514475 warnings.go:110] &#34;Warning: template.kubevirt.io/v1alpha1 VirtualMachineTemplateRequest is deprecated; use template.kubevirt.io/v1beta1 VirtualMachineTemplateRequest&#34;
09:40:52:   I0720 09:40:45.852170  514475 warnings.go:110] &#34;Warning: template.kubevirt.io/v1alpha1 VirtualMachineTemplateRequest is deprecated; use template.kubevirt.io/v1beta1 VirtualMachineTemplateRequest&#34;
09:40:52:   I0720 09:40:45.858926  514475 warnings.go:110] &#34;Warning: template.kubevirt.io/v1alpha1 VirtualMachineTemplate is deprecated; use template.kubevirt.io/v1beta1 VirtualMachineTemplate&#34;</pre>
</details>


</details>

<hr/>
</details>

### internal (4x / 30.77%)

<details>
<summary> make cluster lifecycle target failure (2x / 15.38%) </summary>

<hr/>

**2x**: _2026-07-14 11:54:26 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18454/pull-kubevirt-e2e-kind-1.34-sev/2076998450771136512#1:build-log.txt%3A1242)
<details>
<summary>all...</summary>

* _2026-07-20 07:53:33 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.9/2079112330968305664#1:build-log.txt%3A1029)
<details><summary>context</summary>
<pre>08:26:05: rsync: [sender] failed to connect to 127.0.0.1 (127.0.0.1): Connection refused (111)
08:26:05: rsync error: error in socket IO (code 10) at clientserver.c(139) [sender=3.4.1]
08:26:08: &#43; rm -f /tmp/kubevirt.deploy.zpf9
make: *** [Makefile:189: cluster-sync] Error 10
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>


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
<details>
<summary> kind cluster creation failure (2x / 15.38%) </summary>

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

### needs-investigation (2x / 15.38%)

<details>
<summary> no error snippets (2x / 15.38%) </summary>

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


### main (9x / 69.23%)


#### external (4x / 44.44%)

<details>
<summary> transient kube-apiserver body decode noise (3x / 33.33%) </summary>

<hr/>

**3x**: _2026-07-20 19:53:48 &#43;0000 UTC_: <code>19:57:33: I0720 15:57:33.518428    1640 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079250202144804864#1:build-log.txt%3A706)
<details>
<summary>all...</summary>

* _2026-07-20 19:53:48 &#43;0000 UTC_: <code>19:57:33: I0720 15:57:33.518428    1640 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079250202144804864#1:build-log.txt%3A706)

* _2026-07-19 13:53:33 &#43;0000 UTC_: <code>13:57:42: I0719 09:57:42.070577    1613 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18374/pull-kubevirt-e2e-k8s-1.35-sig-storage/2078840537326030848#1:build-log.txt%3A694)

* _2026-07-18 12:17:09 &#43;0000 UTC_: <code>12:21:01: I0718 08:21:01.640723    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18492/pull-kubevirt-e2e-k8s-1.34-sig-network/2078453882106679296#1:build-log.txt%3A691)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-07-16 07:36:08 &#43;0000 UTC_: <code>07:46:34: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2077658380150771712#1:build-log.txt%3A4236)
<details>
<summary>all...</summary>

* _2026-07-16 07:36:08 &#43;0000 UTC_: <code>07:46:34: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2077658380150771712#1:build-log.txt%3A4236)

</details>

<hr/>
</details>

#### internal (3x / 33.33%)

<details>
<summary> kind cluster creation failure (2x / 22.22%) </summary>

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
<summary> make cluster lifecycle target failure (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-07-14 11:54:26 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18454/pull-kubevirt-e2e-kind-1.34-sev/2076998450771136512#1:build-log.txt%3A1242)
<details>
<summary>all...</summary>

* _2026-07-14 11:54:26 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18454/pull-kubevirt-e2e-kind-1.34-sev/2076998450771136512#1:build-log.txt%3A1242)

</details>

<hr/>
</details>

#### needs-investigation (2x / 22.22%)

<details>
<summary> no error snippets (2x / 22.22%) </summary>

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

### release-1.9 (4x / 30.77%)


#### external (3x / 75.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 50.00%) </summary>

<hr/>

**2x**: _2026-07-20 07:53:50 &#43;0000 UTC_: <code>09:40:52:   {&#34;[namespace kubevirt-test-default4 name testvmi-z9dd6-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid f28e9809-0ef3-4a6d-a97c-2613435f1a6f]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-07-20T09:40:25.101823Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-07-20T09:40:25.119943Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.36-sig-storage-1.9/2079112335531708416#1:build-log.txt%3A5631)
<details>
<summary>all...</summary>

* _2026-07-20 07:54:08 &#43;0000 UTC_: <code>10:27:52:   {&#34;[namespace kubevirt-test-default3 name testvmi-2lwrh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 33f00b10-c4c9-4651-a61c-85b3e8669791]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-07-20T10:27:24.832263Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-07-20T10:27:24.877758Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.9/2079112332193042432#1:build-log.txt%3A9862)

* _2026-07-20 07:53:50 &#43;0000 UTC_: <code>09:40:52:   {&#34;[namespace kubevirt-test-default4 name testvmi-z9dd6-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid f28e9809-0ef3-4a6d-a97c-2613435f1a6f]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-07-20T09:40:25.101823Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-07-20T09:40:25.119943Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.36-sig-storage-1.9/2079112335531708416#1:build-log.txt%3A5631)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-07-15 02:53:43 &#43;0000 UTC_: <code>03:14:41: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18406/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2077224550935826432#1:build-log.txt%3A4402)
<details>
<summary>all...</summary>

* _2026-07-15 02:53:43 &#43;0000 UTC_: <code>03:14:41: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18406/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2077224550935826432#1:build-log.txt%3A4402)

</details>

<hr/>
</details>

#### internal (1x / 25.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-07-20 07:53:33 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.9/2079112330968305664#1:build-log.txt%3A1029)
<details>
<summary>all...</summary>

* _2026-07-20 07:53:33 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.9/2079112330968305664#1:build-log.txt%3A1029)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-storage (4x / 30.77%)


#### external (3x / 75.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 50.00%) </summary>

<hr/>

**2x**: _2026-07-20 07:53:50 &#43;0000 UTC_: <code>09:40:52:   {&#34;[namespace kubevirt-test-default4 name testvmi-z9dd6-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid f28e9809-0ef3-4a6d-a97c-2613435f1a6f]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-07-20T09:40:25.101823Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-07-20T09:40:25.119943Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.36-sig-storage-1.9/2079112335531708416#1:build-log.txt%3A5631)
<details>
<summary>all...</summary>

* _2026-07-20 07:54:08 &#43;0000 UTC_: <code>10:27:52:   {&#34;[namespace kubevirt-test-default3 name testvmi-2lwrh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 33f00b10-c4c9-4651-a61c-85b3e8669791]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-07-20T10:27:24.832263Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-07-20T10:27:24.877758Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.9/2079112332193042432#1:build-log.txt%3A9862)

* _2026-07-20 07:53:50 &#43;0000 UTC_: <code>09:40:52:   {&#34;[namespace kubevirt-test-default4 name testvmi-z9dd6-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid f28e9809-0ef3-4a6d-a97c-2613435f1a6f]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-07-20T09:40:25.101823Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-07-20T09:40:25.119943Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.36-sig-storage-1.9/2079112335531708416#1:build-log.txt%3A5631)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-07-19 13:53:33 &#43;0000 UTC_: <code>13:57:42: I0719 09:57:42.070577    1613 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18374/pull-kubevirt-e2e-k8s-1.35-sig-storage/2078840537326030848#1:build-log.txt%3A694)
<details>
<summary>all...</summary>

* _2026-07-19 13:53:33 &#43;0000 UTC_: <code>13:57:42: I0719 09:57:42.070577    1613 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18374/pull-kubevirt-e2e-k8s-1.35-sig-storage/2078840537326030848#1:build-log.txt%3A694)

</details>

<hr/>
</details>

#### needs-investigation (1x / 25.00%)

<details>
<summary> no error snippets (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-07-18 07:31:20 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18494/pull-kubevirt-e2e-k8s-1.34-sig-storage/2078381951114285056)
<details>
<summary>all...</summary>

* _2026-07-18 07:31:20 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18494/pull-kubevirt-e2e-k8s-1.34-sig-storage/2078381951114285056)

</details>

<hr/>
</details>

### sig-network (2x / 15.38%)


#### external (2x / 100.00%)

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

### sig-compute (7x / 53.85%)


#### external (2x / 28.57%)

<details>
<summary> transient kube-apiserver body decode noise (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-07-20 19:53:48 &#43;0000 UTC_: <code>19:57:33: I0720 15:57:33.518428    1640 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079250202144804864#1:build-log.txt%3A706)
<details>
<summary>all...</summary>

* _2026-07-20 19:53:48 &#43;0000 UTC_: <code>19:57:33: I0720 15:57:33.518428    1640 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079250202144804864#1:build-log.txt%3A706)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-07-16 07:36:08 &#43;0000 UTC_: <code>07:46:34: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2077658380150771712#1:build-log.txt%3A4236)
<details>
<summary>all...</summary>

* _2026-07-16 07:36:08 &#43;0000 UTC_: <code>07:46:34: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2077658380150771712#1:build-log.txt%3A4236)

</details>

<hr/>
</details>

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
<summary> make cluster lifecycle target failure (2x / 28.57%) </summary>

<hr/>

**2x**: _2026-07-14 11:54:26 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18454/pull-kubevirt-e2e-kind-1.34-sev/2076998450771136512#1:build-log.txt%3A1242)
<details>
<summary>all...</summary>

* _2026-07-20 07:53:33 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.9/2079112330968305664#1:build-log.txt%3A1029)

* _2026-07-14 11:54:26 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 6</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18454/pull-kubevirt-e2e-kind-1.34-sev/2076998450771136512#1:build-log.txt%3A1242)

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

Last updated: 2026-07-21 10:01:39
