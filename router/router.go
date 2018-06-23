package router

import (
	"gin-web-base-template/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

var CommonPanicHandle func(action func(c *gin.Context)) func(c *gin.Context) = controller.CommonPanicHandle

func InitRouters(version *gin.RouterGroup) {
	version.GET("/probe", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
	})
}
