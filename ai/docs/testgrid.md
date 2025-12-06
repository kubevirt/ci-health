# testgrid

TestGrid is a highly-configurable, interactive dashboard for viewing your test results in a grid!

More information in the github repository: https://github.com/GoogleCloudPlatform/testgrid

## mapping prow job names to testgrid URLs

There's several dashboards for kubevirt prow jobs. For example there's the dashboard URL
"https://testgrid.k8s.io/kubevirt-presubmits" for all presubmit jobs for the kubevirt/kubevirt repository.
Also there's the dashboard URL "https://testgrid.k8s.io/kubevirt-periodics" for all periodic jobs for
the kubevirt/kubevirt repository.

Mapping the job name to a testgrid URL is to use the pattern "https://testgrid.k8s.io/kubevirt-{job-type}#{job-name}" and then replacing "{job-type}" either with "periodics" for periodic jobs or with "presubmits" for presubmit (also known als "pull" jobs), and then replacing 
"{job-name}" with the job name.

### Mapping examples

* mapping the presubmit job named "pull-kubevirt-e2e-k8s-1.32-sig-storage" to a testgrid URL will result in the URL "https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.32-sig-storage"
* mapping the periodic job named "periodic-kubevirt-e2e-k8s-1.33-sig-compute" to a testgrid URL will result in the URL "https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.33-sig-compute"
