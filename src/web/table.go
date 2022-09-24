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
	var itemList []Item
	if r.Method == "GET" {
		urlString := html.EscapeString(r.URL.Path)
		tags := strings.Split(urlString, "/")

		CurrentTable = tags[2]
		
		itemList = db.SelectOneTable(AppConfig.DbPath, CurrentTable)
	} else {
		itemList = []Item{}
	}

	tmpl, _ := template.ParseFiles("templates/table.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "table", itemList)
}