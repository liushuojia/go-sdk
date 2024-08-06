package toml

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"testing"
)

// 目录下 go test -run TestConfig
func TestConfig(t *testing.T) {
	t.Log("send email")

	var config Config
	configPath := "config.toml"
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		fmt.Errorf("Error decoding TOML file: %v", err)
	}
	fmt.Println(config.App)
}
