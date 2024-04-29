package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseTagGetCorpTagList struct {
	response.ResponseWork

	TagGroups []*CorpTagGroup `json:"tag_group"`
}
