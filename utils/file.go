package BuilderUtils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// 命令行执行目录
func FileBaseWdPath() string{
	str, _ := os.Getwd()
	return str
}

//读取文件
func FileReadAll(fileName string)  (string, error)  {
	strByte,err :=ioutil.ReadFile(fileName)
	return string(strByte),err
}
// 判断所给路径文件/文件夹是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//调用os.MkdirAll递归创建文件夹
func FileMkdirs(filePath string)  error  {
	if !FileExists(filePath) {
		err := os.MkdirAll(filePath,os.ModePerm)
		return err
	}
	return nil
}

func FileWrite(path, name, str string) bool {
	if FileIsDir(path) {
		fullFileName := filepath.Join(path, name)
		if FileExists(fullFileName) {
			return false
		}
		f, _ := os.Create(fullFileName)
		defer f.Close()
		f.Write([]byte(str))
		return true
	}else{
		return false
	}
}

// 判断所给路径是否为文件夹
func FileIsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}