package user

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/repositories"
	"github.com/twinj/uuid"
	"log"
	"time"
)

type userRepository struct {
	Timeout time.Duration
	client  *dynamodb.DynamoDB
}

func (u userRepository) Insert(ctx context.Context, user models.User) error {
	c, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.DeletedAt = nil
	UUID := uuid.NewV4()
	user.UUID = UUID.String()
	item, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		log.Println(err)
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("Users"),
		Item:      item,
		ExpressionAttributeNames: map[string]*string{
			"#UUID":     aws.String("UUID"),
			"#Username": aws.String("Username"),
		},
		ConditionExpression: aws.String("attribute_not_exists(#UUID) and attribute_not_exists(#Username)"),
	}

	if _, err := u.client.PutItemWithContext(c, input); err != nil {
		log.Println(err)

		if _, ok := err.(*dynamodb.ConditionalCheckFailedException); ok {
			return err
		}
		return err
	}
	return nil
}
func (u userRepository) FindByUUID(ctx context.Context, id string) (models.User, error) {
	c, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()

	input := &dynamodb.GetItemInput{
		TableName: aws.String("Users"),
		Key: map[string]*dynamodb.AttributeValue{
			"UUID": {S: aws.String(id)},
		},
	}

	res, err := u.client.GetItemWithContext(c, input)
	if err != nil {
		log.Println(err)

		return models.User{}, err
	}

	if res.Item == nil {
		return models.User{}, err
	}

	var user models.User
	if err := dynamodbattribute.UnmarshalMap(res.Item, &user); err != nil {
		log.Println(err)
		return models.User{}, err
	}
	return user, nil
}
func (u userRepository) FindByUsername(ctx context.Context, username string) ([]models.User, error) {
	c, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()

	input := &dynamodb.ScanInput{
		TableName:        aws.String("Users"),
		FilterExpression: aws.String("contains(Username, :Username)"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":Username": {S: aws.String(username)},
		},
	}

	res, err := u.client.ScanWithContext(c, input)
	if err != nil {
		log.Println(err)
		return []models.User{}, err
	}

	if *res.Count == 0 {
		return []models.User{}, err
	}

	var users []models.User
	for _, userItem := range res.Items {
		var userToScan models.User
		err := dynamodbattribute.UnmarshalMap(userItem, &userToScan)
		if err != nil {
			log.Println(err)
			return users, err
		}
		if userToScan.Username != username {
			continue
		}
		users = append(users, userToScan)
	}
	return users, nil
}

func (u userRepository) Delete(ctx context.Context, uuid string) error {
	c, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String("Users"),
		Key: map[string]*dynamodb.AttributeValue{
			"UUID": {S: aws.String(uuid)},
		},
	}

	if _, err := u.client.DeleteItemWithContext(c, input); err != nil {
		log.Println(err)

		return err
	}
	return nil
}

func (u userRepository) Update(ctx context.Context, user models.User) error {
	c, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("Users"),
		Key: map[string]*dynamodb.AttributeValue{
			"UUID": {S: aws.String(user.UUID)},
		},
		ExpressionAttributeNames: map[string]*string{
			"#Username":  aws.String("Username"),
			"#Password":  aws.String("Password"),
			"#UpdatedAt": aws.String("UpdatedAt"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":Username":  {S: aws.String(user.Username)},
			":Password":  {S: aws.String(user.Password)},
			":UpdatedAt": {S: aws.String(time.Now().Format(time.RFC3339))},
		},
		UpdateExpression: aws.String("set #Username = :Username, #Password = :Password, #UpdatedAt = :UpdatedAt"),
		ReturnValues:     aws.String("UPDATED_NEW"),
	}

	if _, err := u.client.UpdateItemWithContext(c, input); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (u userRepository) CreateTable(ctx context.Context) error {
	result, err := u.listTables(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	if contains(result.TableNames, "Users") {
		return nil
	}
	input := &dynamodb.CreateTableInput{
		TableName: aws.String("Users"),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("UUID"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("UUID"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
	}

	out, err := u.client.CreateTableWithContext(ctx, input)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("Successfully created table %s", out)
	return nil
}

func contains(list []*string, compareItem string) bool {
	for _, listItem := range list {
		if *listItem == compareItem {
			return true
		}
	}
	return false
}
func (u userRepository) listTables(ctx context.Context) (*dynamodb.ListTablesOutput, error) {
	c, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()

	input := &dynamodb.ListTablesInput{}
	result, err := u.client.ListTablesWithContext(c, input)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

//ProvideUserRepository Provide user repository
func ProvideUserRepository(timeout time.Duration, db *dynamodb.DynamoDB) repositories.UserRepository {
	return userRepository{Timeout: timeout, client: db}
}
