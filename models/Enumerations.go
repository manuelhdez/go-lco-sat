package models

import "encoding/xml"

// Enumerations ...
type Enumerations struct {
	XMLName      xml.Name `xml:"EnumerationResults"`
	Enumerations []Blob   `xml:"Blobs"`
}
