package main

import (
	"flag"
	"net/http"

	"github.com/wolfogre/qiniu-download/internal/handler"
	"github.com/wolfogre/qiniu-download/internal/log"
	"github.com/wolfogre/qiniu-download/internal/dao"
)

var (
	bind = flag.String("bind", ":80","bind address, like ip:port" )
	addr = flag.String("addr", "","redis address" )
	pass = flag.String("pass", "","redis password" )
	cdnAddr = flag.String("cdn", "","cdn address" )
	db = flag.Int("db", 1, "redis db" )
)

func main() {
	flag.Parse()
	log.Logger.Info("start")
	err := dao.Init(*addr, *pass, *db)
	if err != nil {
		log.Logger.Panic(err)
	}
	log.Logger.Info("redis init")
	http.ListenAndServe(*bind, handler.NewHandler(*cdnAddr))
	log.Logger.Info("stop")
	log.Logger.Sync()
}
