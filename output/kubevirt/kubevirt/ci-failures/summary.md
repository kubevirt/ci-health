



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-06-18 (1x / 3.85%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-06-18 08:42:15 &#43;0000 UTC_: <code>08:43:22: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17263/pull-kubevirt-e2e-k8s-1.34-sig-performance/2067528172320067584#1:build-log.txt%3A86)
<details>
<summary>all...</summary>

* _2026-06-18 08:42:15 &#43;0000 UTC_: <code>08:43:22: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17263/pull-kubevirt-e2e-k8s-1.34-sig-performance/2067528172320067584#1:build-log.txt%3A86)

</details>

<hr/>
</details>

### 2026-06-17 (1x / 3.85%)


#### external (1x / 100.00%)

<details>
<summary> podman container removal timeout (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-06-17 07:01:38 &#43;0000 UTC_: <code>Error: cannot remove container e672de5d181574bc76d99016452be15084c22ab8f3d5895e5e5f8d0a986dc6f7 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17976/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2067140165721133056#1:build-log.txt%3A1868)
<details>
<summary>all...</summary>

* _2026-06-17 07:01:38 &#43;0000 UTC_: <code>Error: cannot remove container e672de5d181574bc76d99016452be15084c22ab8f3d5895e5e5f8d0a986dc6f7 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17976/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2067140165721133056#1:build-log.txt%3A1868)

</details>

<hr/>
</details>

### 2026-06-16 (2x / 7.69%)


#### external (1x / 50.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-06-16 12:46:25 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2066864704415862784#1:build-log.txt%3A1383)
<details>
<summary>all...</summary>

* _2026-06-16 12:46:25 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2066864704415862784#1:build-log.txt%3A1383)

</details>

<hr/>
</details>

#### internal (1x / 50.00%)

<details>
<summary> kind cluster creation failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-06-16 11:48:08 &#43;0000 UTC_: <code>11:55:03: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-kind-1.34-sev/2066849993645363200#1:build-log.txt%3A729)
<details>
<summary>all...</summary>

* _2026-06-16 11:48:08 &#43;0000 UTC_: <code>11:55:03: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-kind-1.34-sev/2066849993645363200#1:build-log.txt%3A729)

</details>

<hr/>
</details>

### 2026-06-14 (6x / 23.08%)


#### external (2x / 33.33%)

<details>
<summary> download failure from external URL (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-06-14 06:31:19 &#43;0000 UTC_: <code>06:32:23: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp3938845111312673746/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17913/pull-kubevirt-e2e-k8s-1.35-sig-storage/2066045650159538176#1:build-log.txt%3A311)
<details>
<summary>all...</summary>

* _2026-06-14 06:31:19 &#43;0000 UTC_: <code>06:32:23: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp3938845111312673746/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17913/pull-kubevirt-e2e-k8s-1.35-sig-storage/2066045650159538176#1:build-log.txt%3A311)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-06-14 17:57:20 &#43;0000 UTC_: <code>18:13:46: I0614 14:13:46.824946    1620 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18122/pull-kubevirt-e2e-k8s-1.35-sig-operator-1.8/2066180130513883136#1:build-log.txt%3A768)
<details>
<summary>all...</summary>

* _2026-06-14 17:57:20 &#43;0000 UTC_: <code>18:13:46: I0614 14:13:46.824946    1620 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18122/pull-kubevirt-e2e-k8s-1.35-sig-operator-1.8/2066180130513883136#1:build-log.txt%3A768)

</details>

<hr/>
</details>

#### internal (2x / 33.33%)

<details>
<summary> make cluster lifecycle target failure (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-06-14 14:43:50 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18130/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2066169290440052736#1:build-log.txt%3A629)
<details>
<summary>all...</summary>

* _2026-06-14 14:43:50 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18130/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2066169290440052736#1:build-log.txt%3A629)

</details>

<hr/>
</details>
<details>
<summary> kind cluster creation failure (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-06-14 11:18:29 &#43;0000 UTC_: <code>11:24:59: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18125/pull-kubevirt-e2e-kind-1.34-sev-1.8/2066117945167187968#1:build-log.txt%3A789)
<details>
<summary>all...</summary>

* _2026-06-14 11:18:29 &#43;0000 UTC_: <code>11:24:59: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18125/pull-kubevirt-e2e-kind-1.34-sev-1.8/2066117945167187968#1:build-log.txt%3A789)

</details>

<hr/>
</details>

#### needs-investigation (2x / 33.33%)

<details>
<summary> no error snippets (2x / 33.33%) </summary>

<hr/>

**1x**: _2026-06-14 11:27:08 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18125/pull-kubevirt-e2e-k8s-1.34-sig-network-1.8/2066116448811487232)
<details>
<summary>all...</summary>

* _2026-06-14 11:27:08 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18125/pull-kubevirt-e2e-k8s-1.34-sig-network-1.8/2066116448811487232)

</details>

<hr/>

**1x**: _2026-06-14 10:58:52 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18121/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network-1.8/2066112953937039360)
<details>
<summary>all...</summary>

* _2026-06-14 10:58:52 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18121/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network-1.8/2066112953937039360)

</details>

<hr/>
</details>

### 2026-06-13 (2x / 7.69%)


#### needs-investigation (2x / 100.00%)

<details>
<summary> no matching pattern (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-06-13 11:08:07 &#43;0000 UTC_: <code>11:09:44: ERROR: Error initializing RemoteModule</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18034/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065748056631939072#1:build-log.txt%3A347)
<details>
<summary>all...</summary>

* _2026-06-13 11:38:34 &#43;0000 UTC_: <code>11:39:30: ERROR: Error initializing RemoteModule</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17689/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065760587530375168#1:build-log.txt%3A355)

* _2026-06-13 11:08:07 &#43;0000 UTC_: <code>11:09:44: ERROR: Error initializing RemoteModule</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18034/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065748056631939072#1:build-log.txt%3A347)

</details>

<hr/>
</details>

### 2026-06-12 (14x / 53.85%)


#### external (14x / 100.00%)

<details>
<summary> bazel remote cache blob fetch failure (10x / 71.43%) </summary>

<hr/>

**5x**: _2026-06-12 09:02:29 &#43;0000 UTC_: <code>09:05:24: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-storage/2065358924223090688#1:build-log.txt%3A398)
<details>
<summary>all...</summary>

* _2026-06-12 09:02:31 &#43;0000 UTC_: <code>09:07:05: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2065358924449583104#1:build-log.txt%3A500)

* _2026-06-12 09:02:30 &#43;0000 UTC_: <code>09:05:29: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-operator/2065358924541857792#1:build-log.txt%3A430)

* _2026-06-12 09:02:29 &#43;0000 UTC_: <code>09:05:24: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-storage/2065358924223090688#1:build-log.txt%3A398)

* _2026-06-12 09:02:28 &#43;0000 UTC_: <code>09:05:43: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-network/2065358924130816000#1:build-log.txt%3A405)

* _2026-06-12 09:02:25 &#43;0000 UTC_: <code>09:05:22: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-compute/2065358924294393856#1:build-log.txt%3A390)

</details>

<hr/>

**2x**: _2026-06-12 20:31:19 &#43;0000 UTC_: <code>20:36:05: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/compute/BUILD.bazel:3:11: GoCompilePkg tests/compute/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065532284278149120#1:build-log.txt%3A598)
<details>
<summary>all...</summary>

* _2026-06-12 20:31:19 &#43;0000 UTC_: <code>20:36:05: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/compute/BUILD.bazel:3:11: GoCompilePkg tests/compute/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065532284278149120#1:build-log.txt%3A598)

* _2026-06-12 13:39:05 &#43;0000 UTC_: <code>13:43:10: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/compute/BUILD.bazel:3:11: GoCompilePkg tests/compute/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065428547740110848#1:build-log.txt%3A601)

</details>

<hr/>

**2x**: _2026-06-12 08:10:14 &#43;0000 UTC_: <code>08:14:17: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/validatingadmissionpolicy/BUILD.bazel:3:11: GoCompilePkg tests/validatingadmissionpolicy/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065345254365597696#1:build-log.txt%3A522)
<details>
<summary>all...</summary>

* _2026-06-12 08:10:14 &#43;0000 UTC_: <code>08:14:17: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/validatingadmissionpolicy/BUILD.bazel:3:11: GoCompilePkg tests/validatingadmissionpolicy/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065345254365597696#1:build-log.txt%3A522)

* _2026-06-12 06:43:59 &#43;0000 UTC_: <code>06:47:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/validatingadmissionpolicy/BUILD.bazel:3:11: GoCompilePkg tests/validatingadmissionpolicy/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065324059809288192#1:build-log.txt%3A660)

</details>

<hr/>

**1x**: _2026-06-12 10:56:48 &#43;0000 UTC_: <code>11:01:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065380361751236608#1:build-log.txt%3A619)
<details>
<summary>all...</summary>

* _2026-06-12 10:56:48 &#43;0000 UTC_: <code>11:01:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065380361751236608#1:build-log.txt%3A619)

</details>

<hr/>
</details>
<details>
<summary> HTTP error in context (3x / 21.43%) </summary>

<hr/>

**3x**: _2026-06-12 09:49:35 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.34-sig-storage/2065370810171789312#1:build-log.txt%3A607)
<details>
<summary>all...</summary>

* _2026-06-12 09:54:34 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.34-windows2016/2065371648579604480#1:build-log.txt%3A627)

* _2026-06-12 09:50:33 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.36-sig-storage/2065370831113949184#1:build-log.txt%3A611)

* _2026-06-12 09:49:35 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.34-sig-storage/2065370810171789312#1:build-log.txt%3A607)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 7.14%) </summary>

<hr/>

**1x**: _2026-06-12 06:47:52 &#43;0000 UTC_: <code>06:49:59: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18035/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065325019264716800#1:build-log.txt%3A652)
<details>
<summary>all...</summary>

* _2026-06-12 06:47:52 &#43;0000 UTC_: <code>06:49:59: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18035/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065325019264716800#1:build-log.txt%3A652)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (19x / 73.08%)

<details>
<summary> bazel remote cache blob fetch failure (10x / 38.46%) </summary>

<hr/>

**5x**: _2026-06-12 09:02:29 &#43;0000 UTC_: <code>09:05:24: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-storage/2065358924223090688#1:build-log.txt%3A398)
<details>
<summary>all...</summary>

* _2026-06-12 09:02:31 &#43;0000 UTC_: <code>09:07:05: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2065358924449583104#1:build-log.txt%3A500)
<details><summary>context</summary>
<pre>09:06:53: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:06:53: /usr/bin/ld: /tmp/go-link-3551560137/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
09:06:53: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:07:05: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:
09:07:05: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: d54ae0ec30d73d043b28853d40940897a0b32ac3bf43790da5a64317e5f0ffb4/351 for bazel-out/k8-fastbuild/bin/images/disks-images-provider/disks-images-provider-image.2.descriptor.json
09:07:05: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 13c628fc69020eccae6684948e365c2a7f67ea64d879ab7fa1f8c483daaaff15/352 for bazel-out/k8-fastbuild/bin/images/disks-images-provider/disks-images-provider-image.0.descriptor.json
09:07:05: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 700bf3967f973df5b7f14d4c1cbd0c7c5db8b72f9aea09c91d86736100e380b1/351 for bazel-out/k8-fastbuild/bin/images/disks-images-provider/disks-images-provider-image.1.descriptor.json</pre>
</details>


* _2026-06-12 09:02:30 &#43;0000 UTC_: <code>09:05:29: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-operator/2065358924541857792#1:build-log.txt%3A430)
<details><summary>context</summary>
<pre>09:05:17:    74 |   asprintf (&amp;err-&gt;error, &#34;%s: &#34;
09:05:17:       |   ^~~~~~~~
09:05:17:       |   vsprintf
09:05:29: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:
09:05:29: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: d54ae0ec30d73d043b28853d40940897a0b32ac3bf43790da5a64317e5f0ffb4/351 for bazel-out/k8-fastbuild/bin/images/disks-images-provider/disks-images-provider-image.2.descriptor.json
09:05:29: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 700bf3967f973df5b7f14d4c1cbd0c7c5db8b72f9aea09c91d86736100e380b1/351 for bazel-out/k8-fastbuild/bin/images/disks-images-provider/disks-images-provider-image.1.descriptor.json
09:05:29: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 13c628fc69020eccae6684948e365c2a7f67ea64d879ab7fa1f8c483daaaff15/352 for bazel-out/k8-fastbuild/bin/images/disks-images-provider/disks-images-provider-image.0.descriptor.json</pre>
</details>


* _2026-06-12 09:02:29 &#43;0000 UTC_: <code>09:05:24: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-storage/2065358924223090688#1:build-log.txt%3A398)
<details><summary>context</summary>
<pre>09:05:14: INFO: From GoLink cmd/virt-launcher/libvirt-hook-client/libvirt-hook-client_/libvirt-hook-client:
09:05:14: /usr/bin/ld: /tmp/go-link-435826452/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
09:05:14: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:05:24: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:
09:05:24: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 13c628fc69020eccae6684948e365c2a7f67ea64d879ab7fa1f8c483daaaff15/352 for bazel-out/k8-fastbuild/bin/images/disks-images-provider/disks-images-provider-image.0.descriptor.json
09:05:24: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: d54ae0ec30d73d043b28853d40940897a0b32ac3bf43790da5a64317e5f0ffb4/351 for bazel-out/k8-fastbuild/bin/images/disks-images-provider/disks-images-provider-image.2.descriptor.json
09:05:24: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 700bf3967f973df5b7f14d4c1cbd0c7c5db8b72f9aea09c91d86736100e380b1/351 for bazel-out/k8-fastbuild/bin/images/disks-images-provider/disks-images-provider-image.1.descriptor.json</pre>
</details>


* _2026-06-12 09:02:28 &#43;0000 UTC_: <code>09:05:43: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-network/2065358924130816000#1:build-log.txt%3A405)
<details><summary>context</summary>
<pre>09:05:32: INFO: From GoLink cmd/virt-launcher/libvirt-hook-client/libvirt-hook-client_/libvirt-hook-client:
09:05:32: /usr/bin/ld: /tmp/go-link-435826452/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
09:05:32: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:05:43: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:
09:05:43: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: d54ae0ec30d73d043b28853d40940897a0b32ac3bf43790da5a64317e5f0ffb4/351 for bazel-out/k8-fastbuild/bin/images/disks-images-provider/disks-images-provider-image.2.descriptor.json
09:05:43: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 700bf3967f973df5b7f14d4c1cbd0c7c5db8b72f9aea09c91d86736100e380b1/351 for bazel-out/k8-fastbuild/bin/images/disks-images-provider/disks-images-provider-image.1.descriptor.json
09:05:43: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 13c628fc69020eccae6684948e365c2a7f67ea64d879ab7fa1f8c483daaaff15/352 for bazel-out/k8-fastbuild/bin/images/disks-images-provider/disks-images-provider-image.0.descriptor.json</pre>
</details>


* _2026-06-12 09:02:25 &#43;0000 UTC_: <code>09:05:22: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-compute/2065358924294393856#1:build-log.txt%3A390)
<details><summary>context</summary>
<pre>09:05:12: INFO: Invocation ID: 248e38c1-743e-4e95-934d-e1a6d35c1369
09:05:12: INFO: Analyzed 30 targets (0 packages loaded, 0 targets configured).
09:05:12: INFO: Found 30 targets...
09:05:22: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:
09:05:22: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 700bf3967f973df5b7f14d4c1cbd0c7c5db8b72f9aea09c91d86736100e380b1/351 for bazel-out/k8-fastbuild/bin/images/disks-images-provider/disks-images-provider-image.1.descriptor.json
09:05:22: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 13c628fc69020eccae6684948e365c2a7f67ea64d879ab7fa1f8c483daaaff15/352 for bazel-out/k8-fastbuild/bin/images/disks-images-provider/disks-images-provider-image.0.descriptor.json
09:05:22: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: d54ae0ec30d73d043b28853d40940897a0b32ac3bf43790da5a64317e5f0ffb4/351 for bazel-out/k8-fastbuild/bin/images/disks-images-provider/disks-images-provider-image.2.descriptor.json</pre>
</details>


</details>

<hr/>

**2x**: _2026-06-12 20:31:19 &#43;0000 UTC_: <code>20:36:05: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/compute/BUILD.bazel:3:11: GoCompilePkg tests/compute/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065532284278149120#1:build-log.txt%3A598)
<details>
<summary>all...</summary>

* _2026-06-12 20:31:19 &#43;0000 UTC_: <code>20:36:05: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/compute/BUILD.bazel:3:11: GoCompilePkg tests/compute/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065532284278149120#1:build-log.txt%3A598)
<details><summary>context</summary>
<pre>20:35:51: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
20:35:51: /usr/bin/ld: /tmp/go-link-3251999922/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
20:35:51: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
20:36:05: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/compute/BUILD.bazel:3:11: GoCompilePkg tests/compute/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x
20:36:05: INFO: Elapsed time: 34.741s, Critical Path: 6.34s
20:36:05: INFO: 7460 processes: 7426 remote cache hit, 31 internal, 3 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-12 13:39:05 &#43;0000 UTC_: <code>13:43:10: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/compute/BUILD.bazel:3:11: GoCompilePkg tests/compute/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065428547740110848#1:build-log.txt%3A601)
<details><summary>context</summary>
<pre>13:42:58: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
13:42:58: /usr/bin/ld: /tmp/go-link-3251999922/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
13:42:58: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
13:43:10: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/compute/BUILD.bazel:3:11: GoCompilePkg tests/compute/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x
13:43:10: INFO: Elapsed time: 32.298s, Critical Path: 25.77s
13:43:10: INFO: 6454 processes: 6423 remote cache hit, 30 internal, 1 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>

**2x**: _2026-06-12 08:10:14 &#43;0000 UTC_: <code>08:14:17: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/validatingadmissionpolicy/BUILD.bazel:3:11: GoCompilePkg tests/validatingadmissionpolicy/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065345254365597696#1:build-log.txt%3A522)
<details>
<summary>all...</summary>

* _2026-06-12 08:10:14 &#43;0000 UTC_: <code>08:14:17: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/validatingadmissionpolicy/BUILD.bazel:3:11: GoCompilePkg tests/validatingadmissionpolicy/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065345254365597696#1:build-log.txt%3A522)
<details><summary>context</summary>
<pre>08:13:58:    74 |   asprintf (&amp;err-&gt;error, &#34;%s: &#34;
08:13:58:       |   ^~~~~~~~
08:13:58:       |   vsprintf
08:14:17: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/validatingadmissionpolicy/BUILD.bazel:3:11: GoCompilePkg tests/validatingadmissionpolicy/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x
08:14:17: INFO: Elapsed time: 20.880s, Critical Path: 5.18s
08:14:17: INFO: 4806 processes: 4776 remote cache hit, 30 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-12 06:43:59 &#43;0000 UTC_: <code>06:47:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/validatingadmissionpolicy/BUILD.bazel:3:11: GoCompilePkg tests/validatingadmissionpolicy/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065324059809288192#1:build-log.txt%3A660)
<details><summary>context</summary>
<pre>06:47:29: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
06:47:29: /usr/bin/ld: /tmp/go-link-231509647/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
06:47:29: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
06:47:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/validatingadmissionpolicy/BUILD.bazel:3:11: GoCompilePkg tests/validatingadmissionpolicy/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x
06:47:37: INFO: Elapsed time: 112.683s, Critical Path: 36.66s
06:47:37: INFO: 6486 processes: 6135 remote cache hit, 305 internal, 46 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>

**1x**: _2026-06-12 10:56:48 &#43;0000 UTC_: <code>11:01:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065380361751236608#1:build-log.txt%3A619)
<details>
<summary>all...</summary>

* _2026-06-12 10:56:48 &#43;0000 UTC_: <code>11:01:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065380361751236608#1:build-log.txt%3A619)
<details><summary>context</summary>
<pre>11:01:31: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
11:01:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/compute/BUILD.bazel:3:11: GoCompilePkg tests/compute/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x
11:01:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/migration/BUILD.bazel:3:11: GoCompilePkg tests/migration/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x
11:01:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x
11:01:47: INFO: Elapsed time: 45.895s, Critical Path: 20.67s
11:01:47: INFO: 7149 processes: 7114 remote cache hit, 30 internal, 5 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> HTTP error in context (3x / 11.54%) </summary>

<hr/>

**3x**: _2026-06-12 09:49:35 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.34-sig-storage/2065370810171789312#1:build-log.txt%3A607)
<details>
<summary>all...</summary>

* _2026-06-12 09:54:34 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.34-windows2016/2065371648579604480#1:build-log.txt%3A627)
<details><summary>context</summary>
<pre>10:01:59: main.main()
10:01:59: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
10:01:59: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:174: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-12 09:50:33 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.36-sig-storage/2065370831113949184#1:build-log.txt%3A611)
<details><summary>context</summary>
<pre>10:02:19: main.main()
10:02:19: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
10:02:19: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:174: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-12 09:49:35 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.34-sig-storage/2065370810171789312#1:build-log.txt%3A607)
<details><summary>context</summary>
<pre>10:00:45: main.main()
10:00:45: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
10:00:45: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:174: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> download failure in context (2x / 7.69%) </summary>

<hr/>

**1x**: _2026-06-18 08:42:15 &#43;0000 UTC_: <code>08:43:22: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17263/pull-kubevirt-e2e-k8s-1.34-sig-performance/2067528172320067584#1:build-log.txt%3A86)
<details>
<summary>all...</summary>

* _2026-06-18 08:42:15 &#43;0000 UTC_: <code>08:43:22: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17263/pull-kubevirt-e2e-k8s-1.34-sig-performance/2067528172320067584#1:build-log.txt%3A86)
<details><summary>context</summary>
<pre>08:43:22: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted:
08:43:22: INFO: Elapsed time: 25.777s
08:43:22: INFO: 0 processes.
08:43:22: ERROR: Build failed. Not running target
make: *** [Makefile:25: bazel-build] Error 1
&#43; EXIT_VALUE=2
&#43; set &#43;o xtrace</pre>
</details>


</details>

<hr/>

**1x**: _2026-06-12 06:47:52 &#43;0000 UTC_: <code>06:49:59: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18035/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065325019264716800#1:build-log.txt%3A652)
<details>
<summary>all...</summary>

* _2026-06-12 06:47:52 &#43;0000 UTC_: <code>06:49:59: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18035/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065325019264716800#1:build-log.txt%3A652)
<details><summary>context</summary>
<pre>06:49:59:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
06:49:59: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
06:49:59: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
06:49:59: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
06:49:59: INFO: Elapsed time: 1.013s
06:49:59: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> podman container removal timeout (1x / 3.85%) </summary>

<hr/>

**1x**: _2026-06-17 07:01:38 &#43;0000 UTC_: <code>Error: cannot remove container e672de5d181574bc76d99016452be15084c22ab8f3d5895e5e5f8d0a986dc6f7 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17976/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2067140165721133056#1:build-log.txt%3A1868)
<details>
<summary>all...</summary>

* _2026-06-17 07:01:38 &#43;0000 UTC_: <code>Error: cannot remove container e672de5d181574bc76d99016452be15084c22ab8f3d5895e5e5f8d0a986dc6f7 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17976/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2067140165721133056#1:build-log.txt%3A1868)
<details><summary>context</summary>
<pre>Error: cannot remove container 8d2ad360e3c6cbb7df3ce927bc2970362c7b6d0e538300b87b8f89a7e3f07d2e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 418d3eebf1f9cff01fa9e12d4bef1cfb14a8d7caaec3fa749d15525b64730a3a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 64299780286f8bb3d71dcf5c99da1840a93124aef44d0040e27f8bbd20f625db as it could not be stopped: given PID did not die within timeout
Error: cannot remove container e672de5d181574bc76d99016452be15084c22ab8f3d5895e5e5f8d0a986dc6f7 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 21 is not a child of this shell
================================================================================
Done cleaning up after podman in container.</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> download failure from external URL (1x / 3.85%) </summary>

<hr/>

**1x**: _2026-06-14 06:31:19 &#43;0000 UTC_: <code>06:32:23: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp3938845111312673746/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17913/pull-kubevirt-e2e-k8s-1.35-sig-storage/2066045650159538176#1:build-log.txt%3A311)
<details>
<summary>all...</summary>

* _2026-06-14 06:31:19 &#43;0000 UTC_: <code>06:32:23: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp3938845111312673746/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17913/pull-kubevirt-e2e-k8s-1.35-sig-storage/2066045650159538176#1:build-log.txt%3A311)
<details><summary>context</summary>
<pre>06:32:23: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 132, column 45, in _http_archive_impl
06:32:23: 		download_info = ctx.download_and_extract(
06:32:23: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp3938845111312673746/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error
06:32:23: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp3938845111312673746/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 3.85%) </summary>

<hr/>

**1x**: _2026-06-14 17:57:20 &#43;0000 UTC_: <code>18:13:46: I0614 14:13:46.824946    1620 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18122/pull-kubevirt-e2e-k8s-1.35-sig-operator-1.8/2066180130513883136#1:build-log.txt%3A768)
<details>
<summary>all...</summary>

* _2026-06-14 17:57:20 &#43;0000 UTC_: <code>18:13:46: I0614 14:13:46.824946    1620 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18122/pull-kubevirt-e2e-k8s-1.35-sig-operator-1.8/2066180130513883136#1:build-log.txt%3A768)
<details><summary>context</summary>
<pre>18:13:45: [control-plane-check] kube-scheduler is healthy after 4.055938187s
18:13:45: I0614 14:13:45.824344    1620 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;
18:13:46: I0614 14:13:46.324758    1620 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;
18:13:46: I0614 14:13:46.824946    1620 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
18:13:46: I0614 14:13:46.824985    1620 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
18:13:47: [control-plane-check] kube-apiserver is healthy after 6.001879309s
18:13:47: I0614 14:13:47.327592    1620 kubeconfig.go:657] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 3.85%) </summary>

<hr/>

**1x**: _2026-06-16 12:46:25 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2066864704415862784#1:build-log.txt%3A1383)
<details>
<summary>all...</summary>

* _2026-06-16 12:46:25 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2066864704415862784#1:build-log.txt%3A1383)
<details><summary>context</summary>
<pre>13:00:48: rsync: [sender] failed to connect to 127.0.0.1 (127.0.0.1): Connection refused (111)
13:00:48: rsync error: error in socket IO (code 10) at clientserver.c(139) [sender=3.4.1]
13:01:16: &#43; rm -f /tmp/kubevirt.deploy.7kv4
make: *** [Makefile:189: cluster-sync] Error 10
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>


</details>

<hr/>
</details>

### needs-investigation (4x / 15.38%)

<details>
<summary> no matching pattern (2x / 7.69%) </summary>

<hr/>

**2x**: _2026-06-13 11:08:07 &#43;0000 UTC_: <code>11:09:44: ERROR: Error initializing RemoteModule</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18034/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065748056631939072#1:build-log.txt%3A347)
<details>
<summary>all...</summary>

* _2026-06-13 11:38:34 &#43;0000 UTC_: <code>11:39:30: ERROR: Error initializing RemoteModule</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17689/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065760587530375168#1:build-log.txt%3A355)
<details><summary>context</summary>
<pre>11:39:30: Another command (pid=9) is running. Waiting for it to complete on the server (server_pid=60)...
11:39:30: INFO: Invocation ID: c5a53ee1-755c-4332-af36-84a31d4f0279
11:39:30: ERROR: Could not open auth credentials file &#39;/etc/bazel-cache-gcs/bazel-cache-sa.json&#39;: /etc/bazel-cache-gcs/bazel-cache-sa.json (No such file or directory)
11:39:30: ERROR: Error initializing RemoteModule
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>


* _2026-06-13 11:08:07 &#43;0000 UTC_: <code>11:09:44: ERROR: Error initializing RemoteModule</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18034/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065748056631939072#1:build-log.txt%3A347)
<details><summary>context</summary>
<pre>11:09:44: Another command (pid=9) is running. Waiting for it to complete on the server (server_pid=60)...
11:09:44: INFO: Invocation ID: 091c086c-d25e-486a-9180-787ecdf5f170
11:09:44: ERROR: Could not open auth credentials file &#39;/etc/bazel-cache-gcs/bazel-cache-sa.json&#39;: /etc/bazel-cache-gcs/bazel-cache-sa.json (No such file or directory)
11:09:44: ERROR: Error initializing RemoteModule
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> no error snippets (2x / 7.69%) </summary>

<hr/>

**1x**: _2026-06-14 11:27:08 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18125/pull-kubevirt-e2e-k8s-1.34-sig-network-1.8/2066116448811487232)
<details>
<summary>all...</summary>

* _2026-06-14 11:27:08 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18125/pull-kubevirt-e2e-k8s-1.34-sig-network-1.8/2066116448811487232)

</details>

<hr/>

**1x**: _2026-06-14 10:58:52 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18121/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network-1.8/2066112953937039360)
<details>
<summary>all...</summary>

* _2026-06-14 10:58:52 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18121/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network-1.8/2066112953937039360)

</details>

<hr/>
</details>

### internal (3x / 11.54%)

<details>
<summary> kind cluster creation failure (2x / 7.69%) </summary>

<hr/>

**2x**: _2026-06-16 11:48:08 &#43;0000 UTC_: <code>11:55:03: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-kind-1.34-sev/2066849993645363200#1:build-log.txt%3A729)
<details>
<summary>all...</summary>

* _2026-06-16 11:48:08 &#43;0000 UTC_: <code>11:55:03: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-kind-1.34-sev/2066849993645363200#1:build-log.txt%3A729)
<details><summary>context</summary>
<pre>11:55:01:  ✓ Ensuring node image (kindest/node:v1.34.3) 🖼
11:55:01:  • Preparing nodes 📦 📦   ...
11:55:03:  ✗ Preparing nodes 📦 📦
11:55:03: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;
11:55:03:
11:55:03: Stack Trace:
11:55:03: sigs.k8s.io/kind/pkg/errors.Errorf</pre>
</details>


* _2026-06-14 11:18:29 &#43;0000 UTC_: <code>11:24:59: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18125/pull-kubevirt-e2e-kind-1.34-sev-1.8/2066117945167187968#1:build-log.txt%3A789)
<details><summary>context</summary>
<pre>11:24:56:  ✓ Ensuring node image (kindest/node:v1.34.3) 🖼
11:24:56:  • Preparing nodes 📦 📦   ...
11:24:59:  ✗ Preparing nodes 📦 📦
11:24:59: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;
11:24:59:
11:24:59: Stack Trace:
11:24:59: sigs.k8s.io/kind/pkg/errors.Errorf</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> make cluster lifecycle target failure (1x / 3.85%) </summary>

<hr/>

**1x**: _2026-06-14 14:43:50 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18130/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2066169290440052736#1:build-log.txt%3A629)
<details>
<summary>all...</summary>

* _2026-06-14 14:43:50 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18130/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2066169290440052736#1:build-log.txt%3A629)
<details><summary>context</summary>
<pre>14:58:51: main.main()
14:58:51: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
14:58:51: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:173: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (21x / 80.77%)


#### external (18x / 85.71%)

<details>
<summary> bazel remote cache blob fetch failure (10x / 47.62%) </summary>

<hr/>

**5x**: _2026-06-12 09:02:29 &#43;0000 UTC_: <code>09:05:24: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-storage/2065358924223090688#1:build-log.txt%3A398)
<details>
<summary>all...</summary>

* _2026-06-12 09:02:31 &#43;0000 UTC_: <code>09:07:05: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2065358924449583104#1:build-log.txt%3A500)

* _2026-06-12 09:02:30 &#43;0000 UTC_: <code>09:05:29: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-operator/2065358924541857792#1:build-log.txt%3A430)

* _2026-06-12 09:02:29 &#43;0000 UTC_: <code>09:05:24: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-storage/2065358924223090688#1:build-log.txt%3A398)

* _2026-06-12 09:02:28 &#43;0000 UTC_: <code>09:05:43: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-network/2065358924130816000#1:build-log.txt%3A405)

* _2026-06-12 09:02:25 &#43;0000 UTC_: <code>09:05:22: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-compute/2065358924294393856#1:build-log.txt%3A390)

</details>

<hr/>

**2x**: _2026-06-12 20:31:19 &#43;0000 UTC_: <code>20:36:05: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/compute/BUILD.bazel:3:11: GoCompilePkg tests/compute/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065532284278149120#1:build-log.txt%3A598)
<details>
<summary>all...</summary>

* _2026-06-12 20:31:19 &#43;0000 UTC_: <code>20:36:05: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/compute/BUILD.bazel:3:11: GoCompilePkg tests/compute/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065532284278149120#1:build-log.txt%3A598)

* _2026-06-12 13:39:05 &#43;0000 UTC_: <code>13:43:10: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/compute/BUILD.bazel:3:11: GoCompilePkg tests/compute/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065428547740110848#1:build-log.txt%3A601)

</details>

<hr/>

**2x**: _2026-06-12 08:10:14 &#43;0000 UTC_: <code>08:14:17: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/validatingadmissionpolicy/BUILD.bazel:3:11: GoCompilePkg tests/validatingadmissionpolicy/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065345254365597696#1:build-log.txt%3A522)
<details>
<summary>all...</summary>

* _2026-06-12 08:10:14 &#43;0000 UTC_: <code>08:14:17: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/validatingadmissionpolicy/BUILD.bazel:3:11: GoCompilePkg tests/validatingadmissionpolicy/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065345254365597696#1:build-log.txt%3A522)

* _2026-06-12 06:43:59 &#43;0000 UTC_: <code>06:47:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/validatingadmissionpolicy/BUILD.bazel:3:11: GoCompilePkg tests/validatingadmissionpolicy/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065324059809288192#1:build-log.txt%3A660)

</details>

<hr/>

**1x**: _2026-06-12 10:56:48 &#43;0000 UTC_: <code>11:01:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065380361751236608#1:build-log.txt%3A619)
<details>
<summary>all...</summary>

* _2026-06-12 10:56:48 &#43;0000 UTC_: <code>11:01:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065380361751236608#1:build-log.txt%3A619)

</details>

<hr/>
</details>
<details>
<summary> HTTP error in context (3x / 14.29%) </summary>

<hr/>

**3x**: _2026-06-12 09:49:35 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.34-sig-storage/2065370810171789312#1:build-log.txt%3A607)
<details>
<summary>all...</summary>

* _2026-06-12 09:54:34 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.34-windows2016/2065371648579604480#1:build-log.txt%3A627)

* _2026-06-12 09:50:33 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.36-sig-storage/2065370831113949184#1:build-log.txt%3A611)

* _2026-06-12 09:49:35 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.34-sig-storage/2065370810171789312#1:build-log.txt%3A607)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (2x / 9.52%) </summary>

<hr/>

**1x**: _2026-06-18 08:42:15 &#43;0000 UTC_: <code>08:43:22: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17263/pull-kubevirt-e2e-k8s-1.34-sig-performance/2067528172320067584#1:build-log.txt%3A86)
<details>
<summary>all...</summary>

* _2026-06-18 08:42:15 &#43;0000 UTC_: <code>08:43:22: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17263/pull-kubevirt-e2e-k8s-1.34-sig-performance/2067528172320067584#1:build-log.txt%3A86)

</details>

<hr/>

**1x**: _2026-06-12 06:47:52 &#43;0000 UTC_: <code>06:49:59: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18035/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065325019264716800#1:build-log.txt%3A652)
<details>
<summary>all...</summary>

* _2026-06-12 06:47:52 &#43;0000 UTC_: <code>06:49:59: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18035/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065325019264716800#1:build-log.txt%3A652)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 4.76%) </summary>

<hr/>

**1x**: _2026-06-16 12:46:25 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2066864704415862784#1:build-log.txt%3A1383)
<details>
<summary>all...</summary>

* _2026-06-16 12:46:25 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2066864704415862784#1:build-log.txt%3A1383)

</details>

<hr/>
</details>
<details>
<summary> podman container removal timeout (1x / 4.76%) </summary>

<hr/>

**1x**: _2026-06-17 07:01:38 &#43;0000 UTC_: <code>Error: cannot remove container e672de5d181574bc76d99016452be15084c22ab8f3d5895e5e5f8d0a986dc6f7 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17976/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2067140165721133056#1:build-log.txt%3A1868)
<details>
<summary>all...</summary>

* _2026-06-17 07:01:38 &#43;0000 UTC_: <code>Error: cannot remove container e672de5d181574bc76d99016452be15084c22ab8f3d5895e5e5f8d0a986dc6f7 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17976/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2067140165721133056#1:build-log.txt%3A1868)

</details>

<hr/>
</details>
<details>
<summary> download failure from external URL (1x / 4.76%) </summary>

<hr/>

**1x**: _2026-06-14 06:31:19 &#43;0000 UTC_: <code>06:32:23: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp3938845111312673746/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17913/pull-kubevirt-e2e-k8s-1.35-sig-storage/2066045650159538176#1:build-log.txt%3A311)
<details>
<summary>all...</summary>

* _2026-06-14 06:31:19 &#43;0000 UTC_: <code>06:32:23: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp3938845111312673746/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17913/pull-kubevirt-e2e-k8s-1.35-sig-storage/2066045650159538176#1:build-log.txt%3A311)

</details>

<hr/>
</details>

#### needs-investigation (2x / 9.52%)

<details>
<summary> no matching pattern (2x / 9.52%) </summary>

<hr/>

**2x**: _2026-06-13 11:08:07 &#43;0000 UTC_: <code>11:09:44: ERROR: Error initializing RemoteModule</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18034/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065748056631939072#1:build-log.txt%3A347)
<details>
<summary>all...</summary>

* _2026-06-13 11:38:34 &#43;0000 UTC_: <code>11:39:30: ERROR: Error initializing RemoteModule</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17689/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065760587530375168#1:build-log.txt%3A355)

* _2026-06-13 11:08:07 &#43;0000 UTC_: <code>11:09:44: ERROR: Error initializing RemoteModule</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18034/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065748056631939072#1:build-log.txt%3A347)

</details>

<hr/>
</details>

#### internal (1x / 4.76%)

<details>
<summary> kind cluster creation failure (1x / 4.76%) </summary>

<hr/>

**1x**: _2026-06-16 11:48:08 &#43;0000 UTC_: <code>11:55:03: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-kind-1.34-sev/2066849993645363200#1:build-log.txt%3A729)
<details>
<summary>all...</summary>

* _2026-06-16 11:48:08 &#43;0000 UTC_: <code>11:55:03: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-kind-1.34-sev/2066849993645363200#1:build-log.txt%3A729)

</details>

<hr/>
</details>

### release-1.8 (5x / 19.23%)


#### internal (2x / 40.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-06-14 14:43:50 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18130/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2066169290440052736#1:build-log.txt%3A629)
<details>
<summary>all...</summary>

* _2026-06-14 14:43:50 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18130/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2066169290440052736#1:build-log.txt%3A629)

</details>

<hr/>
</details>
<details>
<summary> kind cluster creation failure (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-06-14 11:18:29 &#43;0000 UTC_: <code>11:24:59: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18125/pull-kubevirt-e2e-kind-1.34-sev-1.8/2066117945167187968#1:build-log.txt%3A789)
<details>
<summary>all...</summary>

* _2026-06-14 11:18:29 &#43;0000 UTC_: <code>11:24:59: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18125/pull-kubevirt-e2e-kind-1.34-sev-1.8/2066117945167187968#1:build-log.txt%3A789)

</details>

<hr/>
</details>

#### needs-investigation (2x / 40.00%)

<details>
<summary> no error snippets (2x / 40.00%) </summary>

<hr/>

**1x**: _2026-06-14 11:27:08 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18125/pull-kubevirt-e2e-k8s-1.34-sig-network-1.8/2066116448811487232)
<details>
<summary>all...</summary>

* _2026-06-14 11:27:08 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18125/pull-kubevirt-e2e-k8s-1.34-sig-network-1.8/2066116448811487232)

</details>

<hr/>

**1x**: _2026-06-14 10:58:52 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18121/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network-1.8/2066112953937039360)
<details>
<summary>all...</summary>

* _2026-06-14 10:58:52 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18121/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network-1.8/2066112953937039360)

</details>

<hr/>
</details>

#### external (1x / 20.00%)

<details>
<summary> transient kube-apiserver body decode noise (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-06-14 17:57:20 &#43;0000 UTC_: <code>18:13:46: I0614 14:13:46.824946    1620 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18122/pull-kubevirt-e2e-k8s-1.35-sig-operator-1.8/2066180130513883136#1:build-log.txt%3A768)
<details>
<summary>all...</summary>

* _2026-06-14 17:57:20 &#43;0000 UTC_: <code>18:13:46: I0614 14:13:46.824946    1620 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18122/pull-kubevirt-e2e-k8s-1.35-sig-operator-1.8/2066180130513883136#1:build-log.txt%3A768)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-performance (1x / 3.85%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-06-18 08:42:15 &#43;0000 UTC_: <code>08:43:22: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17263/pull-kubevirt-e2e-k8s-1.34-sig-performance/2067528172320067584#1:build-log.txt%3A86)
<details>
<summary>all...</summary>

* _2026-06-18 08:42:15 &#43;0000 UTC_: <code>08:43:22: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17263/pull-kubevirt-e2e-k8s-1.34-sig-performance/2067528172320067584#1:build-log.txt%3A86)

</details>

<hr/>
</details>

### sig-compute (18x / 69.23%)


#### external (13x / 72.22%)

<details>
<summary> bazel remote cache blob fetch failure (8x / 44.44%) </summary>

<hr/>

**3x**: _2026-06-12 09:02:31 &#43;0000 UTC_: <code>09:07:05: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2065358924449583104#1:build-log.txt%3A500)
<details>
<summary>all...</summary>

* _2026-06-12 09:02:31 &#43;0000 UTC_: <code>09:07:05: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2065358924449583104#1:build-log.txt%3A500)

* _2026-06-12 09:02:30 &#43;0000 UTC_: <code>09:05:29: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-operator/2065358924541857792#1:build-log.txt%3A430)

* _2026-06-12 09:02:25 &#43;0000 UTC_: <code>09:05:22: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-compute/2065358924294393856#1:build-log.txt%3A390)

</details>

<hr/>

**2x**: _2026-06-12 20:31:19 &#43;0000 UTC_: <code>20:36:05: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/compute/BUILD.bazel:3:11: GoCompilePkg tests/compute/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065532284278149120#1:build-log.txt%3A598)
<details>
<summary>all...</summary>

* _2026-06-12 20:31:19 &#43;0000 UTC_: <code>20:36:05: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/compute/BUILD.bazel:3:11: GoCompilePkg tests/compute/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065532284278149120#1:build-log.txt%3A598)

* _2026-06-12 13:39:05 &#43;0000 UTC_: <code>13:43:10: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/compute/BUILD.bazel:3:11: GoCompilePkg tests/compute/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065428547740110848#1:build-log.txt%3A601)

</details>

<hr/>

**2x**: _2026-06-12 08:10:14 &#43;0000 UTC_: <code>08:14:17: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/validatingadmissionpolicy/BUILD.bazel:3:11: GoCompilePkg tests/validatingadmissionpolicy/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065345254365597696#1:build-log.txt%3A522)
<details>
<summary>all...</summary>

* _2026-06-12 08:10:14 &#43;0000 UTC_: <code>08:14:17: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/validatingadmissionpolicy/BUILD.bazel:3:11: GoCompilePkg tests/validatingadmissionpolicy/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065345254365597696#1:build-log.txt%3A522)

* _2026-06-12 06:43:59 &#43;0000 UTC_: <code>06:47:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/validatingadmissionpolicy/BUILD.bazel:3:11: GoCompilePkg tests/validatingadmissionpolicy/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065324059809288192#1:build-log.txt%3A660)

</details>

<hr/>

**1x**: _2026-06-12 10:56:48 &#43;0000 UTC_: <code>11:01:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065380361751236608#1:build-log.txt%3A619)
<details>
<summary>all...</summary>

* _2026-06-12 10:56:48 &#43;0000 UTC_: <code>11:01:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 5a073ce792c5779ff76d563eb571c5e95cb5ba1f313a74b303a11e48ea2fb594/25446 for bazel-out/k8-fastbuild/bin/tests/decorators/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18072/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065380361751236608#1:build-log.txt%3A619)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 5.56%) </summary>

<hr/>

**1x**: _2026-06-16 12:46:25 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2066864704415862784#1:build-log.txt%3A1383)
<details>
<summary>all...</summary>

* _2026-06-16 12:46:25 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2066864704415862784#1:build-log.txt%3A1383)

</details>

<hr/>
</details>
<details>
<summary> podman container removal timeout (1x / 5.56%) </summary>

<hr/>

**1x**: _2026-06-17 07:01:38 &#43;0000 UTC_: <code>Error: cannot remove container e672de5d181574bc76d99016452be15084c22ab8f3d5895e5e5f8d0a986dc6f7 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17976/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2067140165721133056#1:build-log.txt%3A1868)
<details>
<summary>all...</summary>

* _2026-06-17 07:01:38 &#43;0000 UTC_: <code>Error: cannot remove container e672de5d181574bc76d99016452be15084c22ab8f3d5895e5e5f8d0a986dc6f7 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17976/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2067140165721133056#1:build-log.txt%3A1868)

</details>

<hr/>
</details>
<details>
<summary> HTTP error in context (1x / 5.56%) </summary>

<hr/>

**1x**: _2026-06-12 09:54:34 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.34-windows2016/2065371648579604480#1:build-log.txt%3A627)
<details>
<summary>all...</summary>

* _2026-06-12 09:54:34 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.34-windows2016/2065371648579604480#1:build-log.txt%3A627)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 5.56%) </summary>

<hr/>

**1x**: _2026-06-14 17:57:20 &#43;0000 UTC_: <code>18:13:46: I0614 14:13:46.824946    1620 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18122/pull-kubevirt-e2e-k8s-1.35-sig-operator-1.8/2066180130513883136#1:build-log.txt%3A768)
<details>
<summary>all...</summary>

* _2026-06-14 17:57:20 &#43;0000 UTC_: <code>18:13:46: I0614 14:13:46.824946    1620 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18122/pull-kubevirt-e2e-k8s-1.35-sig-operator-1.8/2066180130513883136#1:build-log.txt%3A768)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 5.56%) </summary>

<hr/>

**1x**: _2026-06-12 06:47:52 &#43;0000 UTC_: <code>06:49:59: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18035/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065325019264716800#1:build-log.txt%3A652)
<details>
<summary>all...</summary>

* _2026-06-12 06:47:52 &#43;0000 UTC_: <code>06:49:59: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18035/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065325019264716800#1:build-log.txt%3A652)

</details>

<hr/>
</details>

#### internal (3x / 16.67%)

<details>
<summary> kind cluster creation failure (2x / 11.11%) </summary>

<hr/>

**2x**: _2026-06-16 11:48:08 &#43;0000 UTC_: <code>11:55:03: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-kind-1.34-sev/2066849993645363200#1:build-log.txt%3A729)
<details>
<summary>all...</summary>

* _2026-06-16 11:48:08 &#43;0000 UTC_: <code>11:55:03: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-kind-1.34-sev/2066849993645363200#1:build-log.txt%3A729)

* _2026-06-14 11:18:29 &#43;0000 UTC_: <code>11:24:59: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18125/pull-kubevirt-e2e-kind-1.34-sev-1.8/2066117945167187968#1:build-log.txt%3A789)

</details>

<hr/>
</details>
<details>
<summary> make cluster lifecycle target failure (1x / 5.56%) </summary>

<hr/>

**1x**: _2026-06-14 14:43:50 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18130/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2066169290440052736#1:build-log.txt%3A629)
<details>
<summary>all...</summary>

* _2026-06-14 14:43:50 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18130/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2066169290440052736#1:build-log.txt%3A629)

</details>

<hr/>
</details>

#### needs-investigation (2x / 11.11%)

<details>
<summary> no matching pattern (2x / 11.11%) </summary>

<hr/>

**2x**: _2026-06-13 11:08:07 &#43;0000 UTC_: <code>11:09:44: ERROR: Error initializing RemoteModule</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18034/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065748056631939072#1:build-log.txt%3A347)
<details>
<summary>all...</summary>

* _2026-06-13 11:38:34 &#43;0000 UTC_: <code>11:39:30: ERROR: Error initializing RemoteModule</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17689/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065760587530375168#1:build-log.txt%3A355)

* _2026-06-13 11:08:07 &#43;0000 UTC_: <code>11:09:44: ERROR: Error initializing RemoteModule</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18034/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065748056631939072#1:build-log.txt%3A347)

</details>

<hr/>
</details>

### sig-network (3x / 11.54%)


#### needs-investigation (2x / 66.67%)

<details>
<summary> no error snippets (2x / 66.67%) </summary>

<hr/>

**1x**: _2026-06-14 11:27:08 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18125/pull-kubevirt-e2e-k8s-1.34-sig-network-1.8/2066116448811487232)
<details>
<summary>all...</summary>

* _2026-06-14 11:27:08 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18125/pull-kubevirt-e2e-k8s-1.34-sig-network-1.8/2066116448811487232)

</details>

<hr/>

**1x**: _2026-06-14 10:58:52 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18121/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network-1.8/2066112953937039360)
<details>
<summary>all...</summary>

* _2026-06-14 10:58:52 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18121/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network-1.8/2066112953937039360)

</details>

<hr/>
</details>

#### external (1x / 33.33%)

<details>
<summary> bazel remote cache blob fetch failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-06-12 09:02:28 &#43;0000 UTC_: <code>09:05:43: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-network/2065358924130816000#1:build-log.txt%3A405)
<details>
<summary>all...</summary>

* _2026-06-12 09:02:28 &#43;0000 UTC_: <code>09:05:43: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-network/2065358924130816000#1:build-log.txt%3A405)

</details>

<hr/>
</details>

### sig-storage (4x / 15.38%)


#### external (4x / 100.00%)

<details>
<summary> HTTP error in context (2x / 50.00%) </summary>

<hr/>

**2x**: _2026-06-12 09:49:35 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.34-sig-storage/2065370810171789312#1:build-log.txt%3A607)
<details>
<summary>all...</summary>

* _2026-06-12 09:50:33 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.36-sig-storage/2065370831113949184#1:build-log.txt%3A611)

* _2026-06-12 09:49:35 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.34-sig-storage/2065370810171789312#1:build-log.txt%3A607)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache blob fetch failure (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-06-12 09:02:29 &#43;0000 UTC_: <code>09:05:24: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-storage/2065358924223090688#1:build-log.txt%3A398)
<details>
<summary>all...</summary>

* _2026-06-12 09:02:29 &#43;0000 UTC_: <code>09:05:24: ERROR: /root/go/src/kubevirt.io/kubevirt/images/disks-images-provider/BUILD.bazel:50:10: OCI Image //images/disks-images-provider:disks-images-provider-image failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17422/pull-kubevirt-e2e-k8s-1.36-sig-storage/2065358924223090688#1:build-log.txt%3A398)

</details>

<hr/>
</details>
<details>
<summary> download failure from external URL (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-06-14 06:31:19 &#43;0000 UTC_: <code>06:32:23: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp3938845111312673746/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17913/pull-kubevirt-e2e-k8s-1.35-sig-storage/2066045650159538176#1:build-log.txt%3A311)
<details>
<summary>all...</summary>

* _2026-06-14 06:31:19 &#43;0000 UTC_: <code>06:32:23: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp3938845111312673746/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17913/pull-kubevirt-e2e-k8s-1.35-sig-storage/2066045650159538176#1:build-log.txt%3A311)

</details>

<hr/>
</details>

Last updated: 2026-06-18 19:04:28
