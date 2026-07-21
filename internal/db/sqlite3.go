package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func db_exec(path string, sqlStatement string) {
	db, _ := sql.Open("sqlite", path)
	defer db.Close()

	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatal("ERROR: db_exec: ", err)
	}
}

func db_select(path string, table string) *sql.Rows {
	db, _ := sql.Open("sqlite", path)
	defer db.Close()

	sqlStatement := `SELECT * FROM '%s';`
	sqlStatement = fmt.Sprintf(sqlStatement, quote_str(table))

	res, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal("ERROR: db_select: ", err)
	}

	return res
}

func db_select_count(path string, table string, id int) int {
	db, _ := sql.Open("sqlite", path)
	defer db.Close()

	sqlStatement := `SELECT "COUNT" FROM '%s' WHERE ID = ?;`
	sqlStatement = fmt.Sprintf(sqlStatement, quote_str(table))

	var count int
	err := db.QueryRow(sqlStatement, id).Scan(&count)
	if err != nil {
		log.Fatal("ERROR: db_select_count: ", err)
	}

	return count
}
