package main

import (
	_ "gitee.com/xygfm/authorization/apps"
	"gitee.com/xygfm/authorization/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
