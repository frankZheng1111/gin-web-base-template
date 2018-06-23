package main

import (
	"gin-web-base-template/config"
	db "gin-web-base-template/model"
	"gin-web-base-template/router"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	if config.Config.ENV == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	defer db.DB.Close()
	ginRouter := gin.Default()

	store, err := redis.NewStore(config.Config.Redis.MaxIdle, config.Config.Redis.Network, config.Config.Redis.Address, config.Config.Redis.Password, []byte(config.Config.Redis.AuthKey))
	if err != nil {
		panic(err)
	}

	ginRouter.Use(sessions.Sessions("face-match-servie", store))

	v1 := ginRouter.Group("/v1")
	router.InitRouters(v1)

	ginRouter.Run(fmt.Sprintf(":%d", config.Config.Server.Port))
}
