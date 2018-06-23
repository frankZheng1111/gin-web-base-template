package controller

import (
	db "gin-web-base-template/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func CommonPanicHandle(action func(c *gin.Context)) func(c *gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Error: ", time.Now(), err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "SERVER_ERROR",
				})
				c.Abort()
			}
		}()
		action(c)
	}
}

func ValidationErrorResponse(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"msg": "ParamsValidationError",
	})
	return
}

func MissingLockErrorResponse(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"msg": "OverFrequency",
	})
	return
}

func ParsePaginateFromQuery(c *gin.Context) *db.Paginate {
	paginate := new(db.Paginate)
	page := c.Query("page")
	pageSize := c.Query("page_size")
	paginate.Page, _ = strconv.Atoi(page)
	paginate.PageSize, _ = strconv.Atoi(pageSize)
	return paginate
}

func SetSession(c *gin.Context, key string, value interface{}) {
	session := sessions.Default(c)
	session.Set(key, value)
	session.Save()
}

func GetSession(c *gin.Context, key string) interface{} {
	session := sessions.Default(c)
	return session.Get(key)
}
