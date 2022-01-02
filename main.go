package logharvestorgo

import (
	"fmt"
)

type Log struct {
	Lvl string `json:"type"`
	Msg string `json:"msg"`
}

func main() {
	LH := Forwarder{}
	d := LH.init(Config{})

	d.config.token = "123"
	fmt.Println(d)
}
