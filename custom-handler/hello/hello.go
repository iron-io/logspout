package hello

import (
	"log"
	"sync/atomic"

	"github.com/gliderlabs/logspout/router"
)

func New(name string) router.LogHandler {
	return &HelloHandler{Name: name}
}

type HelloHandler struct {
	Name    string
	counter uint64
}

func (a *HelloHandler) HandleLine(message *router.Message) bool {
	if message.Container.Name == "/logspout" {
		return true
	}
	atomic.AddUint64(&a.counter, 1)
	x := atomic.LoadUint64(&a.counter)
	log.Println(a.Name, x, "Hello World! ", "source:", message.Source, "cname:", message.Container.Name, "mdata:", message.Data)
	if x%5 == 0 {
		log.Println("I got this, aborting processing.")
		return false
	}
	return true
}
