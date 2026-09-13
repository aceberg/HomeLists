package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aceberg/HomeLists/internal/check"
	"github.com/aceberg/HomeLists/internal/db"
)

func api_minus(w http.ResponseWriter, r *http.Request) {

	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// CORS

	table := r.URL.Query().Get("table")
	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if check.IfError(err) {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodDelete {
		db.MinusItem(AppConfig.DbPath, table, id)
	}

	count := db.GetItemCountByID(AppConfig.DbPath, table, id)

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, count)
}
