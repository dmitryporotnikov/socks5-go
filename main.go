package main

import (
	"log"

	"github.com/armon/go-socks5"
)

func main() {
	// --- 1. Defines Credentials ---
	// This is a simple in-memory map
	const requiredUser = "dmitry"
	const requiredPass = "<SECRET_PASSWORD>"

	// --- 2. Creates an Authenticator ---
	cator := socks5.UserPassAuthenticator{
		Credentials: socks5.StaticCredentials{
			requiredUser: requiredPass,
		},
	}

	// --- 3. Configures the SOCKS5 Server ---
	conf := &socks5.Config{

		AuthMethods: []socks5.Authenticator{cator},

		// Prolly should restrict source IP? Leaving for now open for any IP.
		// RuleSet: socks5.PermitAll(), // Default is PermitAll

		Logger: log.New(log.Writer(), "[SOCKS5] ", log.LstdFlags),
	}

	server, err := socks5.New(conf)
	if err != nil {
		log.Fatalf("Failed to create SOCKS5 server: %v", err)
	}

	// --- 4. Starts Listening ---
	listenAddr := "0.0.0.0:40048"
	log.Printf("SOCKS5 proxy server starting on %s...", listenAddr)

	// Uses ListenAndServe to start the server.
	if err := server.ListenAndServe("tcp", listenAddr); err != nil {
		log.Fatalf("Failed to start SOCKS5 server: %v", err)
	}
}
