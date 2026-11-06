





<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-11-06 (1x / 8.33%)


#### internal (1x / 100.00%)

<details>
<summary> configured node not found during provisioning (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-11-06 12:02:09 &#43;0000 UTC_: <code>12:18:05: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19813/pull-kubevirt-e2e-k8s-1.36-windows2016/2118675165452701696#1:build-log.txt%3A868)
<details>
<summary>all...</summary>

* _2026-11-06 12:02:09 &#43;0000 UTC_: <code>12:18:05: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19813/pull-kubevirt-e2e-k8s-1.36-windows2016/2118675165452701696#1:build-log.txt%3A868)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>12:18:02: I1106 07:18:02.507609    1888 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19813/pull-kubevirt-e2e-k8s-1.36-windows2016/2118675165452701696#1:build-log.txt%3A805)

<details><summary>context</summary>
<pre>12:18:01: [control-plane-check] kube-controller-manager is healthy after 5.239285ms
12:18:01: [control-plane-check] kube-scheduler is healthy after 7.898644ms
12:18:02: I1106 07:18:02.008369    1888 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
12:18:02: I1106 07:18:02.507609    1888 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:18:02: I1106 07:18:02.507711    1888 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:18:02: [control-plane-check] kube-apiserver is healthy after 1.503624779s
12:18:02: I1106 07:18:03.008716    1888 uploadconfig.go:109] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19813/pull-kubevirt-e2e-k8s-1.36-windows2016/2118675165452701696#1:build-log.txt%3A1150)

<details><summary>context</summary>
<pre>12:18:17: 
12:18:17: &#43; wait 69
12:18:29: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19813/pull-kubevirt-e2e-k8s-1.36-windows2016/2118675165452701696#1:build-log.txt%3A1178)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:18:30: selecting podman as container runtime
12:18:30: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-windows2016 is being used by the following container(s): afe66bfd0ea325b2aedd124aef268e5e6cb650d1d5ccd698ca0210fa9d12b1b3: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

### 2026-11-05 (4x / 33.33%)


#### external (3x / 75.00%)

<details>
<summary> download failure from external URL (2x / 50.00%) </summary>

<hr/>

**1x**: _2026-11-05 14:33:27 &#43;0000 UTC_: <code>14:36:41: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp6849394626262802922/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A276)
<details>
<summary>all...</summary>

* _2026-11-05 14:33:27 &#43;0000 UTC_: <code>14:36:41: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp6849394626262802922/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A276)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>14:34:53: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_cc&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A218)

<details><summary>context</summary>
<pre>14:34:53:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/io_bazel_rules_go/go/private/repositories.bzl:355:18: in _maybe
14:34:53: Repository rule http_archive defined at:
14:34:53:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
14:34:53: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_cc&#39;:
14:34:53:    Traceback (most recent call last):
14:34:53: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
14:34:53: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>14:34:53: ERROR: Skipping &#39;//:gazelle&#39;: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A224)

<details><summary>context</summary>
<pre>14:34:53: 		download_info = ctx.download_and_extract(
14:34:53: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: WARNING: Target pattern parsing failed.
14:34:53: ERROR: Skipping &#39;//:gazelle&#39;: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: ERROR: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: INFO: Elapsed time: 49.216s
14:34:53: INFO: 0 processes.</pre>
</details>



* <strong>diagnostic</strong>: <code>14:34:53: ERROR: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A225)

<details><summary>context</summary>
<pre>14:34:53: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: WARNING: Target pattern parsing failed.
14:34:53: ERROR: Skipping &#39;//:gazelle&#39;: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: ERROR: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: INFO: Elapsed time: 49.216s
14:34:53: INFO: 0 processes.
14:34:53: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>14:34:53: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A228)

<details><summary>context</summary>
<pre>14:34:53: ERROR: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: INFO: Elapsed time: 49.216s
14:34:53: INFO: 0 processes.
14:34:53: ERROR: Build did NOT complete successfully
14:34:53: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>14:34:53: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A229)

<details><summary>context</summary>
<pre>14:34:53: INFO: Elapsed time: 49.216s
14:34:53: INFO: 0 processes.
14:34:53: ERROR: Build did NOT complete successfully
14:34:53: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A230)

<details><summary>context</summary>
<pre>14:34:53: INFO: 0 processes.
14:34:53: ERROR: Build did NOT complete successfully
14:34:53: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>diagnostic</strong>: <code>14:36:03: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;apple_support~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A245)

<details><summary>context</summary>
<pre>14:36:03:   &lt;builtin&gt;: in &lt;toplevel&gt;
14:36:03: Repository rule http_archive defined at:
14:36:03:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
14:36:03: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;apple_support~&#39;:
14:36:03:    Traceback (most recent call last):
14:36:03: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
14:36:03: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>14:36:03: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp1387829800134064954/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A250)

<details><summary>context</summary>
<pre>14:36:03: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
14:36:03: 		download_info = ctx.download_and_extract(
14:36:03: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp1387829800134064954/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:03: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp1387829800134064954/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:03: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
14:36:03: INFO: Elapsed time: 48.439s, Critical Path: 0.60s
14:36:03: INFO: 1 process: 1 internal.</pre>
</details>



* <strong>diagnostic</strong>: <code>14:36:03: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A251)

<details><summary>context</summary>
<pre>14:36:03: 		download_info = ctx.download_and_extract(
14:36:03: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp1387829800134064954/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:03: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp1387829800134064954/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:03: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
14:36:03: INFO: Elapsed time: 48.439s, Critical Path: 0.60s
14:36:03: INFO: 1 process: 1 internal.
14:36:03: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>14:36:03: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A254)

<details><summary>context</summary>
<pre>14:36:03: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
14:36:03: INFO: Elapsed time: 48.439s, Critical Path: 0.60s
14:36:03: INFO: 1 process: 1 internal.
14:36:03: ERROR: Build did NOT complete successfully
14:36:03: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>14:36:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A255)

<details><summary>context</summary>
<pre>14:36:03: INFO: Elapsed time: 48.439s, Critical Path: 0.60s
14:36:03: INFO: 1 process: 1 internal.
14:36:03: ERROR: Build did NOT complete successfully
14:36:03: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A256)

<details><summary>context</summary>
<pre>14:36:03: INFO: 1 process: 1 internal.
14:36:03: ERROR: Build did NOT complete successfully
14:36:03: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>secondary</strong>: <code>14:36:41: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;apple_support~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A271)

<details><summary>context</summary>
<pre>14:36:41:   &lt;builtin&gt;: in &lt;toplevel&gt;
14:36:41: Repository rule http_archive defined at:
14:36:41:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
14:36:41: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;apple_support~&#39;:
14:36:41:    Traceback (most recent call last):
14:36:41: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
14:36:41: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>secondary</strong>: <code>14:36:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A277)

<details><summary>context</summary>
<pre>14:36:41: 		download_info = ctx.download_and_extract(
14:36:41: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp6849394626262802922/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:41: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp6849394626262802922/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
14:36:41: INFO: Elapsed time: 13.016s, Critical Path: 0.60s
14:36:41: INFO: 1 process: 1 internal.
14:36:41: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>secondary</strong>: <code>14:36:41: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A280)

<details><summary>context</summary>
<pre>14:36:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
14:36:41: INFO: Elapsed time: 13.016s, Critical Path: 0.60s
14:36:41: INFO: 1 process: 1 internal.
14:36:41: ERROR: Build did NOT complete successfully
14:36:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>secondary</strong>: <code>14:36:41: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A281)

<details><summary>context</summary>
<pre>14:36:41: INFO: Elapsed time: 13.016s, Critical Path: 0.60s
14:36:41: INFO: 1 process: 1 internal.
14:36:41: ERROR: Build did NOT complete successfully
14:36:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A282)

<details><summary>context</summary>
<pre>14:36:41: INFO: 1 process: 1 internal.
14:36:41: ERROR: Build did NOT complete successfully
14:36:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A293)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
14:36:47: selecting podman as container runtime
14:36:47: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-operator is being used by the following container(s): 9311f6c2ebab024ca15294aafe5d852280dc712623b4835614575fb0de781fe0: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>

**1x**: _2026-11-05 08:58:16 &#43;0000 UTC_: <code>09:01:23: ERROR: no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A276)
<details>
<summary>all...</summary>

* _2026-11-05 08:58:16 &#43;0000 UTC_: <code>09:01:23: ERROR: no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A276)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>08:59:46: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A221)

<details><summary>context</summary>
<pre>08:59:46:   &lt;builtin&gt;: in &lt;toplevel&gt;
08:59:46: Repository rule http_archive defined at:
08:59:46:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
08:59:46: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:
08:59:46:    Traceback (most recent call last):
08:59:46: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:46: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:46: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp18112639366700063061/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A226)

<details><summary>context</summary>
<pre>08:59:46: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:46: 		download_info = ctx.download_and_extract(
08:59:46: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp18112639366700063061/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:46: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp18112639366700063061/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:46: INFO: Elapsed time: 38.429s, Critical Path: 0.60s
08:59:46: INFO: 1 process: 1 internal.
08:59:46: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:46: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A229)

<details><summary>context</summary>
<pre>08:59:46: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp18112639366700063061/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:46: INFO: Elapsed time: 38.429s, Critical Path: 0.60s
08:59:46: INFO: 1 process: 1 internal.
08:59:46: ERROR: Build did NOT complete successfully
08:59:46: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:46: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A230)

<details><summary>context</summary>
<pre>08:59:46: INFO: Elapsed time: 38.429s, Critical Path: 0.60s
08:59:46: INFO: 1 process: 1 internal.
08:59:46: ERROR: Build did NOT complete successfully
08:59:46: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A231)

<details><summary>context</summary>
<pre>08:59:46: INFO: 1 process: 1 internal.
08:59:46: ERROR: Build did NOT complete successfully
08:59:46: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>diagnostic</strong>: <code>09:00:04: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A246)

<details><summary>context</summary>
<pre>09:00:04:   &lt;builtin&gt;: in &lt;toplevel&gt;
09:00:04: Repository rule http_archive defined at:
09:00:04:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
09:00:04: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:
09:00:04:    Traceback (most recent call last):
09:00:04: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
09:00:04: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>09:00:04: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp2637779128729034844/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A251)

<details><summary>context</summary>
<pre>09:00:04: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
09:00:04: 		download_info = ctx.download_and_extract(
09:00:04: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp2637779128729034844/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:04: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp2637779128729034844/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:04: INFO: Elapsed time: 13.746s, Critical Path: 0.53s
09:00:04: INFO: 1 process: 1 internal.
09:00:04: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>09:00:04: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A254)

<details><summary>context</summary>
<pre>09:00:04: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp2637779128729034844/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:04: INFO: Elapsed time: 13.746s, Critical Path: 0.53s
09:00:04: INFO: 1 process: 1 internal.
09:00:04: ERROR: Build did NOT complete successfully
09:00:04: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>09:00:04: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A255)

<details><summary>context</summary>
<pre>09:00:04: INFO: Elapsed time: 13.746s, Critical Path: 0.53s
09:00:04: INFO: 1 process: 1 internal.
09:00:04: ERROR: Build did NOT complete successfully
09:00:04: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A256)

<details><summary>context</summary>
<pre>09:00:04: INFO: 1 process: 1 internal.
09:00:04: ERROR: Build did NOT complete successfully
09:00:04: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_python~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A271)

<details><summary>context</summary>
<pre>09:01:23:   &lt;builtin&gt;: in &lt;toplevel&gt;
09:01:23: Repository rule http_archive defined at:
09:01:23:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_python~&#39;:
09:01:23:    Traceback (most recent call last):
09:01:23: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
09:01:23: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/toolchains/BUILD:274:27: @@rules_java~//toolchains:toolchain_java10_definition depends on @@bazel_tools//tools/jdk:toolchain_type in repository @@bazel_tools which failed to fetch. no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A277)

<details><summary>context</summary>
<pre>09:01:23: 		download_info = ctx.download_and_extract(
09:01:23: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/toolchains/BUILD:274:27: @@rules_java~//toolchains:toolchain_java10_definition depends on @@bazel_tools//tools/jdk:toolchain_type in repository @@bazel_tools which failed to fetch. no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/src/tools/launcher/BUILD:10:14: Analysis failed
09:01:23: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
09:01:23: INFO: Elapsed time: 67.819s, Critical Path: 0.61s</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/src/tools/launcher/BUILD:10:14: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A278)

<details><summary>context</summary>
<pre>09:01:23: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/toolchains/BUILD:274:27: @@rules_java~//toolchains:toolchain_java10_definition depends on @@bazel_tools//tools/jdk:toolchain_type in repository @@bazel_tools which failed to fetch. no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/src/tools/launcher/BUILD:10:14: Analysis failed
09:01:23: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
09:01:23: INFO: Elapsed time: 67.819s, Critical Path: 0.61s
09:01:23: INFO: 1 process: 1 internal.</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A279)

<details><summary>context</summary>
<pre>09:01:23: ERROR: no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/toolchains/BUILD:274:27: @@rules_java~//toolchains:toolchain_java10_definition depends on @@bazel_tools//tools/jdk:toolchain_type in repository @@bazel_tools which failed to fetch. no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/src/tools/launcher/BUILD:10:14: Analysis failed
09:01:23: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
09:01:23: INFO: Elapsed time: 67.819s, Critical Path: 0.61s
09:01:23: INFO: 1 process: 1 internal.
09:01:23: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A282)

<details><summary>context</summary>
<pre>09:01:23: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
09:01:23: INFO: Elapsed time: 67.819s, Critical Path: 0.61s
09:01:23: INFO: 1 process: 1 internal.
09:01:23: ERROR: Build did NOT complete successfully
09:01:23: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A283)

<details><summary>context</summary>
<pre>09:01:23: INFO: Elapsed time: 67.819s, Critical Path: 0.61s
09:01:23: INFO: 1 process: 1 internal.
09:01:23: ERROR: Build did NOT complete successfully
09:01:23: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A284)

<details><summary>context</summary>
<pre>09:01:23: INFO: 1 process: 1 internal.
09:01:23: ERROR: Build did NOT complete successfully
09:01:23: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A295)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
09:02:10: selecting podman as container runtime
09:02:11: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-operator is being used by the following container(s): 8461bb1b9368374a39c102452c3ac4240460428f7aab3a0411def0f08de4de90: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>
<details>
<summary> bazel repository fetch failure (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-11-05 08:58:00 &#43;0000 UTC_: <code>09:00:17: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A266)
<details>
<summary>all...</summary>

* _2026-11-05 08:58:00 &#43;0000 UTC_: <code>09:00:17: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A266)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>08:59:40: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A216)

<details><summary>context</summary>
<pre>08:59:40:   &lt;builtin&gt;: in &lt;toplevel&gt;
08:59:40: Repository rule http_archive defined at:
08:59:40:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
08:59:40: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:
08:59:40:    Traceback (most recent call last):
08:59:40: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:40: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp10655117106702680376/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A221)

<details><summary>context</summary>
<pre>08:59:40: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:40: 		download_info = ctx.download_and_extract(
08:59:40: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp10655117106702680376/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp10655117106702680376/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:41: INFO: Elapsed time: 57.381s, Critical Path: 0.64s
08:59:41: INFO: 1 process: 1 internal.
08:59:41: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:41: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A224)

<details><summary>context</summary>
<pre>08:59:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp10655117106702680376/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:41: INFO: Elapsed time: 57.381s, Critical Path: 0.64s
08:59:41: INFO: 1 process: 1 internal.
08:59:41: ERROR: Build did NOT complete successfully
08:59:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:41: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A225)

<details><summary>context</summary>
<pre>08:59:41: INFO: Elapsed time: 57.381s, Critical Path: 0.64s
08:59:41: INFO: 1 process: 1 internal.
08:59:41: ERROR: Build did NOT complete successfully
08:59:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A226)

<details><summary>context</summary>
<pre>08:59:41: INFO: 1 process: 1 internal.
08:59:41: ERROR: Build did NOT complete successfully
08:59:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:59: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A241)

<details><summary>context</summary>
<pre>08:59:59:   &lt;builtin&gt;: in &lt;toplevel&gt;
08:59:59: Repository rule http_archive defined at:
08:59:59:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
08:59:59: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:
08:59:59:    Traceback (most recent call last):
08:59:59: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:59: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:59: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp1347094595764469151/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A246)

<details><summary>context</summary>
<pre>08:59:59: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:59: 		download_info = ctx.download_and_extract(
08:59:59: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp1347094595764469151/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:59: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp1347094595764469151/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:59: INFO: Elapsed time: 13.436s, Critical Path: 0.58s
08:59:59: INFO: 1 process: 1 internal.
08:59:59: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:59: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A249)

<details><summary>context</summary>
<pre>08:59:59: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp1347094595764469151/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:59: INFO: Elapsed time: 13.436s, Critical Path: 0.58s
08:59:59: INFO: 1 process: 1 internal.
08:59:59: ERROR: Build did NOT complete successfully
08:59:59: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:59: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A250)

<details><summary>context</summary>
<pre>08:59:59: INFO: Elapsed time: 13.436s, Critical Path: 0.58s
08:59:59: INFO: 1 process: 1 internal.
08:59:59: ERROR: Build did NOT complete successfully
08:59:59: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A251)

<details><summary>context</summary>
<pre>08:59:59: INFO: 1 process: 1 internal.
08:59:59: ERROR: Build did NOT complete successfully
08:59:59: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>secondary</strong>: <code>09:00:17: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp7416242368100028705/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A271)

<details><summary>context</summary>
<pre>09:00:17: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
09:00:17: 		download_info = ctx.download_and_extract(
09:00:17: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp7416242368100028705/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:17: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp7416242368100028705/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:17: INFO: Elapsed time: 13.699s, Critical Path: 0.65s
09:00:17: INFO: 1 process: 1 internal.
09:00:17: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>secondary</strong>: <code>09:00:17: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A274)

<details><summary>context</summary>
<pre>09:00:17: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp7416242368100028705/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:17: INFO: Elapsed time: 13.699s, Critical Path: 0.65s
09:00:17: INFO: 1 process: 1 internal.
09:00:17: ERROR: Build did NOT complete successfully
09:00:17: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>secondary</strong>: <code>09:00:17: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A275)

<details><summary>context</summary>
<pre>09:00:17: INFO: Elapsed time: 13.699s, Critical Path: 0.65s
09:00:17: INFO: 1 process: 1 internal.
09:00:17: ERROR: Build did NOT complete successfully
09:00:17: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A276)

<details><summary>context</summary>
<pre>09:00:17: INFO: 1 process: 1 internal.
09:00:17: ERROR: Build did NOT complete successfully
09:00:17: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A287)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
09:00:26: selecting podman as container runtime
09:00:26: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute is being used by the following container(s): c8d4c791d2f91c3a56c506d6b5653b81a2c82208835724ffd468f4be28ac67c2: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

#### internal (1x / 25.00%)

<details>
<summary> taint not found during cluster setup (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-11-05 08:57:54 &#43;0000 UTC_: <code>09:06:39: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A1209)
<details>
<summary>all...</summary>

* _2026-11-05 08:57:54 &#43;0000 UTC_: <code>09:06:39: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A1209)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>09:06:14: I1105 09:06:12.855745     276 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A827)

<details><summary>context</summary>
<pre>09:06:14: I1105 09:06:12.355808     276 round_trippers.go:632] &#34;Response&#34; verb=&#34;GET&#34; url=&#34;https://10.89.0.2:6443/livez?timeout=10s&#34; status=&#34;403 Forbidden&#34; milliseconds=0
09:06:14: I1105 09:06:12.356102     276 wait.go:283] kube-apiserver check failed at https://10.89.0.2:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: rbac: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
09:06:14: I1105 09:06:12.855440     276 round_trippers.go:632] &#34;Response&#34; verb=&#34;GET&#34; url=&#34;https://10.89.0.2:6443/livez?timeout=10s&#34; status=&#34;500 Internal Server Error&#34; milliseconds=0
09:06:14: I1105 09:06:12.855745     276 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
09:06:14: I1105 09:06:12.855795     276 wait.go:283] kube-apiserver check failed at https://10.89.0.2:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
09:06:14: I1105 09:06:13.356423     276 round_trippers.go:632] &#34;Response&#34; verb=&#34;GET&#34; url=&#34;https://10.89.0.2:6443/livez?timeout=10s&#34; status=&#34;200 OK&#34; milliseconds=1
09:06:14: [control-plane-check] kube-apiserver is healthy after 1.502411537s</pre>
</details>



* <strong>diagnostic</strong>: <code>09:06:24: I1105 09:06:16.972362     282 token.go:249] [discovery] Retrying due to error: could not find a JWS signature in the cluster-info ConfigMap for token ID &#34;abcdef&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A1081)

<details><summary>context</summary>
<pre>09:06:24: I1105 09:06:16.963598     282 envvar.go:195] &#34;Feature gate default state&#34; feature=&#34;UnlockWhileProcessingFIFO&#34; enabled=true
09:06:24: I1105 09:06:16.963944     282 token.go:229] [discovery] Waiting for the cluster-info ConfigMap to receive a JWS signature for token ID &#34;abcdef&#34;
09:06:24: I1105 09:06:16.971994     282 round_trippers.go:632] &#34;Response&#34; verb=&#34;GET&#34; url=&#34;https://kind-1.37-control-plane:6443/api/v1/namespaces/kube-public/configmaps/cluster-info?timeout=10s&#34; status=&#34;200 OK&#34; milliseconds=7
09:06:24: I1105 09:06:16.972362     282 token.go:249] [discovery] Retrying due to error: could not find a JWS signature in the cluster-info ConfigMap for token ID &#34;abcdef&#34;
09:06:24: I1105 09:06:21.964117     282 token.go:229] [discovery] Waiting for the cluster-info ConfigMap to receive a JWS signature for token ID &#34;abcdef&#34;
09:06:24: I1105 09:06:21.968475     282 round_trippers.go:632] &#34;Response&#34; verb=&#34;GET&#34; url=&#34;https://kind-1.37-control-plane:6443/api/v1/namespaces/kube-public/configmaps/cluster-info?timeout=10s&#34; status=&#34;200 OK&#34; milliseconds=4
09:06:24: I1105 09:06:21.968916     282 token.go:113] [discovery] Cluster info signature and contents are valid and no TLS pinning was specified, will use API Server &#34;kind-1.37-control-plane:6443&#34;</pre>
</details>



* <strong>secondary</strong>: <code>09:07:09: error: unable to read URL &#34;https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml&#34;, server reported 504 Gateway Timeout, status code=504</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A1349)

<details><summary>context</summary>
<pre>09:06:56: deployment.apps/cdi-operator created
09:06:56: &#43;&#43; _kubectl create -f https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml
09:06:56: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.37/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.37/.kubeconfig create -f https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml
09:07:09: error: unable to read URL &#34;https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml&#34;, server reported 504 Gateway Timeout, status code=504
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A1350)

<details><summary>context</summary>
<pre>09:06:56: &#43;&#43; _kubectl create -f https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml
09:06:56: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.37/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.37/.kubeconfig create -f https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml
09:07:09: error: unable to read URL &#34;https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml&#34;, server reported 504 Gateway Timeout, status code=504
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>



</details>

<hr/>
</details>

### 2026-11-04 (1x / 8.33%)


#### needs-investigation (1x / 100.00%)

<details>
<summary> no error snippets (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-11-04 07:13:51 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19748/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2117877790450651136)
<details>
<summary>all...</summary>

* _2026-11-04 07:13:51 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19748/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2117877790450651136)




</details>

<hr/>
</details>

### 2026-11-03 (4x / 33.33%)


#### internal (4x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (3x / 75.00%) </summary>

<hr/>

**3x**: _2026-11-03 11:59:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2117586999056535552#1:build-log.txt%3A1217)
<details>
<summary>all...</summary>

* _2026-11-03 12:09:09 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2117587001577312256#1:build-log.txt%3A4051)


<details>
<summary>secondary evidence</summary>


* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2117587001577312256#1:build-log.txt%3A4057)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:30:29: selecting podman as container runtime
12:32:15: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.33-sig-network-1.7 is being used by the following container(s): 437e990e2003afcfa461cefe139c5de3abcd06e6208ef33c9e77cb53b350a614, d7746eb1230afd44770adaf9ee2a421590a3d28144b9974b1f14b2ea60f6311e: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



* _2026-11-03 11:59:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2117586999056535552#1:build-log.txt%3A1217)


<details>
<summary>secondary evidence</summary>


* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2117586999056535552#1:build-log.txt%3A1223)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:36:37: selecting podman as container runtime
12:36:52: Error response from daemon: cannot remove container 09207706028f593ad43cb2e4bf5ea652f5970c6efdf3a476d3f2ee015494902d as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



* _2026-11-03 11:58:41 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A3943)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>12:15:30: I1103 07:15:30.373979    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A798)

<details><summary>context</summary>
<pre>12:15:21: [control-plane-check] Checking kube-scheduler at https://127.0.0.1:10259/livez
12:15:25: [control-plane-check] kube-scheduler is healthy after 3.842462827s
12:15:26: [control-plane-check] kube-controller-manager is healthy after 5.068241346s
12:15:30: I1103 07:15:30.373979    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:15:30: [control-plane-check] kube-apiserver is healthy after 9.002895106s
12:15:30: I1103 07:15:30.875756    1605 kubeconfig.go:657] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists
12:15:30: I1103 07:15:30.878733    1605 kubeconfig.go:730] creating the ClusterRoleBinding for the kubeadm:cluster-admins Group by using super-admin.conf</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A3949)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:27:34: selecting podman as container runtime
12:27:49: Error response from daemon: cannot remove container 4d2bdd9539b64bcf49fcea3d1d98ac72c4d24a6d0e4c0180ba3525a5a560ec7c as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>
<details>
<summary> taint not found during cluster setup (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-11-03 11:58:13 &#43;0000 UTC_: <code>12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1231)
<details>
<summary>all...</summary>

* _2026-11-03 11:58:13 &#43;0000 UTC_: <code>12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1231)


<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>12:24:20: error: taint &#34;node-role.kubernetes.io/control-plane:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1234)

<details><summary>context</summary>
<pre>12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found
12:24:20: &#43;&#43; _kubectl taint nodes vgpu-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
12:24:20: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubeconfig taint nodes vgpu-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
12:24:20: error: taint &#34;node-role.kubernetes.io/control-plane:NoSchedule&#34; not found
12:24:20: &#43;&#43;&#43; grep -i worker
12:24:20: &#43;&#43;&#43; _get_nodes
12:24:20: &#43;&#43;&#43; awk &#39;{print $1}&#39;</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1541)

<details><summary>context</summary>
<pre>12:29:21: INFO: 2 processes: 1 internal, 1 processwrapper-sandbox.
12:29:21: INFO: Running command line: bazel-bin/example-guest-agent-copier /root/go/src/kubevirt.io/kubevirt/_out/cmd/example-guest-agent/example-guest-agent
12:29:49: &#43; rm -f /tmp/kubevirt.deploy.oDzW
make: *** [Makefile:174: cluster-sync] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1550)

<details><summary>context</summary>
<pre>12:29:50: using podman due to KIND_EXPERIMENTAL_PROVIDER
12:29:50: enabling experimental podman provider
12:29:50: Deleting cluster &#34;vgpu&#34; ...
12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125
12:30:05:
12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1553)

<details><summary>context</summary>
<pre>12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125
12:30:05:
12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1555)

<details><summary>context</summary>
<pre>12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125
&#43; true
&#43; exit 2</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1556)

<details><summary>context</summary>
<pre>12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1568)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container kind_ardinghelli in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1569)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container kind_ardinghelli in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1570)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell
================================================================================</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1571)

<details><summary>context</summary>
<pre>Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell
================================================================================
Done cleaning up after podman in container.</pre>
</details>


</details>



</details>

<hr/>
</details>

### 2026-11-02 (2x / 16.67%)


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

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (4x / 33.33%)

<details>
<summary> download failure from external URL (2x / 16.67%) </summary>

<hr/>

**1x**: _2026-11-05 14:33:27 &#43;0000 UTC_: <code>14:36:41: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp6849394626262802922/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A276)
<details>
<summary>all...</summary>

* _2026-11-05 14:33:27 &#43;0000 UTC_: <code>14:36:41: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp6849394626262802922/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A276)
<details><summary>context</summary>
<pre>14:36:41: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
14:36:41: 		download_info = ctx.download_and_extract(
14:36:41: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp6849394626262802922/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:41: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp6849394626262802922/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
14:36:41: INFO: Elapsed time: 13.016s, Critical Path: 0.60s
14:36:41: INFO: 1 process: 1 internal.</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>14:34:53: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_cc&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A218)

<details><summary>context</summary>
<pre>14:34:53:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/io_bazel_rules_go/go/private/repositories.bzl:355:18: in _maybe
14:34:53: Repository rule http_archive defined at:
14:34:53:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
14:34:53: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_cc&#39;:
14:34:53:    Traceback (most recent call last):
14:34:53: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
14:34:53: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>14:34:53: ERROR: Skipping &#39;//:gazelle&#39;: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A224)

<details><summary>context</summary>
<pre>14:34:53: 		download_info = ctx.download_and_extract(
14:34:53: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: WARNING: Target pattern parsing failed.
14:34:53: ERROR: Skipping &#39;//:gazelle&#39;: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: ERROR: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: INFO: Elapsed time: 49.216s
14:34:53: INFO: 0 processes.</pre>
</details>



* <strong>diagnostic</strong>: <code>14:34:53: ERROR: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A225)

<details><summary>context</summary>
<pre>14:34:53: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: WARNING: Target pattern parsing failed.
14:34:53: ERROR: Skipping &#39;//:gazelle&#39;: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: ERROR: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: INFO: Elapsed time: 49.216s
14:34:53: INFO: 0 processes.
14:34:53: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>14:34:53: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A228)

<details><summary>context</summary>
<pre>14:34:53: ERROR: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: INFO: Elapsed time: 49.216s
14:34:53: INFO: 0 processes.
14:34:53: ERROR: Build did NOT complete successfully
14:34:53: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>14:34:53: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A229)

<details><summary>context</summary>
<pre>14:34:53: INFO: Elapsed time: 49.216s
14:34:53: INFO: 0 processes.
14:34:53: ERROR: Build did NOT complete successfully
14:34:53: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A230)

<details><summary>context</summary>
<pre>14:34:53: INFO: 0 processes.
14:34:53: ERROR: Build did NOT complete successfully
14:34:53: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>diagnostic</strong>: <code>14:36:03: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;apple_support~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A245)

<details><summary>context</summary>
<pre>14:36:03:   &lt;builtin&gt;: in &lt;toplevel&gt;
14:36:03: Repository rule http_archive defined at:
14:36:03:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
14:36:03: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;apple_support~&#39;:
14:36:03:    Traceback (most recent call last):
14:36:03: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
14:36:03: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>14:36:03: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp1387829800134064954/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A250)

<details><summary>context</summary>
<pre>14:36:03: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
14:36:03: 		download_info = ctx.download_and_extract(
14:36:03: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp1387829800134064954/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:03: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp1387829800134064954/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:03: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
14:36:03: INFO: Elapsed time: 48.439s, Critical Path: 0.60s
14:36:03: INFO: 1 process: 1 internal.</pre>
</details>



* <strong>diagnostic</strong>: <code>14:36:03: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A251)

<details><summary>context</summary>
<pre>14:36:03: 		download_info = ctx.download_and_extract(
14:36:03: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp1387829800134064954/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:03: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp1387829800134064954/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:03: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
14:36:03: INFO: Elapsed time: 48.439s, Critical Path: 0.60s
14:36:03: INFO: 1 process: 1 internal.
14:36:03: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>14:36:03: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A254)

<details><summary>context</summary>
<pre>14:36:03: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
14:36:03: INFO: Elapsed time: 48.439s, Critical Path: 0.60s
14:36:03: INFO: 1 process: 1 internal.
14:36:03: ERROR: Build did NOT complete successfully
14:36:03: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>14:36:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A255)

<details><summary>context</summary>
<pre>14:36:03: INFO: Elapsed time: 48.439s, Critical Path: 0.60s
14:36:03: INFO: 1 process: 1 internal.
14:36:03: ERROR: Build did NOT complete successfully
14:36:03: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A256)

<details><summary>context</summary>
<pre>14:36:03: INFO: 1 process: 1 internal.
14:36:03: ERROR: Build did NOT complete successfully
14:36:03: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>secondary</strong>: <code>14:36:41: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;apple_support~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A271)

<details><summary>context</summary>
<pre>14:36:41:   &lt;builtin&gt;: in &lt;toplevel&gt;
14:36:41: Repository rule http_archive defined at:
14:36:41:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
14:36:41: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;apple_support~&#39;:
14:36:41:    Traceback (most recent call last):
14:36:41: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
14:36:41: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>secondary</strong>: <code>14:36:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A277)

<details><summary>context</summary>
<pre>14:36:41: 		download_info = ctx.download_and_extract(
14:36:41: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp6849394626262802922/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:41: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp6849394626262802922/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
14:36:41: INFO: Elapsed time: 13.016s, Critical Path: 0.60s
14:36:41: INFO: 1 process: 1 internal.
14:36:41: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>secondary</strong>: <code>14:36:41: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A280)

<details><summary>context</summary>
<pre>14:36:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
14:36:41: INFO: Elapsed time: 13.016s, Critical Path: 0.60s
14:36:41: INFO: 1 process: 1 internal.
14:36:41: ERROR: Build did NOT complete successfully
14:36:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>secondary</strong>: <code>14:36:41: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A281)

<details><summary>context</summary>
<pre>14:36:41: INFO: Elapsed time: 13.016s, Critical Path: 0.60s
14:36:41: INFO: 1 process: 1 internal.
14:36:41: ERROR: Build did NOT complete successfully
14:36:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A282)

<details><summary>context</summary>
<pre>14:36:41: INFO: 1 process: 1 internal.
14:36:41: ERROR: Build did NOT complete successfully
14:36:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A293)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
14:36:47: selecting podman as container runtime
14:36:47: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-operator is being used by the following container(s): 9311f6c2ebab024ca15294aafe5d852280dc712623b4835614575fb0de781fe0: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>

**1x**: _2026-11-05 08:58:16 &#43;0000 UTC_: <code>09:01:23: ERROR: no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A276)
<details>
<summary>all...</summary>

* _2026-11-05 08:58:16 &#43;0000 UTC_: <code>09:01:23: ERROR: no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A276)
<details><summary>context</summary>
<pre>09:01:23: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
09:01:23: 		download_info = ctx.download_and_extract(
09:01:23: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/toolchains/BUILD:274:27: @@rules_java~//toolchains:toolchain_java10_definition depends on @@bazel_tools//tools/jdk:toolchain_type in repository @@bazel_tools which failed to fetch. no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/src/tools/launcher/BUILD:10:14: Analysis failed
09:01:23: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>08:59:46: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A221)

<details><summary>context</summary>
<pre>08:59:46:   &lt;builtin&gt;: in &lt;toplevel&gt;
08:59:46: Repository rule http_archive defined at:
08:59:46:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
08:59:46: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:
08:59:46:    Traceback (most recent call last):
08:59:46: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:46: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:46: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp18112639366700063061/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A226)

<details><summary>context</summary>
<pre>08:59:46: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:46: 		download_info = ctx.download_and_extract(
08:59:46: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp18112639366700063061/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:46: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp18112639366700063061/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:46: INFO: Elapsed time: 38.429s, Critical Path: 0.60s
08:59:46: INFO: 1 process: 1 internal.
08:59:46: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:46: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A229)

<details><summary>context</summary>
<pre>08:59:46: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp18112639366700063061/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:46: INFO: Elapsed time: 38.429s, Critical Path: 0.60s
08:59:46: INFO: 1 process: 1 internal.
08:59:46: ERROR: Build did NOT complete successfully
08:59:46: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:46: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A230)

<details><summary>context</summary>
<pre>08:59:46: INFO: Elapsed time: 38.429s, Critical Path: 0.60s
08:59:46: INFO: 1 process: 1 internal.
08:59:46: ERROR: Build did NOT complete successfully
08:59:46: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A231)

<details><summary>context</summary>
<pre>08:59:46: INFO: 1 process: 1 internal.
08:59:46: ERROR: Build did NOT complete successfully
08:59:46: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>diagnostic</strong>: <code>09:00:04: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A246)

<details><summary>context</summary>
<pre>09:00:04:   &lt;builtin&gt;: in &lt;toplevel&gt;
09:00:04: Repository rule http_archive defined at:
09:00:04:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
09:00:04: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:
09:00:04:    Traceback (most recent call last):
09:00:04: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
09:00:04: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>09:00:04: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp2637779128729034844/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A251)

<details><summary>context</summary>
<pre>09:00:04: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
09:00:04: 		download_info = ctx.download_and_extract(
09:00:04: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp2637779128729034844/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:04: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp2637779128729034844/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:04: INFO: Elapsed time: 13.746s, Critical Path: 0.53s
09:00:04: INFO: 1 process: 1 internal.
09:00:04: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>09:00:04: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A254)

<details><summary>context</summary>
<pre>09:00:04: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp2637779128729034844/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:04: INFO: Elapsed time: 13.746s, Critical Path: 0.53s
09:00:04: INFO: 1 process: 1 internal.
09:00:04: ERROR: Build did NOT complete successfully
09:00:04: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>09:00:04: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A255)

<details><summary>context</summary>
<pre>09:00:04: INFO: Elapsed time: 13.746s, Critical Path: 0.53s
09:00:04: INFO: 1 process: 1 internal.
09:00:04: ERROR: Build did NOT complete successfully
09:00:04: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A256)

<details><summary>context</summary>
<pre>09:00:04: INFO: 1 process: 1 internal.
09:00:04: ERROR: Build did NOT complete successfully
09:00:04: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_python~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A271)

<details><summary>context</summary>
<pre>09:01:23:   &lt;builtin&gt;: in &lt;toplevel&gt;
09:01:23: Repository rule http_archive defined at:
09:01:23:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_python~&#39;:
09:01:23:    Traceback (most recent call last):
09:01:23: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
09:01:23: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/toolchains/BUILD:274:27: @@rules_java~//toolchains:toolchain_java10_definition depends on @@bazel_tools//tools/jdk:toolchain_type in repository @@bazel_tools which failed to fetch. no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A277)

<details><summary>context</summary>
<pre>09:01:23: 		download_info = ctx.download_and_extract(
09:01:23: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/toolchains/BUILD:274:27: @@rules_java~//toolchains:toolchain_java10_definition depends on @@bazel_tools//tools/jdk:toolchain_type in repository @@bazel_tools which failed to fetch. no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/src/tools/launcher/BUILD:10:14: Analysis failed
09:01:23: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
09:01:23: INFO: Elapsed time: 67.819s, Critical Path: 0.61s</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/src/tools/launcher/BUILD:10:14: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A278)

<details><summary>context</summary>
<pre>09:01:23: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/toolchains/BUILD:274:27: @@rules_java~//toolchains:toolchain_java10_definition depends on @@bazel_tools//tools/jdk:toolchain_type in repository @@bazel_tools which failed to fetch. no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/src/tools/launcher/BUILD:10:14: Analysis failed
09:01:23: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
09:01:23: INFO: Elapsed time: 67.819s, Critical Path: 0.61s
09:01:23: INFO: 1 process: 1 internal.</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A279)

<details><summary>context</summary>
<pre>09:01:23: ERROR: no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/toolchains/BUILD:274:27: @@rules_java~//toolchains:toolchain_java10_definition depends on @@bazel_tools//tools/jdk:toolchain_type in repository @@bazel_tools which failed to fetch. no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/src/tools/launcher/BUILD:10:14: Analysis failed
09:01:23: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
09:01:23: INFO: Elapsed time: 67.819s, Critical Path: 0.61s
09:01:23: INFO: 1 process: 1 internal.
09:01:23: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A282)

<details><summary>context</summary>
<pre>09:01:23: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
09:01:23: INFO: Elapsed time: 67.819s, Critical Path: 0.61s
09:01:23: INFO: 1 process: 1 internal.
09:01:23: ERROR: Build did NOT complete successfully
09:01:23: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A283)

<details><summary>context</summary>
<pre>09:01:23: INFO: Elapsed time: 67.819s, Critical Path: 0.61s
09:01:23: INFO: 1 process: 1 internal.
09:01:23: ERROR: Build did NOT complete successfully
09:01:23: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A284)

<details><summary>context</summary>
<pre>09:01:23: INFO: 1 process: 1 internal.
09:01:23: ERROR: Build did NOT complete successfully
09:01:23: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A295)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
09:02:10: selecting podman as container runtime
09:02:11: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-operator is being used by the following container(s): 8461bb1b9368374a39c102452c3ac4240460428f7aab3a0411def0f08de4de90: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>
<details>
<summary> bazel repository fetch failure (1x / 8.33%) </summary>

<hr/>

**1x**: _2026-11-05 08:58:00 &#43;0000 UTC_: <code>09:00:17: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A266)
<details>
<summary>all...</summary>

* _2026-11-05 08:58:00 &#43;0000 UTC_: <code>09:00:17: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A266)
<details><summary>context</summary>
<pre>09:00:17:   &lt;builtin&gt;: in &lt;toplevel&gt;
09:00:17: Repository rule http_archive defined at:
09:00:17:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
09:00:17: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:
09:00:17:    Traceback (most recent call last):
09:00:17: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
09:00:17: 		download_info = ctx.download_and_extract(</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>08:59:40: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A216)

<details><summary>context</summary>
<pre>08:59:40:   &lt;builtin&gt;: in &lt;toplevel&gt;
08:59:40: Repository rule http_archive defined at:
08:59:40:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
08:59:40: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:
08:59:40:    Traceback (most recent call last):
08:59:40: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:40: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp10655117106702680376/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A221)

<details><summary>context</summary>
<pre>08:59:40: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:40: 		download_info = ctx.download_and_extract(
08:59:40: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp10655117106702680376/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp10655117106702680376/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:41: INFO: Elapsed time: 57.381s, Critical Path: 0.64s
08:59:41: INFO: 1 process: 1 internal.
08:59:41: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:41: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A224)

<details><summary>context</summary>
<pre>08:59:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp10655117106702680376/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:41: INFO: Elapsed time: 57.381s, Critical Path: 0.64s
08:59:41: INFO: 1 process: 1 internal.
08:59:41: ERROR: Build did NOT complete successfully
08:59:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:41: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A225)

<details><summary>context</summary>
<pre>08:59:41: INFO: Elapsed time: 57.381s, Critical Path: 0.64s
08:59:41: INFO: 1 process: 1 internal.
08:59:41: ERROR: Build did NOT complete successfully
08:59:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A226)

<details><summary>context</summary>
<pre>08:59:41: INFO: 1 process: 1 internal.
08:59:41: ERROR: Build did NOT complete successfully
08:59:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:59: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A241)

<details><summary>context</summary>
<pre>08:59:59:   &lt;builtin&gt;: in &lt;toplevel&gt;
08:59:59: Repository rule http_archive defined at:
08:59:59:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
08:59:59: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:
08:59:59:    Traceback (most recent call last):
08:59:59: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:59: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:59: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp1347094595764469151/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A246)

<details><summary>context</summary>
<pre>08:59:59: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:59: 		download_info = ctx.download_and_extract(
08:59:59: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp1347094595764469151/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:59: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp1347094595764469151/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:59: INFO: Elapsed time: 13.436s, Critical Path: 0.58s
08:59:59: INFO: 1 process: 1 internal.
08:59:59: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:59: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A249)

<details><summary>context</summary>
<pre>08:59:59: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp1347094595764469151/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:59: INFO: Elapsed time: 13.436s, Critical Path: 0.58s
08:59:59: INFO: 1 process: 1 internal.
08:59:59: ERROR: Build did NOT complete successfully
08:59:59: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:59: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A250)

<details><summary>context</summary>
<pre>08:59:59: INFO: Elapsed time: 13.436s, Critical Path: 0.58s
08:59:59: INFO: 1 process: 1 internal.
08:59:59: ERROR: Build did NOT complete successfully
08:59:59: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A251)

<details><summary>context</summary>
<pre>08:59:59: INFO: 1 process: 1 internal.
08:59:59: ERROR: Build did NOT complete successfully
08:59:59: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>secondary</strong>: <code>09:00:17: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp7416242368100028705/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A271)

<details><summary>context</summary>
<pre>09:00:17: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
09:00:17: 		download_info = ctx.download_and_extract(
09:00:17: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp7416242368100028705/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:17: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp7416242368100028705/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:17: INFO: Elapsed time: 13.699s, Critical Path: 0.65s
09:00:17: INFO: 1 process: 1 internal.
09:00:17: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>secondary</strong>: <code>09:00:17: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A274)

<details><summary>context</summary>
<pre>09:00:17: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp7416242368100028705/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:17: INFO: Elapsed time: 13.699s, Critical Path: 0.65s
09:00:17: INFO: 1 process: 1 internal.
09:00:17: ERROR: Build did NOT complete successfully
09:00:17: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>secondary</strong>: <code>09:00:17: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A275)

<details><summary>context</summary>
<pre>09:00:17: INFO: Elapsed time: 13.699s, Critical Path: 0.65s
09:00:17: INFO: 1 process: 1 internal.
09:00:17: ERROR: Build did NOT complete successfully
09:00:17: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A276)

<details><summary>context</summary>
<pre>09:00:17: INFO: 1 process: 1 internal.
09:00:17: ERROR: Build did NOT complete successfully
09:00:17: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A287)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
09:00:26: selecting podman as container runtime
09:00:26: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute is being used by the following container(s): c8d4c791d2f91c3a56c506d6b5653b81a2c82208835724ffd468f4be28ac67c2: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>
<details>
<summary> container image pull failure (1x / 8.33%) </summary>

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

### internal (7x / 58.33%)

<details>
<summary> make cluster lifecycle target failure (3x / 25.00%) </summary>

<hr/>

**3x**: _2026-11-03 11:59:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2117586999056535552#1:build-log.txt%3A1217)
<details>
<summary>all...</summary>

* _2026-11-03 12:09:09 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2117587001577312256#1:build-log.txt%3A4051)
<details><summary>context</summary>
<pre>12:30:01: 2026/11/03 12:30:01 existing blob: sha256:05daf2a1d5cd22bf20bf774e8599d99ee37f8edccf2588aaa86f21aaec88c88d
12:30:02: 2026/11/03 12:30:02 localhost:36275/kubevirt/kv-network-passt-binding-cni:devel: digest: sha256:0d6e045da432423816b2bc90af03985fc7c0432ef9ca321aa853181e18ba8a66 size: 701
12:30:17: &#43; rm -f /tmp/kubevirt.deploy.jwbQ
make: *** [Makefile:174: cluster-sync] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2117587001577312256#1:build-log.txt%3A4057)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:30:29: selecting podman as container runtime
12:32:15: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.33-sig-network-1.7 is being used by the following container(s): 437e990e2003afcfa461cefe139c5de3abcd06e6208ef33c9e77cb53b350a614, d7746eb1230afd44770adaf9ee2a421590a3d28144b9974b1f14b2ea60f6311e: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



* _2026-11-03 11:59:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2117586999056535552#1:build-log.txt%3A1217)
<details><summary>context</summary>
<pre>12:35:36: INFO: 2 processes: 1 internal, 1 processwrapper-sandbox.
12:35:36: INFO: Running command line: bazel-bin/example-guest-agent-copier /root/go/src/kubevirt.io/kubevirt/_out/cmd/example-guest-agent/example-guest-agent
12:36:36: &#43; rm -f /tmp/kubevirt.deploy.1MC8
make: *** [Makefile:174: cluster-sync] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2117586999056535552#1:build-log.txt%3A1223)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:36:37: selecting podman as container runtime
12:36:52: Error response from daemon: cannot remove container 09207706028f593ad43cb2e4bf5ea652f5970c6efdf3a476d3f2ee015494902d as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



* _2026-11-03 11:58:41 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A3943)
<details><summary>context</summary>
<pre>12:27:25: 2026/11/03 12:27:25 existing blob: sha256:cc7f351ef2b73318cc7d1d133ee9c9dc6913543797cd8f8b294edade3700892e
12:27:25: 2026/11/03 12:27:25 localhost:46415/kubevirt/kv-network-passt-binding-cni:devel: digest: sha256:0d6e045da432423816b2bc90af03985fc7c0432ef9ca321aa853181e18ba8a66 size: 701
12:27:34: &#43; rm -f /tmp/kubevirt.deploy.hT3Q
make: *** [Makefile:174: cluster-sync] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>12:15:30: I1103 07:15:30.373979    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A798)

<details><summary>context</summary>
<pre>12:15:21: [control-plane-check] Checking kube-scheduler at https://127.0.0.1:10259/livez
12:15:25: [control-plane-check] kube-scheduler is healthy after 3.842462827s
12:15:26: [control-plane-check] kube-controller-manager is healthy after 5.068241346s
12:15:30: I1103 07:15:30.373979    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:15:30: [control-plane-check] kube-apiserver is healthy after 9.002895106s
12:15:30: I1103 07:15:30.875756    1605 kubeconfig.go:657] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists
12:15:30: I1103 07:15:30.878733    1605 kubeconfig.go:730] creating the ClusterRoleBinding for the kubeadm:cluster-admins Group by using super-admin.conf</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A3949)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:27:34: selecting podman as container runtime
12:27:49: Error response from daemon: cannot remove container 4d2bdd9539b64bcf49fcea3d1d98ac72c4d24a6d0e4c0180ba3525a5a560ec7c as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>
<details>
<summary> taint not found during cluster setup (2x / 16.67%) </summary>

<hr/>

**2x**: _2026-11-05 08:57:54 &#43;0000 UTC_: <code>09:06:39: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A1209)
<details>
<summary>all...</summary>

* _2026-11-05 08:57:54 &#43;0000 UTC_: <code>09:06:39: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A1209)
<details><summary>context</summary>
<pre>09:06:39: &#43;&#43; for node in ${master_nodes[@]}
09:06:39: &#43;&#43; _kubectl taint nodes kind-1.37-control-plane node-role.kubernetes.io/master:NoSchedule-
09:06:39: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.37/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.37/.kubeconfig taint nodes kind-1.37-control-plane node-role.kubernetes.io/master:NoSchedule-
09:06:39: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found
09:06:39: &#43;&#43; _kubectl taint nodes kind-1.37-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
09:06:39: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.37/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.37/.kubeconfig taint nodes kind-1.37-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
09:06:39: node/kind-1.37-control-plane untainted</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>09:06:14: I1105 09:06:12.855745     276 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A827)

<details><summary>context</summary>
<pre>09:06:14: I1105 09:06:12.355808     276 round_trippers.go:632] &#34;Response&#34; verb=&#34;GET&#34; url=&#34;https://10.89.0.2:6443/livez?timeout=10s&#34; status=&#34;403 Forbidden&#34; milliseconds=0
09:06:14: I1105 09:06:12.356102     276 wait.go:283] kube-apiserver check failed at https://10.89.0.2:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: rbac: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
09:06:14: I1105 09:06:12.855440     276 round_trippers.go:632] &#34;Response&#34; verb=&#34;GET&#34; url=&#34;https://10.89.0.2:6443/livez?timeout=10s&#34; status=&#34;500 Internal Server Error&#34; milliseconds=0
09:06:14: I1105 09:06:12.855745     276 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
09:06:14: I1105 09:06:12.855795     276 wait.go:283] kube-apiserver check failed at https://10.89.0.2:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
09:06:14: I1105 09:06:13.356423     276 round_trippers.go:632] &#34;Response&#34; verb=&#34;GET&#34; url=&#34;https://10.89.0.2:6443/livez?timeout=10s&#34; status=&#34;200 OK&#34; milliseconds=1
09:06:14: [control-plane-check] kube-apiserver is healthy after 1.502411537s</pre>
</details>



* <strong>diagnostic</strong>: <code>09:06:24: I1105 09:06:16.972362     282 token.go:249] [discovery] Retrying due to error: could not find a JWS signature in the cluster-info ConfigMap for token ID &#34;abcdef&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A1081)

<details><summary>context</summary>
<pre>09:06:24: I1105 09:06:16.963598     282 envvar.go:195] &#34;Feature gate default state&#34; feature=&#34;UnlockWhileProcessingFIFO&#34; enabled=true
09:06:24: I1105 09:06:16.963944     282 token.go:229] [discovery] Waiting for the cluster-info ConfigMap to receive a JWS signature for token ID &#34;abcdef&#34;
09:06:24: I1105 09:06:16.971994     282 round_trippers.go:632] &#34;Response&#34; verb=&#34;GET&#34; url=&#34;https://kind-1.37-control-plane:6443/api/v1/namespaces/kube-public/configmaps/cluster-info?timeout=10s&#34; status=&#34;200 OK&#34; milliseconds=7
09:06:24: I1105 09:06:16.972362     282 token.go:249] [discovery] Retrying due to error: could not find a JWS signature in the cluster-info ConfigMap for token ID &#34;abcdef&#34;
09:06:24: I1105 09:06:21.964117     282 token.go:229] [discovery] Waiting for the cluster-info ConfigMap to receive a JWS signature for token ID &#34;abcdef&#34;
09:06:24: I1105 09:06:21.968475     282 round_trippers.go:632] &#34;Response&#34; verb=&#34;GET&#34; url=&#34;https://kind-1.37-control-plane:6443/api/v1/namespaces/kube-public/configmaps/cluster-info?timeout=10s&#34; status=&#34;200 OK&#34; milliseconds=4
09:06:24: I1105 09:06:21.968916     282 token.go:113] [discovery] Cluster info signature and contents are valid and no TLS pinning was specified, will use API Server &#34;kind-1.37-control-plane:6443&#34;</pre>
</details>



* <strong>secondary</strong>: <code>09:07:09: error: unable to read URL &#34;https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml&#34;, server reported 504 Gateway Timeout, status code=504</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A1349)

<details><summary>context</summary>
<pre>09:06:56: deployment.apps/cdi-operator created
09:06:56: &#43;&#43; _kubectl create -f https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml
09:06:56: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.37/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.37/.kubeconfig create -f https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml
09:07:09: error: unable to read URL &#34;https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml&#34;, server reported 504 Gateway Timeout, status code=504
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A1350)

<details><summary>context</summary>
<pre>09:06:56: &#43;&#43; _kubectl create -f https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml
09:06:56: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.37/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.37/.kubeconfig create -f https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml
09:07:09: error: unable to read URL &#34;https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml&#34;, server reported 504 Gateway Timeout, status code=504
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>



* _2026-11-03 11:58:13 &#43;0000 UTC_: <code>12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1231)
<details><summary>context</summary>
<pre>12:24:20: &#43;&#43; for node in ${master_nodes[@]}
12:24:20: &#43;&#43; _kubectl taint nodes vgpu-control-plane node-role.kubernetes.io/master:NoSchedule-
12:24:20: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubeconfig taint nodes vgpu-control-plane node-role.kubernetes.io/master:NoSchedule-
12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found
12:24:20: &#43;&#43; _kubectl taint nodes vgpu-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
12:24:20: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubeconfig taint nodes vgpu-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
12:24:20: error: taint &#34;node-role.kubernetes.io/control-plane:NoSchedule&#34; not found</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>12:24:20: error: taint &#34;node-role.kubernetes.io/control-plane:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1234)

<details><summary>context</summary>
<pre>12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found
12:24:20: &#43;&#43; _kubectl taint nodes vgpu-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
12:24:20: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubeconfig taint nodes vgpu-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
12:24:20: error: taint &#34;node-role.kubernetes.io/control-plane:NoSchedule&#34; not found
12:24:20: &#43;&#43;&#43; grep -i worker
12:24:20: &#43;&#43;&#43; _get_nodes
12:24:20: &#43;&#43;&#43; awk &#39;{print $1}&#39;</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1541)

<details><summary>context</summary>
<pre>12:29:21: INFO: 2 processes: 1 internal, 1 processwrapper-sandbox.
12:29:21: INFO: Running command line: bazel-bin/example-guest-agent-copier /root/go/src/kubevirt.io/kubevirt/_out/cmd/example-guest-agent/example-guest-agent
12:29:49: &#43; rm -f /tmp/kubevirt.deploy.oDzW
make: *** [Makefile:174: cluster-sync] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1550)

<details><summary>context</summary>
<pre>12:29:50: using podman due to KIND_EXPERIMENTAL_PROVIDER
12:29:50: enabling experimental podman provider
12:29:50: Deleting cluster &#34;vgpu&#34; ...
12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125
12:30:05:
12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1553)

<details><summary>context</summary>
<pre>12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125
12:30:05:
12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1555)

<details><summary>context</summary>
<pre>12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125
&#43; true
&#43; exit 2</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1556)

<details><summary>context</summary>
<pre>12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1568)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container kind_ardinghelli in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1569)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container kind_ardinghelli in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1570)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell
================================================================================</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1571)

<details><summary>context</summary>
<pre>Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell
================================================================================
Done cleaning up after podman in container.</pre>
</details>


</details>



</details>

<hr/>
</details>
<details>
<summary> configured node not found during provisioning (2x / 16.67%) </summary>

<hr/>

**2x**: _2026-11-06 12:02:09 &#43;0000 UTC_: <code>12:18:05: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19813/pull-kubevirt-e2e-k8s-1.36-windows2016/2118675165452701696#1:build-log.txt%3A868)
<details>
<summary>all...</summary>

* _2026-11-06 12:02:09 &#43;0000 UTC_: <code>12:18:05: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19813/pull-kubevirt-e2e-k8s-1.36-windows2016/2118675165452701696#1:build-log.txt%3A868)
<details><summary>context</summary>
<pre>12:18:05: Warning: would violate PodSecurity &#34;restricted:latest&#34;: host namespaces (hostNetwork=true), allowPrivilegeEscalation != false (containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.allowPrivilegeEscalation=false), unrestricted capabilities (containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.capabilities.drop=[&#34;ALL&#34;]; container &#34;kube-flannel&#34; must not include &#34;NET_ADMIN&#34;, &#34;NET_RAW&#34; in securityContext.capabilities.add), restricted volume types (volumes &#34;run&#34;, &#34;cni-plugin&#34;, &#34;cni&#34;, &#34;xtables-lock&#34; use restricted volume type &#34;hostPath&#34;), runAsNonRoot != true (pod or containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.runAsNonRoot=true), seccompProfile (pod or containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.seccompProfile.type to &#34;RuntimeDefault&#34; or &#34;Localhost&#34;)
12:18:05: daemonset.apps/kube-flannel-ds created
12:18:05: time=&#34;2026-11-06T12:18:05Z&#34; level=info msg=&#34;[node 1]: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-&#34;
12:18:05: Error from server (NotFound): nodes &#34;node01&#34; not found
12:18:06:
12:18:06: ===== ab78c164736eed2ff9ecce5b6ad7cabf7c1d40a64b98015a3518bbac94e4b1d7 ====
12:18:06: &#43; NUM_NODES=2</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>12:18:02: I1106 07:18:02.507609    1888 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19813/pull-kubevirt-e2e-k8s-1.36-windows2016/2118675165452701696#1:build-log.txt%3A805)

<details><summary>context</summary>
<pre>12:18:01: [control-plane-check] kube-controller-manager is healthy after 5.239285ms
12:18:01: [control-plane-check] kube-scheduler is healthy after 7.898644ms
12:18:02: I1106 07:18:02.008369    1888 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
12:18:02: I1106 07:18:02.507609    1888 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:18:02: I1106 07:18:02.507711    1888 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:18:02: [control-plane-check] kube-apiserver is healthy after 1.503624779s
12:18:02: I1106 07:18:03.008716    1888 uploadconfig.go:109] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19813/pull-kubevirt-e2e-k8s-1.36-windows2016/2118675165452701696#1:build-log.txt%3A1150)

<details><summary>context</summary>
<pre>12:18:17: 
12:18:17: &#43; wait 69
12:18:29: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19813/pull-kubevirt-e2e-k8s-1.36-windows2016/2118675165452701696#1:build-log.txt%3A1178)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:18:30: selecting podman as container runtime
12:18:30: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-windows2016 is being used by the following container(s): afe66bfd0ea325b2aedd124aef268e5e6cb650d1d5ccd698ca0210fa9d12b1b3: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



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



</details>

<hr/>
</details>

### needs-investigation (1x / 8.33%)

<details>
<summary> no error snippets (1x / 8.33%) </summary>

<hr/>

**1x**: _2026-11-04 07:13:51 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19748/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2117877790450651136)
<details>
<summary>all...</summary>

* _2026-11-04 07:13:51 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19748/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2117877790450651136)




</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (7x / 58.33%)


#### external (4x / 57.14%)

<details>
<summary> download failure from external URL (2x / 28.57%) </summary>

<hr/>

**1x**: _2026-11-05 14:33:27 &#43;0000 UTC_: <code>14:36:41: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp6849394626262802922/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A276)
<details>
<summary>all...</summary>

* _2026-11-05 14:33:27 &#43;0000 UTC_: <code>14:36:41: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp6849394626262802922/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A276)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>14:34:53: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_cc&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A218)

<details><summary>context</summary>
<pre>14:34:53:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/io_bazel_rules_go/go/private/repositories.bzl:355:18: in _maybe
14:34:53: Repository rule http_archive defined at:
14:34:53:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
14:34:53: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_cc&#39;:
14:34:53:    Traceback (most recent call last):
14:34:53: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
14:34:53: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>14:34:53: ERROR: Skipping &#39;//:gazelle&#39;: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A224)

<details><summary>context</summary>
<pre>14:34:53: 		download_info = ctx.download_and_extract(
14:34:53: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: WARNING: Target pattern parsing failed.
14:34:53: ERROR: Skipping &#39;//:gazelle&#39;: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: ERROR: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: INFO: Elapsed time: 49.216s
14:34:53: INFO: 0 processes.</pre>
</details>



* <strong>diagnostic</strong>: <code>14:34:53: ERROR: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A225)

<details><summary>context</summary>
<pre>14:34:53: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: WARNING: Target pattern parsing failed.
14:34:53: ERROR: Skipping &#39;//:gazelle&#39;: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: ERROR: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: INFO: Elapsed time: 49.216s
14:34:53: INFO: 0 processes.
14:34:53: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>14:34:53: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A228)

<details><summary>context</summary>
<pre>14:34:53: ERROR: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: INFO: Elapsed time: 49.216s
14:34:53: INFO: 0 processes.
14:34:53: ERROR: Build did NOT complete successfully
14:34:53: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>14:34:53: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A229)

<details><summary>context</summary>
<pre>14:34:53: INFO: Elapsed time: 49.216s
14:34:53: INFO: 0 processes.
14:34:53: ERROR: Build did NOT complete successfully
14:34:53: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A230)

<details><summary>context</summary>
<pre>14:34:53: INFO: 0 processes.
14:34:53: ERROR: Build did NOT complete successfully
14:34:53: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>diagnostic</strong>: <code>14:36:03: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;apple_support~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A245)

<details><summary>context</summary>
<pre>14:36:03:   &lt;builtin&gt;: in &lt;toplevel&gt;
14:36:03: Repository rule http_archive defined at:
14:36:03:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
14:36:03: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;apple_support~&#39;:
14:36:03:    Traceback (most recent call last):
14:36:03: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
14:36:03: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>14:36:03: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp1387829800134064954/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A250)

<details><summary>context</summary>
<pre>14:36:03: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
14:36:03: 		download_info = ctx.download_and_extract(
14:36:03: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp1387829800134064954/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:03: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp1387829800134064954/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:03: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
14:36:03: INFO: Elapsed time: 48.439s, Critical Path: 0.60s
14:36:03: INFO: 1 process: 1 internal.</pre>
</details>



* <strong>diagnostic</strong>: <code>14:36:03: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A251)

<details><summary>context</summary>
<pre>14:36:03: 		download_info = ctx.download_and_extract(
14:36:03: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp1387829800134064954/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:03: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp1387829800134064954/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:03: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
14:36:03: INFO: Elapsed time: 48.439s, Critical Path: 0.60s
14:36:03: INFO: 1 process: 1 internal.
14:36:03: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>14:36:03: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A254)

<details><summary>context</summary>
<pre>14:36:03: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
14:36:03: INFO: Elapsed time: 48.439s, Critical Path: 0.60s
14:36:03: INFO: 1 process: 1 internal.
14:36:03: ERROR: Build did NOT complete successfully
14:36:03: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>14:36:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A255)

<details><summary>context</summary>
<pre>14:36:03: INFO: Elapsed time: 48.439s, Critical Path: 0.60s
14:36:03: INFO: 1 process: 1 internal.
14:36:03: ERROR: Build did NOT complete successfully
14:36:03: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A256)

<details><summary>context</summary>
<pre>14:36:03: INFO: 1 process: 1 internal.
14:36:03: ERROR: Build did NOT complete successfully
14:36:03: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>secondary</strong>: <code>14:36:41: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;apple_support~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A271)

<details><summary>context</summary>
<pre>14:36:41:   &lt;builtin&gt;: in &lt;toplevel&gt;
14:36:41: Repository rule http_archive defined at:
14:36:41:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
14:36:41: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;apple_support~&#39;:
14:36:41:    Traceback (most recent call last):
14:36:41: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
14:36:41: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>secondary</strong>: <code>14:36:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A277)

<details><summary>context</summary>
<pre>14:36:41: 		download_info = ctx.download_and_extract(
14:36:41: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp6849394626262802922/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:41: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp6849394626262802922/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
14:36:41: INFO: Elapsed time: 13.016s, Critical Path: 0.60s
14:36:41: INFO: 1 process: 1 internal.
14:36:41: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>secondary</strong>: <code>14:36:41: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A280)

<details><summary>context</summary>
<pre>14:36:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
14:36:41: INFO: Elapsed time: 13.016s, Critical Path: 0.60s
14:36:41: INFO: 1 process: 1 internal.
14:36:41: ERROR: Build did NOT complete successfully
14:36:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>secondary</strong>: <code>14:36:41: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A281)

<details><summary>context</summary>
<pre>14:36:41: INFO: Elapsed time: 13.016s, Critical Path: 0.60s
14:36:41: INFO: 1 process: 1 internal.
14:36:41: ERROR: Build did NOT complete successfully
14:36:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A282)

<details><summary>context</summary>
<pre>14:36:41: INFO: 1 process: 1 internal.
14:36:41: ERROR: Build did NOT complete successfully
14:36:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A293)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
14:36:47: selecting podman as container runtime
14:36:47: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-operator is being used by the following container(s): 9311f6c2ebab024ca15294aafe5d852280dc712623b4835614575fb0de781fe0: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>

**1x**: _2026-11-05 08:58:16 &#43;0000 UTC_: <code>09:01:23: ERROR: no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A276)
<details>
<summary>all...</summary>

* _2026-11-05 08:58:16 &#43;0000 UTC_: <code>09:01:23: ERROR: no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A276)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>08:59:46: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A221)

<details><summary>context</summary>
<pre>08:59:46:   &lt;builtin&gt;: in &lt;toplevel&gt;
08:59:46: Repository rule http_archive defined at:
08:59:46:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
08:59:46: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:
08:59:46:    Traceback (most recent call last):
08:59:46: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:46: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:46: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp18112639366700063061/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A226)

<details><summary>context</summary>
<pre>08:59:46: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:46: 		download_info = ctx.download_and_extract(
08:59:46: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp18112639366700063061/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:46: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp18112639366700063061/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:46: INFO: Elapsed time: 38.429s, Critical Path: 0.60s
08:59:46: INFO: 1 process: 1 internal.
08:59:46: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:46: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A229)

<details><summary>context</summary>
<pre>08:59:46: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp18112639366700063061/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:46: INFO: Elapsed time: 38.429s, Critical Path: 0.60s
08:59:46: INFO: 1 process: 1 internal.
08:59:46: ERROR: Build did NOT complete successfully
08:59:46: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:46: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A230)

<details><summary>context</summary>
<pre>08:59:46: INFO: Elapsed time: 38.429s, Critical Path: 0.60s
08:59:46: INFO: 1 process: 1 internal.
08:59:46: ERROR: Build did NOT complete successfully
08:59:46: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A231)

<details><summary>context</summary>
<pre>08:59:46: INFO: 1 process: 1 internal.
08:59:46: ERROR: Build did NOT complete successfully
08:59:46: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>diagnostic</strong>: <code>09:00:04: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A246)

<details><summary>context</summary>
<pre>09:00:04:   &lt;builtin&gt;: in &lt;toplevel&gt;
09:00:04: Repository rule http_archive defined at:
09:00:04:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
09:00:04: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:
09:00:04:    Traceback (most recent call last):
09:00:04: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
09:00:04: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>09:00:04: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp2637779128729034844/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A251)

<details><summary>context</summary>
<pre>09:00:04: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
09:00:04: 		download_info = ctx.download_and_extract(
09:00:04: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp2637779128729034844/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:04: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp2637779128729034844/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:04: INFO: Elapsed time: 13.746s, Critical Path: 0.53s
09:00:04: INFO: 1 process: 1 internal.
09:00:04: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>09:00:04: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A254)

<details><summary>context</summary>
<pre>09:00:04: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp2637779128729034844/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:04: INFO: Elapsed time: 13.746s, Critical Path: 0.53s
09:00:04: INFO: 1 process: 1 internal.
09:00:04: ERROR: Build did NOT complete successfully
09:00:04: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>09:00:04: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A255)

<details><summary>context</summary>
<pre>09:00:04: INFO: Elapsed time: 13.746s, Critical Path: 0.53s
09:00:04: INFO: 1 process: 1 internal.
09:00:04: ERROR: Build did NOT complete successfully
09:00:04: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A256)

<details><summary>context</summary>
<pre>09:00:04: INFO: 1 process: 1 internal.
09:00:04: ERROR: Build did NOT complete successfully
09:00:04: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_python~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A271)

<details><summary>context</summary>
<pre>09:01:23:   &lt;builtin&gt;: in &lt;toplevel&gt;
09:01:23: Repository rule http_archive defined at:
09:01:23:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_python~&#39;:
09:01:23:    Traceback (most recent call last):
09:01:23: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
09:01:23: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/toolchains/BUILD:274:27: @@rules_java~//toolchains:toolchain_java10_definition depends on @@bazel_tools//tools/jdk:toolchain_type in repository @@bazel_tools which failed to fetch. no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A277)

<details><summary>context</summary>
<pre>09:01:23: 		download_info = ctx.download_and_extract(
09:01:23: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/toolchains/BUILD:274:27: @@rules_java~//toolchains:toolchain_java10_definition depends on @@bazel_tools//tools/jdk:toolchain_type in repository @@bazel_tools which failed to fetch. no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/src/tools/launcher/BUILD:10:14: Analysis failed
09:01:23: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
09:01:23: INFO: Elapsed time: 67.819s, Critical Path: 0.61s</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/src/tools/launcher/BUILD:10:14: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A278)

<details><summary>context</summary>
<pre>09:01:23: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/toolchains/BUILD:274:27: @@rules_java~//toolchains:toolchain_java10_definition depends on @@bazel_tools//tools/jdk:toolchain_type in repository @@bazel_tools which failed to fetch. no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/src/tools/launcher/BUILD:10:14: Analysis failed
09:01:23: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
09:01:23: INFO: Elapsed time: 67.819s, Critical Path: 0.61s
09:01:23: INFO: 1 process: 1 internal.</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A279)

<details><summary>context</summary>
<pre>09:01:23: ERROR: no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/toolchains/BUILD:274:27: @@rules_java~//toolchains:toolchain_java10_definition depends on @@bazel_tools//tools/jdk:toolchain_type in repository @@bazel_tools which failed to fetch. no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/src/tools/launcher/BUILD:10:14: Analysis failed
09:01:23: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
09:01:23: INFO: Elapsed time: 67.819s, Critical Path: 0.61s
09:01:23: INFO: 1 process: 1 internal.
09:01:23: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A282)

<details><summary>context</summary>
<pre>09:01:23: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
09:01:23: INFO: Elapsed time: 67.819s, Critical Path: 0.61s
09:01:23: INFO: 1 process: 1 internal.
09:01:23: ERROR: Build did NOT complete successfully
09:01:23: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A283)

<details><summary>context</summary>
<pre>09:01:23: INFO: Elapsed time: 67.819s, Critical Path: 0.61s
09:01:23: INFO: 1 process: 1 internal.
09:01:23: ERROR: Build did NOT complete successfully
09:01:23: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A284)

<details><summary>context</summary>
<pre>09:01:23: INFO: 1 process: 1 internal.
09:01:23: ERROR: Build did NOT complete successfully
09:01:23: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A295)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
09:02:10: selecting podman as container runtime
09:02:11: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-operator is being used by the following container(s): 8461bb1b9368374a39c102452c3ac4240460428f7aab3a0411def0f08de4de90: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>
<details>
<summary> container image pull failure (1x / 14.29%) </summary>

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
<details>
<summary> bazel repository fetch failure (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-11-05 08:58:00 &#43;0000 UTC_: <code>09:00:17: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A266)
<details>
<summary>all...</summary>

* _2026-11-05 08:58:00 &#43;0000 UTC_: <code>09:00:17: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A266)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>08:59:40: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A216)

<details><summary>context</summary>
<pre>08:59:40:   &lt;builtin&gt;: in &lt;toplevel&gt;
08:59:40: Repository rule http_archive defined at:
08:59:40:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
08:59:40: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:
08:59:40:    Traceback (most recent call last):
08:59:40: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:40: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp10655117106702680376/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A221)

<details><summary>context</summary>
<pre>08:59:40: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:40: 		download_info = ctx.download_and_extract(
08:59:40: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp10655117106702680376/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp10655117106702680376/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:41: INFO: Elapsed time: 57.381s, Critical Path: 0.64s
08:59:41: INFO: 1 process: 1 internal.
08:59:41: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:41: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A224)

<details><summary>context</summary>
<pre>08:59:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp10655117106702680376/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:41: INFO: Elapsed time: 57.381s, Critical Path: 0.64s
08:59:41: INFO: 1 process: 1 internal.
08:59:41: ERROR: Build did NOT complete successfully
08:59:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:41: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A225)

<details><summary>context</summary>
<pre>08:59:41: INFO: Elapsed time: 57.381s, Critical Path: 0.64s
08:59:41: INFO: 1 process: 1 internal.
08:59:41: ERROR: Build did NOT complete successfully
08:59:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A226)

<details><summary>context</summary>
<pre>08:59:41: INFO: 1 process: 1 internal.
08:59:41: ERROR: Build did NOT complete successfully
08:59:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:59: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A241)

<details><summary>context</summary>
<pre>08:59:59:   &lt;builtin&gt;: in &lt;toplevel&gt;
08:59:59: Repository rule http_archive defined at:
08:59:59:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
08:59:59: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:
08:59:59:    Traceback (most recent call last):
08:59:59: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:59: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:59: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp1347094595764469151/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A246)

<details><summary>context</summary>
<pre>08:59:59: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:59: 		download_info = ctx.download_and_extract(
08:59:59: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp1347094595764469151/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:59: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp1347094595764469151/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:59: INFO: Elapsed time: 13.436s, Critical Path: 0.58s
08:59:59: INFO: 1 process: 1 internal.
08:59:59: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:59: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A249)

<details><summary>context</summary>
<pre>08:59:59: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp1347094595764469151/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:59: INFO: Elapsed time: 13.436s, Critical Path: 0.58s
08:59:59: INFO: 1 process: 1 internal.
08:59:59: ERROR: Build did NOT complete successfully
08:59:59: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:59: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A250)

<details><summary>context</summary>
<pre>08:59:59: INFO: Elapsed time: 13.436s, Critical Path: 0.58s
08:59:59: INFO: 1 process: 1 internal.
08:59:59: ERROR: Build did NOT complete successfully
08:59:59: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A251)

<details><summary>context</summary>
<pre>08:59:59: INFO: 1 process: 1 internal.
08:59:59: ERROR: Build did NOT complete successfully
08:59:59: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>secondary</strong>: <code>09:00:17: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp7416242368100028705/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A271)

<details><summary>context</summary>
<pre>09:00:17: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
09:00:17: 		download_info = ctx.download_and_extract(
09:00:17: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp7416242368100028705/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:17: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp7416242368100028705/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:17: INFO: Elapsed time: 13.699s, Critical Path: 0.65s
09:00:17: INFO: 1 process: 1 internal.
09:00:17: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>secondary</strong>: <code>09:00:17: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A274)

<details><summary>context</summary>
<pre>09:00:17: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp7416242368100028705/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:17: INFO: Elapsed time: 13.699s, Critical Path: 0.65s
09:00:17: INFO: 1 process: 1 internal.
09:00:17: ERROR: Build did NOT complete successfully
09:00:17: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>secondary</strong>: <code>09:00:17: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A275)

<details><summary>context</summary>
<pre>09:00:17: INFO: Elapsed time: 13.699s, Critical Path: 0.65s
09:00:17: INFO: 1 process: 1 internal.
09:00:17: ERROR: Build did NOT complete successfully
09:00:17: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A276)

<details><summary>context</summary>
<pre>09:00:17: INFO: 1 process: 1 internal.
09:00:17: ERROR: Build did NOT complete successfully
09:00:17: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A287)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
09:00:26: selecting podman as container runtime
09:00:26: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute is being used by the following container(s): c8d4c791d2f91c3a56c506d6b5653b81a2c82208835724ffd468f4be28ac67c2: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

#### internal (3x / 42.86%)

<details>
<summary> configured node not found during provisioning (2x / 28.57%) </summary>

<hr/>

**2x**: _2026-11-06 12:02:09 &#43;0000 UTC_: <code>12:18:05: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19813/pull-kubevirt-e2e-k8s-1.36-windows2016/2118675165452701696#1:build-log.txt%3A868)
<details>
<summary>all...</summary>

* _2026-11-06 12:02:09 &#43;0000 UTC_: <code>12:18:05: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19813/pull-kubevirt-e2e-k8s-1.36-windows2016/2118675165452701696#1:build-log.txt%3A868)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>12:18:02: I1106 07:18:02.507609    1888 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19813/pull-kubevirt-e2e-k8s-1.36-windows2016/2118675165452701696#1:build-log.txt%3A805)

<details><summary>context</summary>
<pre>12:18:01: [control-plane-check] kube-controller-manager is healthy after 5.239285ms
12:18:01: [control-plane-check] kube-scheduler is healthy after 7.898644ms
12:18:02: I1106 07:18:02.008369    1888 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
12:18:02: I1106 07:18:02.507609    1888 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:18:02: I1106 07:18:02.507711    1888 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:18:02: [control-plane-check] kube-apiserver is healthy after 1.503624779s
12:18:02: I1106 07:18:03.008716    1888 uploadconfig.go:109] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19813/pull-kubevirt-e2e-k8s-1.36-windows2016/2118675165452701696#1:build-log.txt%3A1150)

<details><summary>context</summary>
<pre>12:18:17: 
12:18:17: &#43; wait 69
12:18:29: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19813/pull-kubevirt-e2e-k8s-1.36-windows2016/2118675165452701696#1:build-log.txt%3A1178)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:18:30: selecting podman as container runtime
12:18:30: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-windows2016 is being used by the following container(s): afe66bfd0ea325b2aedd124aef268e5e6cb650d1d5ccd698ca0210fa9d12b1b3: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



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
<details>
<summary> taint not found during cluster setup (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-11-05 08:57:54 &#43;0000 UTC_: <code>09:06:39: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A1209)
<details>
<summary>all...</summary>

* _2026-11-05 08:57:54 &#43;0000 UTC_: <code>09:06:39: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A1209)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>09:06:14: I1105 09:06:12.855745     276 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A827)

<details><summary>context</summary>
<pre>09:06:14: I1105 09:06:12.355808     276 round_trippers.go:632] &#34;Response&#34; verb=&#34;GET&#34; url=&#34;https://10.89.0.2:6443/livez?timeout=10s&#34; status=&#34;403 Forbidden&#34; milliseconds=0
09:06:14: I1105 09:06:12.356102     276 wait.go:283] kube-apiserver check failed at https://10.89.0.2:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: rbac: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
09:06:14: I1105 09:06:12.855440     276 round_trippers.go:632] &#34;Response&#34; verb=&#34;GET&#34; url=&#34;https://10.89.0.2:6443/livez?timeout=10s&#34; status=&#34;500 Internal Server Error&#34; milliseconds=0
09:06:14: I1105 09:06:12.855745     276 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
09:06:14: I1105 09:06:12.855795     276 wait.go:283] kube-apiserver check failed at https://10.89.0.2:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
09:06:14: I1105 09:06:13.356423     276 round_trippers.go:632] &#34;Response&#34; verb=&#34;GET&#34; url=&#34;https://10.89.0.2:6443/livez?timeout=10s&#34; status=&#34;200 OK&#34; milliseconds=1
09:06:14: [control-plane-check] kube-apiserver is healthy after 1.502411537s</pre>
</details>



* <strong>diagnostic</strong>: <code>09:06:24: I1105 09:06:16.972362     282 token.go:249] [discovery] Retrying due to error: could not find a JWS signature in the cluster-info ConfigMap for token ID &#34;abcdef&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A1081)

<details><summary>context</summary>
<pre>09:06:24: I1105 09:06:16.963598     282 envvar.go:195] &#34;Feature gate default state&#34; feature=&#34;UnlockWhileProcessingFIFO&#34; enabled=true
09:06:24: I1105 09:06:16.963944     282 token.go:229] [discovery] Waiting for the cluster-info ConfigMap to receive a JWS signature for token ID &#34;abcdef&#34;
09:06:24: I1105 09:06:16.971994     282 round_trippers.go:632] &#34;Response&#34; verb=&#34;GET&#34; url=&#34;https://kind-1.37-control-plane:6443/api/v1/namespaces/kube-public/configmaps/cluster-info?timeout=10s&#34; status=&#34;200 OK&#34; milliseconds=7
09:06:24: I1105 09:06:16.972362     282 token.go:249] [discovery] Retrying due to error: could not find a JWS signature in the cluster-info ConfigMap for token ID &#34;abcdef&#34;
09:06:24: I1105 09:06:21.964117     282 token.go:229] [discovery] Waiting for the cluster-info ConfigMap to receive a JWS signature for token ID &#34;abcdef&#34;
09:06:24: I1105 09:06:21.968475     282 round_trippers.go:632] &#34;Response&#34; verb=&#34;GET&#34; url=&#34;https://kind-1.37-control-plane:6443/api/v1/namespaces/kube-public/configmaps/cluster-info?timeout=10s&#34; status=&#34;200 OK&#34; milliseconds=4
09:06:24: I1105 09:06:21.968916     282 token.go:113] [discovery] Cluster info signature and contents are valid and no TLS pinning was specified, will use API Server &#34;kind-1.37-control-plane:6443&#34;</pre>
</details>



* <strong>secondary</strong>: <code>09:07:09: error: unable to read URL &#34;https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml&#34;, server reported 504 Gateway Timeout, status code=504</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A1349)

<details><summary>context</summary>
<pre>09:06:56: deployment.apps/cdi-operator created
09:06:56: &#43;&#43; _kubectl create -f https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml
09:06:56: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.37/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.37/.kubeconfig create -f https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml
09:07:09: error: unable to read URL &#34;https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml&#34;, server reported 504 Gateway Timeout, status code=504
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A1350)

<details><summary>context</summary>
<pre>09:06:56: &#43;&#43; _kubectl create -f https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml
09:06:56: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.37/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.37/.kubeconfig create -f https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml
09:07:09: error: unable to read URL &#34;https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml&#34;, server reported 504 Gateway Timeout, status code=504
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>



</details>

<hr/>
</details>

### release-1.8 (1x / 8.33%)


#### needs-investigation (1x / 100.00%)

<details>
<summary> no error snippets (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-11-04 07:13:51 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19748/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2117877790450651136)
<details>
<summary>all...</summary>

* _2026-11-04 07:13:51 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19748/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2117877790450651136)




</details>

<hr/>
</details>

### release-1.7 (4x / 33.33%)


#### internal (4x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (3x / 75.00%) </summary>

<hr/>

**3x**: _2026-11-03 11:59:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2117586999056535552#1:build-log.txt%3A1217)
<details>
<summary>all...</summary>

* _2026-11-03 12:09:09 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2117587001577312256#1:build-log.txt%3A4051)


<details>
<summary>secondary evidence</summary>


* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2117587001577312256#1:build-log.txt%3A4057)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:30:29: selecting podman as container runtime
12:32:15: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.33-sig-network-1.7 is being used by the following container(s): 437e990e2003afcfa461cefe139c5de3abcd06e6208ef33c9e77cb53b350a614, d7746eb1230afd44770adaf9ee2a421590a3d28144b9974b1f14b2ea60f6311e: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



* _2026-11-03 11:59:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2117586999056535552#1:build-log.txt%3A1217)


<details>
<summary>secondary evidence</summary>


* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2117586999056535552#1:build-log.txt%3A1223)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:36:37: selecting podman as container runtime
12:36:52: Error response from daemon: cannot remove container 09207706028f593ad43cb2e4bf5ea652f5970c6efdf3a476d3f2ee015494902d as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



* _2026-11-03 11:58:41 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A3943)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>12:15:30: I1103 07:15:30.373979    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A798)

<details><summary>context</summary>
<pre>12:15:21: [control-plane-check] Checking kube-scheduler at https://127.0.0.1:10259/livez
12:15:25: [control-plane-check] kube-scheduler is healthy after 3.842462827s
12:15:26: [control-plane-check] kube-controller-manager is healthy after 5.068241346s
12:15:30: I1103 07:15:30.373979    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:15:30: [control-plane-check] kube-apiserver is healthy after 9.002895106s
12:15:30: I1103 07:15:30.875756    1605 kubeconfig.go:657] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists
12:15:30: I1103 07:15:30.878733    1605 kubeconfig.go:730] creating the ClusterRoleBinding for the kubeadm:cluster-admins Group by using super-admin.conf</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A3949)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:27:34: selecting podman as container runtime
12:27:49: Error response from daemon: cannot remove container 4d2bdd9539b64bcf49fcea3d1d98ac72c4d24a6d0e4c0180ba3525a5a560ec7c as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>
<details>
<summary> taint not found during cluster setup (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-11-03 11:58:13 &#43;0000 UTC_: <code>12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1231)
<details>
<summary>all...</summary>

* _2026-11-03 11:58:13 &#43;0000 UTC_: <code>12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1231)


<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>12:24:20: error: taint &#34;node-role.kubernetes.io/control-plane:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1234)

<details><summary>context</summary>
<pre>12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found
12:24:20: &#43;&#43; _kubectl taint nodes vgpu-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
12:24:20: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubeconfig taint nodes vgpu-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
12:24:20: error: taint &#34;node-role.kubernetes.io/control-plane:NoSchedule&#34; not found
12:24:20: &#43;&#43;&#43; grep -i worker
12:24:20: &#43;&#43;&#43; _get_nodes
12:24:20: &#43;&#43;&#43; awk &#39;{print $1}&#39;</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1541)

<details><summary>context</summary>
<pre>12:29:21: INFO: 2 processes: 1 internal, 1 processwrapper-sandbox.
12:29:21: INFO: Running command line: bazel-bin/example-guest-agent-copier /root/go/src/kubevirt.io/kubevirt/_out/cmd/example-guest-agent/example-guest-agent
12:29:49: &#43; rm -f /tmp/kubevirt.deploy.oDzW
make: *** [Makefile:174: cluster-sync] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1550)

<details><summary>context</summary>
<pre>12:29:50: using podman due to KIND_EXPERIMENTAL_PROVIDER
12:29:50: enabling experimental podman provider
12:29:50: Deleting cluster &#34;vgpu&#34; ...
12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125
12:30:05:
12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1553)

<details><summary>context</summary>
<pre>12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125
12:30:05:
12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1555)

<details><summary>context</summary>
<pre>12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125
&#43; true
&#43; exit 2</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1556)

<details><summary>context</summary>
<pre>12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1568)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container kind_ardinghelli in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1569)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container kind_ardinghelli in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1570)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell
================================================================================</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1571)

<details><summary>context</summary>
<pre>Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell
================================================================================
Done cleaning up after podman in container.</pre>
</details>


</details>



</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (9x / 75.00%)


#### internal (5x / 55.56%)

<details>
<summary> taint not found during cluster setup (2x / 22.22%) </summary>

<hr/>

**2x**: _2026-11-05 08:57:54 &#43;0000 UTC_: <code>09:06:39: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A1209)
<details>
<summary>all...</summary>

* _2026-11-05 08:57:54 &#43;0000 UTC_: <code>09:06:39: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A1209)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>09:06:14: I1105 09:06:12.855745     276 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A827)

<details><summary>context</summary>
<pre>09:06:14: I1105 09:06:12.355808     276 round_trippers.go:632] &#34;Response&#34; verb=&#34;GET&#34; url=&#34;https://10.89.0.2:6443/livez?timeout=10s&#34; status=&#34;403 Forbidden&#34; milliseconds=0
09:06:14: I1105 09:06:12.356102     276 wait.go:283] kube-apiserver check failed at https://10.89.0.2:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: rbac: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
09:06:14: I1105 09:06:12.855440     276 round_trippers.go:632] &#34;Response&#34; verb=&#34;GET&#34; url=&#34;https://10.89.0.2:6443/livez?timeout=10s&#34; status=&#34;500 Internal Server Error&#34; milliseconds=0
09:06:14: I1105 09:06:12.855745     276 request.go:1502] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
09:06:14: I1105 09:06:12.855795     276 wait.go:283] kube-apiserver check failed at https://10.89.0.2:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
09:06:14: I1105 09:06:13.356423     276 round_trippers.go:632] &#34;Response&#34; verb=&#34;GET&#34; url=&#34;https://10.89.0.2:6443/livez?timeout=10s&#34; status=&#34;200 OK&#34; milliseconds=1
09:06:14: [control-plane-check] kube-apiserver is healthy after 1.502411537s</pre>
</details>



* <strong>diagnostic</strong>: <code>09:06:24: I1105 09:06:16.972362     282 token.go:249] [discovery] Retrying due to error: could not find a JWS signature in the cluster-info ConfigMap for token ID &#34;abcdef&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A1081)

<details><summary>context</summary>
<pre>09:06:24: I1105 09:06:16.963598     282 envvar.go:195] &#34;Feature gate default state&#34; feature=&#34;UnlockWhileProcessingFIFO&#34; enabled=true
09:06:24: I1105 09:06:16.963944     282 token.go:229] [discovery] Waiting for the cluster-info ConfigMap to receive a JWS signature for token ID &#34;abcdef&#34;
09:06:24: I1105 09:06:16.971994     282 round_trippers.go:632] &#34;Response&#34; verb=&#34;GET&#34; url=&#34;https://kind-1.37-control-plane:6443/api/v1/namespaces/kube-public/configmaps/cluster-info?timeout=10s&#34; status=&#34;200 OK&#34; milliseconds=7
09:06:24: I1105 09:06:16.972362     282 token.go:249] [discovery] Retrying due to error: could not find a JWS signature in the cluster-info ConfigMap for token ID &#34;abcdef&#34;
09:06:24: I1105 09:06:21.964117     282 token.go:229] [discovery] Waiting for the cluster-info ConfigMap to receive a JWS signature for token ID &#34;abcdef&#34;
09:06:24: I1105 09:06:21.968475     282 round_trippers.go:632] &#34;Response&#34; verb=&#34;GET&#34; url=&#34;https://kind-1.37-control-plane:6443/api/v1/namespaces/kube-public/configmaps/cluster-info?timeout=10s&#34; status=&#34;200 OK&#34; milliseconds=4
09:06:24: I1105 09:06:21.968916     282 token.go:113] [discovery] Cluster info signature and contents are valid and no TLS pinning was specified, will use API Server &#34;kind-1.37-control-plane:6443&#34;</pre>
</details>



* <strong>secondary</strong>: <code>09:07:09: error: unable to read URL &#34;https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml&#34;, server reported 504 Gateway Timeout, status code=504</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A1349)

<details><summary>context</summary>
<pre>09:06:56: deployment.apps/cdi-operator created
09:06:56: &#43;&#43; _kubectl create -f https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml
09:06:56: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.37/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.37/.kubeconfig create -f https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml
09:07:09: error: unable to read URL &#34;https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml&#34;, server reported 504 Gateway Timeout, status code=504
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-kind-1.37-sev/2118266411544481792#1:build-log.txt%3A1350)

<details><summary>context</summary>
<pre>09:06:56: &#43;&#43; _kubectl create -f https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml
09:06:56: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.37/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.37/.kubeconfig create -f https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml
09:07:09: error: unable to read URL &#34;https://github.com/kubevirt/containerized-data-importer/releases/download/v1.63.0/cdi-cr.yaml&#34;, server reported 504 Gateway Timeout, status code=504
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>



* _2026-11-03 11:58:13 &#43;0000 UTC_: <code>12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1231)


<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>12:24:20: error: taint &#34;node-role.kubernetes.io/control-plane:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1234)

<details><summary>context</summary>
<pre>12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found
12:24:20: &#43;&#43; _kubectl taint nodes vgpu-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
12:24:20: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubeconfig taint nodes vgpu-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
12:24:20: error: taint &#34;node-role.kubernetes.io/control-plane:NoSchedule&#34; not found
12:24:20: &#43;&#43;&#43; grep -i worker
12:24:20: &#43;&#43;&#43; _get_nodes
12:24:20: &#43;&#43;&#43; awk &#39;{print $1}&#39;</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1541)

<details><summary>context</summary>
<pre>12:29:21: INFO: 2 processes: 1 internal, 1 processwrapper-sandbox.
12:29:21: INFO: Running command line: bazel-bin/example-guest-agent-copier /root/go/src/kubevirt.io/kubevirt/_out/cmd/example-guest-agent/example-guest-agent
12:29:49: &#43; rm -f /tmp/kubevirt.deploy.oDzW
make: *** [Makefile:174: cluster-sync] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1550)

<details><summary>context</summary>
<pre>12:29:50: using podman due to KIND_EXPERIMENTAL_PROVIDER
12:29:50: enabling experimental podman provider
12:29:50: Deleting cluster &#34;vgpu&#34; ...
12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125
12:30:05:
12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1553)

<details><summary>context</summary>
<pre>12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125
12:30:05:
12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1555)

<details><summary>context</summary>
<pre>12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125
&#43; true
&#43; exit 2</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1556)

<details><summary>context</summary>
<pre>12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1568)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container kind_ardinghelli in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1569)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container kind_ardinghelli in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1570)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell
================================================================================</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1571)

<details><summary>context</summary>
<pre>Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell
================================================================================
Done cleaning up after podman in container.</pre>
</details>


</details>



</details>

<hr/>
</details>
<details>
<summary> configured node not found during provisioning (2x / 22.22%) </summary>

<hr/>

**2x**: _2026-11-06 12:02:09 &#43;0000 UTC_: <code>12:18:05: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19813/pull-kubevirt-e2e-k8s-1.36-windows2016/2118675165452701696#1:build-log.txt%3A868)
<details>
<summary>all...</summary>

* _2026-11-06 12:02:09 &#43;0000 UTC_: <code>12:18:05: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19813/pull-kubevirt-e2e-k8s-1.36-windows2016/2118675165452701696#1:build-log.txt%3A868)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>12:18:02: I1106 07:18:02.507609    1888 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19813/pull-kubevirt-e2e-k8s-1.36-windows2016/2118675165452701696#1:build-log.txt%3A805)

<details><summary>context</summary>
<pre>12:18:01: [control-plane-check] kube-controller-manager is healthy after 5.239285ms
12:18:01: [control-plane-check] kube-scheduler is healthy after 7.898644ms
12:18:02: I1106 07:18:02.008369    1888 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
12:18:02: I1106 07:18:02.507609    1888 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:18:02: I1106 07:18:02.507711    1888 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:18:02: [control-plane-check] kube-apiserver is healthy after 1.503624779s
12:18:02: I1106 07:18:03.008716    1888 uploadconfig.go:109] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19813/pull-kubevirt-e2e-k8s-1.36-windows2016/2118675165452701696#1:build-log.txt%3A1150)

<details><summary>context</summary>
<pre>12:18:17: 
12:18:17: &#43; wait 69
12:18:29: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19813/pull-kubevirt-e2e-k8s-1.36-windows2016/2118675165452701696#1:build-log.txt%3A1178)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:18:30: selecting podman as container runtime
12:18:30: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-windows2016 is being used by the following container(s): afe66bfd0ea325b2aedd124aef268e5e6cb650d1d5ccd698ca0210fa9d12b1b3: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



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
<details>
<summary> make cluster lifecycle target failure (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-11-03 11:58:41 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A3943)
<details>
<summary>all...</summary>

* _2026-11-03 11:58:41 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A3943)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>12:15:30: I1103 07:15:30.373979    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A798)

<details><summary>context</summary>
<pre>12:15:21: [control-plane-check] Checking kube-scheduler at https://127.0.0.1:10259/livez
12:15:25: [control-plane-check] kube-scheduler is healthy after 3.842462827s
12:15:26: [control-plane-check] kube-controller-manager is healthy after 5.068241346s
12:15:30: I1103 07:15:30.373979    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:15:30: [control-plane-check] kube-apiserver is healthy after 9.002895106s
12:15:30: I1103 07:15:30.875756    1605 kubeconfig.go:657] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists
12:15:30: I1103 07:15:30.878733    1605 kubeconfig.go:730] creating the ClusterRoleBinding for the kubeadm:cluster-admins Group by using super-admin.conf</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A3949)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:27:34: selecting podman as container runtime
12:27:49: Error response from daemon: cannot remove container 4d2bdd9539b64bcf49fcea3d1d98ac72c4d24a6d0e4c0180ba3525a5a560ec7c as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

#### external (4x / 44.44%)

<details>
<summary> download failure from external URL (2x / 22.22%) </summary>

<hr/>

**1x**: _2026-11-05 14:33:27 &#43;0000 UTC_: <code>14:36:41: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp6849394626262802922/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A276)
<details>
<summary>all...</summary>

* _2026-11-05 14:33:27 &#43;0000 UTC_: <code>14:36:41: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp6849394626262802922/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A276)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>14:34:53: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_cc&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A218)

<details><summary>context</summary>
<pre>14:34:53:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/io_bazel_rules_go/go/private/repositories.bzl:355:18: in _maybe
14:34:53: Repository rule http_archive defined at:
14:34:53:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
14:34:53: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_cc&#39;:
14:34:53:    Traceback (most recent call last):
14:34:53: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
14:34:53: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>14:34:53: ERROR: Skipping &#39;//:gazelle&#39;: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A224)

<details><summary>context</summary>
<pre>14:34:53: 		download_info = ctx.download_and_extract(
14:34:53: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: WARNING: Target pattern parsing failed.
14:34:53: ERROR: Skipping &#39;//:gazelle&#39;: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: ERROR: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: INFO: Elapsed time: 49.216s
14:34:53: INFO: 0 processes.</pre>
</details>



* <strong>diagnostic</strong>: <code>14:34:53: ERROR: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A225)

<details><summary>context</summary>
<pre>14:34:53: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: WARNING: Target pattern parsing failed.
14:34:53: ERROR: Skipping &#39;//:gazelle&#39;: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: ERROR: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: INFO: Elapsed time: 49.216s
14:34:53: INFO: 0 processes.
14:34:53: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>14:34:53: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A228)

<details><summary>context</summary>
<pre>14:34:53: ERROR: no such package &#39;@@rules_cc//cc/common&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_cc/releases/download/0.1.5/rules_cc-0.1.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_cc/temp7421451160340069632/rules_cc-0.1.5.tar.gz: GET returned 504 Gateway Time-out
14:34:53: INFO: Elapsed time: 49.216s
14:34:53: INFO: 0 processes.
14:34:53: ERROR: Build did NOT complete successfully
14:34:53: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>14:34:53: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A229)

<details><summary>context</summary>
<pre>14:34:53: INFO: Elapsed time: 49.216s
14:34:53: INFO: 0 processes.
14:34:53: ERROR: Build did NOT complete successfully
14:34:53: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A230)

<details><summary>context</summary>
<pre>14:34:53: INFO: 0 processes.
14:34:53: ERROR: Build did NOT complete successfully
14:34:53: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>diagnostic</strong>: <code>14:36:03: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;apple_support~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A245)

<details><summary>context</summary>
<pre>14:36:03:   &lt;builtin&gt;: in &lt;toplevel&gt;
14:36:03: Repository rule http_archive defined at:
14:36:03:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
14:36:03: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;apple_support~&#39;:
14:36:03:    Traceback (most recent call last):
14:36:03: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
14:36:03: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>14:36:03: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp1387829800134064954/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A250)

<details><summary>context</summary>
<pre>14:36:03: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
14:36:03: 		download_info = ctx.download_and_extract(
14:36:03: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp1387829800134064954/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:03: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp1387829800134064954/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:03: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
14:36:03: INFO: Elapsed time: 48.439s, Critical Path: 0.60s
14:36:03: INFO: 1 process: 1 internal.</pre>
</details>



* <strong>diagnostic</strong>: <code>14:36:03: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A251)

<details><summary>context</summary>
<pre>14:36:03: 		download_info = ctx.download_and_extract(
14:36:03: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp1387829800134064954/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:03: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp1387829800134064954/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:03: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
14:36:03: INFO: Elapsed time: 48.439s, Critical Path: 0.60s
14:36:03: INFO: 1 process: 1 internal.
14:36:03: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>14:36:03: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A254)

<details><summary>context</summary>
<pre>14:36:03: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
14:36:03: INFO: Elapsed time: 48.439s, Critical Path: 0.60s
14:36:03: INFO: 1 process: 1 internal.
14:36:03: ERROR: Build did NOT complete successfully
14:36:03: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>14:36:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A255)

<details><summary>context</summary>
<pre>14:36:03: INFO: Elapsed time: 48.439s, Critical Path: 0.60s
14:36:03: INFO: 1 process: 1 internal.
14:36:03: ERROR: Build did NOT complete successfully
14:36:03: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A256)

<details><summary>context</summary>
<pre>14:36:03: INFO: 1 process: 1 internal.
14:36:03: ERROR: Build did NOT complete successfully
14:36:03: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>secondary</strong>: <code>14:36:41: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;apple_support~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A271)

<details><summary>context</summary>
<pre>14:36:41:   &lt;builtin&gt;: in &lt;toplevel&gt;
14:36:41: Repository rule http_archive defined at:
14:36:41:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
14:36:41: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;apple_support~&#39;:
14:36:41:    Traceback (most recent call last):
14:36:41: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
14:36:41: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>secondary</strong>: <code>14:36:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A277)

<details><summary>context</summary>
<pre>14:36:41: 		download_info = ctx.download_and_extract(
14:36:41: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp6849394626262802922/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:41: ERROR: /root/go/src/kubevirt.io/kubevirt/BUILD.bazel:104:8: //:gazelle depends on //:gazelle-runner in repository @@ which failed to fetch. no such package &#39;@@apple_support~//crosstool&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/apple_support~/temp6849394626262802922/apple_support.1.5.0.tar.gz: GET returned 504 Gateway Time-out
14:36:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
14:36:41: INFO: Elapsed time: 13.016s, Critical Path: 0.60s
14:36:41: INFO: 1 process: 1 internal.
14:36:41: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>secondary</strong>: <code>14:36:41: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A280)

<details><summary>context</summary>
<pre>14:36:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
14:36:41: INFO: Elapsed time: 13.016s, Critical Path: 0.60s
14:36:41: INFO: 1 process: 1 internal.
14:36:41: ERROR: Build did NOT complete successfully
14:36:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>secondary</strong>: <code>14:36:41: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A281)

<details><summary>context</summary>
<pre>14:36:41: INFO: Elapsed time: 13.016s, Critical Path: 0.60s
14:36:41: INFO: 1 process: 1 internal.
14:36:41: ERROR: Build did NOT complete successfully
14:36:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A282)

<details><summary>context</summary>
<pre>14:36:41: INFO: 1 process: 1 internal.
14:36:41: ERROR: Build did NOT complete successfully
14:36:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19707/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118329433034067968#1:build-log.txt%3A293)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
14:36:47: selecting podman as container runtime
14:36:47: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-operator is being used by the following container(s): 9311f6c2ebab024ca15294aafe5d852280dc712623b4835614575fb0de781fe0: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>

**1x**: _2026-11-05 08:58:16 &#43;0000 UTC_: <code>09:01:23: ERROR: no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A276)
<details>
<summary>all...</summary>

* _2026-11-05 08:58:16 &#43;0000 UTC_: <code>09:01:23: ERROR: no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A276)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>08:59:46: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A221)

<details><summary>context</summary>
<pre>08:59:46:   &lt;builtin&gt;: in &lt;toplevel&gt;
08:59:46: Repository rule http_archive defined at:
08:59:46:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
08:59:46: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:
08:59:46:    Traceback (most recent call last):
08:59:46: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:46: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:46: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp18112639366700063061/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A226)

<details><summary>context</summary>
<pre>08:59:46: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:46: 		download_info = ctx.download_and_extract(
08:59:46: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp18112639366700063061/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:46: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp18112639366700063061/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:46: INFO: Elapsed time: 38.429s, Critical Path: 0.60s
08:59:46: INFO: 1 process: 1 internal.
08:59:46: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:46: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A229)

<details><summary>context</summary>
<pre>08:59:46: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp18112639366700063061/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:46: INFO: Elapsed time: 38.429s, Critical Path: 0.60s
08:59:46: INFO: 1 process: 1 internal.
08:59:46: ERROR: Build did NOT complete successfully
08:59:46: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:46: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A230)

<details><summary>context</summary>
<pre>08:59:46: INFO: Elapsed time: 38.429s, Critical Path: 0.60s
08:59:46: INFO: 1 process: 1 internal.
08:59:46: ERROR: Build did NOT complete successfully
08:59:46: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A231)

<details><summary>context</summary>
<pre>08:59:46: INFO: 1 process: 1 internal.
08:59:46: ERROR: Build did NOT complete successfully
08:59:46: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>diagnostic</strong>: <code>09:00:04: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A246)

<details><summary>context</summary>
<pre>09:00:04:   &lt;builtin&gt;: in &lt;toplevel&gt;
09:00:04: Repository rule http_archive defined at:
09:00:04:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
09:00:04: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:
09:00:04:    Traceback (most recent call last):
09:00:04: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
09:00:04: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>09:00:04: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp2637779128729034844/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A251)

<details><summary>context</summary>
<pre>09:00:04: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
09:00:04: 		download_info = ctx.download_and_extract(
09:00:04: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp2637779128729034844/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:04: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp2637779128729034844/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:04: INFO: Elapsed time: 13.746s, Critical Path: 0.53s
09:00:04: INFO: 1 process: 1 internal.
09:00:04: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>09:00:04: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A254)

<details><summary>context</summary>
<pre>09:00:04: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp2637779128729034844/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:04: INFO: Elapsed time: 13.746s, Critical Path: 0.53s
09:00:04: INFO: 1 process: 1 internal.
09:00:04: ERROR: Build did NOT complete successfully
09:00:04: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>09:00:04: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A255)

<details><summary>context</summary>
<pre>09:00:04: INFO: Elapsed time: 13.746s, Critical Path: 0.53s
09:00:04: INFO: 1 process: 1 internal.
09:00:04: ERROR: Build did NOT complete successfully
09:00:04: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A256)

<details><summary>context</summary>
<pre>09:00:04: INFO: 1 process: 1 internal.
09:00:04: ERROR: Build did NOT complete successfully
09:00:04: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_python~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A271)

<details><summary>context</summary>
<pre>09:01:23:   &lt;builtin&gt;: in &lt;toplevel&gt;
09:01:23: Repository rule http_archive defined at:
09:01:23:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_python~&#39;:
09:01:23:    Traceback (most recent call last):
09:01:23: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
09:01:23: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/toolchains/BUILD:274:27: @@rules_java~//toolchains:toolchain_java10_definition depends on @@bazel_tools//tools/jdk:toolchain_type in repository @@bazel_tools which failed to fetch. no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A277)

<details><summary>context</summary>
<pre>09:01:23: 		download_info = ctx.download_and_extract(
09:01:23: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/toolchains/BUILD:274:27: @@rules_java~//toolchains:toolchain_java10_definition depends on @@bazel_tools//tools/jdk:toolchain_type in repository @@bazel_tools which failed to fetch. no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/src/tools/launcher/BUILD:10:14: Analysis failed
09:01:23: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
09:01:23: INFO: Elapsed time: 67.819s, Critical Path: 0.61s</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/src/tools/launcher/BUILD:10:14: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A278)

<details><summary>context</summary>
<pre>09:01:23: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/toolchains/BUILD:274:27: @@rules_java~//toolchains:toolchain_java10_definition depends on @@bazel_tools//tools/jdk:toolchain_type in repository @@bazel_tools which failed to fetch. no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/src/tools/launcher/BUILD:10:14: Analysis failed
09:01:23: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
09:01:23: INFO: Elapsed time: 67.819s, Critical Path: 0.61s
09:01:23: INFO: 1 process: 1 internal.</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A279)

<details><summary>context</summary>
<pre>09:01:23: ERROR: no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/toolchains/BUILD:274:27: @@rules_java~//toolchains:toolchain_java10_definition depends on @@bazel_tools//tools/jdk:toolchain_type in repository @@bazel_tools which failed to fetch. no such package &#39;@@rules_python~//python&#39;: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_python/releases/download/0.22.1/rules_python-0.22.1.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_python~/temp11516252301565034293/rules_python-0.22.1.tar.gz: GET returned 504 Gateway Time-out
09:01:23: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/src/tools/launcher/BUILD:10:14: Analysis failed
09:01:23: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
09:01:23: INFO: Elapsed time: 67.819s, Critical Path: 0.61s
09:01:23: INFO: 1 process: 1 internal.
09:01:23: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A282)

<details><summary>context</summary>
<pre>09:01:23: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: Analysis failed
09:01:23: INFO: Elapsed time: 67.819s, Critical Path: 0.61s
09:01:23: INFO: 1 process: 1 internal.
09:01:23: ERROR: Build did NOT complete successfully
09:01:23: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>secondary</strong>: <code>09:01:23: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A283)

<details><summary>context</summary>
<pre>09:01:23: INFO: Elapsed time: 67.819s, Critical Path: 0.61s
09:01:23: INFO: 1 process: 1 internal.
09:01:23: ERROR: Build did NOT complete successfully
09:01:23: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A284)

<details><summary>context</summary>
<pre>09:01:23: INFO: 1 process: 1 internal.
09:01:23: ERROR: Build did NOT complete successfully
09:01:23: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-operator/2118266411066331136#1:build-log.txt%3A295)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
09:02:10: selecting podman as container runtime
09:02:11: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-operator is being used by the following container(s): 8461bb1b9368374a39c102452c3ac4240460428f7aab3a0411def0f08de4de90: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>
<details>
<summary> container image pull failure (1x / 11.11%) </summary>

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
<details>
<summary> bazel repository fetch failure (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-11-05 08:58:00 &#43;0000 UTC_: <code>09:00:17: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A266)
<details>
<summary>all...</summary>

* _2026-11-05 08:58:00 &#43;0000 UTC_: <code>09:00:17: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A266)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>08:59:40: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A216)

<details><summary>context</summary>
<pre>08:59:40:   &lt;builtin&gt;: in &lt;toplevel&gt;
08:59:40: Repository rule http_archive defined at:
08:59:40:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
08:59:40: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:
08:59:40:    Traceback (most recent call last):
08:59:40: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:40: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp10655117106702680376/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A221)

<details><summary>context</summary>
<pre>08:59:40: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:40: 		download_info = ctx.download_and_extract(
08:59:40: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp10655117106702680376/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp10655117106702680376/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:41: INFO: Elapsed time: 57.381s, Critical Path: 0.64s
08:59:41: INFO: 1 process: 1 internal.
08:59:41: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:41: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A224)

<details><summary>context</summary>
<pre>08:59:41: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp10655117106702680376/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:41: INFO: Elapsed time: 57.381s, Critical Path: 0.64s
08:59:41: INFO: 1 process: 1 internal.
08:59:41: ERROR: Build did NOT complete successfully
08:59:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:41: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A225)

<details><summary>context</summary>
<pre>08:59:41: INFO: Elapsed time: 57.381s, Critical Path: 0.64s
08:59:41: INFO: 1 process: 1 internal.
08:59:41: ERROR: Build did NOT complete successfully
08:59:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A226)

<details><summary>context</summary>
<pre>08:59:41: INFO: 1 process: 1 internal.
08:59:41: ERROR: Build did NOT complete successfully
08:59:41: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:59: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A241)

<details><summary>context</summary>
<pre>08:59:59:   &lt;builtin&gt;: in &lt;toplevel&gt;
08:59:59: Repository rule http_archive defined at:
08:59:59:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in &lt;toplevel&gt;
08:59:59: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:136:45: An error occurred during the fetch of repository &#39;rules_java~&#39;:
08:59:59:    Traceback (most recent call last):
08:59:59: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:59: 		download_info = ctx.download_and_extract(</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:59: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp1347094595764469151/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A246)

<details><summary>context</summary>
<pre>08:59:59: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
08:59:59: 		download_info = ctx.download_and_extract(
08:59:59: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp1347094595764469151/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:59: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp1347094595764469151/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:59: INFO: Elapsed time: 13.436s, Critical Path: 0.58s
08:59:59: INFO: 1 process: 1 internal.
08:59:59: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:59: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A249)

<details><summary>context</summary>
<pre>08:59:59: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp1347094595764469151/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
08:59:59: INFO: Elapsed time: 13.436s, Critical Path: 0.58s
08:59:59: INFO: 1 process: 1 internal.
08:59:59: ERROR: Build did NOT complete successfully
08:59:59: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>diagnostic</strong>: <code>08:59:59: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A250)

<details><summary>context</summary>
<pre>08:59:59: INFO: Elapsed time: 13.436s, Critical Path: 0.58s
08:59:59: INFO: 1 process: 1 internal.
08:59:59: ERROR: Build did NOT complete successfully
08:59:59: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)</pre>
</details>



* <strong>diagnostic</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A251)

<details><summary>context</summary>
<pre>08:59:59: INFO: 1 process: 1 internal.
08:59:59: ERROR: Build did NOT complete successfully
08:59:59: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; for i in $(seq 1 $tries)
&#43; make bazel-build-images</pre>
</details>



* <strong>secondary</strong>: <code>09:00:17: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp7416242368100028705/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A271)

<details><summary>context</summary>
<pre>09:00:17: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 136, column 45, in _http_archive_impl
09:00:17: 		download_info = ctx.download_and_extract(
09:00:17: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp7416242368100028705/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:17: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp7416242368100028705/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:17: INFO: Elapsed time: 13.699s, Critical Path: 0.65s
09:00:17: INFO: 1 process: 1 internal.
09:00:17: ERROR: Build did NOT complete successfully</pre>
</details>



* <strong>secondary</strong>: <code>09:00:17: ERROR: Build did NOT complete successfully</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A274)

<details><summary>context</summary>
<pre>09:00:17: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted: java.io.IOException: Error downloading [https://github.com/bazelbuild/rules_java/releases/download/7.6.5/rules_java-7.6.5.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_java~/temp7416242368100028705/rules_java-7.6.5.tar.gz: GET returned 504 Gateway Time-out
09:00:17: INFO: Elapsed time: 13.699s, Critical Path: 0.65s
09:00:17: INFO: 1 process: 1 internal.
09:00:17: ERROR: Build did NOT complete successfully
09:00:17: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2</pre>
</details>



* <strong>secondary</strong>: <code>09:00:17: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A275)

<details><summary>context</summary>
<pre>09:00:17: INFO: Elapsed time: 13.699s, Critical Path: 0.65s
09:00:17: INFO: 1 process: 1 internal.
09:00:17: ERROR: Build did NOT complete successfully
09:00:17: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A276)

<details><summary>context</summary>
<pre>09:00:17: INFO: 1 process: 1 internal.
09:00:17: ERROR: Build did NOT complete successfully
09:00:17: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19585/pull-kubevirt-e2e-k8s-1.36-sig-compute/2118266410957279232#1:build-log.txt%3A287)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
09:00:26: selecting podman as container runtime
09:00:26: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute is being used by the following container(s): c8d4c791d2f91c3a56c506d6b5653b81a2c82208835724ffd468f4be28ac67c2: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

### sig-monitoring (1x / 8.33%)


#### needs-investigation (1x / 100.00%)

<details>
<summary> no error snippets (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-11-04 07:13:51 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19748/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2117877790450651136)
<details>
<summary>all...</summary>

* _2026-11-04 07:13:51 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19748/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2117877790450651136)




</details>

<hr/>
</details>

### sig-network (1x / 8.33%)


#### internal (1x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-11-03 12:09:09 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2117587001577312256#1:build-log.txt%3A4051)
<details>
<summary>all...</summary>

* _2026-11-03 12:09:09 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2117587001577312256#1:build-log.txt%3A4051)


<details>
<summary>secondary evidence</summary>


* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2117587001577312256#1:build-log.txt%3A4057)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:30:29: selecting podman as container runtime
12:32:15: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.33-sig-network-1.7 is being used by the following container(s): 437e990e2003afcfa461cefe139c5de3abcd06e6208ef33c9e77cb53b350a614, d7746eb1230afd44770adaf9ee2a421590a3d28144b9974b1f14b2ea60f6311e: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

### sig-storage (1x / 8.33%)


#### internal (1x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-11-03 11:59:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2117586999056535552#1:build-log.txt%3A1217)
<details>
<summary>all...</summary>

* _2026-11-03 11:59:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2117586999056535552#1:build-log.txt%3A1217)


<details>
<summary>secondary evidence</summary>


* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2117586999056535552#1:build-log.txt%3A1223)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:36:37: selecting podman as container runtime
12:36:52: Error response from daemon: cannot remove container 09207706028f593ad43cb2e4bf5ea652f5970c6efdf3a476d3f2ee015494902d as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

Last updated: 2026-11-06 21:15:36
