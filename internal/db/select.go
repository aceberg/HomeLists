package db

import (
	"log"
	"sort"

	"github.com/aceberg/HomeLists/internal/models"
)

func SelectTableList(path string) []models.Table {
	rows := db_select(path, MainTable)

	tableList := []models.Table{}
	for rows.Next() {
		var oneItem models.Table
		err := rows.Scan(&oneItem.ID, &oneItem.Name, &oneItem.Date, &oneItem.Lines)
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

func SelectOneTable(path string, tableName string) []models.Item {
	rows := db_select(path, tableName)

	itemList := []models.Item{}
	for rows.Next() {
		var oneItem models.Item
		err := rows.Scan(&oneItem.ID, &oneItem.Date, &oneItem.Name, &oneItem.Color, &oneItem.Count, &oneItem.Place, &oneItem.Sort)
		if err != nil {
			log.Fatal("ERROR: SelectOneTable: ", err)
		}

		oneItem.Date = unquote_str(oneItem.Date)
		oneItem.Name = unquote_str(oneItem.Name)
		oneItem.Color = unquote_str(oneItem.Color)

		itemList = append(itemList, oneItem)
	}
	return itemList
}

func SelectWatchList(path string) []models.WatchItem {
	rows := db_select(path, WatchTable)

	watchList := []models.WatchItem{}
	for rows.Next() {
		var oneItem models.WatchItem
		err := rows.Scan(&oneItem.ID, &oneItem.Table, &oneItem.ItemID, &oneItem.Name, &oneItem.ByDate, &oneItem.Date, &oneItem.ByCount, &oneItem.Count)
		if err != nil {
			log.Fatal("ERROR: SelectWatchList: ", err)
		}

		oneItem.Table = unquote_str(oneItem.Table)
		oneItem.Name = unquote_str(oneItem.Name)

		watchList = append(watchList, oneItem)
	}

	sort.SliceStable(watchList, func(i, j int) bool {
		return watchList[i].Table < watchList[j].Table
	})

	return watchList
}
