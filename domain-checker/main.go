package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"log"
	"net"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the domain name: ")

	for scanner.Scan() {
		domain := scanner.Text()
		checkDomain(domain)
	}

	if err := scanner.Err(); err!= nil {
        log.Fatal(err)
    }


}

func checkDomain(domain string){
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecord, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: ", err)
		return
	}

	if len(mxRecord) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	
	if err!= nil {
        log.Printf("Error: ", err)
        return
    }

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

    fmt.Printf("Domain: %s\n", domain)
	fmt.Printf("Has MX record: %t\n", hasMX)
	fmt.Printf("Has SPF record: %t\n", hasSPF)
	fmt.Printf("SPF record: %s\n", spfRecord)
	fmt.Printf("Has DMARC record: %t\n", hasDMARC)
	fmt.Printf("DMARC record: %s\n", dmarcRecord)
	fmt.Println("-------------------------------------------------")
	fmt.Println()

}