package web

import (
	"fmt"
	"log"
	"net/http"
	"github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
)

var TableList []Table

func Webgui (AppConfig Conf) {

	TableList = db.SelectTableList(AppConfig.DbPath)

	address := AppConfig.GuiIP + ":" + AppConfig.GuiPort

	log.Println("=================================== ")
	log.Println(fmt.Sprintf("Web GUI at http://%s", address))
	log.Println("=================================== ")

	http.HandleFunc("/", dashboard)
	http.ListenAndServe(address, nil)

	// fmt.Println("Webgui will be here")
	// fmt.Println("Config:", AppConfig)

	// http.HandleFunc("/", index)
	// http.HandleFunc("/minus_count/", minus_count)
	// http.HandleFunc("/new_line/", new_line)
	// http.HandleFunc("/remove_line/", remove_line)
	// http.HandleFunc("/update_db/", update_db)
	// http.ListenAndServe(":8834", nil)
}