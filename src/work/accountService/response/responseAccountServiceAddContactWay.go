package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseAccountServiceAddContactWay struct {
	response.ResponseWork

	URL string `json:"url"`
}
