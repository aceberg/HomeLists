package web

import (
	"fmt"
	"strconv"
	"net/http"
	"html/template"
	// "github.com/aceberg/HomeLists/db"
	. "github.com/aceberg/HomeLists/models"
)

func edit_line(w http.ResponseWriter, r *http.Request) {
	var item Item

	idStr := r.FormValue("id")
	item.Date = r.FormValue("date")
	item.Name = r.FormValue("name")
	countStr := r.FormValue("count")
	placeStr := r.FormValue("place")
  
	// fmt.Println("ID TO EDIT:", idStr)

	if idStr == "" || countStr == "" {
	  fmt.Fprintf(w, "No data!")
	} else {
  
		id, _ := strconv.Atoi(idStr)
		count, _ := strconv.Atoi(countStr)
		place, _ := strconv.Atoi(placeStr)
		item.Id = uint16(id)
		item.Count = uint16(count)
		item.Place = uint16(place)
	
		tmpl, _ := template.ParseFiles("templates/table-edit.html", "templates/header.html", "templates/footer.html")
		tmpl.ExecuteTemplate(w, "table-edit", item)
	}
}