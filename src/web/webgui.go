package web

import (
	"fmt"
	"log"
	"net/http"
	"github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
)

var Data GuiData

func Webgui(appConfig Conf) {

	Data.Config = appConfig
	Data.TableList = db.SelectTableList(Data.Config.DbPath)

	address := Data.Config.GuiIP + ":" + Data.Config.GuiPort

	log.Println("=================================== ")
	log.Println(fmt.Sprintf("Web GUI at http://%s", address))
	log.Println("=================================== ")

	http.HandleFunc("/", dashboard)
	http.HandleFunc("/add_table/", add_table)
	http.HandleFunc("/config/", config)
	http.HandleFunc("/del_line/", del_line)
	http.HandleFunc("/edit_line/", edit_line)
	http.HandleFunc("/new_line/", new_line)
	http.HandleFunc("/save_config/", save_config)
	http.HandleFunc("/table/", table)
	http.HandleFunc("/update_line/", update_line)
	http.ListenAndServe(address, nil)
}