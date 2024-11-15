package logger

import (
	"os"

	"github.com/AmitSuresh/my-microservices/order/src/infra/cloudwatch"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetLogger(cwLogger *cloudwatch.CloudWatchWriter) (logger *zap.Logger) {
	// Configure Zap to write to both CloudWatch and standard output
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stdout"}      // Write to stdout
	cfg.ErrorOutputPaths = []string{"stderr"} // Write to stderr

	// Modify core settings to include CloudWatch logger
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Format timestamps if needed

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig), // Use JSON encoder from cfg
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout), // Write to stdout
			zapcore.AddSync(cwLogger),  // Write to CloudWatch
		),
		zap.NewAtomicLevel(), // Set default log level
	)

	// Build the logger using cfg and set it globally
	logger = zap.New(core)
	// defer logger.Sync() // Ensure buffered logs are flushed on exit
	return
}
