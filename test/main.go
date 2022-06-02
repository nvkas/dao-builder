package main

import (
	"dao-builder/core"
	"dao-builder/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"test/dao"
	"test/model"
)

func main() {
	defer func() {
		//恢复程序的控制权
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()

	fmt.Println("hello word!")

	_ = BuilderCore.NewGormBuilderRunner("", "","test",true,
		//此处获取db连接可以替换成自己框架的db连接或连接池连接
		func(i ...interface{}) *gorm.DB {

			dsn := fmt.Sprintf(
				"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
				"db_username",
				"db_password",
				"127.0.0.1",
				3306,
				"db_name",
			)
			//fmt.Println("mysql dsn:>", dsn )
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				panic(err)
			}
			return db
		})

	_ = BuilderCore.GetNewSession().AutoMigrate(
		&Models.Category{},
		)
	BuilderCore.GetBuilderRunner().AutoMigrate(
		&Models.Category{},
		)

	var Current, PageSize int = 1,10
	var querys []BuilderModel.QueryCondition
	querys = append(querys, BuilderModel.QueryCondition{false,"name like ?",[]interface{}{"%name%"}})
	var orders []BuilderModel.OrderCondition
	orders = append(orders, BuilderModel.OrderCondition{"id","desc"})
	result,total,err := dao.CategoryList(Current, PageSize, querys, orders)
	if err != nil {
		log.Println("ERROR: ",err)
	}
	fmt.Println("result:",result)
	fmt.Println("total:",total)
}
