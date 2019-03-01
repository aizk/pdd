package pdd

import (
	"encoding/json"
)

type PIdQueryListResponse struct {
	PIDList    []*PIdQuery `json:"p_id_list"`
	TotalCount int         `json:"total_count"`
}

type PIdQuery struct {
	PId        string `json:"p_id"`
	PName      string `json:"pid_name"`
	CreateTime int64  `json:"create_time"`
}

// 分页查询生成的推广位信息
func (d *DDK) GoodsPidQuery(notMustparams ...Params) (res *PIdQueryListResponse, err error) {
	params := NewParamsWithType(DDK_GoodsPidQuery, notMustparams...)

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "p_id_query_response")
	res = new(PIdQueryListResponse)
	err = json.Unmarshal(bytes, res)
	return
}

type PIdGenerateListResponse struct {
	PIDList []*PIdGenerate `json:"p_id_list"`
}

type PIdGenerate struct {
	PId        string `json:"p_id"`
	PIdName    string `json:"pid_name"`
	CreateTime int64  `json:"create_time"`
}

// 创建推广位
func (d *DDK) GoodsPidGenerate(number int, notMustparams ...Params) (res *PIdGenerateListResponse, err error) {
	params := NewParamsWithType(DDK_GoodsPidGenerate, notMustparams...)
	params.Set("number", number)

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "p_id_generate_response")
	res = new(PIdGenerateListResponse)
	err = json.Unmarshal(bytes, res)
	return
}
