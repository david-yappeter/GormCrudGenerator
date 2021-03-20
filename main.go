package main

import (
	"os"

	"github.com/david-yappeter/GormCrudGenerator/cmd"
)

func main() {
	arguments := os.Args[1:]
	cmd.GenerateService(arguments)
}
