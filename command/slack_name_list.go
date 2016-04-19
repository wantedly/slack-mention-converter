package command

import (
	"fmt"
	"log"
	"strings"

	"github.com/wantedly/slack-mention-converter/service"
)

type SlackNameListCommand struct {
	Meta
}

func (c *SlackNameListCommand) Run(args []string) int {
	users, err := service.ListSlackUsers()
	if err != nil {
		log.Println(err)
		return 1
	}
	fmt.Println(users)

	return 0
}

func (c *SlackNameListCommand) Synopsis() string {
	return "List up slack users id and name mapping"
}

func (c *SlackNameListCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
