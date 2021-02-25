package repositorydna

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/gusaul/go-dynamock"
	"github.com/jjoc007/meli-xmen-detector/xmen/config"
	dnamodel "github.com/jjoc007/meli-xmen-detector/xmen/model/dna"
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

func Test_repository_Create(t *testing.T) {
	resource := &dnamodel.DNA{
		ID: "123456789",
		DNA: []string{
			"AAAAAA",
			"AAAAAA",
		},
		IsMutant: true,
	}

	av, _ := dynamodbattribute.MarshalMap(resource)
	mock.ExpectPutItem().ToTable(config.POCDNATable).WithItems(av).WillReturns(dynamodb.PutItemOutput{})
	t.Run("Create DNA Record", func(t *testing.T) {
		r := NewRepository(db)
		err := r.Create(resource)
		if err != nil {
			t.Errorf("Create() error = %v", err)
		}
	})
}

func Test_repository_GetByID(t *testing.T) {
	id := "11111111"
	av := map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String(id),
		},
	}

	resourceResult := &dnamodel.DNA{
		ID: "11111111",
		DNA: []string{
			"AAAAAA",
			"AAAAAA",
		},
		IsMutant: true,
	}

	avResult, _ := dynamodbattribute.MarshalMap(resourceResult)
	expectedResult := dynamodb.GetItemOutput{
		ConsumedCapacity: nil,
		Item:             avResult,
	}

	mock.ExpectGetItem().ToTable(config.POCDNATable).WithKeys(av).WillReturns(expectedResult)
	t.Run("Get DNA Record by ID", func(t *testing.T) {
		r := NewRepository(db)
		res, err := r.GetByID(id)
		if err != nil {
			t.Errorf("GetByID() error = %v", err)
		}
		if res == nil {
			t.Error("GetByID() error = result nil")
		}
		if res.IsMutant != resourceResult.IsMutant {
			t.Errorf("GetByID() IsMutant = %v, want %v", res.IsMutant, resourceResult.IsMutant)
		}
		if res.ID != id {
			t.Errorf("GetByID() id = %v, want %v", res.ID, id)
		}
		if len(res.DNA) != len(resourceResult.DNA) {
			t.Errorf("GetByID() len dna = %v, want %v", len(res.DNA), len(resourceResult.DNA))
		}
	})
}