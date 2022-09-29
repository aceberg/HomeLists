package db

import (
	"fmt"
	// "log"
	. "github.com/aceberg/HomeLists/models"
)

func UpdateOneTable(path string, table string, item Item) {
	sqlStatement := `UPDATE "%s" SET 
			DATE = '%s', NAME = '%s', COLOR = '%s', 
			COUNT = '%d', PLACE = '%s', SORT = '%d'
			WHERE ID = '%d';`
  	sqlStatement = fmt.Sprintf(sqlStatement, table, item.Date, item.Name, item.Color, item.Count, item.Place, item.Sort, item.Id)
	db_exec(path, sqlStatement)
}

func DeleteItem(path string, table string, id uint16) {
	sqlStatement := `DELETE FROM "%s" WHERE ID = '%d';`
  	sqlStatement = fmt.Sprintf(sqlStatement, table, id)
	db_exec(path, sqlStatement)
}