package cmd

import (
	"github.com/david-yappeter/GormCrudGenerator/gormgenerator"
	"github.com/david-yappeter/GormCrudGenerator/setting"

	//External Library Dependency
	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"
)

//GenerateService Generate
func GenerateService(arguments []string) {
	if ok := argsCheck(arguments); ok {
		return
	}
	defaultCommand()
}

func defaultCommand() {
	var settingsYaml = setting.ReadYamlConfig()
	goModName := gormgenerator.GetGoModName()
	gormgenerator.GormLogGenerator(settingsYaml)
	gormgenerator.GormConnectionGenerator(settingsYaml, goModName)
	gormgenerator.GormQueryToolsGenerator(settingsYaml)

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

// func
