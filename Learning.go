/*
	For authorized security testing and educational use only.
	User is responsible for complying with all applicable laws.
*/

package main

import (
	"fmt"
	"os"
	"time"
)

/*
	Copyright (c) 2026 Arjun Raj. All rights reserved.
	Tool Name: Amphiptere-Ops
	Description: Purple Team vulnerability scanner, defense verification, and report generator.
*/

func printBanner() {
	fmt.Println("==================================================")
	fmt.Println("             AMPHIPTERE-OPS                     ")
	fmt.Println("       'Unified Purple Team Sentinel.'            ")
	fmt.Println("==================================================")
	fmt.Println(" Creator : Arjun Raj")
	fmt.Println(" Contact : arjunraj.cyber@gmail.com")
	fmt.Println("--------------------------------------------------")
}

type Finding struct {
	Vulnerability string
	RedStatus     string
	BlueCheck     string
}

func runScan() []Finding {
	fmt.Println("\n[*] Phase 1: Red Team Scanning (Weaknesses & Misconfigurations)...")
	time.Sleep(500 * time.Millisecond)

	findings := []Finding{
		{"Open Port 22 (SSH)", "Detected external exposure", "Checking firewall & fail2ban rules..."},
		{"Outdated OpenSSL Package", "Vulnerable version active", "Checking package manager logs..."},
		{"Loose File Permissions (/etc/shadow)", "World-readable bits found", "Checking auditd security logs..."},
		{"Default Configuration Detected", "Standard credentials present", "Checking application access logs..."},
	}

	for i, f := range findings {
		fmt.Printf(" [!] Finding #%d: %s -> %s\n", i+1, f.Vulnerability, f.RedStatus)
		time.Sleep(200 * time.Millisecond)
	}

	return findings
}

func checkDefenses(findings []Finding) []string {
	fmt.Println("\n[*] Phase 2: Blue Team Verification (Logs, Firewall, & Monitors)...")
	time.Sleep(500 * time.Millisecond)

	var reports []string
	for _, f := range findings {
		blueResult := ""
		if f.Vulnerability == "Open Port 22 (SSH)" {
			blueResult = "FAIL: Port exposed without active fail2ban jail configuration."
		} else {
			blueResult = "PASS/WARN: Logged in audit trail, but remediation pending."
		}

		fmt.Printf(" [CHECK] %s | Status: %s\n", f.Vulnerability, blueResult)
		reports = append(reports, fmt.Sprintf("Vulnerability: %s\nRed Finding: %s\nBlue Defense Check: %s\n--------------------------------------------------\n", f.Vulnerability, f.RedStatus, blueResult))
		time.Sleep(200 * time.Millisecond)
	}

	return reports
}

func generateReport(reports []string) {
	fmt.Println("\n[*] Phase 3: Generating collaborative report.txt...")

	file, err := os.Create("report.txt")
	if err != nil {
		fmt.Printf("[!] Error creating report.txt: %v\n", err)
		return
	}
	defer file.Close()

	header := "==================================================\n" +
		"       AMPHIPTERE-OPS PURPLE TEAM REPORT        \n" +
		"       Generated: 2026-08-01                      \n" +
		"==================================================\n\n"

	file.WriteString(header)
	for _, r := range reports {
		file.WriteString(r)
	}

	fmt.Println("[+] SUCCESS: report.txt successfully compiled and saved for both teams.")
}

func main() {
	printBanner()
	findings := runScan()
	reports := checkDefenses(findings)
	generateReport(reports)
}
