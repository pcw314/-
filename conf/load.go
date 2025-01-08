package conf

import (
	"errors"
	"gitee.com/xygfm/authorization/logger"
	"gopkg.in/yaml.v3"
	"os"
)

var config *Config

func GetConfig() *Config {
	if config == nil {
		logger.Error(errors.New("请先加载配置"))
	}
	return config
}

func LoadConfigFromYaml(path string) (*Config, error) {
	defaultConfig := NewDefaultConfig()
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, defaultConfig)
	if err != nil {
		panic(err)
	}
	config = defaultConfig
	logger.Info("加载yaml配置成功")
	return defaultConfig, nil
}
