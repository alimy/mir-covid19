// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"context"
	"net/http"
	"os"

	"github.com/alimy/mgo/grace"
	"github.com/alimy/mir-covid19/internal/config"
	"github.com/alimy/mir-covid19/servants"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	rt "github.com/alimy/mir-covid19/mirc/gen/api"
	ds "github.com/alimy/mir-covid19/mirc/gen/api/ncovh5api/THPneumoniaOuterDataService"
	pn "github.com/alimy/mir-covid19/mirc/gen/api/ncovh5api/THPneumoniaService"
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
	server := &http.Server{Addr: conf.Serve.Addr, Handler: e}
	go func() {
		logrus.Infof("listening and serving HTTP on %s", server.Addr)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logrus.Fatal(err)
		}
	}()

	// graceful shutdown server
	grace.Listen(func(signal os.Signal) (isContinue bool) {
		logrus.Infof("shutdown server because received signal: %s", signal)
		if err := server.Shutdown(context.Background()); err != nil {
			logrus.Fatal(err)
		}
		return false
	})
	grace.Wait()
}

func registerServants(e *gin.Engine) {
	rt.RegisterPortalServant(e, servants.NewPortal())
	pn.RegisterPneumoniaServant(e, servants.NewPneumonia())
	ds.RegisterEpidemicServant(e, servants.NewEpidemic())
}
