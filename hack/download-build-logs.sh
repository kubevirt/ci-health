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

      downloads the log files for urls of failed jobs located in output/tmp/ci-failure-jobs.txt
      into the folder output/tmp/build-logs. Each build log file can be identified by the
      file name that holds the build number being the same as the last part of the build url from the
      output/tmp/ci-failure-jobs.txt.
EOF
}

if [ "$#" -gt 0 ]; then
    if [[ "$1" == -h ]] || [[ "$1" == --help ]]; then
        usage
        exit 0
    fi
fi

if [ ! -f 'output/tmp/ci-failure-jobs.txt' ]; then
    echo "output/tmp/ci-failure-jobs.txt does not exist!"
    exit 1

fi

if [ ! -d 'output/tmp/build-logs' ]; then
    mkdir -p output/tmp/build-logs
fi

while read -r url; do
    build_id=$(echo "$url" | awk -F/ '{print $NF}')
    gcs_url=$(echo "$url" | sed -re 's#^https://prow.ci.kubevirt.io//view/gs/(.*)#https://storage.googleapis.com/\1#g')/build-log.txt
    log_file="output/tmp/build-logs/${build_id}.txt"
    cat <<EOF >"$log_file"
prow job
    urls
        job: ${url}
        build-log.txt: ${gcs_url}

-----------------------------------------------------
EOF
    curl -s "$gcs_url" >>"$log_file"
done <output/tmp/ci-failure-jobs.txt
