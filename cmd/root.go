// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"log"

	"github.com/alimy/mir-covid19/internal/config"
	"github.com/alimy/mir-covid19/internal/logus"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	inConfigFile string // 自定义配置文件路径
	inDebug      bool   // 是否debug模式
)

var (
	rootCmd = &cobra.Command{
		Use:   "covid",
		Short: "covid19 information service",
		Long:  `covid19 information service`,
	}
)

// Setup set root command name,short-describe, long-describe
// return &cobra.Command to custom other options
func Setup(use, short, long string) *cobra.Command {
	rootCmd.Use = use
	rootCmd.Short = short
	rootCmd.Long = long
	return rootCmd
}

// register add sub-command
func register(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

// Execute start application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func inSetup(postInit func(*config.Config)) *config.Config {
	conf := config.InitFrom(inConfigFile)
	inDev := false
	if conf.Application.RunMode == "develop" || inDebug {
		inDev = true
	}
	if inDev {
		logus.SetLevel(logus.LevelDebug)
		gin.SetMode(gin.DebugMode)
	} else {
		logus.SetLevel(logus.LevelInfo)
		gin.SetMode(gin.ReleaseMode)
	}
	postInit(conf)
	return conf
}
