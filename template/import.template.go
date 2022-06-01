package BuilderTemplate

func ImportTemplate(modelPkgNames []string) string {
	str := "import ("+"\r\n"+
		"	\"dao-builder/core\""+"\r\n"

	for _,path := range modelPkgNames {
		str +=	"	\""+path+"\""+"\r\n"
	}
	str +=	")"+"\r\n"
	str += "\r\n"
	return str
}
