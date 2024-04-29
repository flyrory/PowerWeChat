package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseUserIDToOpenID struct {
	response.ResponseWork

	OpenID string `json:"openid"`
}
