package entity

import "time"

type UserProvider struct {
	CustomerID string    `json:"customer_id" db:"customer_id"`
	Token      string    `json:"token" db:"token"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
	CreatedBy  string    `json:"created_by" db:"created_by"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
	UpdatedBy  string    `json:"updated_by" db:"updated_by"`
}
