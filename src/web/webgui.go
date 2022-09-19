package web

import (
	"fmt"
	"log"
	"net/http"
	"github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
)

var TableList []Table
var AppConfig Conf

func Webgui (appConfig Conf) {

	AppConfig = appConfig
	TableList = db.SelectTableList(AppConfig.DbPath)

	address := AppConfig.GuiIP + ":" + AppConfig.GuiPort

	log.Println("=================================== ")
	log.Println(fmt.Sprintf("Web GUI at http://%s", address))
	log.Println("=================================== ")

	http.HandleFunc("/", dashboard)
	http.HandleFunc("/table/", table)
	http.ListenAndServe(address, nil)
}