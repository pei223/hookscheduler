package types

type SortOrder string

const (
	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
)

type SortParams struct {
	SortBy    string
	SortOrder SortOrder
}

type ListParams struct {
	Limit  int
	Offset int
	Sort   *SortParams
}
