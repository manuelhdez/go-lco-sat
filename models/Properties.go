package models

import "encoding/xml"

// Properties ...
type Properties struct {
	XMLName    xml.Name `xml:"Properties"`
	ContentMD5 string   `xml:"Content-MD5"`
}
