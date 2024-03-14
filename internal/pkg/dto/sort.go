package dto

type SortOptions struct {
	Type  key
	Order key
}

const (
	TitleSort key = iota
	RatingSort
	NameSort
)

const (
	AscSort key = iota
	DescSort
)
