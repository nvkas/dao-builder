/**
	Build By DaoBuilder. It is not recommended to modify it

 */
package dao

import (
	"dao-builder/core"
	"test/model"
)

/**
	Category Create
 */
func CategoryCreateBack(category Models.Category) error {
	if err := BuilderCore.GetNewSession().Model(&Models.Category{}).Create(&category).Error; err != nil {
		return err
	}
	return nil
}
