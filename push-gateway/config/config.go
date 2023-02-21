package config

import (
	"fmt"
	"push-gateway/pkg/utils"

	"github.com/spf13/viper"
)

func InitConfig() {

	workDir := utils.WorkDir()
	fmt.Printf("workDir: %v\n", workDir)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
