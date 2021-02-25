package repositorystats

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/gusaul/go-dynamock"
	"github.com/jjoc007/meli-xmen-detector/xmen/config"
	statsmodel "github.com/jjoc007/meli-xmen-detector/xmen/model/stats"
	"testing"
)

var mock *dynamock.DynaMock
var db dynamodbiface.DynamoDBAPI

func init() {
	db, mock = dynamock.New()
}

func TestNewRepository(t *testing.T) {
	t.Run("Test New", func(t *testing.T) {
		got := NewRepository(db)
		if got == nil {
			t.Errorf("NewRepository() = %v is nil", got)
		}
	})
}

func Test_repository_Update(t *testing.T) {
	resource := &statsmodel.Stats{
		ID:             "1",
		CountMutantDNA: 10,
		CountHumanDNA:  10,
		Ratio:          10,
	}

	av, _ := dynamodbattribute.MarshalMap(resource)
	mock.ExpectPutItem().ToTable(config.POCStatsTable).WithItems(av).WillReturns(dynamodb.PutItemOutput{})
	t.Run("Update Stats Record", func(t *testing.T) {
		r := NewRepository(db)
		err := r.Update(resource)
		if err != nil {
			t.Errorf("Update() error = %v", err)
		}
	})
}

func Test_repository_GetByID(t *testing.T) {
	id := "1"
	av := map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String(id),
		},
	}

	resourceResult := &statsmodel.Stats{
		ID:             "1",
		CountMutantDNA: 10,
		CountHumanDNA:  10,
		Ratio:          10,
	}

	avResult, _ := dynamodbattribute.MarshalMap(resourceResult)
	expectedResult := dynamodb.GetItemOutput{
		ConsumedCapacity: nil,
		Item:             avResult,
	}

	mock.ExpectGetItem().ToTable(config.POCStatsTable).WithKeys(av).WillReturns(expectedResult)
	t.Run("Get Stats Record by ID", func(t *testing.T) {
		r := NewRepository(db)
		res, err := r.GetByID(id)
		if err != nil {
			t.Errorf("GetByID() error = %v", err)
		}
		if res == nil {
			t.Error("GetByID() error = result nil")
		}
		if res.ID != id {
			t.Errorf("GetByID() id = %v, want %v", res.ID, id)
		}
		if res.CountHumanDNA != resourceResult.CountHumanDNA {
			t.Errorf("GetByID() CountHumanDNA = %v, want %v", res.CountHumanDNA, resourceResult.CountHumanDNA)
		}
		if res.CountMutantDNA != resourceResult.CountMutantDNA {
			t.Errorf("GetByID() CountMutantDNA = %v, want %v", res.CountMutantDNA, resourceResult.CountMutantDNA)
		}
		if res.Ratio != resourceResult.Ratio {
			t.Errorf("GetByID() Ratio = %v, want %v", res.Ratio, resourceResult.Ratio)
		}
	})
}