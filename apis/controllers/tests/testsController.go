package tests_controller

import (
	"github.com/gin-gonic/gin"
)

func Add(n int, m int)int{
	return n + m
}

func Test(c *gin.Context) {
	c.JSON(200, "test controller is working fine")
}
