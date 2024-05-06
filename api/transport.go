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
	"github.com/ftuyama/reviews-microservice/reviews"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	ErrInvalidRequest = errors.New("Invalid request")
)

// MakeHTTPHandler mounts the endpoints into a REST-y HTTP handler.
func MakeHTTPHandler(e Endpoints, logger log.Logger, tracer stdopentracing.Tracer) *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	// GET /reviews/customer/{id}       GetReviewsByCustomerId
	// GET /reviews/item/{id}           GetReviewsByItemId
	// POST /reviews                    CreateReview
	// DELETE /reviews/{id}             DeleteReview

	r.Methods("GET").Path("/reviews/customer/{id}").Handler(httptransport.NewServer(
		e.GetReviewsByCustomerIdEndpoint,
		decodeGetRequest,
		encodeResponse,
		append(options, httptransport.ServerBefore(opentracing.FromHTTPRequest(tracer, "GET /reviews/customer/{id}", logger)))...,
	))
	r.Methods("GET").Path("/reviews/item/{id}").Handler(httptransport.NewServer(
		e.GetReviewsByItemIdEndpoint,
		decodeGetRequest,
		encodeResponse,
		append(options, httptransport.ServerBefore(opentracing.FromHTTPRequest(tracer, "GET /reviews/item/{id}", logger)))...,
	))
	r.Methods("POST").Path("/reviews").Handler(httptransport.NewServer(
		e.CreateReviewEndpoint,
		decodeReviewRequest,
		encodeResponse,
		append(options, httptransport.ServerBefore(opentracing.FromHTTPRequest(tracer, "POST /reviews", logger)))...,
	))
	r.Methods("DELETE").Path("/reviews/{id}").Handler(httptransport.NewServer(
		e.DeleteReviewEndpoint,
		decodeDeleteReviewRequest,
		encodeResponse,
		append(options, httptransport.ServerBefore(opentracing.FromHTTPRequest(tracer, "DELETE /reviews/{id}", logger)))...,
	))

	r.Handle("/metrics", promhttp.Handler())
	return r
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
