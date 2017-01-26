# slack-mention-converter

[![Build Status](https://travis-ci.org/wantedly/slack-mention-converter.svg?branch=master)](https://travis-ci.org/wantedly/slack-mention-converter)

## DEPRECATED

This project is **no longer** actively maintained by Wantedly, Inc.
Newly developed [developers-account-mapper](https://github.com/wantedly/developers-account-mapper) covers almost all features on this project.
Please consider using this.

## Summary

Convert login name or account name to Slack mention format.

```bash
$ slack-mention-converter to-slack-mention dtan4
<@U02XXXXXX|dai>
```

## Description

Convert login name or account name to slack mention format. Mappings of login_name and account_name are stored at Amazon DynamoDB.

The most simple example usage is

```
$ slack-mention-converter register your_login_name your_slack_name
user your_login_name:@your_slack_name added.
$ slack-mention-converter to-slack-mention your_login_name
<@U02XXXXXX|your_slack_name>
```

## Usage

### AWS configuration

2 DynamoDB tables named `SlackNames` and `SlackIDs` must be created.

#### `SlackNames` table

|Key|Type| |
|---|----|---|
|LoginName|String|Primary key|
|SlackName|String||

#### `SlackIDs` table

|Key|Type| |
|---|----|---|
|SlackName|String|Primary key|
|SlackID|String||

In addition, `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY` and `AWS_REGION` must be set at your shell. This IAM user/role must be allowed to read/write the DynamoDB tables above.

### Command usage

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
$ make docker-build
```

### Run

```
docker run --rm \
  -e SLACK_API_TOKEN=<slack token get by https://api.slack.com/docs/oauth-test-tokens>  \
  -e AWS_ACCESS_KEY_ID=yourawsaccesskeyid \
  -e AWS_SECRET_ACCESS_KEY=yourawssecretaccesskey \
  -e AWS_REGION=ap-northeast-1 \
  quay.io/wantedly/slack-mention-converter \
  <command>
```


## Install

To install, use `go get` and `make`:

```bash
$ go get -d github.com/wantedly/slack-mention-converter
$ make
$ make install
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

## License

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
