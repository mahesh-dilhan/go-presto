package main

import (
	"database/sql"
	"fmt"
	_ "github.com/prestodb/presto-go-client/presto"
)

type Country struct {
	name          string
	positivecases int
}

func main() {
	//catalog -> this is the properties file name in prestro
	// in mycase path  -> usr/local/opt/prestodb/libexec/etc/catalog/postgresql.properties
	//schema -> postgres schema
	dsn := "http://user@localhost:8080?catalog=postgresql&schema=play"
	db, err := sql.Open("presto", dsn)

	if err != nil {
		panic("Unble to connect")
	}

	defer db.Close()

	rows, err := db.Query(`SELECT name, positivecases FROM country`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		country := Country{}
		err = rows.Scan(&country.name, &country.positivecases)
		if err != nil {
			panic(err)
		}
		fmt.Println(country)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
