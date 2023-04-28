package veseeta

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
	Searchable             bool   `json:"searchable"`
}
