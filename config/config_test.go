package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	InitConfig("../conf/app.yaml")
	t.Log(Config)
}
