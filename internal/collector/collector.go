package collector

import (
	"github.com/OchirSan/Traffic-rotator/internal/database"
	"sync"
)

type collector struct {
	mu      *sync.Mutex
	storage []database.Event
}

func New() *collector {
	eventCollector := &collector{
		mu:      &sync.Mutex{},
		storage: make([]database.Event, 0),
	}
	return eventCollector
}

func (c *collector) Set(event *database.Event) {
	c.mu.Lock()
	c.storage = append(c.storage, *event)
	c.mu.Unlock()
}

func (c *collector) Get() []database.Event {
	c.mu.Lock()
	events := c.storage
	c.storage = make([]database.Event, 0)
	c.mu.Unlock()
	return events
}
