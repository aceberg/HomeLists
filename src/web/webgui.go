package web

import (
	"fmt"
	"log"
	"net/http"
	"github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
)

var AppConfig Conf
var TableList []Table

func Webgui(appConfig Conf) {

	AppConfig = appConfig
	TableList = db.SelectTableList(AppConfig.DbPath)

	address := AppConfig.GuiIP + ":" + AppConfig.GuiPort

	log.Println("=================================== ")
	log.Println(fmt.Sprintf("Web GUI at http://%s", address))
	log.Println("=================================== ")

	http.HandleFunc("/", dashboard)
	http.HandleFunc("/add_table/", add_table)
	http.HandleFunc("/backup/", backup)
	http.HandleFunc("/config/", config)
	http.HandleFunc("/dashboard_delete/", dashboard_delete)
	http.HandleFunc("/dashboard_rename/", dashboard_rename)
	http.HandleFunc("/del_line/", del_line)
	http.HandleFunc("/edit_line/", edit_line)
	http.HandleFunc("/new_line/", new_line)
	http.HandleFunc("/save_config/", save_config)
	http.HandleFunc("/sort_before/", sort_before)
	http.HandleFunc("/sort_by_id/", sort_by_id)
	http.HandleFunc("/table/", table)
	http.HandleFunc("/update_line/", update_line)
	http.HandleFunc("/upload/", upload)
	http.ListenAndServe(address, nil)
}