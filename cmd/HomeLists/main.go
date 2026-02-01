package main

import (
	// "fmt"
	"github.com/aceberg/HomeLists/internal/conf"
	"github.com/aceberg/HomeLists/internal/db"
	"github.com/aceberg/HomeLists/internal/web"

	_ "time/tzdata"
)

func main() {
	appConfig := conf.GetConfig()

	db.CreateDB(appConfig.DbPath)

	web.Webgui(appConfig)
}
