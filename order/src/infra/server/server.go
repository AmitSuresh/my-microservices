package server

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/AmitSuresh/my-microservices/order/src/infra/config"
	"github.com/gin-gonic/gin"
)

func GetHttpServer(t *tls.Config, g *gin.Engine, cfg *config.Config) *http.Server {
	return &http.Server{
		Addr:         cfg.ServerAddr,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  10 * time.Second,
		Handler:      g,
		// TLSConfig:    t,
	}
}
