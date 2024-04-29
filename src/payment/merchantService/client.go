package merchantService

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	payment "github.com/flyrory/PowerWeChat/v3/src/payment/kernel"
	"github.com/flyrory/PowerWeChat/v3/src/payment/merchantService/request"
	"github.com/flyrory/PowerWeChat/v3/src/payment/merchantService/response"
	"net/http"
)

type Client struct {
	*payment.BaseClient
}

func NewClient(app *payment.ApplicationPaymentInterface) (*Client, error) {
	baseClient, err := payment.NewBaseClient(app)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 查询投诉单列表
// https://pay.weixin.qq.com/docs/merchant/apis/consumer-complaint/complaints/list-complaints-v2.html
func (comp *Client) Complaints(ctx context.Context, params *request.RequestComplaints) (*response.ResponseComplaints, error) {

	result := &response.ResponseComplaints{}

	data, _ := object.StructToStringMap(params)

	endpoint := "/v3/merchant-service/complaints-v2"
	_, err := comp.Request(ctx, endpoint, data, http.MethodGet, &object.HashMap{}, false, nil, result)

	return result, err
}

// 查询投诉单详情
// https://pay.weixin.qq.com/docs/merchant/apis/consumer-complaint/complaints/list-complaints-v2.html
func (comp *Client) ComplaintInfo(ctx context.Context, params *request.RequestComplaintInfo) (*response.ResponseComplaintInfo, error) {

	result := &response.ResponseComplaintInfo{}

	endpoint := "/v3/merchant-service/complaints-v2/" + params.ComplaintId
	_, err := comp.Request(ctx, endpoint, nil, http.MethodGet, &object.HashMap{}, false, nil, result)

	return result, err
}

// 查询投诉单详情
// https://pay.weixin.qq.com/docs/merchant/apis/consumer-complaint/complaints/list-complaints-v2.html
func (comp *Client) ComplaintHistory(ctx context.Context, params *request.RequestComplaintHistory) (*response.ResponseComplaintHistory, error) {

	result := &response.ResponseComplaintHistory{}

	endpoint := "/v3/merchant-service/complaints-v2/" + params.ComplaintId + "/negotiation-historys"
	_, err := comp.Request(ctx, endpoint, nil, http.MethodGet, &object.HashMap{}, false, nil, result)

	return result, err
}
