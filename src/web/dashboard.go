package web

import (
	"fmt"
	// "log"
	"net/http"
	"html/template"
	"github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
	"time"
)

func dashboard(w http.ResponseWriter, r *http.Request) {

	Data.CurrentTable = "Dashboard"

	tmpl, _ := template.ParseFiles("templates/dashboard.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "header", Data)
	tmpl.ExecuteTemplate(w, "dashboard", Data)
}

func add_table(w http.ResponseWriter, r *http.Request) {
	var newTable Table
	newTable.Name = r.FormValue("name")

	if newTable.Name == "" {
		fmt.Fprintf(w, "No data!")
	} else {
		currentTime := time.Now()
		newTable.Date = currentTime.Format("2006-01-02")

		db.InsertTableList(Data.Config.DbPath, newTable)
		db.CreateTable(Data.Config.DbPath, newTable.Name)
		Data.TableList = db.SelectTableList(Data.Config.DbPath)
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}