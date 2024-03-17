package dto

type SortOptions struct {
	Type  int
	Order int
}

const (
	TitleSort int = iota
	RatingSort
	ReleaseSort
)

const (
	AscSort int = iota
	DescSort
)

var DefaultSort = map[int]int{
	TitleSort:   AscSort,
	RatingSort:  DescSort,
	ReleaseSort: AscSort,
}
