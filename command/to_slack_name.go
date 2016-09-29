package command

import (
	"fmt"
	"log"
	"strings"

	"github.com/wantedly/slack-mention-converter/service"
	"github.com/wantedly/slack-mention-converter/store"
)

type ToSlackNameCommand struct {
	Meta
}

func (c *ToSlackNameCommand) Run(args []string) int {
	var loginName string
	if len(args) != 1 {
		log.Println(c.Help())
		return 1
	}

	loginName = args[0]

	var s store.Store

	// dir, _ := os.Getwd()
	// dir = filepath.Join(dir, "data")
	// s = store.NewCSV(dir)
	s = store.NewDynamoDB()

	user, err := service.GetUser(s, loginName)
	if err != nil {
		log.Printf("Login name '%v' not found\n", loginName)
		return 1
	}
	fmt.Printf("@%v\n", user.SlackName)

	return 0
}

func (c *ToSlackNameCommand) Synopsis() string {
	return "Get slack name from login name"
}

func (c *ToSlackNameCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
