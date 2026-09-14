package main

import (
	"fmt"

	"linistic/internal/networking"
)

func main() {

	diag := networking.Correlate()

	fmt.Println("========== LINISTIC ==========")
	fmt.Println()

	fmt.Println("Issue:", diag.Issue)
	fmt.Println("Confidence:", diag.Confidence)
	fmt.Println()

	fmt.Println("Reasoning:")

	for _, reason := range diag.Reasoning {
		fmt.Println("-", reason)
	}
}
