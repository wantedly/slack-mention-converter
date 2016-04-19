package command

import (
	"fmt"
	"log"
	"strings"

	"github.com/wantedly/slack_mention_converter/service"
)

type To_slack_nameCommand struct {
	Meta
}

func (c *To_slack_nameCommand) Run(args []string) int {
	var loginName string
	if len(args) == 1 {
		loginName = args[0]
	} else {
		log.Println(c.Help())
		return 1
	}

	user, err := service.GetUser(loginName)
	if err != nil {
		log.Printf("Login name '%v' not found\n", loginName)
	}
	fmt.Printf("@%v\n", user.SlackName)

	return 0
}

func (c *To_slack_nameCommand) Synopsis() string {
	return "Get slack name from login name"
}

func (c *To_slack_nameCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
