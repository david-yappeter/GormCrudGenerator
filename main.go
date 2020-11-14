package main

import (
	"github.com/davidyap2002/GormCrudGenerator/gormgenerator"
)

func main() {
	GenerateService()
}

//GenerateService Generate
func GenerateService() {
	goModName := gormgenerator.GetGoModName()

	gormgenerator.GormLogGenerator()
	gormgenerator.GormConnectionGenerator(goModName)
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
