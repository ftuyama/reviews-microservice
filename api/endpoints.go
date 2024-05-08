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
	GetReviewsEndpoint 						 				endpoint.Endpoint
	GetReviewsByItemIdCustomerIdEndpoint  endpoint.Endpoint
	GetReviewsByItemIdEndpoint     				endpoint.Endpoint
	CreateReviewEndpoint           				endpoint.Endpoint
	DeleteReviewEndpoint           				endpoint.Endpoint
}

// MakeEndpoints returns an Endpoints structure, where each endpoint is
// backed by the given service.
func MakeEndpoints(s Service, tracer stdopentracing.Tracer) Endpoints {
	return Endpoints{
		GetReviewsEndpoint: 									opentracing.TraceServer(tracer, "GET /reviews")(MakeGetReviewsEndpoint(s)),
		GetReviewsByItemIdCustomerIdEndpoint: opentracing.TraceServer(tracer, "GET /reviews/item/{item_id}/customer/{customer_id}")(MakeGetReviewsByItemIdCustomerIdEndpoint(s)),
		GetReviewsByItemIdEndpoint:     			opentracing.TraceServer(tracer, "GET /reviews/item/{item_id}")(MakeGetReviewsByItemIdEndpoint(s)),
		CreateReviewEndpoint:           			opentracing.TraceServer(tracer, "POST /reviews")(MakeCreateReviewEndpoint(s)),
		DeleteReviewEndpoint:           			opentracing.TraceServer(tracer, "DELETE /reviews/{id}")(MakeDeleteReviewEndpoint(s)),
	}
}

// MakeGetReviewsEndpoint returns an endpoint via the given service.
func MakeGetReviewsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		reviews, err := s.GetReviews()
		return reviewsResponse{Reviews: reviews}, err
	}
}

// MakeGetReviewsByItemIdCustomerIdEndpoint returns an endpoint via the given service.
func MakeGetReviewsByItemIdCustomerIdEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetByItemIdCustomerIdRequest)
		reviews, err := s.GetReviewsByItemIdCustomerId(req.ItemId, req.CustomerId)
		return reviewsResponse{Reviews: reviews}, err
	}
}

// MakeGetReviewsByItemIdEndpoint returns an endpoint via the given service.
func MakeGetReviewsByItemIdEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetByItemIdRequest)
		reviews, err := s.GetReviewsByItemId(req.ItemId)
		return reviewsResponse{Reviews: reviews}, err
	}
}

// MakeCreateReviewEndpoint returns an endpoint via the given service.
func MakeCreateReviewEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateReviewRequest)
		review, err := s.CreateReview(&req.Review)
		return review, err
	}
}

// MakeDeleteReviewEndpoint returns an endpoint via the given service.
func MakeDeleteReviewEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteRequest)
		err = s.DeleteReview(req.Id)
		if err == nil {
			return statusResponse{Status: true}, err
		}
		return statusResponse{Status: false}, err
	}
}

type GetRequest struct {
	Attr   string
}

type GetByItemIdRequest struct {
	ItemId string
	Attr   string
}

type GetByItemIdCustomerIdRequest struct {
	ItemId 		 string
	CustomerId string
	Attr   		 string
}

type reviewsResponse struct {
	Reviews []reviews.Review `json:"reviews"`
}

type CreateReviewRequest struct {
	Review reviews.Review `json:"review"`
}

type DeleteRequest struct {
	Id string `json:"id"`
}

type statusResponse struct {
	Status bool `json:"status"`
}

type postResponse struct {
	Id string `json:"id"`
}
