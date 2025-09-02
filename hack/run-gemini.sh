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

set -e
set -u
set -o pipefail

function usage() {
    cat <<EOF
usage: $0 [prompt-file-name]

Prompt files available:
EOF
( cd ./ai && ls ./*.md )
}

if [ "$#" -lt 1 ]; then
    usage
    exit 1
fi

prompt_file_name="./ai/$1"
if [ ! -f "${prompt_file_name}" ]; then
    echo "${prompt_file_name} does not exist"
    usage
    exit 1
fi

set -x
gemini --debug --yolo <"${prompt_file_name}"
