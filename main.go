package main

import (
	"os"

	"github.com/graduation-fci/multivendor-scrapper/handler"
	"github.com/graduation-fci/multivendor-scrapper/search"
	"github.com/graduation-fci/multivendor-scrapper/veseeta"
)

var (
	FILE_PATH string
	SHEET     string
)

func main() {
	SetupEnvs()
	theif := search.NewTheif(
		search.TheifOpts{
			Threads:         search.SINGLE_THREADED,
			YieldThereshold: search.NO_YIELD,
			YieldMillis:     search.NO_SLEEP,
		},
	)

	theif.SetDriver(veseeta.NewScrapper()).SetGoal(
		handler.LoadProducts(FILE_PATH, SHEET),
	)

	theif.StartRobbery()
}

func SetupEnvs() {
	FILE_PATH = os.Getenv("filepath")
	SHEET = os.Getenv("sheet")
	if FILE_PATH == "" || SHEET == "" {
		panic("Missing env vars")
	}
}
