package main

import (
	"fmt"
	"flag"
	"log"
)

const (
	// VERSION of this software
	VERSION = "0.10";
)

func main() {

	// allow some parameters as flags
	mailDataFile := flag.String("data", "", "defines the data of the email")
	flag.Parse()
	if *mailDataFile == "" {
		log.Printf("Falta el parametro --data ( nombre del json con los datos del mail)")
		return
	}

	fmt.Println("go-send-an-email started, version: ", VERSION)
	fmt.Println("loading ", *mailDataFile)

	mailData, err := ReadMailData( *mailDataFile)
	if err != nil {
		log.Printf("Error al leer archivo de datos del mail (%s) - error %s", *mailDataFile, err.Error())
	}

	fmt.Println("loading config.json")
	cfg, err := ReadConfigFile()
	if err != nil {
		log.Printf("Error al leer archivo de configuracion (config.json) - error %s", err.Error())
	}

	fmt.Println("sending email to: ", mailData.To, " with ", len(mailData.Attachments)," attachments.")
	err = GoSendMail(cfg, mailData)
	if err != nil {
		log.Printf("SendMail returned an error: %s", err.Error())
	}
	fmt.Println("email sent!")
}