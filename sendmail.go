package main

import (
	"bufio"
	"os"
	"gopkg.in/gomail.v2"
    "log"
	"strings"
)

// GoSendMail sends a mail
func GoSendMail(cfg CfgData, mailData MailData) (error) {

	m := gomail.NewMessage()
	m.SetHeader("From", cfg.SMTPUser)
	m.SetHeader("To", mailData.To)
	if mailData.Cc != "" {
		adresses := strings.Split(mailData.Cc,";")
		// m.SetHeader("Cc", mailData.Cc)
		m.SetHeader("Cc", adresses...)
	}

	m.SetHeader("Subject", mailData.Subject)

	body := ""
	fileBody, err := os.Open(mailData.Body)
    if err != nil {
		// log.Fatal(err)
		log.Printf("Error al leer mailData.Body, Error : %s", err.Error())
    } else {

		defer fileBody.Close()
	
		scanner := bufio.NewScanner(fileBody)
		for scanner.Scan() {
			body += scanner.Text()
		}
	
		if err := scanner.Err(); err != nil {
			// log.Fatal(err)
			log.Printf("Error en bufio.NewScanner(fileBody), Error : %s", err.Error())

		}
	}

	// Interpolate Fixed Data
	Interpolate( &body, "titular", cfg.Titular)
	Interpolate( &body, "banco", cfg.Banco)
	Interpolate( &body, "cuit", cfg.Cuit)
	Interpolate( &body, "cbu", cfg.Cbu)
	Interpolate( &body, "cta", cfg.Cta)

	// Interpolate Slots
	for _, slot := range mailData.Slots {
		Interpolate( &body, slot.ID, slot.Value)
	}

	m.SetBody("text/html",body)

	for _, att := range mailData.Attachments {
		m.Attach(att.Attachment)
	}

	d := gomail.NewDialer(cfg.SMTPUrl, cfg.SMTPPort, cfg.SMTPUser, cfg.SMTPPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}



