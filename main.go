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

	// ARGS: Date
	date := time.Now().Format("2006-01-02")
	args := os.Args[1:]

	if len(args) > 0 {
		if args[0] == "--date" {
			date = args[1]
		}
	}

	ch := make(chan string)

	// INICIO
	fmt.Printf("Iniciado: %v\n\n", time.Now())

	fileURL := "https://cfdisat.blob.core.windows.net/lco?restype=container&comp=list&prefix=LCO_" + date
	fmt.Println("Downloading LCO ... ")
	go f.DownloadFile("LCO.xml", fileURL, ch)
	fmt.Println(<-ch)

	xmlFile, err := os.Open("LCO.xml")
	defer os.Remove("LCO.xml")
	defer xmlFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var enums m.Enumerations
	xml.Unmarshal(byteValue, &enums)

	cChannel := 0
	for i := 0; i < len(enums.Enumerations); i++ {
		lenBlobs := len(enums.Enumerations[i].Blobs)
		for j := 0; j < lenBlobs; j++ {
			fileName := enums.Enumerations[i].Blobs[j]
			fmt.Printf("Analizing ... %v :: %v\n", fileName.Name, fileName.Properties.ContentMD5)

			fmt.Println("Downloading... ", fileName.URL)
			go f.DownloadFile(fileName.Name, fileName.URL, ch)

			// fmt.Println("Extracting GZ... ", fileName.Name)
			// lcoXMLFile, errGz := f.UnGZip(fileName.Name, ".")
			// if errGz != nil {
			// 	log.Fatal(errGz)
			// }

			// fmt.Println("Deleting GZ... ", fileName.Name)
			// err := os.Remove(fileName.Name)
			// if err != nil {
			// 	fmt.Println(err)
			// }

			// fmt.Println("Processing XML... ", lcoXMLFile)
			// f.ProcessLCOFile(lcoXMLFile)

			fmt.Printf("-----:::-----:::-----\n\n")
			cChannel++
		}
	}

	cnP := make(chan string)
	for i := 0; i < cChannel; i++ {
		nf := <-ch
		fmt.Println(i, nf)
		lcoXMLFile, errGz := f.UnGZip(nf, ".")
		if errGz != nil {
			log.Fatal(errGz)
		}
		go f.ProcessLCOFile(lcoXMLFile, cnP)
	}

	for i := 0; i < cChannel; i++ {
		nf := <-cnP
		fmt.Println(i, nf)
	}
	fmt.Printf("Finalizado: %v\n", time.Now())

}
