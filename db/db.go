package db

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/ftuyama/reviews-microservice/reviews"
)

// ReviewDatabase represents an interface for managing reviews.
type ReviewDatabase interface {
	CreateReview(*reviews.Review) error
	GetReviewsByCustomerId(string) ([]reviews.Review, error)
	GetReviewsByItemId(string) ([]reviews.Review, error)
	DeleteReview(string) error
}

var (
	// DefaultReviewDb is the review database set for the microservice.
	DefaultReviewDb ReviewDatabase
)

func init() {
	// Initialize the review database based on configuration.
	flag.StringVar(&database, "database", os.Getenv("USER_DATABASE"), "Database to use, Mongodb or ...")
}

// InitReview initializes the selected review database in DefaultReviewDb.
func InitReview() error {
	if database == "" {
		return ErrNoDatabaseSelected
	}
	err := SetReview()
	if err != nil {
		return err
	}
	return DefaultReviewDb.Init()
}

// SetReview sets the DefaultReviewDb.
func SetReview() error {
	if v, ok := DBTypes[database]; ok {
		DefaultReviewDb = v
		return nil
	}
	return fmt.Errorf(ErrNoDatabaseFound, database)
}

// CreateReview invokes DefaultReviewDb method to create a review.
func CreateReview(r *reviews.Review) error {
	return DefaultReviewDb.CreateReview(r)
}

// GetReviewsByCustomerId invokes DefaultReviewDb method to get reviews by customer ID.
func GetReviewsByCustomerId(customerId string) ([]reviews.Review, error) {
	return DefaultReviewDb.GetReviewsByCustomerId(customerId)
}

// GetReviewsByItemId invokes DefaultReviewDb method to get reviews by item ID.
func GetReviewsByItemId(itemId string) ([]reviews.Review, error) {
	return DefaultReviewDb.GetReviewsByItemId(itemId)
}

// DeleteReview invokes DefaultReviewDb method to delete a review.
func DeleteReview(id string) error {
	return DefaultReviewDb.DeleteReview(id)
}
