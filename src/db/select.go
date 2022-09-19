package db

import (
	"fmt"
	"log"
	// "database/sql"
	. "github.com/aceberg/HomeLists/models"
)

func SelectTableList(path string) ([]Table) {
	rows := db_select(path, "fTBZ96")

	tableList := []Table{}
	for rows.Next() {
		var oneItem Table
		err := rows.Scan(&oneItem.Id, &oneItem.Name, &oneItem.Date)
		if err != nil {
			log.Fatal("ERROR: SelectTableList: ", err)
		}
		
		tableList = append(tableList, oneItem)
	}
	fmt.Println("TABLES:", tableList)
	return tableList
}