package gormgenerator

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/david-yappeter/GormCrudGenerator/setting"
)

//GetGoModName Get Go Mod Name
func GetGoModName() string {
	fileHandle, _ := os.Open("./go.mod")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	var temp string

	for fileScanner.Scan() {
		goModName := strings.Split(fileScanner.Text(), " ")

		if len(goModName) > 1 {
			if goModName[0] == "module" {
				temp = goModName[1]
			}
		}
	}

	return temp
}

//GetStructAndAttribute Get Structs and Its Attributes
func GetStructAndAttribute(setting setting.YamlSettings) ([]string, map[string][]string) {
	fileHandle, _ := os.Open(fmt.Sprintf("%s/%s.go", pathRemoveLastStrip(setting.Service.From.Path), setting.Service.From.Name))
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	var listStruct []string
	listAttributes := make(map[string][]string)

	for fileScanner.Scan() {
		temp := strings.Split(fileScanner.Text(), " ")

		if len(temp) > 2 {
			if temp[0] == "type" && temp[2] == "struct" {
				listStruct = append(listStruct, temp[1])

				var tempAttributesList []string

				for fileScanner.Scan() {
					attributesSplit := strings.Split(fileScanner.Text(), " ")

					if len(attributesSplit) > 0 {
						if attributesSplit[0] == "}" {
							break
						}
						tempAttributesList = append(tempAttributesList, attributesSplit[0])
					}
				}
				listAttributes[temp[1]] = tempAttributesList
			}
		}
	}

	return listStruct, listAttributes
}
