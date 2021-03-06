// Package Common provides configuration and constants to be used commonly accross
// all packages
package common

import (
	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
)

// Config Configration
type Config struct {
	Server struct {
		Port string `yaml:"port" env:"SERVER_PORT" env-default:"9999"`
		Host string `yaml:"host" env:"SERVER_HOST" env-default:""`
	} `yaml:"server"`

	Database struct {
		Uri        string `yaml:"uri" env:"DB_URI" env-default:"mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getircase-study?retryWrites=true"`
		Name       string `yaml:"name" env:"DB_NAME" env-default:"getir-case-study"`
		Collection string `yaml:"port" env:"DB_COLLECTION" env-default:"records"`
	} `yaml:"database"`
}

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Dir(filename)
}

// Cfg Exported Config access object
var Cfg Config

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	var configPath string
	// setup config file path for dev/production environment
	if os.Getenv("DEV") != "" {
		configPath = getCurrentPath() + "/config.yaml"
	} else {
		configPath = "common/config.yaml"
	}
	log.Debug("Reading config from ", configPath)
	err := cleanenv.ReadConfig(configPath, &Cfg)
	if err != nil {
		log.Fatal("couldn't Read app configuration")
	}
	// update port if set from env, needed for heroku deployment
	if port := os.Getenv("PORT"); port != "" {
		Cfg.Server.Port = port
	}
	if host := os.Getenv("HOST"); host != "" {
		Cfg.Server.Host = host
	}
}
