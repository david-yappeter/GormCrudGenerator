package cmd

import (
	"github.com/david-yappeter/GormCrudGenerator/generator"
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
	goModName := generator.GetGoModName()
	generator.GormLogGenerator(settingsYaml)
	generator.GormConnectionGenerator(settingsYaml, goModName)
	generator.GormQueryToolsGenerator(settingsYaml)

	if settingsYaml.Service.Apply {
		listStruct, attributesList := generator.GetStructAndAttribute(settingsYaml)
		for _, val := range listStruct {
			err := generator.CrudGenerator(settingsYaml, goModName, val, attributesList)
			if err != nil {
				panic(err)
			}
		}
	}
}

// func
