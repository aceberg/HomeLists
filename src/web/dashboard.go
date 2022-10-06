package web

import (
	// "fmt"
	"log"
	"net/http"
	"html/template"
	"github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
	"time"
	"strings"
)

func dashboard(w http.ResponseWriter, r *http.Request) {
	var guiData GuiData

	guiData.Config = AppConfig
	guiData.TableList = db.SelectTableList(AppConfig.DbPath)
	guiData.CurrentTable = "Dashboard"

	tmpl, _ := template.ParseFiles("templates/dashboard.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "header", guiData)
	tmpl.ExecuteTemplate(w, "dashboard", guiData)
}

func add_table(w http.ResponseWriter, r *http.Request) {
	var newTable Table
	newTable.Name = r.FormValue("name")

	check := strings.ContainsAny(newTable.Name, "<>?/#&'\"")

	if newTable.Name != "" && newTable.Name != "." && !check {
		currentTime := time.Now()
		newTable.Date = currentTime.Format("2006-01-02")

		db.CreateTable(AppConfig.DbPath, newTable.Name)
		db.InsertTableList(AppConfig.DbPath, newTable)
		TableList = db.SelectTableList(AppConfig.DbPath)
	} else {
		log.Println("ERROR: incorrect table name", newTable.Name)
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}