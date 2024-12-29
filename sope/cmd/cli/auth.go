package main

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

func doAuth() error {
	// migrations
	dbType := sop.DB.DataType
	fileName := fmt.Sprintf("%d_create_auth_tables", time.Now().UnixMicro())
	upFile := sop.RootPath + "/migrations/" + fileName + ".up.sql"
	downFile := sop.RootPath + "/migrations/" + fileName + ".down.sql"

	err := copyFilefromTemplate("templates/migrations/auth_tables."+dbType+".sql", upFile)
	if err != nil {
		exitGracefully(err)
	}

	err = copyDataToFile([]byte("drop table if exists users cascade; drop table if exists tokens cascade; drop table if exists remember_tokens;"), downFile)
	if err != nil {
		exitGracefully(err)
	}

	// run migrations
	err = doMigrate("up", "")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/data/user.go.txt", sop.RootPath+"/data/user.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/data/token.go.txt", sop.RootPath+"/data/token.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/middleware/auth-token.go.txt", sop.RootPath+"/middleware/auth-token.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/middleware/auth.go.txt", sop.RootPath+"/middleware/auth.go")
	if err != nil {
		exitGracefully(err)
	}

	color.Green("-Created users Table")
	color.Green("-Created tokens Table")
	color.Green("-Created remember_tokens Table")
	color.Green("-user and token model created")
	color.Green("-auth middleware created")
	color.Green("-Migration ran successfully")
	color.Yellow("Don't forget to uncomment user and token models in data/models.go and add appropriate middleware to your routes")

	return nil
}
