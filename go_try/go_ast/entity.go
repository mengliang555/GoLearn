package go_ast

import (
	"encoding/json"
)

type Param struct {
	ParamName string `json:"param_name"`
	ParamType string `json:"param_type"`
}

type FuncStruct struct {
	MethodName    string   `json:"method_name"`
	ReqParamList  []*Param `json:"req_param_list"`
	RespParamList []*Param `json:"resp_param_list"`
}

func (s *FuncStruct) Print() string {
	bb, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(bb)
}

