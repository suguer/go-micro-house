package config

import (
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

func InitConfig() {
	_, workDir, _, _ := runtime.Caller(0)
	workDir = filepath.Dir(filepath.Dir(workDir))
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
