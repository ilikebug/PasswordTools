package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	c := gin.Default()
	SetRoute(&c.RouterGroup)
	c.Run("0.0.0.0:8080")
}

