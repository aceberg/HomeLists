package main

import (
	_ "time/tzdata"

	"github.com/aceberg/HomeLists/internal/conf"
	"github.com/aceberg/HomeLists/internal/db"
	"github.com/aceberg/HomeLists/internal/web"
)

func main() {
	appConfig := conf.GetConfig()

	db.CreateDB(appConfig.DbPath)

	web.Webgui(appConfig)
}
