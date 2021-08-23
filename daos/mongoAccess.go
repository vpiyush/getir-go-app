// Packages gaos provides data access objects for mongo and in-memory database
package daos

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/vpiyush/getir-go-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var collection *mongo.Collection

// RecordDAO fetches data from mongo DB
type RecordDAO struct {
}

// NewRecordDAO creates a new RecordDAO
func NewRecordDAO() *RecordDAO {
	return &RecordDAO{}
}

//Find does a query to mongo database, filters records and returns record Array
func (r RecordDAO) Find(sDate time.Time, eDate time.Time, minCount int, maxCount int) ([]models.Record, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// needs to be canceled to avoid memory leak
	defer cancel()
	// create mongo pipe lint for filtering
	pipeline := mongo.Pipeline{
		// project stage, project only required members
		{
			{"$project", bson.D{
				{"_id", 0},
				{"key", 1},
				{"createdAt", 1},
				{"totalCount", bson.D{{"$sum", "$counts"}}},
			}},
		},
		// match stage, filter records based on below matching conditions
		{
			{"$match", bson.D{
				{"totalCount", bson.D{
					{"$lte", maxCount},
					{"$gte", minCount},
				}},
				{"createdAt", bson.D{
					{"$lte", eDate},
					{"$gte", sDate},
				}},
			}},
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Error("Error in record Aggregation: ", err)
		return nil, err
	}
	defer cursor.Close(ctx)
	var records []models.Record
	if err = cursor.All(ctx, &records); err != nil {
		log.Error("Error while decoing records from cursor: ", err)
		return nil, err
	}
	return records, err

}
