package config


type Ssl struct {
	KeyStr string `mapstructure:"keystr" json:"keystr" yaml:"keystr"`
	Pem    string `mapstructure:"pem" json:"pem" yaml:"pem"`
	Enable bool `mapstructure:"enable" json:"enable" yaml:"enable"`
	Domain string `mapstructure:"domain" json:"domain" yaml:"domain"`
}