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

// CSVFileName para escribir una lista de los archivos descargados
const CSVFileName = "csv_files.txt"

func main() {

	var createdFiles []string

	date := time.Now().Format("2006-01-02")
	args := os.Args[1:]

	if len(args) > 0 {
		if args[0] == "--date" {
			date = args[1]
		}
	}

	ch := make(chan string)

	// INICIO
	fmt.Printf("Starting... %v\n\n", time.Now())

	fileURL := "https://cfdisat.blob.core.windows.net/lco?restype=container&comp=list&prefix=LCO_" + date
	fmt.Println("Downloading LCO...")
	go f.DownloadFile("LCO.xml", fileURL, ch)
	_ = <-ch

	xmlFile, err := os.Open("LCO.xml")

	createdFiles = append(createdFiles, "LCO.xml")

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

			fmt.Println("Downloading...", fileName.URL)
			go f.DownloadFile(fileName.Name, fileName.URL, ch)

			createdFiles = append(createdFiles, fileName.Name)

			cChannel++
		}
	}

	cnP := make(chan string)
	for i := 0; i < cChannel; i++ {
		nf := <-ch
		fmt.Println("Extracting...", nf)
		lcoXMLFile, errGz := f.UnGZip(nf, ".")
		if errGz != nil {
			log.Fatal(errGz)
		}
		fmt.Println("Processing...", nf)
		go f.ProcessLCOFile(lcoXMLFile, cnP)
		createdFiles = append(createdFiles, lcoXMLFile)
	}

	csvFiles := ""
	for i := 0; i < cChannel; i++ {
		nf := <-cnP
		fmt.Println(nf)
		csvFiles = csvFiles + nf + "\r\n"
	}

	fmt.Println("Deleting files...")
	xmlFile.Close()
	for _, f := range createdFiles {
		os.Remove(f)
	}

	err = ioutil.WriteFile(CSVFileName, []byte(csvFiles), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Finishing... %v\n", time.Now())

}
