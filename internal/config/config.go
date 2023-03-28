package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const (
	CfgFileNameDefault = "config"
	CfgFileTypeDefault = "json"
	CfgFilePathDefault = "config"

	CfgFileEnvName = "CONFIG_FILE_NAME"
)

func New() *viper.Viper {
	config := viper.New()

	cfgFile, ok := os.LookupEnv(CfgFileEnvName)
	if !ok {
		cfgFile = fmt.Sprintf("%s.%s", CfgFileNameDefault, CfgFileTypeDefault)
	}
	cfgFile = fmt.Sprintf("./%s", strings.Join([]string{CfgFilePathDefault, cfgFile}, "/"))
	log.Println("reading log.file:", cfgFile)
	config.SetConfigFile(cfgFile)

	err := config.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		log.Printf("fatal error config file: %s\n", err.Error())
	}

	return config
}
