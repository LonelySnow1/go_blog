package global

import (
	"go.uber.org/zap"
	"server/config"
)

var (
	Config *config.Config
	Log    *zap.Logger
)
