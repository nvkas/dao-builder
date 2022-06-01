package BuilderUtils

import (
	"encoding/json"
	"strings"
)

func InterFaceToMap(data interface{}) map[string]interface{} {
	bytes,_ := json.Marshal(data)
	maps := make(map[string]interface{})
	json.Unmarshal(bytes, &maps)
	return maps
}

//首字母大写
func IndexToUpper(str string) string {
	result:= ""
	strArr := strings.Split(str,"")

	for index,s := range strArr {
		if index == 0 {
			result += strings.ToUpper(s)
		}else {
			result += s
		}
	}
	return result
}

//首字母小写
func IndexToLower(str string) string {
	result:= ""
	strArr := strings.Split(str,"")

	for index,s := range strArr {
		if index == 0 {
			result += strings.ToLower(s)
		}else {
			result += s
		}
	}
	return result
}

//驼峰转自定义符号
func HumpToPoint(str,point string) string {

	str = strings.Replace(str,"ID","id",-1)

	result:= ""
	strArr := strings.Split(str,"")

	for index,s := range strArr {
		upperStr := strings.ToUpper(s)
		lowerStr := strings.ToLower(s)
		if index == 0 {
			result += lowerStr
		}else {
			str := lowerStr
			if s == upperStr {
				str = point+lowerStr
			}
			result += str
		}
	}
	return result
}