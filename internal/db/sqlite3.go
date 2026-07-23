package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"

	"github.com/aceberg/HomeLists/internal/check"
)

func db_exec(path string, sqlStatement string) {
	db, _ := sql.Open("sqlite", path)
	defer db.Close()

	_, err := db.Exec(sqlStatement)
	check.IfError(err)
}

func db_select(path string, table string) *sql.Rows {
	db, _ := sql.Open("sqlite", path)
	defer db.Close()

	sqlStatement := `SELECT * FROM '%s';`
	sqlStatement = fmt.Sprintf(sqlStatement, quote_str(table))

	res, err := db.Query(sqlStatement)
	check.IfError(err)

	return res
}

func db_select_count(path string, table string, id int) int {
	db, _ := sql.Open("sqlite", path)
	defer db.Close()

	sqlStatement := `SELECT "COUNT" FROM '%s' WHERE ID = ?;`
	sqlStatement = fmt.Sprintf(sqlStatement, quote_str(table))

	var count int
	err := db.QueryRow(sqlStatement, id).Scan(&count)
	check.IfError(err)

	return count
}

func db_select_name(path string, table string, id int) string {
	db, _ := sql.Open("sqlite", path)
	defer db.Close()

	sqlStatement := `SELECT "NAME" FROM '%s' WHERE ID = ?;`
	sqlStatement = fmt.Sprintf(sqlStatement, quote_str(table))

	var name string
	err := db.QueryRow(sqlStatement, id).Scan(&name)
	check.IfError(err)

	return name
}
