package web

import (
	"net/http"
	"strconv"

	"github.com/aceberg/HomeLists/internal/db"
	"github.com/aceberg/HomeLists/internal/models"
)

func update_watch(w http.ResponseWriter, r *http.Request) {
	var wItem models.WatchItem

	idStr := r.FormValue("id")
	wItem.ID, _ = strconv.Atoi(idStr)

	wItem.ByDate = r.FormValue("bydate")
	wItem.Date = r.FormValue("date")
	wItem.ByCount = r.FormValue("bycount")

	countStr := r.FormValue("count")
	wItem.Count, _ = strconv.Atoi(countStr)

	db.UpdateWatchItem(AppConfig.DbPath, wItem)

	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}

func del_watch(w http.ResponseWriter, r *http.Request) {

	idStr := r.FormValue("id")
	id, _ := strconv.Atoi(idStr)

	db.DeleteWatchItem(AppConfig.DbPath, id)

	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}
