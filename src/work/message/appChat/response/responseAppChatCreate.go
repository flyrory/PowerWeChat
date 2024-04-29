package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseAppChatCreate struct {
	response.ResponseWork

	ChatID string `json:"chatid"`
}
