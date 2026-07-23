package db

import (
	"fmt"
)

func MinusItem(path string, table string, id int) {

	count := db_select_count(path, table, id)

	if count > 0 {
		count = count - 1
	}

	sqlStatement := `UPDATE '%s' SET COUNT = '%d' WHERE ID = '%d';`
	sqlStatement = fmt.Sprintf(sqlStatement, quote_str(table), count, id)

	db_exec(path, sqlStatement)
}

func GetItemNameByID(path string, table string, id int) string {

	name := db_select_name(path, table, id)

	return name
}
