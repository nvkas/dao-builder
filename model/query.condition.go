package BuilderModel

//查询条件
type QueryCondition struct {
	//是否判断(例如: = ?  if(XXX){ = ?})
	JudgeFlag bool `json:"judge_flag"`
	//字段
	Query string `json:"query"`
	//值
	Arges []interface{} `json:"arges"`
}

//排序条件
type OrderCondition struct {
	//字段
	Field string `json:"field"`
	//值(desc asc)
	Condition string `json:"condition"`
}