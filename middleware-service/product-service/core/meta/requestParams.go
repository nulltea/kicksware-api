package meta

type RequestParams interface {
	Limit() int
	SetLimit(limit int)

	Offset() int
	SetOffset(offset int)

	SortBy() string
	SetSortBy(sortBy string)

	SortDirection() string
	SortDirectionNum() int
	SetSortDirection(direction string)
}