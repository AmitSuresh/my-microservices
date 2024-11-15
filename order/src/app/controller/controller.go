package controller

import (
	"net/http"

	customprometheus "github.com/AmitSuresh/my-microservices/order/src/infra/prometheus" // import customprometheus
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func InitReq(l *zap.Logger, ge *gin.Engine) {
	ge.Use(func(c *gin.Context) {
		customprometheus.HttpRequests.WithLabelValues(c.Request.Method, c.Request.RequestURI).Inc()
		c.Next()
	})
	gr := ge.Group("/api")
	gr.GET("/ping", func(c *gin.Context) {
		customprometheus.HttpRequests.WithLabelValues(c.Request.Method, c.Request.RequestURI).Inc()
		timer := prometheus.NewTimer(customprometheus.HttpDuration.WithLabelValues(c.Request.Method, c.Request.RequestURI))
		defer timer.ObserveDuration()
		c.Set("timer", timer)

		l.Info("pong")
		c.JSON(http.StatusOK, gin.H{"success_message": "pong"})
	})

	ge.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
