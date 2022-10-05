package web

import (
	// "fmt"
	"log"
	"net/http"
	"github.com/aceberg/HomeLists/db"
	// . "github.com/aceberg/HomeLists/models"
	"strconv"
)

func dashboard_delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	name := r.FormValue("name")

	id, _ := strconv.Atoi(idStr)

	db.DeleteTable(AppConfig.DbPath, name, uint16(id))
	TableList = db.SelectTableList(AppConfig.DbPath)

	log.Println("INFO: Deleted table", name)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func dashboard_rename(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	name := r.FormValue("name")

	oldName := ""
	for _, oneTable := range TableList {
		if oneTable.Id == idStr {
			oldName = oneTable.Name
		}
	} 
	id, _ := strconv.Atoi(idStr)

	db.RenameTable(AppConfig.DbPath, oldName, name, uint16(id))
	TableList = db.SelectTableList(AppConfig.DbPath)

	log.Println("INFO: Updated table name:", oldName, "->", name)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}