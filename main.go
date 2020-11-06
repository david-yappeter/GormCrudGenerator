package main

import (
	"fmt"
	"reflect"
)

// type Student struct {
// 	Fname  string `gorm:"type:varchar(255);not null"`
// 	Lname  string `gorm:"type:varchar(255);not null"`
// 	City   string `gorm:"type:varchar(255);not null"`
// 	Mobile int    `gorm:"type:int;not null"`
// }

// func main() {
// 	// s := Student{}
// 	// v := reflect.ValueOf(s)
// 	// typeOfS := v.Type()

// 	// for i := 0; i < v.NumField(); i++ {
// 	// 	fmt.Printf("Field: %s\tType: %s\tValue: %v\n", typeOfS.Field(i).Name, typeOfS.Field(i).Type, v.Field(i).Interface())
// 	// }

// 	// fmt.Println(typeOfS.Field(0).Anonymous)
// 	// fmt.Println(typeOfS.Field(0).Index)
// 	// fmt.Println(typeOfS.Field(0).Offset)
// 	// fmt.Println(typeOfS.Field(0).PkgPath)
// 	// fmt.Println(typeOfS.Field(0).Tag)

// 	var temp interface{} = {
// 		Student{},
// 	}

// }

type Student struct {
	ID   string
	Name string
}

func main() {
	item := Student{ID: "10", Name: "david"}
	doSomethinWithThisParam(&item)
	fmt.Printf("%+v", item)
}

// func doSomethinWithThisParam(item *interface{}) {
// 	*item = &Student{
// 		ID:   "124",
// 		Name: "Iman Tumorang",
// 	}
// }

func doSomethinWithThisParam(item interface{}) {
	origin := reflect.ValueOf(item)

	tempLen := origin.Elem()

	fmt.Println(tempLen.Type().Field(0).Name)
	// for i := 0; i < tempLen; i++ {
	// 	fmt.Println(origin.Field(i).Interface())
	// }
	// origin = &Student{
	// 	ID:   "124",
	// 	Name: "Iman Tumorang",
	// }
	// item = origin

	fmt.Println(origin)
}
