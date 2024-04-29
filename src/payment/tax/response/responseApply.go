package response

import "github.com/flyrory/PowerWeChat/v3/src/kernel/response"

type ResponseApplyForCardTemplate struct {
	response.ResponsePayment

	CardAppid string `json:"card_appid"`
	CardId    string `json:"card_id"`
}
