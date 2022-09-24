package db

import (
	"fmt"
	// "log"
	. "github.com/aceberg/HomeLists/models"
)

func InsertTableList(path string, table Table) {
	sqlStatement := `INSERT INTO "fTBZ96" (NAME, DATE) 
					 VALUES ('%s','%s');`
  	sqlStatement = fmt.Sprintf(sqlStatement, table.Name, table.Date)
	db_exec(path, sqlStatement)
}

func InsertOneTable(path string, table string, item Item) {
	sqlStatement := `INSERT INTO "%s" (DATE, NAME, COLOR, COUNT, PLACE) 
					 VALUES ('%s','%s','%s','%d','%d');`
  	sqlStatement = fmt.Sprintf(sqlStatement, table, item.Date, item.Name, item.Color, item.Count, item.Place)
	db_exec(path, sqlStatement)
}