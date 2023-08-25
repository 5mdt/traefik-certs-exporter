// exports traefik's certificates as files
package main

import (
	"encoding/base64"
	"encoding/json"
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
	// Open json file
	jsonFile, err := os.Open("input/acme.json")
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
	outputDir := "output"
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	for _, cert := range traefik.Le.Certificates {
		// Print certificate hostname
		log.Println("Certificate host:", cert.Domain.Main)

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

		// Create certificate file
		certPath := filepath.Join(outputDir, cert.Domain.Main+".cer")
		if err := ioutil.WriteFile(certPath, certBody, os.ModePerm); err != nil {
			log.Println(err)
		}

		// Create key file
		keyPath := filepath.Join(outputDir, cert.Domain.Main+".key")
		if err := ioutil.WriteFile(keyPath, certKey, os.ModePerm); err != nil {
			log.Println(err)
		}

		log.Println("Saved certificate and key for", cert.Domain.Main)
	}
}
