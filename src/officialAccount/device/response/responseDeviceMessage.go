package response

import "github.com/flyrory/PowerWeChat/v3/src/kernel/response"

type ResponseDeviceMessage struct {
	response.ResponseOfficialAccount

	Ret     int    `json:"ret"`
	RetInfo string `json:"ret_info"`
}
