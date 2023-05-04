package veseeta

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"

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
}

func NewScrapper() Scrapper {
	return Scrapper{
		baseLink: &url.URL{
			Scheme: SCHEMA,
			Host:   HOST,
			Path:   PATH,
		},
	}
}

func (s *Scrapper) Search(query string) (search.Response, error) {
	response, err := http.Get(s.URL(query))
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
	log.Printf("%v", searchResponse.Product)
	return search.Response{}, nil
}

func (s *Scrapper) URL(query string) string {
	params := url.Values{}
	params.Add(VEERSION_FIELD, VERSION_VALUE)
	params.Add(SIZE_FIELD, SIZE_VALUE)
	params.Add(QUERY_FIELD, query)
	url := *s.baseLink
	url.RawQuery = params.Encode()
	return url.String()
}
