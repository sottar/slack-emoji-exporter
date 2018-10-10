package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	apiToken := os.Getenv("SLACK_API_TOKEN")
	if apiToken == "" {
		panic("SLACK_API_TOKEN is not set")
	}
	api := slack.New(apiToken)

	emojiList, err := api.GetEmoji()

	if err != nil {
		panic(err)
	}

	if err := os.Mkdir("emoji", 0777); err != nil {
		panic(err)
	}

	for name, emojiURL := range emojiList {
		if strings.Index(emojiURL, "alias") != -1 {
			continue
		}

		res, err := http.Get(emojiURL)
		if err != nil {
			panic(err)
		}

		file, err := os.Create(filepath.Join("emoji", name+".png"))
		if err != nil {
			panic(err)
		}
		io.Copy(file, res.Body)
	}
}
