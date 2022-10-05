package web

import (
	"sort"
	"strings"
	"net/http"
	"html"
	"html/template"
	"github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
)

func table(w http.ResponseWriter, r *http.Request) {
	var guiData GuiData

	if r.Method == "GET" {
		urlString := html.EscapeString(r.URL.Path)
		tags := strings.Split(urlString, "/")

		guiData.CurrentTable = tags[2]
		
		guiData.ItemList = db.SelectOneTable(AppConfig.DbPath, guiData.CurrentTable)

		sort.SliceStable(guiData.ItemList, func(i, j int) bool {
			return guiData.ItemList[i].Sort < guiData.ItemList[j].Sort
		})

		lines := len(guiData.ItemList)
		db.UpdateTable(AppConfig.DbPath, uint16(lines), guiData.CurrentTable)
		TableList = db.SelectTableList(AppConfig.DbPath)

	} else {
		guiData.ItemList = []Item{}
	}

	guiData.Config = AppConfig
	guiData.TableList = TableList

	tmpl, _ := template.ParseFiles("templates/table.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "header", guiData)
	tmpl.ExecuteTemplate(w, "table", guiData)
}