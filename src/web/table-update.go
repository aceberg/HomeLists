package web

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
)

func del_line(w http.ResponseWriter, r *http.Request) {
	var currentTable string

	currentTable = r.FormValue("cur_table")

	idInt, _ := strconv.Atoi(r.FormValue("id"))
	id := uint16(idInt)

	db.DeleteItem(AppConfig.DbPath, currentTable, id)
  
	path := "/table/" + currentTable

	http.Redirect(w, r, path, 302)
}

func new_line(w http.ResponseWriter, r *http.Request) {
	var item Item
	var currentTable string

	currentTable = r.FormValue("cur_table")
	item.Place = r.FormValue("place")

	db.InsertItem(AppConfig.DbPath, currentTable, item)

	itemList := db.SelectOneTable(AppConfig.DbPath, currentTable)
	item = itemList[len(itemList)-1]
	item.Sort = item.Id
	db.UpdateItem(AppConfig.DbPath, currentTable, item)
  
	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func update_line(w http.ResponseWriter, r *http.Request) {
	var item Item
	var currentTable string

	currentTable = r.FormValue("cur_table")

	idStr := r.FormValue("id")
	item.Date = r.FormValue("date")
	item.Name = r.FormValue("name")
	item.Color = r.FormValue("color")
	countStr := r.FormValue("count")
	item.Place = r.FormValue("place")
	sortStr := r.FormValue("sort")
	minus := r.FormValue("minus")
  
	if idStr == "" || countStr == "" {
	  fmt.Fprintf(w, "No data!")
	} else {
		id, _ := strconv.Atoi(idStr)
		count, _ := strconv.Atoi(countStr)
		sort, _ := strconv.Atoi(sortStr)
		
		if minus == "yes" {
			count = count - 1
		}
		if count < 0 {
			count = 0
		}

		item.Id = uint16(id)
		item.Count = uint16(count)
		item.Sort = uint16(sort)

		db.UpdateItem(AppConfig.DbPath, currentTable, item)
	
		path := "/table/" + currentTable

		http.Redirect(w, r, path, 302)
	}
}