// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package config

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/alimy/mir-covid19/internal/assets"
	"github.com/sirupsen/logrus"
)

var (
	config *Config
)

// InitFrom initialize models with custom config file
func InitFrom(path string) *Config {
	config = &Config{}
	// init config
	if err := loadConfig(config); err != nil {
		logrus.Error("load config error", err)
	}
	if path == "" {
		// Empty
	} else if fileIsExist(path) { // set config from custom config file
		customConfig(config, path)
	} else {
		logrus.Infof("custom config file is not exist so use default config path: %s", path)
	}
	return config
}

// MyConfig return application's config
func MyConfig() *Config {
	return config
}

// fileIsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func fileIsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func loadConfig(config *Config) error {
	data := assets.MustAsset("config/app.toml")
	_, err := toml.Decode(string(data), config)
	return err
}

func customConfig(config *Config, path string) {
	customConfig := &Config{}
	meta, err := toml.DecodeFile(path, customConfig)
	if err != nil {
		logrus.Errorf("decode custom config error %v", err)
		return
	}
	for _, key := range meta.Keys() {
		if len(key) == 1 { // top section just continue
			continue
		}
		switch key[0] {
		case "application":
			myAppSection(config, customConfig, key)
		case "serve":
			myServeSection(config, customConfig, key)
		case "database":
			myDatabaseSection(config, customConfig, key)
		case "develop":
			myDevelopSection(config, customConfig, key)
		}
	}
}

func myAppSection(config *Config, custom *Config, key toml.Key) {
	switch key[1] {
	case "name":
		config.Application.Name = custom.Application.Name
	case "version":
		config.Application.Version = custom.Application.Version
	case "authors":
		config.Application.Authors = custom.Application.Authors
	case "description":
		config.Application.Description = custom.Application.Description
	case "run_mode":
		config.Application.RunMode = custom.Application.RunMode
	}
}

func myServeSection(config *Config, custom *Config, key toml.Key) {
	switch key[1] {
	case "addr":
		config.Serve.Addr = custom.Serve.Addr
	}
}

func myDatabaseSection(config *Config, custom *Config, key toml.Key) {
	switch key[1] {
	case "type":
		config.Database.Type = custom.Database.Type
	case "host":
		config.Database.Host = custom.Database.Host
	case "name":
		config.Database.Name = custom.Database.Name
	case "user":
		config.Database.User = custom.Database.User
	case "password":
		config.Database.Password = custom.Database.Password
	case "param":
		config.Database.Param = custom.Database.Param
	case "sslMode":
		config.Database.SslMode = custom.Database.SslMode
	case "path":
		config.Database.Path = custom.Database.Path
	}
}

func myDevelopSection(config *Config, custom *Config, key toml.Key) {
	switch key[1] {
	case "dataware_fake":
		config.Develop.DatawareFake = custom.Develop.DatawareFake
	}
}
