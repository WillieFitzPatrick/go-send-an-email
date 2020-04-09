package main

import (
	"time"
	"io/ioutil"
    "encoding/json"
	"strings"
    "os"
    "log"
)

// FileExists helper func
func FileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

func checkErr(err error) bool {
	if err != nil {
        // panic(err)
	    log.Printf("Error %s",err.Error())
        return true
    }
    return false
}



// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
func Find(a []string, x string) int {
    for i, n := range a {
        if x == n {
            return i
        }
    }
    return len(a)
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
    for _, n := range a {
        if x == n {
            return true
        }
    }
    return false
}


// ReadMailData opens config.json and returns its data
func ReadMailData( file string) ( MailData, error ) {
    mailData := MailData{}
    fileData, err := os.Open(file)

	// if we os.Open returns an error then handle it
	if err != nil {
		log.Printf("Error al abrir %s, Error : %s",file,  err.Error())
		return mailData, err
	} 

	defer fileData.Close()

	data, err := ioutil.ReadAll(fileData)
	if err != nil {
		log.Printf("Error al leer datos de config.json : Error %s", err.Error())
		return mailData, err
	}

	// unmarshall it
	err = json.Unmarshal(data, &mailData)
	if err != nil  {
		log.Printf("Error al procesar datos de config.json, Error : %s", err.Error())
		return mailData, err
    }
    
    return mailData, nil

}

// ReadConfigFile opens config.json and returns its data
func ReadConfigFile() ( CfgData,error ) {
    cfgData := CfgData{}
    configFile, err := os.Open("./config.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		log.Printf("Error al abrir config.json, Error : %s", err.Error())
		return cfgData, err
	} 

	defer configFile.Close()

	data, err := ioutil.ReadAll(configFile)
	if err != nil {
		log.Printf("Error al leer datos de config.json : Error %s", err.Error())
		return cfgData ,err
	}

	// unmarshall it
	err = json.Unmarshal(data, &cfgData)
	if err != nil  {
		log.Printf("Error al procesar datos de config.json, Error : %s", err.Error())
		return cfgData, err
    }
    
    return cfgData, nil

}

//Interpolate replaces a variable with a value in a string
func Interpolate ( text *string, search string, value string)  {
    search = "{{" + search + "}}"
    *text = strings.Replace(*text, search, value, -1)
}

// LogHeader returns a header
func LogHeader() string {
    header := `
	<h1 style="width: 100%; text-align: center;">Afip-Arba Connector</h1>
	<h3 style="width: 100%; text-align: center;">Reporte de Eventos</h3>
	<hr />
	<br />
	<p><span style="font-size: 14px;padding-left: 10px;">Archivo : {{archivo}}</span></p>
	<p><span style="font-size: 14px;padding-left: 10px;">Fecha : {{fecha}}</span></p>
	`
	Interpolate( &header, "archivo", "aa-connector.log")
	currentTime := time.Now()
	Interpolate( &header, "fecha", currentTime.Format("2006-01-02 15:04:05") )

	header += `<br />`
    header += `<h3 style="width: 100%; text-align: start; padding-left: 10px;">Eventos Registrados</h3>`
    return header
}

