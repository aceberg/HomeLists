package web

import (
	"log"
	"strings"
	"html"
	"net/http"
	"io"
	"os"
	"time"
)

func backup(w http.ResponseWriter, r *http.Request) {
	urlString := html.EscapeString(r.URL.Path)
	tags := strings.Split(urlString, "/")

	method := tags[2]
	currentTime := time.Now()

	switch method {
	case "create":
		backupString := "-" + currentTime.Format("2006-01-02T15-04")

		sourceFile, _ := os.Open(AppConfig.DbPath)
		defer sourceFile.Close()

		newFile, _ := os.Create(AppConfig.DbPath + backupString)
		defer newFile.Close()

		io.Copy(newFile, sourceFile)
		log.Println("INFO: Backup file created", AppConfig.DbPath + backupString)

		http.Redirect(w, r, r.Header.Get("Referer"), 302)

	case "download":
		filename := "HomeLists-backup-" + currentTime.Format("2006-01-02T15-04") + ".db"

		w.Header().Set("Content-Disposition", "attachment; filename="+filename)
		w.Header().Set("Content-Type", "application/octet-stream")
		http.ServeFile(w, r, AppConfig.DbPath)
	default:
		http.Redirect(w, r, "/config/", 302)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
    uploadFile, _, err := r.FormFile("dbfile")
	if err != nil {
		log.Println("ERROR: Upload error:", err)
	} else {
		defer uploadFile.Close()

		newFile, _ := os.Create(AppConfig.DbPath)
		defer newFile.Close()

		io.Copy(newFile, uploadFile)
		log.Println("INFO: DB uploaded from backup")
	}
	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}