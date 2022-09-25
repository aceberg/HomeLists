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
var CurrentTable string
var OneItem Item

func Webgui (appConfig Conf) {

	AppConfig = appConfig
	TableList = db.SelectTableList(AppConfig.DbPath)

	address := AppConfig.GuiIP + ":" + AppConfig.GuiPort

	log.Println("=================================== ")
	log.Println(fmt.Sprintf("Web GUI at http://%s", address))
	log.Println("=================================== ")

	http.HandleFunc("/", dashboard)
	http.HandleFunc("/add_table/", add_table)
	http.HandleFunc("/del_line/", del_line)
	http.HandleFunc("/edit_line/", edit_line)
	http.HandleFunc("/new_line/", new_line)
	http.HandleFunc("/table/", table)
	http.HandleFunc("/update_line/", update_line)
	http.ListenAndServe(address, nil)
}