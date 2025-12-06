# SIG Failure Summary

This document provides a summary of the SIG failure reports.

## SIG Compute

*   **VMI, Lifecycle, Timeout** - `VMIlifecycle [rfe_id:1287][crit:high][vendor:cnv-qe@redhat.com][level:component]Creating a VirtualMachineInstance definition creation should be accepted [test_id:1608]with a restartPolicy of Always` failed due to a timeout waiting for the VMI to be in the `Running` phase.
*   **VMI, Lifecycle, Timeout** - `VMIlifecycle [rfe_id:1287][crit:high][vendor:cnv-qe@redhat.com][level:component]Creating a VirtualMachineInstance definition creation should be accepted [test_id:1609]with a restartPolicy of Never` failed due to a timeout waiting for the VMI to be in the `Running` phase.
*   **VMI, Lifecycle, Timeout** - `VMIlifecycle [rfe_id:1287][crit:high][vendor:cnv-qe@redhat.com][level:component]Creating a VirtualMachineInstance and deleting it should be resilient to virt-handler restart [test_id:1612]by executing a graceful shutdown of the VirtualMachineInstance` failed due to a timeout.
*   **VMI, virt-handler, Pod** - `[sig-network] [crit:high][vendor:cnv-qe@redhat.com][level:component] [crit:high][vendor:cnv-qe@redhat.com][level:component]Creating a VirtualMachineInstance when virt-handler is responsive VMIs shouldn't fail after the kubelet restarts [sig-network]with bridge networking` failed because a pod remained in `Pending` phase.
*   **CloudInit, UserData, Timeout** - `[rfe_id:151][crit:high][vendor:cnv-qe@redhat.com][level:component][sig-compute]CloudInit UserData [rfe_id:151][crit:medium][vendor:cnv-qe@redhat.com][level:component]A new VirtualMachineInstance with cloudInitNoCloud with injected ssh-key [test_id:1616]should have ssh-key under authorized keys` failed due to a timeout.
*   **Migration, Policy, Timeout** - `VM Live Migration Starting a VirtualMachineInstance migration monitor [test_id:6980]Migration should fail if target pod fails during target preparation` failed due to a timeout.
*   **Migration, Namespace, DataVolume** - `Live Migration across namespaces datavolume disk should live migrate regular disk several times` failed due to an unexpected migration phase.
*   **Configuration, GuestAgent, Condition** - `[sig-compute]Configurations VirtualMachineInstance definition [rfe_id:140][crit:medium][vendor:cnv-qe@redhat.com][level:component]with guestAgent with cluster config changes [test_id:6958]VMI condition should not signal unsupported agent presence for optional commands` failed due to an unexpected `AgentVersionNotSupported` condition.
*   **GuestAccess, Credentials, Condition** - `[sig-compute] Guest Access Credentials with qemu guest agent should update to unsupported agent [test_id:6222]for public ssh keys` failed because the `AgentConnected` condition was not found.
*   **SEV, Lifecycle, Timeout** - `[sig-compute]AMD Secure Encrypted Virtualization (SEV) lifecycle should start a SEV or SEV-ES VM It should launch with base SEV features enabled` failed due to a timeout and the VMI unexpectedly stopping.

[Link to full report](https://storage.googleapis.com/kubevirt-prow/reports/sig-failure-reports/sig-compute-failure-report.html)

## SIG Storage

*   **Hotplug, Filesystem, BlockVolume** - `[sig-storage] Hotplug [storage-req] Online VM should allow hotplugging both a filesystem and block volume` timed out waiting for the VirtualMachine to be ready.
*   **Migration, Namespace, Policy** - `[rfe_id:393][crit:high][vendor:cnv-qe@redhat.com][level:system][sig-compute] Live Migration across namespaces with migration policy should be able to cancel a migration by deleting the migration resource delete target migration` timed out.
*   **Migration, Namespace, Policy** - `[rfe_id:393][crit:high][vendor:cnv-qe@redhat.com][level:system][sig-compute] Live Migration across namespaces with migration policy should be able to cancel a migration by deleting the migration resource delete source migration` timed out.
*   **Hotplug, DataVolume, Timeout** - `[sig-storage] Hotplug [storage-req] Online VM should add/remove volume with DataVolume immediate attach (virtio)` timed out.
*   **Hotplug, Migration, PersistentVolume** - `[sig-storage] Hotplug [storage-req] VMI migration should allow live migration with attached hotplug volumes persistent disk VMI` timed out.

[Link to full report](https://storage.googleapis.com/kubevirt-prow/reports/sig-failure-reports/sig-storage-failure-report.html)

## SIG Network

*   **Masquerade, Migration, IPv4** - `[sig-network] [rfe_id:694][crit:medium][vendor:cnv-qe@redhat.com][level:component]Networking VirtualMachineInstance with masquerade binding mechanism when performing migration preserves connectivity - IPv4 with explicit ports used by live migration` failed because the dhcp client failed to run.
*   **SRIOV, Hotplug, Memory** - `SRIOV VMI connected to single SRIOV network memory hotplug Should successfully reattach host-device` failed due to timeouts and unexpected network interface configurations.
*   **Subdomain, FQDN, Masquerade** - `[sig-network] Subdomain with a headless service given VMI should have the expected FQDN with Masquerade binding without subdomain` timed out.
*   **Bridge, Hotplug, Migration** - `[sig-network] bridge nic-hotplug a running VM is able to hotplug multiple network interfaces Migration based` failed due to missing network interfaces after a timeout.
*   **Bridge, Hotunplug, Migration** - `[sig-network] bridge nic-hotunplug a running VM hot-unplug network interface succeed Migration based` failed due to an "admission webhook denied the request: in-flight migration detected" error.

[Link to full report](https://storage.googleapis.com/kubevirt-prow/reports/sig-failure-reports/sig-network-failure-report.html)

## SIG Operator

*   **Operator, Reconcile, DaemonSet** - `[sig-operator]Operator should reconcile components checking updating resource is reverted to original state for [test_id:6308] daemonsets` failed due to a timeout.
*   **Operator, Update, KubeVirtCR** - `[sig-operator]Operator [rfe_id:2291][crit:high][vendor:cnv-qe@redhat.com][level:component]should update kubevirt [release-blocker][test_id:3145]from previous release to target tested release by patching KubeVirt CR` timed out.

[Link to full report](https://storage.googleapis.com/kubevirt-prow/reports/sig-failure-reports/sig-operator-failure-report.html)

## SIG Monitoring

No failures to report.

[Link to full report](https://storage.googleapis.com/kubevirt-prow/reports/sig-failure-reports/sig-monitoring-failure-report.html)
