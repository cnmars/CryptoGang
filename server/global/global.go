package global

import (
	object "golddigger/app/system/models/config"
	"golddigger/common/models"
	"golddigger/common/wsserver"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

const (
	// Version go-admin version info
	Version = "1.0.0"
)

var (
	GD_CONFIG object.ConfigCollection
	GD_DB     *gorm.DB
	GD_REDIS  *redis.Client
	GD_VP     *viper.Viper
	// GD_LOG    *oplogging.Logger
	GD_LOG   *zap.Logger
	GD_ENGIN *gin.Engine

	BlackCache             local_cache.Cache
	GD_Concurrency_Control = &singleflight.Group{}
	GD_WS                  *wsserver.PubSub
	GD_PROVIDERS           models.ChainProviders
)
