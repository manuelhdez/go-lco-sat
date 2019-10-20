package models

import "encoding/xml"

// Blob ...
type Blob struct {
	XMLName xml.Name   `xml:"Blobs"`
	Blobs   []BlobData `xml:"Blob"`
}
