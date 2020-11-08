package main

import (
	"fmt"

	//Comment m
	"github.com/dave/jennifer/jen"
)

//TestingJennifer T
func TestingJennifer() {

	structName := "Students"

	f := jen.NewFile("main")
	f.Func().Id(structName  +"Create").Params().Block(
		jen.Qual("fmt", "Println").Call(jen.Lit("Hello, world")),
	)
	fmt.Printf("%#v", f)

	// file, err := os.Create("tester.go")

	// defer file.Close()

	// if err != nil {
	// 	panic(err)
	// }

	// consume, err := file.WriteString(fmt.Sprintf("%#v", f))

	// fmt.Println(consume)

	// if err != nil {
	// 	panic(err)
	// }

}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {

// 	fileHandle, _ := os.Open("tester.go")
// 	defer fileHandle.Close()
// 	fileScanner := bufio.NewScanner(fileHandle)

// 	var listStruct []string

// 	for fileScanner.Scan() {
// 		temp := strings.Split(fileScanner.Text(), " ")

// 		if len(temp) >= 2 {
// 			if temp[0] == "type" && temp[2] == "struct" {
// 				listStruct = append(listStruct, temp[1])
// 			}
// 		}
// 	}

// 	for _, val := range listStruct {
// 		fmt.Println(val)
// 	}
// }
