package main

import (
	web "github.com/hhandhuan/ku-bbs/cmd/webserver"
	_ "github.com/hhandhuan/ku-bbs/pkg/config"
	_ "github.com/hhandhuan/ku-bbs/pkg/db"
	_ "github.com/hhandhuan/ku-bbs/pkg/redis"
)

func main() {
	web.Run()
}
