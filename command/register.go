package command

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/wantedly/slack-mention-converter/service"
)

type RegisterCommand struct {
	Meta
}

func (c *RegisterCommand) Run(args []string) int {
	var loginName, slackName string
	if len(args) == 1 {
		loginName = os.Getenv("USER")
		slackName = args[0]
	} else if len(args) == 2 {
		loginName = args[0]
		slackName = args[1]
	} else {
		log.Println(c.Help())
		return 1
	}

	user := service.NewUser(loginName, slackName)
	err := service.AddUser(user)
	if err != nil {
		log.Println(err)
		return 1
	}
	fmt.Printf("user %v added.\n", user)

	return 0
}

func (c *RegisterCommand) Synopsis() string {
	return "Register LoginName and SlackName mapping"
}

func (c *RegisterCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
