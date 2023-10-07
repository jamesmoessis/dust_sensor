package storage

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/jamesmoessis/dust_sensor/backend/handlers"
)

var TABLE_NAME = "settings"
var PRIMARY_KEY = "setting"

// SETTING_KEY defines the row in the ddb table which is updated.
// It may be altered for testing purposes.
var SETTING_KEY = "main"
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
		if tableName == TABLE_NAME {
			tableAlreadyExists = true
			break
		}
	}

	if tableAlreadyExists {
		return nil
	}

	_, err = db.dynamo.CreateTable(ctx, &dynamodb.CreateTableInput{
		TableName: &TABLE_NAME,
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: &PRIMARY_KEY,
				KeyType:       types.KeyTypeHash,
			},
		},
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: &PRIMARY_KEY,
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
	_, err := db.dynamo.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &TABLE_NAME,
		Item: map[string]types.AttributeValue{
			"setting":   &types.AttributeValueMemberS{Value: SETTING_KEY},
			"isOn":      &types.AttributeValueMemberBOOL{Value: settings.IsOn},
			"threshold": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", settings.Threshold)},
		},
	})
	return err
}

func (db *DynamoSettingsDb) GetSettings(ctx context.Context) (*handlers.Settings, error) {
	output, err := db.dynamo.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &TABLE_NAME,
		Key:       map[string]types.AttributeValue{"setting": &types.AttributeValueMemberS{Value: SETTING_KEY}},
	})
	if err != nil {
		return nil, err
	}

	isOn, ok := output.Item["isOn"]
	if !ok {
		return nil, fmt.Errorf("isOn not in item attributes")
	}
	threshold, ok := output.Item["threshold"]
	if !ok {
		return nil, fmt.Errorf("threshold not in item attributes")
	}

	isOnAttribute, ok := isOn.(*types.AttributeValueMemberBOOL)
	if !ok {
		return nil, fmt.Errorf("isOn is not of attribute type BOOL")
	}
	thresholdAttribute, ok := threshold.(*types.AttributeValueMemberN)
	if !ok {
		return nil, fmt.Errorf("threshold is not of attribute type N")
	}
	thresholdNum, err := strconv.Atoi(thresholdAttribute.Value)
	if err != nil {
		return nil, fmt.Errorf("Could not convert threshold %s to string", thresholdAttribute.Value)
	}

	return &handlers.Settings{
		IsOn:      isOnAttribute.Value,
		Threshold: thresholdNum,
	}, nil
}
