package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Mushroom struct {
    ObservationID string `json:"observation_id"`
    ImageURL      string `json:"image_url"`
}

func main() {
    cfg, err := config.LoadDefaultConfig(context.TODO(),
        config.WithRegion("us-east-1"),
    )
    if err != nil {
        log.Fatalf("unable to load SDK config, %v", err)
    }

    s3Client := s3.NewFromConfig(cfg)
    csvFile, err := os.Open("mushrooms.csv")
    if err != nil {
        log.Fatalf("failed to open file, %v", err)
    }
    defer csvFile.Close()

    reader := csv.NewReader(csvFile)
    reader.Comma = ','

    bucket := "fungis3"

    if _, err := reader.Read(); err != nil {
        log.Fatalf("failed to read header row, %v", err)
    }

    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatalf("failed to read record, %v", err)
        }

        mushroom := Mushroom{
            ObservationID: record[0],
            ImageURL:      record[8],
        }

        fileName := fmt.Sprintf("%s.jpg", mushroom.ObservationID)
        objectKey := fmt.Sprintf("iNaturalist/%s", fileName)

        // Stream the image directly from the URL to S3
		response, err := http.Get(mushroom.ImageURL)
		if err != nil {
			log.Fatalf("failed to download image, %v", err)
		}
		defer response.Body.Close()
	
		// Buffer the image data
		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, response.Body)
		if err != nil {
			log.Fatalf("failed to buffer image data, %v", err)
		}
	
		// Use the buffered data for S3 upload
		_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(objectKey),
			Body:   bytes.NewReader(buf.Bytes()), // Use buffered data here
		})
		if err != nil {
			log.Fatalf("failed to upload image to S3, %v", err)
		}
	
		fmt.Printf("Successfully uploaded %s to S3\n", objectKey)
    }
}