package command

import (
	"fmt"
	"log"
	"strings"

	"github.com/wantedly/slack-mention-converter/service"
	"github.com/wantedly/slack-mention-converter/store"
)

type ToSlackMentionCommand struct {
	Meta
}

func (c *ToSlackMentionCommand) Run(args []string) int {
	var loginName, slackName string
	if len(args) == 1 {
		loginName = args[0]
	} else {
		log.Println(c.Help())
		return 1
	}

	var s store.Store

	// dir, _ := os.Getwd()
	// dir = filepath.Join(dir, "data")
	// s = store.NewCSV(dir)
	s = store.NewDynamoDB()

	user, err := service.GetUser(s, loginName)
	if err != nil {
		slackName = loginName
		log.Printf("Login name '%v' not found. Treat it as slack name\n", loginName)
		return 1
	}

	slackName = user.SlackName

	slackUser, err := service.GetSlackUser(s, slackName)
	if err != nil {
		log.Printf("%v. Slack Name: %v\n", err, slackName)
		return 1
	}
	fmt.Println(slackUser)

	return 0
}

func (c *ToSlackMentionCommand) Synopsis() string {
	return "Get slack mention format from login name"
}

func (c *ToSlackMentionCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
