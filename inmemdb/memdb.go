//Package inmemdb is a simple implemention of a in memory database
package inmemdb

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/vpiyush/getir-go-app/models"
	"os"
	"sync"
)

// in memmory data base, it supports parallel reads and exclusive writes
type cache struct {
	items map[string]string
	mu    sync.RWMutex
}

// Exports acces to in-memory DB
var Cache *cache

// init initalizes the in-memory DB
func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})
	Cache = &cache{
		items: make(map[string]string),
		mu:    sync.RWMutex{},
	}
}

//Insert creates a new key value pair in in-memory db
func (c *cache) Insert(key string, value string) (*models.Pair, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.items[key]; ok {
		return nil, errors.New("Key Already Exists")
	}
	c.items[key] = value
	return &models.Pair{
		Key:   key,
		Value: value,
	}, nil
}

//Get fetches the key value pair based on given key
func (c *cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if value, ok := c.items[key]; ok {
		return value, true
	}
	return "", false
}
