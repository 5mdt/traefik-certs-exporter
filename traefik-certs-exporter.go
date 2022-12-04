// exports traefik's certificates as files
package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// ACME is an imported global traefik acme.json sctructure
type ACME struct {
	Le struct {
		Account struct {
			Email string
		} `json: "Account"`
		Certificates []Certificates
	} `json: "le"`
}

// Certificates list in acme.json
type Certificates struct {
	Domain struct {
		Main string `json: "main"`
	} `json: "domain"`
	Certificate string `json: "certificate"`
	Key         string `json: "key"`
}

func main() {
	// open json file
	jsonFile, err := os.Open("input/acme.json")
	// if os.Open returns an error then print out it
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully Opened acme.json")
	// defer the closing of the json file
	defer jsonFile.Close()

	// create directory for certs
	path := "output"
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var traefik ACME
	err = json.Unmarshal(byteValue, &traefik)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(traefik.Le.Certificates); i++ {
		// print cert hostname
		log.Println("Certificate host:", traefik.Le.Certificates[i].Domain.Main)
		// decode and print certificate
		certBody, err := base64.StdEncoding.DecodeString(traefik.Le.Certificates[i].Certificate)
		if err != nil {
			log.Println(err)
		}
		log.Println("Save certificate body", traefik.Le.Certificates[i].Domain.Main)
		// create cert file
		certFile, err := os.Create("output/" + traefik.Le.Certificates[i].Domain.Main + ".cer")
		if err != nil {
			log.Println(err)
		}
		defer certFile.Close()

		_, err = certFile.Write(certBody)
		if err != nil {
			log.Println(err)
		}

		// decode and print key
		certKey, err := base64.StdEncoding.DecodeString(traefik.Le.Certificates[i].Key)
		if err != nil {
			log.Println(err)
		}
		log.Println("Save certificate key", traefik.Le.Certificates[i].Domain.Main)
		// create cert file
		keyFile, err := os.Create("output/" + traefik.Le.Certificates[i].Domain.Main + ".key")
		if err != nil {
			log.Println(err)
		}
		defer keyFile.Close()

		_, err = keyFile.Write(certKey)
		if err != nil {
			log.Println(err)
		}

	}
}
