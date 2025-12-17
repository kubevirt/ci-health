# CI Failure Summary

This summary provides an overview of recent CI job failures that did not reach the testing stage.

## Non-Fixable Failures

### High Volume: `make cluster-sync` Error 125

**Quantity:** 3

**Description:** These jobs failed with a generic `Error 125` from `make cluster-sync`, followed by a shell error `wait: pid ... is not a child of this shell`. This suggests that a critical child process was crashing or being killed unexpectedly. This error is most likely caused by the termination of the job itself by the Prow system.

<details>
<summary>Log Snippet</summary>

```
make: *** [Makefile:174: cluster-sync] Error 125
+ ret=2
+ check_for_panics
+ set +x
+ make cluster-down
./kubevirtci/cluster-up/down.sh
Error response from daemon: volume ... is being used by the following container(s): ...
make: *** [Makefile:162: cluster-down] Error 1
...
/usr/local/bin/runner.sh: line 50: wait: pid ... is not a child of this shell
```

</details>

<details>
<summary>Failed Job URLs</summary>

* https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16343/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations/2000900961081495552
* https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16340/pull-kubevirt-e2e-k8s-1.34-ipv6-sig-network/2000504107663626240
* https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16174/pull-kubevirt-e2e-k8s-1.32-sig-operator/1998670897703030784

</details>

<details>
<summary>Pull Request History</summary>

* https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16343
* https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16340
* https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16174

</details>

| SIGs | Branches | Kubernetes Versions |
|---|---|---|
| compute, network, operator | main | 1.34, 1.32 |

---

### External Dependency Failures

**Quantity:** 1

**Description:** The build process failed because it was unable to download several RPM packages from `http://mirror.stream.centos.org/9-stream/...`, resulting in `404 Not Found` errors. This was followed by a timeout when pulling container images from `quay.io`.

<details>
<summary>Log Snippet</summary>

```
17:39:28: WARNING: Download from http://mirror.stream.centos.org/9-stream/BaseOS/x86_64/os/Packages/libsss_nss_idmap-2.9.6-2.el9.x86_64.rpm failed: class java.io.FileNotFoundException GET returned 404 Not Found
...
17:49:40: ERROR: An error occurred during the fetch of repository 'fedora_realtime':
17:49:40:    Traceback (most recent call last):
17:49:40:       File "/tmp/cache/bazel/18316b1300bb8985bc913139d5cc6323/external/io_bazel_rules_docker/container/pull.bzl", line 189, column 13, in _impl
17:49:40:               fail("Pull command failed: %s (%s)" % (result.stderr, " ".join([str(a) for a in args])))
17:49:40: Error in fail: Pull command failed: Timed out (... -name quay.io/kubevirt/fedora-realtime-container-disk@sha256:437f4e02986daf0058239f4a282d32304dcac629d5d1b4c75a74025f1ce22811)
```

</details>

<details>
<summary>Failed Job URL</summary>

* https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15990/pull-kubevirt-e2e-kind-1.30-vgpu-1.5/1983948169213382656

</details>

<details>
<summary>Pull Request History</summary>

* https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15990

</details>

| SIGs | Branches | Kubernetes Versions |
|---|---|---|
| N/A | release-1.5 | 1.30 |

---

### QEMU/ALSA Audio Device Failure

**Quantity:** 1

**Description:** The QEMU process failed to initialize audio devices, spamming the log with "ALSA: Couldn't open audio device: Host is down". This appears to be a QEMU or host environment issue.

<details>
<summary>Log Snippet</summary>

```
sdl: SDL_OpenAudioDevice for playback failed
sdl: Reason: ALSA: Couldn't open audio device: Host is down
audio: Could not create a backend for voice `dac'
...
/usr/local/bin/runner.sh: line 50: wait: pid 32 is not a child of this shell
```

</details>

<details>
<summary>Failed Job URL</summary>

* https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16357/pull-kubevirt-e2e-k8s-1.32-ipv6-sig-network-1.5/2000922274189807616

</details>

<details>
<summary>Pull Request History</summary>

* https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16357

</details>

| SIGs | Branches | Kubernetes Versions |
|---|---|---|
| network | release-1.5 | 1.32 |