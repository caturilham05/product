package web

type WebResponse struct {
	Code   int         `json:"code" example:"400"`
	Status string      `json:"status" example:"BAD REQUEST | ERROR NOT FOUND | UNAUTHORIZED | CONFLICT"`
	Data   interface{} `json:"data"`
}
