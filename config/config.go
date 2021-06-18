package config

import (
	y "gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const (
	EnvConfigPath = "CHIA_CLIENT_PATH"
)

var Config = struct {
	Url string `yaml:"url"`
	SSL SSL    `yaml:"ssl"`
}{}

type SSL struct {
	CertPath string `yaml:"cert_path"`
	KeyPath  string `yaml:"key_path"`
}

func InitConfig(configPath string) {
	if configPath == "" {
		configPath = os.Getenv("CHIA_CLIENT_PATH")
	}
	if configPath == "" {
		panic("plz set config path first")
	}
	configBytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic("config file not found, see conf/app.yaml.example")
	}
	if err = y.Unmarshal(configBytes, &Config); err != nil {
		panic("load config file error")
	}
}
