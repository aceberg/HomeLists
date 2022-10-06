package db

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func db_exec(path string, sqlStatement string) {
	db, _ := sql.Open("sqlite3", path)
	defer db.Close()
  
	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatal("ERROR: db_exec: ", err)
	}
}

func db_select(path string, table string) (*sql.Rows) {
	db, _ := sql.Open("sqlite3", path)
	defer db.Close()
  
	sqlStatement := `SELECT * FROM '%s';`
	sqlStatement = fmt.Sprintf(sqlStatement, quote_str(table))

	res, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal("ERROR: db_select: ", err)
	}
	
	// fmt.Println("RES (db_select): ", res)
	// fmt.Printf("RES TYPE: %T", res)
	return res
}