package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseCorpGroupTransferSession struct {
	response.ResponseWork
	Userid     string `json:"userid"`
	SessionKey string `json:"session_key"`
}
