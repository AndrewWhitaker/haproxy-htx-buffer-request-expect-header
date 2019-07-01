package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func respond(c *gin.Context) {
	hostname, _ := os.Hostname()

	c.JSON(200, gin.H{
		"hostname":       hostname,
		"gin_client_ip":  c.ClientIP(),
		"headers":        c.Request.Header,
		"host":           c.Request.Host,
		"remote_addr":    c.Request.RemoteAddr,
		"request_uri":    c.Request.RequestURI,
		"close":          c.Request.Close,
		"content_length": c.Request.ContentLength,
	})
}

func sleep(c *gin.Context) {
	d, _ := time.ParseDuration("1s")
	time.Sleep(d)
	respond(c)
}

func main() {
	r := gin.New()
	r.Use(gin.Recovery())

	r.POST("/", sleep)

	r.Run()
}
