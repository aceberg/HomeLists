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
	if r.Method == "GET" {
		urlString := html.EscapeString(r.URL.Path)
		tags := strings.Split(urlString, "/")

		Data.CurrentTable = tags[2]
		
		Data.ItemList = db.SelectOneTable(Data.Config.DbPath, Data.CurrentTable)
	} else {
		Data.ItemList = []Item{}
	}

	tmpl, _ := template.ParseFiles("templates/table.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "header", Data)
	tmpl.ExecuteTemplate(w, "table", Data)
}