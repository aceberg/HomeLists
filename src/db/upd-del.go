package db

import (
	"fmt"
	// "log"
	. "github.com/aceberg/HomeLists/models"
)

func UpdateItem(path string, table string, item Item) {
	sqlStatement := `UPDATE '%s' SET 
			DATE = '%s', NAME = '%s', COLOR = '%s', 
			COUNT = '%d', PLACE = '%s', SORT = '%d'
			WHERE ID = '%d';`
	sqlStatement = fmt.Sprintf(sqlStatement, quote_str(table), quote_str(item.Date), quote_str(item.Name), quote_str(item.Color), item.Count, item.Place, item.Sort, item.Id)
	db_exec(path, sqlStatement)
}

func DeleteItem(path string, table string, id uint16) {
	sqlStatement := `DELETE FROM '%s' WHERE ID = '%d';`
	sqlStatement = fmt.Sprintf(sqlStatement, quote_str(table), id)
	db_exec(path, sqlStatement)
}

func DeleteTable(path string, table string, id uint16) {
	sqlStatement := `DELETE FROM '%s' WHERE ID = '%d';`
	sqlStatement = fmt.Sprintf(sqlStatement, MainTable, id)
	db_exec(path, sqlStatement)

	sqlStatement = `DROP TABLE IF EXISTS '%s';`
	sqlStatement = fmt.Sprintf(sqlStatement, quote_str(table))
	db_exec(path, sqlStatement)
}

func RenameTable(path string, oldName, newName string, id uint16) {
	sqlStatement := `ALTER TABLE '%s' RENAME TO '%s';`
	sqlStatement = fmt.Sprintf(sqlStatement, quote_str(oldName), quote_str(newName))
	db_exec(path, sqlStatement)

	sqlStatement = `UPDATE '%s' SET NAME = '%s' WHERE ID = '%d';`
	sqlStatement = fmt.Sprintf(sqlStatement, MainTable, quote_str(newName), id)
	db_exec(path, sqlStatement)
}

func UpdateTable(path string, lines uint16, table string) {
	sqlStatement := `UPDATE '%s' SET LINES = '%d' WHERE NAME = '%s';`
	sqlStatement = fmt.Sprintf(sqlStatement, MainTable, lines, quote_str(table))
	db_exec(path, sqlStatement)
}

func UpdateWatchItem(path string, wItem WatchItem) {
	sqlStatement := `UPDATE '%s' SET 
			NAME = '%s', BYDATE = '%s', DATE = '%s',
			BYCOUNT = '%s', COUNT = '%d'
			WHERE ID = '%d';`
	sqlStatement = fmt.Sprintf(sqlStatement, WatchTable, quote_str(wItem.Name), wItem.ByDate, quote_str(wItem.Date), wItem.ByCount, wItem.Count, wItem.Id)
	db_exec(path, sqlStatement)
}

func DeleteWatchItem(path string, id int) {
	sqlStatement := `DELETE FROM '%s' WHERE ID = '%d';`
	sqlStatement = fmt.Sprintf(sqlStatement, WatchTable, id)
	db_exec(path, sqlStatement)
}
