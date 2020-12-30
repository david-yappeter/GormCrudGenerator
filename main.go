package main

import (
	"os"
	"strings"

	"github.com/davidyap2002/GormCrudGenerator/gormgenerator"
)

func main() {
	arguments := os.Args[1:]
	GenerateService(arguments)
}

//GenerateService Generate
func GenerateService(arguments []string) {
	var connectionType string

	if len(arguments) == 0 {
		connectionType = "mysql"
	} else {
		if strings.EqualFold(arguments[0], "mysql") {
			connectionType = "mysql"
		} else if strings.EqualFold(arguments[0], "postgre") || strings.EqualFold(arguments[0], "postgres") {
			connectionType = "postgres"
		} else {
			panic("Invalid Arguments")
		}
	}

	goModName := gormgenerator.GetGoModName()

	gormgenerator.GormLogGenerator()
	gormgenerator.GormConnectionGenerator(goModName, connectionType)
	gormgenerator.GormQueryToolsGenerator()
	gormgenerator.PaginationVariableGenerator()

	listStruct, attributesList := gormgenerator.GetStructAndAttribute()

	for _, val := range listStruct {
		err := gormgenerator.CrudGenerator(goModName, val, attributesList)

		if err != nil {
			panic(err)
		}
	}

}
