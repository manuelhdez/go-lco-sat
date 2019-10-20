package models

import "encoding/xml"

// Certificado ...
type Certificado struct {
	XMLName             xml.Name `xml:"Certificado"`
	ValidezObligaciones string   `xml:"ValidezObligaciones,attr"`
	EstatusCertificado  string   `xml:"EstatusCertificado,attr"`
	NoCertificado       string   `xml:"noCertificado,attr"`
	FechaFinal          string   `xml:"FechaFinal,attr"`
	FechaInicio         string   `xml:"FechaInicio,attr"`
}
