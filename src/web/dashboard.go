package web

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
	"time"
)

func dashboard(w http.ResponseWriter, r *http.Request) {

	CurrentTable = "fTBZ96"

	tmpl, _ := template.ParseFiles("templates/dashboard.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "dashboard", TableList)
}

func add_table(w http.ResponseWriter, r *http.Request) {
	var newTable Table
	newTable.Name = r.FormValue("name")

	if newTable.Name == "" {
		fmt.Fprintf(w, "No data!")
	} else {
		currentTime := time.Now()
		newTable.Date = currentTime.Format("2006-01-02")
		log.Println("INFO: Added table", newTable)

		db.InsertTableList(AppConfig.DbPath, newTable)
		db.CreateTable(AppConfig.DbPath, newTable.Name)
		TableList = db.SelectTableList(AppConfig.DbPath)
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}