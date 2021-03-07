package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/david-yappeter/GormCrudGenerator/gormgenerator"
	"github.com/david-yappeter/GormCrudGenerator/setting"
	"gopkg.in/yaml.v3"

	//External Library Dependency
	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"
)

func main() {
	arguments := os.Args[1:]
	GenerateService(arguments)
}

//GenerateService Generate
func GenerateService(arguments []string) {

	if len(arguments) > 0 {
		if strings.EqualFold(arguments[0], "init") {
			defaultConfig := `
database:
    type:
        # Only The First One Will Be Applied
        - mysql
        - postgre
    path: ./config
    name: databaseGorm
    apply: true
    setting:
        path: ./logger
        name: logMode
        singularTable: true
        tablePrefix: ""
        logLevel:
            # Only The First One Will Be Applied
            - Info
            - Silent
            - Warn
            - Error
        slowThreshold: 1
        apply: true
service:
    from:
        path: ./gormgenerator
        name: model
        # Ignore Model (Case-Sensitive)
        ignore:
            - 
    to:
        path: ./service
        postfix: "Generated"
    apply: true
queryTools:
    path: ./tools
    name: dbGenerator
    apply: true
`

			var d yaml.Node
			err := yaml.Unmarshal([]byte(defaultConfig), &d)
			if err != nil {
				panic(err)
			}

			settingOutput, err := yaml.Marshal(&d)
			if err != nil {
				panic(err)
			}

			file, err := os.Create(fmt.Sprintf("./gormCrud.yaml"))

			if err != nil {
				panic(err)
			}

			_, err = file.WriteString(string(settingOutput))

			if err != nil {
				panic(err)
			}

		} else {
			panic("Not A Valid Arguments")
		}

		return
	}

	var settingsYaml setting.YamlSettings

	body, err := ioutil.ReadFile("./gormCrud.yaml")

	if err != nil {
		log.Println("Please Run With Arguments 'init' if you didn't have the config file")
		log.Println("Please Check Your Yaml File, Name it 'gormCrud.yaml'")
		panic(err)
	}

	err = yaml.Unmarshal(body, &settingsYaml)

	if err != nil {
		panic(err)
	}

	connectionType := settingsYaml.Database.Type[0]

	// Get gomod name of your workspace
	goModName := gormgenerator.GetGoModName()

	gormgenerator.GormLogGenerator(settingsYaml, settingsYaml.Database.Setting.Apply)
	gormgenerator.GormConnectionGenerator(settingsYaml, goModName, connectionType, settingsYaml.Database.Apply)
	gormgenerator.GormQueryToolsGenerator(settingsYaml, settingsYaml.QueryTools.Apply)
	// gormgenerator.PaginationVariableGenerator()

	if settingsYaml.Service.Apply {
		listStruct, attributesList := gormgenerator.GetStructAndAttribute(settingsYaml)

		for _, val := range listStruct {

			err := gormgenerator.CrudGenerator(settingsYaml, goModName, val, attributesList)

			if err != nil {
				panic(err)
			}
		}
	}
}
