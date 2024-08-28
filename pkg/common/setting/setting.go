package setting

import (
	"log/slog"
	"sync/atomic"

	"github.com/spf13/viper"
)

var gConfig atomic.Value

type Config struct {
	ServiceName string
	Port        int
}

func GlobalConfig() *Config {
	return gConfig.Load().(*Config)
}

func SetGlobalConfig(c *Config) {
	gConfig.Store(c)
}

func InitGlobalConfig(confFilePath string) error {
	cfg := &Config{}
	v := viper.New()
	v.SetDefault("ServiceName", "echo")
	v.SetDefault("Port", "45505")
	v.SetConfigFile(confFilePath)
	if err := v.ReadInConfig();err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&cfg);err != nil {
		panic(err)
	}
	SetGlobalConfig(cfg)
	slog.Info("config: %v", cfg)
	return nil
}
