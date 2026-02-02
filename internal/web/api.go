package web

import (
	"net/http"
	"strconv"

	"github.com/aceberg/HomeLists/internal/db"
)

func api_minus(w http.ResponseWriter, r *http.Request) {
	table := r.URL.Query().Get("table")
	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err == nil {
		db.MinusItem(AppConfig.DbPath, table, uint16(id))
	}
}
