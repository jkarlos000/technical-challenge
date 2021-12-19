package entity

import "time"

// Beer represents a beer record.
//
// swagger:model
type Beer struct {
	// required: true
	// example: 1
	ID        int       `json:"id" db:"beer_id"`

	// required: true
	// example: Golden
	Name      string    `json:"name" db:"name"`

	// required: true
	// example: Kross
	Brewery   string    `json:"brewery" db:"brewery"`

	// required: true
	// example: Chile
	Country   string    `json:"country" db:"country"`

	// required: true
	// example: 10.5
	Price     float32   `json:"price" db:"price"`

	// required: true
	// example: EUR
	Currency  string    `json:"currency" db:"currency"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt *time.Time `json:"-" db:"updated_at"`
}

// TableName represents the table name
func (b *Beer) TableName() string {
	return "beers"
}
