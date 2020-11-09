package gormgenerator

import (
	"bufio"
	"os"
	"strings"
)

//GetGoModName Get Go Mod Name
func GetGoModName() string {
	fileHandle, _ := os.Open("./go.mod")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	fileScanner.Scan()

	goModName := strings.Split(fileScanner.Text(), " ")[1]

	return goModName
}

//GetStructAndAttribute Get Structs and Its Attributes
func GetStructAndAttribute() ([]string, map[string][]string) {
	fileHandle, _ := os.Open("./gormGenerator/model.go")
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
