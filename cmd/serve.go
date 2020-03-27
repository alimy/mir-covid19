// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"github.com/alimy/mir-covid19/internal/config"
	"github.com/alimy/mir-covid19/mirc/gen/api"
	"github.com/alimy/mir-covid19/mirc/gen/api/ncovh5api"
	"github.com/alimy/mir-covid19/servants"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	serveAddr string
)

func init() {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "start covid19 information service",
		Long:  "start covid19 information service",
		Run:   serveRun,
	}
	serveCmd.Flags().BoolVarP(&inDebug, "debug", "d", false, "whether in debug mode")
	serveCmd.Flags().StringVarP(&inConfigFile, "config", "c", "", "custom config file path used to init application")
	serveCmd.Flags().StringVarP(&serveAddr, "addr", "a", "", "service listen address")
	register(serveCmd)
}

func serveRun(_cmd *cobra.Command, _args []string) {
	conf := inSetup(func(conf *config.Config) {
		conf.Serve.SetAddr(serveAddr)
	})
	logrus.Infof("start with config:%s", conf)

	// register servants to gin
	e := gin.New()
	registerServants(e)

	// start servant service
	if err := e.Run(conf.Serve.Addr); err != nil {
		logrus.Fatal(err)
	}
}

func registerServants(e *gin.Engine) {
	api.RegisterPortalServant(e, servants.NewPortal())
	ncovh5api.RegisterEpidemicServant(e, servants.NewEpidemic())
}
