package veseeta

import "github.com/graduation-fci/multivendor-scrapper/search"

type SearchResponse struct {
	From       int       `json:"from"`
	Size       int       `json:"size"`
	TotalCount int       `json:"totalCount"`
	Product    []Product `json:"productShapes"`
}

type Product struct {
	ID                     int    `json:"id"`
	ProductShapeTypeName   string `json:"productShapeTypeName"`
	ProductShapeTypeNameAr string `json:"productShapeTypeNameAr"`
	ProductShapeIconURL    string `json:"productShapeIconUrl"`
	ProductNameEn          string `json:"productNameEn"`
	ProductNameAr          string `json:"productNameAr"`
	CategoryURLEn          string `json:"categoryUrlEn"`
	CategoryURLAr          string `json:"categoryUrlAr"`
	Category               string `json:"category"`
	MainImageURL           string `json:"mainImageUrl"`
}

func (p Product) ToGeneric(internalTerm string) search.Response {
	return search.Response{
		ID:                     p.ID,
		ScrapperInternalName:   internalTerm,
		ProductShapeTypeName:   p.ProductShapeTypeName,
		ProductShapeTypeNameAr: p.ProductShapeTypeNameAr,
		ProductShapeIconURL:    p.ProductShapeIconURL,
		ProductNameEn:          p.ProductNameEn,
		ProductNameAr:          p.ProductNameAr,
		CategoryURLEn:          p.CategoryURLEn,
		CategoryURLAr:          p.CategoryURLAr,
		Category:               p.Category,
		MainImageURL:           p.MainImageURL,
	}
}
