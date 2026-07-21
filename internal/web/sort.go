package web

import (
	"net/http"
	"strconv"

	"github.com/aceberg/HomeLists/internal/db"
	"github.com/aceberg/HomeLists/internal/models"
)

func sort_by_id(w http.ResponseWriter, r *http.Request) {

	currentTable := r.FormValue("cur_table")

	itemList := db.SelectOneTable(AppConfig.DbPath, currentTable)

	for _, oneItem := range itemList {
		oneItem.Sort = oneItem.ID
		db.UpdateItem(AppConfig.DbPath, currentTable, oneItem)
	}

	path := "/table/" + currentTable

	http.Redirect(w, r, path, http.StatusFound)
}

func sort_before(w http.ResponseWriter, r *http.Request) {
	var item models.Item

	currentTable := r.FormValue("cur_table")
	idStr := r.FormValue("id")
	beforeStr := r.FormValue("before")

	if beforeStr != "" {
		id, _ := strconv.Atoi(idStr)
		before, _ := strconv.Atoi(beforeStr)

		itemList := db.SelectOneTable(AppConfig.DbPath, currentTable)

		for _, oneItem := range itemList {
			if oneItem.ID == id {
				item = oneItem
			}
		}

		for _, oneItem := range itemList {
			if oneItem.Sort == before {
				item.Sort = oneItem.Sort
				oneItem.Sort = oneItem.Sort + 1
				db.UpdateItem(AppConfig.DbPath, currentTable, oneItem)
			}
			if oneItem.Sort > before {
				oneItem.Sort = oneItem.Sort + 1
				db.UpdateItem(AppConfig.DbPath, currentTable, oneItem)
			}
		}
		db.UpdateItem(AppConfig.DbPath, currentTable, item)
	}

	path := "/table/" + currentTable

	http.Redirect(w, r, path, http.StatusFound)
}
