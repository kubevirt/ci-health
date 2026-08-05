package cifailures

import (
	"fmt"
	"regexp"
)

var groups = map[string]string{
	"sig-compute":     "sig-compute|sig-operator|sev|vgpu|windows",
	"sig-network":     "sig-network|sriov",
	"sig-storage":     "sig-storage",
	"sig-monitoring":  "sig-monitoring",
	"sig-performance": "sig-performance",
}

const UnknownSIG = "(sig-unknown)"

func SIGForGroup(whatever string) (sig string, err error) {
	for sig, group := range groups {
		groupMatcher := regexp.MustCompile(group)
		if groupMatcher.MatchString(whatever) {
			return sig, nil
		}
	}
	return UnknownSIG, fmt.Errorf("no sig found for %s", whatever)
}
