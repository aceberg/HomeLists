package db

import (
	"fmt"
	// "log"
	. "github.com/aceberg/HomeLists/models"
)

func UpdateItem(path string, table string, item Item) {
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

func DeleteTable(path string, table string, id uint16) {
	sqlStatement := `DELETE FROM "%s" WHERE ID = '%d';`
  	sqlStatement = fmt.Sprintf(sqlStatement, MainTable, id)
	db_exec(path, sqlStatement)

	sqlStatement = `DROP TABLE IF EXISTS "%s";`
  	sqlStatement = fmt.Sprintf(sqlStatement, table)
	db_exec(path, sqlStatement)
}

func UpdateTable(path string, oldName, newName string, id uint16) {
	sqlStatement := `ALTER TABLE "%s" RENAME TO "%s";`
	sqlStatement = fmt.Sprintf(sqlStatement, oldName, newName)
  	db_exec(path, sqlStatement)

	sqlStatement = `UPDATE "%s" SET NAME = '%s' WHERE ID = '%d';`
  	sqlStatement = fmt.Sprintf(sqlStatement, MainTable, newName, id)
	db_exec(path, sqlStatement)
}