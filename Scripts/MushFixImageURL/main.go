package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-east-1")},
    )
    if err != nil {
        fmt.Println("Error initializing session:", err)
        return
    }

    svc := dynamodb.New(sess)

    params := &dynamodb.ScanInput{
        TableName: aws.String("Mushroom"),
    }

    for {
        result, err := svc.Scan(params)
        if err != nil {
            fmt.Println("Failed to make Query API call:", err)
            return
        }

        for _, i := range result.Items {
            if i["Observation"] == nil || i["Source#Species"] == nil {
                continue
            }
            newImageURL := "s3://fungis3/iNaturalist/" + *i["Observation"].S + ".jpg"

            key := map[string]*dynamodb.AttributeValue{
                "Source#Species": i["Source#Species"],
                "Observation":    i["Observation"],
            }

            input := &dynamodb.UpdateItemInput{
                Key:       key,
                TableName: aws.String("Mushroom"),
                UpdateExpression: aws.String("set ImageURL = :i"),
                ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
                    ":i": {S: aws.String(newImageURL)},
                },
            }

            _, err = svc.UpdateItem(input)
            if err != nil {
                fmt.Println("Failed to update item:", err)
                continue
            }

            fmt.Println("Updated ImageURL:", newImageURL)
        }

        if result.LastEvaluatedKey == nil {
            break
        }
        
        params.ExclusiveStartKey = result.LastEvaluatedKey
    }
}