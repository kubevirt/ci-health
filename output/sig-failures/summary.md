# KubeVirt SIG Failure Report Summary

This report summarizes the recent failures for each KubeVirt Special Interest Group (SIG).

## SIG Compute

*   **VMI, Lifecycle, Timeout:** Several tests related to the VirtualMachineInstance lifecycle, such as creation and pausing, timed out waiting for the VMI to enter the "Running" phase.
*   **VMI, virt-handler, Restart:** A test for creating a VMI when virt-handler is responsive failed after a kubelet restart with bridge networking, expecting the PodPhase to be "Succeeded" but it was "Pending".
*   **CloudInit, UserData, Timeout:** A CloudInit UserData test for injecting an ssh-key timed out.
*   **VMI, Deletion, Timeout:** A test for deleting a VMI with a grace period greater than 0 timed out, as the VMI object still existed.
*   **VMI, Softreboot, Timeout:** Two tests for soft rebooting a VMI with the ACPI feature enabled timed out waiting for a reboot record.
*   **Migration, Target, Failure:** A VMI live migration test failed because the target pod failed during target preparation.
*   **Migration, Evacuation, Backoff:** A VMI live migration triggered by evacuation failed, with the migration stuck in the "Pending" phase.
*   **Migration, Datavolume, Failure:** A live migration across namespaces with a datavolume disk failed unexpectedly.
*   **Migration, Paused, Timeout:** A VMI live migration with a paused VMI timed out, with the migration stuck in the "Running" phase.
*   **Configuration, GuestAgent, Condition:** A test for VMI configuration with guestAgent failed because it found an unexpected 'AgentVersionNotSupported' condition.
*   **Guest, Access, Credentials:** A test for guest access credentials with the qemu guest agent timed out, as the 'AgentConnected' condition was not found.
*   **SEV, Lifecycle, Failure:** An AMD SEV test for starting a SEV or SEV-ES VM timed out and the VMI unexpectedly stopped with a "Failed" state.

[SIG Compute Failure Report](https://storage.googleapis.com/kubevirt-prow/reports/sig-failure-reports/sig-compute-failure-report.html)

## SIG Storage

*   **Hotplug, Filesystem, Timeout:** A test for hotplugging both a filesystem and a block volume to an online VM timed out, expecting the VirtualMachine Status.Ready to be true.
*   **Migration, Policy, Cancellation:** Two tests for cancelling a live migration across namespaces by deleting the migration resource timed out.
*   **Hotplug, DataVolume, Timeout:** A test for adding/removing a volume with DataVolume immediate attach (virtio) to an online VM timed out.
*   **Hotplug, Migration, Timeout:** A VMI migration test with attached hotplug persistent disk volumes timed out.

[SIG Storage Failure Report](https://storage.googleapis.com/kubevirt-prow/reports/sig-failure-reports/sig-storage-failure-report.html)

## SIG Network

*   **Masquerade, Migration, Connectivity:** A test for a VMI with masquerade binding mechanism failed to preserve connectivity after migration because it "failed to run dhcp client".
*   **SRIOV, Hotplug, Timeout:** Multiple tests for SRIOV VMI connected to a single SRIOV network with memory hotplug timed out.
*   **Subdomain, Headless, FQDN:** A test for a VMI with a subdomain and a headless service timed out.
*   **Bridge, Hotplug, Migration:** A test for hotplugging multiple network interfaces to a running VM with migration timed out.
*   **Bridge, Hotunplug, Migration:** Two tests for hot-unplugging a network interface from a running VM with migration failed due to an "admission webhook 'migration-create-validator.kubevirt.io' denied the request: in-flight migration detected".

[SIG Network Failure Report](https://storage.googleapis.com/kubevirt-prow/reports/sig-failure-reports/sig-network-failure-report.html)

## SIG Operator

*   **Operator, Reconcile, Daemonsets:** A test for the operator reconciling components and reverting changes for daemonsets timed out.
*   **Operator, Update, KubeVirt:** A test for the operator updating KubeVirt from a previous release to the target tested release by patching the KubeVirt CR timed out.

[SIG Operator Failure Report](https://storage.googleapis.com/kubevirt-prow/reports/sig-failure-reports/sig-operator-failure-report.html)

## SIG Monitoring

*   **No, SIG, Failures:** The SIG Monitoring failure report shows no SIG failures to display.

[SIG Monitoring Failure Report](https://storage.googleapis.com/kubevirt-prow/reports/sig-failure-reports/sig-monitoring-failure-report.html)