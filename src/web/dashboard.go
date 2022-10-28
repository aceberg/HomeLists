package web

import (
	"fmt"
	"github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

func dashboard(w http.ResponseWriter, r *http.Request) {
	var guiData GuiData

	TableList = db.SelectTableList(AppConfig.DbPath)
	watchList := db.SelectWatchList(AppConfig.DbPath)

	guiData.Config = AppConfig
	guiData.TableList = TableList
	guiData.CurrentTable = "Dashboard"

	itemList := []Item{}
	tableName := ""
	tableItems := []Item{}
	for _, watchItem := range watchList {
		if (watchItem.ByCount == "yes" || watchItem.ByDate == "yes") {
			if tableName != watchItem.Table {
				tableName = watchItem.Table
				tableItems = db.SelectOneTable(AppConfig.DbPath, tableName)
			}
			for _, searchItem := range tableItems {
				if uint16(watchItem.ItemId) == searchItem.Id {
					if (searchItem.Count <= uint16(watchItem.Count) && watchItem.ByCount == "yes") {
						searchItem.Place = tableName
						searchItem.Color = "#ffe0b3" // orange

						if (watchItem.Count > 0 && searchItem.Count == 0) {
							searchItem.Color = "#ffb3b3" // red
						}

						itemList = append(itemList, searchItem)
					}
				}
			}
		}
	}
	guiData.ItemList = itemList

	fmt.Println("ITEM LIST", itemList)

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
