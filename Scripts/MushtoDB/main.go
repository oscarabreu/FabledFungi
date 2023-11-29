package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Mushroom struct {
	SourceSpecies    string `dynamodbav:"Source#Species"`
    ObservationID    string `dynamodbav:"Observation"`
    TaxonID       	 string `json:"taxon_id"`
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
	// Open CSV File
	csvFile, err := os.Open("mushrooms.csv")
	if err != nil {
		log.Fatalf("failed to open file, %v", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = ',' 
	if _, err := reader.Read(); err != nil {
		log.Fatalf("failed to read header row, %v", err)
	}
	
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

    svc := dynamodb.NewFromConfig(cfg)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to read record, %v", err)
		}
		
		mushroom := Mushroom{
			ObservationID:    record[0],
			ObservedOn:       record[1],
			UserID:           record[2],
			UserLogin:        record[3],
			CreatedAt:        record[4],
			UpdatedAt:        record[5],
			License:          record[6],
			URL:              record[7],
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
			ImageURL:         "s3://fungis3/iNaturalist/" + record[0],
			SourceSpecies: "iNaturalist#" + record[24],
			
		}

		av, err := attributevalue.MarshalMap(mushroom)
		if err != nil {
    		log.Fatalf("failed to marshal Mushroom struct, %v", err)
		}
		fmt.Printf("Imported: %s\n", record[0])

	
        _, err = svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
			TableName: aws.String("Mushroom"),
			Item: av, 
		})
        if err != nil {
            log.Fatalf("failed to put record to DynamoDB, %v", err)
        }

    }
    log.Println("CSV data imported to DynamoDB successfully.")
}