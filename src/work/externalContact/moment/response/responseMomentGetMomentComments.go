package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseMomentGetMomentComments struct {
	response.ResponseWork

	CommentList []*power.HashMap `json:"comment_list"`
	LikeList    []*power.HashMap `json:"like_list"`
	NextCursor  string           `json:"next_cursor"`
}
