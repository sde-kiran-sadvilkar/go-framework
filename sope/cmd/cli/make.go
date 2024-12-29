package main

import (
	"errors"
	"fmt"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"os"
	"strings"
	"time"
)

func doMake(arg2, arg3 string) error {

	switch arg2 {
	case "migration":
		dbType := sop.DB.DataType
		if arg3 == "" {
			exitGracefully(errors.New("you must give the migration a name"))
		}

		fileName := fmt.Sprintf("%d_%s", time.Now().UnixMicro(), arg3)

		upFile := sop.RootPath + "/migrations/" + fileName + "." + dbType + ".up.sql"
		downFile := sop.RootPath + "/migrations/" + fileName + "." + dbType + ".down.sql"

		err := copyFilefromTemplate("templates/migrations/migration."+dbType+".up.sql", upFile)
		if err != nil {
			exitGracefully(err)
		}

		err = copyFilefromTemplate("templates/migrations/migration."+dbType+".down.sql", downFile)
		if err != nil {
			exitGracefully(err)
		}

	case "auth":
		err := doAuth()
		if err != nil {
			exitGracefully(err)
		}

	case "controller":
		if arg3 == "" {
			exitGracefully(errors.New("you must give the controller a name"))
		}
		fileName := sop.RootPath + "/controllers/" + strings.ToLower(arg3) + ".go"

		if fileExists(fileName) {
			exitGracefully(errors.New(fileName + " file already exists"))
		}
		data, err := templateFS.ReadFile("templates/controllers/controller.go.txt")

		if err != nil {
			exitGracefully(err)
		}

		controller := string(data)
		controller = strings.ReplaceAll(controller, "$CONTROLLERNAME$", strcase.ToCamel(arg3))

		err = os.WriteFile(fileName, []byte(controller), 0644)
		if err != nil {
			exitGracefully(err)
		}

	case "model":
		if arg3 == "" {
			exitGracefully(errors.New("you must give the model a name"))
		}

		data, err := templateFS.ReadFile("templates/data/model.go.txt")
		if err != nil {
			exitGracefully(err)
		}

		model := string(data)

		plur := pluralize.NewClient()

		var modelName = arg3
		var tableName = arg3

		if plur.IsPlural(arg3) {
			modelName = plur.Singular(arg3)
			tableName = strings.ToLower(tableName)
		} else {
			tableName = strings.ToLower(plur.Plural(arg3))
		}

		fileName := sop.RootPath + "/data/" + strings.ToLower(modelName) + ".go"
		if fileExists(fileName) {
			exitGracefully(errors.New(fileName + " already exists!"))
		}

		model = strings.ReplaceAll(model, "$MODELNAME$", strcase.ToCamel(modelName))
		model = strings.ReplaceAll(model, "$TABLENAME$", tableName)

		err = copyDataToFile([]byte(model), fileName)
		if err != nil {
			exitGracefully(err)
		}

	case "session":
		err := doSessionTable()
		if err != nil {
			exitGracefully(err)
		}

	}

	return nil
}
