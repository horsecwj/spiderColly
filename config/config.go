package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var (
	appConfig  APPConfig
	configOnce sync.Once
)

type APPConfig struct {
	Debug    bool            `json:"debug"`
	Database *DatabaseConfig `json:"database"`
	Server   *ServerConfig   `json:"server"`
}

// 数据库配置信息
type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	LogMode  bool   `json:"log_mode"`
}

// Server配置信息
type ServerConfig struct {
	Port        string `json:"port"`
	SummedCheck bool   `json:"summed_check"`
	VerifyHost  string `json:"verify_host"`
}

// 项目配置
func APPConf() *APPConfig {

	configOnce.Do(func() {

		if err := viper.Unmarshal(&appConfig); err != nil {

			fmt.Println("读取配置出错:", err)
		}
	})
	return &appConfig
}

// 数据库配置
func DBConf() *DatabaseConfig {

	return APPConf().Database
}

func RefreshConf() {

	configOnce = sync.Once{}
}
