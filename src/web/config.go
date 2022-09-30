package web

import (
	"net/http"
	"html/template"
	"github.com/aceberg/HomeLists/conf"
)

func config(w http.ResponseWriter, r *http.Request) {
	Data.CurrentTable = "Config"

	Data.Themes = []string{"cerulean", "cosmo", "cyborg", "darkly", "flatly", "journal", "litera", "lumen", "lux", "materia", "minty", "morph", "pulse", "quartz", "sandstone", "simplex", "sketchy", "slate", "solar", "spacelab", "superhero", "united", "vapor", "yeti", "zephyr"}

	tmpl, _ := template.ParseFiles("templates/config.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "header", Data)
	tmpl.ExecuteTemplate(w, "config", Data)
}

func save_config(w http.ResponseWriter, r *http.Request) {

	Data.Config.Theme = r.FormValue("theme")
	conf.WriteConfig(Data.Config.Theme)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}