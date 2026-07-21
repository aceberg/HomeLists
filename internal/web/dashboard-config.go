package web

import (
	"log"
	"net/http"
	"strconv"

	"github.com/aceberg/HomeLists/internal/db"
)

func dashboard_delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	name := r.FormValue("name")

	id, _ := strconv.Atoi(idStr)

	db.DeleteTable(AppConfig.DbPath, name, id)
	TableList = db.SelectTableList(AppConfig.DbPath)

	log.Println("INFO: Deleted table", name)

	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}

func dashboard_rename(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	name := r.FormValue("name")

	oldName := ""
	for _, oneTable := range TableList {
		if oneTable.ID == idStr {
			oldName = oneTable.Name
		}
	}
	id, _ := strconv.Atoi(idStr)

	db.RenameTable(AppConfig.DbPath, oldName, name, id)
	TableList = db.SelectTableList(AppConfig.DbPath)

	log.Println("INFO: Updated table name:", oldName, "->", name)

	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}
