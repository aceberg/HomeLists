package web

import (
	"net/http"
	"html/template"
	"github.com/aceberg/HomeLists/conf"
	. "github.com/aceberg/HomeLists/models"
)

func config(w http.ResponseWriter, r *http.Request) {
	var guiData GuiData

	guiData.Config = AppConfig
	guiData.TableList = TableList
	guiData.CurrentTable = "Config"

	guiData.Themes = []string{"cerulean", "cosmo", "cyborg", "darkly", "flatly", "journal", "litera", "lumen", "lux", "materia", "minty", "morph", "pulse", "quartz", "sandstone", "simplex", "sketchy", "slate", "solar", "spacelab", "superhero", "united", "vapor", "yeti", "zephyr"}

	tmpl, _ := template.ParseFiles("templates/config.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "header", guiData)
	tmpl.ExecuteTemplate(w, "config", guiData)
}

func save_config(w http.ResponseWriter, r *http.Request) {

	AppConfig.Theme = r.FormValue("theme")
	conf.WriteConfig(AppConfig.Theme)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}