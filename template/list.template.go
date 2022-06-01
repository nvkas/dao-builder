package BuilderTemplate

import (
	"dao-builder/utils"
	"strings"
)

//查询模板
func ListTemplate(model string,fields []BuilderUtils.ModelFields) string {
	maxName := model[strings.LastIndexAny(model,".")+1 :]
	//miniName := BuilderUtils.IndexToLower(maxName)

	str := "/**"+"\r\n"+
	"	"+maxName+" List"+"\r\n"+
	"*/"+"\r\n"+
	"func "+maxName+"List(Current, PageSize int, querys []BuilderModel.QueryCondition, orders []BuilderModel.OrderCondition) ([]"+model+", int64, error) {"+"\r\n"+
	"	db := BuilderCore.GetNewSession().Model(&"+model+"{})"+"\r\n"+
	"	"+"\r\n"+
	"	//查询条件"+"\r\n"+
	"	for _, query := range querys {"+"\r\n"+
	"		arges := query.Arges"+"\r\n"+
	"		if query.JudgeFlag {"+"\r\n"+
	"			db.Where(query.Query, arges...)"+"\r\n"+
	"		} else {"+"\r\n"+
	"			if len(arges) > 0 {"+"\r\n"+
	"				db.Where(query.Query, arges...)"+"\r\n"+
	"			}"+"\r\n"+
	"		}"+"\r\n"+
	"	}"+"\r\n"+
	"	//排序信息"+"\r\n"+
	"	ordersStr := \"\""+"\r\n"+
	"	for _, order := range orders {"+"\r\n"+
	"		ordersStr += order.Field + \" \" + order.Condition + \" \""+"\r\n"+
	"	}"+"\r\n"+
	"	if ordersStr != \"\" {"+"\r\n"+
	"		db = db.Order(ordersStr)"+"\r\n"+
	"	}"+"\r\n"+
	""+"\r\n"+
	"结果集"+"\r\n"+
	"	var ("+"\r\n"+
	"		result []"+model+""+"\r\n"+
	"		Total  int64"+"\r\n"+
	"	)"+"\r\n"+
	"	if err := db.Count(&Total)."+"\r\n"+
	"		Offset((Current - 1) * PageSize)."+"\r\n"+
	"		Limit(PageSize)."+"\r\n"+
	"		Find(&result).Error; err != nil {"+"\r\n"+
	"		return result, Total, err"+"\r\n"+
	"	}"+"\r\n"+
	"	"+"\r\n"+
	"	return result, Total, nil"+"\r\n"+
	"}"+"\r\n"

	return str
}
