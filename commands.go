package main

import (
	"github.com/mitchellh/cli"
	"github.com/wantedly/slack-mention-converter/command"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"register": func() (cli.Command, error) {
			return &command.RegisterCommand{
				Meta: *meta,
			}, nil
		},
		"list": func() (cli.Command, error) {
			return &command.ListCommand{
				Meta: *meta,
			}, nil
		},
		"slack-name-list": func() (cli.Command, error) {
			return &command.SlackNameListCommand{
				Meta: *meta,
			}, nil
		},
		"to-slack-name": func() (cli.Command, error) {
			return &command.ToSlackNameCommand{
				Meta: *meta,
			}, nil
		},
		"to-slack-mention": func() (cli.Command, error) {
			return &command.ToSlackMentionCommand{
				Meta: *meta,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}
