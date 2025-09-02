# About this document

This document describes the naming conventions being used for the prow jobs in the KubeVirt organization.

# Prow Job Naming Conventions and Branching Strategy

The Prow job configurations in the project-infra repository follow a consistent naming convention that reflects the branching strategy of the repositories they test. This convention allows
for targeted testing of the main branch, release branches, and tags.

Important: jobs for main branches do not contain a version suffix, while jobs for release branches do contain a version suffix! If a job name lacks a version suffix, it targets the `main` branch.

Example:
* job name for branch `main`: `pull-kubevirt-e2e-k8s-1.25-sig-compute` 
* job name for branch `release-1.0`: `pull-kubevirt-e2e-k8s-1.25-sig-compute-1.0` , where the suffix `-1.0` indicates that the job targets the release branch of that version


## Presubmit Jobs (run on pull requests)

Presubmit jobs are defined in files that are typically named [repo]-presubmits.yaml for the main branch and [repo]-presubmits-[version].yaml for release branches.

* `main` Branch Jobs:
    * Jobs that run on the main branch have generic names, such as pull-kubevirt-unit-test.
    * These jobs use the skip_branches field to prevent them from running on release branches. For example:

```yaml
         skip_branches:
           - release-\d+\.\d+
```

* Release Branch Jobs:
    * Jobs that run on release branches have a version suffix in their name, such as pull-kubevirt-unit-test-0.59.
    * These jobs use the branches field to target a specific release branch. For example:

```yaml
         branches:
           - release-0.59
```

* To maintain a consistent set of checks on pull requests, jobs for different release branches often use the context field to report their status to the same GitHub check. For example,
pull-kubevirt-unit-test-0.59 might have the context pull-kubevirt-unit-test.

## Postsubmit Jobs (run after a merge)

Postsubmit jobs are defined in files typically named [repo]-postsubmits.yaml. They are configured to run on specific branches or tags using regular expressions in the branches field.

* `main` Branch Jobs:
    * These jobs are configured to run only on the main branch. For example:
        ```yaml
              branches:
                - main

* Release Branch Jobs:
    * These jobs use a regular expression to match release branch names, such as ^release-\d+\.\d+$. This allows a single job definition to cover all release branches.

* Tag Jobs:
    * These jobs are triggered when a new tag is pushed. They use a regular expression to match semantic versioning patterns, such as ^v\d+\.\d+\.\d+$.

## Periodic Jobs (run on a schedule)

Periodic jobs are not directly tied to pull requests or merges, but they are often configured to run against a specific branch. This is specified in the extra_refs section of the job
definition.

* Example:

```yaml
     extra_refs:
       - org: kubevirt
         repo: kubevirt
         base_ref: main
```

This configuration ensures that the periodic job will clone the kubevirt/kubevirt repository and check out the main branch before running.
