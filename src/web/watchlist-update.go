package web

import (
	"fmt"
	"github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
	"net/http"
	"strconv"
)

func update_watch(w http.ResponseWriter, r *http.Request) {
	var wItem WatchItem

	idStr := r.FormValue("id")
	wItem.Id, _ = strconv.Atoi(idStr)

	wItem.Name = r.FormValue("name")
	wItem.ByDate = r.FormValue("bydate")
	wItem.Date = r.FormValue("date")
	wItem.ByCount = r.FormValue("bycount")
	
	countStr := r.FormValue("count")
	wItem.Count, _ = strconv.Atoi(countStr)

	fmt.Println("WATCH ITEM:", wItem)

	db.UpdateWatchItem(AppConfig.DbPath, wItem)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func del_watch(w http.ResponseWriter, r *http.Request) {

	idStr := r.FormValue("id")
	id, _ := strconv.Atoi(idStr)

	db.DeleteWatchItem(AppConfig.DbPath, id)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}