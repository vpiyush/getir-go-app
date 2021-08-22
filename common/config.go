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

var Cfg Config

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})
	configPath := "common/config.yaml"
	log.Debug("Reading config from ", configPath)
	err := cleanenv.ReadConfig(configPath, &Cfg)
	if err != nil {
		log.Fatal("couldn't Read app configuration")
	}
}
