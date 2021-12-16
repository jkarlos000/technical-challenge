package entity

import "time"

// Beer represents a beer record.
type Beer struct {
	ID        int       `json:"id" db:"beer_id"`
	Name      string    `json:"name" db:"name"`
	Brewery   string    `json:"brewery" db:"brewery"`
	Country   string    `json:"country" db:"country"`
	Price     float32   `json:"price" db:"price"`
	Currency  string    `json:"currency" db:"currency"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

// TableName represents the table name
func (b *Beer) TableName() string {
	return "beers"
}
