package BuilderTemplate

import (
	"dao-builder/utils"
	"strings"
)

//删除模板
func DeleteTemplate(model string,fields []BuilderUtils.ModelFields) string {
	maxName := model[strings.LastIndexAny(model,".")+1 :]

	//获取主键类型
	primarykeyType := ""
	for _,field := range fields {
		gorm := field.Tags["gorm"]
		for _,v := range gorm {
			if v == "primarykey" {
				primarykeyType = field.Types.String()
			}
		}
	}
	if primarykeyType == "" {
		primarykeyType = fields[0].Types.String()
	}

	str := "/**"+"\r\n"+
		"	"+maxName+" Delete"+"\r\n"+
	"*/"+"\r\n"+
	"func "+maxName+"Delete(ids []"+primarykeyType+") error {"+"\r\n"+
	"	if err := BuilderCore.GetNewSession().Model(&"+model+"{}).Where(\"id in (?)\",ids).Delete(&"+model+"{}).Error; err != nil {"+"\r\n"+
	"		return err"+"\r\n"+
	"	}"+"\r\n"+
	"	return nil"+"\r\n"+
	"}"+"\r\n"
	return str
}
