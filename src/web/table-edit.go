package web

import (
	// "fmt"
	"strconv"
	"net/http"
	"html/template"
	// "github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
)

func edit_line(w http.ResponseWriter, r *http.Request) {
	var guiData GuiData

	guiData.Config = AppConfig
	guiData.TableList = TableList

	guiData.CurrentTable = r.FormValue("cur_table")

	idStr := r.FormValue("id")
	guiData.OneItem.Date = r.FormValue("date")
	guiData.OneItem.Name = r.FormValue("name")
	guiData.OneItem.Color = r.FormValue("color")
	countStr := r.FormValue("count")
	guiData.OneItem.Place = r.FormValue("place")
	sortStr := r.FormValue("sort")

	id, _ := strconv.Atoi(idStr)
	count, _ := strconv.Atoi(countStr)
	sort, _ := strconv.Atoi(sortStr)
	
	guiData.OneItem.Id = uint16(id)
	guiData.OneItem.Count = uint16(count)
	guiData.OneItem.Sort = uint16(sort)

	tmpl, _ := template.ParseFiles("templates/table-edit.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "header", guiData)
	tmpl.ExecuteTemplate(w, "table-edit", guiData)
}