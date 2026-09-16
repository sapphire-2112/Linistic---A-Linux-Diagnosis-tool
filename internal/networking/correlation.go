package networking

type Diagnosis struct {
	Issue      string
	Confidence string
	Reasoning  []string
}

func Correlate() Diagnosis {

	rfkill := GetRFKillState()
	driver := GetDriverState()
	net := GetCurrentNetworkState()

	if rfkill.HardBlocked {
		return Diagnosis{
			Issue:      "Wireless Hardware Disabled",
			Confidence: "High",
			Reasoning: []string{
				"RFKill hard block detected",
				"Wireless radio disabled by hardware",
			},
		}
	}

	if rfkill.SoftBlocked {
		return Diagnosis{
			Issue:      "Wireless Software Disabled",
			Confidence: "High",
			Reasoning: []string{
				"RFKill soft block detected",
				"Wireless radio disabled by software",
			},
		}
	}

	if !driver.HardwarePresent {
		return Diagnosis{
			Issue:      "Network Hardware Missing",
			Confidence: "High",
			Reasoning: []string{
				"No network controller detected",
			},
		}
	}

	if !driver.DriverBound {
		return Diagnosis{
			Issue:      "Network Driver Missing",
			Confidence: "High",
			Reasoning: []string{
				"Hardware detected",
				"No driver bound to device",
			},
		}
	}

	if !driver.ModuleLoaded {
		return Diagnosis{
			Issue:      "Kernel Module Missing",
			Confidence: "High",
			Reasoning: []string{
				"Driver detected",
				"Kernel module not loaded",
			},
		}
	}

	if driver.FirmwareFailure {
		return Diagnosis{
			Issue:      "Firmware Failure",
			Confidence: "High",
			Reasoning: []string{
				"Driver reported firmware loading failure",
			},
		}
	}

	if driver.PCIeFatal {
		return Diagnosis{
			Issue:      "PCIe Communication Failure",
			Confidence: "High",
			Reasoning: []string{
				"Fatal PCIe errors detected",
			},
		}
	}

	if driver.DriverReset {
		return Diagnosis{
			Issue:      "Driver Reset Detected",
			Confidence: "Medium",
			Reasoning: []string{
				"Network driver repeatedly reset device",
			},
		}
	}

	if driver.DriverTimeout {
		return Diagnosis{
			Issue:      "Driver Timeout",
			Confidence: "Medium",
			Reasoning: []string{
				"Driver timed out waiting for hardware response",
			},
		}
	}

	if !net.Connected {
		return Diagnosis{
			Issue:      "Not Connected To WiFi",
			Confidence: "High",
			Reasoning: []string{
				"Wireless stack healthy",
				"No active WiFi association",
			},
		}
	}

	dhcp := GetDHCPState()

	if dhcp.TimeoutDetected {
		return Diagnosis{
			Issue:      "DHCP Timeout",
			Confidence: "High",
			Reasoning: []string{
				"DHCP requests timed out",
				"No response received from DHCP server",
			},
		}
	}

	if dhcp.LeaseFailure {
		return Diagnosis{
			Issue:      "DHCP Lease Failure",
			Confidence: "High",
			Reasoning: []string{
				"WiFi association successful",
				"DHCP lease could not be obtained",
			},
		}
	}

	if !net.IPAssigned {
		return Diagnosis{
			Issue:      "No IP Address",
			Confidence: "High",
			Reasoning: []string{
				"WiFi connected",
				"No IPv4 address assigned",
				"Likely DHCP failure",
			},
		}
	}

	if !net.DefaultRoute {
		return Diagnosis{
			Issue:      "Routing Failure",
			Confidence: "High",
			Reasoning: []string{
				"IP address assigned",
				"No default route configured",
			},
		}
	}

	arp := GetARPState()

	if !arp.GatewayFound {
		return Diagnosis{
			Issue:      "No Default Gateway",
			Confidence: "High",
			Reasoning: []string{
				"No gateway found in ARP table",
			},
		}
	}

	if arp.Failed {
		return Diagnosis{
			Issue:      "Gateway Unreachable",
			Confidence: "High",
			Reasoning: []string{
				"ARP resolution to gateway failed",
			},
		}
	}

	if arp.Incomplete {
		return Diagnosis{
			Issue:      "ARP Resolution Failure",
			Confidence: "High",
			Reasoning: []string{
				"Gateway MAC address could not be resolved",
			},
		}
	}

	if !net.InternetReachable {
		return Diagnosis{
			Issue:      "Internet Unreachable",
			Confidence: "Medium",
			Reasoning: []string{
				"WiFi connected",
				"IP assigned",
				"Default route present",
				"Internet reachability failed",
			},
		}
	}

	if !net.DNSWorking {
		return Diagnosis{
			Issue:      "DNS Failure",
			Confidence: "High",
			Reasoning: []string{
				"Internet reachable",
				"DNS resolution failed",
			},
		}
	}

	if net.Signal != 0 && net.Signal < -75 {
		return Diagnosis{
			Issue:      "Weak WiFi Signal",
			Confidence: "Medium",
			Reasoning: []string{
				"Signal strength below recommended threshold",
			},
		}
	}

	if dhcp.ExcessiveRenewals {
		return Diagnosis{
			Issue:      "Frequent DHCP Renewals Detected",
			Confidence: "Low",
			Reasoning: []string{
				"Connection is working",
				"DHCP lease is renewing unusually often",
			},
		}
	}

	if driver.PCIeCorrectable {
		return Diagnosis{
			Issue:      "System Healthy (Recoverable PCIe Errors Observed)",
			Confidence: "Medium",
			Reasoning: []string{
				"Network functioning normally",
				"Recoverable PCIe errors were corrected by Linux",
			},
		}
	}

	return Diagnosis{
		Issue:      "System Healthy",
		Confidence: "High",
		Reasoning: []string{
			"RFKill healthy",
			"Network hardware present",
			"Driver loaded",
			"Kernel module loaded",
			"WiFi connected",
			"IP assigned",
			"Route present",
			"Gateway reachable",
			"Internet reachable",
			"DNS working",
		},
	}
}
