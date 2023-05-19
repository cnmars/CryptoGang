package service

import (
	"errors"
	"golddigger/app/system/models"
	"golddigger/app/system/models/config"
	"golddigger/common/jobs"
	"golddigger/global"
	"golddigger/utils"

	"go.uber.org/zap"
)

//@description: 读取配置文件
//@return: conf config.Server, err error

type SystemConfigService struct{}

//@description: 读取配置文件
//@return: conf config.Server, err error

func (systemConfigService *SystemConfigService) GetSystemConfig() (conf config.ConfigCollection, err error) {
	return global.GD_CONFIG, nil
}

//@description: 设置配置文件
//@param: system model.System
//@return: err error

func (systemConfigService *SystemConfigService) SetSystemConfig(system models.System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.GD_VP.Set(k, v)
	}
	err = global.GD_VP.WriteConfig()
	return err
}

// @description: 获取服务器信息
// @return: server *utils.Server, err error
type UpgradePrice struct {
	PriceUnit  string  `json:"unit"`
	Receiptor  string  `json:"receiptor"`
	PriceMonth float32 `json:"priceMonth"`
	PriceYear  float32 `json:"priceYear"`
}

func (systemConfigService *SystemConfigService) GetServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		global.GD_LOG.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Ram, err = utils.InitRAM(); err != nil {
		global.GD_LOG.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		global.GD_LOG.Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}

func (systemConfigService *SystemConfigService) GetUpgradePrice() (price *UpgradePrice, err error) {
	data, err := jobs.FetchEthPriceOnce()
	if global.GD_CONFIG.System.ReceptorAccount == "" {
		return nil, errors.New("系统未设置收款账户")
	}
	if global.GD_CONFIG.System.PriceMonth < 0.1 || global.GD_CONFIG.System.PriceYear < 0.1 {
		return nil, errors.New("系统未设置价格")
	}
	if err == nil {
		//美金换算成ETH后用eth支付
		price = &UpgradePrice{
			PriceUnit:  "eth",
			Receiptor:  global.GD_CONFIG.System.ReceptorAccount,
			PriceMonth: global.GD_CONFIG.System.PriceMonth / data.Price,
			PriceYear:  global.GD_CONFIG.System.PriceYear / data.Price,
		}

	}
	return price, err
}
