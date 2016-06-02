package main

import (
	// These are all the regular ones included in logspout
	_ "github.com/gliderlabs/logspout/adapters/raw"
	_ "github.com/gliderlabs/logspout/adapters/syslog"
	_ "github.com/gliderlabs/logspout/httpstream"
	"github.com/gliderlabs/logspout/router"
	_ "github.com/gliderlabs/logspout/routesapi"
	_ "github.com/gliderlabs/logspout/transports/tcp"
	_ "github.com/gliderlabs/logspout/transports/tls"
	_ "github.com/gliderlabs/logspout/transports/udp"

	// This is the one we're adding.
	hello "github.com/gliderlabs/logspout/custom-handler/hello"
)

func init() {
	router.Handlers.Add(hello.New("foo"))
	router.Handlers.Add(hello.New("bar"))
	// router.Routes.AddFromUri("syslog://logs4.papertrailapp.com:25630")
}
