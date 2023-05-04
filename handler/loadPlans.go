package handler

import (
	"log"

	"github.com/graduation-fci/multivendor-scrapper/search"
	excelize "github.com/xuri/excelize/v2"
)

const TAG = "[Handler]"

func LoadProducts(path, sheet string) func() ([]search.Item, error) {
	return func() ([]search.Item, error) {
		file, err := excelize.OpenFile(path)
		if err != nil {
			log.Println(TAG, "cannot open file", path, "to load products with error", err)

			return []search.Item{}, err
		}

		rows, err := file.GetRows(sheet)
		if err != nil {
			log.Println(TAG, "cannot open sheet", sheet, "on file", path, "with error", err)
			return []search.Item{}, err
		}

		var searchTerms []search.Item
		for _, row := range rows {
			searchTerms = append(searchTerms, row[0])
		}

		return searchTerms, nil
	}
}
