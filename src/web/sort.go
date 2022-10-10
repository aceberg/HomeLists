package web

import (
	// "fmt"
	"github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
	"net/http"
	"strconv"
)

func sort_by_id(w http.ResponseWriter, r *http.Request) {

	currentTable := r.FormValue("cur_table")

	itemList := db.SelectOneTable(AppConfig.DbPath, currentTable)

	for _, oneItem := range itemList {
		oneItem.Sort = oneItem.Id
		db.UpdateItem(AppConfig.DbPath, currentTable, oneItem)
	}

	path := "/table/" + currentTable

	http.Redirect(w, r, path, 302)
}

func sort_before(w http.ResponseWriter, r *http.Request) {
	var item Item

	currentTable := r.FormValue("cur_table")
	idStr := r.FormValue("id")
	beforeStr := r.FormValue("before")

	if beforeStr != "" {
		id, _ := strconv.Atoi(idStr)
		before, _ := strconv.Atoi(beforeStr)

		itemList := db.SelectOneTable(AppConfig.DbPath, currentTable)

		for _, oneItem := range itemList {
			if oneItem.Id == uint16(id) {
				item = oneItem
			}
		}

		for _, oneItem := range itemList {
			if oneItem.Sort == uint16(before) {
				item.Sort = oneItem.Sort
				oneItem.Sort = oneItem.Sort + 1
				db.UpdateItem(AppConfig.DbPath, currentTable, oneItem)
			}
			if oneItem.Sort > uint16(before) {
				oneItem.Sort = oneItem.Sort + 1
				db.UpdateItem(AppConfig.DbPath, currentTable, oneItem)
			}
		}
		db.UpdateItem(AppConfig.DbPath, currentTable, item)
	}

	path := "/table/" + currentTable

	http.Redirect(w, r, path, 302)
}
