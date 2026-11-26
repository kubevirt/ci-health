





<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-11-23 (3x / 75.00%)


#### internal (2x / 66.67%)

<details>
<summary> configured node not found during provisioning (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-11-23 18:50:51 &#43;0000 UTC_: <code>18:52:09: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19479/pull-kubevirt-e2e-k8s-1.36-sig-network/2124925311807131648#1:build-log.txt%3A698)
<details>
<summary>all...</summary>

* _2026-11-23 18:50:51 &#43;0000 UTC_: <code>18:52:09: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19479/pull-kubevirt-e2e-k8s-1.36-sig-network/2124925311807131648#1:build-log.txt%3A698)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>18:52:07: I1123 13:52:06.874315    1962 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19479/pull-kubevirt-e2e-k8s-1.36-sig-network/2124925311807131648#1:build-log.txt%3A634)

<details><summary>context</summary>
<pre>18:52:06: [control-plane-check] kube-controller-manager is healthy after 8.514229ms
18:52:06: [control-plane-check] kube-scheduler is healthy after 9.410102ms
18:52:06: I1123 13:52:06.374210    1962 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
18:52:07: I1123 13:52:06.874315    1962 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
18:52:07: I1123 13:52:06.874358    1962 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
18:52:07: I1123 13:52:07.377532    1962 uploadconfig.go:109] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap
18:52:07: [control-plane-check] kube-apiserver is healthy after 1.504955242s</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19479/pull-kubevirt-e2e-k8s-1.36-sig-network/2124925311807131648#1:build-log.txt%3A1022)

<details><summary>context</summary>
<pre>18:52:09: 
18:52:09: &#43; wait 69
18:52:20: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
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
<summary> make cluster lifecycle target failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-11-23 10:05:27 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/20015/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2124803835099615232#1:build-log.txt%3A1193)
<details>
<summary>all...</summary>

* _2026-11-23 10:05:27 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/20015/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2124803835099615232#1:build-log.txt%3A1193)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>10:19:51: I1123 05:19:51.532976    1591 request.go:1664] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/20015/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2124803835099615232#1:build-log.txt%3A763)

<details><summary>context</summary>
<pre>10:19:48: [control-plane-check] Checking kube-scheduler at https://127.0.0.1:10259/livez
10:19:50: [control-plane-check] kube-scheduler is healthy after 2.118400411s
10:19:50: [control-plane-check] kube-controller-manager is healthy after 2.603560978s
10:19:51: I1123 05:19:51.532976    1591 request.go:1664] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
10:19:52: [control-plane-check] kube-apiserver is healthy after 4.003281102s
10:19:52: I1123 05:19:52.038341    1591 kubeconfig.go:665] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists
10:19:52: I1123 05:19:52.041457    1591 kubeconfig.go:738] creating the ClusterRoleBinding for the kubeadm:cluster-admins Group by using super-admin.conf</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/20015/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2124803835099615232#1:build-log.txt%3A1199)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
10:25:11: selecting podman as container runtime
10:25:45: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7 is being used by the following container(s): 78694673e778d74e782e667f11be6db389a0f55e5c27c234fc87aebe8e9f674a, 05ff6ad545874ae0691c6a1bfba03f1a54dcd66dd4ef9c64d883795eafa6648b: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

#### needs-investigation (1x / 33.33%)

<details>
<summary> no matching pattern (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-11-23 10:23:49 &#43;0000 UTC_: <code>10:27:48: ERROR: no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1215)
<details>
<summary>all...</summary>

* _2026-11-23 10:23:49 &#43;0000 UTC_: <code>10:27:48: ERROR: no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1215)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>10:24:21: I1123 10:24:20.740427    2085 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A635)

<details><summary>context</summary>
<pre>10:24:20: [control-plane-check] kube-scheduler is healthy after 4.814576ms
10:24:20: [control-plane-check] kube-controller-manager is healthy after 8.72848ms
10:24:20: I1123 10:24:20.244157    2085 wait.go:283] kube-apiserver check failed at https://[fd00::101]:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: rbac: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
10:24:21: I1123 10:24:20.740427    2085 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
10:24:21: I1123 10:24:20.740464    2085 wait.go:283] kube-apiserver check failed at https://[fd00::101]:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
10:24:21: [control-plane-check] kube-apiserver is healthy after 1.502020037s
10:24:21: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>10:27:47: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/internal/go_repository.bzl:297:13: An error occurred during the fetch of repository &#39;gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1209)

<details><summary>context</summary>
<pre>10:27:47: INFO: Repo gazelle&#43;&#43;go_deps&#43;org_golang_x_mod defined by rule go_repository in @@gazelle&#43;//internal:go_repository.bzl
10:27:47: INFO: Repo gazelle&#43;&#43;go_deps&#43;com_github_bmatcuk_doublestar_v4 defined by rule go_repository in @@gazelle&#43;//internal:go_repository.bzl
10:27:47: INFO: Repo gazelle&#43;&#43;go_deps&#43;com_github_bazelbuild_buildtools defined by rule go_repository in @@gazelle&#43;//internal:go_repository.bzl
10:27:47: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/internal/go_repository.bzl:297:13: An error occurred during the fetch of repository &#39;gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs&#39;:
10:27:47:    Traceback (most recent call last):
10:27:47: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/internal/go_repository.bzl&#34;, line 297, column 13, in _go_repository_impl
10:27:47: 		fail(&#34;%s: %s&#34; % (ctx.name, result.stderr))</pre>
</details>



* <strong>secondary</strong>: <code>10:27:49: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/repo/BUILD.bazel:3:11: @@gazelle&#43;//repo:repo depends on @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//:vcs in repository @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs which failed to fetch. no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1217)

<details><summary>context</summary>
<pre>10:27:47: exit status 1
10:27:48: ERROR: no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout
10:27:48: exit status 1
10:27:49: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/repo/BUILD.bazel:3:11: @@gazelle&#43;//repo:repo depends on @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//:vcs in repository @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs which failed to fetch. no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout
10:27:49: exit status 1
10:27:49: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
10:27:49: INFO: Elapsed time: 75.163s, Critical Path: 0.61s</pre>
</details>



* <strong>secondary</strong>: <code>10:27:49: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1219)

<details><summary>context</summary>
<pre>10:27:48: exit status 1
10:27:49: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/repo/BUILD.bazel:3:11: @@gazelle&#43;//repo:repo depends on @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//:vcs in repository @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs which failed to fetch. no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout
10:27:49: exit status 1
10:27:49: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
10:27:49: INFO: Elapsed time: 75.163s, Critical Path: 0.61s
10:27:49: INFO: 1 process: 1 internal.
10:27:49: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>secondary</strong>: <code>10:27:49: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1222)

<details><summary>context</summary>
<pre>10:27:49: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
10:27:49: INFO: Elapsed time: 75.163s, Critical Path: 0.61s
10:27:49: INFO: 1 process: 1 internal.
10:27:49: ERROR: Build did NOT complete successfully
10:27:49: ERROR: Build failed. Not running target
10:27:51: &#43; rm -f /tmp/kubevirt.deploy.9ro5
make: *** [Makefile:194: cluster-sync] Error 1</pre>
</details>



* <strong>secondary</strong>: <code>10:27:49: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1223)

<details><summary>context</summary>
<pre>10:27:49: INFO: Elapsed time: 75.163s, Critical Path: 0.61s
10:27:49: INFO: 1 process: 1 internal.
10:27:49: ERROR: Build did NOT complete successfully
10:27:49: ERROR: Build failed. Not running target
10:27:51: &#43; rm -f /tmp/kubevirt.deploy.9ro5
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:194: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1225)

<details><summary>context</summary>
<pre>10:27:49: ERROR: Build did NOT complete successfully
10:27:49: ERROR: Build failed. Not running target
10:27:51: &#43; rm -f /tmp/kubevirt.deploy.9ro5
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1234)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
10:27:52: selecting podman as container runtime
10:28:35: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network is being used by the following container(s): 73f69f92b0ee12ba15fe06ff081855efc21f25241a9104c96303a06d01ed229d: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

### 2026-11-21 (1x / 25.00%)


#### needs-investigation (1x / 100.00%)

<details>
<summary> no matching pattern (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-11-21 03:28:48 &#43;0000 UTC_: <code>03:39:24: Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1542)
<details>
<summary>all...</summary>

* _2026-11-21 03:28:48 &#43;0000 UTC_: <code>03:39:24: Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1542)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>03:29:18: I1121 03:29:18.058174    2039 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A632)

<details><summary>context</summary>
<pre>03:29:17: [control-plane-check] kube-scheduler is healthy after 10.206452ms
03:29:17: [control-plane-check] kube-controller-manager is healthy after 12.069006ms
03:29:17: I1121 03:29:17.559246    2039 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: rbac: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
03:29:18: I1121 03:29:18.058174    2039 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
03:29:18: I1121 03:29:18.058228    2039 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
03:29:18: [control-plane-check] kube-apiserver is healthy after 1.502576755s
03:29:18: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>03:39:24: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1543)

<details><summary>context</summary>
<pre>03:39:22:     layer transfers: 0 from disk, 14 from container registry, 0 from remote cache
03:39:22: INFO: Invocation ID: eb79cc26-99a6-4945-aeae-b12f2226c7ff
03:39:24: Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:24: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: INFO: Repo &#43;virt_template_ext&#43;virt_template_controller defined by rule pull in @@rules_img&#43;//img/private/repository_rules:pull.bzl
03:39:25: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/download.bzl:441:13: An error occurred during the fetch of repository &#39;&#43;virt_template_ext&#43;virt_template_controller&#39;:
03:39:25:    Traceback (most recent call last):</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/download.bzl:441:13: An error occurred during the fetch of repository &#39;&#43;virt_template_ext&#43;virt_template_controller&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1545)

<details><summary>context</summary>
<pre>03:39:24: Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:24: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: INFO: Repo &#43;virt_template_ext&#43;virt_template_controller defined by rule pull in @@rules_img&#43;//img/private/repository_rules:pull.bzl
03:39:25: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/download.bzl:441:13: An error occurred during the fetch of repository &#39;&#43;virt_template_ext&#43;virt_template_controller&#39;:
03:39:25:    Traceback (most recent call last):
03:39:25: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/pull.bzl&#34;, line 48, column 28, in _pull_impl
03:39:25: 		_download_with_tool(</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: Error in fail: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1551)

<details><summary>context</summary>
<pre>03:39:25: 		_download_with_tool(
03:39:25: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/download.bzl&#34;, line 441, column 13, in download_with_tool
03:39:25: 		fail(&#34;img tool failed with exit code {} and message {}&#34;.format(result.return_code, result.stderr))
03:39:25: Error in fail: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1552)

<details><summary>context</summary>
<pre>03:39:25: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/download.bzl&#34;, line 441, column 13, in download_with_tool
03:39:25: 		fail(&#34;img tool failed with exit code {} and message {}&#34;.format(result.return_code, result.stderr))
03:39:25: Error in fail: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1553)

<details><summary>context</summary>
<pre>03:39:25: 		fail(&#34;img tool failed with exit code {} and message {}&#34;.format(result.return_code, result.stderr))
03:39:25: Error in fail: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1554)

<details><summary>context</summary>
<pre>03:39:25: Error in fail: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1555)

<details><summary>context</summary>
<pre>03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed
03:39:25: INFO: Elapsed time: 2.254s, Critical Path: 0.16s</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1556)

<details><summary>context</summary>
<pre>03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed
03:39:25: INFO: Elapsed time: 2.254s, Critical Path: 0.16s
03:39:25: INFO: 1 process: 1 internal.</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1557)

<details><summary>context</summary>
<pre>03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed
03:39:25: INFO: Elapsed time: 2.254s, Critical Path: 0.16s
03:39:25: INFO: 1 process: 1 internal.
03:39:25: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1560)

<details><summary>context</summary>
<pre>03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed
03:39:25: INFO: Elapsed time: 2.254s, Critical Path: 0.16s
03:39:25: INFO: 1 process: 1 internal.
03:39:25: ERROR: Build did NOT complete successfully
03:39:25: ERROR: Build failed. Not running target
03:39:26: &#43; rm -f /tmp/kubevirt.deploy.Yb3J
make: *** [Makefile:194: cluster-sync] Error 1</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1561)

<details><summary>context</summary>
<pre>03:39:25: INFO: Elapsed time: 2.254s, Critical Path: 0.16s
03:39:25: INFO: 1 process: 1 internal.
03:39:25: ERROR: Build did NOT complete successfully
03:39:25: ERROR: Build failed. Not running target
03:39:26: &#43; rm -f /tmp/kubevirt.deploy.Yb3J
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:194: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1563)

<details><summary>context</summary>
<pre>03:39:25: ERROR: Build did NOT complete successfully
03:39:25: ERROR: Build failed. Not running target
03:39:26: &#43; rm -f /tmp/kubevirt.deploy.Yb3J
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1572)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
03:39:26: selecting podman as container runtime
03:39:57: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.37-sig-storage is being used by the following container(s): 4db2389c3fdcc1a55455a1885b7cd013063cca177a8b8947fd884f25a3975020: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
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


### internal (2x / 50.00%)

<details>
<summary> configured node not found during provisioning (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-11-23 18:50:51 &#43;0000 UTC_: <code>18:52:09: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19479/pull-kubevirt-e2e-k8s-1.36-sig-network/2124925311807131648#1:build-log.txt%3A698)
<details>
<summary>all...</summary>

* _2026-11-23 18:50:51 &#43;0000 UTC_: <code>18:52:09: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19479/pull-kubevirt-e2e-k8s-1.36-sig-network/2124925311807131648#1:build-log.txt%3A698)
<details><summary>context</summary>
<pre>18:52:09: Warning: would violate PodSecurity &#34;restricted:latest&#34;: host namespaces (hostNetwork=true), allowPrivilegeEscalation != false (containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.allowPrivilegeEscalation=false), unrestricted capabilities (containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.capabilities.drop=[&#34;ALL&#34;]; container &#34;kube-flannel&#34; must not include &#34;NET_ADMIN&#34;, &#34;NET_RAW&#34; in securityContext.capabilities.add), restricted volume types (volumes &#34;run&#34;, &#34;cni-plugin&#34;, &#34;cni&#34;, &#34;xtables-lock&#34; use restricted volume type &#34;hostPath&#34;), runAsNonRoot != true (pod or containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.runAsNonRoot=true), seccompProfile (pod or containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.seccompProfile.type to &#34;RuntimeDefault&#34; or &#34;Localhost&#34;)
18:52:09: daemonset.apps/kube-flannel-ds created
18:52:09: time=&#34;2026-11-23T18:52:09Z&#34; level=info msg=&#34;[node 1]: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-&#34;
18:52:09: Error from server (NotFound): nodes &#34;node01&#34; not found
18:52:09:
18:52:09: ===== b9dda4cde9a36505f90294a03e151cba907a32cdc6c40b26be796a4aed6c516f ====
18:52:09: &#43; NUM_NODES=3</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>18:52:07: I1123 13:52:06.874315    1962 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19479/pull-kubevirt-e2e-k8s-1.36-sig-network/2124925311807131648#1:build-log.txt%3A634)

<details><summary>context</summary>
<pre>18:52:06: [control-plane-check] kube-controller-manager is healthy after 8.514229ms
18:52:06: [control-plane-check] kube-scheduler is healthy after 9.410102ms
18:52:06: I1123 13:52:06.374210    1962 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
18:52:07: I1123 13:52:06.874315    1962 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
18:52:07: I1123 13:52:06.874358    1962 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
18:52:07: I1123 13:52:07.377532    1962 uploadconfig.go:109] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap
18:52:07: [control-plane-check] kube-apiserver is healthy after 1.504955242s</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19479/pull-kubevirt-e2e-k8s-1.36-sig-network/2124925311807131648#1:build-log.txt%3A1022)

<details><summary>context</summary>
<pre>18:52:09: 
18:52:09: &#43; wait 69
18:52:20: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
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
<summary> make cluster lifecycle target failure (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-11-23 10:05:27 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/20015/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2124803835099615232#1:build-log.txt%3A1193)
<details>
<summary>all...</summary>

* _2026-11-23 10:05:27 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/20015/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2124803835099615232#1:build-log.txt%3A1193)
<details><summary>context</summary>
<pre>10:24:20: INFO: 2 processes: 1 internal, 1 processwrapper-sandbox.
10:24:20: INFO: Running command line: bazel-bin/example-guest-agent-copier /root/go/src/kubevirt.io/kubevirt/_out/cmd/example-guest-agent/example-guest-agent
10:25:10: &#43; rm -f /tmp/kubevirt.deploy.HRnv
make: *** [Makefile:174: cluster-sync] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>10:19:51: I1123 05:19:51.532976    1591 request.go:1664] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/20015/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2124803835099615232#1:build-log.txt%3A763)

<details><summary>context</summary>
<pre>10:19:48: [control-plane-check] Checking kube-scheduler at https://127.0.0.1:10259/livez
10:19:50: [control-plane-check] kube-scheduler is healthy after 2.118400411s
10:19:50: [control-plane-check] kube-controller-manager is healthy after 2.603560978s
10:19:51: I1123 05:19:51.532976    1591 request.go:1664] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
10:19:52: [control-plane-check] kube-apiserver is healthy after 4.003281102s
10:19:52: I1123 05:19:52.038341    1591 kubeconfig.go:665] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists
10:19:52: I1123 05:19:52.041457    1591 kubeconfig.go:738] creating the ClusterRoleBinding for the kubeadm:cluster-admins Group by using super-admin.conf</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/20015/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2124803835099615232#1:build-log.txt%3A1199)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
10:25:11: selecting podman as container runtime
10:25:45: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7 is being used by the following container(s): 78694673e778d74e782e667f11be6db389a0f55e5c27c234fc87aebe8e9f674a, 05ff6ad545874ae0691c6a1bfba03f1a54dcd66dd4ef9c64d883795eafa6648b: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

### needs-investigation (2x / 50.00%)

<details>
<summary> no matching pattern (2x / 50.00%) </summary>

<hr/>

**1x**: _2026-11-23 10:23:49 &#43;0000 UTC_: <code>10:27:48: ERROR: no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1215)
<details>
<summary>all...</summary>

* _2026-11-23 10:23:49 &#43;0000 UTC_: <code>10:27:48: ERROR: no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1215)
<details><summary>context</summary>
<pre>10:27:47: 		fail(&#34;%s: %s&#34; % (ctx.name, result.stderr))
10:27:47: Error in fail: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout
10:27:47: exit status 1
10:27:48: ERROR: no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout
10:27:48: exit status 1
10:27:49: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/repo/BUILD.bazel:3:11: @@gazelle&#43;//repo:repo depends on @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//:vcs in repository @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs which failed to fetch. no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout
10:27:49: exit status 1</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>10:24:21: I1123 10:24:20.740427    2085 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A635)

<details><summary>context</summary>
<pre>10:24:20: [control-plane-check] kube-scheduler is healthy after 4.814576ms
10:24:20: [control-plane-check] kube-controller-manager is healthy after 8.72848ms
10:24:20: I1123 10:24:20.244157    2085 wait.go:283] kube-apiserver check failed at https://[fd00::101]:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: rbac: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
10:24:21: I1123 10:24:20.740427    2085 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
10:24:21: I1123 10:24:20.740464    2085 wait.go:283] kube-apiserver check failed at https://[fd00::101]:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
10:24:21: [control-plane-check] kube-apiserver is healthy after 1.502020037s
10:24:21: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>10:27:47: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/internal/go_repository.bzl:297:13: An error occurred during the fetch of repository &#39;gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1209)

<details><summary>context</summary>
<pre>10:27:47: INFO: Repo gazelle&#43;&#43;go_deps&#43;org_golang_x_mod defined by rule go_repository in @@gazelle&#43;//internal:go_repository.bzl
10:27:47: INFO: Repo gazelle&#43;&#43;go_deps&#43;com_github_bmatcuk_doublestar_v4 defined by rule go_repository in @@gazelle&#43;//internal:go_repository.bzl
10:27:47: INFO: Repo gazelle&#43;&#43;go_deps&#43;com_github_bazelbuild_buildtools defined by rule go_repository in @@gazelle&#43;//internal:go_repository.bzl
10:27:47: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/internal/go_repository.bzl:297:13: An error occurred during the fetch of repository &#39;gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs&#39;:
10:27:47:    Traceback (most recent call last):
10:27:47: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/internal/go_repository.bzl&#34;, line 297, column 13, in _go_repository_impl
10:27:47: 		fail(&#34;%s: %s&#34; % (ctx.name, result.stderr))</pre>
</details>



* <strong>secondary</strong>: <code>10:27:49: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/repo/BUILD.bazel:3:11: @@gazelle&#43;//repo:repo depends on @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//:vcs in repository @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs which failed to fetch. no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1217)

<details><summary>context</summary>
<pre>10:27:47: exit status 1
10:27:48: ERROR: no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout
10:27:48: exit status 1
10:27:49: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/repo/BUILD.bazel:3:11: @@gazelle&#43;//repo:repo depends on @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//:vcs in repository @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs which failed to fetch. no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout
10:27:49: exit status 1
10:27:49: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
10:27:49: INFO: Elapsed time: 75.163s, Critical Path: 0.61s</pre>
</details>



* <strong>secondary</strong>: <code>10:27:49: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1219)

<details><summary>context</summary>
<pre>10:27:48: exit status 1
10:27:49: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/repo/BUILD.bazel:3:11: @@gazelle&#43;//repo:repo depends on @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//:vcs in repository @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs which failed to fetch. no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout
10:27:49: exit status 1
10:27:49: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
10:27:49: INFO: Elapsed time: 75.163s, Critical Path: 0.61s
10:27:49: INFO: 1 process: 1 internal.
10:27:49: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>secondary</strong>: <code>10:27:49: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1222)

<details><summary>context</summary>
<pre>10:27:49: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
10:27:49: INFO: Elapsed time: 75.163s, Critical Path: 0.61s
10:27:49: INFO: 1 process: 1 internal.
10:27:49: ERROR: Build did NOT complete successfully
10:27:49: ERROR: Build failed. Not running target
10:27:51: &#43; rm -f /tmp/kubevirt.deploy.9ro5
make: *** [Makefile:194: cluster-sync] Error 1</pre>
</details>



* <strong>secondary</strong>: <code>10:27:49: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1223)

<details><summary>context</summary>
<pre>10:27:49: INFO: Elapsed time: 75.163s, Critical Path: 0.61s
10:27:49: INFO: 1 process: 1 internal.
10:27:49: ERROR: Build did NOT complete successfully
10:27:49: ERROR: Build failed. Not running target
10:27:51: &#43; rm -f /tmp/kubevirt.deploy.9ro5
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:194: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1225)

<details><summary>context</summary>
<pre>10:27:49: ERROR: Build did NOT complete successfully
10:27:49: ERROR: Build failed. Not running target
10:27:51: &#43; rm -f /tmp/kubevirt.deploy.9ro5
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1234)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
10:27:52: selecting podman as container runtime
10:28:35: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network is being used by the following container(s): 73f69f92b0ee12ba15fe06ff081855efc21f25241a9104c96303a06d01ed229d: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>

**1x**: _2026-11-21 03:28:48 &#43;0000 UTC_: <code>03:39:24: Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1542)
<details>
<summary>all...</summary>

* _2026-11-21 03:28:48 &#43;0000 UTC_: <code>03:39:24: Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1542)
<details><summary>context</summary>
<pre>03:39:22: localhost:39399/kubevirt/virt-template-apiserver:v0.3.0
03:39:22:     layer transfers: 0 from disk, 14 from container registry, 0 from remote cache
03:39:22: INFO: Invocation ID: eb79cc26-99a6-4945-aeae-b12f2226c7ff
03:39:24: Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:24: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: INFO: Repo &#43;virt_template_ext&#43;virt_template_controller defined by rule pull in @@rules_img&#43;//img/private/repository_rules:pull.bzl
03:39:25: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/download.bzl:441:13: An error occurred during the fetch of repository &#39;&#43;virt_template_ext&#43;virt_template_controller&#39;:</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>03:29:18: I1121 03:29:18.058174    2039 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A632)

<details><summary>context</summary>
<pre>03:29:17: [control-plane-check] kube-scheduler is healthy after 10.206452ms
03:29:17: [control-plane-check] kube-controller-manager is healthy after 12.069006ms
03:29:17: I1121 03:29:17.559246    2039 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: rbac: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
03:29:18: I1121 03:29:18.058174    2039 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
03:29:18: I1121 03:29:18.058228    2039 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
03:29:18: [control-plane-check] kube-apiserver is healthy after 1.502576755s
03:29:18: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>03:39:24: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1543)

<details><summary>context</summary>
<pre>03:39:22:     layer transfers: 0 from disk, 14 from container registry, 0 from remote cache
03:39:22: INFO: Invocation ID: eb79cc26-99a6-4945-aeae-b12f2226c7ff
03:39:24: Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:24: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: INFO: Repo &#43;virt_template_ext&#43;virt_template_controller defined by rule pull in @@rules_img&#43;//img/private/repository_rules:pull.bzl
03:39:25: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/download.bzl:441:13: An error occurred during the fetch of repository &#39;&#43;virt_template_ext&#43;virt_template_controller&#39;:
03:39:25:    Traceback (most recent call last):</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/download.bzl:441:13: An error occurred during the fetch of repository &#39;&#43;virt_template_ext&#43;virt_template_controller&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1545)

<details><summary>context</summary>
<pre>03:39:24: Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:24: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: INFO: Repo &#43;virt_template_ext&#43;virt_template_controller defined by rule pull in @@rules_img&#43;//img/private/repository_rules:pull.bzl
03:39:25: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/download.bzl:441:13: An error occurred during the fetch of repository &#39;&#43;virt_template_ext&#43;virt_template_controller&#39;:
03:39:25:    Traceback (most recent call last):
03:39:25: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/pull.bzl&#34;, line 48, column 28, in _pull_impl
03:39:25: 		_download_with_tool(</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: Error in fail: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1551)

<details><summary>context</summary>
<pre>03:39:25: 		_download_with_tool(
03:39:25: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/download.bzl&#34;, line 441, column 13, in download_with_tool
03:39:25: 		fail(&#34;img tool failed with exit code {} and message {}&#34;.format(result.return_code, result.stderr))
03:39:25: Error in fail: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1552)

<details><summary>context</summary>
<pre>03:39:25: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/download.bzl&#34;, line 441, column 13, in download_with_tool
03:39:25: 		fail(&#34;img tool failed with exit code {} and message {}&#34;.format(result.return_code, result.stderr))
03:39:25: Error in fail: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1553)

<details><summary>context</summary>
<pre>03:39:25: 		fail(&#34;img tool failed with exit code {} and message {}&#34;.format(result.return_code, result.stderr))
03:39:25: Error in fail: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1554)

<details><summary>context</summary>
<pre>03:39:25: Error in fail: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1555)

<details><summary>context</summary>
<pre>03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed
03:39:25: INFO: Elapsed time: 2.254s, Critical Path: 0.16s</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1556)

<details><summary>context</summary>
<pre>03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed
03:39:25: INFO: Elapsed time: 2.254s, Critical Path: 0.16s
03:39:25: INFO: 1 process: 1 internal.</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1557)

<details><summary>context</summary>
<pre>03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed
03:39:25: INFO: Elapsed time: 2.254s, Critical Path: 0.16s
03:39:25: INFO: 1 process: 1 internal.
03:39:25: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1560)

<details><summary>context</summary>
<pre>03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed
03:39:25: INFO: Elapsed time: 2.254s, Critical Path: 0.16s
03:39:25: INFO: 1 process: 1 internal.
03:39:25: ERROR: Build did NOT complete successfully
03:39:25: ERROR: Build failed. Not running target
03:39:26: &#43; rm -f /tmp/kubevirt.deploy.Yb3J
make: *** [Makefile:194: cluster-sync] Error 1</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1561)

<details><summary>context</summary>
<pre>03:39:25: INFO: Elapsed time: 2.254s, Critical Path: 0.16s
03:39:25: INFO: 1 process: 1 internal.
03:39:25: ERROR: Build did NOT complete successfully
03:39:25: ERROR: Build failed. Not running target
03:39:26: &#43; rm -f /tmp/kubevirt.deploy.Yb3J
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:194: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1563)

<details><summary>context</summary>
<pre>03:39:25: ERROR: Build did NOT complete successfully
03:39:25: ERROR: Build failed. Not running target
03:39:26: &#43; rm -f /tmp/kubevirt.deploy.Yb3J
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1572)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
03:39:26: selecting podman as container runtime
03:39:57: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.37-sig-storage is being used by the following container(s): 4db2389c3fdcc1a55455a1885b7cd013063cca177a8b8947fd884f25a3975020: volume is being used
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


### main (3x / 75.00%)


#### needs-investigation (2x / 66.67%)

<details>
<summary> no matching pattern (2x / 66.67%) </summary>

<hr/>

**1x**: _2026-11-23 10:23:49 &#43;0000 UTC_: <code>10:27:48: ERROR: no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1215)
<details>
<summary>all...</summary>

* _2026-11-23 10:23:49 &#43;0000 UTC_: <code>10:27:48: ERROR: no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1215)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>10:24:21: I1123 10:24:20.740427    2085 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A635)

<details><summary>context</summary>
<pre>10:24:20: [control-plane-check] kube-scheduler is healthy after 4.814576ms
10:24:20: [control-plane-check] kube-controller-manager is healthy after 8.72848ms
10:24:20: I1123 10:24:20.244157    2085 wait.go:283] kube-apiserver check failed at https://[fd00::101]:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: rbac: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
10:24:21: I1123 10:24:20.740427    2085 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
10:24:21: I1123 10:24:20.740464    2085 wait.go:283] kube-apiserver check failed at https://[fd00::101]:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
10:24:21: [control-plane-check] kube-apiserver is healthy after 1.502020037s
10:24:21: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>10:27:47: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/internal/go_repository.bzl:297:13: An error occurred during the fetch of repository &#39;gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1209)

<details><summary>context</summary>
<pre>10:27:47: INFO: Repo gazelle&#43;&#43;go_deps&#43;org_golang_x_mod defined by rule go_repository in @@gazelle&#43;//internal:go_repository.bzl
10:27:47: INFO: Repo gazelle&#43;&#43;go_deps&#43;com_github_bmatcuk_doublestar_v4 defined by rule go_repository in @@gazelle&#43;//internal:go_repository.bzl
10:27:47: INFO: Repo gazelle&#43;&#43;go_deps&#43;com_github_bazelbuild_buildtools defined by rule go_repository in @@gazelle&#43;//internal:go_repository.bzl
10:27:47: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/internal/go_repository.bzl:297:13: An error occurred during the fetch of repository &#39;gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs&#39;:
10:27:47:    Traceback (most recent call last):
10:27:47: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/internal/go_repository.bzl&#34;, line 297, column 13, in _go_repository_impl
10:27:47: 		fail(&#34;%s: %s&#34; % (ctx.name, result.stderr))</pre>
</details>



* <strong>secondary</strong>: <code>10:27:49: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/repo/BUILD.bazel:3:11: @@gazelle&#43;//repo:repo depends on @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//:vcs in repository @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs which failed to fetch. no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1217)

<details><summary>context</summary>
<pre>10:27:47: exit status 1
10:27:48: ERROR: no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout
10:27:48: exit status 1
10:27:49: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/repo/BUILD.bazel:3:11: @@gazelle&#43;//repo:repo depends on @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//:vcs in repository @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs which failed to fetch. no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout
10:27:49: exit status 1
10:27:49: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
10:27:49: INFO: Elapsed time: 75.163s, Critical Path: 0.61s</pre>
</details>



* <strong>secondary</strong>: <code>10:27:49: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1219)

<details><summary>context</summary>
<pre>10:27:48: exit status 1
10:27:49: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/repo/BUILD.bazel:3:11: @@gazelle&#43;//repo:repo depends on @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//:vcs in repository @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs which failed to fetch. no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout
10:27:49: exit status 1
10:27:49: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
10:27:49: INFO: Elapsed time: 75.163s, Critical Path: 0.61s
10:27:49: INFO: 1 process: 1 internal.
10:27:49: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>secondary</strong>: <code>10:27:49: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1222)

<details><summary>context</summary>
<pre>10:27:49: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
10:27:49: INFO: Elapsed time: 75.163s, Critical Path: 0.61s
10:27:49: INFO: 1 process: 1 internal.
10:27:49: ERROR: Build did NOT complete successfully
10:27:49: ERROR: Build failed. Not running target
10:27:51: &#43; rm -f /tmp/kubevirt.deploy.9ro5
make: *** [Makefile:194: cluster-sync] Error 1</pre>
</details>



* <strong>secondary</strong>: <code>10:27:49: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1223)

<details><summary>context</summary>
<pre>10:27:49: INFO: Elapsed time: 75.163s, Critical Path: 0.61s
10:27:49: INFO: 1 process: 1 internal.
10:27:49: ERROR: Build did NOT complete successfully
10:27:49: ERROR: Build failed. Not running target
10:27:51: &#43; rm -f /tmp/kubevirt.deploy.9ro5
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:194: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1225)

<details><summary>context</summary>
<pre>10:27:49: ERROR: Build did NOT complete successfully
10:27:49: ERROR: Build failed. Not running target
10:27:51: &#43; rm -f /tmp/kubevirt.deploy.9ro5
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1234)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
10:27:52: selecting podman as container runtime
10:28:35: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network is being used by the following container(s): 73f69f92b0ee12ba15fe06ff081855efc21f25241a9104c96303a06d01ed229d: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>

**1x**: _2026-11-21 03:28:48 &#43;0000 UTC_: <code>03:39:24: Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1542)
<details>
<summary>all...</summary>

* _2026-11-21 03:28:48 &#43;0000 UTC_: <code>03:39:24: Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1542)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>03:29:18: I1121 03:29:18.058174    2039 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A632)

<details><summary>context</summary>
<pre>03:29:17: [control-plane-check] kube-scheduler is healthy after 10.206452ms
03:29:17: [control-plane-check] kube-controller-manager is healthy after 12.069006ms
03:29:17: I1121 03:29:17.559246    2039 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: rbac: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
03:29:18: I1121 03:29:18.058174    2039 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
03:29:18: I1121 03:29:18.058228    2039 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
03:29:18: [control-plane-check] kube-apiserver is healthy after 1.502576755s
03:29:18: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>03:39:24: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1543)

<details><summary>context</summary>
<pre>03:39:22:     layer transfers: 0 from disk, 14 from container registry, 0 from remote cache
03:39:22: INFO: Invocation ID: eb79cc26-99a6-4945-aeae-b12f2226c7ff
03:39:24: Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:24: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: INFO: Repo &#43;virt_template_ext&#43;virt_template_controller defined by rule pull in @@rules_img&#43;//img/private/repository_rules:pull.bzl
03:39:25: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/download.bzl:441:13: An error occurred during the fetch of repository &#39;&#43;virt_template_ext&#43;virt_template_controller&#39;:
03:39:25:    Traceback (most recent call last):</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/download.bzl:441:13: An error occurred during the fetch of repository &#39;&#43;virt_template_ext&#43;virt_template_controller&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1545)

<details><summary>context</summary>
<pre>03:39:24: Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:24: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: INFO: Repo &#43;virt_template_ext&#43;virt_template_controller defined by rule pull in @@rules_img&#43;//img/private/repository_rules:pull.bzl
03:39:25: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/download.bzl:441:13: An error occurred during the fetch of repository &#39;&#43;virt_template_ext&#43;virt_template_controller&#39;:
03:39:25:    Traceback (most recent call last):
03:39:25: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/pull.bzl&#34;, line 48, column 28, in _pull_impl
03:39:25: 		_download_with_tool(</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: Error in fail: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1551)

<details><summary>context</summary>
<pre>03:39:25: 		_download_with_tool(
03:39:25: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/download.bzl&#34;, line 441, column 13, in download_with_tool
03:39:25: 		fail(&#34;img tool failed with exit code {} and message {}&#34;.format(result.return_code, result.stderr))
03:39:25: Error in fail: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1552)

<details><summary>context</summary>
<pre>03:39:25: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/download.bzl&#34;, line 441, column 13, in download_with_tool
03:39:25: 		fail(&#34;img tool failed with exit code {} and message {}&#34;.format(result.return_code, result.stderr))
03:39:25: Error in fail: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1553)

<details><summary>context</summary>
<pre>03:39:25: 		fail(&#34;img tool failed with exit code {} and message {}&#34;.format(result.return_code, result.stderr))
03:39:25: Error in fail: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1554)

<details><summary>context</summary>
<pre>03:39:25: Error in fail: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1555)

<details><summary>context</summary>
<pre>03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed
03:39:25: INFO: Elapsed time: 2.254s, Critical Path: 0.16s</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1556)

<details><summary>context</summary>
<pre>03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed
03:39:25: INFO: Elapsed time: 2.254s, Critical Path: 0.16s
03:39:25: INFO: 1 process: 1 internal.</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1557)

<details><summary>context</summary>
<pre>03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed
03:39:25: INFO: Elapsed time: 2.254s, Critical Path: 0.16s
03:39:25: INFO: 1 process: 1 internal.
03:39:25: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1560)

<details><summary>context</summary>
<pre>03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed
03:39:25: INFO: Elapsed time: 2.254s, Critical Path: 0.16s
03:39:25: INFO: 1 process: 1 internal.
03:39:25: ERROR: Build did NOT complete successfully
03:39:25: ERROR: Build failed. Not running target
03:39:26: &#43; rm -f /tmp/kubevirt.deploy.Yb3J
make: *** [Makefile:194: cluster-sync] Error 1</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1561)

<details><summary>context</summary>
<pre>03:39:25: INFO: Elapsed time: 2.254s, Critical Path: 0.16s
03:39:25: INFO: 1 process: 1 internal.
03:39:25: ERROR: Build did NOT complete successfully
03:39:25: ERROR: Build failed. Not running target
03:39:26: &#43; rm -f /tmp/kubevirt.deploy.Yb3J
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:194: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1563)

<details><summary>context</summary>
<pre>03:39:25: ERROR: Build did NOT complete successfully
03:39:25: ERROR: Build failed. Not running target
03:39:26: &#43; rm -f /tmp/kubevirt.deploy.Yb3J
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1572)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
03:39:26: selecting podman as container runtime
03:39:57: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.37-sig-storage is being used by the following container(s): 4db2389c3fdcc1a55455a1885b7cd013063cca177a8b8947fd884f25a3975020: volume is being used
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

**1x**: _2026-11-23 18:50:51 &#43;0000 UTC_: <code>18:52:09: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19479/pull-kubevirt-e2e-k8s-1.36-sig-network/2124925311807131648#1:build-log.txt%3A698)
<details>
<summary>all...</summary>

* _2026-11-23 18:50:51 &#43;0000 UTC_: <code>18:52:09: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19479/pull-kubevirt-e2e-k8s-1.36-sig-network/2124925311807131648#1:build-log.txt%3A698)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>18:52:07: I1123 13:52:06.874315    1962 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19479/pull-kubevirt-e2e-k8s-1.36-sig-network/2124925311807131648#1:build-log.txt%3A634)

<details><summary>context</summary>
<pre>18:52:06: [control-plane-check] kube-controller-manager is healthy after 8.514229ms
18:52:06: [control-plane-check] kube-scheduler is healthy after 9.410102ms
18:52:06: I1123 13:52:06.374210    1962 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
18:52:07: I1123 13:52:06.874315    1962 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
18:52:07: I1123 13:52:06.874358    1962 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
18:52:07: I1123 13:52:07.377532    1962 uploadconfig.go:109] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap
18:52:07: [control-plane-check] kube-apiserver is healthy after 1.504955242s</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19479/pull-kubevirt-e2e-k8s-1.36-sig-network/2124925311807131648#1:build-log.txt%3A1022)

<details><summary>context</summary>
<pre>18:52:09: 
18:52:09: &#43; wait 69
18:52:20: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>



</details>

<hr/>
</details>

### release-1.7 (1x / 25.00%)


#### internal (1x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-11-23 10:05:27 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/20015/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2124803835099615232#1:build-log.txt%3A1193)
<details>
<summary>all...</summary>

* _2026-11-23 10:05:27 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/20015/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2124803835099615232#1:build-log.txt%3A1193)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>10:19:51: I1123 05:19:51.532976    1591 request.go:1664] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/20015/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2124803835099615232#1:build-log.txt%3A763)

<details><summary>context</summary>
<pre>10:19:48: [control-plane-check] Checking kube-scheduler at https://127.0.0.1:10259/livez
10:19:50: [control-plane-check] kube-scheduler is healthy after 2.118400411s
10:19:50: [control-plane-check] kube-controller-manager is healthy after 2.603560978s
10:19:51: I1123 05:19:51.532976    1591 request.go:1664] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
10:19:52: [control-plane-check] kube-apiserver is healthy after 4.003281102s
10:19:52: I1123 05:19:52.038341    1591 kubeconfig.go:665] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists
10:19:52: I1123 05:19:52.041457    1591 kubeconfig.go:738] creating the ClusterRoleBinding for the kubeadm:cluster-admins Group by using super-admin.conf</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/20015/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2124803835099615232#1:build-log.txt%3A1199)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
10:25:11: selecting podman as container runtime
10:25:45: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7 is being used by the following container(s): 78694673e778d74e782e667f11be6db389a0f55e5c27c234fc87aebe8e9f674a, 05ff6ad545874ae0691c6a1bfba03f1a54dcd66dd4ef9c64d883795eafa6648b: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
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


### sig-network (2x / 50.00%)


#### internal (1x / 50.00%)

<details>
<summary> configured node not found during provisioning (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-11-23 18:50:51 &#43;0000 UTC_: <code>18:52:09: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19479/pull-kubevirt-e2e-k8s-1.36-sig-network/2124925311807131648#1:build-log.txt%3A698)
<details>
<summary>all...</summary>

* _2026-11-23 18:50:51 &#43;0000 UTC_: <code>18:52:09: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19479/pull-kubevirt-e2e-k8s-1.36-sig-network/2124925311807131648#1:build-log.txt%3A698)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>18:52:07: I1123 13:52:06.874315    1962 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19479/pull-kubevirt-e2e-k8s-1.36-sig-network/2124925311807131648#1:build-log.txt%3A634)

<details><summary>context</summary>
<pre>18:52:06: [control-plane-check] kube-controller-manager is healthy after 8.514229ms
18:52:06: [control-plane-check] kube-scheduler is healthy after 9.410102ms
18:52:06: I1123 13:52:06.374210    1962 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
18:52:07: I1123 13:52:06.874315    1962 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
18:52:07: I1123 13:52:06.874358    1962 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
18:52:07: I1123 13:52:07.377532    1962 uploadconfig.go:109] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap
18:52:07: [control-plane-check] kube-apiserver is healthy after 1.504955242s</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19479/pull-kubevirt-e2e-k8s-1.36-sig-network/2124925311807131648#1:build-log.txt%3A1022)

<details><summary>context</summary>
<pre>18:52:09: 
18:52:09: &#43; wait 69
18:52:20: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>



</details>

<hr/>
</details>

#### needs-investigation (1x / 50.00%)

<details>
<summary> no matching pattern (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-11-23 10:23:49 &#43;0000 UTC_: <code>10:27:48: ERROR: no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1215)
<details>
<summary>all...</summary>

* _2026-11-23 10:23:49 &#43;0000 UTC_: <code>10:27:48: ERROR: no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1215)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>10:24:21: I1123 10:24:20.740427    2085 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A635)

<details><summary>context</summary>
<pre>10:24:20: [control-plane-check] kube-scheduler is healthy after 4.814576ms
10:24:20: [control-plane-check] kube-controller-manager is healthy after 8.72848ms
10:24:20: I1123 10:24:20.244157    2085 wait.go:283] kube-apiserver check failed at https://[fd00::101]:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: rbac: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
10:24:21: I1123 10:24:20.740427    2085 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
10:24:21: I1123 10:24:20.740464    2085 wait.go:283] kube-apiserver check failed at https://[fd00::101]:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
10:24:21: [control-plane-check] kube-apiserver is healthy after 1.502020037s
10:24:21: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>10:27:47: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/internal/go_repository.bzl:297:13: An error occurred during the fetch of repository &#39;gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1209)

<details><summary>context</summary>
<pre>10:27:47: INFO: Repo gazelle&#43;&#43;go_deps&#43;org_golang_x_mod defined by rule go_repository in @@gazelle&#43;//internal:go_repository.bzl
10:27:47: INFO: Repo gazelle&#43;&#43;go_deps&#43;com_github_bmatcuk_doublestar_v4 defined by rule go_repository in @@gazelle&#43;//internal:go_repository.bzl
10:27:47: INFO: Repo gazelle&#43;&#43;go_deps&#43;com_github_bazelbuild_buildtools defined by rule go_repository in @@gazelle&#43;//internal:go_repository.bzl
10:27:47: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/internal/go_repository.bzl:297:13: An error occurred during the fetch of repository &#39;gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs&#39;:
10:27:47:    Traceback (most recent call last):
10:27:47: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/internal/go_repository.bzl&#34;, line 297, column 13, in _go_repository_impl
10:27:47: 		fail(&#34;%s: %s&#34; % (ctx.name, result.stderr))</pre>
</details>



* <strong>secondary</strong>: <code>10:27:49: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/repo/BUILD.bazel:3:11: @@gazelle&#43;//repo:repo depends on @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//:vcs in repository @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs which failed to fetch. no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1217)

<details><summary>context</summary>
<pre>10:27:47: exit status 1
10:27:48: ERROR: no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout
10:27:48: exit status 1
10:27:49: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/repo/BUILD.bazel:3:11: @@gazelle&#43;//repo:repo depends on @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//:vcs in repository @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs which failed to fetch. no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout
10:27:49: exit status 1
10:27:49: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
10:27:49: INFO: Elapsed time: 75.163s, Critical Path: 0.61s</pre>
</details>



* <strong>secondary</strong>: <code>10:27:49: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1219)

<details><summary>context</summary>
<pre>10:27:48: exit status 1
10:27:49: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/gazelle&#43;/repo/BUILD.bazel:3:11: @@gazelle&#43;//repo:repo depends on @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//:vcs in repository @@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs which failed to fetch. no such package &#39;@@gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs//&#39;: gazelle&#43;&#43;go_deps&#43;org_golang_x_tools_go_vcs: fetch_repo: golang.org/x/tools/go/vcs@v0.1.0-deprecated: Get &#34;https://proxy.golang.org/golang.org/x/tools/go/vcs/@v/v0.1.0-deprecated.info&#34;: net/http: TLS handshake timeout
10:27:49: exit status 1
10:27:49: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
10:27:49: INFO: Elapsed time: 75.163s, Critical Path: 0.61s
10:27:49: INFO: 1 process: 1 internal.
10:27:49: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>secondary</strong>: <code>10:27:49: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1222)

<details><summary>context</summary>
<pre>10:27:49: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
10:27:49: INFO: Elapsed time: 75.163s, Critical Path: 0.61s
10:27:49: INFO: 1 process: 1 internal.
10:27:49: ERROR: Build did NOT complete successfully
10:27:49: ERROR: Build failed. Not running target
10:27:51: &#43; rm -f /tmp/kubevirt.deploy.9ro5
make: *** [Makefile:194: cluster-sync] Error 1</pre>
</details>



* <strong>secondary</strong>: <code>10:27:49: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1223)

<details><summary>context</summary>
<pre>10:27:49: INFO: Elapsed time: 75.163s, Critical Path: 0.61s
10:27:49: INFO: 1 process: 1 internal.
10:27:49: ERROR: Build did NOT complete successfully
10:27:49: ERROR: Build failed. Not running target
10:27:51: &#43; rm -f /tmp/kubevirt.deploy.9ro5
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:194: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1225)

<details><summary>context</summary>
<pre>10:27:49: ERROR: Build did NOT complete successfully
10:27:49: ERROR: Build failed. Not running target
10:27:51: &#43; rm -f /tmp/kubevirt.deploy.9ro5
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19802/pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network/2124807956347228160#1:build-log.txt%3A1234)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
10:27:52: selecting podman as container runtime
10:28:35: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.37-ipv6-sig-network is being used by the following container(s): 73f69f92b0ee12ba15fe06ff081855efc21f25241a9104c96303a06d01ed229d: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

### sig-storage (2x / 50.00%)


#### internal (1x / 50.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-11-23 10:05:27 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/20015/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2124803835099615232#1:build-log.txt%3A1193)
<details>
<summary>all...</summary>

* _2026-11-23 10:05:27 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/20015/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2124803835099615232#1:build-log.txt%3A1193)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>10:19:51: I1123 05:19:51.532976    1591 request.go:1664] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/20015/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2124803835099615232#1:build-log.txt%3A763)

<details><summary>context</summary>
<pre>10:19:48: [control-plane-check] Checking kube-scheduler at https://127.0.0.1:10259/livez
10:19:50: [control-plane-check] kube-scheduler is healthy after 2.118400411s
10:19:50: [control-plane-check] kube-controller-manager is healthy after 2.603560978s
10:19:51: I1123 05:19:51.532976    1591 request.go:1664] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
10:19:52: [control-plane-check] kube-apiserver is healthy after 4.003281102s
10:19:52: I1123 05:19:52.038341    1591 kubeconfig.go:665] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists
10:19:52: I1123 05:19:52.041457    1591 kubeconfig.go:738] creating the ClusterRoleBinding for the kubeadm:cluster-admins Group by using super-admin.conf</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/20015/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7/2124803835099615232#1:build-log.txt%3A1199)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
10:25:11: selecting podman as container runtime
10:25:45: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.33-sig-storage-1.7 is being used by the following container(s): 78694673e778d74e782e667f11be6db389a0f55e5c27c234fc87aebe8e9f674a, 05ff6ad545874ae0691c6a1bfba03f1a54dcd66dd4ef9c64d883795eafa6648b: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
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
<summary> no matching pattern (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-11-21 03:28:48 &#43;0000 UTC_: <code>03:39:24: Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1542)
<details>
<summary>all...</summary>

* _2026-11-21 03:28:48 &#43;0000 UTC_: <code>03:39:24: Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1542)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>03:29:18: I1121 03:29:18.058174    2039 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A632)

<details><summary>context</summary>
<pre>03:29:17: [control-plane-check] kube-scheduler is healthy after 10.206452ms
03:29:17: [control-plane-check] kube-controller-manager is healthy after 12.069006ms
03:29:17: I1121 03:29:17.559246    2039 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: rbac: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
03:29:18: I1121 03:29:18.058174    2039 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
03:29:18: I1121 03:29:18.058228    2039 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
03:29:18: [control-plane-check] kube-apiserver is healthy after 1.502576755s
03:29:18: [upload-config] Storing the configuration used in ConfigMap &#34;kubeadm-config&#34; in the &#34;kube-system&#34; Namespace</pre>
</details>



* <strong>secondary</strong>: <code>03:39:24: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1543)

<details><summary>context</summary>
<pre>03:39:22:     layer transfers: 0 from disk, 14 from container registry, 0 from remote cache
03:39:22: INFO: Invocation ID: eb79cc26-99a6-4945-aeae-b12f2226c7ff
03:39:24: Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:24: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: INFO: Repo &#43;virt_template_ext&#43;virt_template_controller defined by rule pull in @@rules_img&#43;//img/private/repository_rules:pull.bzl
03:39:25: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/download.bzl:441:13: An error occurred during the fetch of repository &#39;&#43;virt_template_ext&#43;virt_template_controller&#39;:
03:39:25:    Traceback (most recent call last):</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/download.bzl:441:13: An error occurred during the fetch of repository &#39;&#43;virt_template_ext&#43;virt_template_controller&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1545)

<details><summary>context</summary>
<pre>03:39:24: Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:24: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: INFO: Repo &#43;virt_template_ext&#43;virt_template_controller defined by rule pull in @@rules_img&#43;//img/private/repository_rules:pull.bzl
03:39:25: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/download.bzl:441:13: An error occurred during the fetch of repository &#39;&#43;virt_template_ext&#43;virt_template_controller&#39;:
03:39:25:    Traceback (most recent call last):
03:39:25: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/pull.bzl&#34;, line 48, column 28, in _pull_impl
03:39:25: 		_download_with_tool(</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: Error in fail: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1551)

<details><summary>context</summary>
<pre>03:39:25: 		_download_with_tool(
03:39:25: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/download.bzl&#34;, line 441, column 13, in download_with_tool
03:39:25: 		fail(&#34;img tool failed with exit code {} and message {}&#34;.format(result.return_code, result.stderr))
03:39:25: Error in fail: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1552)

<details><summary>context</summary>
<pre>03:39:25: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_img&#43;/img/private/repository_rules/download.bzl&#34;, line 441, column 13, in download_with_tool
03:39:25: 		fail(&#34;img tool failed with exit code {} and message {}&#34;.format(result.return_code, result.stderr))
03:39:25: Error in fail: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1553)

<details><summary>context</summary>
<pre>03:39:25: 		fail(&#34;img tool failed with exit code {} and message {}&#34;.format(result.return_code, result.stderr))
03:39:25: Error in fail: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1554)

<details><summary>context</summary>
<pre>03:39:25: Error in fail: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1555)

<details><summary>context</summary>
<pre>03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed
03:39:25: INFO: Elapsed time: 2.254s, Critical Path: 0.16s</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1556)

<details><summary>context</summary>
<pre>03:39:25: ERROR: no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed
03:39:25: INFO: Elapsed time: 2.254s, Critical Path: 0.16s
03:39:25: INFO: 1 process: 1 internal.</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1557)

<details><summary>context</summary>
<pre>03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: /root/go/src/kubevirt.io/kubevirt/images/virt-template/BUILD.bazel:13:15: //images/virt-template:virt-template-controller-image depends on @@&#43;virt_template_ext&#43;virt_template_controller//:virt_template_controller in repository @@&#43;virt_template_ext&#43;virt_template_controller which failed to fetch. no such package &#39;@@&#43;virt_template_ext&#43;virt_template_controller//&#39;: img tool failed with exit code 1 and message Failed to download from quay.io: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: Error: Failed to download blob from all registries: downloading image: getting config digest: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/0a/0a17165576a1d92d85e4dfe3054c628730a4b0140735d4b990d405ef28ce5794?X-Amz-Algorithm=REDACTED&amp;X-Amz-Credential=REDACTED&amp;X-Amz-Date=REDACTED&amp;X-Amz-Expires=REDACTED&amp;X-Amz-Signature=REDACTED&amp;X-Amz-SignedHeaders=REDACTED&amp;akamai_signature=REDACTED&amp;namespace=REDACTED&amp;region=REDACTED&amp;repo_name=REDACTED&#34;: remote error: tls: internal error
03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed
03:39:25: INFO: Elapsed time: 2.254s, Critical Path: 0.16s
03:39:25: INFO: 1 process: 1 internal.
03:39:25: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1560)

<details><summary>context</summary>
<pre>03:39:25: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted: Analysis failed
03:39:25: INFO: Elapsed time: 2.254s, Critical Path: 0.16s
03:39:25: INFO: 1 process: 1 internal.
03:39:25: ERROR: Build did NOT complete successfully
03:39:25: ERROR: Build failed. Not running target
03:39:26: &#43; rm -f /tmp/kubevirt.deploy.Yb3J
make: *** [Makefile:194: cluster-sync] Error 1</pre>
</details>



* <strong>secondary</strong>: <code>03:39:25: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1561)

<details><summary>context</summary>
<pre>03:39:25: INFO: Elapsed time: 2.254s, Critical Path: 0.16s
03:39:25: INFO: 1 process: 1 internal.
03:39:25: ERROR: Build did NOT complete successfully
03:39:25: ERROR: Build failed. Not running target
03:39:26: &#43; rm -f /tmp/kubevirt.deploy.Yb3J
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:194: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1563)

<details><summary>context</summary>
<pre>03:39:25: ERROR: Build did NOT complete successfully
03:39:25: ERROR: Build failed. Not running target
03:39:26: &#43; rm -f /tmp/kubevirt.deploy.Yb3J
make: *** [Makefile:194: cluster-sync] Error 1
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19616/pull-kubevirt-e2e-k8s-1.37-sig-storage/2123981778073948160#1:build-log.txt%3A1572)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
03:39:26: selecting podman as container runtime
03:39:57: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.37-sig-storage is being used by the following container(s): 4db2389c3fdcc1a55455a1885b7cd013063cca177a8b8947fd884f25a3975020: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

Last updated: 2026-11-26 12:11:10
