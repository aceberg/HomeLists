package db

import (
	"fmt"
	"log"
	"strconv"
)

func MinusItem(path string, table string, id int) {

	count := db_select_count(path, table, id)
	countBefore := count

	if count > 0 {
		count = count - 1
	}

	sqlStatement := `UPDATE '%s' SET COUNT = '%d' WHERE ID = '%d';`
	sqlStatement = fmt.Sprintf(sqlStatement, quote_str(table), count, id)

	db_exec(path, sqlStatement)

	log.Println("Minus 1 on item "+db_select_name(path, table, id)+", count:", strconv.Itoa(countBefore)+" -> "+strconv.Itoa(count))
}

func GetItemNameByID(path string, table string, id int) string {

	name := db_select_name(path, table, id)

	return name
}

func GetItemCountByID(path string, table string, id int) int {

	count := db_select_count(path, table, id)

	return count
}
