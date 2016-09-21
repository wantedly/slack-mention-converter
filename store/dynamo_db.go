package store

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/wantedly/slack-mention-converter/models"
)

const (
	slackIDTable      = "SlackIDs"
	userMapTable      = "SlackUsers"
	batchWriteItemMax = 25
)

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

func (d *DynamoDB) refreshSlackIDs(token string) error {
	users, err := models.RetrieveFromSlack(token)
	if err != nil {
		return err
	}

	// BatchWriteItem accepts up to 25 items at the same time
	for i := 0; i < len(users)/batchWriteItemMax+1; i++ {
		var reqs []*dynamodb.WriteRequest
		var last int

		// e.g. 60 users are divided as: [0:25], [25:50], [50:60]
		if len(users)-i*batchWriteItemMax >= batchWriteItemMax {
			last = (i + 1) * batchWriteItemMax
		} else {
			last = len(users)
		}

		for _, u := range users[i*batchWriteItemMax : last] {
			reqs = append(reqs, &dynamodb.WriteRequest{
				PutRequest: &dynamodb.PutRequest{
					Item: map[string]*dynamodb.AttributeValue{
						"SlackName": {
							S: aws.String(u.Name),
						},
						"SlackID": {
							S: aws.String(u.ID),
						},
					},
				},
			})
		}

		reqItems := make(map[string][]*dynamodb.WriteRequest)
		reqItems[slackIDTable] = reqs

		_, err = d.db.BatchWriteItem(&dynamodb.BatchWriteItemInput{
			RequestItems: reqItems,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *DynamoDB) retrieveUser(name string) (*models.SlackUser, error) {
	resp, err := d.db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(slackIDTable),
		Key: map[string]*dynamodb.AttributeValue{
			"SlackName": {
				S: aws.String(name),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return models.NewSlackUser(*resp.Item["SlackID"].S, name), nil
}

func (d *DynamoDB) retrieveUsers() ([]*models.SlackUser, error) {
	resp, err := d.db.Scan(&dynamodb.ScanInput{
		TableName: aws.String(slackIDTable),
	})
	if err != nil {
		return nil, err
	}

	var users []*models.SlackUser

	for _, item := range resp.Items {
		users = append(users, models.NewSlackUser(*item["SlackID"].S, *item["SlackName"].S))
	}

	return users, nil
}

func (d *DynamoDB) GetSlackUser(name string) (*models.SlackUser, error) {
	token := os.Getenv("SLACK_API_TOKEN")

	user, err := d.retrieveUser(name)
	if err == nil {
		return user, nil
	}

	if err := d.refreshSlackIDs(token); err != nil {
		return nil, err
	}

	user, err = d.retrieveUser(name)
	if err == nil {
		return user, nil
	}

	return nil, err
}

func (d *DynamoDB) ListSlackUsers() ([]*models.SlackUser, error) {
	token := os.Getenv("SLACK_API_TOKEN")

	users, err := d.retrieveUsers()
	if err != nil {
		return nil, err
	}

	if len(users) > 0 {
		return users, nil
	}

	if err := d.refreshSlackIDs(token); err != nil {
		return nil, err
	}

	users, err = d.retrieveUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
