package reviews

import (
	"time"
)

type Review struct {
	ID         string    `json:"id,omitempty" bson:"id,omitempty"`
	CustomerId string    `json:"customerId" bson:"customerId"`
	ItemId     string    `json:"itemId" bson:"itemId"`
	CreatedAt  time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	Rating     int       `json:"rating" bson:"rating"`
	Comment    string    `json:"comment" bson:"comment"`
}
