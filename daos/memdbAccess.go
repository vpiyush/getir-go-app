package daos

import (
	memdb "github.com/vpiyush/getir-go-app/inmemdb"
	"github.com/vpiyush/getir-go-app/models"
)

var cache *memdb.Cache

// PairDAO fetches key value from in-memory db
type PairDAO struct {
}

// NewPairDAO returns a new pair Dao
func NewPairDAO() *PairDAO {
	return &PairDAO{}
}

// Insert create a new key value pair in in-memory DB
func (s PairDAO) Insert(key string, value string) (*models.Pair, error) {
	return cache.Insert(key, value)
}

// Get fetches a value corresponding to given key
func (p PairDAO) Get(key string) (string, bool) {
	return cache.Get(key)
}
