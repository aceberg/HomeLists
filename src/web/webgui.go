package web

import (
	// "fmt"
	"embed"
	"github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
	"log"
	"net/http"
)

var AppConfig Conf
var TableList []Table

//go:embed templates/*
var TemplHTML embed.FS

func Webgui(appConfig Conf) {

	AppConfig = appConfig
	TableList = db.SelectTableList(AppConfig.DbPath)

	address := AppConfig.GuiIP + ":" + AppConfig.GuiPort

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", dashboard)
	http.HandleFunc("/add_table/", add_table)
	http.HandleFunc("/add_to_watchlist/", add_to_watchlist)
	http.HandleFunc("/backup/", backup)
	http.HandleFunc("/config/", config)
	http.HandleFunc("/dashboard_delete/", dashboard_delete)
	http.HandleFunc("/dashboard_rename/", dashboard_rename)
	http.HandleFunc("/del_line/", del_line)
	http.HandleFunc("/del_watch/", del_watch)
	http.HandleFunc("/edit_line/", edit_line)
	http.HandleFunc("/new_line/", new_line)
	http.HandleFunc("/save_config/", save_config)
	http.HandleFunc("/sort_before/", sort_before)
	http.HandleFunc("/sort_by_id/", sort_by_id)
	http.HandleFunc("/table/", table)
	http.HandleFunc("/update_line/", update_line)
	http.HandleFunc("/update_watch/", update_watch)
	http.HandleFunc("/upload/", upload)
	http.HandleFunc("/watchlist/", watchlist)
	http.ListenAndServe(address, nil)
}
