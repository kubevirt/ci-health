package stats

import (
	"github.com/fgimenez/cihealth/pkg/gh"
)

func Run(tokenPath, source string, dataDays int) (*Results, error) {
	results := &Results{
		DataDays: dataDays,
		Source:   source,
	}

	_, err := gh.GetClient(tokenPath)
	if err != nil {
		return nil, err
	}

	return results, nil
}
