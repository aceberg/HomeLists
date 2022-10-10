package web

import (
	"github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
	"html"
	"html/template"
	"net/http"
	"sort"
	"strings"
)

func table(w http.ResponseWriter, r *http.Request) {
	var guiData GuiData

	TableList = db.SelectTableList(AppConfig.DbPath)

	urlString := html.EscapeString(r.URL.Path)
	tags := strings.Split(urlString, "/")

	guiData.CurrentTable = tags[2]

	check := false
	for _, oneTable := range TableList {
		if oneTable.Name == guiData.CurrentTable {
			check = true
		}
	}

	if check {
		guiData.ItemList = db.SelectOneTable(AppConfig.DbPath, guiData.CurrentTable)

		sort.SliceStable(guiData.ItemList, func(i, j int) bool {
			return guiData.ItemList[i].Sort < guiData.ItemList[j].Sort
		})

		lines := len(guiData.ItemList)
		db.UpdateTable(AppConfig.DbPath, uint16(lines), guiData.CurrentTable)
		TableList = db.SelectTableList(AppConfig.DbPath)

		guiData.Config = AppConfig
		guiData.TableList = TableList

		tmpl, _ := template.ParseFiles("templates/table.html", "templates/header.html", "templates/footer.html")
		tmpl.ExecuteTemplate(w, "header", guiData)
		tmpl.ExecuteTemplate(w, "table", guiData)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
