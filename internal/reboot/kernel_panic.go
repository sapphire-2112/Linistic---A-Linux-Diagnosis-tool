package reboot

import (
	"strings"

	"linistic/internal/utils"
)

type KernelPanicState struct {
	PanicDetected bool
	PanicReason   string
}

func GetKernelPanicState() KernelPanicState {

	logs := strings.ToLower(
		utils.Run(
			"journalctl",
			"-k",
			"-b",
			"-1",
			"--no-pager",
		),
	)

	state := KernelPanicState{}

	indicators := []string{
		"kernel panic",
		"not syncing",
		"oops:",
		"bug:",
		"general protection fault",
		"unable to handle kernel",
		"fatal exception",
		"kernel bug",
		"panic occurred",
	}

	for _, indicator := range indicators {

		if strings.Contains(logs, indicator) {

			state.PanicDetected = true
			state.PanicReason = indicator

			break
		}
	}

	return state
}