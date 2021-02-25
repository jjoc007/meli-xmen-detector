package repositorydna

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/jjoc007/meli-xmen-detector/xmen/config"
	"github.com/jjoc007/meli-xmen-detector/xmen/log"
	dnamodel "github.com/jjoc007/meli-xmen-detector/xmen/model/dna"
)

// DNARepository describes the dna repository.
type DNARepository interface {
	Create(*dnamodel.DNA) error
	GetByID(string) (*dnamodel.DNA, error)
}

// NewRepository creates and returns a new dna repository instance
func NewRepository(database dynamodbiface.DynamoDBAPI) DNARepository {
	return &repository{
		database: database,
		table:    config.POCDNATable,
	}
}

type repository struct {
	database dynamodbiface.DynamoDBAPI
	table    string
}

func (s *repository) Create(resource *dnamodel.DNA) (err error) {
	log.Logger.Debug().Msg("Adding a new dna")

	av, err := dynamodbattribute.MarshalMap(resource)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = s.database.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(s.table),
		Item:      av,
	})

	log.Logger.Debug().Msgf("ID %v inserted.\n", resource.ID)
	return nil
}

func (s *repository) GetByID(id string) (dna *dnamodel.DNA, err error) {
	log.Logger.Debug().Msgf("Getting dna by ID")

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

	if result.Item == nil {
		return nil, nil
	}

	dna = &dnamodel.DNA{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &dna)
	if err != nil {
		return
	}
	return
}