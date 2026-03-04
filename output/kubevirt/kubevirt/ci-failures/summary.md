

# CI failures for kubevirt/kubevirt

## per branch (7x)

<details>
<summary> main (6x / 85.71%) </summary>

<hr/>

**3x**: _2026-02-18 20:51:17 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16604/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2024225096696074240#1:build-log.txt%3A304)
<details>
<summary>all...</summary>

* _2026-02-18 20:51:17 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16604/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2024225096696074240#1:build-log.txt%3A304)

* _2026-02-18 20:50:11 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16659/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2024224812649418752#1:build-log.txt%3A267)

* _2026-02-18 20:40:28 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16812/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2024222098527883264#1:build-log.txt%3A290)

</details>

<hr/>

**1x**: _2026-02-26 20:46:13 &#43;0000 UTC_: <code>20:48:19: ERROR: Analysis of target &#39;//cmd/virt-api:virt-api-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16604/pull-kubevirt-e2e-k8s-1.33-sig-storage/2027074770331242496#1:build-log.txt%3A487)
<details>
<summary>all...</summary>

* _2026-02-26 20:46:13 &#43;0000 UTC_: <code>20:48:19: ERROR: Analysis of target &#39;//cmd/virt-api:virt-api-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16604/pull-kubevirt-e2e-k8s-1.33-sig-storage/2027074770331242496#1:build-log.txt%3A487)

</details>

<hr/>

**1x**: _2026-02-10 09:23:25 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16659/pull-kubevirt-e2e-k8s-1.33-sig-compute/2021152888784424960#1:build-log.txt%3A984)
<details>
<summary>all...</summary>

* _2026-02-10 09:23:25 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16659/pull-kubevirt-e2e-k8s-1.33-sig-compute/2021152888784424960#1:build-log.txt%3A984)

</details>

<hr/>

**1x**: _2026-02-27 09:03:02 &#43;0000 UTC_: <code>09:21:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16776/pull-kubevirt-e2e-k8s-1.35-sig-operator/2027308341545406464#1:build-log.txt%3A4616)
<details>
<summary>all...</summary>

* _2026-02-27 09:03:02 &#43;0000 UTC_: <code>09:21:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16776/pull-kubevirt-e2e-k8s-1.35-sig-operator/2027308341545406464#1:build-log.txt%3A4616)

</details>

<hr/>
</details>
<details>
<summary> release-1.7 (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-01-20 10:59:02 &#43;0000 UTC_: <code>11:24:04: I0120 06:24:04.513526    1600 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16588/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.7/2013566717766144000#1:build-log.txt%3A769)
<details>
<summary>all...</summary>

* _2026-01-20 10:59:02 &#43;0000 UTC_: <code>11:24:04: I0120 06:24:04.513526    1600 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16588/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.7/2013566717766144000#1:build-log.txt%3A769)

</details>

<hr/>
</details>

## per day (7x)

<details>
<summary> 2026-02-18 (3x / 42.86%) </summary>

<hr/>

**3x**: _2026-02-18 20:51:17 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16604/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2024225096696074240#1:build-log.txt%3A304)
<details>
<summary>all...</summary>

* _2026-02-18 20:51:17 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16604/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2024225096696074240#1:build-log.txt%3A304)

* _2026-02-18 20:50:11 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16659/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2024224812649418752#1:build-log.txt%3A267)

* _2026-02-18 20:40:28 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16812/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2024222098527883264#1:build-log.txt%3A290)

</details>

<hr/>
</details>
<details>
<summary> 2026-01-20 (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-01-20 10:59:02 &#43;0000 UTC_: <code>11:24:04: I0120 06:24:04.513526    1600 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16588/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.7/2013566717766144000#1:build-log.txt%3A769)
<details>
<summary>all...</summary>

* _2026-01-20 10:59:02 &#43;0000 UTC_: <code>11:24:04: I0120 06:24:04.513526    1600 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16588/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.7/2013566717766144000#1:build-log.txt%3A769)

</details>

<hr/>
</details>
<details>
<summary> 2026-02-26 (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-02-26 20:46:13 &#43;0000 UTC_: <code>20:48:19: ERROR: Analysis of target &#39;//cmd/virt-api:virt-api-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16604/pull-kubevirt-e2e-k8s-1.33-sig-storage/2027074770331242496#1:build-log.txt%3A487)
<details>
<summary>all...</summary>

* _2026-02-26 20:46:13 &#43;0000 UTC_: <code>20:48:19: ERROR: Analysis of target &#39;//cmd/virt-api:virt-api-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16604/pull-kubevirt-e2e-k8s-1.33-sig-storage/2027074770331242496#1:build-log.txt%3A487)

</details>

<hr/>
</details>
<details>
<summary> 2026-02-10 (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-02-10 09:23:25 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16659/pull-kubevirt-e2e-k8s-1.33-sig-compute/2021152888784424960#1:build-log.txt%3A984)
<details>
<summary>all...</summary>

* _2026-02-10 09:23:25 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16659/pull-kubevirt-e2e-k8s-1.33-sig-compute/2021152888784424960#1:build-log.txt%3A984)

</details>

<hr/>
</details>
<details>
<summary> 2026-02-27 (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-02-27 09:03:02 &#43;0000 UTC_: <code>09:21:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16776/pull-kubevirt-e2e-k8s-1.35-sig-operator/2027308341545406464#1:build-log.txt%3A4616)
<details>
<summary>all...</summary>

* _2026-02-27 09:03:02 &#43;0000 UTC_: <code>09:21:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16776/pull-kubevirt-e2e-k8s-1.35-sig-operator/2027308341545406464#1:build-log.txt%3A4616)

</details>

<hr/>
</details>

## per SIG (7x)

<details>
<summary> sig-compute (5x / 71.43%) </summary>

<hr/>

**3x**: _2026-02-18 20:51:17 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16604/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2024225096696074240#1:build-log.txt%3A304)
<details>
<summary>all...</summary>

* _2026-02-18 20:51:17 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16604/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2024225096696074240#1:build-log.txt%3A304)

* _2026-02-18 20:50:11 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16659/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2024224812649418752#1:build-log.txt%3A267)

* _2026-02-18 20:40:28 &#43;0000 UTC_: <code>make: *** [Makefile:37: bazel-build-images] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16812/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2024222098527883264#1:build-log.txt%3A290)

</details>

<hr/>

**1x**: _2026-02-10 09:23:25 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16659/pull-kubevirt-e2e-k8s-1.33-sig-compute/2021152888784424960#1:build-log.txt%3A984)
<details>
<summary>all...</summary>

* _2026-02-10 09:23:25 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16659/pull-kubevirt-e2e-k8s-1.33-sig-compute/2021152888784424960#1:build-log.txt%3A984)

</details>

<hr/>

**1x**: _2026-02-27 09:03:02 &#43;0000 UTC_: <code>09:21:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16776/pull-kubevirt-e2e-k8s-1.35-sig-operator/2027308341545406464#1:build-log.txt%3A4616)
<details>
<summary>all...</summary>

* _2026-02-27 09:03:02 &#43;0000 UTC_: <code>09:21:03: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16776/pull-kubevirt-e2e-k8s-1.35-sig-operator/2027308341545406464#1:build-log.txt%3A4616)

</details>

<hr/>
</details>
<details>
<summary> sig-monitoring (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-01-20 10:59:02 &#43;0000 UTC_: <code>11:24:04: I0120 06:24:04.513526    1600 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16588/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.7/2013566717766144000#1:build-log.txt%3A769)
<details>
<summary>all...</summary>

* _2026-01-20 10:59:02 &#43;0000 UTC_: <code>11:24:04: I0120 06:24:04.513526    1600 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16588/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.7/2013566717766144000#1:build-log.txt%3A769)

</details>

<hr/>
</details>
<details>
<summary> sig-storage (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-02-26 20:46:13 &#43;0000 UTC_: <code>20:48:19: ERROR: Analysis of target &#39;//cmd/virt-api:virt-api-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16604/pull-kubevirt-e2e-k8s-1.33-sig-storage/2027074770331242496#1:build-log.txt%3A487)
<details>
<summary>all...</summary>

* _2026-02-26 20:46:13 &#43;0000 UTC_: <code>20:48:19: ERROR: Analysis of target &#39;//cmd/virt-api:virt-api-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16604/pull-kubevirt-e2e-k8s-1.33-sig-storage/2027074770331242496#1:build-log.txt%3A487)

</details>

<hr/>
</details>

Last updated: 2026-03-04 07:33:42
