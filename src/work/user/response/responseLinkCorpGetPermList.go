package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseLinkCorpGetPermList struct {
	response.ResponseWork

	UserIDs       []string `json:"userids"`
	DepartmentIDs []string `json:"department_ids"`
}
