package networking

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type NetworkState struct {
	Connected         bool
	IPAssigned        bool
	DefaultRoute      bool
	InternetReachable bool
	DNSWorking        bool
	Signal            int
}

func run(cmd string, args ...string) string {
	out, _ := exec.Command(cmd, args...).CombinedOutput()
	return string(out)
}

func GetCurrentNetworkState() NetworkState {

	var state NetworkState

	//-----------------------------------
	// WiFi Connection
	//-----------------------------------

	link := run("iw", "dev", "wlan0", "link")

	if strings.Contains(link, "Connected to") {
		state.Connected = true
	}

	//-----------------------------------
	// Signal Strength
	//-----------------------------------

	re := regexp.MustCompile(`signal:\s*(-?\d+)`)
	match := re.FindStringSubmatch(link)

	if len(match) == 2 {

		dbm, err := strconv.Atoi(match[1])

		if err == nil {
			state.Signal = dbm
		}
	}

	//-----------------------------------
	// IP Address
	//-----------------------------------

	ip := run("ip", "-4", "addr", "show", "wlan0")

	if strings.Contains(ip, "inet ") {
		state.IPAssigned = true
	}

	//-----------------------------------
	// Default Route
	//-----------------------------------

	route := run("ip", "route")

	if strings.Contains(route, "default via") {
		state.DefaultRoute = true
	}

	//-----------------------------------
	// Internet Reachability
	//-----------------------------------

	if exec.Command(
		"ping",
		"-c", "1",
		"-W", "2",
		"8.8.8.8",
	).Run() == nil {

		state.InternetReachable = true
	}

	//-----------------------------------
	// DNS
	//-----------------------------------

	if exec.Command(
		"getent",
		"hosts",
		"google.com",
	).Run() == nil {

		state.DNSWorking = true
	}

	return state
}

func DiagnoseWiFi() {

	state := GetCurrentNetworkState()

	fmt.Println("=== Current State ===")
	fmt.Printf("Connected: %v\n", state.Connected)
	fmt.Printf("IP Assigned: %v\n", state.IPAssigned)
	fmt.Printf("Default Route: %v\n", state.DefaultRoute)
	fmt.Printf("Internet Reachable: %v\n", state.InternetReachable)
	fmt.Printf("DNS Working: %v\n", state.DNSWorking)
	fmt.Printf("Signal: %d dBm\n", state.Signal)
	fmt.Println()

	//---------------------------------------------------
	// NOT CONNECTED
	//---------------------------------------------------

	if !state.Connected {

		nm := run(
			"journalctl",
			"-u",
			"NetworkManager",
			"-n",
			"50",
			"--no-pager",
		)

		kernel := run(
			"journalctl",
			"-k",
			"-n",
			"50",
			"--no-pager",
		)

		fmt.Println("=== Diagnosis ===")

		switch {

		case strings.Contains(strings.ToLower(nm), "handshake failed"):

			fmt.Println("Issue: Authentication Failure")
			fmt.Println()
			fmt.Println("Reason:")
			fmt.Println("The access point rejected authentication.")
			fmt.Println()
			fmt.Println("Possible Causes:")
			fmt.Println("- Wrong password")
			fmt.Println("- WPA mismatch")
			fmt.Println("- AP compatibility issue")

		case strings.Contains(strings.ToLower(kernel), "rfkill"):

			fmt.Println("Issue: Wireless Disabled")
			fmt.Println()
			fmt.Println("Reason:")
			fmt.Println("The wireless adapter is blocked by RFKill.")
			fmt.Println()
			fmt.Println("Possible Causes:")
			fmt.Println("- Airplane mode")
			fmt.Println("- Hardware switch")
			fmt.Println("- BIOS wireless setting")

		case strings.Contains(strings.ToLower(kernel), "firmware crash"):

			fmt.Println("Issue: WiFi Driver Failure")
			fmt.Println()
			fmt.Println("Reason:")
			fmt.Println("Firmware crash detected.")
			fmt.Println()
			fmt.Println("Possible Causes:")
			fmt.Println("- Driver bug")
			fmt.Println("- Firmware bug")
			fmt.Println("- Hardware instability")

		default:

			fmt.Println("Issue: Not Connected To WiFi")
			fmt.Println()
			fmt.Println("Reason:")
			fmt.Println("No active association exists.")
		}

		return
	}

	//---------------------------------------------------
	// NO IP
	//---------------------------------------------------

	if !state.IPAssigned {

		fmt.Println("=== Diagnosis ===")
		fmt.Println("Issue: No IP Address")
		fmt.Println()
		fmt.Println("Reason:")
		fmt.Println("WiFi connection exists but DHCP did not assign an address.")
		fmt.Println()
		fmt.Println("Possible Causes:")
		fmt.Println("- DHCP timeout")
		fmt.Println("- DHCP server unavailable")
		fmt.Println("- Router misconfiguration")

		return
	}

	//---------------------------------------------------
	// NO ROUTE
	//---------------------------------------------------

	if !state.DefaultRoute {

		fmt.Println("=== Diagnosis ===")
		fmt.Println("Issue: Missing Default Route")
		fmt.Println()
		fmt.Println("Reason:")
		fmt.Println("An IP exists but no route to external networks was configured.")
		fmt.Println()
		fmt.Println("Possible Causes:")
		fmt.Println("- DHCP routing failure")
		fmt.Println("- Static route issue")

		return
	}

	//---------------------------------------------------
	// INTERNET FAILURE
	//---------------------------------------------------

	if !state.InternetReachable {

		fmt.Println("=== Diagnosis ===")
		fmt.Println("Issue: Connected To WiFi But Internet Unavailable")
		fmt.Println()

		fmt.Println("Evidence:")
		fmt.Println("✓ WiFi Connected")
		fmt.Println("✓ IP Assigned")
		fmt.Println("✓ Route Present")
		fmt.Println("✗ Internet Reachability Failed")

		fmt.Println()

		fmt.Println("Likely Causes:")
		fmt.Println("- Mobile hotspot not forwarding traffic")
		fmt.Println("- Router upstream outage")
		fmt.Println("- ISP issue")
		fmt.Println("- Captive portal")

		return
	}

	//---------------------------------------------------
	// DNS FAILURE
	//---------------------------------------------------

	if !state.DNSWorking {

		fmt.Println("=== Diagnosis ===")
		fmt.Println("Issue: DNS Failure")
		fmt.Println()

		fmt.Println("Evidence:")
		fmt.Println("✓ Internet Reachable")
		fmt.Println("✗ DNS Resolution Failed")

		fmt.Println()

		fmt.Println("Likely Causes:")
		fmt.Println("- DNS server unavailable")
		fmt.Println("- Broken resolver configuration")
		fmt.Println("- VPN DNS issue")

		return
	}

	//---------------------------------------------------
	// WEAK SIGNAL
	//---------------------------------------------------

	if state.Signal != 0 && state.Signal < -75 {

		fmt.Println("=== Diagnosis ===")
		fmt.Println("Issue: Weak WiFi Signal")
		fmt.Println()

		fmt.Printf("Signal Strength: %d dBm\n", state.Signal)

		fmt.Println()

		fmt.Println("Likely Causes:")
		fmt.Println("- Distance from access point")
		fmt.Println("- RF interference")
		fmt.Println("- Weak hotspot signal")

		return
	}

	//---------------------------------------------------
	// HEALTHY
	//---------------------------------------------------

	fmt.Println("=== Diagnosis ===")
	fmt.Println("Issue: Network Healthy")
	fmt.Println()

	fmt.Println("Summary:")
	fmt.Println("WiFi association, IP assignment, routing, internet connectivity and DNS resolution are all functioning correctly.")
}


