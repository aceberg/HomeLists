package web

import (
	"time"

	"github.com/aceberg/HomeLists/internal/db"
)

func renameWatchList() {

	for {
		wl := db.SelectWatchList(AppConfig.DbPath)

		for _, item := range wl {
			name := db.GetItemNameByID(AppConfig.DbPath, item.Table, item.ItemID)
			if item.Name != name {
				item.Name = name
				db.UpdateWatchItem(AppConfig.DbPath, item)
			}

		}

		time.Sleep(time.Duration(12) * time.Hour)
	}
}
