package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	f "lco/misc"
	m "lco/models"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Printf("Iniciado: %v\n", time.Now())

	fileURL := "https://cfdisat.blob.core.windows.net/lco?restype=container&comp=list&prefix=LCO_2019-10-18"
	fmt.Println("Downloading LCO ... ")
	if err := f.DownloadFile("LCO.xml", fileURL); err != nil {
		panic(err)
	}

	xmlFile, err := os.Open("LCO.xml")
	defer os.Remove("LCO.xml")
	defer xmlFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var enums m.Enumerations
	xml.Unmarshal(byteValue, &enums)

	for i := 0; i < len(enums.Enumerations); i++ {
		lenBlobs := len(enums.Enumerations[i].Blobs)
		for j := 0; j < lenBlobs; j++ {
			fileName := enums.Enumerations[i].Blobs[j]
			fmt.Printf("Analizing ... %v :: %v\n", fileName.Name, fileName.Properties.ContentMD5)

			fmt.Println("Downloading... ", fileName.URL)
			if err := f.DownloadFile(fileName.Name, fileName.URL); err != nil {
				panic(err)
			}

			fmt.Println("Extracting GZ... ", fileName.Name)
			lcoXMLFile, errGz := f.UnGZip(fileName.Name, ".")
			if errGz != nil {
				log.Fatal(errGz)
			}

			fmt.Println("Deleting GZ... ", fileName.Name)
			err := os.Remove(fileName.Name)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Processing XML... ", fileName.Name)
			f.ProcessLCOFile(lcoXMLFile)

			fmt.Printf("-----:::-----:::-----\n\n")
		}
	}

	fmt.Printf("Finalizado: %v\n", time.Now())

}
