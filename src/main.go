package main

import (
	// "fmt"
	"github.com/aceberg/HomeLists/conf"
	"github.com/aceberg/HomeLists/db"
	"github.com/aceberg/HomeLists/web"

	_ "time/tzdata"
)

func main() {
	appConfig := conf.GetConfig()

	db.CreateDB(appConfig.DbPath)

	web.Webgui(appConfig)
}
