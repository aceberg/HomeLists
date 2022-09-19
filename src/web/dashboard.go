package web

import (
	// "fmt"
	"net/http"
	"html/template"
)

func dashboard(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("templates/dashboard.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "dashboard", TableList)
}