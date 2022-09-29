package db

import (
	"os"
	"log"
	"fmt"
)

func CreateDB(path string) {
	if _, err := os.Stat(path); err == nil {
        log.Println("INFO: DB exists")
    } else {
		sqlStatement := `CREATE TABLE "fTBZ96" (
			"ID"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
			"NAME"	TEXT NOT NULL,
			"DATE"	TEXT NOT NULL
		);`
    	db_exec(path, sqlStatement)
		log.Println("INFO: DB created!")
    }
}

func CreateTable(path string, tableName string) {
	sqlStatement := `CREATE TABLE "%s" (
		"ID"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		"DATE"	TEXT NOT NULL,
		"NAME"	TEXT NOT NULL,
		"COLOR"	TEXT NOT NULL,
		"COUNT"	INTEGER DEFAULT 0,
		"PLACE"	TEXT NOT NULL,
		"SORT"	INTEGER DEFAULT 0
	);`
	sqlStatement = fmt.Sprintf(sqlStatement, tableName)
    db_exec(path, sqlStatement)
	log.Println("INFO: Created table", tableName)
}