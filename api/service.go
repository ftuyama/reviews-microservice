package api

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"

	"reviews/reviews"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
)

// Service is the reviews service, providing operations for creating, retrieving, and deleting reviews.
type Service interface {
	CreateReview(review *reviews.Review) (reviews.Review, error)
	GetReviewsByCustomerId(customerId string) ([]reviews.Review, error)
	GetReviewsByItemId(itemId string) ([]reviews.Review, error)
	DeleteReview(id string) error
}

// NewFixedService returns a simple implementation of the Service interface.
func NewFixedService() Service {
	return &fixedService{}
}

type fixedService struct{}

func (s *fixedService) CreateReview(review *reviews.Review) (reviews.Review, error) {
	// Your implementation to create a review in the database
	return reviews.Review{}, nil
}

func (s *fixedService) GetReviewsByCustomerId(customerId string) ([]reviews.Review, error) {
	// Your implementation to retrieve reviews by customer ID from the database
	return nil, nil
}

func (s *fixedService) GetReviewsByItemId(itemId string) ([]reviews.Review, error) {
	// Your implementation to retrieve reviews by item ID from the database
	return nil, nil
}

func (s *fixedService) DeleteReview(id string) error {
	// Your implementation to delete a review from the database
	return nil
}

func calculatePassHash(pass, salt string) string {
	h := sha1.New()
	io.WriteString(h, salt)
	io.WriteString(h, pass)
	return fmt.Sprintf("%x", h.Sum(nil))
}
