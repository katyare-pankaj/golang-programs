package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/miekg/dns"
)

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Set security headers
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'")
	w.Header().Set("Strict-Transport-Security", "max-age=31536000")
	w.Header().Set("X-Frame-Options", "SAMEORIGIN")

	// Perform DNS lookup for the request hostname using DNSSEC
	domain := r.Host
	if !strings.HasSuffix(domain, ".") {
		domain = domain + "."
	}

	m := new(dns.Msg)
	m.SetQuestion(domain, dns.TypeSOA)
	m.SetEdns0(4096, true)

	c := new(dns.Client)
	in, _, err := c.Exchange(m, "8.8.8.8:53") // Use Google DNS as an example
	if err != nil {
		fmt.Println("DNS lookup failed:", err)
		return
	}

	// Validate the DNSSEC response
	if !in.Rcode == dns.RcodeSuccess {
		fmt.Println("DNS lookup failed: Rcode", in.Rcode)
		return
	}

	err = in.Validate()
	if err != nil {
		fmt.Println("DNSSEC validation failed:", err)
		// You can decide to handle this error appropriately, such as returning a 503 Service Unavailable response.
		return
	}

	// Handle the request here
	fmt.Fprintf(w, "Welcome to our nonprofit application!")
}
