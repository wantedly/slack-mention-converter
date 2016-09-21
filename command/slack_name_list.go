package command

import (
	"fmt"
	"log"
	"strings"

	"github.com/wantedly/slack-mention-converter/service"
	"github.com/wantedly/slack-mention-converter/store"
)

type SlackNameListCommand struct {
	Meta
}

func (c *SlackNameListCommand) Run(args []string) int {
	var s store.Store

	// dir, _ := os.Getwd()
	// dir = filepath.Join(dir, "data")
	// s = store.NewCSV(dir)
	s = store.NewDynamoDB()

	users, err := service.ListSlackUsers(s)
	if err != nil {
		log.Println(err)
		return 1
	}
	for _, user := range users {
		fmt.Println(user)
	}

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
