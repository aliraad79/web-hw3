package main

import (
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	cmd := flag.String("cmd", "", "")
	flag.Parse()

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"S": cmd})
	})

	router.Run(":8081")
}
