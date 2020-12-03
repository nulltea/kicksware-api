package service

type RatingService interface {
	IncrementViews(entity string) (int64, error)
	IncrementOrders(entity string) (int64, error)
	IncrementSearches(entity string) (int64, error)
	CalculateRating(entity string) (int64, error)
	RetrieveRating(entity string) (int64, error)
	UpdateRating(entity string) error
}