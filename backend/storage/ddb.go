package storage

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/jamesmoessis/dust_sensor/backend/handlers"
)

var SETTINGS = "settings"
var SETTING = "setting"
var AWS_REGION = "ap-southeast-2"
var READ_CAPACITY int64 = 1
var WRITE_CAPACITY int64 = 1

type DynamoSettingsDb struct {
	dynamo *dynamodb.Client
}

var _ handlers.SettingsDB = (*DynamoSettingsDb)(nil)

func NewDynamoSettingsDb(ctx context.Context) *DynamoSettingsDb {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(AWS_REGION),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := dynamodb.NewFromConfig(cfg)
	return &DynamoSettingsDb{
		dynamo: svc,
	}
}

func (db *DynamoSettingsDb) CreateSettingsTableIfNotExists(ctx context.Context) error {
	resp, err := db.dynamo.ListTables(ctx, &dynamodb.ListTablesInput{
		Limit: aws.Int32(100),
	})
	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
	}
	tableAlreadyExists := false
	for _, tableName := range resp.TableNames {
		if tableName == SETTINGS {
			tableAlreadyExists = true
		}
	}

	if tableAlreadyExists {
		return nil
	}

	_, err = db.dynamo.CreateTable(ctx, &dynamodb.CreateTableInput{
		TableName: &SETTINGS,
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: &SETTING,
				KeyType:       types.KeyTypeHash,
			},
		},
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: &SETTING,
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  &READ_CAPACITY,
			WriteCapacityUnits: &WRITE_CAPACITY,
		},
	})
	return err
}

func (db *DynamoSettingsDb) UpdateSettings(ctx context.Context, settings handlers.Settings) error {
	// db.dynamo.PutItem(ctx, &dynamodb.PutItemInput{
	// 	TableName: &SETTINGS,

	// })
	return nil
}

func (db *DynamoSettingsDb) GetSettings(ctx context.Context) (handlers.Settings, error) {
	return handlers.Settings{}, nil
}
