package response

import "github.com/flyrory/PowerWeChat/v3/src/kernel/response"

type ResponseAddGroupWelcomeTemplate struct {
	response.ResponseWork

	TemplateID string `json:"template_id"`
}
