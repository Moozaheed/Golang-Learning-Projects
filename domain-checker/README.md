
# Domain Checker

This Go program checks whether a domain has the necessary DNS records, including MX, SPF, and DMARC. It performs DNS lookups to verify the existence of these records for a given domain and provides detailed output regarding each record type.

## Features:
- Checks if a domain has an MX (Mail Exchange) record.
- Verifies if the domain has an SPF (Sender Policy Framework) record.
- Verifies if the domain has a DMARC (Domain-based Message Authentication, Reporting & Conformance) record.
- Displays the SPF and DMARC records if they are found.

## Prerequisites:
- Go 1.16+ installed on your system.

## Installation:

1. Clone the repository or download the Go source file.

    ```bash
    git clone https://github.com/moozaheed/Golang-Learning-Projects/domain-checker.git
    cd domain-checker
    ```

2. If you haven't installed Go on your machine, download and install it from the official website:
   [Go Downloads](https://golang.org/dl/)

3. Run the Go program:

    ```bash
    go run main.go
    ```

## Usage:

1. When the program runs, it will prompt you to enter a domain name.

    ```
    Enter the domain name:
    ```

2. Enter the domain you want to check (e.g., `mailchimp.com`).

3. The program will check the domain for the following DNS records:
    - **MX (Mail Exchange)**: Checks whether the domain has an MX record (used for mail routing).
    - **SPF (Sender Policy Framework)**: Checks if the domain has an SPF record to specify allowed mail senders.
    - **DMARC (Domain-based Message Authentication, Reporting & Conformance)**: Checks if the domain has a DMARC record to protect against email spoofing.

4. After checking, the program will print the results in the following format:

    ```
    Domain: mailchimp.com
    Has MX record: true
    Has SPF record: true
    SPF record: v=spf1 ip4:205.201.128.0/20 ip4:198.2.128.0/18 ip4:148.105.0.0/16 ip4:129.145.74.12 include:_spf.google.com include:mailsenders.netsuite.com include:_spf2.intuit.com include:_spf.qualtrics.com ip4:199.33.145.1 ip4:199.33.145.32 ip4:35.176.132.251 ip4:52.60.115.116 ~all
    Has DMARC record: false
    DMARC record: 
    -------------------------------------------------
    ```


## Error Handling:
- If an error occurs while looking up DNS records, the program logs the error and provides a message indicating the failure to fetch records.
  
  Example error message:
  ```
  Error: lookup example.com: no such host
  ```

