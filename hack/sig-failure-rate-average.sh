#!/bin/bash
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

DAYS=${1:-90}
INTERVAL_DAYS=7

if ! [[ "$DAYS" =~ ^[0-9]+$ ]]; then
    echo "Error: Number of days must be a positive integer"
    echo "Usage: $0 [days]"
    exit 1
fi

echo "Collecting SIG failure rate data over $DAYS days with $INTERVAL_DAYS day intervals..." >&2

TEMP_FILE=$(mktemp)
trap "rm -f $TEMP_FILE" EXIT

# Create CSV header
echo "Date,SIG,Retests,Total,FailureRate" > "$TEMP_FILE"

NUM_SAMPLES=$((DAYS / INTERVAL_DAYS))

# Get list of commits to sample from
COMMITS=$(git log --oneline --format="%h %cd %s" --date=short --since="${DAYS} days ago" | grep -i "badges updated" | awk '{print $1 "," $2}')

if [ -z "$COMMITS" ]; then
    echo "Error: No 'Badges updated' commits found in the last $DAYS days" >&2
    echo "Available commits:" >&2
    git log --oneline --since="${DAYS} days ago" | head -5 >&2
    exit 1
fi

# Sample commits at intervals
COMMIT_ARRAY=()
while IFS= read -r line; do
    COMMIT_ARRAY+=("$line")
done <<< "$COMMITS"

TOTAL_COMMITS=${#COMMIT_ARRAY[@]}
echo "Found $TOTAL_COMMITS badge update commits" >&2

for ((i = 0; i < NUM_SAMPLES; i++)); do
    INDEX=$((i * INTERVAL_DAYS * TOTAL_COMMITS / DAYS))
    if [ "$INDEX" -ge "$TOTAL_COMMITS" ]; then
        INDEX=$((TOTAL_COMMITS - 1))
    fi

    COMMIT_INFO=${COMMIT_ARRAY[$INDEX]}
    COMMIT_HASH=$(echo "$COMMIT_INFO" | cut -d',' -f1)
    COMMIT_DATE=$(echo "$COMMIT_INFO" | cut -d',' -f2)

    echo "Processing commit $COMMIT_HASH from $COMMIT_DATE..." >&2

    # Extract SIG data from this commit
    RESULTS_JSON=$(git show "$COMMIT_HASH:output/kubevirt/kubevirt/results.json" 2>/dev/null || echo "{}")

    if [ "$RESULTS_JSON" = "{}" ]; then
        echo "Warning: Could not find results.json in commit $COMMIT_HASH, skipping..." >&2
        continue
    fi

    # Extract SIG retest data
    SIG_COMPUTE_RETESTS=$(echo "$RESULTS_JSON" | jq -r '.Data.SIGRetests.SIGComputeRetest // 0')
    SIG_COMPUTE_TOTAL=$(echo "$RESULTS_JSON" | jq -r '.Data.SIGRetests.SIGComputeTotal // 0')

    SIG_STORAGE_RETESTS=$(echo "$RESULTS_JSON" | jq -r '.Data.SIGRetests.SIGStorageRetest // 0')
    SIG_STORAGE_TOTAL=$(echo "$RESULTS_JSON" | jq -r '.Data.SIGRetests.SIGStorageTotal // 0')

    SIG_NETWORK_RETESTS=$(echo "$RESULTS_JSON" | jq -r '.Data.SIGRetests.SIGNetworkRetest // 0')
    SIG_NETWORK_TOTAL=$(echo "$RESULTS_JSON" | jq -r '.Data.SIGRetests.SIGNetworkTotal // 0')

    SIG_OPERATOR_RETESTS=$(echo "$RESULTS_JSON" | jq -r '.Data.SIGRetests.SIGOperatorRetest // 0')
    SIG_OPERATOR_TOTAL=$(echo "$RESULTS_JSON" | jq -r '.Data.SIGRetests.SIGOperatorTotal // 0')

    SIG_CI_RETESTS=$(echo "$RESULTS_JSON" | jq -r '.Data.SIGRetests.SIGCIRetest // 0')
    SIG_CI_TOTAL=$(echo "$RESULTS_JSON" | jq -r '.Data.SIGRetests.SIGCITotal // 0')

    SIG_MONITORING_RETESTS=$(echo "$RESULTS_JSON" | jq -r '.Data.SIGRetests.SIGMonitoringRetest // 0')
    SIG_MONITORING_TOTAL=$(echo "$RESULTS_JSON" | jq -r '.Data.SIGRetests.SIGMonitoringTotal // 0')

    # Calculate failure rates
    for SIG in "Compute" "Operator" "Storage" "Network" "CI" "Monitoring"; do
        case $SIG in
            "Compute")
                RETESTS=$SIG_COMPUTE_RETESTS
                TOTAL=$SIG_COMPUTE_TOTAL
                ;;
            "Storage")
                RETESTS=$SIG_STORAGE_RETESTS
                TOTAL=$SIG_STORAGE_TOTAL
                ;;
            "Network")
                RETESTS=$SIG_NETWORK_RETESTS
                TOTAL=$SIG_NETWORK_TOTAL
                ;;
            "Operator")
                RETESTS=$SIG_OPERATOR_RETESTS
                TOTAL=$SIG_OPERATOR_TOTAL
                ;;
            "CI")
                RETESTS=$SIG_CI_RETESTS
                TOTAL=$SIG_CI_TOTAL
                ;;
            "Monitoring")
                RETESTS=$SIG_MONITORING_RETESTS
                TOTAL=$SIG_MONITORING_TOTAL
                ;;
        esac

        if [ "$TOTAL" -gt 0 ]; then
            FAILURE_RATE=$(echo "scale=4; $RETESTS / $TOTAL" | bc)
        else
            FAILURE_RATE="0.0000"
        fi

        echo "$COMMIT_DATE,SIG-$SIG,$RETESTS,$TOTAL,$FAILURE_RATE" >> "$TEMP_FILE"
    done
done


echo ""
echo "Average Failure Rates by SIG over $DAYS days:"
echo "============================================="

for SIG in "Compute" "Operator" "Storage" "Network" "CI" "Monitoring"; do
    AVG_RATE=$(awk -F',' -v sig="SIG-$SIG" '$2 == sig { sum += $5; count++ } END { if (count > 0) print sum/count; else print 0 }' "$TEMP_FILE")
    TOTAL_RETESTS=$(awk -F',' -v sig="SIG-$SIG" '$2 == sig { sum += $3 } END { print sum+0 }' "$TEMP_FILE")
    TOTAL_TESTS=$(awk -F',' -v sig="SIG-$SIG" '$2 == sig { sum += $4 } END { print sum+0 }' "$TEMP_FILE")

    printf "%-12s: %.4f%% (%d retests / %d total tests)\n" "$SIG" "$(echo "$AVG_RATE * 100" | bc)" "$TOTAL_RETESTS" "$TOTAL_TESTS"
done
