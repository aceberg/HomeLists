package db

import (
	"fmt"
	// "log"
	. "github.com/aceberg/HomeLists/models"
)

func UpdateOneTable(path string, table string, item Item) {
	sqlStatement := `UPDATE "%s" SET 
			DATE = '%s', NAME = '%s', COLOR = '%s', 
			COUNT = '%d', PLACE = '%d' 
			WHERE ID = '%d';`
  	sqlStatement = fmt.Sprintf(sqlStatement, table, item.Date, item.Name, item.Color, item.Count, item.Place, item.Id)
	db_exec(path, sqlStatement)
}