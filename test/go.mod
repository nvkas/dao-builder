module test

go 1.16

require (
	gorm.io/driver/mysql v1.1.2
	gorm.io/gorm v1.21.15
	dao-builder v0.0.1
)

replace dao-builder => ../../dao-builder
