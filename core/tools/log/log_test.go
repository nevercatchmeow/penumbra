package log_test

import (
	"testing"
	"time"

	"github.com/nevercatchmeow/penumbra/core/tools/log"
)

func TestLogger(t *testing.T) {
	logger := log.NewLog(log.WithLogDir("./logs", "./logs"), log.WithRunMode(log.RunModeTest, nil))
	log.SetLogger(logger)

	logger.Info("info msg")
	log.Info("info msg", log.Duration("duration", time.Hour*7))
}
