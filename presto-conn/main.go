package main

import (
	"database/sql"
	_ "github.com/prestodb/presto-go-client/presto"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)

func main() {
	dsn := "http://user@localhost:8080?catalog=default&schema=test"
	db, err := sql.Open("presto", dsn)

	if err != nil {
		panic("Unble to connect")
	}
	fmt.Println(db)
}
