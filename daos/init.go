package daos

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/vpiyush/getir-go-app/common"
	"github.com/vpiyush/getir-go-app/inmemdb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

// init sets up the mongo database
func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})

	log.Debug("Connecting to mongo..")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(common.Cfg.Database.Uri)
	client, _ := mongo.Connect(ctx, clientOptions)
	// check if the connection was successful
	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	collection = client.Database(common.Cfg.Database.Name).Collection(common.Cfg.Database.Collection)
	log.Debug("Connected to mongo database ", common.Cfg.Database.Name, " and collection ", common.Cfg.Database.Collection)

	// setup in memory cache
	cache = inmemdb.NewMemCache()
}
