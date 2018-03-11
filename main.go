package main

import (
	"flag"
	"net/http"
	"github.com/wolfogre/qiniuauth/internal/handler"
	"github.com/wolfogre/qiniuauth/internal/log"
	"github.com/wolfogre/qiniuauth/internal/dao"
)

var (
	bind = flag.String("bind", ":80","bind address, like ip:port" )
	addr = flag.String("addr", "","redis address" )
	pass = flag.String("pass", "","redis password" )
	db = flag.Int("db", 1,"redis db" )
)

func main() {
	flag.Parse()
	log.Logger.Info("start")
	err := dao.Init(*addr, *pass, *db)
	if err != nil {
		log.Logger.Panic(err)
	}
	log.Logger.Info("redis init")
	http.ListenAndServe(*bind, handler.NewHandler())
	log.Logger.Info("stop")
	log.Logger.Sync()
}
