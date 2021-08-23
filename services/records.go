package services

import (
	"github.com/vpiyush/getir-go-app/models"
	"time"
)

type recordDAO interface {
	Find(startDate time.Time, endDate time.Time, minCount int, maxCount int) ([]models.Record, error)
}

type RecordService struct {
	dao recordDAO
}

// NewRecordService creates a new RecordService with the given record DAO.
func NewRecordService(dao recordDAO) *RecordService {
	return &RecordService{dao}
}

// Get just retrieves records using record DAO, here can be additional logic for processing data retrieved by DAOs
func (s *RecordService) Find(startDate time.Time, endDate time.Time, minCount int, maxCount int) ([]models.Record, error) {
	return s.dao.Find(startDate, endDate, minCount, maxCount)
}
