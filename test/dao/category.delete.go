/**
	Build By DaoBuilder. It is not recommended to modify it
*/
package dao

import (
	"dao-builder/core"
	"test/model"
)

/**
	Category Delete
*/
func CategoryDelete(ids []int64) error {
	if err := BuilderCore.GetNewSession().Where("id in (?)",ids).Delete(&Models.Category{}).Error; err != nil {
		return err
	}
	return nil
}
