





<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-11-17 (1x / 25.00%)


#### external (1x / 100.00%)

<details>
<summary> container image pull failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-11-17 00:02:38 &#43;0000 UTC_: <code>00:18:43: Error: unable to copy from source docker://quay.io/kubevirt/builder:2610291047-177548615f: copying system image from manifest list: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1473617964/14&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/92/9250c70b3a36f0e8c52cce7cdfd8ca593cb87ac83647e43ff4455345c7b077d6?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T001820Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e6e8470c7d72bae4e2ee61190df3a3a1a13879e84e916329f5d572869ba0d850&amp;region=us-east-1&amp;namespace=kubevirt&amp;repo_name=builder&amp;akamai_signature=exp=1794875600~hmac=03a4000e343af8fd33d156c3a2fe1c9d34ab9ff4bfb4fe3a69d07e1367240492&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19931/pull-kubevirt-e2e-k8s-1.36-sig-compute/2122480343905734656#1:build-log.txt%3A951)
<details>
<summary>all...</summary>

* _2026-11-17 00:02:38 &#43;0000 UTC_: <code>00:18:43: Error: unable to copy from source docker://quay.io/kubevirt/builder:2610291047-177548615f: copying system image from manifest list: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1473617964/14&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/92/9250c70b3a36f0e8c52cce7cdfd8ca593cb87ac83647e43ff4455345c7b077d6?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T001820Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e6e8470c7d72bae4e2ee61190df3a3a1a13879e84e916329f5d572869ba0d850&amp;region=us-east-1&amp;namespace=kubevirt&amp;repo_name=builder&amp;akamai_signature=exp=1794875600~hmac=03a4000e343af8fd33d156c3a2fe1c9d34ab9ff4bfb4fe3a69d07e1367240492&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19931/pull-kubevirt-e2e-k8s-1.36-sig-compute/2122480343905734656#1:build-log.txt%3A951)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>00:17:16: I1116 19:17:15.813983    1908 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19931/pull-kubevirt-e2e-k8s-1.36-sig-compute/2122480343905734656#1:build-log.txt%3A614)

<details><summary>context</summary>
<pre>00:17:15: [control-plane-check] kube-controller-manager is healthy after 3.477922ms
00:17:15: [control-plane-check] kube-scheduler is healthy after 4.532682ms
00:17:15: I1116 19:17:15.314190    1908 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
00:17:16: I1116 19:17:15.813983    1908 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
00:17:16: I1116 19:17:15.814046    1908 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
00:17:16: [control-plane-check] kube-apiserver is healthy after 1.502841152s
00:17:16: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:194: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19931/pull-kubevirt-e2e-k8s-1.36-sig-compute/2122480343905734656#1:build-log.txt%3A953)

<details><summary>context</summary>
<pre>00:18:20: Copying blob sha256:a81cc226408ac493733771cc7935adb8774b5f5fa6ca02a471c7e62dbadff8f2
00:18:43: Error: unable to copy from source docker://quay.io/kubevirt/builder:2610291047-177548615f: copying system image from manifest list: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1473617964/14&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/92/9250c70b3a36f0e8c52cce7cdfd8ca593cb87ac83647e43ff4455345c7b077d6?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T001820Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e6e8470c7d72bae4e2ee61190df3a3a1a13879e84e916329f5d572869ba0d850&amp;region=us-east-1&amp;namespace=kubevirt&amp;repo_name=builder&amp;akamai_signature=exp=1794875600~hmac=03a4000e343af8fd33d156c3a2fe1c9d34ab9ff4bfb4fe3a69d07e1367240492&#34;: EOF)
00:18:43: &#43; rm -f /tmp/kubevirt.deploy.PIw7
make: *** [Makefile:194: cluster-sync] Error 125
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>


</details>



</details>

<hr/>
</details>

### 2026-11-13 (1x / 25.00%)


#### internal (1x / 100.00%)

<details>
<summary> VM SSH connection timeout during cluster provisioning (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-11-13 15:08:32 &#43;0000 UTC_: <code>15:19:01: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19097/pull-kubevirt-e2e-k8s-1.37-sig-monitoring/2121258758930698240#1:build-log.txt%3A793)
<details>
<summary>all...</summary>

* _2026-11-13 15:08:32 &#43;0000 UTC_: <code>15:19:01: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19097/pull-kubevirt-e2e-k8s-1.37-sig-monitoring/2121258758930698240#1:build-log.txt%3A793)


<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19097/pull-kubevirt-e2e-k8s-1.37-sig-monitoring/2121258758930698240#1:build-log.txt%3A794)

<details><summary>context</summary>
<pre>15:18:49: 
15:18:49: &#43; wait 72
15:19:01: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>



</details>

<hr/>
</details>

### 2026-11-11 (1x / 25.00%)


#### internal (1x / 100.00%)

<details>
<summary> configured node not found during provisioning (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-11-11 18:32:22 &#43;0000 UTC_: <code>18:39:11: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19610/pull-kubevirt-e2e-k8s-1.36-sig-compute/2120584309881442304#1:build-log.txt%3A885)
<details>
<summary>all...</summary>

* _2026-11-11 18:32:22 &#43;0000 UTC_: <code>18:39:11: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19610/pull-kubevirt-e2e-k8s-1.36-sig-compute/2120584309881442304#1:build-log.txt%3A885)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>18:39:08: I1111 13:39:08.307873    1915 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19610/pull-kubevirt-e2e-k8s-1.36-sig-compute/2120584309881442304#1:build-log.txt%3A821)

<details><summary>context</summary>
<pre>18:39:07: [control-plane-check] kube-controller-manager is healthy after 5.728524ms
18:39:07: [control-plane-check] kube-scheduler is healthy after 5.74229ms
18:39:07: I1111 13:39:07.808597    1915 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
18:39:08: I1111 13:39:08.307873    1915 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
18:39:08: I1111 13:39:08.307933    1915 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
18:39:08: [control-plane-check] kube-apiserver is healthy after 1.502758867s
18:39:08: I1111 13:39:08.807449    1915 uploadconfig.go:109] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19610/pull-kubevirt-e2e-k8s-1.36-sig-compute/2120584309881442304#1:build-log.txt%3A1194)

<details><summary>context</summary>
<pre>18:39:12: 
18:39:12: &#43; wait 75
18:39:23: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19610/pull-kubevirt-e2e-k8s-1.36-sig-compute/2120584309881442304#1:build-log.txt%3A1222)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
18:39:23: selecting podman as container runtime
18:39:23: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute is being used by the following container(s): fb15418c2e14cb45163c8d0b8a0c6872a738dc087c7b3cb0154f8b394c7585d2: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

### 2026-11-10 (1x / 25.00%)


#### internal (1x / 100.00%)

<details>
<summary> VM SSH connection timeout during cluster provisioning (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-11-10 10:02:30 &#43;0000 UTC_: <code>11:06:30: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19857/pull-kubevirt-e2e-k8s-1.34-sig-network-1.9/2120093328576876544#1:build-log.txt%3A1072)
<details>
<summary>all...</summary>

* _2026-11-10 10:02:30 &#43;0000 UTC_: <code>11:06:30: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19857/pull-kubevirt-e2e-k8s-1.34-sig-network-1.9/2120093328576876544#1:build-log.txt%3A1072)


<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19857/pull-kubevirt-e2e-k8s-1.34-sig-network-1.9/2120093328576876544#1:build-log.txt%3A1074)

<details><summary>context</summary>
<pre>11:06:19: 6audio: Could not create a backend for voice `ac97.mc&#39;
11:06:30: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed
11:06:31: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:174: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19857/pull-kubevirt-e2e-k8s-1.34-sig-network-1.9/2120093328576876544#1:build-log.txt%3A1100)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
11:06:31: selecting podman as container runtime
11:06:31: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-network-1.9 is being used by the following container(s): c127fab2def9926feeb32a7d5cb5b9252514f248d8bf91af09c29158931b8fc7: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### internal (3x / 75.00%)

<details>
<summary> VM SSH connection timeout during cluster provisioning (2x / 50.00%) </summary>

<hr/>

**2x**: _2026-11-13 15:08:32 &#43;0000 UTC_: <code>15:19:01: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19097/pull-kubevirt-e2e-k8s-1.37-sig-monitoring/2121258758930698240#1:build-log.txt%3A793)
<details>
<summary>all...</summary>

* _2026-11-13 15:08:32 &#43;0000 UTC_: <code>15:19:01: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19097/pull-kubevirt-e2e-k8s-1.37-sig-monitoring/2121258758930698240#1:build-log.txt%3A793)
<details><summary>context</summary>
<pre>15:18:49: &#43; [[ -n &#39;&#39; ]]
15:18:49: 
15:18:49: &#43; wait 72
15:19:01: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19097/pull-kubevirt-e2e-k8s-1.37-sig-monitoring/2121258758930698240#1:build-log.txt%3A794)

<details><summary>context</summary>
<pre>15:18:49: 
15:18:49: &#43; wait 72
15:19:01: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>



* _2026-11-10 10:02:30 &#43;0000 UTC_: <code>11:06:30: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19857/pull-kubevirt-e2e-k8s-1.34-sig-network-1.9/2120093328576876544#1:build-log.txt%3A1072)
<details><summary>context</summary>
<pre>11:06:19: .sdl: SDL_OpenAudioDevice for recording failed
11:06:19: &lt;sdl: Reason: ALSA: Couldn&#39;t open audio device: Host is down
11:06:19: 6audio: Could not create a backend for voice `ac97.mc&#39;
11:06:30: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed
11:06:31: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:174: cluster-up] Error 1
&#43;&#43; collect_debug_logs</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19857/pull-kubevirt-e2e-k8s-1.34-sig-network-1.9/2120093328576876544#1:build-log.txt%3A1074)

<details><summary>context</summary>
<pre>11:06:19: 6audio: Could not create a backend for voice `ac97.mc&#39;
11:06:30: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed
11:06:31: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:174: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19857/pull-kubevirt-e2e-k8s-1.34-sig-network-1.9/2120093328576876544#1:build-log.txt%3A1100)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
11:06:31: selecting podman as container runtime
11:06:31: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-network-1.9 is being used by the following container(s): c127fab2def9926feeb32a7d5cb5b9252514f248d8bf91af09c29158931b8fc7: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>
<details>
<summary> configured node not found during provisioning (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-11-11 18:32:22 &#43;0000 UTC_: <code>18:39:11: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19610/pull-kubevirt-e2e-k8s-1.36-sig-compute/2120584309881442304#1:build-log.txt%3A885)
<details>
<summary>all...</summary>

* _2026-11-11 18:32:22 &#43;0000 UTC_: <code>18:39:11: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19610/pull-kubevirt-e2e-k8s-1.36-sig-compute/2120584309881442304#1:build-log.txt%3A885)
<details><summary>context</summary>
<pre>18:39:11: Warning: would violate PodSecurity &#34;restricted:latest&#34;: host namespaces (hostNetwork=true), allowPrivilegeEscalation != false (containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.allowPrivilegeEscalation=false), unrestricted capabilities (containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.capabilities.drop=[&#34;ALL&#34;]; container &#34;kube-flannel&#34; must not include &#34;NET_ADMIN&#34;, &#34;NET_RAW&#34; in securityContext.capabilities.add), restricted volume types (volumes &#34;run&#34;, &#34;cni-plugin&#34;, &#34;cni&#34;, &#34;xtables-lock&#34; use restricted volume type &#34;hostPath&#34;), runAsNonRoot != true (pod or containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.runAsNonRoot=true), seccompProfile (pod or containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.seccompProfile.type to &#34;RuntimeDefault&#34; or &#34;Localhost&#34;)
18:39:11: daemonset.apps/kube-flannel-ds created
18:39:11: time=&#34;2026-11-11T18:39:11Z&#34; level=info msg=&#34;[node 1]: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-&#34;
18:39:11: Error from server (NotFound): nodes &#34;node01&#34; not found
18:39:11:
18:39:11: ===== d8f5f171b7b397292aded2e19587182f749f342550350e9f2a379b56657fde64 ====
18:39:11: &#43; NUM_NODES=2</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>18:39:08: I1111 13:39:08.307873    1915 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19610/pull-kubevirt-e2e-k8s-1.36-sig-compute/2120584309881442304#1:build-log.txt%3A821)

<details><summary>context</summary>
<pre>18:39:07: [control-plane-check] kube-controller-manager is healthy after 5.728524ms
18:39:07: [control-plane-check] kube-scheduler is healthy after 5.74229ms
18:39:07: I1111 13:39:07.808597    1915 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
18:39:08: I1111 13:39:08.307873    1915 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
18:39:08: I1111 13:39:08.307933    1915 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
18:39:08: [control-plane-check] kube-apiserver is healthy after 1.502758867s
18:39:08: I1111 13:39:08.807449    1915 uploadconfig.go:109] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19610/pull-kubevirt-e2e-k8s-1.36-sig-compute/2120584309881442304#1:build-log.txt%3A1194)

<details><summary>context</summary>
<pre>18:39:12: 
18:39:12: &#43; wait 75
18:39:23: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19610/pull-kubevirt-e2e-k8s-1.36-sig-compute/2120584309881442304#1:build-log.txt%3A1222)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
18:39:23: selecting podman as container runtime
18:39:23: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute is being used by the following container(s): fb15418c2e14cb45163c8d0b8a0c6872a738dc087c7b3cb0154f8b394c7585d2: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

### external (1x / 25.00%)

<details>
<summary> container image pull failure (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-11-17 00:02:38 &#43;0000 UTC_: <code>00:18:43: Error: unable to copy from source docker://quay.io/kubevirt/builder:2610291047-177548615f: copying system image from manifest list: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1473617964/14&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/92/9250c70b3a36f0e8c52cce7cdfd8ca593cb87ac83647e43ff4455345c7b077d6?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T001820Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e6e8470c7d72bae4e2ee61190df3a3a1a13879e84e916329f5d572869ba0d850&amp;region=us-east-1&amp;namespace=kubevirt&amp;repo_name=builder&amp;akamai_signature=exp=1794875600~hmac=03a4000e343af8fd33d156c3a2fe1c9d34ab9ff4bfb4fe3a69d07e1367240492&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19931/pull-kubevirt-e2e-k8s-1.36-sig-compute/2122480343905734656#1:build-log.txt%3A951)
<details>
<summary>all...</summary>

* _2026-11-17 00:02:38 &#43;0000 UTC_: <code>00:18:43: Error: unable to copy from source docker://quay.io/kubevirt/builder:2610291047-177548615f: copying system image from manifest list: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1473617964/14&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/92/9250c70b3a36f0e8c52cce7cdfd8ca593cb87ac83647e43ff4455345c7b077d6?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T001820Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e6e8470c7d72bae4e2ee61190df3a3a1a13879e84e916329f5d572869ba0d850&amp;region=us-east-1&amp;namespace=kubevirt&amp;repo_name=builder&amp;akamai_signature=exp=1794875600~hmac=03a4000e343af8fd33d156c3a2fe1c9d34ab9ff4bfb4fe3a69d07e1367240492&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19931/pull-kubevirt-e2e-k8s-1.36-sig-compute/2122480343905734656#1:build-log.txt%3A951)
<details><summary>context</summary>
<pre>00:18:19: Copying blob sha256:f32e658155c5b1add211a8a321dcc673c4f17fa3e86b2428ca8dc4ebf37a1fb9
00:18:19: Copying blob sha256:12d1784f2fa801320b2dc270f5056eefd21725f80c85d4e2a4a17ee3ef2b2e3e
00:18:20: Copying blob sha256:a81cc226408ac493733771cc7935adb8774b5f5fa6ca02a471c7e62dbadff8f2
00:18:43: Error: unable to copy from source docker://quay.io/kubevirt/builder:2610291047-177548615f: copying system image from manifest list: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1473617964/14&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/92/9250c70b3a36f0e8c52cce7cdfd8ca593cb87ac83647e43ff4455345c7b077d6?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T001820Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e6e8470c7d72bae4e2ee61190df3a3a1a13879e84e916329f5d572869ba0d850&amp;region=us-east-1&amp;namespace=kubevirt&amp;repo_name=builder&amp;akamai_signature=exp=1794875600~hmac=03a4000e343af8fd33d156c3a2fe1c9d34ab9ff4bfb4fe3a69d07e1367240492&#34;: EOF)
00:18:43: &#43; rm -f /tmp/kubevirt.deploy.PIw7
make: *** [Makefile:194: cluster-sync] Error 125
&#43; ret=2</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>00:17:16: I1116 19:17:15.813983    1908 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19931/pull-kubevirt-e2e-k8s-1.36-sig-compute/2122480343905734656#1:build-log.txt%3A614)

<details><summary>context</summary>
<pre>00:17:15: [control-plane-check] kube-controller-manager is healthy after 3.477922ms
00:17:15: [control-plane-check] kube-scheduler is healthy after 4.532682ms
00:17:15: I1116 19:17:15.314190    1908 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
00:17:16: I1116 19:17:15.813983    1908 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
00:17:16: I1116 19:17:15.814046    1908 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
00:17:16: [control-plane-check] kube-apiserver is healthy after 1.502841152s
00:17:16: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:194: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19931/pull-kubevirt-e2e-k8s-1.36-sig-compute/2122480343905734656#1:build-log.txt%3A953)

<details><summary>context</summary>
<pre>00:18:20: Copying blob sha256:a81cc226408ac493733771cc7935adb8774b5f5fa6ca02a471c7e62dbadff8f2
00:18:43: Error: unable to copy from source docker://quay.io/kubevirt/builder:2610291047-177548615f: copying system image from manifest list: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1473617964/14&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/92/9250c70b3a36f0e8c52cce7cdfd8ca593cb87ac83647e43ff4455345c7b077d6?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T001820Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e6e8470c7d72bae4e2ee61190df3a3a1a13879e84e916329f5d572869ba0d850&amp;region=us-east-1&amp;namespace=kubevirt&amp;repo_name=builder&amp;akamai_signature=exp=1794875600~hmac=03a4000e343af8fd33d156c3a2fe1c9d34ab9ff4bfb4fe3a69d07e1367240492&#34;: EOF)
00:18:43: &#43; rm -f /tmp/kubevirt.deploy.PIw7
make: *** [Makefile:194: cluster-sync] Error 125
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>


</details>



</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (3x / 75.00%)


#### internal (2x / 66.67%)

<details>
<summary> VM SSH connection timeout during cluster provisioning (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-11-13 15:08:32 &#43;0000 UTC_: <code>15:19:01: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19097/pull-kubevirt-e2e-k8s-1.37-sig-monitoring/2121258758930698240#1:build-log.txt%3A793)
<details>
<summary>all...</summary>

* _2026-11-13 15:08:32 &#43;0000 UTC_: <code>15:19:01: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19097/pull-kubevirt-e2e-k8s-1.37-sig-monitoring/2121258758930698240#1:build-log.txt%3A793)


<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19097/pull-kubevirt-e2e-k8s-1.37-sig-monitoring/2121258758930698240#1:build-log.txt%3A794)

<details><summary>context</summary>
<pre>15:18:49: 
15:18:49: &#43; wait 72
15:19:01: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>



</details>

<hr/>
</details>
<details>
<summary> configured node not found during provisioning (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-11-11 18:32:22 &#43;0000 UTC_: <code>18:39:11: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19610/pull-kubevirt-e2e-k8s-1.36-sig-compute/2120584309881442304#1:build-log.txt%3A885)
<details>
<summary>all...</summary>

* _2026-11-11 18:32:22 &#43;0000 UTC_: <code>18:39:11: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19610/pull-kubevirt-e2e-k8s-1.36-sig-compute/2120584309881442304#1:build-log.txt%3A885)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>18:39:08: I1111 13:39:08.307873    1915 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19610/pull-kubevirt-e2e-k8s-1.36-sig-compute/2120584309881442304#1:build-log.txt%3A821)

<details><summary>context</summary>
<pre>18:39:07: [control-plane-check] kube-controller-manager is healthy after 5.728524ms
18:39:07: [control-plane-check] kube-scheduler is healthy after 5.74229ms
18:39:07: I1111 13:39:07.808597    1915 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
18:39:08: I1111 13:39:08.307873    1915 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
18:39:08: I1111 13:39:08.307933    1915 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
18:39:08: [control-plane-check] kube-apiserver is healthy after 1.502758867s
18:39:08: I1111 13:39:08.807449    1915 uploadconfig.go:109] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19610/pull-kubevirt-e2e-k8s-1.36-sig-compute/2120584309881442304#1:build-log.txt%3A1194)

<details><summary>context</summary>
<pre>18:39:12: 
18:39:12: &#43; wait 75
18:39:23: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19610/pull-kubevirt-e2e-k8s-1.36-sig-compute/2120584309881442304#1:build-log.txt%3A1222)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
18:39:23: selecting podman as container runtime
18:39:23: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute is being used by the following container(s): fb15418c2e14cb45163c8d0b8a0c6872a738dc087c7b3cb0154f8b394c7585d2: volume is being used
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

**1x**: _2026-11-17 00:02:38 &#43;0000 UTC_: <code>00:18:43: Error: unable to copy from source docker://quay.io/kubevirt/builder:2610291047-177548615f: copying system image from manifest list: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1473617964/14&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/92/9250c70b3a36f0e8c52cce7cdfd8ca593cb87ac83647e43ff4455345c7b077d6?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T001820Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e6e8470c7d72bae4e2ee61190df3a3a1a13879e84e916329f5d572869ba0d850&amp;region=us-east-1&amp;namespace=kubevirt&amp;repo_name=builder&amp;akamai_signature=exp=1794875600~hmac=03a4000e343af8fd33d156c3a2fe1c9d34ab9ff4bfb4fe3a69d07e1367240492&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19931/pull-kubevirt-e2e-k8s-1.36-sig-compute/2122480343905734656#1:build-log.txt%3A951)
<details>
<summary>all...</summary>

* _2026-11-17 00:02:38 &#43;0000 UTC_: <code>00:18:43: Error: unable to copy from source docker://quay.io/kubevirt/builder:2610291047-177548615f: copying system image from manifest list: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1473617964/14&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/92/9250c70b3a36f0e8c52cce7cdfd8ca593cb87ac83647e43ff4455345c7b077d6?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T001820Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e6e8470c7d72bae4e2ee61190df3a3a1a13879e84e916329f5d572869ba0d850&amp;region=us-east-1&amp;namespace=kubevirt&amp;repo_name=builder&amp;akamai_signature=exp=1794875600~hmac=03a4000e343af8fd33d156c3a2fe1c9d34ab9ff4bfb4fe3a69d07e1367240492&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19931/pull-kubevirt-e2e-k8s-1.36-sig-compute/2122480343905734656#1:build-log.txt%3A951)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>00:17:16: I1116 19:17:15.813983    1908 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19931/pull-kubevirt-e2e-k8s-1.36-sig-compute/2122480343905734656#1:build-log.txt%3A614)

<details><summary>context</summary>
<pre>00:17:15: [control-plane-check] kube-controller-manager is healthy after 3.477922ms
00:17:15: [control-plane-check] kube-scheduler is healthy after 4.532682ms
00:17:15: I1116 19:17:15.314190    1908 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
00:17:16: I1116 19:17:15.813983    1908 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
00:17:16: I1116 19:17:15.814046    1908 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
00:17:16: [control-plane-check] kube-apiserver is healthy after 1.502841152s
00:17:16: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:194: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19931/pull-kubevirt-e2e-k8s-1.36-sig-compute/2122480343905734656#1:build-log.txt%3A953)

<details><summary>context</summary>
<pre>00:18:20: Copying blob sha256:a81cc226408ac493733771cc7935adb8774b5f5fa6ca02a471c7e62dbadff8f2
00:18:43: Error: unable to copy from source docker://quay.io/kubevirt/builder:2610291047-177548615f: copying system image from manifest list: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1473617964/14&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/92/9250c70b3a36f0e8c52cce7cdfd8ca593cb87ac83647e43ff4455345c7b077d6?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T001820Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e6e8470c7d72bae4e2ee61190df3a3a1a13879e84e916329f5d572869ba0d850&amp;region=us-east-1&amp;namespace=kubevirt&amp;repo_name=builder&amp;akamai_signature=exp=1794875600~hmac=03a4000e343af8fd33d156c3a2fe1c9d34ab9ff4bfb4fe3a69d07e1367240492&#34;: EOF)
00:18:43: &#43; rm -f /tmp/kubevirt.deploy.PIw7
make: *** [Makefile:194: cluster-sync] Error 125
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>


</details>



</details>

<hr/>
</details>

### release-1.9 (1x / 25.00%)


#### internal (1x / 100.00%)

<details>
<summary> VM SSH connection timeout during cluster provisioning (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-11-10 10:02:30 &#43;0000 UTC_: <code>11:06:30: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19857/pull-kubevirt-e2e-k8s-1.34-sig-network-1.9/2120093328576876544#1:build-log.txt%3A1072)
<details>
<summary>all...</summary>

* _2026-11-10 10:02:30 &#43;0000 UTC_: <code>11:06:30: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19857/pull-kubevirt-e2e-k8s-1.34-sig-network-1.9/2120093328576876544#1:build-log.txt%3A1072)


<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19857/pull-kubevirt-e2e-k8s-1.34-sig-network-1.9/2120093328576876544#1:build-log.txt%3A1074)

<details><summary>context</summary>
<pre>11:06:19: 6audio: Could not create a backend for voice `ac97.mc&#39;
11:06:30: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed
11:06:31: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:174: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19857/pull-kubevirt-e2e-k8s-1.34-sig-network-1.9/2120093328576876544#1:build-log.txt%3A1100)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
11:06:31: selecting podman as container runtime
11:06:31: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-network-1.9 is being used by the following container(s): c127fab2def9926feeb32a7d5cb5b9252514f248d8bf91af09c29158931b8fc7: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
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


### sig-compute (2x / 50.00%)


#### external (1x / 50.00%)

<details>
<summary> container image pull failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-11-17 00:02:38 &#43;0000 UTC_: <code>00:18:43: Error: unable to copy from source docker://quay.io/kubevirt/builder:2610291047-177548615f: copying system image from manifest list: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1473617964/14&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/92/9250c70b3a36f0e8c52cce7cdfd8ca593cb87ac83647e43ff4455345c7b077d6?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T001820Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e6e8470c7d72bae4e2ee61190df3a3a1a13879e84e916329f5d572869ba0d850&amp;region=us-east-1&amp;namespace=kubevirt&amp;repo_name=builder&amp;akamai_signature=exp=1794875600~hmac=03a4000e343af8fd33d156c3a2fe1c9d34ab9ff4bfb4fe3a69d07e1367240492&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19931/pull-kubevirt-e2e-k8s-1.36-sig-compute/2122480343905734656#1:build-log.txt%3A951)
<details>
<summary>all...</summary>

* _2026-11-17 00:02:38 &#43;0000 UTC_: <code>00:18:43: Error: unable to copy from source docker://quay.io/kubevirt/builder:2610291047-177548615f: copying system image from manifest list: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1473617964/14&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/92/9250c70b3a36f0e8c52cce7cdfd8ca593cb87ac83647e43ff4455345c7b077d6?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T001820Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e6e8470c7d72bae4e2ee61190df3a3a1a13879e84e916329f5d572869ba0d850&amp;region=us-east-1&amp;namespace=kubevirt&amp;repo_name=builder&amp;akamai_signature=exp=1794875600~hmac=03a4000e343af8fd33d156c3a2fe1c9d34ab9ff4bfb4fe3a69d07e1367240492&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19931/pull-kubevirt-e2e-k8s-1.36-sig-compute/2122480343905734656#1:build-log.txt%3A951)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>00:17:16: I1116 19:17:15.813983    1908 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19931/pull-kubevirt-e2e-k8s-1.36-sig-compute/2122480343905734656#1:build-log.txt%3A614)

<details><summary>context</summary>
<pre>00:17:15: [control-plane-check] kube-controller-manager is healthy after 3.477922ms
00:17:15: [control-plane-check] kube-scheduler is healthy after 4.532682ms
00:17:15: I1116 19:17:15.314190    1908 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
00:17:16: I1116 19:17:15.813983    1908 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
00:17:16: I1116 19:17:15.814046    1908 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
00:17:16: [control-plane-check] kube-apiserver is healthy after 1.502841152s
00:17:16: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:194: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19931/pull-kubevirt-e2e-k8s-1.36-sig-compute/2122480343905734656#1:build-log.txt%3A953)

<details><summary>context</summary>
<pre>00:18:20: Copying blob sha256:a81cc226408ac493733771cc7935adb8774b5f5fa6ca02a471c7e62dbadff8f2
00:18:43: Error: unable to copy from source docker://quay.io/kubevirt/builder:2610291047-177548615f: copying system image from manifest list: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1473617964/14&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/92/9250c70b3a36f0e8c52cce7cdfd8ca593cb87ac83647e43ff4455345c7b077d6?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T001820Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e6e8470c7d72bae4e2ee61190df3a3a1a13879e84e916329f5d572869ba0d850&amp;region=us-east-1&amp;namespace=kubevirt&amp;repo_name=builder&amp;akamai_signature=exp=1794875600~hmac=03a4000e343af8fd33d156c3a2fe1c9d34ab9ff4bfb4fe3a69d07e1367240492&#34;: EOF)
00:18:43: &#43; rm -f /tmp/kubevirt.deploy.PIw7
make: *** [Makefile:194: cluster-sync] Error 125
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>


</details>



</details>

<hr/>
</details>

#### internal (1x / 50.00%)

<details>
<summary> configured node not found during provisioning (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-11-11 18:32:22 &#43;0000 UTC_: <code>18:39:11: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19610/pull-kubevirt-e2e-k8s-1.36-sig-compute/2120584309881442304#1:build-log.txt%3A885)
<details>
<summary>all...</summary>

* _2026-11-11 18:32:22 &#43;0000 UTC_: <code>18:39:11: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19610/pull-kubevirt-e2e-k8s-1.36-sig-compute/2120584309881442304#1:build-log.txt%3A885)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>18:39:08: I1111 13:39:08.307873    1915 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19610/pull-kubevirt-e2e-k8s-1.36-sig-compute/2120584309881442304#1:build-log.txt%3A821)

<details><summary>context</summary>
<pre>18:39:07: [control-plane-check] kube-controller-manager is healthy after 5.728524ms
18:39:07: [control-plane-check] kube-scheduler is healthy after 5.74229ms
18:39:07: I1111 13:39:07.808597    1915 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
18:39:08: I1111 13:39:08.307873    1915 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
18:39:08: I1111 13:39:08.307933    1915 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
18:39:08: [control-plane-check] kube-apiserver is healthy after 1.502758867s
18:39:08: I1111 13:39:08.807449    1915 uploadconfig.go:109] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19610/pull-kubevirt-e2e-k8s-1.36-sig-compute/2120584309881442304#1:build-log.txt%3A1194)

<details><summary>context</summary>
<pre>18:39:12: 
18:39:12: &#43; wait 75
18:39:23: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19610/pull-kubevirt-e2e-k8s-1.36-sig-compute/2120584309881442304#1:build-log.txt%3A1222)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
18:39:23: selecting podman as container runtime
18:39:23: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute is being used by the following container(s): fb15418c2e14cb45163c8d0b8a0c6872a738dc087c7b3cb0154f8b394c7585d2: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

### sig-monitoring (1x / 25.00%)


#### internal (1x / 100.00%)

<details>
<summary> VM SSH connection timeout during cluster provisioning (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-11-13 15:08:32 &#43;0000 UTC_: <code>15:19:01: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19097/pull-kubevirt-e2e-k8s-1.37-sig-monitoring/2121258758930698240#1:build-log.txt%3A793)
<details>
<summary>all...</summary>

* _2026-11-13 15:08:32 &#43;0000 UTC_: <code>15:19:01: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19097/pull-kubevirt-e2e-k8s-1.37-sig-monitoring/2121258758930698240#1:build-log.txt%3A793)


<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19097/pull-kubevirt-e2e-k8s-1.37-sig-monitoring/2121258758930698240#1:build-log.txt%3A794)

<details><summary>context</summary>
<pre>15:18:49: 
15:18:49: &#43; wait 72
15:19:01: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>



</details>

<hr/>
</details>

### sig-network (1x / 25.00%)


#### internal (1x / 100.00%)

<details>
<summary> VM SSH connection timeout during cluster provisioning (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-11-10 10:02:30 &#43;0000 UTC_: <code>11:06:30: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19857/pull-kubevirt-e2e-k8s-1.34-sig-network-1.9/2120093328576876544#1:build-log.txt%3A1072)
<details>
<summary>all...</summary>

* _2026-11-10 10:02:30 &#43;0000 UTC_: <code>11:06:30: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19857/pull-kubevirt-e2e-k8s-1.34-sig-network-1.9/2120093328576876544#1:build-log.txt%3A1072)


<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>make: *** [Makefile:174: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19857/pull-kubevirt-e2e-k8s-1.34-sig-network-1.9/2120093328576876544#1:build-log.txt%3A1074)

<details><summary>context</summary>
<pre>11:06:19: 6audio: Could not create a backend for voice `ac97.mc&#39;
11:06:30: could not establish a connection to the node after a generous timeout: ssh.sh echo VM is up failed
11:06:31: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:174: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19857/pull-kubevirt-e2e-k8s-1.34-sig-network-1.9/2120093328576876544#1:build-log.txt%3A1100)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
11:06:31: selecting podman as container runtime
11:06:31: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-sig-network-1.9 is being used by the following container(s): c127fab2def9926feeb32a7d5cb5b9252514f248d8bf91af09c29158931b8fc7: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

Last updated: 2026-11-17 09:10:08
