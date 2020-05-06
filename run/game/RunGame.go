package main

import (
	"github.com/ouczbs/goworld"
	"os"
	"strings"
)

var (
	_SERVICE_NAMES = []string{
		"OnlineService",
		"SpaceService",
	}
)

func main() {
	cmd := "cmd -gid 10086"
	os.Args = strings.Split(cmd," ")
	// 运行游戏服务器
	goworld.Run()
}
