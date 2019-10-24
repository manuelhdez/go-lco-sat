package db

import "fmt"

// ConfMySQL ...
func ConfMySQL() string {
	user := "lco"
	pass := "123"
	port := "3306"
	host := "127.0.0.1"
	db := "lco"

	hpp := "tcp(" + host + ":" + port + ")"
	out := fmt.Sprintf("%s:%s@%s/%s?allowAllFiles=true", user, pass, hpp, db)
	return out
}

// ConfPG ...
func ConfPG() string {
	user := "postgres"
	pass := "123"
	port := "5432"
	host := "127.0.0.1"
	db := "lco"

	out := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, db)
	return out
}
