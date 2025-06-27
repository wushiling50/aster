package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/wushiling50/aster/api/internal/config"
	"github.com/wushiling50/aster/api/internal/handler"
	"github.com/wushiling50/aster/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "api/etc/applet.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(func(header http.Header) {
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		header.Set("Access-Control-Allow-Headers", "*")
		header.Set("Access-Control-Expose-Headers", "Content-Length")
		header.Set("Access-Control-Max-Age", "43200")
	}, func(w http.ResponseWriter) {}))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
