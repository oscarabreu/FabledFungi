package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"

	// "net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Mushroom struct {
    TaxonID       	 string `json:"taxon_id"`
	ObserverID       string `json:"observer_id"`
    ObservedOn       string `json:"observed_on"`
    UserID           string `json:"user_id"`
    UserLogin        string `json:"user_login"`
    CreatedAt        string `json:"created_at"`
    UpdatedAt        string `json:"updated_at"`
    License          string `json:"license"`
    URL              string `json:"url"`
    ImageURL         string `json:"image_url"`
    TagList          string `json:"tag_list"`
    Description      string `json:"description"`
    PlaceGuess       string `json:"place_guess"`
    Latitude         string `json:"latitude"`
    Longitude        string `json:"longitude"`
    SpeciesGuess     string `json:"species_guess"`
    ScientificName   string `json:"scientific_name"`
    CommonName       string `json:"common_name"`
    TaxonKingdomName string `json:"taxon_kingdom_name"`
    TaxonPhylumName  string `json:"taxon_phylum_name"`
    TaxonClassName   string `json:"taxon_class_name"`
    TaxonOrderName   string `json:"taxon_order_name"`
    TaxonFamilyName  string `json:"taxon_family_name"`
    TaxonGenusName   string `json:"taxon_genus_name"`
    TaxonSpeciesName string `json:"taxon_species_name"`
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
	reader.Comma = ',' // CSV delimiter is a tab
	bucket := "mushrooms3"

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
			ObserverID:       record[0],
			ObservedOn:       record[1],
			UserID:           record[2],
			UserLogin:        record[3],
			CreatedAt:        record[4],
			UpdatedAt:        record[5],
			License:          record[6],
			URL:              record[7],
			ImageURL:         record[8],
			TagList:          record[9],
			Description:      record[10],
			PlaceGuess:       record[11],
			Latitude:         record[12],
			Longitude:        record[13],
			SpeciesGuess:     record[14],
			ScientificName:   record[15],
			CommonName:       record[16],
			TaxonID:          record[17],
			TaxonKingdomName: record[18],
			TaxonPhylumName:  record[19],
			TaxonClassName:   record[20],
			TaxonOrderName:   record[21],
			TaxonFamilyName:  record[22],
			TaxonGenusName:   record[23],
			TaxonSpeciesName: record[24],
		}

		jsonBytes, err := json.Marshal(mushroom)
		if err != nil {
			log.Fatalf("failed to marshal JSON, %v", err)
		}

		objectKey := fmt.Sprintf("mushrooms/%s.json", mushroom.ObserverID)
		_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(objectKey),
			Body:   bytes.NewReader(jsonBytes),
			ContentType: aws.String("application/json"),
		})
		if err != nil {
			log.Fatalf("failed to upload to S3, %v", err)
		}
		fmt.Printf("Successfully uploaded %s to S3\n", objectKey)
	}
}