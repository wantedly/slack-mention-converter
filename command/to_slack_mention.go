package command

import (
	"fmt"
	"log"
	"strings"

	"github.com/wantedly/slack_mention_converter/service"
)

type To_slack_mentionCommand struct {
	Meta
}

func (c *To_slack_mentionCommand) Run(args []string) int {
	var loginName, slackName string
	if len(args) == 1 {
		loginName = args[0]
	} else {
		log.Println(c.Help())
		return 1
	}

	user, err := service.GetUser(loginName)
	if err != nil {
		slackName = loginName
		log.Printf("Login name '%v' not found. Treat it as slack name\n", loginName)
	} else {
		slackName = user.SlackName
	}
	slackUser, err := service.GetSlackUser(slackName)
	if err != nil {
		log.Printf("%v. Slack Name: %v\n", err, slackName)
	}
	fmt.Printf("<@%v|%v>\n", slackUser.ID, slackName)

	return 0
}

func (c *To_slack_mentionCommand) Synopsis() string {
	return "Get slack mention format from login name"
}

func (c *To_slack_mentionCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
