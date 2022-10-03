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

	db.DeleteTable(Data.Config.DbPath, name, uint16(id))
	Data.TableList = db.SelectTableList(Data.Config.DbPath)

	log.Println("INFO: Deleted table", name)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func dashboard_rename(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	name := r.FormValue("name")

	oldName := ""
	for _, oneTable := range Data.TableList {
		if oneTable.Id == idStr {
			oldName = oneTable.Name
		}
	} 
	id, _ := strconv.Atoi(idStr)

	db.UpdateTable(Data.Config.DbPath, oldName, name, uint16(id))
	Data.TableList = db.SelectTableList(Data.Config.DbPath)

	log.Println("INFO: Updated table name:", oldName, "->", name)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}