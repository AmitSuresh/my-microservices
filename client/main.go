package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load the client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair("../client.crt", "../client.key")
	if err != nil {
		log.Fatalf("failed to load client certificate: %v", err)
	}

	// Load the root CA cert to verify the server certificate
	caCert, err := os.ReadFile("../rootCA.crt")
	if err != nil {
		log.Fatalf("failed to read CA certificate: %v", err)
	}

	// Create a certificate pool from the CA certificate
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Configure the TLS client
	clientTLSConfig := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      caCertPool, // Verify the server's certificate against the root CA
	}

	// Create an HTTP client with the TLS configuration
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: clientTLSConfig,
		},
	}

	// Make a request to the server
	resp, err := client.Get("https://localhost:9090/api/ping")
	if err != nil {
		log.Fatalf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Print the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}
	fmt.Printf("Response: %s\n", body)
}
