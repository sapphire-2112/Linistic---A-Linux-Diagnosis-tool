package networking

import (
	"strings"
	"linistic/internal/utils"
)

type DHCPState struct {
	LeaseFailure      bool
	TimeoutDetected   bool
	ExcessiveRenewals bool

	LeaseCount   int
	NoLeaseCount int
	TimeoutCount int
}

func GetDHCPState() DHCPState {

	state := DHCPState{}

	net := GetCurrentNetworkState()

	

	if !net.Connected {
		return state
	}

	logs := strings.ToLower(
		utils.Run(
			"journalctl",
			"-u",
			"NetworkManager",
			"-n",
			"200",
			"--no-pager",
		),
	)

	state.LeaseCount =
		strings.Count(logs, "state changed new lease")

	state.NoLeaseCount =
		strings.Count(logs, "state changed no lease")

	state.TimeoutCount =
		strings.Count(logs, "request timed out")

	if state.TimeoutCount > 0 {
		state.TimeoutDetected = true
	}

	

	if net.Connected &&
		!net.IPAssigned &&
		state.NoLeaseCount > 0 {

		state.LeaseFailure = true
	}

	
	if net.Connected &&
		net.IPAssigned &&
		state.LeaseCount > 10 {

		state.ExcessiveRenewals = true
	}

	return state
}