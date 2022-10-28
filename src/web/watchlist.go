package web

import (
	// "fmt"
	"github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
	"html/template"
	"net/http"
)

func watchlist(w http.ResponseWriter, r *http.Request) {
	var guiData GuiData

	TableList = db.SelectTableList(AppConfig.DbPath)

	guiData.Config = AppConfig
	guiData.TableList = TableList
	guiData.CurrentTable = "Watchlist"
	guiData.WatchList = db.SelectWatchList(AppConfig.DbPath)

	// fmt.Println("WL:", guiData.WatchList)

	tmpl, _ := template.ParseFiles("templates/watchlist.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "header", guiData)
	tmpl.ExecuteTemplate(w, "watchlist", guiData)
}
