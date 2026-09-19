package networking

import (
	"strings"
	"linistic/internal/utils"
)

type FirmwareFinding struct {
	Severity string
	Device   string
	Message  string
}


func AnalyzeFirmware() []FirmwareFinding {

	var findings []FirmwareFinding

	logs := utils.Run(
		"journalctl",
		"-k",
		"-b",
		"-n",
		"500",
		"--no-pager",
	)

	lines := strings.Split(logs, "\n")

	seen := make(map[string]bool)

	for _, line := range lines {

		l := strings.ToLower(line)


		if strings.Contains(l, "failed to load firmware") {

			key := "failed_load"

			if !seen[key] {

				findings = append(findings, FirmwareFinding{
					Severity: "FAIL",
					Device:   "Unknown Device",
					Message:  "Firmware load failure detected",
				})

				seen[key] = true
			}
		}

		if strings.Contains(l, "direct firmware load failed") {

			key := "direct_load_fail"

			if !seen[key] {

				findings = append(findings, FirmwareFinding{
					Severity: "FAIL",
					Device:   "Unknown Device",
					Message:  "Direct firmware load failed",
				})

				seen[key] = true
			}
		}

		if strings.Contains(l, "firmware crash") {

			key := "Firmware_crash"

			if !seen[key] {

				findings = append(findings, FirmwareFinding{
					Severity: "FAIL",
					Device:   "Unknown Device",
					Message:  "Firmware crash detected",
				})

				seen[key] = true
			}
		}

		if strings.Contains(l, "firmware error") {

			key := "Firmware_error"

			if !seen[key] {

				findings = append(findings, FirmwareFinding{
					Severity: "FAIL",
					Device:   "Unknown Device",
					Message:  "Firmware error reported by kernel",
				})

				seen[key] = true
			}
		}

		
		if strings.Contains(line, "[Firmware Bug]") {

			key := "Firmware_bug"

			if !seen[key] {

				findings = append(findings, FirmwareFinding{
					Severity: "WARN",
					Device:   "BIOS/Firmware",
					Message:  "Firmware bug detected but kernel may have applied a workaround",
				})

				seen[key] = true
			}
		}

		//------------------------------------------------
		// GPU FIRMWARE
		//------------------------------------------------

		if strings.Contains(line, "Finished loading DMC firmware") {

			key := "intel_dmc"

			if !seen[key] {

				findings = append(findings, FirmwareFinding{
					Severity: "PASS",
					Device:   "Intel GPU",
					Message:  "DMC firmware loaded successfully",
				})

				seen[key] = true
			}
		}

		if strings.Contains(line, "GuC firmware") {

			key := "intel_guc"

			if !seen[key] {

				findings = append(findings, FirmwareFinding{
					Severity: "PASS",
					Device:   "Intel GPU",
					Message:  "GuC firmware loaded successfully",
				})

				seen[key] = true
			}
		}

		if strings.Contains(line, "HuC firmware") {

			key := "intel_huc"

			if !seen[key] {

				findings = append(findings, FirmwareFinding{
					Severity: "PASS",
					Device:   "Intel GPU",
					Message:  "HuC firmware loaded successfully",
				})

				seen[key] = true
			}
		}

		//------------------------------------------------
		// REALTEK WIFI
		//------------------------------------------------

		if strings.Contains(line, "rtw88") &&
			strings.Contains(line, "Firmware version") {

			key := "realtek_wifi"

			if !seen[key] {

				findings = append(findings, FirmwareFinding{
					Severity: "PASS",
					Device:   "Realtek WiFi",
					Message:  "Firmware loaded successfully",
				})

				seen[key] = true
			}
		}

		//------------------------------------------------
		// INTEL WIFI
		//------------------------------------------------

		if strings.Contains(l, "loaded firmware version") &&
			strings.Contains(l, "iwlwifi") {

			key := "intel_wifi"

			if !seen[key] {

				findings = append(findings, FirmwareFinding{
					Severity: "PASS",
					Device:   "Intel WiFi",
					Message:  "Firmware loaded successfully",
				})

				seen[key] = true
			}
		}
	}

	return findings
}