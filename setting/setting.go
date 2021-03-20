package setting

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

//YamlSettings Yaml Setting
type YamlSettings struct {
	Database struct {
		Type    []string `yaml:"type"`
		Path    string   `yaml:"path"`
		Name    string   `yaml:"name"`
		Apply   bool     `yaml:"apply"`
		Setting struct {
			Path          string   `yaml:"path"`
			Name          string   `yaml:"name"`
			SingularTable bool     `yaml:"singularTable"`
			TablePrefix   string   `yaml:"tablePrefix"`
			LogLevel      []string `yaml:"logLevel"`
			SlowThreshold int      `yaml:"slowThreshold"`
			Apply         bool     `yaml:"apply"`
		} `yaml:"setting"`
	} `yaml:"database"`
	Service struct {
		From struct {
			Path   string   `yaml:"path"`
			Name   string   `yaml:"name"`
			Ignore []string `yaml:"ignore"`
		} `yaml:"from"`
		To struct {
			Path    string `yaml:"path"`
			Postfix string `yaml:"postfix"`
		} `yaml:"to"`
		Apply bool `yaml:"apply"`
	} `yaml:"service"`
	QueryTools struct {
		Path  string `yaml:"path"`
		Name  string `yaml:"name"`
		Apply bool   `yaml:"apply"`
	} `yaml:"queryTools"`
}

//ReadYamlConfig Read Yaml to Model
func ReadYamlConfig() YamlSettings {
	var settingsYaml YamlSettings
	body, err := ioutil.ReadFile("./gormCrud.yaml")
	if err != nil {
		log.Println("Please Run With Arguments 'init' if you didn't have the config file")
		log.Println("Please Check Your Yaml File, Name it 'gormCrud.yaml'")
		panic(err)
	}
	if err = yaml.Unmarshal(body, &settingsYaml); err != nil {
		panic(err)
	}
	return settingsYaml
}
