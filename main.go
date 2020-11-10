package gormcrudgenerator

import (
	"github.com/davidyap2002/GormCrudGenerator/gormgenerator"
)

// func main(){
// 	GenerateService()
// }

//GenerateService Generate
func GenerateService() {
	// gormgenerator.GormGenerator()
	goModName := gormgenerator.GetGoModName()

	listStruct, attributesList := gormgenerator.GetStructAndAttribute()

	for _, val := range listStruct {
		err := gormgenerator.GormGenerator(goModName, val, attributesList)

		if err != nil {
			panic(err)
		}
	}

}

// //GenerateCRUD g
// func GenerateCRUD(item interface{}) string {

// 	origin := reflect.ValueOf(item)

// 	// fmt.Println(reflect.TypeOf(item))
// 	// fmt.Printf("%s\n", origin.Type())

// 	elem := origin.Elem()

// 	temp := strings.Split(fmt.Sprintf("%s", origin.Type()), ".")

// 	structName := temp[len(temp)-1]
// 	lowerStruct := strcase.ToLowerCamel(structName)
// 	dbName := strcase.ToSnake(structName)

// 	var attributes []string

// 	attributeLen := elem.NumField()

// 	for i := 0; i < attributeLen; i++ {
// 		attributes = append(attributes, elem.Type().Field(i).Name)
// 	}

// 	generated := ""

// 	create := `

// ` +
// 		structName + `Create(ctx context.Context, input model.New` + structName + `) (*model.` + structName + `, error)` +
// 		`			db := config.ConnectGorm()
// 	defer db.Close()

// 	` +
// 		lowerStruct + ` := model.` + structName + `{` + "\n"

// 	flag := true
// 	for _, val := range attributes {
// 		if flag {
// 			flag = false
// 			continue
// 		}
// 		create += `		` + val + `: ` + `input.` + val + ",\n"
// 	}

// 	create += `	}

// 	err := db.Table("` + dbName + `").Create(&` + lowerStruct + `).Error

// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}

// 	return &` + lowerStruct + `, nil
// }`

// 	update := structName + `Update(ctx context.Context, input model.Update` + structName + `) (*model.` + structName + `, error)` +
// 		`			db := config.ConnectGorm()
// 	defer db.Close()

// 	` +
// 		lowerStruct + ` := model.` + structName + `{` + "\n"

// 	for _, val := range attributes {
// 		update += `		` + val + `: ` + `input.` + val + ",\n"
// 	}

// 	update += `	}

// 	err := db.Table("` + dbName + `").Where("id = ?", input.ID).Updates(&` + lowerStruct + `).Error

// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}

// 	return &` + lowerStruct + `, nil
// }`

// 	delete := structName + `Delete(ctx context.Context, id int) (string, error)
// 	db := config.ConnectGorm()
// 	defer db.Close()

// 	err := db.Table("` + dbName + `").Where("id = ?", id).Delete(&model.` + structName + `{}).Error

// 	if err != nil {
// 		fmt.Println(err)
// 		return "Fail", err
// 	}

// 	return "Success", nil
// }`

// 	readOne := structName + `GetByID(ctx context.Context, id int) (*model.` + structName + `, error)` +
// 		`
// 	db := config.ConnectGorm()
// 	defer db.Close()

// 	var ` + lowerStruct + ` model.` + structName + `

// 	err := db.Table("` + dbName + `").Where("id = ?", id).First(&` + lowerStruct + `).Error

// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}

// 	return &` + lowerStruct + `, nil
// }`

// 	readAll := structName + `GetAll(ctx context.Context) ([]*model.` + structName + `, error)` +
// 		`
// 	db := config.ConnectGorm()
// 	defer db.Close()

// 	var ` + lowerStruct + ` []*model.` + structName + `

// 	err := db.Table("` + dbName + `").Find(&` + lowerStruct + `).Error

// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}

// 	return ` + lowerStruct + `, nil
// }`

// 	generated += create + "\n\n" + update + "\n\n" + delete + "\n\n" + readOne + "\n\n" + readAll

// 	return generated
// }
