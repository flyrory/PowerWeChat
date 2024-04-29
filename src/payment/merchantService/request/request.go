package request

type RequestComplaints struct {
	Limit            string `json:"limit"`
	Offset           string `json:"offset"`
	BeginDate        string `json:"begin_date"`
	EndDate          string `json:"end_date"`
	ComplaintedMchId string `json:"complainted_mchid"`
}

type RequestComplaintInfo struct {
	ComplaintId string `json:"complaint_id"`
}

type RequestComplaintHistory struct {
	ComplaintId string `json:"complaint_id"`
}
