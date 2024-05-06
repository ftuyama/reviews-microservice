package db

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"reviews/reviews"
)

// Database represents an interface for managing reviews.
type Database interface {
	Init() error
	CreateReview(*reviews.Review) error
	GetReviewsByCustomerId(string) ([]reviews.Review, error)
	GetReviewsByItemId(string) ([]reviews.Review, error)
	DeleteReview(string) error
}

var (
	database string
	//DefaultDb is the database set for the microservice
	DefaultDb Database
	//DBTypes is a map of DB interfaces that can be used for this service
	DBTypes = map[string]Database{}
	//ErrNoDatabaseFound error returnes when database interface does not exists in DBTypes
	ErrNoDatabaseFound = "No database with name %v registered"
	//ErrNoDatabaseSelected is returned when no database was designated in the flag or env
	ErrNoDatabaseSelected = errors.New("No DB selected")
)

func init() {
	// Initialize the review database based on configuration.
	flag.StringVar(&database, "database", os.Getenv("USER_DATABASE"), "Database to use, Mongodb or ...")
}

//Init inits the selected DB in DefaultDb
func Init() error {
	if database == "" {
		return ErrNoDatabaseSelected
	}
	return DefaultDb.Init()
}

//Register registers the database interface in the DBTypes
func Register(name string, db Database) {
	DBTypes[name] = db
}

// InitReview initializes the selected review database in DefaultDb.
func InitReview() error {
	if database == "" {
		return ErrNoDatabaseSelected
	}
	err := SetReview()
	if err != nil {
		return err
	}
	return DefaultDb.Init()
}

// SetReview sets the DefaultDb.
func SetReview() error {
	if v, ok := DBTypes[database]; ok {
		DefaultDb = v
		return nil
	}
	return fmt.Errorf(ErrNoDatabaseFound, database)
}

// CreateReview invokes DefaultDb method to create a review.
func CreateReview(r *reviews.Review) error {
	return DefaultDb.CreateReview(r)
}

// GetReviewsByCustomerId invokes DefaultDb method to get reviews by customer ID.
func GetReviewsByCustomerId(customerId string) ([]reviews.Review, error) {
	return DefaultDb.GetReviewsByCustomerId(customerId)
}

// GetReviewsByItemId invokes DefaultDb method to get reviews by item ID.
func GetReviewsByItemId(itemId string) ([]reviews.Review, error) {
	return DefaultDb.GetReviewsByItemId(itemId)
}

// DeleteReview invokes DefaultDb method to delete a review.
func DeleteReview(id string) error {
	return DefaultDb.DeleteReview(id)
}
