# CI Failure Analysis Summary

## Non-fixable errors

### Unexpected Build Process Termination

This error is characterized by the build process terminating unexpectedly, often resulting in a `make: *** [Makefile:174: cluster-sync] Error 125` or similar error, followed by a failure to clean up the test environment. This is likely caused by the Prow job being terminated due to resource constraints or other infrastructure issues.

**Quantity total:** 4

**Branch name(s):** release-1.7, main

**SIG(s):** compute, operator, windows, performance, storage

**Kubernetes version(s):** 1.34, 1.32

<details>
<summary>Example log snippet</summary>

```
make: *** [Makefile:174: cluster-sync] Error 125
+ ret=2
+ make cluster-down
./kubevirtci/cluster-up/down.sh
13:40:32: selecting podman as container runtime
13:41:19: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.34-windows2016-1.7 is being used by the following container(s): 5c9c43d4c824671cc2ff9f901d37a0abe42b0f22baad949aa97b94694cb288fe, 1c3c3aeadc6100ee41fc51a5ce82f5df514d206fbbe71d310cca5bd67cc3660c: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
+ true
+ exit 2
+ EXIT_VALUE=2
+ set +o xtrace
Cleaning up after podman in container.
================================================================================
Cleaning up after podman
time="2025-12-29T13:41:29Z" level=warning msg="StopSignal SIGTERM failed to stop container pull-kubevirt-e2e-k8s-1.34-windows2016-1.7-bazel-server in 10 seconds, resorting to SIGKILL"
1c3c3aeadc6100ee41fc51a5ce82f5df514d206fbbe71d310cca5bd67cc3660c
5c9c43d4c824671cc2ff9f901d37a0abe42b0f22baad949aa97b94694cb288fe
/usr/local/bin/runner.sh: line 50: wait: pid 1206 is not a child of this shell
================================================================================
Done cleaning up after podman in container.
```

</details>

<details>
<summary>Failed build job URLs</summary>

* https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16424/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2005626330367922176
* https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16424/pull-kubevirt-e2e-k8s-1.34-sig-compute-migrations-1.7/2005626338702004224
* https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/15975/pull-kubevirt-e2e-k8s-1.34-sig-performance/2005630385106456576
* https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16403/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2003739616691097600

</details>

<details>
<summary>Pull request history URLs</summary>

* https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16424
* https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=15975
* https://prow.ci.kubevirt.io/pr-history/?org=kubevirt&repo=kubevirt&pr=16403

</details>
