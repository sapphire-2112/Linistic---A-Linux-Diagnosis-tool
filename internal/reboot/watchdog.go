package reboot

import (
	"strings"

	"linistic/internal/utils"
)

type WatchdogState struct {
	WatchdogDetected bool
	HardLockup       bool
	SoftLockup       bool

	Evidence []string
}

func GetWatchdogState() WatchdogState {

	logs := strings.ToLower(
		utils.Run(
			"journalctl",
			"-b",
			"-1",
			"--no-pager",
		),
	)

	state := WatchdogState{}

	hardLockupIndicators := []string{
		"hard lockup",
		"watchdog detected hard lockup",
		"nmi watchdog",
	}

	for _, indicator := range hardLockupIndicators {
		if strings.Contains(logs, indicator) {
			state.HardLockup = true
			state.Evidence = append(state.Evidence, indicator)
		}
	}

	softLockupIndicators := []string{
		"soft lockup",
		"bug: soft lockup",
	}

	for _, indicator := range softLockupIndicators {
		if strings.Contains(logs, indicator) {
			state.SoftLockup = true
			state.Evidence = append(state.Evidence, indicator)
		}
	}

	watchdogIndicators := []string{
		"watchdog timeout",
		"watchdog bite",
		"watchdog reset",
	}

	for _, indicator := range watchdogIndicators {
		if strings.Contains(logs, indicator) {
			state.WatchdogDetected = true
			state.Evidence = append(state.Evidence, indicator)
		}
	}

	return state
}