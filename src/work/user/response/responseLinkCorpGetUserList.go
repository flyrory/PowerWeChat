package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseLinkCorpGetUserList struct {
	response.ResponseWork

	UserList []*power.HashMap `json:"userlist"`
}
