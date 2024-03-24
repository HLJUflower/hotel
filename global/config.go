package global

//所有配置对应的结构体

// App 应用环境基本配置

type apps struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`
	Port    string `mapstructure:"port" json:"port" yaml:"port"`
	AppName string `mapstructure:"app_name" json:"app_name" yaml:"app_name"`
	AppUrl  string `mapstructure:"app_url" json:"app_url" yaml:"app_url"`
}

// Configuration 全局基本配置
type Configuration struct {
	App apps `mapstructure:"api" json:"api" yaml:"api"`
}
