package main

import (
	"github.com/mitchellh/cli"
	"github.com/wantedly/slack_mention_converter/command"
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
		"slack_name_list": func() (cli.Command, error) {
			return &command.Slack_name_listCommand{
				Meta: *meta,
			}, nil
		},
		"to_slack_name": func() (cli.Command, error) {
			return &command.To_slack_nameCommand{
				Meta: *meta,
			}, nil
		},
		"to_slack_mention": func() (cli.Command, error) {
			return &command.To_slack_mentionCommand{
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
