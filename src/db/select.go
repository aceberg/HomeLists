package db

import (
	// "fmt"
	"log"
	// "database/sql"
	. "github.com/aceberg/HomeLists/models"
)

func SelectTableList(path string) ([]Table) {
	rows := db_select(path, MainTable)

	tableList := []Table{}
	for rows.Next() {
		var oneItem Table
		err := rows.Scan(&oneItem.Id, &oneItem.Name, &oneItem.Date)
		if err != nil {
			log.Fatal("ERROR: SelectTableList: ", err)
		}
		
		tableList = append(tableList, oneItem)
	}
	// fmt.Println("TABLES:", tableList)
	return tableList
}

func SelectOneTable(path string, tableName string) ([]Item) {
	rows := db_select(path, tableName)

	itemList := []Item{}
	for rows.Next() {
		var oneItem Item
		err := rows.Scan(&oneItem.Id, &oneItem.Date, &oneItem.Name, &oneItem.Color, &oneItem.Count, &oneItem.Place, &oneItem.Sort)
		if err != nil {
			log.Fatal("ERROR: SelectOneTable: ", err)
		}
		
		itemList = append(itemList, oneItem)
	}
	// fmt.Println("ITEMS:", itemList)
	return itemList
}