package veseeta

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	"github.com/graduation-fci/multivendor-scrapper/search"
)

const (
	TAG            = "[veseeta]"
	SCHEMA         = "https"
	HOST           = "v-gateway.vezeetaservices.com"
	PATH           = "/inventory/api/V2/ProductShapes"
	VEERSION_FIELD = "version"
	VERSION_VALUE  = "2"
	SIZE_FIELD     = "size"
	SIZE_VALUE     = "5"
	QUERY_FIELD    = "query"
)

type Scrapper struct {
	baseLink *url.URL
	metric   strutil.StringMetric
}

func NewScrapper() Scrapper {
	metric := metrics.NewLevenshtein()
	metric.CaseSensitive = false

	return Scrapper{
		baseLink: &url.URL{
			Scheme: SCHEMA,
			Host:   HOST,
			Path:   PATH,
		},
		metric: metric,
	}
}

func (s Scrapper) Identifier() string {
	return TAG
}

func (s Scrapper) Search(term string) (search.Response, error) {
	response, err := http.Get(s.URL(term))
	if err != nil {
		log.Printf("%s, Error in HTTP Request. info: %s", TAG, err)
		return search.Response{}, err
	}

	if response.StatusCode != http.StatusOK {
		log.Printf("%s, Status code is not okay code: %d with body: %v", TAG, response.StatusCode, response.Header)
		return search.Response{}, errors.New("NOT_OK_STATUS_CODE")
	}

	var searchResponse SearchResponse
	if err := json.NewDecoder(response.Body).Decode(&searchResponse); err != nil {
		log.Printf("%s, Error in Decoding Body. Error: %s", TAG, err)
		return search.Response{}, err
	}

	return s.highestSimularity(term, searchResponse.Product).ToGeneric(term), nil
}

func (s *Scrapper) URL(term string) string {
	params := url.Values{}
	params.Add(VEERSION_FIELD, VERSION_VALUE)
	params.Add(SIZE_FIELD, SIZE_VALUE)
	params.Add(QUERY_FIELD, term)
	url := *s.baseLink
	url.RawQuery = params.Encode()
	return url.String()
}

func (s *Scrapper) highestSimularity(term string, products []Product) Product {
	var highestSimularProduct Product
	highestScore := 0.0
	for _, product := range products {
		score := strutil.Similarity(term, product.ProductNameEn, s.metric)
		if score > highestScore {
			highestScore = score
			highestSimularProduct = product
		}
	}

	return highestSimularProduct
}
