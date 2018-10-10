# slack emoji exporter

## how to use
1. clone this repo (`$ git clone git@github.com:sottar/slack-emoji-exporter.git`)
2. install `go/dep` (`$ go get -u github.com/golang/dep/cmd/dep` or `$ brew install dep`)
3. `$ dep ensure`
4. set Slack api token in `.env`
5. set up in slack.com
  1. Create Slack App at https://api.slack.com/apps
  2. Add `emoji:read` permission to created App
6. `$ go run main.go`
