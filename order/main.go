package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AmitSuresh/my-microservices/order/src/app/controller"
	"github.com/AmitSuresh/my-microservices/order/src/infra/cloudwatch"
	gconfig "github.com/AmitSuresh/my-microservices/order/src/infra/config"
	"github.com/AmitSuresh/my-microservices/order/src/infra/ginserver"
	customlogger "github.com/AmitSuresh/my-microservices/order/src/infra/logger"
	customprometheus "github.com/AmitSuresh/my-microservices/order/src/infra/prometheus"
	hs "github.com/AmitSuresh/my-microservices/order/src/infra/server"
	//"github.com/AmitSuresh/my-microservices/order/src/infra/tlsconfig"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	//"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	//"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	//"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

var (
	cfg          *gconfig.Config
	l            *zap.Logger
	ge           *gin.Engine
	httpServer   *http.Server
	srvStartTime time.Time
)

func init() {
	srvStartTime = time.Now()
	cfg = gconfig.LoadConfig()

	awsConfig, err := config.LoadDefaultConfig(context.Background(), config.WithRegion("ap-southeast-1"))
	if err != nil {
		panic(err)
	}
	cwWriter, err := cloudwatch.NewCloudWatchWriter(context.Background(), cfg.LogGroupName, cfg.LogStreamName, awsConfig)
	if err != nil {
		panic(err)
	}

	l = customlogger.GetLogger(cwWriter)

	customprometheus.InitializePrometheus()
	ge = ginserver.GetGinServer(cfg)

	controller.InitReq(l, ge)

	httpServer = hs.GetHttpServer(nil, ge, cfg)
}

func main() {
	defer func() {
		_ = l.Sync()
	}()
	go func() {
		for {
			time.Sleep(1 * time.Second)
			customprometheus.Uptime.Set(time.Since(srvStartTime).Seconds())
		}
	}()
	l.Info("here", zap.Any("cfg", cfg))
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Fatal("Server failed to start", zap.Error(err))
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		l.Error("Server shutdown failed", zap.Error(err))
	} else {
		l.Info("Server gracefully stopped")
	}
}
