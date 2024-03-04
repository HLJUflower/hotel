package config

import "hotel/config/app"

//所有配置对应的结构体

// Configuration 全局基本配置
type Configuration struct {
	App app.App `mapstructure:"app" json:"app" yaml:"app"`
}
