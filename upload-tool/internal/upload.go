package internal

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// A function that adds items to a given table
func AddModelItems(svc *dynamodb.DynamoDB, items []Model, tableName string) error {

	var returnError error

	for _, item := range items {

		value, err := dynamodbattribute.MarshalMap(item)

		if err != nil {
			returnError = err
		}

		// Create item in table Model
		input := &dynamodb.PutItemInput{
			Item:      value,
			TableName: aws.String(tableName),
		}

		_, err = svc.PutItem(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case dynamodb.ErrCodeProvisionedThroughputExceededException:
					log.Println("ERROR: ", dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
					returnError = err
				case dynamodb.ErrCodeResourceNotFoundException:
					log.Println("ERROR: ", dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
					returnError = err

				case dynamodb.ErrCodeRequestLimitExceeded:
					log.Println("ERROR: ", dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
					returnError = err

				case dynamodb.ErrCodeInternalServerError:
					log.Println("ERROR: ", dynamodb.ErrCodeInternalServerError, aerr.Error())
					returnError = err

				default:
					log.Println("ERROR: ", aerr.Error())
					returnError = err
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				log.Println(input)
				log.Println("ERROR: ", err.(awserr.Error))
				returnError = err
			}
		}

	}

	if returnError == nil {
		log.Printf("%d items successfully uploaded to the table\n\n", len(items))
	}

	return returnError
}
