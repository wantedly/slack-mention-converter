package command

import (
	"fmt"
	"log"
	"strings"

	"github.com/wantedly/slack_mention_converter/service"
)

type ListCommand struct {
	Meta
}

func (c *ListCommand) Run(args []string) int {
	users, err := service.ListUsers()
	if err != nil {
		log.Println(err)
		return 1
	}
	fmt.Println(users)

	return 0
}

func (c *ListCommand) Synopsis() string {
	return "List mapping of <login_name> and <slack_name"
}

func (c *ListCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
