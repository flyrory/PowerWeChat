package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseLinkCorpGetDepartmentList struct {
	response.ResponseWork

	DepartmentList []*power.HashMap `json:"department_list"`
}
