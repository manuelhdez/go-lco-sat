package misc

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	m "lco/models"
	"os"
)

// ProcessLCOFile ...
func ProcessLCOFile(file string) {

	rfcFile, _ := os.Open(file)
	defer rfcFile.Close()

	re := bufio.NewReader(rfcFile)
	_, _ = io.ReadFull(re, make([]byte, 200))

	decoder := xml.NewDecoder(re)

	for {
		t, err := decoder.Token()
		if err != nil {
			//panic(err)
			fmt.Println("Terminando de procesar...")
			break
		}
		switch x := t.(type) {
		case xml.StartElement:
			switch x.Name {
			//case xml.Name{Space: "http:/www.sat.gob.mx/cfd/LCO", Local: "LCO"}:
			case xml.Name{Space: "lco", Local: "lco"}:

			//case xml.Name{Space: "http:/www.sat.gob.mx/cfd/LCO", Local: "Contribuyente"}:
			case xml.Name{Space: "lco", Local: "Contribuyente"}:
				var rfc m.Lco
				err = decoder.DecodeElement(&rfc, &x)
				if err != nil {
					panic(err)
				}
				fmt.Printf("%v :: %v \n", rfc.Rfc, rfc.Certificado.NoCertificado)
			default:
				//fmt.Println(t)
				fmt.Printf("Unexpected SE {%s}%s\n", x.Name.Space, x.Name.Local)
			}
		case xml.EndElement:
			switch x.Name {
			case xml.Name{Space: "", Local: "input"}:
				fmt.Println("end of input")
				return
			default:
				fmt.Printf("Unexpected EE {%s}%s\n", x.Name.Space, x.Name.Local)
			}
		}
	}

}
