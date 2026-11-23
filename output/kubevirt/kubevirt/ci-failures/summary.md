





<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-11-21 (1x / 33.33%)


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

### 2026-11-17 (2x / 66.67%)


#### external (2x / 100.00%)

<details>
<summary> container image pull failure (2x / 100.00%) </summary>

<hr/>

**1x**: _2026-11-17 07:12:58 &#43;0000 UTC_: <code>07:19:39: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage2707048589/2&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/5c/5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T071939Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e8cf4826f6d12c45abde65ca089e80040721b710854a7e9d3935017927358cc8&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1794900879~hmac=96c9e138c0d821d123b2ec5af3093d15468523626067d8ff236643746e84ed98&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19932/pull-kubevirt-e2e-kind-1.37-vgpu/2122588605732884480#1:build-log.txt%3A456)
<details>
<summary>all...</summary>

* _2026-11-17 07:12:58 &#43;0000 UTC_: <code>07:19:39: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage2707048589/2&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/5c/5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T071939Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e8cf4826f6d12c45abde65ca089e80040721b710854a7e9d3935017927358cc8&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1794900879~hmac=96c9e138c0d821d123b2ec5af3093d15468523626067d8ff236643746e84ed98&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19932/pull-kubevirt-e2e-kind-1.37-vgpu/2122588605732884480#1:build-log.txt%3A456)


<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19932/pull-kubevirt-e2e-kind-1.37-vgpu/2122588605732884480#1:build-log.txt%3A457)

<details><summary>context</summary>
<pre>07:19:37: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
07:19:37: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
07:19:39: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage2707048589/2&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/5c/5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T071939Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e8cf4826f6d12c45abde65ca089e80040721b710854a7e9d3935017927358cc8&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1794900879~hmac=96c9e138c0d821d123b2ec5af3093d15468523626067d8ff236643746e84ed98&#34;: EOF)
make: *** [Makefile:179: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>



</details>

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

<a id="per-error-category"></a>

## per error category [⬆](#top)


### needs-investigation (1x / 33.33%)

<details>
<summary> no matching pattern (1x / 33.33%) </summary>

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

### external (2x / 66.67%)

<details>
<summary> container image pull failure (2x / 66.67%) </summary>

<hr/>

**1x**: _2026-11-17 07:12:58 &#43;0000 UTC_: <code>07:19:39: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage2707048589/2&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/5c/5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T071939Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e8cf4826f6d12c45abde65ca089e80040721b710854a7e9d3935017927358cc8&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1794900879~hmac=96c9e138c0d821d123b2ec5af3093d15468523626067d8ff236643746e84ed98&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19932/pull-kubevirt-e2e-kind-1.37-vgpu/2122588605732884480#1:build-log.txt%3A456)
<details>
<summary>all...</summary>

* _2026-11-17 07:12:58 &#43;0000 UTC_: <code>07:19:39: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage2707048589/2&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/5c/5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T071939Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e8cf4826f6d12c45abde65ca089e80040721b710854a7e9d3935017927358cc8&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1794900879~hmac=96c9e138c0d821d123b2ec5af3093d15468523626067d8ff236643746e84ed98&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19932/pull-kubevirt-e2e-kind-1.37-vgpu/2122588605732884480#1:build-log.txt%3A456)
<details><summary>context</summary>
<pre>07:19:37: Getting image source signatures
07:19:37: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
07:19:37: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
07:19:39: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage2707048589/2&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/5c/5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T071939Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e8cf4826f6d12c45abde65ca089e80040721b710854a7e9d3935017927358cc8&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1794900879~hmac=96c9e138c0d821d123b2ec5af3093d15468523626067d8ff236643746e84ed98&#34;: EOF)
make: *** [Makefile:179: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19932/pull-kubevirt-e2e-kind-1.37-vgpu/2122588605732884480#1:build-log.txt%3A457)

<details><summary>context</summary>
<pre>07:19:37: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
07:19:37: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
07:19:39: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage2707048589/2&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/5c/5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T071939Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e8cf4826f6d12c45abde65ca089e80040721b710854a7e9d3935017927358cc8&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1794900879~hmac=96c9e138c0d821d123b2ec5af3093d15468523626067d8ff236643746e84ed98&#34;: EOF)
make: *** [Makefile:179: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>



</details>

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


### main (3x / 100.00%)


#### needs-investigation (1x / 33.33%)

<details>
<summary> no matching pattern (1x / 33.33%) </summary>

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

#### external (2x / 66.67%)

<details>
<summary> container image pull failure (2x / 66.67%) </summary>

<hr/>

**1x**: _2026-11-17 07:12:58 &#43;0000 UTC_: <code>07:19:39: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage2707048589/2&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/5c/5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T071939Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e8cf4826f6d12c45abde65ca089e80040721b710854a7e9d3935017927358cc8&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1794900879~hmac=96c9e138c0d821d123b2ec5af3093d15468523626067d8ff236643746e84ed98&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19932/pull-kubevirt-e2e-kind-1.37-vgpu/2122588605732884480#1:build-log.txt%3A456)
<details>
<summary>all...</summary>

* _2026-11-17 07:12:58 &#43;0000 UTC_: <code>07:19:39: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage2707048589/2&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/5c/5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T071939Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e8cf4826f6d12c45abde65ca089e80040721b710854a7e9d3935017927358cc8&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1794900879~hmac=96c9e138c0d821d123b2ec5af3093d15468523626067d8ff236643746e84ed98&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19932/pull-kubevirt-e2e-kind-1.37-vgpu/2122588605732884480#1:build-log.txt%3A456)


<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19932/pull-kubevirt-e2e-kind-1.37-vgpu/2122588605732884480#1:build-log.txt%3A457)

<details><summary>context</summary>
<pre>07:19:37: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
07:19:37: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
07:19:39: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage2707048589/2&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/5c/5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T071939Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e8cf4826f6d12c45abde65ca089e80040721b710854a7e9d3935017927358cc8&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1794900879~hmac=96c9e138c0d821d123b2ec5af3093d15468523626067d8ff236643746e84ed98&#34;: EOF)
make: *** [Makefile:179: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>



</details>

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

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-storage (1x / 33.33%)


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

### sig-compute (2x / 66.67%)


#### external (2x / 100.00%)

<details>
<summary> container image pull failure (2x / 100.00%) </summary>

<hr/>

**1x**: _2026-11-17 07:12:58 &#43;0000 UTC_: <code>07:19:39: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage2707048589/2&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/5c/5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T071939Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e8cf4826f6d12c45abde65ca089e80040721b710854a7e9d3935017927358cc8&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1794900879~hmac=96c9e138c0d821d123b2ec5af3093d15468523626067d8ff236643746e84ed98&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19932/pull-kubevirt-e2e-kind-1.37-vgpu/2122588605732884480#1:build-log.txt%3A456)
<details>
<summary>all...</summary>

* _2026-11-17 07:12:58 &#43;0000 UTC_: <code>07:19:39: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage2707048589/2&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/5c/5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T071939Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e8cf4826f6d12c45abde65ca089e80040721b710854a7e9d3935017927358cc8&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1794900879~hmac=96c9e138c0d821d123b2ec5af3093d15468523626067d8ff236643746e84ed98&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19932/pull-kubevirt-e2e-kind-1.37-vgpu/2122588605732884480#1:build-log.txt%3A456)


<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19932/pull-kubevirt-e2e-kind-1.37-vgpu/2122588605732884480#1:build-log.txt%3A457)

<details><summary>context</summary>
<pre>07:19:37: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
07:19:37: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
07:19:39: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage2707048589/2&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/5c/5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261117%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261117T071939Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=e8cf4826f6d12c45abde65ca089e80040721b710854a7e9d3935017927358cc8&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1794900879~hmac=96c9e138c0d821d123b2ec5af3093d15468523626067d8ff236643746e84ed98&#34;: EOF)
make: *** [Makefile:179: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>



</details>

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

Last updated: 2026-11-23 00:09:30
