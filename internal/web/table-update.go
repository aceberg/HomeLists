package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aceberg/HomeLists/internal/check"
	"github.com/aceberg/HomeLists/internal/db"
	"github.com/aceberg/HomeLists/internal/models"
)

func del_line(w http.ResponseWriter, r *http.Request) {
	currentTable := r.FormValue("cur_table")

	id, _ := strconv.Atoi(r.FormValue("id"))

	db.DeleteItem(AppConfig.DbPath, currentTable, id)

	path := "/table/" + currentTable

	http.Redirect(w, r, path, http.StatusFound)
}

func new_line(w http.ResponseWriter, r *http.Request) {
	var item models.Item

	currentTable := r.FormValue("cur_table")
	item.Place = r.FormValue("place")

	db.InsertItem(AppConfig.DbPath, currentTable, item)

	itemList := db.SelectOneTable(AppConfig.DbPath, currentTable)
	item = itemList[len(itemList)-1]
	item.Sort = item.ID
	db.UpdateItem(AppConfig.DbPath, currentTable, item)

	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}

func update_line(w http.ResponseWriter, r *http.Request) {
	var item models.Item

	currentTable := r.FormValue("cur_table")

	idStr := r.FormValue("id")
	item.Date = r.FormValue("date")
	item.Name = r.FormValue("name")
	item.Color = r.FormValue("color")
	countStr := r.FormValue("count")
	item.Place = r.FormValue("place")
	sortStr := r.FormValue("sort")
	minus := r.FormValue("minus")

	if idStr == "" {
		_, err := fmt.Fprintf(w, "No data!")
		check.IfError(err)
	} else {
		id, _ := strconv.Atoi(idStr)
		count, err := strconv.Atoi(countStr)
		sort, _ := strconv.Atoi(sortStr)

		if err != nil {
			count = 0
		}
		if minus == "yes" {
			count = count - 1
		}
		if count < 0 {
			count = 0
		}

		item.ID = id
		item.Count = count
		item.Sort = sort

		db.UpdateItem(AppConfig.DbPath, currentTable, item)

		path := "/table/" + currentTable

		http.Redirect(w, r, path, http.StatusFound)
	}
}
