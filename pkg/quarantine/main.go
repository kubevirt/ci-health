package quarantine

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/avast/retry-go"
	log "github.com/sirupsen/logrus"
)

const org = "kubevirt"
const repo = "kubevirt"

type report struct {
	FilesStats []fileStat `json:"files_stats"`
}

type fileStat struct {
	Path      string    `json:"path"`
	TestStats testStats `json:"test_stats"`
}

type testStats struct {
	SpecsTotal        int         `json:"specs_total"`
	MatchingSpecPaths interface{} `json:"matching_spec_paths"`
}

type QuarantineStats struct {
	TotalQuarantineCount    float64
	SigComputeQuarantine    float64
	SigStorageQuarantine    float64
	SigNetworkQuarantine    float64
	SigMonitoringQuarantine float64
}

func quarantineReportURL(org string, repo string) string {
	return fmt.Sprintf("https://storage.googleapis.com/kubevirt-prow/reports/quarantined-tests/%s/%s/quarantine.json", org, repo)
}

func httpGetWithRetry(url string) (resp *http.Response, err error) {
	httpRetryLog := log.WithField("url", url)
	retry.Do(
		func() error {
			resp, err = http.Get(url)
			switch {
			case resp.StatusCode == http.StatusOK:
				httpRetryLog.Debugf("http get succeeded")
				return nil
			case resp.StatusCode == http.StatusGatewayTimeout:
				httpRetryLog.Debugf("failed to http get, will retry")
				return fmt.Errorf("failed to http get %s (status %d): %v", url, resp.StatusCode, err)
			case err != nil:
				httpRetryLog.Debugf("failed to http get, aborting")
				return retry.Unrecoverable(err)
			default:
				httpRetryLog.Debugf("failed to http get, aborting")
				return retry.Unrecoverable(fmt.Errorf("failed to http get %s (status %d): %v", url, resp.StatusCode, err))
			}
		},
	)
	return
}

func GetQuarantineStats() (stats QuarantineStats, err error) {
	var (
		quarantineReport report
		qStats           QuarantineStats
	)

	quarantineURL := quarantineReportURL(org, repo)
	resp, err := httpGetWithRetry(quarantineURL)
	if err != nil {
		log.Errorf("Failed to download quarantine report")
		return qStats, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&quarantineReport)
	if err != nil {
		log.Errorf("Failed to decode JSON quarantine report")
		return qStats, err
	}
	for _, f := range quarantineReport.FilesStats {
		if f.TestStats.MatchingSpecPaths != nil {
			paths, ok := f.TestStats.MatchingSpecPaths.([]interface{})
			if ok {
				for _, path := range paths {
					pathMap, ok := path.(map[string]interface{})
					if ok {
						if _, found := pathMap["lines"].([]interface{}); found {
							qStats.TotalQuarantineCount += 1
							switch {
							case strings.Contains(f.Path, "storage"):
								qStats.SigStorageQuarantine += 1
							case strings.Contains(f.Path, "network"):
								qStats.SigNetworkQuarantine += 1
							case strings.Contains(f.Path, "monitoring"):
								qStats.SigMonitoringQuarantine += 1
							default:
								qStats.SigComputeQuarantine += 1
							}
						}
					}
				}
			}
		}
	}
	return qStats, nil
}
