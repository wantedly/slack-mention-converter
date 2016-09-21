package command

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/wantedly/slack-mention-converter/service"
	"github.com/wantedly/slack-mention-converter/store"
)

type ListCommand struct {
	Meta
}

func (c *ListCommand) Run(args []string) int {
	var s store.Store

	dir, _ := os.Getwd()
	dir = filepath.Join(dir, "data")
	s = store.NewCSV(dir)

	users, err := service.ListUsers(s)
	if err != nil {
		log.Println(err)
		return 1
	}
	for _, user := range users {
		fmt.Println(user)
	}

	return 0
}

func (c *ListCommand) Synopsis() string {
	return "List mapping of <login_name> and <slack_name>"
}

func (c *ListCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
