package ginserver

import (
	"github.com/AmitSuresh/my-microservices/order/src/infra/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetGinServer(cfg *config.Config) *gin.Engine {
	gin.SetMode(cfg.GinMode)
	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	g.Use(cors.Default())
	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	return g
}
