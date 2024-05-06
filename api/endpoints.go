package api

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/tracing/opentracing"
	"reviews/reviews"
	stdopentracing "github.com/opentracing/opentracing-go"
)

// Endpoints collects the endpoints that comprise the Service.
type Endpoints struct {
	GetReviewsByCustomerIdEndpoint endpoint.Endpoint
	GetReviewsByItemIdEndpoint     endpoint.Endpoint
	CreateReviewEndpoint           endpoint.Endpoint
	DeleteReviewEndpoint           endpoint.Endpoint
}

// MakeEndpoints returns an Endpoints structure, where each endpoint is
// backed by the given service.
func MakeEndpoints(s Service, tracer stdopentracing.Tracer) Endpoints {
	return Endpoints{
		GetReviewsByCustomerIdEndpoint: opentracing.TraceServer(tracer, "GET /reviews/customer/{id}")(MakeGetReviewsByCustomerIdEndpoint(s)),
		GetReviewsByItemIdEndpoint:     opentracing.TraceServer(tracer, "GET /reviews/item/{id}")(MakeGetReviewsByItemIdEndpoint(s)),
		CreateReviewEndpoint:           opentracing.TraceServer(tracer, "POST /reviews")(MakeCreateReviewEndpoint(s)),
		DeleteReviewEndpoint:           opentracing.TraceServer(tracer, "DELETE /reviews/{id}")(MakeDeleteReviewEndpoint(s)),
	}
}

// MakeGetReviewsByCustomerIdEndpoint returns an endpoint via the given service.
func MakeGetReviewsByCustomerIdEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetRequest)
		reviews, err := s.GetReviewsByCustomerId(req.ID)
		return reviewsResponse{Reviews: reviews}, err
	}
}

// MakeGetReviewsByItemIdEndpoint returns an endpoint via the given service.
func MakeGetReviewsByItemIdEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetRequest)
		reviews, err := s.GetReviewsByItemId(req.ID)
		return reviewsResponse{Reviews: reviews}, err
	}
}

// MakeCreateReviewEndpoint returns an endpoint via the given service.
func MakeCreateReviewEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// req := request.(CreateReviewRequest)
		// id, err := s.CreateReview(req.Review)
		// return postResponse{ID: id}, err
		return nil, nil
	}
}

// MakeDeleteReviewEndpoint returns an endpoint via the given service.
func MakeDeleteReviewEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteRequest)
		err = s.DeleteReview(req.ID)
		if err == nil {
			return statusResponse{Status: true}, err
		}
		return statusResponse{Status: false}, err
	}
}

type GetRequest struct {
	ID   string
	Attr string
}

type reviewsResponse struct {
	Reviews []reviews.Review `json:"reviews"`
}

type CreateReviewRequest struct {
	Review reviews.Review `json:"review"`
}

type DeleteRequest struct {
	ID string `json:"id"`
}

type statusResponse struct {
	Status bool `json:"status"`
}

type postResponse struct {
	ID string `json:"id"`
}
