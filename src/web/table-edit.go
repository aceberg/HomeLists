package web

import (
	// "fmt"
	// "strconv"
	"net/http"
	"html/template"
	// "github.com/aceberg/HomeLists/db"
	//. "github.com/aceberg/HomeLists/models"
)

func edit_line(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("templates/table-edit.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "header", Data)
	tmpl.ExecuteTemplate(w, "table-edit", Data)
}