package request

type Shipping struct {
	ID           uint   `json:"id"`
	ShippingNo   string `json:"shipping_no"`
	CustomerNote string `json:"customer_note"`
	Status       string `json:"status"`
	ItemID       int32  `json:"item_id"`
}

type GetShippingResponse struct {
	Data interface{} `json:"data"`
}

type ShippingResponse struct {
	Message    string      `json:"string"`
	Data       interface{} `json:"data"`
	StatusCode int32       `json:"status_code"`
	Err        bool        `json:"error,omitempty"`
}

type CreateNewShippingRequest struct {
	Data Shipping `json:"data"`
}
