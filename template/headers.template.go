package BuilderTemplate

func HeadersTemplate(pkg string) string {
	str := "/**"+"\r\n"+
		"	Build By DaoBuilder. It is not recommended to modify it"+"\r\n"+
		"*/"+"\r\n"
	str += "package "+pkg+"\r\n"
	str += "\r\n"
	return str
}
