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