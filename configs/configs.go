package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
)

func Init(dir string, logger *zap.Logger) Yamlcfg {
	viper.SetConfigName("config")    // name of config file (without extension)
	viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(dir)         // path to look for the config file in
	viper.AddConfigPath("./configs") // optionally look for config in the working directory
	viper.AddConfigPath("./")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Print(err)
		logger.Fatal(fmt.Sprint("fatal error config file: %w", err))
	}
	var yamlcfg Yamlcfg
	err = viper.Unmarshal(&yamlcfg)
	if err != nil {
		log.Print(fmt.Sprintf("Ошибка маршалинга %v", err))
		logger.Fatal(fmt.Sprintf("Ошибка маршалинга %v", err))
	}
	return yamlcfg
}
