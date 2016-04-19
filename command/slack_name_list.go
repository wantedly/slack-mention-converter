package command

import (
	"fmt"
	"log"
	"strings"

	"github.com/wantedly/slack_mention_converter/service"
)

type Slack_name_listCommand struct {
	Meta
}

func (c *Slack_name_listCommand) Run(args []string) int {
	users, err := service.ListSlackUsers()
	if err != nil {
		log.Println(err)
		return 1
	}
	fmt.Println(users)

	return 0
}

func (c *Slack_name_listCommand) Synopsis() string {
	return ""
}

func (c *Slack_name_listCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
