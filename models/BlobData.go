package models

import "encoding/xml"

// BlobData ...
type BlobData struct {
	XMLName    xml.Name   `xml:"Blob"`
	Name       string     `xml:"Name"`
	URL        string     `xml:"Url"`
	Properties Properties `xml:"Properties"`
}
