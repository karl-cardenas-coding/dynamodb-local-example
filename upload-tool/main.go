package main

import (
	"log"
	"os"
	. "upload-data/internal"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Go SDK Github https://github.com/aws/aws-sdk-go
// Go AWS SDK API https://docs.aws.amazon.com/sdk-for-go/api/

var (
	svc           *dynamodb.DynamoDB
	tableName     = os.Getenv("TABLE_NAME")
	region        = os.Getenv("REGION")
	localDynamoDB = os.Getenv("DYNAMODB_ADDR")
	jsonFile      = os.Getenv("JSON_PATH")
)

func main() {

	if jsonFile != "" && tableName != "" && region != "" && localDynamoDB != "" {

		// Create a new DynomoDB sessions
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String(region)},
		)

		// Create DynamoDB client
		svc = dynamodb.New(sess, &aws.Config{Endpoint: aws.String(localDynamoDB)})
		if err != nil {
			log.Println("Error creating session:")
			log.Fatal(err.Error())

		}

		objectsList := GetObjectsFromFile(jsonFile)

		err = AddModelItems(svc, objectsList, tableName)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("########################################")
		log.Println("## ALL DATA LOADED SUCCESSFULLY ##")
		log.Println("########################################")

	} else {
		log.Fatal("MISSING ENVIRONMENT VARIABLE PARAMETERS - EXITING")
	}
}
