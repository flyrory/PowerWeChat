package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseTagDetail struct {
	response.ResponseWork

	TagName   string              `json:"tagname"`
	UserList  []*UserSimpleDetail `json:"userlist"`
	PartyList []int               `json:"partylist"`
}
