package BuilderTemplate

import (
	"dao-builder/utils"
	"strings"
)

//修改模板
func UpdateTemplate(model string,fields []BuilderUtils.ModelFields) string {
	maxName := model[strings.LastIndexAny(model,".")+1 :]
	miniName := BuilderUtils.IndexToLower(maxName)

	//获取主键类型
	var primarykey BuilderUtils.ModelFields
	for _,field := range fields {
		gorm := field.Tags["gorm"]
		for _,v := range gorm {
			if v == "primarykey" {
				primarykey = field
			}
		}
	}
	if primarykey.Name == "" {
		primarykey = fields[0]
	}

	primarykeyName := primarykey.Name
	primarykeyNameLower := BuilderUtils.HumpToPoint(primarykey.Name,"_")

	str := "/**"+"\r\n"+
		"	"+maxName+" Update"+"\r\n"+
		"*/"+"\r\n"+
		"func "+maxName+"Update("+miniName+" "+model+") error {"+"\r\n"+
		"	if err := BuilderCore.GetNewSession().Model(&"+model+"{}).Where(\""+primarykeyNameLower+" = ?\","+miniName+"."+primarykeyName+").Updates("+miniName+").Error; err != nil {"+"\r\n"+
		"		return err"+"\r\n"+
		"	}"+"\r\n"+
		"	return nil"+"\r\n"+
		"}"
	return str
}
