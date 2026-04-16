---
name: analyze-build
description: >
    Analyze the root cause inside a failed prowjob.
    Use when user mentions to analyze a prowjob failure or adds a prowjob url
allowed-tools: [Read, Write, Bash(go run:*)]
---

## Overview

This skill helps the user to find the root cause and mitigations for failing prowjobs.

## Analysis data generation

As a base for analyzing the failure, the go tool `cmd/ci-failures` from this project is to be used.

Example how to generate the data file for a specific prowjob, given is the url to the prowjob build:

```bash
$ go run ./cmd/ci-failures analyze-build https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16885/pull-kubevirt-e2e-k8s-1.34-windows2016/2034979182789791744
{"level":"info","msg":"wrote analysis to output/tmp/pull-kubevirt-e2e-k8s-1.34-windows2016.yaml","time":"2026-03-30T13:37:48+02:00"}
```

Note: the output shows the path to the generated file containing the base data for the analysis. 

## Analysis

The output is to be looked at and the reason and possible mitigations are to be deduced from this.
