package models

import "encoding/xml"

// Lco ...
type Lco struct {
	XMLName     xml.Name    `xml:"Contribuyente"`
	Rfc         string      `xml:"RFC,attr"`
	Certificado Certificado `xml:"Certificado"`
}
