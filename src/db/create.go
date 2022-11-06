package db

import (
	"fmt"
	"log"
	"os"
)

func CreateDB(path string) {
	if _, err := os.Stat(path); err == nil {
		log.Println("INFO: DB exists")
	} else {
		log.Println("INFO: DB created")
	}

	sqlStatement := `CREATE TABLE IF NOT EXISTS '%s' (
		"ID"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		"NAME"	TEXT NOT NULL,
		"DATE"	TEXT NOT NULL,
		"LINES"	INTEGER DEFAULT 0
	);`
	sqlStatement = fmt.Sprintf(sqlStatement, MainTable)
	db_exec(path, sqlStatement)

	sqlStatement = `CREATE TABLE IF NOT EXISTS '%s' (
		"ID"		INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		"TABLENAME"	TEXT NOT NULL,
		"ITEMID"	INTEGER DEFAULT 0,
		"NAME"		TEXT NOT NULL,
		"BYDATE"	TEXT DEFAULT "no",
		"DATE"		TEXT DEFAULT "",
		"BYCOUNT"	TEXT DEFAULT "yes",
		"COUNT"		INTEGER DEFAULT 0
	);`
	sqlStatement = fmt.Sprintf(sqlStatement, WatchTable)
	db_exec(path, sqlStatement)
}

func CreateTable(path string, tableName string) {
	sqlStatement := `CREATE TABLE IF NOT EXISTS '%s' (
		"ID"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		"DATE"	TEXT NOT NULL,
		"NAME"	TEXT NOT NULL,
		"COLOR"	TEXT NOT NULL,
		"COUNT"	INTEGER DEFAULT 0,
		"PLACE"	TEXT NOT NULL,
		"SORT"	INTEGER DEFAULT 0
	);`
	sqlStatement = fmt.Sprintf(sqlStatement, quote_str(tableName))
	db_exec(path, sqlStatement)
	log.Println("INFO: Created table", tableName)
}
