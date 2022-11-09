package web

import (
	// "fmt"
	"github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func checkDate(days string, expire string) bool {
	if expire == "" || days == "" {
		return false
	} else {
		expireDate, err := time.Parse("2006-01-02", expire)
		if err == nil {
			daysInt, _ := strconv.Atoi(days)

			nowDate := time.Now()
			plusDate := nowDate.Add(time.Duration(daysInt) * 24 * time.Hour)

			if plusDate.Before(expireDate) {
				return false
			} else {
				return true
			}
		} else {
			log.Println("ERROR: checkDate, wrong expire date", err)
			return false
		}
	}
}

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
		if tableName != watchItem.Table {
			tableName = watchItem.Table
			tableItems = db.SelectOneTable(AppConfig.DbPath, tableName)
		}
		for _, searchItem := range tableItems {
			if uint16(watchItem.ItemId) == searchItem.Id {
				if watchItem.ByCount == "yes" && searchItem.Count <= uint16(watchItem.Count) {
					searchItem.Place = tableName
					searchItem.Color = "#ffe0b3" // orange

					if watchItem.Count > 0 && searchItem.Count == 0 {
						searchItem.Color = "#ffb3b3" // red
					}

					itemList = append(itemList, searchItem)
				} else if watchItem.ByDate == "yes" && checkDate(watchItem.Date, searchItem.Date) {
					searchItem.Place = tableName
					searchItem.Color = "#ffe0b3" // orange

					itemList = append(itemList, searchItem)
				}
			}
		}
	}
	guiData.ItemList = itemList

	tmpl, _ := template.ParseFS(TemplHTML, "templates/dashboard.html", "templates/header.html", "templates/footer.html")
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
