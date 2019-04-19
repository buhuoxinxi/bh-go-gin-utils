package main

import (
	_ "github.com/buhuoxinxi/bh-go-gin-utils/testdata"

	"github.com/buhuoxinxi/bh-go-gin-utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// new server
	engine := bhginutils.NewServer()

	// hand
	var handler = func(c *gin.Context) {
		c.JSON(200, gin.H{
			"url":     c.Request.URL.Path,
			"message": "pong",
		})
	}

	// route
	// curl -k https://127.0.0.1:8099/ping
	// curl -k http://127.0.0.1:8099/ping
	bhginutils.RegisterRoutes(engine, []*bhginutils.Route{
		bhginutils.NewRoute("GET", "ping", handler),
	})

	// route group
	// curl -k https://127.0.0.1:8099/v1/ping
	// curl -k http://127.0.0.1:8099/v1/ping
	bhginutils.RegisterRouteGroup(engine, "v1", []*bhginutils.Route{
		bhginutils.NewRoute("GET", "ping", handler),
	})

	// start
	bhginutils.RunServer(engine)
}
