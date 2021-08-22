// Packages gaos provides data access objects for mongo and in-memory database
package daos

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/vpiyush/getir-go-app/common"
	"github.com/vpiyush/getir-go-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
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

// init sets up the mongo database
func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})
	log.Debug("Connecting to mongo..")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(common.Cfg.Database.Uri)
	client, _ := mongo.Connect(ctx, clientOptions)

	// check if the connection was successful
	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	collection = client.Database(common.Cfg.Database.Name).Collection(common.Cfg.Database.Collection)
	log.Debug("Connected to mongo database ", common.Cfg.Database.Name, " and collection ", common.Cfg.Database.Collection)

	// context needs to be canceled to avoid memory leak
	//TODO:REVIEW
	//defer cancel()
}

//Find does a query to mongo database, filters records and returns record Array
func (r RecordDAO) Find(startDate string, endDate string, minCount int, maxCount int) ([]models.Record, error) {
	//TODO: review Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	sDate, _ := time.Parse(common.DateFromatYYYYMMDD, startDate)
	eDate, _ := time.Parse(common.DateFromatYYYYMMDD, endDate)
	pipeline := mongo.Pipeline{
		{
			{"$project", bson.D{
				{"_id", 0},
				{"key", 1},
				{"createdAt", 1},
				{"totalCount", bson.D{{"$sum", "$counts"}}},
			}},
		},
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
