package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	
	client := dynamodb.NewFromConfig(cfg)

	tableName := "Mushroom"

	params := &dynamodb.ScanInput{
		TableName: &tableName,
	}
	result, err := client.Scan(context.TODO(), params)
	if err != nil {
		fmt.Println("Error scanning table:", err)
		return
	}

	for _, i := range result.Items {
		newUUID := uuid.New().String()

		key := map[string]types.AttributeValue{
            "Source#Species": i["Source#Species"], 
            "Observation":    i["Observation"],   
        }

		update := &dynamodb.UpdateItemInput{
			TableName:                 &tableName,
			Key:                       key,
            UpdateExpression: 		   aws.String("SET #U = :uuid"),
			ExpressionAttributeNames: map[string]string{
                "#U": "UUID",
            },
            ExpressionAttributeValues: map[string]types.AttributeValue{
                ":uuid": &types.AttributeValueMemberS{Value: newUUID},
            },
		}

		_, updateErr := client.UpdateItem(context.TODO(), update)
		if updateErr != nil {
			fmt.Println("Got error updating item:", updateErr)
			continue
		}
		fmt.Println("Updated item:", key)
	}
}