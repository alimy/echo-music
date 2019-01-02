package config

import (
	_ "github.com/BurntSushi/toml"
)

const (
	defaultConfigPath = "assets/config/app.toml"
)

var CustomConfigPath = ""
