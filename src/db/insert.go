package db

import (
	"fmt"
	// "log"
	. "github.com/aceberg/HomeLists/models"
)

func InsertTableList(path string, table Table) {
	sqlStatement := `INSERT INTO '%s' (NAME, DATE) 
					 VALUES ('%s','%s');`
	sqlStatement = fmt.Sprintf(sqlStatement, MainTable, quote_str(table.Name), quote_str(table.Date))
	db_exec(path, sqlStatement)
}

func InsertItem(path string, table string, item Item) {
	sqlStatement := `INSERT INTO '%s' (DATE, NAME, COLOR, COUNT, PLACE, SORT) 
					 VALUES ('%s','%s','%s','%d','%s','%d');`
	sqlStatement = fmt.Sprintf(sqlStatement, quote_str(table), quote_str(item.Date), quote_str(item.Name), quote_str(item.Color), item.Count, item.Place, item.Sort)
	db_exec(path, sqlStatement)
}

func InsertWatchItem(path string, wItem WatchItem) {
	sqlStatement := `INSERT INTO '%s' (TABLENAME, ITEMID, NAME) 
					 VALUES ('%s','%d','%s');`
	sqlStatement = fmt.Sprintf(sqlStatement, WatchTable, quote_str(wItem.Table), wItem.ItemId, quote_str(wItem.Name))

	db_exec(path, sqlStatement)
}
