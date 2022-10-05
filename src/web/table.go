package web

import (
	// "fmt"
	"strings"
	"net/http"
	"html"
	"html/template"
	"github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
)

func table(w http.ResponseWriter, r *http.Request) {
	var guiData GuiData

	guiData.Config = AppConfig
	guiData.TableList = TableList

	if r.Method == "GET" {
		urlString := html.EscapeString(r.URL.Path)
		tags := strings.Split(urlString, "/")

		guiData.CurrentTable = tags[2]
		
		guiData.ItemList = db.SelectOneTable(AppConfig.DbPath, guiData.CurrentTable)
	} else {
		guiData.ItemList = []Item{}
	}

	tmpl, _ := template.ParseFiles("templates/table.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "header", guiData)
	tmpl.ExecuteTemplate(w, "table", guiData)
}