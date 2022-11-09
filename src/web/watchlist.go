package web

import (
	// "fmt"
	"github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
	"html/template"
	"net/http"
	"strconv"
)

func watchlist(w http.ResponseWriter, r *http.Request) {
	var guiData GuiData

	TableList = db.SelectTableList(AppConfig.DbPath)

	guiData.Config = AppConfig
	guiData.TableList = TableList
	guiData.CurrentTable = "Watchlist"
	guiData.WatchList = db.SelectWatchList(AppConfig.DbPath)

	// fmt.Println("WL:", guiData.WatchList)

	tmpl, _ := template.ParseFS(TemplHTML, "templates/watchlist.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "header", guiData)
	tmpl.ExecuteTemplate(w, "watchlist", guiData)
}

func add_to_watchlist(w http.ResponseWriter, r *http.Request) {
	var wItem WatchItem

	wItem.Table = r.FormValue("cur_table")
	wItem.Name = r.FormValue("name")
	idStr := r.FormValue("id")
	wItem.ItemId, _ = strconv.Atoi(idStr)

	watchList := db.SelectWatchList(AppConfig.DbPath)

	check := true
	for _, searchItem := range watchList {
		if searchItem.Table == wItem.Table && searchItem.ItemId == wItem.ItemId {
			check = false
			break
		}
	}

	if check {
		db.InsertWatchItem(AppConfig.DbPath, wItem)
	}

	path := "/table/" + wItem.Table

	http.Redirect(w, r, path, 302)
}
