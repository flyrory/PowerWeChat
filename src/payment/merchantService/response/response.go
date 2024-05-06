package response

import "time"

type ReturnAddressInfo struct {
	ReturnAddress string `json:"return_address"`
	Longitude     string `json:"longitude"`
	Latitude      string `json:"latitude"`
}

type SharePowerInfo struct {
	ReturnTime        string            `json:"return_time"`
	ReturnAddressInfo ReturnAddressInfo `json:"return_address_info"`
}

type AdditionalInfo struct {
	Type           string         `json:"type"`
	SharePowerInfo SharePowerInfo `json:"share_power_info"`
}

type ServiceOrderInfo struct {
	OrderId    string `json:"order_id"`
	OutOrderNo string `json:"out_order_no"`
	State      string `json:"state"`
}

type ComplaintMediaList struct {
	MediaType string   `json:"media_type"`
	MediaUrl  []string `json:"media_url"`
}

type ComplainOrderInfo struct {
	TransactionId string `json:"transaction_id"`
	OutTradeNo    string `json:"out_trade_no"`
	Amount        int    `json:"amount"`
}

type Data struct {
	ComplaintId           string               `json:"complaint_id"`
	ComplaintTime         time.Time            `json:"complaint_time"`
	ComplaintDetail       string               `json:"complaint_detail"`
	ComplaintState        string               `json:"complaint_state"`
	PayerPhone            string               `json:"payer_phone"`
	ComplaintOrderInfo    []ComplainOrderInfo  `json:"complaint_order_info"`
	ComplaintFullRefunded bool                 `json:"complaint_full_refunded"`
	IncomingUserResponse  bool                 `json:"incoming_user_response"`
	UserComplaintTimes    int                  `json:"user_complaint_times"`
	ComplaintMediaList    []ComplaintMediaList `json:"complaint_media_list"`
	ProblemDescription    string               `json:"problem_description"`
	ProblemType           string               `json:"problem_type"`
	ApplyRefundAmount     int                  `json:"apply_refund_amount"`
	UserTagList           []string             `json:"user_tag_list"`
	ServiceOrderInfo      []ServiceOrderInfo   `json:"service_order_info"`
	AdditionalInfo        AdditionalInfo       `json:"additional_info"`
}

type ResponseComplaints struct {
	Data       []Data `json:"data"`
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
	TotalCount int    `json:"total_count"`
}

type ResponseComplaintInfo struct {
	ComplaintId           string               `json:"complaint_id"`
	ComplaintTime         time.Time            `json:"complaint_time"`
	ComplaintDetail       string               `json:"complaint_detail"`
	ComplaintState        string               `json:"complaint_state"`
	PayerPhone            string               `json:"payer_phone"`
	ComplaintOrderInfo    []ComplainOrderInfo  `json:"complaint_order_info"`
	ComplaintFullRefunded bool                 `json:"complaint_full_refunded"`
	IncomingUserResponse  bool                 `json:"incoming_user_response"`
	UserComplaintTimes    int                  `json:"user_complaint_times"`
	ComplaintMediaList    []ComplaintMediaList `json:"complaint_media_list"`
	ProblemDescription    string               `json:"problem_description"`
	ProblemType           string               `json:"problem_type"`
	ApplyRefundAmount     int                  `json:"apply_refund_amount"`
	UserTagList           []string             `json:"user_tag_list"`
	ServiceOrderInfo      []ServiceOrderInfo   `json:"service_order_info"`
	AdditionalInfo        AdditionalInfo       `json:"additional_info"`
}

type ResponseComplaintHistoryData struct {
	LogId              string               `json:"log_id"`
	Operator           string               `json:"operator"`
	OperateTime        time.Time            `json:"operate_time"`
	OperateType        string               `json:"operate_type"`
	OperateDetails     string               `json:"operate_details"`
	ImageList          []string             `json:"image_list"`
	ComplaintMediaList []ComplaintMediaList `json:"complaint_media_list,omitempty"`
}

type ResponseComplaintHistory struct {
	Data       []ResponseComplaintHistoryData `json:"data"`
	Limit      int                            `json:"limit"`
	Offset     int                            `json:"offset"`
	TotalCount int                            `json:"total_count"`
}
