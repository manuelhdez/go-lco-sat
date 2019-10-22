package misc

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	m "lco/models"
	"os"
)

// MaxToFlush es el maximo de bytes antes de hacer flush al
// NewWriter de bufio
const MaxToFlush = 250000

// ProcessLCOFile ...
func ProcessLCOFile(file string, ch chan<- string) {

	rfcFile, _ := os.Open(file)
	defer rfcFile.Close()

	re := bufio.NewReader(rfcFile)
	_, _ = io.ReadFull(re, make([]byte, 200))

	decoder := xml.NewDecoder(re)

	// Crear archivo CSV: START
	txtFile, errFile := os.Create(file + ".csv")
	if errFile != nil {
		panic(errFile)
	}
	defer txtFile.Close()

	w := bufio.NewWriter(txtFile)
	lines := "RFC,CERTIFICADO,VALIDEZ,ESTATUS,FECHA_INICIO,FECHA_FINAL"
	// Crear archivo CSV: END

	for {
		// Escribir en Archivo: START
		if len(lines) > MaxToFlush {
			w.WriteString(lines)
			w.Flush()
			lines = ""
		}
		// Escribir en Archivo: END

		t, err := decoder.Token()
		if err != nil {
			//panic(err)
			if len(lines) > 0 {
				w.WriteString(lines)
				w.Flush()
			}
			fmt.Printf("Terminando de procesar...%v\n", file)
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
				// fmt.Printf("%v :: %v \n", rfc.Rfc, rfc.Certificado.NoCertificado)
				lines += "\n" + RfcLineData(rfc)
			default:
				// fmt.Printf("Unexpected SE {%s}%s\n", x.Name.Space, x.Name.Local)
			}
		case xml.EndElement:
			switch x.Name {
			case xml.Name{Space: "", Local: "input"}:
				fmt.Println("end of input")
				return
			default:
				// fmt.Printf("Unexpected EE {%s}%s\n", x.Name.Space, x.Name.Local)
			}
		}
	}
	ch <- file
}

// RfcLineData ...
func RfcLineData(rfc m.Lco) string {
	return rfc.Rfc + "," + rfc.Certificado.NoCertificado + "," + rfc.Certificado.ValidezObligaciones + "," + rfc.Certificado.EstatusCertificado + "," + rfc.Certificado.FechaInicio + "," + rfc.Certificado.FechaFinal
}
