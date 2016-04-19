# slack_mention_converter

## Description

Convert login_name or account_name to slack mention format.

The most simple example usage is

```
$ slack_mention_converter register <your_login_name> <your_slack_name>
$ slack_mention_converter to_slack_mention <your_login_name>
```

## Usage

```
usage: slack_mention_converter [--version] [--help] <command> [<args>]

Available commands are:
    list                List mapping of <login_name> and <slack_name
    register            Register LoginName and SlackName mapping
    slack_name_list     List up slack users id and name mapping
    to_slack_mention    Get slack mention format from login name
    to_slack_name       Get slack name from login name
    version             Print slack_mention_converter version and quit
```

## Use by Docker

### Build

```
docker build -t quay.io/wantedly/slack_mention_converter .
```

### Run

```
docker run --rm \
  -e SLACK_TOKEN=<slack token get by https://api.slack.com/docs/oauth-test-tokens>  \
  -v data:/data \
  quay.io/wantedly/slack_mention_converter \
  <command>
```


## Install

To install, use `go get`:

```bash
$ go get -d github.com/wantedly/slack_mention_converter
```

## Contribution

1. Fork ([https://github.com/wantedly/slack_mention_converter/fork](https://github.com/wantedly/slack_mention_converter/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[wantedly](https://github.com/wantedly)
