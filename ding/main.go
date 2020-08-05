package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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
	} else {
		data, _ := ioutil.ReadAll(os.Stdin)
		text = string(data)
	}

	if strings.HasPrefix(text, "-h") {
		fmt.Fprintf(os.Stderr, "help: %s [text]\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "\tif text is not set, it will read text from STDIN.")
		return
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
