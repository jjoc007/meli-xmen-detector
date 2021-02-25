package dna

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/jjoc007/meli-xmen-detector/xmen/config/db"
	"github.com/jjoc007/meli-xmen-detector/xmen/log"
	repositorydna "github.com/jjoc007/meli-xmen-detector/xmen/repository/dna"
	repositorystats "github.com/jjoc007/meli-xmen-detector/xmen/repository/stats"
	servicedna "github.com/jjoc007/meli-xmen-detector/xmen/service/dna"
	servicestats "github.com/jjoc007/meli-xmen-detector/xmen/service/stats"
)

// Instances is a global map that contain all object instances of app
var Instances = MakeDependencyInjection()

// MakeDependencyInjection Initialize all dependencies
func MakeDependencyInjection() map[string]interface{} {
	log.Logger.Debug().Msg("Start bootstrap app objects")
	instances := make(map[string]interface{})

	database, err := db.NewDynamoDBStorage()
	if err != nil {
		panic(err)
	}
	instances["dataBase"] = database

	instances["statsRepository"] = repositorystats.NewRepository(database.GetConnection().(dynamodbiface.DynamoDBAPI))
	instances["statsService"] = servicestats.New(
		instances["statsRepository"].(repositorystats.StatsRepository))

	instances["dnaRepository"] = repositorydna.NewRepository(database.GetConnection().(dynamodbiface.DynamoDBAPI))
	instances["dnaService"] = servicedna.New(
		instances["dnaRepository"].(repositorydna.DNARepository),
		instances["statsService"].(servicestats.StatsService),
		)

	log.Logger.Debug().Msg("End bootstrap app objects")
	return instances
}
