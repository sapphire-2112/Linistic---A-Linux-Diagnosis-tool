package networking

import (
	"fmt"
	"regexp"
	"strings"
)


type DriverState struct {
	HardwarePresent bool
	DriverBound     bool
	ModuleLoaded    bool
	DriverName      string

	FirmwareFailure bool
	DriverReset     bool
	DriverTimeout   bool

	PCIeFatal       bool
	PCIeCorrectable bool
}


func GetDriverState() DriverState {

	var state DriverState

	lspci := strings.ToLower(run("lspci"))

	if strings.Contains(lspci, "network") ||
		strings.Contains(lspci, "wireless") ||
		strings.Contains(lspci, "ethernet") {

		state.HardwarePresent = true
	}

	driverInfo := run("ethtool", "-i", "wlan0")

	re := regexp.MustCompile(`driver:\s*(\S+)`)
	match := re.FindStringSubmatch(driverInfo)

	if len(match) >= 2 {

		state.DriverBound = true
		state.DriverName = match[1]
	}

	if state.DriverBound {

		lsmod := run("lsmod")

		if strings.Contains(lsmod, state.DriverName) {
			state.ModuleLoaded = true
		}
	}

	kernelLogs := strings.ToLower(
		run(
			"journalctl",
			"-k",
			"-n",
			"1000",
			"--no-pager",
		),
	)

	if strings.Contains(kernelLogs, "failed to load firmware") ||
		strings.Contains(kernelLogs, "firmware crash") ||
		strings.Contains(kernelLogs, "firmware not found") {

		state.FirmwareFailure = true
	}

	if strings.Contains(kernelLogs, "resetting adapter") ||
		strings.Contains(kernelLogs, "recovery triggered") {

		state.DriverReset = true
	}

	if strings.Contains(kernelLogs, "watchdog timeout") ||
		strings.Contains(kernelLogs, "timeout waiting") {

		state.DriverTimeout = true
	}

	if strings.Contains(kernelLogs, "pcie bus error") &&
		strings.Contains(kernelLogs, "correctable") {

		state.PCIeCorrectable = true
	}

	if strings.Contains(kernelLogs, "pcie bus error") &&
		(strings.Contains(kernelLogs, "fatal") ||
			strings.Contains(kernelLogs, "uncorrectable")) {

		state.PCIeFatal = true
	}

	return state
}
func DiagnoseDriver() {

	state := GetDriverState()

	fmt.Println("=== Driver Analysis ===")

	if !state.HardwarePresent {

		fmt.Println("Issue: Network Hardware Not Detected")
		fmt.Println()
		fmt.Println("Reason:")
		fmt.Println("No network controller was detected by the kernel.")
		return
	}

	if !state.DriverBound {

		fmt.Println("Issue: No Driver Bound To Device")
		fmt.Println()
		fmt.Println("Reason:")
		fmt.Println("Network hardware exists but no driver is attached.")
		return
	}

	fmt.Println("Driver:", state.DriverName)
	fmt.Println()

	if !state.ModuleLoaded {

		fmt.Println("Issue: Kernel Module Missing")
		fmt.Println()
		fmt.Println("Reason:")
		fmt.Println("Driver detected but kernel module is not loaded.")
		return
	}

	if state.FirmwareFailure {

		fmt.Println("Issue: Firmware Load Failure")
		fmt.Println()
		fmt.Println("Reason:")
		fmt.Println("Driver could not load required firmware.")
		return
	}

	if state.DriverReset {

		fmt.Println("Issue: Driver Reset Detected")
		fmt.Println()
		fmt.Println("Reason:")
		fmt.Println("Driver is repeatedly resetting the device.")
		return
	}

	if state.DriverTimeout {

		fmt.Println("Issue: Driver Timeout")
		fmt.Println()
		fmt.Println("Reason:")
		fmt.Println("Driver failed to receive expected hardware response.")
		return
	}

	if state.PCIeFatal {

		fmt.Println("Issue: PCIe Communication Failure")
		fmt.Println()
		fmt.Println("Reason:")
		fmt.Println("Fatal PCIe communication errors detected.")
		return
	}

	if state.PCIeCorrectable {

		fmt.Println("Info: Recoverable PCIe Errors")
		fmt.Println()
		fmt.Println("Reason:")
		fmt.Println("The device reported PCIe errors but Linux corrected them successfully.")
		return
	}

	fmt.Println("Issue: Driver Layer Healthy")
	fmt.Println()
	fmt.Println("Reason:")
	fmt.Println("Hardware detected, driver loaded and no driver failures observed.")
}