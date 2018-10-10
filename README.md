# slack emoji exporter

## how to use

### in slack
1. Create Slack App at https://api.slack.com/apps
2. Add `emoji:read` permission to created App

### in terminal
1. clone this repo (`$ git clone git@github.com:sottar/slack-emoji-exporter.git`)
2. install `go/dep` (`$ go get -u github.com/golang/dep/cmd/dep` or `$ brew install dep`)
3. run `$ dep ensure`
4. set SLACK_API_TOKEN in `.env`
5. run `$ go run main.go`
