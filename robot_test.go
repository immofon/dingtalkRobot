package dingtalkRobot

import (
	"fmt"
	"os"
	"testing"
)

func Test_Robot(t *testing.T) {
	p := NewPusher(os.Getenv("token"), os.Getenv("secret"))
	err := p.PushText("Hello world!\n From golang!")
	fmt.Println(err)
}
