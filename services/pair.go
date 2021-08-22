package services

import (
	log "github.com/sirupsen/logrus"
	"github.com/vpiyush/getir-go-app/models"
	"os"
)

type pairDAO interface {
	Insert(key string, value string) (*models.Pair, error)
	Get(key string) (string, bool)
}

type PairService struct {
	dao pairDAO
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

// NewRecordService creates a new PairService with the given record DAO.
func NewPairService(dao pairDAO) *PairService {
	return &PairService{dao}
}

// Insert creates key value using pair DAO, here additional logic can be for processing data retrieved by DAOs
func (s *PairService) Insert(key string, value string) (*models.Pair, error) {
	return s.dao.Insert(key, value)
}

// Get just retrieves pairs using pair DAO, here additional logic can be for processing data retrieved by DAOs
func (s *PairService) Get(key string) (string, bool) {
	return s.dao.Get(key)
}
