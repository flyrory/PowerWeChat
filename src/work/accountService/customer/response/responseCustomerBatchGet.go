package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseCustomerBatchGet struct {
	response.ResponseWork

	CustomerList          []*power.HashMap `json:"customer_list"`
	InvalidExternalUserID []*string        `json:"invalid_external_userid"`
}
