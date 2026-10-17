package stats

import (
	"sync"

	log "github.com/sirupsen/logrus"

	cifailures "github.com/kubevirt/ci-health/pkg/ci-failures"
	"github.com/kubevirt/ci-health/pkg/types"
)

var (
	loadFlakefinderFn  = cifailures.LoadFlakefinderIndex
	fetchFailedTestsFn = cifailures.FetchFailedJunitTests
)

type e2eClassifier struct {
	days int

	reportOnce sync.Once
	report     *cifailures.FlakefinderReport
}

func newE2EClassifier(dataDays int) *e2eClassifier {
	if dataDays < 1 {
		dataDays = 7
	}
	return &e2eClassifier{days: dataDays}
}

func (c *e2eClassifier) flakefinder() *cifailures.FlakefinderReport {
	c.reportOnce.Do(func() {
		report, err := loadFlakefinderFn(c.days)
		if err != nil {
			log.WithError(err).Warn("failed to load flakefinder; e2e jobs will not be labeled flake")
			return
		}
		c.report = report
	})
	return c.report
}

func (c *e2eClassifier) classify(jobURL string) (cause, reason string) {
	tests, err := fetchFailedTestsFn(jobURL)
	if err != nil {
		log.WithError(err).Warnf("failed to fetch junit tests for %s; treating as ci", jobURL)
		return types.FailureCauseCI, "failed to fetch junit; treating as ci"
	}
	return cifailures.ClassifyE2ECause(tests, c.flakefinder())
}
