package main

import (
	"fmt"
	"os"

	"github.com/immofon/dingtalkRobot"
)

func getenv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		fmt.Fprintf(os.Stderr, "You should set ENV %q\n", key)
		os.Exit(2)
	}
	return v
}

func main() {
	access_token := getenv("DINGTALK_ROBOT_ACCESS_TOKEN")
	secret := getenv("DINGTALK_ROBOT_SECRET")
	text := ""
	if len(os.Args) > 1 {
		text = os.Args[1]
	}

	if text == "" {
		return
	}

	p := dingtalkRobot.NewPusher(access_token, secret)

	err := p.PushText(text)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
