package web

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/aceberg/HomeLists/internal/check"
	"github.com/aceberg/HomeLists/internal/db"
	"github.com/aceberg/HomeLists/internal/models"
)

func watchlist(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	TableList = db.SelectTableList(AppConfig.DbPath)

	guiData.Config = AppConfig
	guiData.TableList = TableList
	guiData.CurrentTable = "Watchlist"
	guiData.WatchList = db.SelectWatchList(AppConfig.DbPath)

	tmpl, _ := template.ParseFS(TemplHTML, "templates/watchlist.html", "templates/header.html", "templates/footer.html")
	err := tmpl.ExecuteTemplate(w, "header", guiData)
	check.IfError(err)
	err = tmpl.ExecuteTemplate(w, "watchlist", guiData)
	check.IfError(err)
}

func add_to_watchlist(w http.ResponseWriter, r *http.Request) {
	var wItem models.WatchItem

	wItem.Table = r.FormValue("cur_table")
	wItem.Name = r.FormValue("name")
	idStr := r.FormValue("id")
	wItem.ItemID, _ = strconv.Atoi(idStr)

	watchList := db.SelectWatchList(AppConfig.DbPath)

	found := false
	for _, searchItem := range watchList {
		if searchItem.Table == wItem.Table && searchItem.ItemID == wItem.ItemID {
			found = true
			break
		}
	}

	if !found {
		db.InsertWatchItem(AppConfig.DbPath, wItem)
	}

	path := "/table/" + wItem.Table

	http.Redirect(w, r, path, http.StatusFound)
}
