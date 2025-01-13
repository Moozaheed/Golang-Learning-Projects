package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"log"
	"net"
	"regexp"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the domain name (or type 'exit' to quit): ")

	for scanner.Scan() {
		domain := scanner.Text()
		if domain == "exit" {
			break
		}
		if !isValidDomain(domain) {
			fmt.Println("Invalid domain format. Please try again.")
			continue
		}
		checkDomain(domain)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// isValidDomain checks if the provided domain is in a valid format.
func isValidDomain(domain string) bool {
	// Regex to validate domain name format
	regex := `^(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z0-9-]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(domain)
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// Check for MX record
	mxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Println("Error looking up MX record:", err)
	} else if len(mxRecord) > 0 {
		hasMX = true
	}

	// Check for TXT records (SPF, DMARC)
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Println("Error looking up TXT records:", err)
	} else {
		for _, txtRecord := range txtRecords {
			if strings.Contains(txtRecord, "v=spf1") {
				hasSPF = true
				spfRecord = txtRecord
			}
			if strings.Contains(txtRecord, "v=DMARC1") {
				hasDMARC = true
				dmarcRecord = txtRecord
			}
		}
	}

	// Display results
	fmt.Printf("Domain: %s\n", domain)
	fmt.Printf("Has MX record: %t\n", hasMX)
	if hasSPF {
		fmt.Printf("Has SPF record: %t\n", hasSPF)
		fmt.Printf("SPF record: %s\n", spfRecord)
	} else {
		fmt.Println("Has SPF record: false")
	}
	if hasDMARC {
		fmt.Printf("Has DMARC record: %t\n", hasDMARC)
		fmt.Printf("DMARC record: %s\n", dmarcRecord)
	} else {
		fmt.Println("Has DMARC record: false")
	}
	fmt.Println("-------------------------------------------------")
}
