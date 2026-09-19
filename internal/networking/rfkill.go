package networking

import (
	"fmt"
	"strings"
	"linistic/internal/utils"
)

type RFKillState struct {
	Found       bool
	DeviceName  string
	SoftBlocked bool
	HardBlocked bool
}

func GetRFKillState() RFKillState {

	out := utils.Run("rfkill", "list")

	var state RFKillState

	if strings.TrimSpace(out) == "" {
		return state
	}

	blocks := strings.Split(out, "\n\n")

	for _, block := range blocks {

		lower := strings.ToLower(block)

		if !strings.Contains(lower, "wireless") &&
			!strings.Contains(lower, "wlan") &&
			!strings.Contains(lower, "wifi") {
			continue
		}

		state.Found = true

		lines := strings.Split(block, "\n")

		if len(lines) > 0 {
			state.DeviceName = strings.TrimSpace(lines[0])
		}

		if strings.Contains(lower, "soft blocked: yes") {
			state.SoftBlocked = true
		}

		if strings.Contains(lower, "hard blocked: yes") {
			state.HardBlocked = true
		}

		break
	}

	return state
}

func DiagnoseRFKill() {

	state := GetRFKillState()

	fmt.Println("=== RFKill Analysis ===")

	if !state.Found {

		fmt.Println("Issue: No Wireless RFKill Device Found")
		fmt.Println()
		fmt.Println("Reason:")
		fmt.Println("No wireless adapter was reported by RFKill.")

		return
	}

	fmt.Println("Device:", state.DeviceName)
	fmt.Println()

	if state.HardBlocked {

		fmt.Println("Issue: Hardware Radio Block")
		fmt.Println()

		fmt.Println("Reason:")
		fmt.Println("The wireless adapter is physically disabled.")

		fmt.Println()

		fmt.Println("Possible Causes:")
		fmt.Println("- Hardware WiFi switch disabled")
		fmt.Println("- BIOS wireless setting disabled")
		fmt.Println("- Vendor hotkey disabled radio")

		fmt.Println()

		fmt.Println("Suggested Actions:")
		fmt.Println("- Enable hardware wireless switch")
		fmt.Println("- Check BIOS wireless settings")
		fmt.Println("- Use vendor function key")

		return
	}

	if state.SoftBlocked {

		fmt.Println("Issue: Software Radio Block")
		fmt.Println()

		fmt.Println("Reason:")
		fmt.Println("The wireless adapter is blocked by software.")

		fmt.Println()

		fmt.Println("Possible Causes:")
		fmt.Println("- Airplane mode enabled")
		fmt.Println("- RFKill command used")
		fmt.Println("- Desktop network manager disabled radio")

		fmt.Println()

		fmt.Println("Suggested Actions:")
		fmt.Println("- Run: rfkill unblock wifi")
		fmt.Println("- Disable airplane mode")
		fmt.Println("- Re-enable WiFi from network settings")

		return
	}

	fmt.Println("RFKill State Healthy")
	fmt.Println()

	fmt.Println("Reason:")
	fmt.Println("No software or hardware radio blocks detected.")
}