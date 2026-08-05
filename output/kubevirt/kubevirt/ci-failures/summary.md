



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-08-04 (1x / 25.00%)


#### external (1x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-08-04 01:45:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18671/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2084455542964621312#1:build-log.txt%3A604)
<details>
<summary>all...</summary>

* _2026-08-04 01:45:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18671/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2084455542964621312#1:build-log.txt%3A604)

</details>

<hr/>
</details>

### 2026-07-30 (3x / 75.00%)


#### external (2x / 66.67%)

<details>
<summary> container image pull failure in context (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-07-30 06:56:30 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18592/pull-kubevirt-e2e-kind-1.35-vgpu/2082721834897248256#1:build-log.txt%3A444)
<details>
<summary>all...</summary>

* _2026-07-30 06:56:30 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18592/pull-kubevirt-e2e-kind-1.35-vgpu/2082721834897248256#1:build-log.txt%3A444)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-07-30 05:30:51 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18613/pull-kubevirt-e2e-k8s-1.36-sig-network/2082700298786181120#1:build-log.txt%3A1858)
<details>
<summary>all...</summary>

* _2026-07-30 05:30:51 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18613/pull-kubevirt-e2e-k8s-1.36-sig-network/2082700298786181120#1:build-log.txt%3A1858)

</details>

<hr/>
</details>

#### internal (1x / 33.33%)

<details>
<summary> make cluster lifecycle target failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-07-30 13:41:27 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18648/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2082821112458121216#1:build-log.txt%3A1056)
<details>
<summary>all...</summary>

* _2026-07-30 13:41:27 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18648/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2082821112458121216#1:build-log.txt%3A1056)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (3x / 75.00%)

<details>
<summary> container image pull failure in context (2x / 50.00%) </summary>

<hr/>

**2x**: _2026-08-04 01:45:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18671/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2084455542964621312#1:build-log.txt%3A604)
<details>
<summary>all...</summary>

* _2026-08-04 01:45:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18671/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2084455542964621312#1:build-log.txt%3A604)
<details><summary>context</summary>
<pre>01:54:50: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
01:54:50: Writing manifest to image destination
01:54:50: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9 is being used by the following container(s): 6a071776d5ab7e847352faa75fdfc6084d16e607f072b021e3de5c756d8e0253: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


* _2026-07-30 06:56:30 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18592/pull-kubevirt-e2e-kind-1.35-vgpu/2082721834897248256#1:build-log.txt%3A444)
<details><summary>context</summary>
<pre>07:02:40: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
07:02:40: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
07:02:43: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage2359435640/2&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/5c/5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20260730%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20260730T070243Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f175f684cac3ad3f4b7138ba6cd3fc46f17a809d59715f296ba1210765b89a14&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1785395863~hmac=d2e1138027895eb1a76396c4f3ce5e8605eb87adaa68b6e6ede1aeb2ec91e6e3&#34;: EOF)
make: *** [Makefile:174: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-07-30 05:30:51 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18613/pull-kubevirt-e2e-k8s-1.36-sig-network/2082700298786181120#1:build-log.txt%3A1858)
<details>
<summary>all...</summary>

* _2026-07-30 05:30:51 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18613/pull-kubevirt-e2e-k8s-1.36-sig-network/2082700298786181120#1:build-log.txt%3A1858)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
05:34:51: selecting podman as container runtime
05:35:25: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network is being used by the following container(s): 1a3c5054a4b326fcf6e48875c02bb6ad4340e0af4d54ac695892cf63ebac96ec: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>

### internal (1x / 25.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-07-30 13:41:27 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18648/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2082821112458121216#1:build-log.txt%3A1056)
<details>
<summary>all...</summary>

* _2026-07-30 13:41:27 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18648/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2082821112458121216#1:build-log.txt%3A1056)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
14:02:11: selecting podman as container runtime
14:02:38: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7 is being used by the following container(s): e0a020bd2bd11f60ce26e9525064ea1f4f66c9b590247ed662d259c32a9f9537, 6b2381dd92e851177e8fae6be1d2610d9fd5a53c242b4287982a4e3c899d7cc8: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### release-1.9 (1x / 25.00%)


#### external (1x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-08-04 01:45:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18671/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2084455542964621312#1:build-log.txt%3A604)
<details>
<summary>all...</summary>

* _2026-08-04 01:45:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18671/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2084455542964621312#1:build-log.txt%3A604)

</details>

<hr/>
</details>

### release-1.7 (1x / 25.00%)


#### internal (1x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-07-30 13:41:27 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18648/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2082821112458121216#1:build-log.txt%3A1056)
<details>
<summary>all...</summary>

* _2026-07-30 13:41:27 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18648/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2082821112458121216#1:build-log.txt%3A1056)

</details>

<hr/>
</details>

### main (2x / 50.00%)


#### external (2x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-07-30 06:56:30 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18592/pull-kubevirt-e2e-kind-1.35-vgpu/2082721834897248256#1:build-log.txt%3A444)
<details>
<summary>all...</summary>

* _2026-07-30 06:56:30 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18592/pull-kubevirt-e2e-kind-1.35-vgpu/2082721834897248256#1:build-log.txt%3A444)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-07-30 05:30:51 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18613/pull-kubevirt-e2e-k8s-1.36-sig-network/2082700298786181120#1:build-log.txt%3A1858)
<details>
<summary>all...</summary>

* _2026-07-30 05:30:51 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18613/pull-kubevirt-e2e-k8s-1.36-sig-network/2082700298786181120#1:build-log.txt%3A1858)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (3x / 75.00%)


#### external (2x / 66.67%)

<details>
<summary> container image pull failure in context (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-08-04 01:45:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18671/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2084455542964621312#1:build-log.txt%3A604)
<details>
<summary>all...</summary>

* _2026-08-04 01:45:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18671/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2084455542964621312#1:build-log.txt%3A604)

* _2026-07-30 06:56:30 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18592/pull-kubevirt-e2e-kind-1.35-vgpu/2082721834897248256#1:build-log.txt%3A444)

</details>

<hr/>
</details>

#### internal (1x / 33.33%)

<details>
<summary> make cluster lifecycle target failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-07-30 13:41:27 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18648/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2082821112458121216#1:build-log.txt%3A1056)
<details>
<summary>all...</summary>

* _2026-07-30 13:41:27 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18648/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2082821112458121216#1:build-log.txt%3A1056)

</details>

<hr/>
</details>

### sig-network (1x / 25.00%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-07-30 05:30:51 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18613/pull-kubevirt-e2e-k8s-1.36-sig-network/2082700298786181120#1:build-log.txt%3A1858)
<details>
<summary>all...</summary>

* _2026-07-30 05:30:51 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18613/pull-kubevirt-e2e-k8s-1.36-sig-network/2082700298786181120#1:build-log.txt%3A1858)

</details>

<hr/>
</details>

Last updated: 2026-08-05 18:40:41
