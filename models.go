package main

import (
	// "encoding/xml"
)

// CfgData struct for configuracion
type CfgData struct {
	SMTPUrl         string `json:"smtpUrl"`
	SMTPPort        int    `json:"smtpPort"`
	SMTPUser        string `json:"smtpUser"`
	SMTPPassword    string `json:"smtpPassword"`
	Titular			string `json:"titular"`
    Banco			string `json:"banco"`
    Cuit			string `json:"cuit"`
    Cbu				string `json:"cbu"`
    Cta				string `json:"cta"`
}

// MailData data struct
type MailData struct {
	To            string `json:"To"`
	Cc            string `json:"Cc"`
	Subject       string `json:"Subject"`
	Body          string `json:"Body"`
	Attachments   []Attachment `json:"Attachments"`
	Slots         []Slot `json:"Slots"`
}

// Attachment data struct
type  Attachment struct {
	Attachment string `json:"Attachment"`
}

// Slot data struct
type  Slot struct {
	ID string `json:"Id"`
	Value string `json:"Value"`
}