package main

import (
  "encoding/json"
  "encoding/base64"
  "io/ioutil"
  "os"
  "log"
)

type ACME struct {
  Le struct {
    Account struct {
      Email string
    } `json: "Account"`
    Certificates []Certificates
  } `json: "le"`
}

type Certificates struct {
  Domain struct {
    Main string `json: "main"`
  }`json: "domain"`
  Certificate string `json: "certificate"`
  Key string `json: "key"`
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
    cert_body, err := base64.StdEncoding.DecodeString(traefik.Le.Certificates[i].Certificate)
    if err != nil {
      log.Println(err)
    }
    log.Println("Save certificate body", traefik.Le.Certificates[i].Domain.Main)
    // create cert file
    cert_file, err := os.Create("output/"+ traefik.Le.Certificates[i].Domain.Main + ".cer")
    if err != nil {
        log.Println(err)
    }
    defer cert_file.Close()

    _, err = cert_file.Write(cert_body)
    if err != nil {
        log.Println(err)
    }

    // decode and print key
    cert_key, err := base64.StdEncoding.DecodeString(traefik.Le.Certificates[i].Key)
    if err != nil {
      log.Println(err)
    }
    log.Println("Save certificate key", traefik.Le.Certificates[i].Domain.Main)
    // create cert file
    key_file, err := os.Create("output/"+ traefik.Le.Certificates[i].Domain.Main + ".key")
    if err != nil {
        log.Println(err)
    }
    defer key_file.Close()

    _, err = key_file.Write(cert_key)
    if err != nil {
        log.Println(err)
    }

  }
}
