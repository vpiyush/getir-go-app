package daos

import (
	memdb "github.com/vpiyush/getir-go-app/inmemdb"
	"github.com/vpiyush/getir-go-app/models"
)

// PairDAO fetches key value from in-memory db
type PairDAO struct {
}

// NewPairDAO returns a new pair Dao
func NewPairDAO() *PairDAO {
	return &PairDAO{}
}

// Insert create a new key value pair in in-memory DB
func (s PairDAO) Insert(key string, value string) (*models.Pair, error) {
	return memdb.Cache.Insert(key, value)
}

// Get fetches a value corresponding to given key
func (p PairDAO) Get(key string) (string, bool) {
	return memdb.Cache.Get(key)
}
