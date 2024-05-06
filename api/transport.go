package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"reviews/reviews"
	stdopentracing "github.com/opentracing/opentracing-go"
)

var (
	ErrInvalidRequest = errors.New("Invalid request")
)

// MakeHTTPHandler mounts the endpoints into a REST-y HTTP handler.
func MakeHTTPHandler(e Endpoints, logger log.Logger, tracer stdopentracing.Tracer) *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	// options := []httptransport.ServerOption{
	// 	httptransport.ServerErrorLogger(logger),
	// 	httptransport.ServerErrorEncoder(encodeError),
	// }

	// GET /reviews/       							GetReviews
	// GET /reviews/customer/{id}       GetReviewsByCustomerId
	// GET /reviews/item/{id}           GetReviewsByItemId
	// POST /reviews                    CreateReview
	// DELETE /reviews/{id}             DeleteReview

	r.Methods("GET").Path("/reviews").Handler(httptransport.NewServer(
		e.GetReviewsEndpoint,
		decodeGetRequest,
		encodeResponse,
	))
	r.Methods("GET").Path("/reviews/customer/{id}").Handler(httptransport.NewServer(
		e.GetReviewsByCustomerIdEndpoint,
		decodeGetRequest,
		encodeResponse,
	))
	r.Methods("GET").Path("/reviews/item/{id}").Handler(httptransport.NewServer(
		e.GetReviewsByItemIdEndpoint,
		decodeGetRequest,
		encodeResponse,
	))
	r.Methods("POST").Path("/reviews").Handler(httptransport.NewServer(
		e.CreateReviewEndpoint,
		decodeReviewRequest,
		encodeResponse,
	))
	r.Methods("DELETE").Path("/reviews/{id}").Handler(httptransport.NewServer(
		e.DeleteReviewEndpoint,
		decodeDeleteReviewRequest,
		encodeResponse,
	))
	return r
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	code := http.StatusInternalServerError
	switch err {
	case ErrUnauthorized:
		code = http.StatusUnauthorized
	}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/hal+json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error":       err.Error(),
		"status_code": code,
		"status_text": http.StatusText(code),
	})
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	// All of our response objects are JSON serializable, so we just do that.
	w.Header().Set("Content-Type", "application/hal+json")
	return json.NewEncoder(w).Encode(response)
}

func decodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	g := GetRequest{}
	u := strings.Split(r.URL.Path, "/")
	if len(u) > 4 {
		g.ID = u[4]
	}
	return g, nil
}

func decodeReviewRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	review := reviews.Review{}
	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		return nil, err
	}
	return review, nil
}

func decodeDeleteReviewRequest(_ context.Context, r *http.Request) (interface{}, error) {
	d := deleteRequest{}
	u := strings.Split(r.URL.Path, "/")
	if len(u) == 3 {
		d.ID = u[2]
		return d, nil
	}
	return d, ErrInvalidRequest
}

type deleteRequest struct {
	Entity string
	ID     string
}
