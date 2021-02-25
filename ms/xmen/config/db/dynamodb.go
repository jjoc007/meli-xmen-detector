package db

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/jjoc007/meli-xmen-detector/xmen/log"
)

// DataBase representation basic actions on data base
type DataBase interface {
	OpenConnection() error
	GetConnection() interface{}
}

type dynamoDataBase struct {
	databaseConnection  dynamodbiface.DynamoDBAPI
}

// NewDynamoDBStorage creates and returns a new Lock dynamo db connection instance
func NewDynamoDBStorage() (DataBase, error) {
	log.Logger.Debug().Msg("New instance Dynamo storage")
	dataBase := &dynamoDataBase{}
	err := dataBase.OpenConnection()
	if err != nil {
		return nil, err
	}
	return dataBase, nil
}

// OpenConnection start dynamo db connection
func (db *dynamoDataBase) OpenConnection() error {
	log.Logger.Info().Msgf("Starting Dynamo connection")

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)
	db.databaseConnection = dynamodbiface.DynamoDBAPI(svc)

	log.Logger.Info().Msg("DynamoDB UP")
	return nil
}

// GetConnection get dynamo db connection
func (db *dynamoDataBase) GetConnection() interface{} {
	return db.databaseConnection
}
