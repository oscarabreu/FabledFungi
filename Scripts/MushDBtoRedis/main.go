package main

import (
	"fmt"
	"os"

	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-redis/redis/v8"
)

func main() {
	awsSession := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	dynamoDBClient := dynamodb.New(awsSession)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "mushroomcache-rtfzbt.serverless.use1.cache.amazonaws.com:6379",
		Password: "", 
		DB:       0,  
	})

	ctx := context.Background()

	
	tableName := "Mushroom"
    primaryKeyAttribute := "Source#Species" 

    scanInput := &dynamodb.ScanInput{
        TableName: aws.String(tableName),
        ProjectionExpression: aws.String("#pk"), 
        ExpressionAttributeNames: map[string]*string{
            "#pk": aws.String(primaryKeyAttribute), 
        },
    }

    result, err := dynamoDBClient.Scan(scanInput)
    if err != nil {
        fmt.Println("Error querying DynamoDB:", err)
        os.Exit(1)
    }
	counter := 0

	for _, item := range result.Items {
		primaryKeyValue := *item[primaryKeyAttribute].S 
		err := redisClient.Set(ctx, fmt.Sprintf("%d", counter), primaryKeyValue, 0).Err()
		if err != nil {
			fmt.Println("Error storing key in Redis:", err)
			os.Exit(1)
		}

		counter++
	}
}