package repositorystats

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/jjoc007/meli-xmen-detector/xmen/config"
	"github.com/jjoc007/meli-xmen-detector/xmen/log"
	statmodel "github.com/jjoc007/meli-xmen-detector/xmen/model/stats"
)

// StatsRepository describes the dna repository.
type StatsRepository interface {
	Update(*statmodel.Stats) error
	GetByID(string) (*statmodel.Stats, error)
}

// NewRepository creates and returns a new dna repository instance
func NewRepository(database dynamodbiface.DynamoDBAPI) StatsRepository {
	return &repository{
		database: database,
		table:    config.POCStatsTable,
	}
}

type repository struct {
	database dynamodbiface.DynamoDBAPI
	table    string
}

func (s *repository) Update(resource *statmodel.Stats) (err error) {
	log.Logger.Debug().Msg("Update stats")

	av, err := dynamodbattribute.MarshalMap(resource)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = s.database.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(s.table),
		Item:      av,
	})

	log.Logger.Debug().Msgf("ID %v updated.\n", resource.ID)
	return nil
}

func (s *repository) GetByID(id string) (stats *statmodel.Stats, err error) {
	log.Logger.Debug().Msgf("Getting stats")

	result, err := s.database.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(s.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	stats = &statmodel.Stats{}
	if result.Item == nil {
		return stats, nil
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &stats)
	if err != nil {
		return
	}

	return
}