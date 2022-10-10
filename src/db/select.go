package db

import (
	// "fmt"
	. "github.com/aceberg/HomeLists/models"
	"log"
	"sort"
)

func SelectTableList(path string) []Table {
	rows := db_select(path, MainTable)

	tableList := []Table{}
	for rows.Next() {
		var oneItem Table
		err := rows.Scan(&oneItem.Id, &oneItem.Name, &oneItem.Date, &oneItem.Lines)
		if err != nil {
			log.Fatal("ERROR: SelectTableList: ", err)
		}

		oneItem.Name = unquote_str(oneItem.Name)

		tableList = append(tableList, oneItem)
	}

	sort.SliceStable(tableList, func(i, j int) bool {
		return tableList[i].Name < tableList[j].Name
	})

	return tableList
}

func SelectOneTable(path string, tableName string) []Item {
	rows := db_select(path, tableName)

	itemList := []Item{}
	for rows.Next() {
		var oneItem Item
		err := rows.Scan(&oneItem.Id, &oneItem.Date, &oneItem.Name, &oneItem.Color, &oneItem.Count, &oneItem.Place, &oneItem.Sort)
		if err != nil {
			log.Fatal("ERROR: SelectOneTable: ", err)
		}

		oneItem.Date = unquote_str(oneItem.Date)
		oneItem.Name = unquote_str(oneItem.Name)
		oneItem.Color = unquote_str(oneItem.Color)

		itemList = append(itemList, oneItem)
	}
	// fmt.Println("ITEMS:", itemList)
	return itemList
}
