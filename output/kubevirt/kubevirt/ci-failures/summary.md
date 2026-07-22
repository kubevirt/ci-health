



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-07-21 (3x / 21.43%)


#### external (3x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (3x / 100.00%) </summary>

<hr/>

**3x**: _2026-07-21 10:31:31 &#43;0000 UTC_: <code>10:35:24: I0721 06:35:24.375362    1616 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.34-sig-network/2079503745925779456#1:build-log.txt%3A695)
<details>
<summary>all...</summary>

* _2026-07-21 10:34:47 &#43;0000 UTC_: <code>10:38:57: I0721 06:38:56.954264    1608 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.35-sig-network/2079503746533953536#1:build-log.txt%3A704)

* _2026-07-21 10:33:48 &#43;0000 UTC_: <code>10:38:56: I0721 06:38:56.963765    1637 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.34-sig-compute/2079503746223575040#1:build-log.txt%3A679)

* _2026-07-21 10:31:31 &#43;0000 UTC_: <code>10:35:24: I0721 06:35:24.375362    1616 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.34-sig-network/2079503745925779456#1:build-log.txt%3A695)

</details>

<hr/>
</details>

### 2026-07-20 (5x / 35.71%)


#### external (3x / 60.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 40.00%) </summary>

<hr/>

**2x**: _2026-07-20 07:54:08 &#43;0000 UTC_: <code>10:27:52:   {&#34;[namespace kubevirt-test-default3 name testvmi-2lwrh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 33f00b10-c4c9-4651-a61c-85b3e8669791]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-07-20T10:27:24.832263Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-07-20T10:27:24.877758Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.9/2079112332193042432#1:build-log.txt%3A9862)
<details>
<summary>all...</summary>

* _2026-07-20 07:54:08 &#43;0000 UTC_: <code>10:27:52:   {&#34;[namespace kubevirt-test-default3 name testvmi-2lwrh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 33f00b10-c4c9-4651-a61c-85b3e8669791]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-07-20T10:27:24.832263Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-07-20T10:27:24.877758Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.9/2079112332193042432#1:build-log.txt%3A9862)

* _2026-07-20 07:53:50 &#43;0000 UTC_: <code>09:40:52:   {&#34;[namespace kubevirt-test-default4 name testvmi-z9dd6-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid f28e9809-0ef3-4a6d-a97c-2613435f1a6f]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-07-20T09:40:25.101823Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-07-20T09:40:25.119943Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.36-sig-storage-1.9/2079112335531708416#1:build-log.txt%3A5631)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-07-20 19:53:48 &#43;0000 UTC_: <code>19:57:33: I0720 15:57:33.518428    1640 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079250202144804864#1:build-log.txt%3A706)
<details>
<summary>all...</summary>

* _2026-07-20 19:53:48 &#43;0000 UTC_: <code>19:57:33: I0720 15:57:33.518428    1640 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079250202144804864#1:build-log.txt%3A706)

</details>

<hr/>
</details>

#### internal (2x / 40.00%)

<details>
<summary> make cluster lifecycle target failure (2x / 40.00%) </summary>

<hr/>

**2x**: _2026-07-20 13:32:28 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18520/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2079121446369497088#1:build-log.txt%3A3940)
<details>
<summary>all...</summary>

* _2026-07-20 13:32:28 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18520/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2079121446369497088#1:build-log.txt%3A3940)

* _2026-07-20 07:53:33 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.9/2079112330968305664#1:build-log.txt%3A1029)

</details>

<hr/>
</details>

### 2026-07-19 (1x / 7.14%)


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

### 2026-07-18 (2x / 14.29%)


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

### 2026-07-17 (1x / 7.14%)


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

### 2026-07-16 (1x / 7.14%)


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

### 2026-07-15 (1x / 7.14%)


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

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (10x / 71.43%)

<details>
<summary> transient kube-apiserver body decode noise (6x / 42.86%) </summary>

<hr/>

**6x**: _2026-07-20 19:53:48 &#43;0000 UTC_: <code>19:57:33: I0720 15:57:33.518428    1640 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079250202144804864#1:build-log.txt%3A706)
<details>
<summary>all...</summary>

* _2026-07-21 10:34:47 &#43;0000 UTC_: <code>10:38:57: I0721 06:38:56.954264    1608 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.35-sig-network/2079503746533953536#1:build-log.txt%3A704)
<details><summary>context</summary>
<pre>10:38:56: I0721 06:38:55.956720    1608 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
10:38:56: I0721 06:38:56.454483    1608 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
10:38:56: I0721 06:38:56.454534    1608 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
10:38:57: I0721 06:38:56.954264    1608 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
10:38:57: I0721 06:38:56.954302    1608 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
10:38:57: [control-plane-check] kube-apiserver is healthy after 2.003071497s
10:38:57: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>


* _2026-07-21 10:33:48 &#43;0000 UTC_: <code>10:38:56: I0721 06:38:56.963765    1637 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.34-sig-compute/2079503746223575040#1:build-log.txt%3A679)
<details><summary>context</summary>
<pre>10:38:55: [control-plane-check] kube-scheduler is healthy after 7.131635ms
10:38:55: [control-plane-check] kube-controller-manager is healthy after 12.95566ms
10:38:56: I0721 06:38:56.465183    1637 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
10:38:56: I0721 06:38:56.963765    1637 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
10:38:57: [control-plane-check] kube-apiserver is healthy after 2.006058787s
10:38:57: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace
10:38:57: I0721 06:38:57.467718    1637 uploadconfig.go:111] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap</pre>
</details>


* _2026-07-21 10:31:31 &#43;0000 UTC_: <code>10:35:24: I0721 06:35:24.375362    1616 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.34-sig-network/2079503745925779456#1:build-log.txt%3A695)
<details><summary>context</summary>
<pre>10:35:23: [control-plane-check] kube-scheduler is healthy after 8.263586ms
10:35:23: [control-plane-check] kube-controller-manager is healthy after 9.248658ms
10:35:23: I0721 06:35:23.876680    1616 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
10:35:24: I0721 06:35:24.375362    1616 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
10:35:24: [control-plane-check] kube-apiserver is healthy after 2.004972692s
10:35:24: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace
10:35:24: I0721 06:35:24.879094    1616 uploadconfig.go:111] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap</pre>
</details>


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
<summary> download failure in context (2x / 14.29%) </summary>

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
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 14.29%) </summary>

<hr/>

**2x**: _2026-07-20 07:54:08 &#43;0000 UTC_: <code>10:27:52:   {&#34;[namespace kubevirt-test-default3 name testvmi-2lwrh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 33f00b10-c4c9-4651-a61c-85b3e8669791]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-07-20T10:27:24.832263Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-07-20T10:27:24.877758Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.9/2079112332193042432#1:build-log.txt%3A9862)
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

### internal (2x / 14.29%)

<details>
<summary> make cluster lifecycle target failure (2x / 14.29%) </summary>

<hr/>

**2x**: _2026-07-20 13:32:28 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18520/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2079121446369497088#1:build-log.txt%3A3940)
<details>
<summary>all...</summary>

* _2026-07-20 13:32:28 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18520/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2079121446369497088#1:build-log.txt%3A3940)
<details><summary>context</summary>
<pre>13:51:33: 2026/07/20 13:51:33 existing blob: sha256:05daf2a1d5cd22bf20bf774e8599d99ee37f8edccf2588aaa86f21aaec88c88d
13:51:36: 2026/07/20 13:51:36 localhost:37295/kubevirt/kv-network-passt-binding-cni:devel: digest: sha256:0d6e045da432423816b2bc90af03985fc7c0432ef9ca321aa853181e18ba8a66 size: 701
13:51:48: &#43; rm -f /tmp/kubevirt.deploy.tB7r
make: *** [Makefile:174: cluster-sync] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>


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


</details>

<hr/>
</details>

### needs-investigation (2x / 14.29%)

<details>
<summary> no error snippets (2x / 14.29%) </summary>

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


### main (9x / 64.29%)


#### external (7x / 77.78%)

<details>
<summary> transient kube-apiserver body decode noise (6x / 66.67%) </summary>

<hr/>

**6x**: _2026-07-20 19:53:48 &#43;0000 UTC_: <code>19:57:33: I0720 15:57:33.518428    1640 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079250202144804864#1:build-log.txt%3A706)
<details>
<summary>all...</summary>

* _2026-07-21 10:34:47 &#43;0000 UTC_: <code>10:38:57: I0721 06:38:56.954264    1608 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.35-sig-network/2079503746533953536#1:build-log.txt%3A704)

* _2026-07-21 10:33:48 &#43;0000 UTC_: <code>10:38:56: I0721 06:38:56.963765    1637 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.34-sig-compute/2079503746223575040#1:build-log.txt%3A679)

* _2026-07-21 10:31:31 &#43;0000 UTC_: <code>10:35:24: I0721 06:35:24.375362    1616 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.34-sig-network/2079503745925779456#1:build-log.txt%3A695)

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

### release-1.7 (1x / 7.14%)


#### internal (1x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-07-20 13:32:28 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18520/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2079121446369497088#1:build-log.txt%3A3940)
<details>
<summary>all...</summary>

* _2026-07-20 13:32:28 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18520/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2079121446369497088#1:build-log.txt%3A3940)

</details>

<hr/>
</details>

### release-1.9 (4x / 28.57%)


#### external (3x / 75.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 50.00%) </summary>

<hr/>

**2x**: _2026-07-20 07:54:08 &#43;0000 UTC_: <code>10:27:52:   {&#34;[namespace kubevirt-test-default3 name testvmi-2lwrh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 33f00b10-c4c9-4651-a61c-85b3e8669791]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-07-20T10:27:24.832263Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-07-20T10:27:24.877758Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.9/2079112332193042432#1:build-log.txt%3A9862)
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


### sig-network (5x / 35.71%)


#### external (4x / 80.00%)

<details>
<summary> transient kube-apiserver body decode noise (3x / 60.00%) </summary>

<hr/>

**3x**: _2026-07-21 10:31:31 &#43;0000 UTC_: <code>10:35:24: I0721 06:35:24.375362    1616 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.34-sig-network/2079503745925779456#1:build-log.txt%3A695)
<details>
<summary>all...</summary>

* _2026-07-21 10:34:47 &#43;0000 UTC_: <code>10:38:57: I0721 06:38:56.954264    1608 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.35-sig-network/2079503746533953536#1:build-log.txt%3A704)

* _2026-07-21 10:31:31 &#43;0000 UTC_: <code>10:35:24: I0721 06:35:24.375362    1616 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.34-sig-network/2079503745925779456#1:build-log.txt%3A695)

* _2026-07-18 12:17:09 &#43;0000 UTC_: <code>12:21:01: I0718 08:21:01.640723    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18492/pull-kubevirt-e2e-k8s-1.34-sig-network/2078453882106679296#1:build-log.txt%3A691)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-07-15 02:53:43 &#43;0000 UTC_: <code>03:14:41: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18406/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2077224550935826432#1:build-log.txt%3A4402)
<details>
<summary>all...</summary>

* _2026-07-15 02:53:43 &#43;0000 UTC_: <code>03:14:41: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18406/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network-1.9/2077224550935826432#1:build-log.txt%3A4402)

</details>

<hr/>
</details>

#### internal (1x / 20.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-07-20 13:32:28 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18520/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2079121446369497088#1:build-log.txt%3A3940)
<details>
<summary>all...</summary>

* _2026-07-20 13:32:28 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18520/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2079121446369497088#1:build-log.txt%3A3940)

</details>

<hr/>
</details>

### sig-compute (5x / 35.71%)


#### external (3x / 60.00%)

<details>
<summary> transient kube-apiserver body decode noise (2x / 40.00%) </summary>

<hr/>

**2x**: _2026-07-20 19:53:48 &#43;0000 UTC_: <code>19:57:33: I0720 15:57:33.518428    1640 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079250202144804864#1:build-log.txt%3A706)
<details>
<summary>all...</summary>

* _2026-07-21 10:33:48 &#43;0000 UTC_: <code>10:38:56: I0721 06:38:56.963765    1637 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17995/pull-kubevirt-e2e-k8s-1.34-sig-compute/2079503746223575040#1:build-log.txt%3A679)

* _2026-07-20 19:53:48 &#43;0000 UTC_: <code>19:57:33: I0720 15:57:33.518428    1640 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2079250202144804864#1:build-log.txt%3A706)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-07-16 07:36:08 &#43;0000 UTC_: <code>07:46:34: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2077658380150771712#1:build-log.txt%3A4236)
<details>
<summary>all...</summary>

* _2026-07-16 07:36:08 &#43;0000 UTC_: <code>07:46:34: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18477/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2077658380150771712#1:build-log.txt%3A4236)

</details>

<hr/>
</details>

#### internal (1x / 20.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-07-20 07:53:33 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.9/2079112330968305664#1:build-log.txt%3A1029)
<details>
<summary>all...</summary>

* _2026-07-20 07:53:33 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.9/2079112330968305664#1:build-log.txt%3A1029)

</details>

<hr/>
</details>

#### needs-investigation (1x / 20.00%)

<details>
<summary> no error snippets (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-07-17 00:02:39 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18455/pull-kubevirt-e2e-k8s-1.35-sig-operator/2077906655319691264)
<details>
<summary>all...</summary>

* _2026-07-17 00:02:39 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18455/pull-kubevirt-e2e-k8s-1.35-sig-operator/2077906655319691264)

</details>

<hr/>
</details>

### sig-storage (4x / 28.57%)


#### external (3x / 75.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 50.00%) </summary>

<hr/>

**2x**: _2026-07-20 07:54:08 &#43;0000 UTC_: <code>10:27:52:   {&#34;[namespace kubevirt-test-default3 name testvmi-2lwrh-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid 33f00b10-c4c9-4651-a61c-85b3e8669791]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-07-20T10:27:24.832263Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-07-20T10:27:24.877758Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18509/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.9/2079112332193042432#1:build-log.txt%3A9862)
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

Last updated: 2026-07-22 00:27:27
