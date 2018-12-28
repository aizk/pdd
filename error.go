package pdd

import "encoding/json"

type Error struct {
	ErrorResponse struct{
		ErrorMsg string `json:"error_msg"`
		SubMsg string `json:"sub_msg"`
		SubCode int `json:"sub_code"`
		ErrorCode int `json:"error_code"`
		RequestId string `json:"request_id"`
	} `json:"error_response"`
}

func (e Error) Error() string {
	r, _ := json.Marshal(e)
	return string(r)
}