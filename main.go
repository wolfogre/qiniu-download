package main

import (
	"flag"
	"net/http"
	"github.com/wolfogre/qiniuauth/internal/handler"
	"github.com/wolfogre/qiniuauth/internal/log"
)

var (
	bind = flag.String("bind", ":1533","bind address, like ip:port" )
)

func main() {
	log.Logger.Info("start")
	http.ListenAndServe(":1533", handler.NewHandler())
	log.Logger.Info("stop")
	log.Logger.Sync()
}
