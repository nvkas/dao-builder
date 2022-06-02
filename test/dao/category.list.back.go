/**
	Build By DaoBuilder. It is not recommended to modify it

 */
package dao

import (
	"dao-builder/core"
	"dao-builder/model"
	"test/model"
)

/**
	Category List
 */
func CategoryListBack(Current, PageSize int, querys []BuilderModel.QueryCondition, orders []BuilderModel.OrderCondition) ([]Models.Category, int64, error) {
	db := BuilderCore.GetNewSession().Model(&Models.Category{})

	//查询条件
	for _, query := range querys {
		arges := query.Arges
		if query.JudgeFlag {
			db.Where(query.Query, arges...)
		} else {
			if len(arges) > 0 {
				db.Where(query.Query, arges...)
			}
		}
	}
	ordersStr := ""
	for _, order := range orders {
		ordersStr += order.Field + " " + order.Condition + " "
	}
	//排序信息
	if ordersStr != "" {
		db = db.Order(ordersStr)
	}

	var (
		result []Models.Category
		Total  int64
	)
	if err := db.Count(&Total).
		Offset((Current - 1) * PageSize).
		Limit(PageSize).
		Find(&result).Error; err != nil {
		return result, Total, err
	}

	return result, Total, nil
}
