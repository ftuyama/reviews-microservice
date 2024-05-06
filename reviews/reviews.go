package reviews

import (
	"time"
)

type Review struct {
	CustomerId string    `json:"customerId" bson:"customerId"`
	ItemId     string    `json:"itemId" bson:"itemId"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt" bson:"updatedAt"`
	Rating     int       `json:"rating" bson:"rating"`
	Comment    string    `json:"comment" bson:"comment"`
}
