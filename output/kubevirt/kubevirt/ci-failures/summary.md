



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-06-14 (1x / 2.44%)


#### external (1x / 100.00%)

<details>
<summary> download failure from external URL (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-06-14 06:31:19 &#43;0000 UTC_: <code>06:32:23: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp3938845111312673746/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17913/pull-kubevirt-e2e-k8s-1.35-sig-storage/2066045650159538176#1:build-log.txt%3A311)
<details>
<summary>all...</summary>

* _2026-06-14 06:31:19 &#43;0000 UTC_: <code>06:32:23: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp3938845111312673746/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17913/pull-kubevirt-e2e-k8s-1.35-sig-storage/2066045650159538176#1:build-log.txt%3A311)

</details>

<hr/>
</details>

### 2026-06-13 (1x / 2.44%)


#### needs-investigation (1x / 100.00%)

<details>
<summary> no matching pattern (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-06-13 11:08:07 &#43;0000 UTC_: <code>11:09:44: ERROR: Error initializing RemoteModule</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18034/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065748056631939072#1:build-log.txt%3A347)
<details>
<summary>all...</summary>

* _2026-06-13 11:08:07 &#43;0000 UTC_: <code>11:09:44: ERROR: Error initializing RemoteModule</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18034/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065748056631939072#1:build-log.txt%3A347)

</details>

<hr/>
</details>

### 2026-06-12 (9x / 21.95%)


#### external (9x / 100.00%)

<details>
<summary> bazel remote cache blob fetch failure (5x / 55.56%) </summary>

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
<summary> HTTP error in context (3x / 33.33%) </summary>

<hr/>

**3x**: _2026-06-12 09:54:34 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.34-windows2016/2065371648579604480#1:build-log.txt%3A627)
<details>
<summary>all...</summary>

* _2026-06-12 09:54:34 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.34-windows2016/2065371648579604480#1:build-log.txt%3A627)

* _2026-06-12 09:50:33 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.36-sig-storage/2065370831113949184#1:build-log.txt%3A611)

* _2026-06-12 09:49:35 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.34-sig-storage/2065370810171789312#1:build-log.txt%3A607)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-06-12 06:47:52 &#43;0000 UTC_: <code>06:49:59: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18035/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065325019264716800#1:build-log.txt%3A652)
<details>
<summary>all...</summary>

* _2026-06-12 06:47:52 &#43;0000 UTC_: <code>06:49:59: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18035/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065325019264716800#1:build-log.txt%3A652)

</details>

<hr/>
</details>

### 2026-06-11 (25x / 60.98%)


#### external (25x / 100.00%)

<details>
<summary> HTTP error in context (23x / 92.00%) </summary>

<hr/>

**23x**: _2026-06-11 07:23:11 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18056/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2064971560589987840#1:build-log.txt%3A623)
<details>
<summary>all...</summary>

* _2026-06-11 08:43:11 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.34-sig-storage/2064991677570879488#1:build-log.txt%3A603)

* _2026-06-11 08:43:11 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.34-windows2016/2064991677390524416#1:build-log.txt%3A620)

* _2026-06-11 08:43:11 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.35-sig-storage/2064991677696708608#1:build-log.txt%3A604)

* _2026-06-11 08:11:20 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16106/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2064983681084166144#1:build-log.txt%3A796)

* _2026-06-11 07:58:25 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18055/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2064980396310794240#1:build-log.txt%3A623)

* _2026-06-11 07:44:55 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18094/pull-kubevirt-e2e-k8s-1.35-sig-storage/2064976950866219008#1:build-log.txt%3A598)

* _2026-06-11 07:44:48 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18094/pull-kubevirt-e2e-k8s-1.34-windows2016/2064976949725368320#1:build-log.txt%3A614)

* _2026-06-11 07:44:47 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18094/pull-kubevirt-e2e-k8s-1.34-sig-storage/2064976950388068352#1:build-log.txt%3A600)

* _2026-06-11 07:24:17 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18055/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations-1.8/2064971753305673728#1:build-log.txt%3A686)

* _2026-06-11 07:24:13 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18055/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2064971753439891456#1:build-log.txt%3A609)

* _2026-06-11 07:24:01 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18055/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2064971753741881344#1:build-log.txt%3A608)

* _2026-06-11 07:24:00 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18055/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2064971753607663616#1:build-log.txt%3A606)

* _2026-06-11 07:23:14 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18056/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2064971560803897344#1:build-log.txt%3A606)

* _2026-06-11 07:23:13 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18056/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2064971560724205568#1:build-log.txt%3A611)

* _2026-06-11 07:23:11 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18056/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2064971560589987840#1:build-log.txt%3A623)

* _2026-06-11 07:23:11 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18056/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2064971560891977728#1:build-log.txt%3A606)

* _2026-06-11 06:58:14 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17651/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2064965270748794880#1:build-log.txt%3A695)

* _2026-06-11 06:58:14 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17651/pull-kubevirt-e2e-k8s-1.35-sig-storage/2064965270971092992#1:build-log.txt%3A616)

* _2026-06-11 06:58:13 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17651/pull-kubevirt-e2e-k8s-1.36-sig-storage/2064965271050784768#1:build-log.txt%3A616)

* _2026-06-11 06:58:13 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17651/pull-kubevirt-e2e-k8s-1.34-windows2016/2064965270564245504#1:build-log.txt%3A632)

* _2026-06-11 06:58:12 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17651/pull-kubevirt-e2e-k8s-1.34-sig-storage/2064965270895595520#1:build-log.txt%3A615)

* _2026-06-11 05:48:18 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.34-windows2016/2064947674259394560#1:build-log.txt%3A698)

* _2026-06-11 05:48:17 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.34-sig-storage/2064947675047923712#1:build-log.txt%3A642)

</details>

<hr/>
</details>
<details>
<summary> failed external fetch in context (2x / 8.00%) </summary>

<hr/>

**2x**: _2026-06-11 05:48:18 &#43;0000 UTC_: <code>05:52:48: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.35-sig-operator/2064947676004225024#1:build-log.txt%3A411)
<details>
<summary>all...</summary>

* _2026-06-11 05:48:18 &#43;0000 UTC_: <code>05:52:48: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.35-sig-operator/2064947676004225024#1:build-log.txt%3A411)

* _2026-06-11 05:48:17 &#43;0000 UTC_: <code>05:52:44: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.35-sig-storage/2064947675668680704#1:build-log.txt%3A417)

</details>

<hr/>
</details>

### 2026-06-09 (3x / 7.32%)


#### internal (2x / 66.67%)

<details>
<summary> make cluster lifecycle target failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-06-09 12:55:48 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2064329652029100032#1:build-log.txt%3A1196)
<details>
<summary>all...</summary>

* _2026-06-09 12:55:48 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2064329652029100032#1:build-log.txt%3A1196)

</details>

<hr/>
</details>
<details>
<summary> make bazel-build target failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-06-09 08:35:56 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2064264517830840320#1:build-log.txt%3A4515)
<details>
<summary>all...</summary>

* _2026-06-09 08:35:56 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2064264517830840320#1:build-log.txt%3A4515)

</details>

<hr/>
</details>

#### external (1x / 33.33%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-06-09 08:33:34 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2064264500244123648#1:build-log.txt%3A4384)
<details>
<summary>all...</summary>

* _2026-06-09 08:33:34 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2064264500244123648#1:build-log.txt%3A4384)

</details>

<hr/>
</details>

### 2026-06-08 (2x / 4.88%)


#### external (1x / 50.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-06-08 13:09:50 &#43;0000 UTC_: <code>15:08:10:   {&#34;[namespace kubevirt-test-default1 name testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid d112888e-d01a-47a1-af22-ff1f17598cf2]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-06-08T15:07:46.018950Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-08T15:07:46.055406Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2063971602751557632#1:build-log.txt%3A6872)
<details>
<summary>all...</summary>

* _2026-06-08 13:09:50 &#43;0000 UTC_: <code>15:08:10:   {&#34;[namespace kubevirt-test-default1 name testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid d112888e-d01a-47a1-af22-ff1f17598cf2]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-06-08T15:07:46.018950Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-08T15:07:46.055406Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2063971602751557632#1:build-log.txt%3A6872)

</details>

<hr/>
</details>

#### needs-investigation (1x / 50.00%)

<details>
<summary> no matching pattern (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-06-08 13:09:41 &#43;0000 UTC_: <code>13:29:32: failed to create new CRI image service: validate service connection: validate CRI v1 image API for endpoint &#34;unix:///var/run/crio/crio.sock&#34;: rpc error: code = DeadlineExceeded desc = stream terminated by RST_STREAM with error code: CANCEL[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2063971600276918272#1:build-log.txt%3A674)
<details>
<summary>all...</summary>

* _2026-06-08 13:09:41 &#43;0000 UTC_: <code>13:29:32: failed to create new CRI image service: validate service connection: validate CRI v1 image API for endpoint &#34;unix:///var/run/crio/crio.sock&#34;: rpc error: code = DeadlineExceeded desc = stream terminated by RST_STREAM with error code: CANCEL[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2063971600276918272#1:build-log.txt%3A674)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (37x / 90.24%)

<details>
<summary> HTTP error in context (26x / 63.41%) </summary>

<hr/>

**26x**: _2026-06-11 07:23:11 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18056/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2064971560589987840#1:build-log.txt%3A623)
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


* _2026-06-11 08:43:11 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.34-windows2016/2064991677390524416#1:build-log.txt%3A620)
<details><summary>context</summary>
<pre>08:47:06: main.main()
08:47:06: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
08:47:06: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:174: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-11 08:43:11 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.34-sig-storage/2064991677570879488#1:build-log.txt%3A603)
<details><summary>context</summary>
<pre>08:47:14: main.main()
08:47:14: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
08:47:14: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:174: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-11 08:43:11 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.35-sig-storage/2064991677696708608#1:build-log.txt%3A604)
<details><summary>context</summary>
<pre>08:47:07: main.main()
08:47:07: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
08:47:07: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:174: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-11 08:11:20 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16106/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2064983681084166144#1:build-log.txt%3A796)
<details><summary>context</summary>
<pre>08:19:03: main.main()
08:19:03: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
08:19:03: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:174: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-11 07:58:25 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18055/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2064980396310794240#1:build-log.txt%3A623)
<details><summary>context</summary>
<pre>08:20:06: main.main()
08:20:06: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
08:20:06: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:173: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-11 07:44:55 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18094/pull-kubevirt-e2e-k8s-1.35-sig-storage/2064976950866219008#1:build-log.txt%3A598)
<details><summary>context</summary>
<pre>07:54:09: main.main()
07:54:09: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
07:54:09: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:174: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-11 07:44:48 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18094/pull-kubevirt-e2e-k8s-1.34-windows2016/2064976949725368320#1:build-log.txt%3A614)
<details><summary>context</summary>
<pre>07:54:08: main.main()
07:54:08: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
07:54:09: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:174: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-11 07:44:47 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18094/pull-kubevirt-e2e-k8s-1.34-sig-storage/2064976950388068352#1:build-log.txt%3A600)
<details><summary>context</summary>
<pre>07:51:28: main.main()
07:51:28: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
07:51:31: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:174: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-11 07:24:17 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18055/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations-1.8/2064971753305673728#1:build-log.txt%3A686)
<details><summary>context</summary>
<pre>07:57:09: main.main()
07:57:09: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
07:57:14: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:173: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-11 07:24:13 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18055/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2064971753439891456#1:build-log.txt%3A609)
<details><summary>context</summary>
<pre>07:56:45: main.main()
07:56:45: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
07:56:49: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:173: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-11 07:24:01 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18055/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2064971753741881344#1:build-log.txt%3A608)
<details><summary>context</summary>
<pre>07:52:49: main.main()
07:52:49: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
07:52:53: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:173: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-11 07:24:00 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18055/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2064971753607663616#1:build-log.txt%3A606)
<details><summary>context</summary>
<pre>07:50:39: main.main()
07:50:39: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
07:50:39: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:173: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-11 07:23:14 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18056/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2064971560803897344#1:build-log.txt%3A606)
<details><summary>context</summary>
<pre>07:52:48: main.main()
07:52:48: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
07:52:52: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:173: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-11 07:23:13 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18056/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2064971560724205568#1:build-log.txt%3A611)
<details><summary>context</summary>
<pre>07:53:23: main.main()
07:53:23: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
07:53:27: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:173: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-11 07:23:11 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18056/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2064971560891977728#1:build-log.txt%3A606)
<details><summary>context</summary>
<pre>07:55:12: main.main()
07:55:12: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
07:55:23: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:173: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-11 07:23:11 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18056/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2064971560589987840#1:build-log.txt%3A623)
<details><summary>context</summary>
<pre>07:56:35: main.main()
07:56:35: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
07:57:29: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:173: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-11 06:58:14 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17651/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2064965270748794880#1:build-log.txt%3A695)
<details><summary>context</summary>
<pre>07:02:20: main.main()
07:02:20: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
07:02:20: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:174: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-11 06:58:14 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17651/pull-kubevirt-e2e-k8s-1.35-sig-storage/2064965270971092992#1:build-log.txt%3A616)
<details><summary>context</summary>
<pre>07:02:14: main.main()
07:02:14: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
07:02:14: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:174: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-11 06:58:13 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17651/pull-kubevirt-e2e-k8s-1.34-windows2016/2064965270564245504#1:build-log.txt%3A632)
<details><summary>context</summary>
<pre>07:02:29: main.main()
07:02:29: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
07:02:29: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:174: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-11 06:58:13 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17651/pull-kubevirt-e2e-k8s-1.36-sig-storage/2064965271050784768#1:build-log.txt%3A616)
<details><summary>context</summary>
<pre>07:02:04: main.main()
07:02:04: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
07:02:06: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:174: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-11 06:58:12 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17651/pull-kubevirt-e2e-k8s-1.34-sig-storage/2064965270895595520#1:build-log.txt%3A615)
<details><summary>context</summary>
<pre>07:02:15: main.main()
07:02:15: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
07:02:15: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:174: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-11 05:48:18 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.34-windows2016/2064947674259394560#1:build-log.txt%3A698)
<details><summary>context</summary>
<pre>05:54:33: main.main()
05:54:33: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
05:54:33: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:174: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


* _2026-06-11 05:48:17 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.34-sig-storage/2064947675047923712#1:build-log.txt%3A642)
<details><summary>context</summary>
<pre>05:52:40: main.main()
05:52:40: 	/home/prow/go/src/github.com/kubevirt/kubevirtci/cluster-provision/gocli/cmd/cli/main.go:8 &#43;0xf
05:52:41: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:174: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> bazel remote cache blob fetch failure (5x / 12.20%) </summary>

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
<summary> failed external fetch in context (2x / 4.88%) </summary>

<hr/>

**2x**: _2026-06-11 05:48:18 &#43;0000 UTC_: <code>05:52:48: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.35-sig-operator/2064947676004225024#1:build-log.txt%3A411)
<details>
<summary>all...</summary>

* _2026-06-11 05:48:18 &#43;0000 UTC_: <code>05:52:48: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.35-sig-operator/2064947676004225024#1:build-log.txt%3A411)
<details><summary>context</summary>
<pre>05:52:48: error: RPC failed; HTTP 503 curl 22 The requested URL returned error: 503
05:52:48: fatal: expected &#39;packfile&#39;
05:52:48: fetch_repo: exit status 128
05:52:48: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:
05:52:48: INFO: Elapsed time: 30.922s
05:52:48: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-11 05:48:17 &#43;0000 UTC_: <code>05:52:44: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.35-sig-storage/2064947675668680704#1:build-log.txt%3A417)
<details><summary>context</summary>
<pre>05:52:44: error: RPC failed; HTTP 503 curl 22 The requested URL returned error: 503
05:52:44: fatal: expected &#39;packfile&#39;
05:52:44: fetch_repo: exit status 128
05:52:44: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:
05:52:44: INFO: Elapsed time: 31.002s
05:52:44: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 4.88%) </summary>

<hr/>

**1x**: _2026-06-09 08:33:34 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2064264500244123648#1:build-log.txt%3A4384)
<details>
<summary>all...</summary>

* _2026-06-09 08:33:34 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2064264500244123648#1:build-log.txt%3A4384)
<details><summary>context</summary>
<pre>09:08:13: INFO: Elapsed time: 0.398s, Critical Path: 0.13s
09:08:13: INFO: 1 process: 1 internal.
09:08:13: INFO: Running command line: bazel-bin/example-guest-agent-copier /root/go/src/kubevirt.io/kubevirt/_out/cmd/example-guest-agent/example-guest-agent
make: *** [Makefile:26: bazel-build-functests] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>


</details>

<hr/>

**1x**: _2026-06-08 13:09:50 &#43;0000 UTC_: <code>15:08:10:   {&#34;[namespace kubevirt-test-default1 name testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid d112888e-d01a-47a1-af22-ff1f17598cf2]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-06-08T15:07:46.018950Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-08T15:07:46.055406Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2063971602751557632#1:build-log.txt%3A6872)
<details>
<summary>all...</summary>

* _2026-06-08 13:09:50 &#43;0000 UTC_: <code>15:08:10:   {&#34;[namespace kubevirt-test-default1 name testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid d112888e-d01a-47a1-af22-ff1f17598cf2]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-06-08T15:07:46.018950Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-08T15:07:46.055406Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2063971602751557632#1:build-log.txt%3A6872)
<details><summary>context</summary>
<pre>15:08:10:
15:08:10:   Captured StdOut/StdErr Output &gt;&gt;
15:08:10:   {&#34;[namespace kubevirt-test-default1 name testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid d112888e-d01a-47a1-af22-ff1f17598cf2]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;Event(v1.ObjectReference{Kind:\&#34;VirtualMachineInstance\&#34;, Namespace:\&#34;kubevirt-test-default1\&#34;, Name:\&#34;testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\&#34;, UID:\&#34;d112888e-d01a-47a1-af22-ff1f17598cf2\&#34;, APIVersion:\&#34;kubevirt.io/v1\&#34;, ResourceVersion:\&#34;145126\&#34;, FieldPath:\&#34;\&#34;}): type: &#39;Normal&#39; reason: &#39;SuccessfulCreate&#39; Created virtual machine pod virt-launcher-testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxfrqsg&#34;,&#34;pos&#34;:&#34;watcher.go:150&#34;,&#34;timestamp&#34;:&#34;2026-06-08T15:07:42.269099Z&#34;}
15:08:10:   {&#34;[namespace kubevirt-test-default1 name testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid d112888e-d01a-47a1-af22-ff1f17598cf2]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-06-08T15:07:46.018950Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-08T15:07:46.055406Z&#34;}
15:08:10:   &lt;&lt; Captured StdOut/StdErr Output
15:08:10: ------------------------------
15:08:48: • [40.115 seconds]</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 2.44%) </summary>

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
<summary> download failure from external URL (1x / 2.44%) </summary>

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

### internal (2x / 4.88%)

<details>
<summary> make cluster lifecycle target failure (1x / 2.44%) </summary>

<hr/>

**1x**: _2026-06-09 12:55:48 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2064329652029100032#1:build-log.txt%3A1196)
<details>
<summary>all...</summary>

* _2026-06-09 12:55:48 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2064329652029100032#1:build-log.txt%3A1196)
<details><summary>context</summary>
<pre>13:15:45: INFO: 2 processes: 1 internal, 1 processwrapper-sandbox.
13:15:45: INFO: Running command line: bazel-bin/example-guest-agent-copier /root/go/src/kubevirt.io/kubevirt/_out/cmd/example-guest-agent/example-guest-agent
13:15:54: &#43; rm -f /tmp/kubevirt.deploy.r5Pz
make: *** [Makefile:174: cluster-sync] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> make bazel-build target failure (1x / 2.44%) </summary>

<hr/>

**1x**: _2026-06-09 08:35:56 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2064264517830840320#1:build-log.txt%3A4515)
<details>
<summary>all...</summary>

* _2026-06-09 08:35:56 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2064264517830840320#1:build-log.txt%3A4515)
<details><summary>context</summary>
<pre>09:11:37: INFO: Elapsed time: 0.318s, Critical Path: 0.10s
09:11:37: INFO: 1 process: 1 internal.
09:11:37: INFO: Running command line: bazel-bin/example-guest-agent-copier /root/go/src/kubevirt.io/kubevirt/_out/cmd/example-guest-agent/example-guest-agent
make: *** [Makefile:26: bazel-build-functests] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>


</details>

<hr/>
</details>

### needs-investigation (2x / 4.88%)

<details>
<summary> no matching pattern (2x / 4.88%) </summary>

<hr/>

**1x**: _2026-06-13 11:08:07 &#43;0000 UTC_: <code>11:09:44: ERROR: Error initializing RemoteModule</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18034/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065748056631939072#1:build-log.txt%3A347)
<details>
<summary>all...</summary>

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

**1x**: _2026-06-08 13:09:41 &#43;0000 UTC_: <code>13:29:32: failed to create new CRI image service: validate service connection: validate CRI v1 image API for endpoint &#34;unix:///var/run/crio/crio.sock&#34;: rpc error: code = DeadlineExceeded desc = stream terminated by RST_STREAM with error code: CANCEL[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2063971600276918272#1:build-log.txt%3A674)
<details>
<summary>all...</summary>

* _2026-06-08 13:09:41 &#43;0000 UTC_: <code>13:29:32: failed to create new CRI image service: validate service connection: validate CRI v1 image API for endpoint &#34;unix:///var/run/crio/crio.sock&#34;: rpc error: code = DeadlineExceeded desc = stream terminated by RST_STREAM with error code: CANCEL[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2063971600276918272#1:build-log.txt%3A674)
<details><summary>context</summary>
<pre>13:29:32: I0608 09:29:28.855786    1636 checks.go:201] validating availability of port 2380
13:29:32: I0608 09:29:28.855877    1636 checks.go:241] validating the existence and emptiness of directory /var/lib/etcd
13:29:32: [preflight] Some fatal errors occurred:
13:29:32: failed to create new CRI image service: validate service connection: validate CRI v1 image API for endpoint &#34;unix:///var/run/crio/crio.sock&#34;: rpc error: code = DeadlineExceeded desc = stream terminated by RST_STREAM with error code: CANCEL[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`
13:29:32: error execution phase preflight
13:29:32: k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow.(*Runner).Run.func1
13:29:32: 	k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow/runner.go:262</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (27x / 65.85%)


#### external (26x / 96.30%)

<details>
<summary> HTTP error in context (17x / 62.96%) </summary>

<hr/>

**17x**: _2026-06-11 07:44:48 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18094/pull-kubevirt-e2e-k8s-1.34-windows2016/2064976949725368320#1:build-log.txt%3A614)
<details>
<summary>all...</summary>

* _2026-06-12 09:54:34 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.34-windows2016/2065371648579604480#1:build-log.txt%3A627)

* _2026-06-12 09:50:33 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.36-sig-storage/2065370831113949184#1:build-log.txt%3A611)

* _2026-06-12 09:49:35 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.34-sig-storage/2065370810171789312#1:build-log.txt%3A607)

* _2026-06-11 08:43:11 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.34-windows2016/2064991677390524416#1:build-log.txt%3A620)

* _2026-06-11 08:43:11 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.34-sig-storage/2064991677570879488#1:build-log.txt%3A603)

* _2026-06-11 08:43:11 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.35-sig-storage/2064991677696708608#1:build-log.txt%3A604)

* _2026-06-11 08:11:20 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16106/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2064983681084166144#1:build-log.txt%3A796)

* _2026-06-11 07:44:55 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18094/pull-kubevirt-e2e-k8s-1.35-sig-storage/2064976950866219008#1:build-log.txt%3A598)

* _2026-06-11 07:44:48 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18094/pull-kubevirt-e2e-k8s-1.34-windows2016/2064976949725368320#1:build-log.txt%3A614)

* _2026-06-11 07:44:47 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18094/pull-kubevirt-e2e-k8s-1.34-sig-storage/2064976950388068352#1:build-log.txt%3A600)

* _2026-06-11 06:58:14 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17651/pull-kubevirt-e2e-k8s-1.35-sig-storage/2064965270971092992#1:build-log.txt%3A616)

* _2026-06-11 06:58:14 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17651/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2064965270748794880#1:build-log.txt%3A695)

* _2026-06-11 06:58:13 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17651/pull-kubevirt-e2e-k8s-1.34-windows2016/2064965270564245504#1:build-log.txt%3A632)

* _2026-06-11 06:58:13 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17651/pull-kubevirt-e2e-k8s-1.36-sig-storage/2064965271050784768#1:build-log.txt%3A616)

* _2026-06-11 06:58:12 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17651/pull-kubevirt-e2e-k8s-1.34-sig-storage/2064965270895595520#1:build-log.txt%3A615)

* _2026-06-11 05:48:18 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.34-windows2016/2064947674259394560#1:build-log.txt%3A698)

* _2026-06-11 05:48:17 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.34-sig-storage/2064947675047923712#1:build-log.txt%3A642)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache blob fetch failure (5x / 18.52%) </summary>

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
<summary> failed external fetch in context (2x / 7.41%) </summary>

<hr/>

**2x**: _2026-06-11 05:48:18 &#43;0000 UTC_: <code>05:52:48: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.35-sig-operator/2064947676004225024#1:build-log.txt%3A411)
<details>
<summary>all...</summary>

* _2026-06-11 05:48:18 &#43;0000 UTC_: <code>05:52:48: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.35-sig-operator/2064947676004225024#1:build-log.txt%3A411)

* _2026-06-11 05:48:17 &#43;0000 UTC_: <code>05:52:44: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.35-sig-storage/2064947675668680704#1:build-log.txt%3A417)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 3.70%) </summary>

<hr/>

**1x**: _2026-06-12 06:47:52 &#43;0000 UTC_: <code>06:49:59: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18035/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065325019264716800#1:build-log.txt%3A652)
<details>
<summary>all...</summary>

* _2026-06-12 06:47:52 &#43;0000 UTC_: <code>06:49:59: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18035/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065325019264716800#1:build-log.txt%3A652)

</details>

<hr/>
</details>
<details>
<summary> download failure from external URL (1x / 3.70%) </summary>

<hr/>

**1x**: _2026-06-14 06:31:19 &#43;0000 UTC_: <code>06:32:23: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp3938845111312673746/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17913/pull-kubevirt-e2e-k8s-1.35-sig-storage/2066045650159538176#1:build-log.txt%3A311)
<details>
<summary>all...</summary>

* _2026-06-14 06:31:19 &#43;0000 UTC_: <code>06:32:23: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp3938845111312673746/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17913/pull-kubevirt-e2e-k8s-1.35-sig-storage/2066045650159538176#1:build-log.txt%3A311)

</details>

<hr/>
</details>

#### needs-investigation (1x / 3.70%)

<details>
<summary> no matching pattern (1x / 3.70%) </summary>

<hr/>

**1x**: _2026-06-13 11:08:07 &#43;0000 UTC_: <code>11:09:44: ERROR: Error initializing RemoteModule</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18034/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065748056631939072#1:build-log.txt%3A347)
<details>
<summary>all...</summary>

* _2026-06-13 11:08:07 &#43;0000 UTC_: <code>11:09:44: ERROR: Error initializing RemoteModule</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18034/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065748056631939072#1:build-log.txt%3A347)

</details>

<hr/>
</details>

### release-1.8 (11x / 26.83%)


#### external (10x / 90.91%)

<details>
<summary> HTTP error in context (9x / 81.82%) </summary>

<hr/>

**9x**: _2026-06-11 07:23:11 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18056/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2064971560589987840#1:build-log.txt%3A623)
<details>
<summary>all...</summary>

* _2026-06-11 07:58:25 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18055/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2064980396310794240#1:build-log.txt%3A623)

* _2026-06-11 07:24:17 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18055/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations-1.8/2064971753305673728#1:build-log.txt%3A686)

* _2026-06-11 07:24:13 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18055/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2064971753439891456#1:build-log.txt%3A609)

* _2026-06-11 07:24:01 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18055/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2064971753741881344#1:build-log.txt%3A608)

* _2026-06-11 07:24:00 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18055/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2064971753607663616#1:build-log.txt%3A606)

* _2026-06-11 07:23:14 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18056/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2064971560803897344#1:build-log.txt%3A606)

* _2026-06-11 07:23:13 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18056/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2064971560724205568#1:build-log.txt%3A611)

* _2026-06-11 07:23:11 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18056/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2064971560589987840#1:build-log.txt%3A623)

* _2026-06-11 07:23:11 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18056/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2064971560891977728#1:build-log.txt%3A606)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 9.09%) </summary>

<hr/>

**1x**: _2026-06-08 13:09:50 &#43;0000 UTC_: <code>15:08:10:   {&#34;[namespace kubevirt-test-default1 name testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid d112888e-d01a-47a1-af22-ff1f17598cf2]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-06-08T15:07:46.018950Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-08T15:07:46.055406Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2063971602751557632#1:build-log.txt%3A6872)
<details>
<summary>all...</summary>

* _2026-06-08 13:09:50 &#43;0000 UTC_: <code>15:08:10:   {&#34;[namespace kubevirt-test-default1 name testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid d112888e-d01a-47a1-af22-ff1f17598cf2]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-06-08T15:07:46.018950Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-08T15:07:46.055406Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2063971602751557632#1:build-log.txt%3A6872)

</details>

<hr/>
</details>

#### needs-investigation (1x / 9.09%)

<details>
<summary> no matching pattern (1x / 9.09%) </summary>

<hr/>

**1x**: _2026-06-08 13:09:41 &#43;0000 UTC_: <code>13:29:32: failed to create new CRI image service: validate service connection: validate CRI v1 image API for endpoint &#34;unix:///var/run/crio/crio.sock&#34;: rpc error: code = DeadlineExceeded desc = stream terminated by RST_STREAM with error code: CANCEL[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2063971600276918272#1:build-log.txt%3A674)
<details>
<summary>all...</summary>

* _2026-06-08 13:09:41 &#43;0000 UTC_: <code>13:29:32: failed to create new CRI image service: validate service connection: validate CRI v1 image API for endpoint &#34;unix:///var/run/crio/crio.sock&#34;: rpc error: code = DeadlineExceeded desc = stream terminated by RST_STREAM with error code: CANCEL[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2063971600276918272#1:build-log.txt%3A674)

</details>

<hr/>
</details>

### release-1.7 (3x / 7.32%)


#### internal (2x / 66.67%)

<details>
<summary> make cluster lifecycle target failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-06-09 12:55:48 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2064329652029100032#1:build-log.txt%3A1196)
<details>
<summary>all...</summary>

* _2026-06-09 12:55:48 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2064329652029100032#1:build-log.txt%3A1196)

</details>

<hr/>
</details>
<details>
<summary> make bazel-build target failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-06-09 08:35:56 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2064264517830840320#1:build-log.txt%3A4515)
<details>
<summary>all...</summary>

* _2026-06-09 08:35:56 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2064264517830840320#1:build-log.txt%3A4515)

</details>

<hr/>
</details>

#### external (1x / 33.33%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-06-09 08:33:34 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2064264500244123648#1:build-log.txt%3A4384)
<details>
<summary>all...</summary>

* _2026-06-09 08:33:34 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2064264500244123648#1:build-log.txt%3A4384)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-storage (19x / 46.34%)


#### external (19x / 100.00%)

<details>
<summary> HTTP error in context (16x / 84.21%) </summary>

<hr/>

**16x**: _2026-06-11 07:23:13 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18056/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2064971560724205568#1:build-log.txt%3A611)
<details>
<summary>all...</summary>

* _2026-06-12 09:50:33 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.36-sig-storage/2065370831113949184#1:build-log.txt%3A611)

* _2026-06-12 09:49:35 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.34-sig-storage/2065370810171789312#1:build-log.txt%3A607)

* _2026-06-11 08:43:11 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.35-sig-storage/2064991677696708608#1:build-log.txt%3A604)

* _2026-06-11 08:43:11 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.34-sig-storage/2064991677570879488#1:build-log.txt%3A603)

* _2026-06-11 07:44:55 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18094/pull-kubevirt-e2e-k8s-1.35-sig-storage/2064976950866219008#1:build-log.txt%3A598)

* _2026-06-11 07:44:47 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18094/pull-kubevirt-e2e-k8s-1.34-sig-storage/2064976950388068352#1:build-log.txt%3A600)

* _2026-06-11 07:24:13 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18055/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2064971753439891456#1:build-log.txt%3A609)

* _2026-06-11 07:24:01 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18055/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2064971753741881344#1:build-log.txt%3A608)

* _2026-06-11 07:24:00 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18055/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2064971753607663616#1:build-log.txt%3A606)

* _2026-06-11 07:23:14 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18056/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2064971560803897344#1:build-log.txt%3A606)

* _2026-06-11 07:23:13 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18056/pull-kubevirt-e2e-k8s-1.33-sig-storage-1.8/2064971560724205568#1:build-log.txt%3A611)

* _2026-06-11 07:23:11 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18056/pull-kubevirt-e2e-k8s-1.35-sig-storage-1.8/2064971560891977728#1:build-log.txt%3A606)

* _2026-06-11 06:58:14 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17651/pull-kubevirt-e2e-k8s-1.35-sig-storage/2064965270971092992#1:build-log.txt%3A616)

* _2026-06-11 06:58:13 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17651/pull-kubevirt-e2e-k8s-1.36-sig-storage/2064965271050784768#1:build-log.txt%3A616)

* _2026-06-11 06:58:12 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17651/pull-kubevirt-e2e-k8s-1.34-sig-storage/2064965270895595520#1:build-log.txt%3A615)

* _2026-06-11 05:48:17 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.34-sig-storage/2064947675047923712#1:build-log.txt%3A642)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 5.26%) </summary>

<hr/>

**1x**: _2026-06-08 13:09:50 &#43;0000 UTC_: <code>15:08:10:   {&#34;[namespace kubevirt-test-default1 name testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid d112888e-d01a-47a1-af22-ff1f17598cf2]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-06-08T15:07:46.018950Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-08T15:07:46.055406Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2063971602751557632#1:build-log.txt%3A6872)
<details>
<summary>all...</summary>

* _2026-06-08 13:09:50 &#43;0000 UTC_: <code>15:08:10:   {&#34;[namespace kubevirt-test-default1 name testvmi-2kxr4-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx kind VirtualMachineInstance uid d112888e-d01a-47a1-af22-ff1f17598cf2]&#34;:&#34;(MISSING)&#34;,&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;server error. command SyncVMI failed: \&#34;LibvirtError(Code=1, Domain=10, Message=&#39;internal error: process exited while connecting to monitor: 2026-06-08T15:07:46.018950Z qemu-kvm: -blockdev {\\\&#34;driver\\\&#34;:\\\&#34;file\\\&#34;,\\\&#34;filename\\\&#34;:\\\&#34;/var/run/kubevirt-private/vmi-disks/disk0/disk.img\\\&#34;,\\\&#34;node-name\\\&#34;:\\\&#34;libvirt-1-storage\\\&#34;,\\\&#34;read-only\\\&#34;:false,\\\&#34;discard\\\&#34;:\\\&#34;unmap\\\&#34;,\\\&#34;cache\\\&#34;:{\\\&#34;direct\\\&#34;:false,\\\&#34;no-flush\\\&#34;:false}}: Could not open &#39;/var/run/kubevirt-private/vmi-disks/disk0/disk.img&#39;: Permission denied&#39;)\&#34;&#34;,&#34;pos&#34;:&#34;watcher.go:157&#34;,&#34;reason&#34;:&#34;warning event received&#34;,&#34;timestamp&#34;:&#34;2026-06-08T15:07:46.055406Z&#34;}</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2063971602751557632#1:build-log.txt%3A6872)

</details>

<hr/>
</details>
<details>
<summary> download failure from external URL (1x / 5.26%) </summary>

<hr/>

**1x**: _2026-06-14 06:31:19 &#43;0000 UTC_: <code>06:32:23: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp3938845111312673746/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17913/pull-kubevirt-e2e-k8s-1.35-sig-storage/2066045650159538176#1:build-log.txt%3A311)
<details>
<summary>all...</summary>

* _2026-06-14 06:31:19 &#43;0000 UTC_: <code>06:32:23: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp3938845111312673746/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17913/pull-kubevirt-e2e-k8s-1.35-sig-storage/2066045650159538176#1:build-log.txt%3A311)

</details>

<hr/>
</details>
<details>
<summary> failed external fetch in context (1x / 5.26%) </summary>

<hr/>

**1x**: _2026-06-11 05:48:17 &#43;0000 UTC_: <code>05:52:44: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.35-sig-storage/2064947675668680704#1:build-log.txt%3A417)
<details>
<summary>all...</summary>

* _2026-06-11 05:48:17 &#43;0000 UTC_: <code>05:52:44: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.35-sig-storage/2064947675668680704#1:build-log.txt%3A417)

</details>

<hr/>
</details>

### sig-compute (20x / 48.78%)


#### external (18x / 90.00%)

<details>
<summary> HTTP error in context (10x / 50.00%) </summary>

<hr/>

**10x**: _2026-06-11 07:23:11 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18056/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2064971560589987840#1:build-log.txt%3A623)
<details>
<summary>all...</summary>

* _2026-06-12 09:54:34 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18052/pull-kubevirt-e2e-k8s-1.34-windows2016/2065371648579604480#1:build-log.txt%3A627)

* _2026-06-11 08:43:11 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.34-windows2016/2064991677390524416#1:build-log.txt%3A620)

* _2026-06-11 08:11:20 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16106/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2064983681084166144#1:build-log.txt%3A796)

* _2026-06-11 07:58:25 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18055/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2064980396310794240#1:build-log.txt%3A623)

* _2026-06-11 07:44:48 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18094/pull-kubevirt-e2e-k8s-1.34-windows2016/2064976949725368320#1:build-log.txt%3A614)

* _2026-06-11 07:24:17 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18055/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations-1.8/2064971753305673728#1:build-log.txt%3A686)

* _2026-06-11 07:23:11 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18056/pull-kubevirt-e2e-k8s-1.34-windows2016-1.8/2064971560589987840#1:build-log.txt%3A623)

* _2026-06-11 06:58:14 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17651/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2064965270748794880#1:build-log.txt%3A695)

* _2026-06-11 06:58:13 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17651/pull-kubevirt-e2e-k8s-1.34-windows2016/2064965270564245504#1:build-log.txt%3A632)

* _2026-06-11 05:48:18 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.34-windows2016/2064947674259394560#1:build-log.txt%3A698)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache blob fetch failure (5x / 25.00%) </summary>

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
<summary> failed external fetch in context (1x / 5.00%) </summary>

<hr/>

**1x**: _2026-06-11 05:48:18 &#43;0000 UTC_: <code>05:52:48: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.35-sig-operator/2064947676004225024#1:build-log.txt%3A411)
<details>
<summary>all...</summary>

* _2026-06-11 05:48:18 &#43;0000 UTC_: <code>05:52:48: ERROR: Analysis of target &#39;//images/winrmcli:winrmcli-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17845/pull-kubevirt-e2e-k8s-1.35-sig-operator/2064947676004225024#1:build-log.txt%3A411)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 5.00%) </summary>

<hr/>

**1x**: _2026-06-12 06:47:52 &#43;0000 UTC_: <code>06:49:59: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18035/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065325019264716800#1:build-log.txt%3A652)
<details>
<summary>all...</summary>

* _2026-06-12 06:47:52 &#43;0000 UTC_: <code>06:49:59: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18035/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065325019264716800#1:build-log.txt%3A652)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 5.00%) </summary>

<hr/>

**1x**: _2026-06-09 08:33:34 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2064264500244123648#1:build-log.txt%3A4384)
<details>
<summary>all...</summary>

* _2026-06-09 08:33:34 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2064264500244123648#1:build-log.txt%3A4384)

</details>

<hr/>
</details>

#### needs-investigation (2x / 10.00%)

<details>
<summary> no matching pattern (2x / 10.00%) </summary>

<hr/>

**1x**: _2026-06-13 11:08:07 &#43;0000 UTC_: <code>11:09:44: ERROR: Error initializing RemoteModule</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18034/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065748056631939072#1:build-log.txt%3A347)
<details>
<summary>all...</summary>

* _2026-06-13 11:08:07 &#43;0000 UTC_: <code>11:09:44: ERROR: Error initializing RemoteModule</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18034/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2065748056631939072#1:build-log.txt%3A347)

</details>

<hr/>

**1x**: _2026-06-08 13:09:41 &#43;0000 UTC_: <code>13:29:32: failed to create new CRI image service: validate service connection: validate CRI v1 image API for endpoint &#34;unix:///var/run/crio/crio.sock&#34;: rpc error: code = DeadlineExceeded desc = stream terminated by RST_STREAM with error code: CANCEL[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2063971600276918272#1:build-log.txt%3A674)
<details>
<summary>all...</summary>

* _2026-06-08 13:09:41 &#43;0000 UTC_: <code>13:29:32: failed to create new CRI image service: validate service connection: validate CRI v1 image API for endpoint &#34;unix:///var/run/crio/crio.sock&#34;: rpc error: code = DeadlineExceeded desc = stream terminated by RST_STREAM with error code: CANCEL[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17962/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2063971600276918272#1:build-log.txt%3A674)

</details>

<hr/>
</details>

### sig-network (2x / 4.88%)


#### internal (2x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-06-09 12:55:48 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2064329652029100032#1:build-log.txt%3A1196)
<details>
<summary>all...</summary>

* _2026-06-09 12:55:48 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.32-sig-network-1.7/2064329652029100032#1:build-log.txt%3A1196)

</details>

<hr/>
</details>
<details>
<summary> make bazel-build target failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-06-09 08:35:56 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2064264517830840320#1:build-log.txt%3A4515)
<details>
<summary>all...</summary>

* _2026-06-09 08:35:56 &#43;0000 UTC_: <code>make: *** [Makefile:26: bazel-build-functests] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17965/pull-kubevirt-e2e-k8s-1.34-sig-network-1.7/2064264517830840320#1:build-log.txt%3A4515)

</details>

<hr/>
</details>

Last updated: 2026-06-15 11:17:21
