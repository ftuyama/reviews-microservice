package api

import (
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/ftuyama/reviews-microservice/reviews"
)

// ReviewMiddleware decorates a review service.
type ReviewMiddleware func(ReviewService) ReviewService

// LoggingReviewMiddleware logs method calls, parameters, results, and elapsed time for the review service.
func LoggingReviewMiddleware(logger log.Logger) ReviewMiddleware {
	return func(next ReviewService) ReviewService {
		return loggingReviewMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingReviewMiddleware struct {
	next   ReviewService
	logger log.Logger
}

func (mw loggingReviewMiddleware) CreateReview(review *reviews.Review) error {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "CreateReview",
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.CreateReview(review)
}

func (mw loggingReviewMiddleware) GetReviewsByCustomerId(customerId string) ([]reviews.Review, error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "GetReviewsByCustomerId",
			"customerId", customerId,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.GetReviewsByCustomerId(customerId)
}

func (mw loggingReviewMiddleware) GetReviewsByItemId(itemId string) ([]reviews.Review, error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "GetReviewsByItemId",
			"itemId", itemId,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.GetReviewsByItemId(itemId)
}

func (mw loggingReviewMiddleware) DeleteReview(id string) error {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "DeleteReview",
			"id", id,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.DeleteReview(id)
}
