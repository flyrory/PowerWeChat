package response

import "github.com/flyrory/PowerWeChat/v3/src/kernel/response"

type ResponseMediaUpload struct {
	response.ResponsePayment
	MediaId string `json:"media_id"`
}
