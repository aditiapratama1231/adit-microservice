package models

type Shipping struct {
	Model
	ShippingNo string `gorm:"shipping_no" json:"shipping_no"` 
	CustomerNote string `gorm:"customer_note" json:"customer_note"`
	Status string `gorm:"status" json:"status"`
	ItemID int32 `gorm:"item_id" json:"item_id"`
}