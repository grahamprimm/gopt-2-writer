package log

import (
	"go.uber.org/zap"
)

func NewDevConfig() zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func init() {
	logger, err := zap.NewDevelopmentConfig().Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	sugarD := logger.Sugar()
	sugarD.Info("sugar init success")

	// The first call to Named sets the Logger name.
	sugarD.Named("dev")

	// Additional calls to Named create a period-separated path.
	//sugarD.Named("subpackage").Info("sub-logger")
}
