/**
	Build By DaoBuilder. It is not recommended to modify it
*/
package dao

import (
	"dao-builder/core"
	"test/model"
)

/**
	Category Update
*/
func CategoryUpdate(category Models.Category) error {
	if err := BuilderCore.GetNewSession().Model(&Models.Category{}).Where("id = ?",category.ID).Updates(category).Error; err != nil {
		return err
	}
	return nil
}