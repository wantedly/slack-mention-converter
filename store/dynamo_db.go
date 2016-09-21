package store

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/wantedly/slack-mention-converter/models"
)

const userMapTable = "SlackUsers"

type DynamoDB struct {
	db *dynamodb.DynamoDB
}

func NewDynamoDB() *DynamoDB {
	db := dynamodb.New(session.New(&aws.Config{}))

	return &DynamoDB{
		db: db,
	}
}

func (d *DynamoDB) GetUser(loginName string) (*models.User, error) {
	resp, err := d.db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(userMapTable),
		Key: map[string]*dynamodb.AttributeValue{
			"LoginName": {
				S: aws.String(loginName),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return &models.User{
		LoginName: loginName,
		SlackName: *resp.Item["SlackName"].S,
	}, nil
}

func (d *DynamoDB) AddUser(user *models.User) error {
	_, err := d.db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(userMapTable),
		Item: map[string]*dynamodb.AttributeValue{
			"LoginName": {
				S: aws.String(user.LoginName),
			},
			"SlackName": {
				S: aws.String(user.SlackName),
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (d *DynamoDB) ListUsers() ([]*models.User, error) {
	resp, err := d.db.Scan(&dynamodb.ScanInput{
		TableName: aws.String(userMapTable),
	})
	if err != nil {
		return nil, err
	}

	var users []*models.User

	for _, item := range resp.Items {
		users = append(users, models.NewUser(*item["LoginName"].S, *item["SlackName"].S))
	}

	return users, nil
}

func (d *DynamoDB) PutUsers(users []*models.User) error {
	for _, user := range users {
		if err := d.AddUser(user); err != nil {
			return err
		}
	}

	return nil
}
