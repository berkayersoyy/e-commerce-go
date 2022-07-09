package providers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

//NewSession Returns new session
func NewSession() (*session.Session, error) {
	return session.NewSessionWithOptions(
		session.Options{
			Config: aws.Config{
				Credentials:      credentials.NewStaticCredentials(os.Getenv("DynamoDBID"), os.Getenv("DynamoDBSECRET"), ""),
				Region:           aws.String(os.Getenv("DynamoDBREGION")),
				S3ForcePathStyle: aws.Bool(true),
				Endpoint:         aws.String(os.Getenv("DynamoDBENDPOINTURL")),
			},
			Profile: os.Getenv("DynamoDBPROFILE"),
		},
	)
}

//NewDynamoDb Returns new dynamodb instance
func NewDynamoDb(ses *session.Session) *dynamodb.DynamoDB {
	return dynamodb.New(ses)
}
