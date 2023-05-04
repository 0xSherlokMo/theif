package search

const (
	TAG = "[SEARCH]"
)

type Item = string

type Response struct {
	ID                     int
	ProductShapeTypeName   string
	ProductShapeTypeNameAr string
	ProductShapeIconURL    string
	ProductNameEn          string
	ProductNameAr          string
	CategoryURLEn          string
	CategoryURLAr          string
	Category               string
	MainImageURL           string
}

type SearchDriver interface {
	Search(query string) (Response, error)
	Identifier() string
}
