package api

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"

	"reviews/db"
	"reviews/reviews"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
)

// Service is the reviews service, providing operations for creating, retrieving, and deleting reviews.
type Service interface {
	CreateReview(review *reviews.Review) (reviews.Review, error)
	GetReviews() ([]reviews.Review, error)
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
	// review, err := db.CreateReview(review)
	return reviews.Review{}, nil
}

func (s *fixedService) GetReviews() ([]reviews.Review, error) {
	reviews, err := db.GetReviews()
	return reviews, err
}

func (s *fixedService) GetReviewsByCustomerId(customerId string) ([]reviews.Review, error) {
	reviews, err := db.GetReviewsByCustomerId(customerId)
	return reviews, err
}

func (s *fixedService) GetReviewsByItemId(itemId string) ([]reviews.Review, error) {
	reviews, err := db.GetReviewsByItemId(itemId)
	return reviews, err
}

func (s *fixedService) DeleteReview(id string) error {
	return db.DeleteReview(id)
}

func calculatePassHash(pass, salt string) string {
	h := sha1.New()
	io.WriteString(h, salt)
	io.WriteString(h, pass)
	return fmt.Sprintf("%x", h.Sum(nil))
}
