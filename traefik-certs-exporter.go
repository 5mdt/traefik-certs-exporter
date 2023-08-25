// exports traefik's certificates as files
package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// ACME is an imported global traefik acme.json structure
type ACME struct {
	Le struct {
		Account      Account
		Certificates []Certificates
	} `json:"le"`
}

// Account represents the account details in acme.json
type Account struct {
	Email string
}

// Certificates list in acme.json
type Certificates struct {
	Domain      Domain
	Certificate string `json:"certificate"`
	Key         string `json:"key"`
}

// Domain represents the domain details in acme.json
type Domain struct {
	Main string `json:"main"`
}

func main() {
	// Parse command-line arguments
	acmeJSONPath := flag.String("acmejson", "input/acme.json", "Path to the acme.json file")
	outputDir := flag.String("output", "output", "Path to the output directory")
	flag.Parse()

	// Open acme.json file
	jsonFile, err := os.Open(*acmeJSONPath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	log.Println("Successfully Opened acme.json")

	// Read JSON content
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var traefik ACME
	if err := json.Unmarshal(byteValue, &traefik); err != nil {
		log.Fatal(err)
	}

	// Create output directory if not exists
	if err := os.MkdirAll(*outputDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	for _, cert := range traefik.Le.Certificates {
		domainDir := filepath.Join(*outputDir, cert.Domain.Main)
		if err := os.MkdirAll(domainDir, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		// Decode certificate and key
		certBody, err := base64.StdEncoding.DecodeString(cert.Certificate)
		if err != nil {
			log.Println(err)
			continue
		}

		certKey, err := base64.StdEncoding.DecodeString(cert.Key)
		if err != nil {
			log.Println(err)
			continue
		}

		// Write certificate and key to files
		certPath := filepath.Join(domainDir, "cert.pem")
		if err := ioutil.WriteFile(certPath, append(certBody, []byte("\n")...), os.ModePerm); err != nil {
			log.Println(err)
		}

		keyPath := filepath.Join(domainDir, "privkey.pem")
		if err := ioutil.WriteFile(keyPath, append(certKey, []byte("\n")...), os.ModePerm); err != nil {
			log.Println(err)
		}

		log.Println("Saved certificate and key for", cert.Domain.Main)
	}
}
