package types

type GeneralResponse struct {
	ResultCode  int         `json:"resultCode"`
	Description string      `json:"description"`
	ErrCode     int         `json:"errCode"`
	Result      interface{} `json:"result"`
}
