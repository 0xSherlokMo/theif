package search

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/graduation-fci/multivendor-scrapper/storage"
	"github.com/schollz/progressbar/v3"
)

const (
	NO_YIELD        = 0
	NO_SLEEP        = 0
	SINGLE_THREADED = 1
)

type TheifOpts struct {
	Threads         int
	YieldThereshold int
	YieldMillis     int
}

type Theif struct {
	Opts TheifOpts

	SessionId string

	driver SearchDriver

	checkList []Item

	fileManager storage.FileManager[Response]
}

func NewTheif(opts TheifOpts) *Theif {
	return &Theif{
		Opts: opts,
	}
}

func (t *Theif) Session() string {
	return fmt.Sprintf("[%s]", t.SessionId)
}

func (t *Theif) SetDriver(driver SearchDriver) *Theif {
	t.driver = driver
	return t
}

func (t *Theif) SetGoal(handler func() ([]Item, error)) *Theif {
	items, err := handler()
	if err != nil {
		panic(TAG + " Cannot load terms with error " + err.Error())
	}
	t.checkList = items
	return t
}

func (t *Theif) StartRobbery() {
	log.Println(TAG, t.driver.Identifier(), "Preparing robbery")
	t.prepareRobbery()
	log.Println(TAG, t.driver.Identifier(), "Started Session", t.Session())

	itemsChannel := make(chan Item, t.Opts.Threads)
	var wg sync.WaitGroup
	wg.Add(t.Opts.Threads)
	for thread := 1; thread <= t.Opts.Threads; thread++ {
		go t.hand(thread, itemsChannel, &wg)
	}

	progressBar := progressbar.Default(int64(len(t.checkList)), "Robbery Percentage From", t.driver.Identifier())

	for counter, item := range t.checkList {
		if t.reachedThereShold(counter) {
			log.Println(TAG, t.driver.Identifier(), "Reached thereshold, resting for", t.Opts.YieldMillis/1000, "Second")
			time.Sleep(time.Millisecond * time.Duration(t.Opts.YieldMillis))
		}
		itemsChannel <- item
		progressBar.Add(1)
	}

	wg.Wait()
}

func (t *Theif) prepareRobbery() {
	t.SessionId = uuid.NewString()
	err := t.fileManager.StartSession(t.SessionId)
	if err != nil {
		panic(TAG + " Cannot Start Session, due to " + err.Error() + " session id " + t.SessionId)
	}
}

func (t *Theif) hand(id int, items chan Item, wg *sync.WaitGroup) {
	threadTag := fmt.Sprintf("[Thread: %d]", id)
	for item := range items {
		response, err := t.driver.Search(item)
		if err != nil {
			log.Println(TAG, t.Session(), t.driver.Identifier(), "cannot search term:", item, "due to:", err)
			continue
		}
		err = t.fileManager.WriteToDisk(response)
		if err != nil {
			log.Println(TAG, t.Session(), threadTag, t.driver.Identifier(), "cannot commit to disk:", item, "due to:", err)
		}
	}

	wg.Done()
}

func (t *Theif) reachedThereShold(counter int) bool {
	const zeroIndexOffest = 1
	counter = counter + zeroIndexOffest

	return (counter % t.Opts.YieldThereshold) == 0
}
