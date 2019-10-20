package models

import "encoding/xml"

// Lcos ...
type Lcos struct {
	XMLName xml.Name `xml:"LCO"`
	Lcos    []Lco    `xml:"Contribuyente"`
}
