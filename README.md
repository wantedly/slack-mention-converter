# slack-mention-converter

## Description

Convert login_name or account_name to slack mention format.

The most simple example usage is

```
$ slack-mention-converter register <your_login_name> <your_slack_name>
$ slack-mention-converter ToSlackMention <your_login_name>
```

## Usage

```
usage: slack-mention-converter [--version] [--help] <command> [<args>]

Available commands are:
    list                List mapping of <login_name> and <slack_name>
    register            Register LoginName and SlackName mapping
    slack-name-list     List up slack users id and name mapping
    to-slack-mention    Get slack mention format from login name
    to-slack-name       Get slack name from login name
    version             Print slack-mention-converter version and quit
```

## Use by Docker

### Build

```
docker build -t quay.io/wantedly/slack-mention-converter .
```

### Run

```
docker run --rm \
  -e SLACK_TOKEN=<slack token get by https://api.slack.com/docs/oauth-test-tokens>  \
  -v data:/data \
  quay.io/wantedly/slack-mention-converter \
  <command>
```


## Install

To install, use `go get`:

```bash
$ go get -d github.com/wantedly/slack-mention-converter
```

## Contribution

1. Fork ([https://github.com/wantedly/slack-mention-converter/fork](https://github.com/wantedly/slack-mention-converter/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[wantedly](https://github.com/wantedly)
