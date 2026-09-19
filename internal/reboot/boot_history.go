package reboot

import (
	"linistic/internal/utils"
	"strings"
)

type BootHistory struct {
	PreviousBootFound bool
	CleanShutdown     bool
	RebootDetected    bool
}

func GetBootHistory() BootHistory {

	logs := utils.Run(
		"journalctl",
		"-b",
		"-1",
		"--no-pager",
	)

	state := BootHistory{}

	if strings.TrimSpace(logs) == "" {
		return state
	}

	state.PreviousBootFound = true

	lower := strings.ToLower(logs)

	cleanShutdownIndicators := []string{
		"reached target shutdown.target",
		"reached target shutdown",
		"system shutdown",
		"systemd-shutdown",
		"reached target final.target",
	}

	for _, indicator := range cleanShutdownIndicators {
		if strings.Contains(lower, indicator) {
			state.CleanShutdown = true
			break
		}
	}

	rebootIndicators := []string{
		"system is rebooting",
		"reached target reboot.target",
		"system reboot",
		"systemd-reboot.service",
		"the system will reboot now",
	}

	for _, indicator := range rebootIndicators {
		if strings.Contains(lower, indicator) {
			state.RebootDetected = true
			break
		}
	}

	return state
}
