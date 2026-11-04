





<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-11-02 (2x / 40.00%)


#### internal (1x / 50.00%)

<details>
<summary> configured node not found during provisioning (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-11-02 12:10:13 &#43;0000 UTC_: <code>12:19:49: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A875)
<details>
<summary>all...</summary>

* _2026-11-02 12:10:13 &#43;0000 UTC_: <code>12:19:49: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A875)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A810)

<details><summary>context</summary>
<pre>12:19:44: [control-plane-check] kube-controller-manager is healthy after 8.488191ms
12:19:44: [control-plane-check] kube-scheduler is healthy after 10.387052ms
12:19:45: I1102 07:19:45.589042    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:45: I1102 07:19:46.088280    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:46: I1102 07:19:46.586947    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding</pre>
</details>



* <strong>diagnostic</strong>: <code>12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A812)

<details><summary>context</summary>
<pre>12:19:45: I1102 07:19:45.589042    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:45: I1102 07:19:46.088280    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:46: I1102 07:19:46.586947    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:47.090991    1901 uploadconfig.go:109] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap
12:19:46: [control-plane-check] kube-apiserver is healthy after 2.005710556s</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A1156)

<details><summary>context</summary>
<pre>12:20:04: 
12:20:04: &#43; wait 69
12:20:53: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A1184)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:20:53: selecting podman as container runtime
12:20:54: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-windows2016 is being used by the following container(s): 6a6a2973d141a1448dc31509ecfd824ede38d3435e163c24d8b9cba7915f160b: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

#### external (1x / 50.00%)

<details>
<summary> container image pull failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-11-02 09:09:03 &#43;0000 UTC_: <code>09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A433)
<details>
<summary>all...</summary>

* _2026-11-02 09:09:03 &#43;0000 UTC_: <code>09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A433)


<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A434)

<details><summary>context</summary>
<pre>09:16:19: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
09:16:19: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)
make: *** [Makefile:179: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>



</details>

<hr/>
</details>

### 2026-10-28 (3x / 60.00%)


#### internal (1x / 33.33%)

<details>
<summary> configured node not found during provisioning (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-10-28 16:02:01 &#43;0000 UTC_: <code>16:12:51: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A862)
<details>
<summary>all...</summary>

* _2026-10-28 16:02:01 &#43;0000 UTC_: <code>16:12:51: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A862)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>16:12:48: I1028 12:12:48.374608    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A796)

<details><summary>context</summary>
<pre>16:12:47: [control-plane-check] kube-scheduler is healthy after 6.01446ms
16:12:47: [control-plane-check] kube-controller-manager is healthy after 7.342355ms
16:12:47: I1028 12:12:47.876825    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
16:12:48: I1028 12:12:48.374608    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
16:12:48: I1028 12:12:48.374663    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
16:12:48: I1028 12:12:48.875519    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
16:12:48: I1028 12:12:48.875556    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding</pre>
</details>



* <strong>diagnostic</strong>: <code>16:12:48: I1028 12:12:48.875519    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A798)

<details><summary>context</summary>
<pre>16:12:47: I1028 12:12:47.876825    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
16:12:48: I1028 12:12:48.374608    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
16:12:48: I1028 12:12:48.374663    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
16:12:48: I1028 12:12:48.875519    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
16:12:48: I1028 12:12:48.875556    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
16:12:49: [control-plane-check] kube-apiserver is healthy after 2.003454173s
16:12:49: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1157)

<details><summary>context</summary>
<pre>16:14:19: 
16:14:19: &#43; wait 66
16:14:39: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1185)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
16:14:40: selecting podman as container runtime
16:14:40: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-storage is being used by the following container(s): ab5ac3b271a5edd5373cc1acd2a27a8884d4400a81f5eed93a634e9d5f377457: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

#### needs-investigation (2x / 66.67%)

<details>
<summary> image layer unavailable in original registries (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-10-28 15:32:58 &#43;0000 UTC_: <code>15:59:16: Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2209)
<details>
<summary>all...</summary>

* _2026-10-28 15:34:15 &#43;0000 UTC_: <code>15:40:36: Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A1339)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>15:38:17: I1028 11:38:17.589226    1900 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A822)

<details><summary>context</summary>
<pre>15:38:16: [control-plane-check] kube-scheduler is healthy after 2.953464ms
15:38:16: [control-plane-check] kube-controller-manager is healthy after 4.656113ms
15:38:17: I1028 11:38:17.089314    1900 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
15:38:17: I1028 11:38:17.589226    1900 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
15:38:17: I1028 11:38:17.589278    1900 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
15:38:18: [control-plane-check] kube-apiserver is healthy after 1.502106535s
15:38:18: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:194: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A1341)

<details><summary>context</summary>
<pre>15:40:34: INFO: Running command line: bazel-bin/push-fedora-with-test-tooling-container-disk.exe --registry localhost:45413/kubevirt --repository fedora-with-test-tooling-container-disk --tag devel
15:40:36: Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries
15:40:38: &#43; rm -f /tmp/kubevirt.deploy.VzYI
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A1350)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
15:40:38: selecting podman as container runtime
15:41:00: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-operator is being used by the following container(s): 072ddc794764862cead4ae6b6c5479ed96468412c055fb6419931fe946ae6cc7: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



* _2026-10-28 15:32:58 &#43;0000 UTC_: <code>15:59:16: Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2209)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>15:38:26: I1028 15:38:26.005087    2040 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A831)

<details><summary>context</summary>
<pre>15:38:25: [control-plane-check] kube-scheduler is healthy after 5.33362ms
15:38:25: [control-plane-check] kube-controller-manager is healthy after 7.324903ms
15:38:25: I1028 15:38:25.506495    2040 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: rbac: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
15:38:26: I1028 15:38:26.005087    2040 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
15:38:26: I1028 15:38:26.005128    2040 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
15:38:26: [control-plane-check] kube-apiserver is healthy after 1.501955359s
15:38:26: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:194: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2211)

<details><summary>context</summary>
<pre>15:59:14: INFO: Running command line: bazel-bin/push-fedora-with-test-tooling-container-disk.exe --registry localhost:45605/kubevirt --repository kv-fedora-with-test-tooling-container-disk --tag devel
15:59:16: Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries
15:59:23: &#43; rm -f /tmp/kubevirt.deploy.EZG2
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2220)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
15:59:23: selecting podman as container runtime
15:59:38: Error response from daemon: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2236)

<details><summary>context</summary>
<pre>bf3625c46ab363fc7049174e5d93200f7e4750ad44ba342936e9b2e64bc10de4
24814c8573d55105ee1958eebc65d6bbadec60390bf3034e207bf414b25d324a
065a95016f2c9afd305c46ed8fdf873e16712de179534c5dccf2dfb28fd45dd4
Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1020 is not a child of this shell
================================================================================</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2237)

<details><summary>context</summary>
<pre>24814c8573d55105ee1958eebc65d6bbadec60390bf3034e207bf414b25d324a
065a95016f2c9afd305c46ed8fdf873e16712de179534c5dccf2dfb28fd45dd4
Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1020 is not a child of this shell
================================================================================
Done cleaning up after podman in container.</pre>
</details>


</details>



</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### needs-investigation (2x / 40.00%)

<details>
<summary> image layer unavailable in original registries (2x / 40.00%) </summary>

<hr/>

**2x**: _2026-10-28 15:32:58 &#43;0000 UTC_: <code>15:59:16: Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2209)
<details>
<summary>all...</summary>

* _2026-10-28 15:34:15 &#43;0000 UTC_: <code>15:40:36: Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A1339)
<details><summary>context</summary>
<pre>15:40:34: INFO: 7 processes: 2 remote cache hit, 5 internal.
15:40:34: INFO: Build completed successfully, 7 total actions
15:40:34: INFO: Running command line: bazel-bin/push-fedora-with-test-tooling-container-disk.exe --registry localhost:45413/kubevirt --repository fedora-with-test-tooling-container-disk --tag devel
15:40:36: Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries
15:40:38: &#43; rm -f /tmp/kubevirt.deploy.VzYI
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>15:38:17: I1028 11:38:17.589226    1900 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A822)

<details><summary>context</summary>
<pre>15:38:16: [control-plane-check] kube-scheduler is healthy after 2.953464ms
15:38:16: [control-plane-check] kube-controller-manager is healthy after 4.656113ms
15:38:17: I1028 11:38:17.089314    1900 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
15:38:17: I1028 11:38:17.589226    1900 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
15:38:17: I1028 11:38:17.589278    1900 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
15:38:18: [control-plane-check] kube-apiserver is healthy after 1.502106535s
15:38:18: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:194: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A1341)

<details><summary>context</summary>
<pre>15:40:34: INFO: Running command line: bazel-bin/push-fedora-with-test-tooling-container-disk.exe --registry localhost:45413/kubevirt --repository fedora-with-test-tooling-container-disk --tag devel
15:40:36: Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries
15:40:38: &#43; rm -f /tmp/kubevirt.deploy.VzYI
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A1350)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
15:40:38: selecting podman as container runtime
15:41:00: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-operator is being used by the following container(s): 072ddc794764862cead4ae6b6c5479ed96468412c055fb6419931fe946ae6cc7: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



* _2026-10-28 15:32:58 &#43;0000 UTC_: <code>15:59:16: Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2209)
<details><summary>context</summary>
<pre>15:59:14: INFO: 1 process: 1 internal.
15:59:14: INFO: Build completed successfully, 1 total action
15:59:14: INFO: Running command line: bazel-bin/push-fedora-with-test-tooling-container-disk.exe --registry localhost:45605/kubevirt --repository kv-fedora-with-test-tooling-container-disk --tag devel
15:59:16: Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries
15:59:23: &#43; rm -f /tmp/kubevirt.deploy.EZG2
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>15:38:26: I1028 15:38:26.005087    2040 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A831)

<details><summary>context</summary>
<pre>15:38:25: [control-plane-check] kube-scheduler is healthy after 5.33362ms
15:38:25: [control-plane-check] kube-controller-manager is healthy after 7.324903ms
15:38:25: I1028 15:38:25.506495    2040 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: rbac: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
15:38:26: I1028 15:38:26.005087    2040 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
15:38:26: I1028 15:38:26.005128    2040 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
15:38:26: [control-plane-check] kube-apiserver is healthy after 1.501955359s
15:38:26: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:194: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2211)

<details><summary>context</summary>
<pre>15:59:14: INFO: Running command line: bazel-bin/push-fedora-with-test-tooling-container-disk.exe --registry localhost:45605/kubevirt --repository kv-fedora-with-test-tooling-container-disk --tag devel
15:59:16: Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries
15:59:23: &#43; rm -f /tmp/kubevirt.deploy.EZG2
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2220)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
15:59:23: selecting podman as container runtime
15:59:38: Error response from daemon: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2236)

<details><summary>context</summary>
<pre>bf3625c46ab363fc7049174e5d93200f7e4750ad44ba342936e9b2e64bc10de4
24814c8573d55105ee1958eebc65d6bbadec60390bf3034e207bf414b25d324a
065a95016f2c9afd305c46ed8fdf873e16712de179534c5dccf2dfb28fd45dd4
Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1020 is not a child of this shell
================================================================================</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2237)

<details><summary>context</summary>
<pre>24814c8573d55105ee1958eebc65d6bbadec60390bf3034e207bf414b25d324a
065a95016f2c9afd305c46ed8fdf873e16712de179534c5dccf2dfb28fd45dd4
Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1020 is not a child of this shell
================================================================================
Done cleaning up after podman in container.</pre>
</details>


</details>



</details>

<hr/>
</details>

### external (1x / 20.00%)

<details>
<summary> container image pull failure (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-11-02 09:09:03 &#43;0000 UTC_: <code>09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A433)
<details>
<summary>all...</summary>

* _2026-11-02 09:09:03 &#43;0000 UTC_: <code>09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A433)
<details><summary>context</summary>
<pre>09:16:19: Getting image source signatures
09:16:19: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
09:16:19: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)
make: *** [Makefile:179: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A434)

<details><summary>context</summary>
<pre>09:16:19: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
09:16:19: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)
make: *** [Makefile:179: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>



</details>

<hr/>
</details>

### internal (2x / 40.00%)

<details>
<summary> configured node not found during provisioning (2x / 40.00%) </summary>

<hr/>

**2x**: _2026-10-28 16:02:01 &#43;0000 UTC_: <code>16:12:51: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A862)
<details>
<summary>all...</summary>

* _2026-11-02 12:10:13 &#43;0000 UTC_: <code>12:19:49: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A875)
<details><summary>context</summary>
<pre>12:19:49: Warning: would violate PodSecurity &#34;restricted:latest&#34;: host namespaces (hostNetwork=true), allowPrivilegeEscalation != false (containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.allowPrivilegeEscalation=false), unrestricted capabilities (containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.capabilities.drop=[&#34;ALL&#34;]; container &#34;kube-flannel&#34; must not include &#34;NET_ADMIN&#34;, &#34;NET_RAW&#34; in securityContext.capabilities.add), restricted volume types (volumes &#34;run&#34;, &#34;cni-plugin&#34;, &#34;cni&#34;, &#34;xtables-lock&#34; use restricted volume type &#34;hostPath&#34;), runAsNonRoot != true (pod or containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.runAsNonRoot=true), seccompProfile (pod or containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.seccompProfile.type to &#34;RuntimeDefault&#34; or &#34;Localhost&#34;)
12:19:49: daemonset.apps/kube-flannel-ds created
12:19:49: time=&#34;2026-11-02T12:19:49Z&#34; level=info msg=&#34;[node 1]: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-&#34;
12:19:49: Error from server (NotFound): nodes &#34;node01&#34; not found
12:19:49:
12:19:49: ===== 21fb777d537fc8ba8050063cc17395841f375e1ae321c53c470f287880e4f2da ====
12:19:49: &#43; NUM_NODES=2</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A810)

<details><summary>context</summary>
<pre>12:19:44: [control-plane-check] kube-controller-manager is healthy after 8.488191ms
12:19:44: [control-plane-check] kube-scheduler is healthy after 10.387052ms
12:19:45: I1102 07:19:45.589042    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:45: I1102 07:19:46.088280    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:46: I1102 07:19:46.586947    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding</pre>
</details>



* <strong>diagnostic</strong>: <code>12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A812)

<details><summary>context</summary>
<pre>12:19:45: I1102 07:19:45.589042    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:45: I1102 07:19:46.088280    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:46: I1102 07:19:46.586947    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:47.090991    1901 uploadconfig.go:109] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap
12:19:46: [control-plane-check] kube-apiserver is healthy after 2.005710556s</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A1156)

<details><summary>context</summary>
<pre>12:20:04: 
12:20:04: &#43; wait 69
12:20:53: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A1184)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:20:53: selecting podman as container runtime
12:20:54: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-windows2016 is being used by the following container(s): 6a6a2973d141a1448dc31509ecfd824ede38d3435e163c24d8b9cba7915f160b: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



* _2026-10-28 16:02:01 &#43;0000 UTC_: <code>16:12:51: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A862)
<details><summary>context</summary>
<pre>16:12:51: Warning: would violate PodSecurity &#34;restricted:latest&#34;: host namespaces (hostNetwork=true), allowPrivilegeEscalation != false (containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.allowPrivilegeEscalation=false), unrestricted capabilities (containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.capabilities.drop=[&#34;ALL&#34;]; container &#34;kube-flannel&#34; must not include &#34;NET_ADMIN&#34;, &#34;NET_RAW&#34; in securityContext.capabilities.add), restricted volume types (volumes &#34;run&#34;, &#34;cni-plugin&#34;, &#34;cni&#34;, &#34;xtables-lock&#34; use restricted volume type &#34;hostPath&#34;), runAsNonRoot != true (pod or containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.runAsNonRoot=true), seccompProfile (pod or containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.seccompProfile.type to &#34;RuntimeDefault&#34; or &#34;Localhost&#34;)
16:12:51: daemonset.apps/kube-flannel-ds created
16:12:51: time=&#34;2026-10-28T16:12:51Z&#34; level=info msg=&#34;[node 1]: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-&#34;
16:12:51: Error from server (NotFound): nodes &#34;node01&#34; not found
16:12:52:
16:12:52: ===== 4d14d8585cba6436535b5549add1301313c33b5144286f93182a37c4e8d6be5a ====
16:12:52: &#43; NUM_NODES=2</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>16:12:48: I1028 12:12:48.374608    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A796)

<details><summary>context</summary>
<pre>16:12:47: [control-plane-check] kube-scheduler is healthy after 6.01446ms
16:12:47: [control-plane-check] kube-controller-manager is healthy after 7.342355ms
16:12:47: I1028 12:12:47.876825    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
16:12:48: I1028 12:12:48.374608    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
16:12:48: I1028 12:12:48.374663    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
16:12:48: I1028 12:12:48.875519    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
16:12:48: I1028 12:12:48.875556    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding</pre>
</details>



* <strong>diagnostic</strong>: <code>16:12:48: I1028 12:12:48.875519    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A798)

<details><summary>context</summary>
<pre>16:12:47: I1028 12:12:47.876825    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
16:12:48: I1028 12:12:48.374608    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
16:12:48: I1028 12:12:48.374663    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
16:12:48: I1028 12:12:48.875519    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
16:12:48: I1028 12:12:48.875556    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
16:12:49: [control-plane-check] kube-apiserver is healthy after 2.003454173s
16:12:49: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1157)

<details><summary>context</summary>
<pre>16:14:19: 
16:14:19: &#43; wait 66
16:14:39: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1185)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
16:14:40: selecting podman as container runtime
16:14:40: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-storage is being used by the following container(s): ab5ac3b271a5edd5373cc1acd2a27a8884d4400a81f5eed93a634e9d5f377457: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (5x / 100.00%)


#### needs-investigation (2x / 40.00%)

<details>
<summary> image layer unavailable in original registries (2x / 40.00%) </summary>

<hr/>

**2x**: _2026-10-28 15:32:58 &#43;0000 UTC_: <code>15:59:16: Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2209)
<details>
<summary>all...</summary>

* _2026-10-28 15:34:15 &#43;0000 UTC_: <code>15:40:36: Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A1339)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>15:38:17: I1028 11:38:17.589226    1900 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A822)

<details><summary>context</summary>
<pre>15:38:16: [control-plane-check] kube-scheduler is healthy after 2.953464ms
15:38:16: [control-plane-check] kube-controller-manager is healthy after 4.656113ms
15:38:17: I1028 11:38:17.089314    1900 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
15:38:17: I1028 11:38:17.589226    1900 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
15:38:17: I1028 11:38:17.589278    1900 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
15:38:18: [control-plane-check] kube-apiserver is healthy after 1.502106535s
15:38:18: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:194: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A1341)

<details><summary>context</summary>
<pre>15:40:34: INFO: Running command line: bazel-bin/push-fedora-with-test-tooling-container-disk.exe --registry localhost:45413/kubevirt --repository fedora-with-test-tooling-container-disk --tag devel
15:40:36: Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries
15:40:38: &#43; rm -f /tmp/kubevirt.deploy.VzYI
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A1350)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
15:40:38: selecting podman as container runtime
15:41:00: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-operator is being used by the following container(s): 072ddc794764862cead4ae6b6c5479ed96468412c055fb6419931fe946ae6cc7: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



* _2026-10-28 15:32:58 &#43;0000 UTC_: <code>15:59:16: Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2209)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>15:38:26: I1028 15:38:26.005087    2040 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A831)

<details><summary>context</summary>
<pre>15:38:25: [control-plane-check] kube-scheduler is healthy after 5.33362ms
15:38:25: [control-plane-check] kube-controller-manager is healthy after 7.324903ms
15:38:25: I1028 15:38:25.506495    2040 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: rbac: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
15:38:26: I1028 15:38:26.005087    2040 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
15:38:26: I1028 15:38:26.005128    2040 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
15:38:26: [control-plane-check] kube-apiserver is healthy after 1.501955359s
15:38:26: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:194: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2211)

<details><summary>context</summary>
<pre>15:59:14: INFO: Running command line: bazel-bin/push-fedora-with-test-tooling-container-disk.exe --registry localhost:45605/kubevirt --repository kv-fedora-with-test-tooling-container-disk --tag devel
15:59:16: Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries
15:59:23: &#43; rm -f /tmp/kubevirt.deploy.EZG2
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2220)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
15:59:23: selecting podman as container runtime
15:59:38: Error response from daemon: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2236)

<details><summary>context</summary>
<pre>bf3625c46ab363fc7049174e5d93200f7e4750ad44ba342936e9b2e64bc10de4
24814c8573d55105ee1958eebc65d6bbadec60390bf3034e207bf414b25d324a
065a95016f2c9afd305c46ed8fdf873e16712de179534c5dccf2dfb28fd45dd4
Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1020 is not a child of this shell
================================================================================</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2237)

<details><summary>context</summary>
<pre>24814c8573d55105ee1958eebc65d6bbadec60390bf3034e207bf414b25d324a
065a95016f2c9afd305c46ed8fdf873e16712de179534c5dccf2dfb28fd45dd4
Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1020 is not a child of this shell
================================================================================
Done cleaning up after podman in container.</pre>
</details>


</details>



</details>

<hr/>
</details>

#### external (1x / 20.00%)

<details>
<summary> container image pull failure (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-11-02 09:09:03 &#43;0000 UTC_: <code>09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A433)
<details>
<summary>all...</summary>

* _2026-11-02 09:09:03 &#43;0000 UTC_: <code>09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A433)


<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A434)

<details><summary>context</summary>
<pre>09:16:19: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
09:16:19: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)
make: *** [Makefile:179: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>



</details>

<hr/>
</details>

#### internal (2x / 40.00%)

<details>
<summary> configured node not found during provisioning (2x / 40.00%) </summary>

<hr/>

**2x**: _2026-10-28 16:02:01 &#43;0000 UTC_: <code>16:12:51: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A862)
<details>
<summary>all...</summary>

* _2026-11-02 12:10:13 &#43;0000 UTC_: <code>12:19:49: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A875)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A810)

<details><summary>context</summary>
<pre>12:19:44: [control-plane-check] kube-controller-manager is healthy after 8.488191ms
12:19:44: [control-plane-check] kube-scheduler is healthy after 10.387052ms
12:19:45: I1102 07:19:45.589042    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:45: I1102 07:19:46.088280    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:46: I1102 07:19:46.586947    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding</pre>
</details>



* <strong>diagnostic</strong>: <code>12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A812)

<details><summary>context</summary>
<pre>12:19:45: I1102 07:19:45.589042    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:45: I1102 07:19:46.088280    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:46: I1102 07:19:46.586947    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:47.090991    1901 uploadconfig.go:109] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap
12:19:46: [control-plane-check] kube-apiserver is healthy after 2.005710556s</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A1156)

<details><summary>context</summary>
<pre>12:20:04: 
12:20:04: &#43; wait 69
12:20:53: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A1184)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:20:53: selecting podman as container runtime
12:20:54: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-windows2016 is being used by the following container(s): 6a6a2973d141a1448dc31509ecfd824ede38d3435e163c24d8b9cba7915f160b: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



* _2026-10-28 16:02:01 &#43;0000 UTC_: <code>16:12:51: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A862)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>16:12:48: I1028 12:12:48.374608    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A796)

<details><summary>context</summary>
<pre>16:12:47: [control-plane-check] kube-scheduler is healthy after 6.01446ms
16:12:47: [control-plane-check] kube-controller-manager is healthy after 7.342355ms
16:12:47: I1028 12:12:47.876825    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
16:12:48: I1028 12:12:48.374608    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
16:12:48: I1028 12:12:48.374663    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
16:12:48: I1028 12:12:48.875519    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
16:12:48: I1028 12:12:48.875556    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding</pre>
</details>



* <strong>diagnostic</strong>: <code>16:12:48: I1028 12:12:48.875519    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A798)

<details><summary>context</summary>
<pre>16:12:47: I1028 12:12:47.876825    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
16:12:48: I1028 12:12:48.374608    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
16:12:48: I1028 12:12:48.374663    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
16:12:48: I1028 12:12:48.875519    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
16:12:48: I1028 12:12:48.875556    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
16:12:49: [control-plane-check] kube-apiserver is healthy after 2.003454173s
16:12:49: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1157)

<details><summary>context</summary>
<pre>16:14:19: 
16:14:19: &#43; wait 66
16:14:39: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1185)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
16:14:40: selecting podman as container runtime
16:14:40: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-storage is being used by the following container(s): ab5ac3b271a5edd5373cc1acd2a27a8884d4400a81f5eed93a634e9d5f377457: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (3x / 60.00%)


#### needs-investigation (1x / 33.33%)

<details>
<summary> image layer unavailable in original registries (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-10-28 15:34:15 &#43;0000 UTC_: <code>15:40:36: Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A1339)
<details>
<summary>all...</summary>

* _2026-10-28 15:34:15 &#43;0000 UTC_: <code>15:40:36: Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A1339)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>15:38:17: I1028 11:38:17.589226    1900 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A822)

<details><summary>context</summary>
<pre>15:38:16: [control-plane-check] kube-scheduler is healthy after 2.953464ms
15:38:16: [control-plane-check] kube-controller-manager is healthy after 4.656113ms
15:38:17: I1028 11:38:17.089314    1900 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
15:38:17: I1028 11:38:17.589226    1900 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
15:38:17: I1028 11:38:17.589278    1900 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
15:38:18: [control-plane-check] kube-apiserver is healthy after 1.502106535s
15:38:18: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:194: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A1341)

<details><summary>context</summary>
<pre>15:40:34: INFO: Running command line: bazel-bin/push-fedora-with-test-tooling-container-disk.exe --registry localhost:45413/kubevirt --repository fedora-with-test-tooling-container-disk --tag devel
15:40:36: Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries
15:40:38: &#43; rm -f /tmp/kubevirt.deploy.VzYI
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.36-sig-operator/2115466961587539968#1:build-log.txt%3A1350)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
15:40:38: selecting podman as container runtime
15:41:00: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-operator is being used by the following container(s): 072ddc794764862cead4ae6b6c5479ed96468412c055fb6419931fe946ae6cc7: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

#### internal (1x / 33.33%)

<details>
<summary> configured node not found during provisioning (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-11-02 12:10:13 &#43;0000 UTC_: <code>12:19:49: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A875)
<details>
<summary>all...</summary>

* _2026-11-02 12:10:13 &#43;0000 UTC_: <code>12:19:49: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A875)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A810)

<details><summary>context</summary>
<pre>12:19:44: [control-plane-check] kube-controller-manager is healthy after 8.488191ms
12:19:44: [control-plane-check] kube-scheduler is healthy after 10.387052ms
12:19:45: I1102 07:19:45.589042    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:45: I1102 07:19:46.088280    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:46: I1102 07:19:46.586947    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding</pre>
</details>



* <strong>diagnostic</strong>: <code>12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A812)

<details><summary>context</summary>
<pre>12:19:45: I1102 07:19:45.589042    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:45: I1102 07:19:46.088280    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:46: I1102 07:19:46.586947    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:47.090991    1901 uploadconfig.go:109] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap
12:19:46: [control-plane-check] kube-apiserver is healthy after 2.005710556s</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A1156)

<details><summary>context</summary>
<pre>12:20:04: 
12:20:04: &#43; wait 69
12:20:53: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A1184)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:20:53: selecting podman as container runtime
12:20:54: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-windows2016 is being used by the following container(s): 6a6a2973d141a1448dc31509ecfd824ede38d3435e163c24d8b9cba7915f160b: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

#### external (1x / 33.33%)

<details>
<summary> container image pull failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-11-02 09:09:03 &#43;0000 UTC_: <code>09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A433)
<details>
<summary>all...</summary>

* _2026-11-02 09:09:03 &#43;0000 UTC_: <code>09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A433)


<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A434)

<details><summary>context</summary>
<pre>09:16:19: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
09:16:19: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)
make: *** [Makefile:179: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>



</details>

<hr/>
</details>

### sig-storage (2x / 40.00%)


#### internal (1x / 50.00%)

<details>
<summary> configured node not found during provisioning (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-10-28 16:02:01 &#43;0000 UTC_: <code>16:12:51: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A862)
<details>
<summary>all...</summary>

* _2026-10-28 16:02:01 &#43;0000 UTC_: <code>16:12:51: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A862)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>16:12:48: I1028 12:12:48.374608    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A796)

<details><summary>context</summary>
<pre>16:12:47: [control-plane-check] kube-scheduler is healthy after 6.01446ms
16:12:47: [control-plane-check] kube-controller-manager is healthy after 7.342355ms
16:12:47: I1028 12:12:47.876825    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
16:12:48: I1028 12:12:48.374608    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
16:12:48: I1028 12:12:48.374663    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
16:12:48: I1028 12:12:48.875519    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
16:12:48: I1028 12:12:48.875556    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding</pre>
</details>



* <strong>diagnostic</strong>: <code>16:12:48: I1028 12:12:48.875519    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A798)

<details><summary>context</summary>
<pre>16:12:47: I1028 12:12:47.876825    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
16:12:48: I1028 12:12:48.374608    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
16:12:48: I1028 12:12:48.374663    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
16:12:48: I1028 12:12:48.875519    1883 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
16:12:48: I1028 12:12:48.875556    1883 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
16:12:49: [control-plane-check] kube-apiserver is healthy after 2.003454173s
16:12:49: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1157)

<details><summary>context</summary>
<pre>16:14:19: 
16:14:19: &#43; wait 66
16:14:39: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/2115473968377892864#1:build-log.txt%3A1185)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
16:14:40: selecting podman as container runtime
16:14:40: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-storage is being used by the following container(s): ab5ac3b271a5edd5373cc1acd2a27a8884d4400a81f5eed93a634e9d5f377457: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

#### needs-investigation (1x / 50.00%)

<details>
<summary> image layer unavailable in original registries (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-10-28 15:32:58 &#43;0000 UTC_: <code>15:59:16: Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2209)
<details>
<summary>all...</summary>

* _2026-10-28 15:32:58 &#43;0000 UTC_: <code>15:59:16: Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2209)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>15:38:26: I1028 15:38:26.005087    2040 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A831)

<details><summary>context</summary>
<pre>15:38:25: [control-plane-check] kube-scheduler is healthy after 5.33362ms
15:38:25: [control-plane-check] kube-controller-manager is healthy after 7.324903ms
15:38:25: I1028 15:38:25.506495    2040 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: rbac: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
15:38:26: I1028 15:38:26.005087    2040 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
15:38:26: I1028 15:38:26.005128    2040 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
15:38:26: [control-plane-check] kube-apiserver is healthy after 1.501955359s
15:38:26: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:194: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2211)

<details><summary>context</summary>
<pre>15:59:14: INFO: Running command line: bazel-bin/push-fedora-with-test-tooling-container-disk.exe --registry localhost:45605/kubevirt --repository kv-fedora-with-test-tooling-container-disk --tag devel
15:59:16: Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries
15:59:23: &#43; rm -f /tmp/kubevirt.deploy.EZG2
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2220)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
15:59:23: selecting podman as container runtime
15:59:38: Error response from daemon: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2236)

<details><summary>context</summary>
<pre>bf3625c46ab363fc7049174e5d93200f7e4750ad44ba342936e9b2e64bc10de4
24814c8573d55105ee1958eebc65d6bbadec60390bf3034e207bf414b25d324a
065a95016f2c9afd305c46ed8fdf873e16712de179534c5dccf2dfb28fd45dd4
Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1020 is not a child of this shell
================================================================================</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19294/pull-kubevirt-e2e-k8s-1.37-sig-storage/2115466429074509824#1:build-log.txt%3A2237)

<details><summary>context</summary>
<pre>24814c8573d55105ee1958eebc65d6bbadec60390bf3034e207bf414b25d324a
065a95016f2c9afd305c46ed8fdf873e16712de179534c5dccf2dfb28fd45dd4
Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container b487b02b7bd863b67c35b5b6ac7145b787206836b350ba5ccfc14c766051509a as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1020 is not a child of this shell
================================================================================
Done cleaning up after podman in container.</pre>
</details>


</details>



</details>

<hr/>
</details>

Last updated: 2026-11-04 15:10:18
