package BuilderUtils

import (
	"fmt"
	"reflect"
	"strings"
)

type ModelFields struct {
	Name  string              `json:"name"`
	Tags  map[string][]string `json:"tags"`
	Types reflect.Type        `json:"types"`
}

func GetFields(model interface{}) []ModelFields {
	var result []ModelFields
	// 获取结构体实例的反射类型对象
	t := reflect.TypeOf(model).Elem()

	for i := 0; i < t.NumField(); i++ {
		//struct 字段
		child := t.Field(i).Type
		if child.Kind() == reflect.Struct {
			for j := 0; j < child.NumField(); j++ {
				jField := child.Field(j)
				jFieldName := jField.Name
				jFieldType := jField.Type

				//fmt.Println("name:", jFieldName)

				tagMap := make(map[string][]string)
				//builder := strings.Split(jField.Tag.Get("builder"), ";")
				//tagMap["builder"] = builder
				json := strings.Split(jField.Tag.Get("json"), ";")
				tagMap["json"] = json
				gorm := strings.Split(jField.Tag.Get("gorm"), ";")
				tagMap["gorm"] = gorm
				//fmt.Println("tagMap:",tagMap)
				result = append(result, ModelFields{jFieldName, tagMap, jFieldType})
			}
		} else {
			field := t.Field(i)
			fieldName := field.Name
			fieldType := field.Type

			//fmt.Println("name:", fieldName)

			tagMap := make(map[string][]string)
			//builder := strings.Split(field.Tag.Get("builder"), ";")
			//tagMap["builder"] = builder
			json := strings.Split(field.Tag.Get("json"), ";")
			tagMap["json"] = json
			gorm := strings.Split(field.Tag.Get("gorm"), ";")
			tagMap["gorm"] = gorm
			//fmt.Println("tagMap:",tagMap)
			result = append(result, ModelFields{fieldName, tagMap, fieldType})
		}
	}
	fmt.Println(result)
	return result
}
