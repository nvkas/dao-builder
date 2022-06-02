/**
	Build By DaoBuilder. It is not recommended to modify it
*/
package dao

import (
	"dao-builder/core"
	"test/model"
)

/**
	category Create
*/
func CategoryCreate(category Models.Category) error {
	var Category Models.Category
	Category.ID = category.ID
	Category.CreatedAt = category.CreatedAt
	Category.UpdatedAt = category.UpdatedAt
	Category.DeletedAt = category.DeletedAt
	Category.CreatedBy = category.CreatedBy
	Category.UpdatedBy = category.UpdatedBy
	Category.Name = category.Name
	Category.Image = category.Image
	Category.SortNo = category.SortNo
	Category.Enable = category.Enable
	if err := BuilderCore.GetNewSession().Create(&Category).Error; err != nil {
		return err
	}
	return nil
}
