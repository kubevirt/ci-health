



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-05-18 (2x / 9.52%)


#### pr-build (2x / 100.00%)

<details>
<summary> ginkgo test failure marker (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-05-18 07:45:55 &#43;0000 UTC_: <code>11:09:36: • [FAILED] [570.279 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17828/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2056279962301763584#1:build-log.txt%3A9570)
<details>
<summary>all...</summary>

* _2026-05-18 07:52:18 &#43;0000 UTC_: <code>11:09:10: • [FAILED] [549.986 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17829/pull-kubevirt-e2e-k8s-1.35-sig-storage/2056281442882686976#1:build-log.txt%3A9108)

* _2026-05-18 07:45:55 &#43;0000 UTC_: <code>11:09:36: • [FAILED] [570.279 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17828/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2056279962301763584#1:build-log.txt%3A9570)

</details>

<hr/>
</details>

### 2026-05-15 (5x / 23.81%)


#### needs-investigation (4x / 80.00%)

<details>
<summary> no error snippets (4x / 80.00%) </summary>

<hr/>

**1x**: _2026-05-15 11:34:55 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16257/pull-kubevirt-e2e-k8s-1.35-sig-operator/2055241605253697536)
<details>
<summary>all...</summary>

* _2026-05-15 11:34:55 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16257/pull-kubevirt-e2e-k8s-1.35-sig-operator/2055241605253697536)

</details>

<hr/>

**1x**: _2026-05-15 08:44:00 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17351/pull-kubevirt-e2e-k8s-1.34-sig-storage/2055207418115133440)
<details>
<summary>all...</summary>

* _2026-05-15 08:44:00 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17351/pull-kubevirt-e2e-k8s-1.34-sig-storage/2055207418115133440)

</details>

<hr/>

**1x**: _2026-05-15 07:29:39 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.35-sig-storage/2055186770911301632)
<details>
<summary>all...</summary>

* _2026-05-15 07:29:39 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.35-sig-storage/2055186770911301632)

</details>

<hr/>

**1x**: _2026-05-15 07:29:22 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.34-sig-operator/2055186770034692096)
<details>
<summary>all...</summary>

* _2026-05-15 07:29:22 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.34-sig-operator/2055186770034692096)

</details>

<hr/>
</details>

#### external (1x / 20.00%)

<details>
<summary> transient kube-apiserver body decode noise (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-05-15 04:52:36 &#43;0000 UTC_: <code>04:56:39: I0515 00:56:39.079139    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-network/2055146025328316416#1:build-log.txt%3A753)
<details>
<summary>all...</summary>

* _2026-05-15 04:52:36 &#43;0000 UTC_: <code>04:56:39: I0515 00:56:39.079139    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-network/2055146025328316416#1:build-log.txt%3A753)

</details>

<hr/>
</details>

### 2026-05-14 (4x / 19.05%)


#### needs-investigation (3x / 75.00%)

<details>
<summary> no error snippets (3x / 75.00%) </summary>

<hr/>

**1x**: _2026-05-14 09:41:14 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054858348410441728)
<details>
<summary>all...</summary>

* _2026-05-14 09:41:14 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054858348410441728)

</details>

<hr/>

**1x**: _2026-05-14 08:35:19 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054827079924453376)
<details>
<summary>all...</summary>

* _2026-05-14 08:35:19 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054827079924453376)

</details>

<hr/>

**1x**: _2026-05-14 03:11:53 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054750758351409152)
<details>
<summary>all...</summary>

* _2026-05-14 03:11:53 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054750758351409152)

</details>

<hr/>
</details>

#### pr-build (1x / 25.00%)

<details>
<summary> ginkgo test failure marker (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-05-14 18:29:24 &#43;0000 UTC_: <code>20:18:29: • [FAILED] [269.183 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17801/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.6/2054959179134668800#1:build-log.txt%3A4093)
<details>
<summary>all...</summary>

* _2026-05-14 18:29:24 &#43;0000 UTC_: <code>20:18:29: • [FAILED] [269.183 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17801/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.6/2054959179134668800#1:build-log.txt%3A4093)

</details>

<hr/>
</details>

### 2026-05-13 (10x / 47.62%)


#### needs-investigation (8x / 80.00%)

<details>
<summary> no error snippets (8x / 80.00%) </summary>

<hr/>

**1x**: _2026-05-13 20:39:40 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054662742870069248)
<details>
<summary>all...</summary>

* _2026-05-13 20:39:40 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054662742870069248)

</details>

<hr/>

**1x**: _2026-05-13 19:09:36 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054551663221411840)
<details>
<summary>all...</summary>

* _2026-05-13 19:09:36 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054551663221411840)

</details>

<hr/>

**1x**: _2026-05-13 13:19:37 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054551663053639680)
<details>
<summary>all...</summary>

* _2026-05-13 13:19:37 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054551663053639680)

</details>

<hr/>

**1x**: _2026-05-13 12:02:10 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17794/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054532503019982848)
<details>
<summary>all...</summary>

* _2026-05-13 12:02:10 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17794/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054532503019982848)

</details>

<hr/>

**1x**: _2026-05-13 09:33:56 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054495210582315008)
<details>
<summary>all...</summary>

* _2026-05-13 09:33:56 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054495210582315008)

</details>

<hr/>

**1x**: _2026-05-13 09:01:51 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-network/2054487143505465344)
<details>
<summary>all...</summary>

* _2026-05-13 09:01:51 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-network/2054487143505465344)

</details>

<hr/>

**1x**: _2026-05-13 06:33:53 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17680/pull-kubevirt-e2e-k8s-1.34-sig-network/2054449906545856512)
<details>
<summary>all...</summary>

* _2026-05-13 06:33:53 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17680/pull-kubevirt-e2e-k8s-1.34-sig-network/2054449906545856512)

</details>

<hr/>

**1x**: _2026-05-13 06:33:53 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17680/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054449906466164736)
<details>
<summary>all...</summary>

* _2026-05-13 06:33:53 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17680/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054449906466164736)

</details>

<hr/>
</details>

#### external (2x / 20.00%)

<details>
<summary> transient kube-apiserver body decode noise (2x / 20.00%) </summary>

<hr/>

**2x**: _2026-05-13 16:43:43 &#43;0000 UTC_: <code>16:49:17: I0513 12:49:17.114198    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054569150063316992#1:build-log.txt%3A783)
<details>
<summary>all...</summary>

* _2026-05-13 16:43:43 &#43;0000 UTC_: <code>16:49:17: I0513 12:49:17.114198    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054569150063316992#1:build-log.txt%3A783)

* _2026-05-13 15:06:33 &#43;0000 UTC_: <code>15:12:22: I0513 11:12:22.171698    1602 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-network/2054578675570970624#1:build-log.txt%3A1887)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### pr-build (3x / 14.29%)

<details>
<summary> ginkgo test failure marker (3x / 14.29%) </summary>

<hr/>

**2x**: _2026-05-18 07:45:55 &#43;0000 UTC_: <code>11:09:36: • [FAILED] [570.279 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17828/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2056279962301763584#1:build-log.txt%3A9570)
<details>
<summary>all...</summary>

* _2026-05-18 07:52:18 &#43;0000 UTC_: <code>11:09:10: • [FAILED] [549.986 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17829/pull-kubevirt-e2e-k8s-1.35-sig-storage/2056281442882686976#1:build-log.txt%3A9108)
<details><summary>context</summary>
<pre>11:00:00:   {&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;system is in sync with kubevirt config resource version 206657&#34;,&#34;pos&#34;:&#34;kvconfig.go:101&#34;,&#34;timestamp&#34;:&#34;2026-05-18T11:00:00.128573Z&#34;}
11:00:00:   &lt;&lt; Captured StdOut/StdErr Output
11:00:00: ------------------------------
11:09:10: • [FAILED] [549.986 seconds]
11:09:10: [sig-storage] Hotplug with PCI hostdev [It] should restart a VM after hotplugging a block volume [sig-storage, Serial, RequiresBlockStorage]
11:09:10: tests/storage/hotplug.go:2216
11:09:10:</pre>
</details>


* _2026-05-18 07:45:55 &#43;0000 UTC_: <code>11:09:36: • [FAILED] [570.279 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17828/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2056279962301763584#1:build-log.txt%3A9570)
<details><summary>context</summary>
<pre>11:00:05:   {&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;system is in sync with kubevirt config resource version 200331&#34;,&#34;pos&#34;:&#34;kvconfig.go:102&#34;,&#34;timestamp&#34;:&#34;2026-05-18T11:00:05.879219Z&#34;}
11:00:05:   &lt;&lt; Captured StdOut/StdErr Output
11:00:05: ------------------------------
11:09:36: • [FAILED] [570.279 seconds]
11:09:36: [sig-storage] SCSI persistent reservation [BeforeEach] with PersistentReservation feature gate toggled should delete and recreate virt-handler [sig-storage, Serial]
11:09:36:   [BeforeEach] tests/storage/reservation.go:186
11:09:36:   [It] tests/storage/reservation.go:344</pre>
</details>


</details>

<hr/>

**1x**: _2026-05-14 18:29:24 &#43;0000 UTC_: <code>20:18:29: • [FAILED] [269.183 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17801/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.6/2054959179134668800#1:build-log.txt%3A4093)
<details>
<summary>all...</summary>

* _2026-05-14 18:29:24 &#43;0000 UTC_: <code>20:18:29: • [FAILED] [269.183 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17801/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.6/2054959179134668800#1:build-log.txt%3A4093)
<details><summary>context</summary>
<pre>20:14:00:
20:14:00:   There were additional failures detected.  To view them in detail run ginkgo -vv
20:14:00: ------------------------------
20:18:29: • [FAILED] [269.183 seconds]
20:18:29: [rfe_id:393][crit:high][vendor:cnv-qe@redhat.com][level:system][sig-compute] Live Migration with a live-migrate eviction strategy set [ref_id:2293</pre>
</details>


</details>

<hr/>
</details>

### needs-investigation (15x / 71.43%)

<details>
<summary> no error snippets (15x / 71.43%) </summary>

<hr/>

**1x**: _2026-05-15 11:34:55 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16257/pull-kubevirt-e2e-k8s-1.35-sig-operator/2055241605253697536)
<details>
<summary>all...</summary>

* _2026-05-15 11:34:55 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16257/pull-kubevirt-e2e-k8s-1.35-sig-operator/2055241605253697536)

</details>

<hr/>

**1x**: _2026-05-15 08:44:00 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17351/pull-kubevirt-e2e-k8s-1.34-sig-storage/2055207418115133440)
<details>
<summary>all...</summary>

* _2026-05-15 08:44:00 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17351/pull-kubevirt-e2e-k8s-1.34-sig-storage/2055207418115133440)

</details>

<hr/>

**1x**: _2026-05-15 07:29:39 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.35-sig-storage/2055186770911301632)
<details>
<summary>all...</summary>

* _2026-05-15 07:29:39 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.35-sig-storage/2055186770911301632)

</details>

<hr/>

**1x**: _2026-05-15 07:29:22 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.34-sig-operator/2055186770034692096)
<details>
<summary>all...</summary>

* _2026-05-15 07:29:22 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.34-sig-operator/2055186770034692096)

</details>

<hr/>

**1x**: _2026-05-14 09:41:14 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054858348410441728)
<details>
<summary>all...</summary>

* _2026-05-14 09:41:14 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054858348410441728)

</details>

<hr/>

**1x**: _2026-05-14 08:35:19 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054827079924453376)
<details>
<summary>all...</summary>

* _2026-05-14 08:35:19 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054827079924453376)

</details>

<hr/>

**1x**: _2026-05-14 03:11:53 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054750758351409152)
<details>
<summary>all...</summary>

* _2026-05-14 03:11:53 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054750758351409152)

</details>

<hr/>

**1x**: _2026-05-13 20:39:40 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054662742870069248)
<details>
<summary>all...</summary>

* _2026-05-13 20:39:40 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054662742870069248)

</details>

<hr/>

**1x**: _2026-05-13 19:09:36 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054551663221411840)
<details>
<summary>all...</summary>

* _2026-05-13 19:09:36 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054551663221411840)

</details>

<hr/>

**1x**: _2026-05-13 13:19:37 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054551663053639680)
<details>
<summary>all...</summary>

* _2026-05-13 13:19:37 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054551663053639680)

</details>

<hr/>

**1x**: _2026-05-13 12:02:10 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17794/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054532503019982848)
<details>
<summary>all...</summary>

* _2026-05-13 12:02:10 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17794/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054532503019982848)

</details>

<hr/>

**1x**: _2026-05-13 09:33:56 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054495210582315008)
<details>
<summary>all...</summary>

* _2026-05-13 09:33:56 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054495210582315008)

</details>

<hr/>

**1x**: _2026-05-13 09:01:51 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-network/2054487143505465344)
<details>
<summary>all...</summary>

* _2026-05-13 09:01:51 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-network/2054487143505465344)

</details>

<hr/>

**1x**: _2026-05-13 06:33:53 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17680/pull-kubevirt-e2e-k8s-1.34-sig-network/2054449906545856512)
<details>
<summary>all...</summary>

* _2026-05-13 06:33:53 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17680/pull-kubevirt-e2e-k8s-1.34-sig-network/2054449906545856512)

</details>

<hr/>

**1x**: _2026-05-13 06:33:53 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17680/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054449906466164736)
<details>
<summary>all...</summary>

* _2026-05-13 06:33:53 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17680/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054449906466164736)

</details>

<hr/>
</details>

### external (3x / 14.29%)

<details>
<summary> transient kube-apiserver body decode noise (3x / 14.29%) </summary>

<hr/>

**3x**: _2026-05-15 04:52:36 &#43;0000 UTC_: <code>04:56:39: I0515 00:56:39.079139    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-network/2055146025328316416#1:build-log.txt%3A753)
<details>
<summary>all...</summary>

* _2026-05-15 04:52:36 &#43;0000 UTC_: <code>04:56:39: I0515 00:56:39.079139    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-network/2055146025328316416#1:build-log.txt%3A753)
<details><summary>context</summary>
<pre>04:56:31: [control-plane-check] Checking kube-scheduler at https://127.0.0.1:10259/livez
04:56:34: [control-plane-check] kube-scheduler is healthy after 3.507988268s
04:56:35: [control-plane-check] kube-controller-manager is healthy after 4.052522673s
04:56:39: I0515 00:56:39.079139    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
04:56:39: [control-plane-check] kube-apiserver is healthy after 8.502646674s
04:56:39: I0515 00:56:39.582491    1610 kubeconfig.go:657] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists
04:56:39: I0515 00:56:39.584077    1610 kubeconfig.go:730] creating the ClusterRoleBinding for the kubeadm:cluster-admins Group by using super-admin.conf</pre>
</details>


* _2026-05-13 16:43:43 &#43;0000 UTC_: <code>16:49:17: I0513 12:49:17.114198    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054569150063316992#1:build-log.txt%3A783)
<details><summary>context</summary>
<pre>16:49:15: I0513 12:49:15.610988    1610 wait.go:283] kube-apiserver check failed at https://[fd00::101]:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;
16:49:16: I0513 12:49:16.113898    1610 wait.go:283] kube-apiserver check failed at https://[fd00::101]:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;
16:49:16: I0513 12:49:16.612987    1610 wait.go:283] kube-apiserver check failed at https://[fd00::101]:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;
16:49:17: I0513 12:49:17.114198    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
16:49:17: I0513 12:49:17.114263    1610 wait.go:283] kube-apiserver check failed at https://[fd00::101]:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
16:49:17: [control-plane-check] kube-apiserver is healthy after 8.003487542s
16:49:17: I0513 12:49:17.618445    1610 kubeconfig.go:657] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists</pre>
</details>


* _2026-05-13 15:06:33 &#43;0000 UTC_: <code>15:12:22: I0513 11:12:22.171698    1602 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-network/2054578675570970624#1:build-log.txt%3A1887)
<details><summary>context</summary>
<pre>15:12:21: I0513 11:12:20.703751    1602 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;
15:12:21: I0513 11:12:21.160993    1602 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;
15:12:22: I0513 11:12:21.658924    1602 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;
15:12:22: I0513 11:12:22.171698    1602 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
15:12:22: I0513 11:12:22.171810    1602 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
15:12:22: [control-plane-check] kube-apiserver is healthy after 5.004516706s
15:12:22: I0513 11:12:22.664856    1602 kubeconfig.go:657] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (19x / 90.48%)


#### needs-investigation (15x / 78.95%)

<details>
<summary> no error snippets (15x / 78.95%) </summary>

<hr/>

**1x**: _2026-05-15 11:34:55 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16257/pull-kubevirt-e2e-k8s-1.35-sig-operator/2055241605253697536)
<details>
<summary>all...</summary>

* _2026-05-15 11:34:55 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16257/pull-kubevirt-e2e-k8s-1.35-sig-operator/2055241605253697536)

</details>

<hr/>

**1x**: _2026-05-15 08:44:00 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17351/pull-kubevirt-e2e-k8s-1.34-sig-storage/2055207418115133440)
<details>
<summary>all...</summary>

* _2026-05-15 08:44:00 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17351/pull-kubevirt-e2e-k8s-1.34-sig-storage/2055207418115133440)

</details>

<hr/>

**1x**: _2026-05-15 07:29:39 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.35-sig-storage/2055186770911301632)
<details>
<summary>all...</summary>

* _2026-05-15 07:29:39 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.35-sig-storage/2055186770911301632)

</details>

<hr/>

**1x**: _2026-05-15 07:29:22 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.34-sig-operator/2055186770034692096)
<details>
<summary>all...</summary>

* _2026-05-15 07:29:22 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.34-sig-operator/2055186770034692096)

</details>

<hr/>

**1x**: _2026-05-14 09:41:14 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054858348410441728)
<details>
<summary>all...</summary>

* _2026-05-14 09:41:14 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054858348410441728)

</details>

<hr/>

**1x**: _2026-05-14 08:35:19 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054827079924453376)
<details>
<summary>all...</summary>

* _2026-05-14 08:35:19 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054827079924453376)

</details>

<hr/>

**1x**: _2026-05-14 03:11:53 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054750758351409152)
<details>
<summary>all...</summary>

* _2026-05-14 03:11:53 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054750758351409152)

</details>

<hr/>

**1x**: _2026-05-13 20:39:40 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054662742870069248)
<details>
<summary>all...</summary>

* _2026-05-13 20:39:40 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054662742870069248)

</details>

<hr/>

**1x**: _2026-05-13 19:09:36 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054551663221411840)
<details>
<summary>all...</summary>

* _2026-05-13 19:09:36 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054551663221411840)

</details>

<hr/>

**1x**: _2026-05-13 13:19:37 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054551663053639680)
<details>
<summary>all...</summary>

* _2026-05-13 13:19:37 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054551663053639680)

</details>

<hr/>

**1x**: _2026-05-13 12:02:10 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17794/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054532503019982848)
<details>
<summary>all...</summary>

* _2026-05-13 12:02:10 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17794/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054532503019982848)

</details>

<hr/>

**1x**: _2026-05-13 09:33:56 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054495210582315008)
<details>
<summary>all...</summary>

* _2026-05-13 09:33:56 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054495210582315008)

</details>

<hr/>

**1x**: _2026-05-13 09:01:51 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-network/2054487143505465344)
<details>
<summary>all...</summary>

* _2026-05-13 09:01:51 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-network/2054487143505465344)

</details>

<hr/>

**1x**: _2026-05-13 06:33:53 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17680/pull-kubevirt-e2e-k8s-1.34-sig-network/2054449906545856512)
<details>
<summary>all...</summary>

* _2026-05-13 06:33:53 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17680/pull-kubevirt-e2e-k8s-1.34-sig-network/2054449906545856512)

</details>

<hr/>

**1x**: _2026-05-13 06:33:53 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17680/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054449906466164736)
<details>
<summary>all...</summary>

* _2026-05-13 06:33:53 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17680/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054449906466164736)

</details>

<hr/>
</details>

#### external (3x / 15.79%)

<details>
<summary> transient kube-apiserver body decode noise (3x / 15.79%) </summary>

<hr/>

**3x**: _2026-05-15 04:52:36 &#43;0000 UTC_: <code>04:56:39: I0515 00:56:39.079139    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-network/2055146025328316416#1:build-log.txt%3A753)
<details>
<summary>all...</summary>

* _2026-05-15 04:52:36 &#43;0000 UTC_: <code>04:56:39: I0515 00:56:39.079139    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-network/2055146025328316416#1:build-log.txt%3A753)

* _2026-05-13 16:43:43 &#43;0000 UTC_: <code>16:49:17: I0513 12:49:17.114198    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054569150063316992#1:build-log.txt%3A783)

* _2026-05-13 15:06:33 &#43;0000 UTC_: <code>15:12:22: I0513 11:12:22.171698    1602 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-network/2054578675570970624#1:build-log.txt%3A1887)

</details>

<hr/>
</details>

#### pr-build (1x / 5.26%)

<details>
<summary> ginkgo test failure marker (1x / 5.26%) </summary>

<hr/>

**1x**: _2026-05-18 07:52:18 &#43;0000 UTC_: <code>11:09:10: • [FAILED] [549.986 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17829/pull-kubevirt-e2e-k8s-1.35-sig-storage/2056281442882686976#1:build-log.txt%3A9108)
<details>
<summary>all...</summary>

* _2026-05-18 07:52:18 &#43;0000 UTC_: <code>11:09:10: • [FAILED] [549.986 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17829/pull-kubevirt-e2e-k8s-1.35-sig-storage/2056281442882686976#1:build-log.txt%3A9108)

</details>

<hr/>
</details>

### release-1.8 (1x / 4.76%)


#### pr-build (1x / 100.00%)

<details>
<summary> ginkgo test failure marker (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-05-18 07:45:55 &#43;0000 UTC_: <code>11:09:36: • [FAILED] [570.279 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17828/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2056279962301763584#1:build-log.txt%3A9570)
<details>
<summary>all...</summary>

* _2026-05-18 07:45:55 &#43;0000 UTC_: <code>11:09:36: • [FAILED] [570.279 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17828/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2056279962301763584#1:build-log.txt%3A9570)

</details>

<hr/>
</details>

### release-1.6 (1x / 4.76%)


#### pr-build (1x / 100.00%)

<details>
<summary> ginkgo test failure marker (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-05-14 18:29:24 &#43;0000 UTC_: <code>20:18:29: • [FAILED] [269.183 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17801/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.6/2054959179134668800#1:build-log.txt%3A4093)
<details>
<summary>all...</summary>

* _2026-05-14 18:29:24 &#43;0000 UTC_: <code>20:18:29: • [FAILED] [269.183 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17801/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.6/2054959179134668800#1:build-log.txt%3A4093)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-storage (7x / 33.33%)


#### pr-build (2x / 28.57%)

<details>
<summary> ginkgo test failure marker (2x / 28.57%) </summary>

<hr/>

**2x**: _2026-05-18 07:45:55 &#43;0000 UTC_: <code>11:09:36: • [FAILED] [570.279 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17828/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2056279962301763584#1:build-log.txt%3A9570)
<details>
<summary>all...</summary>

* _2026-05-18 07:52:18 &#43;0000 UTC_: <code>11:09:10: • [FAILED] [549.986 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17829/pull-kubevirt-e2e-k8s-1.35-sig-storage/2056281442882686976#1:build-log.txt%3A9108)

* _2026-05-18 07:45:55 &#43;0000 UTC_: <code>11:09:36: • [FAILED] [570.279 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17828/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2056279962301763584#1:build-log.txt%3A9570)

</details>

<hr/>
</details>

#### needs-investigation (5x / 71.43%)

<details>
<summary> no error snippets (5x / 71.43%) </summary>

<hr/>

**1x**: _2026-05-15 08:44:00 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17351/pull-kubevirt-e2e-k8s-1.34-sig-storage/2055207418115133440)
<details>
<summary>all...</summary>

* _2026-05-15 08:44:00 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17351/pull-kubevirt-e2e-k8s-1.34-sig-storage/2055207418115133440)

</details>

<hr/>

**1x**: _2026-05-15 07:29:39 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.35-sig-storage/2055186770911301632)
<details>
<summary>all...</summary>

* _2026-05-15 07:29:39 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.35-sig-storage/2055186770911301632)

</details>

<hr/>

**1x**: _2026-05-14 03:11:53 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054750758351409152)
<details>
<summary>all...</summary>

* _2026-05-14 03:11:53 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054750758351409152)

</details>

<hr/>

**1x**: _2026-05-13 20:39:40 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054662742870069248)
<details>
<summary>all...</summary>

* _2026-05-13 20:39:40 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054662742870069248)

</details>

<hr/>

**1x**: _2026-05-13 19:09:36 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054551663221411840)
<details>
<summary>all...</summary>

* _2026-05-13 19:09:36 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-storage/2054551663221411840)

</details>

<hr/>
</details>

### sig-compute (6x / 28.57%)


#### needs-investigation (5x / 83.33%)

<details>
<summary> no error snippets (5x / 83.33%) </summary>

<hr/>

**1x**: _2026-05-15 11:34:55 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16257/pull-kubevirt-e2e-k8s-1.35-sig-operator/2055241605253697536)
<details>
<summary>all...</summary>

* _2026-05-15 11:34:55 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16257/pull-kubevirt-e2e-k8s-1.35-sig-operator/2055241605253697536)

</details>

<hr/>

**1x**: _2026-05-15 07:29:22 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.34-sig-operator/2055186770034692096)
<details>
<summary>all...</summary>

* _2026-05-15 07:29:22 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.34-sig-operator/2055186770034692096)

</details>

<hr/>

**1x**: _2026-05-14 09:41:14 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054858348410441728)
<details>
<summary>all...</summary>

* _2026-05-14 09:41:14 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054858348410441728)

</details>

<hr/>

**1x**: _2026-05-14 08:35:19 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054827079924453376)
<details>
<summary>all...</summary>

* _2026-05-14 08:35:19 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054827079924453376)

</details>

<hr/>

**1x**: _2026-05-13 12:02:10 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17794/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054532503019982848)
<details>
<summary>all...</summary>

* _2026-05-13 12:02:10 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17794/pull-kubevirt-e2e-k8s-1.35-sig-operator/2054532503019982848)

</details>

<hr/>
</details>

#### pr-build (1x / 16.67%)

<details>
<summary> ginkgo test failure marker (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-05-14 18:29:24 &#43;0000 UTC_: <code>20:18:29: • [FAILED] [269.183 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17801/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.6/2054959179134668800#1:build-log.txt%3A4093)
<details>
<summary>all...</summary>

* _2026-05-14 18:29:24 &#43;0000 UTC_: <code>20:18:29: • [FAILED] [269.183 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17801/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.6/2054959179134668800#1:build-log.txt%3A4093)

</details>

<hr/>
</details>

### sig-network (8x / 38.10%)


#### external (3x / 37.50%)

<details>
<summary> transient kube-apiserver body decode noise (3x / 37.50%) </summary>

<hr/>

**3x**: _2026-05-15 04:52:36 &#43;0000 UTC_: <code>04:56:39: I0515 00:56:39.079139    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-network/2055146025328316416#1:build-log.txt%3A753)
<details>
<summary>all...</summary>

* _2026-05-15 04:52:36 &#43;0000 UTC_: <code>04:56:39: I0515 00:56:39.079139    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-network/2055146025328316416#1:build-log.txt%3A753)

* _2026-05-13 16:43:43 &#43;0000 UTC_: <code>16:49:17: I0513 12:49:17.114198    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054569150063316992#1:build-log.txt%3A783)

* _2026-05-13 15:06:33 &#43;0000 UTC_: <code>15:12:22: I0513 11:12:22.171698    1602 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-network/2054578675570970624#1:build-log.txt%3A1887)

</details>

<hr/>
</details>

#### needs-investigation (5x / 62.50%)

<details>
<summary> no error snippets (5x / 62.50%) </summary>

<hr/>

**1x**: _2026-05-13 13:19:37 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054551663053639680)
<details>
<summary>all...</summary>

* _2026-05-13 13:19:37 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054551663053639680)

</details>

<hr/>

**1x**: _2026-05-13 09:33:56 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054495210582315008)
<details>
<summary>all...</summary>

* _2026-05-13 09:33:56 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054495210582315008)

</details>

<hr/>

**1x**: _2026-05-13 09:01:51 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-network/2054487143505465344)
<details>
<summary>all...</summary>

* _2026-05-13 09:01:51 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17539/pull-kubevirt-e2e-k8s-1.35-sig-network/2054487143505465344)

</details>

<hr/>

**1x**: _2026-05-13 06:33:53 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17680/pull-kubevirt-e2e-k8s-1.34-sig-network/2054449906545856512)
<details>
<summary>all...</summary>

* _2026-05-13 06:33:53 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17680/pull-kubevirt-e2e-k8s-1.34-sig-network/2054449906545856512)

</details>

<hr/>

**1x**: _2026-05-13 06:33:53 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17680/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054449906466164736)
<details>
<summary>all...</summary>

* _2026-05-13 06:33:53 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17680/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2054449906466164736)

</details>

<hr/>
</details>

Last updated: 2026-05-19 16:57:15
