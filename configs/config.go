package configs

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)



type (
	Config struct {
		Server ServerConfig
		Database    DatabaseConfig
		Log         LogConfig
	}

	ServerConfig struct {
		Host string
		Port int
	}

	DatabaseConfig struct {
		Host                string
		Port                int
		DbName              string
		Username            string
		Password            string
	}

	LogConfig struct {
		Level string
	}
)

var config Config

func Init() {
		// if env == "local" {
		// 	env = "local"
		// }

		viper.AddConfigPath("configs/")
		viper.SetConfigName("config-local")
		viper.SetConfigType("yaml")

		// get application config
		err := viper.ReadInConfig()
		panicOnError(err)

		err = viper.Unmarshal(&config)
		panicOnError(err)

		err = viper.Unmarshal(&config)
		panicOnError(err)

		viper.OnConfigChange(func(e fsnotify.Event) {
			viper.Unmarshal(&config)
		})
		viper.WatchConfig()

		logrus.Print("Config file read loaded successfully")
}




func Get() Config {
	return config
}

// for unit test purpose
func Set(cfg Config) {
	config = cfg
}

func panicOnError(err error) {
	if err != nil {
		log.Printf("panic on config %v", err)
		panic(err)
	}
}