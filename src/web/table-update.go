package web

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
)

func del_line(w http.ResponseWriter, r *http.Request) {

	idInt, _ := strconv.Atoi(r.FormValue("id"))
	id := uint16(idInt)

	db.DeleteItem(Data.Config.DbPath, Data.CurrentTable, id)
  
	path := "/table/" + Data.CurrentTable

	http.Redirect(w, r, path, 302)
}

func new_line(w http.ResponseWriter, r *http.Request) {
	var item Item

	item.Place = r.FormValue("place")

	db.InsertOneTable(Data.Config.DbPath, Data.CurrentTable, item)
  
	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func update_line(w http.ResponseWriter, r *http.Request) {
	var item Item

	idStr := r.FormValue("id")
	item.Date = r.FormValue("date")
	item.Name = r.FormValue("name")
	item.Color = r.FormValue("color")
	countStr := r.FormValue("count")
	item.Place = r.FormValue("place")
	minus := r.FormValue("minus")
	edit := r.FormValue("edit")
  
	if idStr == "" || countStr == "" {
	  fmt.Fprintf(w, "No data!")
	} else {
		id, _ := strconv.Atoi(idStr)
		count, _ := strconv.Atoi(countStr)
		
		if minus == "yes" {
			count = count - 1
		}
		if count < 0 {
			count = 0
		}

		item.Id = uint16(id)
		item.Count = uint16(count)

		if edit == "yes" {
			Data.OneItem = item
			edit_line(w, r)

		} else { 
			db.UpdateOneTable(Data.Config.DbPath, Data.CurrentTable, item)
	
			path := "/table/" + Data.CurrentTable

			http.Redirect(w, r, path, 302)
		}
	}
}