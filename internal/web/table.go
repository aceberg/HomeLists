package web

import (
	"html"
	"html/template"
	"net/http"
	"sort"
	"strings"

	"github.com/aceberg/HomeLists/internal/check"
	"github.com/aceberg/HomeLists/internal/db"
	"github.com/aceberg/HomeLists/internal/models"
)

func table(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	TableList = db.SelectTableList(AppConfig.DbPath)

	urlString := html.EscapeString(r.URL.Path)
	tags := strings.Split(urlString, "/")

	guiData.CurrentTable = tags[2]

	found := false
	for _, oneTable := range TableList {
		if oneTable.Name == guiData.CurrentTable {
			found = true
		}
	}

	if found {
		guiData.ItemList = db.SelectOneTable(AppConfig.DbPath, guiData.CurrentTable)

		sort.SliceStable(guiData.ItemList, func(i, j int) bool {
			return guiData.ItemList[i].Sort < guiData.ItemList[j].Sort
		})

		lines := len(guiData.ItemList)
		db.UpdateTable(AppConfig.DbPath, lines, guiData.CurrentTable)
		TableList = db.SelectTableList(AppConfig.DbPath)

		guiData.Config = AppConfig
		guiData.TableList = TableList

		tmpl, _ := template.ParseFS(TemplHTML, "templates/table.html", "templates/header.html", "templates/footer.html")
		err := tmpl.ExecuteTemplate(w, "header", guiData)
		check.IfError(err)
		err = tmpl.ExecuteTemplate(w, "table", guiData)
		check.IfError(err)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
