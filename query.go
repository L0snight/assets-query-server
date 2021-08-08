package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Query struct {
	IP string `form:"ip"`
}

func query(c *gin.Context)  {
	ip := c.PostForm("ip")
	log.Println(ip)
	c.String(200, ip);
}
func main()  {
	route := gin.Default()
	route.POST("/query", query)
	route.Run(":8088")
}
