package BuilderCore

import (
	"dao-builder/utils"
	"gorm.io/gorm"
	"os"
)

//构建者
type builderRunner struct {
	//初始化完成标志
	InitFlag bool
	//Mod文件名称 例如:module test
	ModName string
	//项目路径
	RootPath string
	//文件生成目录
	BuildPath string
	//事务支持
	TransactionFlag bool
	//获取数据库连接
	NewSession func(...interface{}) *gorm.DB
}

var br = &builderRunner{}

func NewGormBuilderRunner(RootPath,BuildPath,ModName string,TransactionFlag bool,NewSession func(...interface{}) *gorm.DB) error {
	if RootPath == "" {
		RootPath = BuilderUtils.FileBaseWdPath()
	}
	if BuildPath == "" {
		BuildPath = RootPath+string(os.PathSeparator)+"dao"
		if !BuilderUtils.FileExists(BuildPath) {
			BuilderUtils.FileMkdirs(BuildPath)
		}
	}
	br.RootPath = RootPath
	br.BuildPath = BuildPath
	br.ModName = ModName
	br.TransactionFlag = TransactionFlag
	br.NewSession = NewSession
	br.InitFlag = true
	return nil
}

func GetBuilderRunner() *builderRunner {
	if !br.InitFlag {
		return nil
	}
	return br
}

func GetNewSession() *gorm.DB {
	return br.NewSession()
}

//生成文件
func (builder *builderRunner)AutoMigrate(models ...interface{})  {
	for _,model := range models {
		builder.buildModel(model)
	}
}

