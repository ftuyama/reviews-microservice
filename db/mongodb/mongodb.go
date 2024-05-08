package mongodb

import (
	"errors"
	"flag"
	"net/url"
	"os"
	"time"

	"reviews/reviews"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	name     string
	password string
	host     string
	db       = "reviews"
	// ErrInvalidHexID represents an entity id that is not a valid bson ObjectID
	ErrInvalidHexID = errors.New("invalid Id Hex")
)

func init() {
	flag.StringVar(&name, "mongo-user", os.Getenv("MONGO_USER"), "Mongo user")
	flag.StringVar(&password, "mongo-password", os.Getenv("MONGO_PASS"), "Mongo password")
	flag.StringVar(&host, "mongo-host", os.Getenv("MONGO_HOST"), "Mongo host")
}

// Mongo meets the Database interface requirements
type Mongo struct {
	// Session is a MongoDB Session
	Session *mgo.Session
}

// Init MongoDB
func (m *Mongo) Init() error {
	u := getURL()
	var err error
	m.Session, err = mgo.DialWithTimeout(u.String(), time.Duration(5)*time.Second)
	if err != nil {
		return err
	}
	return m.EnsureIndexes()
}

// MongoReview is a wrapper for reviews
type MongoReview struct {
	reviews.Review `bson:",inline"`
	ID             bson.ObjectId `bson:"_id"`
}

// NewReview returns a new MongoReview
func NewReview() MongoReview {
	return MongoReview{}
}

// CreateReview inserts a review into MongoDB
func (m *Mongo) CreateReview(r *reviews.Review) error {
	s := m.Session.Copy()
	defer s.Close()
	id := bson.NewObjectId()
	mr := MongoReview{Review: *r, ID: id}
	c := s.DB("").C("reviews")
	_, err := c.UpsertId(mr.ID, mr)
	if err != nil {
		return err
	}
	r.ID = id.Hex()
	return nil
}

// GetReviews retrieves reviews
func (m *Mongo) GetReviews() ([]reviews.Review, error) {
	s := m.Session.Copy()
	defer s.Close()
	var mrs []MongoReview
	c := s.DB("").C("reviews")
	err := c.Find(nil).All(&mrs)
	if err != nil {
		return nil, err
	}
	reviews := make([]reviews.Review, len(mrs))
	for i, mr := range mrs {
		reviews[i] = mr.Review
	}
	return reviews, nil
}

// GetReviewsByItemIdCustomerId retrieves reviews by item Id and customer Id
func (m *Mongo) GetReviewsByItemIdCustomerId(itemId, customerId string) ([]reviews.Review, error) {
	s := m.Session.Copy()
	defer s.Close()
	var mrs []MongoReview
	c := s.DB("").C("reviews")
	err := c.Find(bson.M{"itemId": itemId, "customerId": customerId}).All(&mrs)
	if err != nil {
		return nil, err
	}
	reviews := make([]reviews.Review, len(mrs))
	for i, mr := range mrs {
		reviews[i] = mr.Review
	}
	return reviews, nil
}

// GetReviewsByItemId retrieves reviews by item ID
func (m *Mongo) GetReviewsByItemId(itemId string) ([]reviews.Review, error) {
	s := m.Session.Copy()
	defer s.Close()
	var mrs []MongoReview
	c := s.DB("").C("reviews")
	err := c.Find(bson.M{"itemId": itemId}).All(&mrs)
	if err != nil {
		return nil, err
	}
	reviews := make([]reviews.Review, len(mrs))
	for i, mr := range mrs {
		reviews[i] = mr.Review
	}
	return reviews, nil
}

// DeleteReview deletes a review from MongoDB
func (m *Mongo) DeleteReview(id string) error {
	s := m.Session.Copy()
	defer s.Close()
	if !bson.IsObjectIdHex(id) {
		return ErrInvalidHexID
	}
	c := s.DB("").C("reviews")
	return c.RemoveId(bson.ObjectIdHex(id))
}

func getURL() *url.URL {
	ur := url.URL{
		Scheme: "mongodb",
		Host:   host,
		Path:   db,
	}
	if name != "" {
		u := url.UserPassword(name, password)
		ur.User = u
	}
	return &ur
}

// EnsureIndexes ensures username is unique
func (m *Mongo) EnsureIndexes() error {
	s := m.Session.Copy()
	defer s.Close()
	i := mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     false,
	}
	c := s.DB("").C("customers")
	return c.EnsureIndex(i)
}

func (m *Mongo) Ping() error {
	s := m.Session.Copy()
	defer s.Close()
	return s.Ping()
}
