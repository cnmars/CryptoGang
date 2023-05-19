package config

type System struct {
	Env          	 string `mapstructure:"env" json:"env" yaml:"env"`                                       // 环境值
	ReceptorAccount  string `mapstructure:"receptorAccount" json:"receptorAccount" yaml:"receptorAccount"`   // 收款账户
	PriceMonth       float32 `mapstructure:"priceMonth" json:"priceMonth" yaml:"priceMonth"`                              		 // 端口值
	PriceYear        float32 `mapstructure:"priceYear" json:"priceYear" yaml:"priceYear"`                              		 // 端口值
	Port          	 int    `mapstructure:"port" json:"port" yaml:"port"`                              		 // 端口值
	DbType           string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`                      	 // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType          string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"`                   	 // Oss类型
	UseMultipoint    bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"` 	 // 多点登录拦截
	UseRedis         bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`                	 // 使用redis
	LimitCountIP     int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP      int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
}
