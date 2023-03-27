package entity

import "time"

type Wallet struct {
	ID         int64     `json:"id" db:"id"`
	CustomerID string    `json:"customer_id" db:"customer_id"`
	IsEnable   bool      `json:"is_enable" db:"is_enable"`
	Balance    float64   `json:"balance" db:"balance"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
	CreatedBy  string    `json:"created_by" db:"created_by"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
	UpdatedBy  string    `json:"updated_by" db:"updated_by"`
}
