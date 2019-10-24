package db

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/lib/pq"
)

// MyPG ...
func MyPG(fileName string) {

	db, err := sql.Open("postgres", ConfPG())
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

	txn, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := txn.Prepare(pq.CopyIn("lco", "rfc", "certificado", "validez", "estatus", "fecha_inicio", "fecha_final"))
	fmt.Println(stmt)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	c := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if c == 0 {
			c++
			continue
		}
		line := scanner.Text()
		lineStr := strings.Split(line, ",")
		validez, _ := strconv.ParseUint(lineStr[2], 10, 8)
		_, err = stmt.Exec(lineStr[0], lineStr[1], validez, lineStr[3], lineStr[4], lineStr[5])
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = txn.Commit()
	if err != nil {
		log.Fatal(err)
	}

}
