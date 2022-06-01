package BuilderCore

import (
	"dao-builder/template"
	"dao-builder/utils"
	"fmt"
	"reflect"
)

func (builder *builderRunner)buildModel(model interface{})  {
	typeOf := reflect.TypeOf(model)
	//fmt.Println("typeOf:",typeOf.PkgPath(),typeOf.Name(),typeOf.FieldAlign())
	// 取类型的元素
	elem := typeOf.Elem()
	//8 Models.Category test/model Category
	//fmt.Println("elem:",elem.FieldAlign(),elem.String(),elem.PkgPath(),elem.Name())

	point := "."
	buildPkg := "dao"
	headers := BuilderTemplate.HeadersTemplate(buildPkg)
	importer := BuilderTemplate.ImportTemplate([]string{elem.PkgPath()})

	modName := elem.String()
	structName := BuilderUtils.HumpToPoint(elem.Name(),point)
	fields:= BuilderUtils.GetFields(model)

	//增加
	create:= BuilderTemplate.CreateTemplate(modName,fields)
	createStr := headers + importer + create
	createSuccess := BuilderUtils.FileWrite(builder.BuildPath+"",structName+point+"create.go",createStr)
	fmt.Println(createSuccess)

	//删除
	delete:= BuilderTemplate.DeleteTemplate(modName,fields)
	deleteStr := headers + importer + delete
	deleteSuccess := BuilderUtils.FileWrite(builder.BuildPath+"",structName+point+"delete.go",deleteStr)
	fmt.Println(deleteSuccess)

	//修改
	update:= BuilderTemplate.UpdateTemplate(modName,fields)
	updateStr := headers + importer + update
	updateSuccess := BuilderUtils.FileWrite(builder.BuildPath+"",structName+point+"update.go",updateStr)
	fmt.Println(updateSuccess)

	//查询
	importer = BuilderTemplate.ImportTemplate([]string{"dao-builder/model",elem.PkgPath()})
	list:= BuilderTemplate.ListTemplate(modName,fields)
	listStr := headers + importer + list
	listSuccess := BuilderUtils.FileWrite(builder.BuildPath+"",structName+point+"list.go",listStr)
	fmt.Println(listSuccess)
}
