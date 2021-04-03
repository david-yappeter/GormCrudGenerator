package cmd

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func argsCheck(arguments []string) bool {
	if len(arguments) > 0 {
		if strings.EqualFold(arguments[0], "init") {
			return yamlInit()
		} else {
			panic("Not A Valid Arguments")
		}
	}

	return false
}

func yamlInit() bool {

	var d yaml.Node
	err := yaml.Unmarshal([]byte(defaultConfig), &d)
	if err != nil {
		panic(err)
	}

	settingOutput, err := yaml.Marshal(&d)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("./gormCrud.yaml")

	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(string(settingOutput))

	if err != nil {
		panic(err)
	}

	return true
}
