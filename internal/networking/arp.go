package networking

import (
	"regexp"
	"strings"
)

type ARPState struct {
	GatewayIP      string
	GatewayFound   bool
	ARPEntryExists bool

	State string

	Reachable bool
	Failed    bool
	Incomplete bool
}


func GetARPState() ARPState {

	var state ARPState

	route := run("ip", "route", "show", "default")

	re := regexp.MustCompile(`default via (\S+)`)
	match := re.FindStringSubmatch(route)

	if len(match) < 2 {
		return state
	}

	gateway := match[1]

	state.GatewayIP = gateway
	state.GatewayFound = true

	neigh := run("ip", "neigh", "show", gateway)

	if strings.TrimSpace(neigh) == "" {
		return state
	}

	state.ARPEntryExists = true

	switch {

	case strings.Contains(neigh, "REACHABLE"):
		state.State = "REACHABLE"
		state.Reachable = true

	case strings.Contains(neigh, "STALE"):
		state.State = "STALE"

	case strings.Contains(neigh, "DELAY"):
		state.State = "DELAY"

	case strings.Contains(neigh, "PROBE"):
		state.State = "PROBE"

	case strings.Contains(neigh, "FAILED"):
		state.State = "FAILED"
		state.Failed = true

	case strings.Contains(neigh, "INCOMPLETE"):
		state.State = "INCOMPLETE"
		state.Incomplete = true
	}

	return state
}

func DiagnoseARP() {

	state := GetARPState()

	println("=== ARP Analysis ===")

	if !state.GatewayFound {

		println("Issue: No Default Gateway")
		println()
		println("Reason:")
		println("No default route configured.")

		return
	}

	if !state.ARPEntryExists {

		println("Issue: Gateway Resolution Failure")
		println()
		println("Reason:")
		println("Gateway exists but no ARP entry found.")

		return
	}

	if state.Failed {

		println("Issue: Gateway Unreachable")
		println()
		println("Reason:")
		println("ARP requests to gateway failed.")

		return
	}

	if state.Incomplete {

		println("Issue: ARP Resolution Incomplete")
		println()
		println("Reason:")
		println("Linux could not resolve gateway MAC address.")

		return
	}

	println("ARP Layer Healthy")
	println()
	println("Gateway:", state.GatewayIP)
	println("State:", state.State)
}