#!/usr/bin/env bash
#
# This file is part of the KubeVirt project
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# Copyright the KubeVirt Authors.
#
#

set -euo pipefail

function usage() {
    cat <<EOF
usage: $0 [{file_with_job_urls} default 'output/tmp/ci-failure-jobs.txt']

    fetches failures from the build logs of build urls given through the file - each line correlates to a build url.
    writes matching lines into one file per group, prefixed with 'output/tmp/errors-'
    outputs the file names of the files created
EOF
}

if [ "$#" -gt 0 ]; then
    if [[ "$1" == -h ]] || [[ "$1" == --help ]]; then
        usage
        exit 0
    fi

    if [ ! -f "$1" ]; then
        usage
        exit 1
    fi
fi

urls_file="${1-output/tmp/ci-failure-jobs.txt}"

work_dir=$(mktemp -d)

function cleanup() {
    rm -rf "${work_dir}"
}

trap 'cleanup' SIGINT SIGTERM EXIT

current_commit="$(git rev-parse HEAD)"

# see https://github.com/kubevirt/project-infra/blob/4aa6433517a4a3d09ccf312adeb8fedf211dbdc8/github/ci/prow-deploy/kustom/base/configs/current/config/config.yaml#L39C1-L47C1
rg_expression='^E\d{4} \d\d:\d\d:\d\d\.\d+|(Error|ERROR|error)s?:|(FAIL|Failure \[)\b|timed out|panic\b|\[FAILED\]|fatal: '

declare -a groups=( 'sig-compute|sig-operator|sev|vgpu|windows' 'sig-network|sriov' 'sig-storage' 'sig-monitoring')
for group in "${groups[@]}"; do
    output_file="output/tmp/errors-${group//|/-}-${current_commit:0:7}.txt"
    truncate --size 0 "${output_file}"
    echo '' > "${output_file}"

    while IFS= read -r job_url; do
        build_log_file="${work_dir}/build-log.txt"
        job_name=$(echo "${job_url}" | sed -re 's#.*(pull-kubevirt-[^/]+).*#\1#g')
        curl --silent --fail \
            "$(echo "${job_url}" | sed -re 's#^https://prow.ci.kubevirt.io//view/gs/(.*)#https://storage.googleapis.com/\1#g')/build-log.txt" \
            -o "${build_log_file}"
        cat <<EOF >>"${output_file}"
prow job ${job_name} with URL ${job_url} has following errors:
-----------------------------------------
EOF
        if ! rg -B 3 -A 1 "${rg_expression}" "${build_log_file}" >>"${output_file}"; then
            echo "no matches found for ERROR in ${job_url}" >>"${output_file}"
        fi
        echo "-----------------------------------------" >>"${output_file}"
        echo "" >>"${output_file}"
    done< <(grep -E "$group" "${urls_file}")

    echo "${output_file}"
done
