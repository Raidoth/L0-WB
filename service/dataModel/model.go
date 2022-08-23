package dataModel

import "time"

type Delivery_t struct {
	Name    string `json:"name,omitempty" validate:"required"`
	Phone   string `json:"phone,omitempty" validate:"required"`
	Zip     string `json:"zip,omitempty" validate:"required"`
	City    string `json:"city,omitempty" validate:"required"`
	Address string `json:"address,omitempty" validate:"required"`
	Region  string `json:"region,omitempty" validate:"required"`
	Email   string `json:"email,omitempty" validate:"required"`
}

type Payment_t struct {
	Transaction  string `json:"transaction,omitempty" validate:"required"`
	RequestId    string `json:"request_id,omitempty"`
	Currency     string `json:"currency,omitempty" validate:"required"`
	Provider     string `json:"provider,omitempty" validate:"required"`
	Amount       uint   `json:"amount,omitempty" validate:"required"`
	PaymentDT    uint   `json:"payment_dt,omitempty" validate:"required"`
	Bank         string `json:"bank,omitempty" validate:"required"`
	DeliveryCost uint   `json:"delivery_cost,omitempty" validate:"required"`
	GoodsTotal   uint   `json:"goods_total,omitempty" validate:"required"`
	CustomFee    int    `json:"custom_fee,omitempty"`
}

type Items_t struct {
	ChrtId      uint   `json:"chrt_id" validate:"required"`
	TrackNumber string `json:"track_number,omitempty" validate:"required"`
	Price       uint   `json:"price,omitempty" validate:"required"`
	Rid         string `json:"rid,omitempty" validate:"required"`
	Name        string `json:"name,omitempty" validate:"required" `
	Sale        uint   `json:"sale,omitempty" validate:"required"`
	Size        string `json:"size,omitempty" validate:"required"`
	TotalPrice  uint   `json:"total_price,omitempty" validate:"required"`
	NmId        uint   `json:"nm_id,omitempty" validate:"required"`
	Brand       string `json:"brand,omitempty" validate:"required"`
	Status      int    `json:"status,omitempty" validate:"required"`
}

type Order_t struct {
	OrderUid          string     `json:"order_uid" validate:"required"`
	TrackNumber       string     `json:"track_number,omitempty" validate:"required"`
	Entry             string     `json:"entry,omitempty" validate:"required"`
	Delivery          Delivery_t `json:"delivery,omitempty" validate:"required"`
	Payment           Payment_t  `json:"payment,omitempty" validate:"required"`
	Items             []Items_t  `json:"items,omitempty" validate:"required"`
	Locale            string     `json:"locale,omitempty" validate:"required"`
	InternalSignature string     `json:"internal_signature,omitempty"`
	CustomerId        string     `json:"customer_id,omitempty" validate:"required"`
	DeliveryService   string     `json:"delivery_service,omitempty" validate:"required"`
	ShardKey          string     `json:"shardkey,omitempty" validate:"required"`
	SmId              uint       `json:"sm_id,omitempty" validate:"required"`
	DateCreated       time.Time  `json:"date_created,omitempty" validate:"required"`
	OofShard          string     `json:"oof_shard,omitempty" validate:"required"`
}
