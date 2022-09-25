package web

import (
	// "fmt"
	// "strconv"
	"net/http"
	"html/template"
	// "github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
)

func edit_line(w http.ResponseWriter, r *http.Request) {
	var item Item

	item = OneItem

	tmpl, _ := template.ParseFiles("templates/table-edit.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "table-edit", item)
}