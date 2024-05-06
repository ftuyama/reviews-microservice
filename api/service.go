package api

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/ftuyama/reviews-microservice/reviews"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
)

// ReviewService is the reviews service, providing operations for creating, retrieving, and deleting reviews.
type ReviewService interface {
	CreateReview(review *reviews.Review) error
	GetReviewsByCustomerId(customerId string) ([]reviews.Review, error)
	GetReviewsByItemId(itemId string) ([]reviews.Review, error)
	DeleteReview(id string) error
}

// NewFixedReviewService returns a simple implementation of the ReviewService interface.
func NewFixedReviewService() ReviewService {
	return &fixedReviewService{}
}

type fixedReviewService struct{}

func (s *fixedReviewService) CreateReview(review *reviews.Review) error {
	// Your implementation to create a review in the database
	return nil
}

func (s *fixedReviewService) GetReviewsByCustomerId(customerId string) ([]reviews.Review, error) {
	// Your implementation to retrieve reviews by customer ID from the database
	return nil, nil
}

func (s *fixedReviewService) GetReviewsByItemId(itemId string) ([]reviews.Review, error) {
	// Your implementation to retrieve reviews by item ID from the database
	return nil, nil
}

func (s *fixedReviewService) DeleteReview(id string) error {
	// Your implementation to delete a review from the database
	return nil
}

func calculatePassHash(pass, salt string) string {
	h := sha1.New()
	io.WriteString(h, salt)
	io.WriteString(h, pass)
	return fmt.Sprintf("%x", h.Sum(nil))
}
