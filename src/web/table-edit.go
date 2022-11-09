package web

import (
	// "fmt"
	"html/template"
	"net/http"
	"strconv"
	// "github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
)

func edit_line(w http.ResponseWriter, r *http.Request) {
	var guiData GuiData

	guiData.Colors = []LineColor{
		{Name: "Blue", Code: "#b3d1ff"},
		{Name: "Cyan", Code: "#b3ffff"},
		{Name: "Green", Code: "#b3ffb3"},
		{Name: "Orange", Code: "#ffe0b3"},
		{Name: "Pink", Code: "#ffccff"},
		{Name: "Purple", Code: "#ecb3ff"},
		{Name: "Red", Code: "#ffb3b3"},
		{Name: "Yellow", Code: "#ffffb3"},
	}

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

	tmpl, _ := template.ParseFS(TemplHTML, "templates/table-edit.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "header", guiData)
	tmpl.ExecuteTemplate(w, "table-edit", guiData)
}
