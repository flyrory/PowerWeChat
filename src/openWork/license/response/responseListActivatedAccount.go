package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
	"github.com/flyrory/PowerWeChat/v3/src/openWork/license/model"
)

type ResponseListActivatedAccount struct {
	response.ResponseWork
	NextCursor  string             `json:"next_cursor,omitempty"`
	HasMore     int                `json:"has_more,omitempty"`
	AccountList []model.ActiveInfo `json:"account_list,omitempty"`
}
