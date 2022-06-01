package BuilderTemplate

import (
	"dao-builder/utils"
	"strings"
)

//添加模板
func CreateTemplate(model string,fields []BuilderUtils.ModelFields) string {
	maxName := model[strings.LastIndexAny(model,".")+1 :]
	miniName := BuilderUtils.IndexToLower(maxName)

	str := "/**"+"\r\n"+
		"	"+maxName+" Create"+"\r\n"+
		"*/"+"\r\n"+
	"func "+maxName+"Create("+miniName+" "+model+") error {"+"\r\n"+
	//	"	var "+maxName+" "+model+""+"\r\n"
	//for _,field := range fields {
	//	str += "	"+maxName+"."+field.Name+" = "+miniName+"."+field.Name+""+"\r\n"
	//}
	"	if err := BuilderCore.GetNewSession().Model(&"+model+"{}).Create(&"+miniName+").Error; err != nil {"+"\r\n"+
		"		return err"+"\r\n"+
		"	}"+"\r\n"+
		"	return nil"+"\r\n"+
		"}"
	return str
}
