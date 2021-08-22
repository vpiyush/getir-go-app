package services

import (
	"github.com/vpiyush/getir-go-app/models"
)

type recordDAO interface {
	Find(startDate string, endDate string, minCount int, maxCount int) ([]models.Record, error)
}

type RecordService struct {
	dao recordDAO
}

// NewRecordService creates a new RecordService with the given record DAO.
func NewRecordService(dao recordDAO) *RecordService {
	return &RecordService{dao}
}

// Get just retrieves records using record DAO, here can be additional logic for processing data retrieved by DAOs
func (s *RecordService) Find(startDate string, endDate string, minCount int, maxCount int) ([]models.Record, error) {
	return s.dao.Find(startDate, endDate, minCount, maxCount)
}
