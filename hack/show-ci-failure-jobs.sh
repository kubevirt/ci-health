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
usage: $0

    fetches URLs for the CI failure runs from the data of the latest run -
    covering last 7 days.
EOF
}

if [ "$#" -gt 0 ]; then
    if [[ "$1" == -h ]] || [[ "$1" == --help ]]; then
        usage
        exit 0
    else
        usage
        exit 1
    fi
fi

for url in $(jq -r '.Data.SIGRetests.FailedJobLeaderBoard[].FailureURLs | flatten[]' ./output/kubevirt/kubevirt/results.json); do
    if ! curl -s --head --fail "$(echo "$url" | sed -re 's#^https://prow.ci.kubevirt.io//view/gs/(.*)#https://storage.googleapis.com/\1/artifacts/junit.functest.xml#g')" -o /dev/null; then
        echo "$url"
    fi
done
