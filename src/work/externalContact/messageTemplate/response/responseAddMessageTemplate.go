package response

import "github.com/flyrory/PowerWeChat/v3/src/kernel/response"

type ResponseAddMessageTemplate struct {
	response.ResponseWork
	FailList []string `json:"fail_list"`
	MsgID    string   `json:"msgid"`
}
