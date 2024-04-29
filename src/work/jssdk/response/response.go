package response

import "github.com/flyrory/PowerWeChat/v3/src/kernel/response"

type ResponseGetTicket struct {
	response.ResponseWork

	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}
