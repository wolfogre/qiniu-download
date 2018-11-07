package main
//
//import (
//	"flag"
//	"net/http"
//
//	"github.com/wolfogre/qiniu-download/internal/handler"
//	"github.com/wolfogre/qiniu-download/internal/log"
//)
//
//var (
//	bind = flag.String("bind", ":80","bind address, like ip:port" )
//	cdn = flag.String("cdn", "","cdn address" )
//)
//
//func main() {
//	flag.Parse()
//	log.Logger.Info("start")
//	log.Logger.Info("listen on %v, cdn is %v", bind, cdn)
//	http.ListenAndServe(*bind, handler.NewHandler(*cdn))
//	log.Logger.Info("stop")
//	log.Logger.Sync()
//}

import (
	"log"
	"sync"

	"github.com/judwhite/go-svc/svc"
)

// program implements svc.Service
type program struct {
	wg   sync.WaitGroup
	quit chan bool
}

func main() {
	if err := svc.Run(&program{}); err != nil {
		panic(err)
	}
}

func (p *program) Init(env svc.Environment) error {
	log.Printf("is win service? %v\n", env.IsWindowsService())
	return nil
}

func (p *program) Start() error {
	// The Start method must not block, or Windows may assume your service failed
	// to start. Launch a Goroutine here to do something interesting/blocking.

	p.quit = make(chan bool)

	p.wg.Add(1)
	go func() {
		log.Println("Starting...")
		<-p.quit
		log.Println("Quit signal received...")
		p.wg.Done()
	}()

	return nil
}

func (p *program) Stop() error {
	// The Stop method is invoked by stopping the Windows service, or by pressing Ctrl+C on the console.
	// This method may block, but it's a good idea to finish quickly or your process may be killed by
	// Windows during a shutdown/reboot. As a general rule you shouldn't rely on graceful shutdown.

	log.Println("Stopping...")
	close(p.quit)
	p.wg.Wait()
	log.Println("Stopped.")
	return nil
}