package entity

import "time"

// Currency represent a currency record.
type Currency struct {
	ID          int        `json:"id" db:"id"`
	Base        string     `json:"base" db:"base"`
	Destination string     `json:"destination" db:"destination"`
	Rate        float32    `json:"rate" db:"rate"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
}

// TableName represents the table name
func (b *Currency) TableName() string {
	return "currencies"
}
