# About this document

This document describes the different branch types and how to identify target branches of prow jobs.

# KubeVirt repository - branching strategy

## Goal

Ensure that the release branches remain stable while development of new features continues on the `main` branch.

## Main Branch

The `main` branch is the primary development branch. All new features and bug fixes should be based on this branch. Pull requests from contributors are targeted to the `main` branch.

The `main` branch is expected to be in a runnable state at all times, with all tests passing.

## Release Branches

When a new release is planned, a release branch is created from the `main` branch. These branches are named using the convention `release-X.Y`, where `X` is the major version and `Y` is the minor version. For example, `release-0.58`.

Once a release branch is created, it is considered "frozen." Only critical bug fixes are cherry-picked from the `main` branch into the release branch. No new features are added to a release branch.

