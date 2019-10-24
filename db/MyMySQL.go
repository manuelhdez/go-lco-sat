package db

import (
	"bufio"
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"os"

	"github.com/go-sql-driver/mysql"
)

// MyMySQL ...
func MyMySQL(fileName string) {
	db, err := sql.Open("mysql", ConfMySQL())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	mysql.RegisterReaderHandler("data", func() io.Reader {
		var csvReader io.Reader
		myfile, _ := os.Open(fileName)
		defer myfile.Close()

		byts := &bytes.Buffer{}

		csvReader = bufio.NewReader(myfile)
		io.Copy(byts, csvReader)
		return byts
	})
	defer mysql.DeregisterReaderHandler("data")

	strData := `
		LOAD DATA LOCAL INFILE 'Reader::data' 
		INTO TABLE lco 
		FIELDS TERMINATED BY ',' 
		IGNORE 1 LINES 
		(rfc, certificado, validez, estatus, fecha_inicio, fecha_final)`
	sentenciaPreparada, err := db.Exec(strData)
	fmt.Println(sentenciaPreparada)
}
