package web

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
)

// func minus_count() {}

// func edit_line() {}

func new_line(w http.ResponseWriter, r *http.Request) {
	var item Item

	place, _ := strconv.ParseUint(r.FormValue("place"), 10, 16)
	item.Place = uint16(place)

	db.InsertOneTable(AppConfig.DbPath, CurrentTable, item)
  
	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func update_line(w http.ResponseWriter, r *http.Request) {
	var item Item

	idStr := r.FormValue("id")
	item.Date = r.FormValue("date")
	item.Name = r.FormValue("name")
	countStr := r.FormValue("count")
  
	if idStr == "" || countStr == "" {
	  fmt.Fprintf(w, "No data!")
	} else {
  
	id, _ := strconv.Atoi(idStr)
	count, _ := strconv.Atoi(countStr)
	item.Id = uint16(id)
	item.Count = uint16(count)
  
	if count >= 0 {
		db.UpdateOneTable(AppConfig.DbPath, CurrentTable, item)
	}
  
	http.Redirect(w, r, r.Header.Get("Referer"), 302)
	}
}