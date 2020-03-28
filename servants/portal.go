// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"net/http"

	"github.com/alimy/mir-covid19/mirc/gen/api"
	"github.com/alimy/mir-covid19/web"
	"github.com/gin-gonic/gin"
)

// portal implement api.portal
type portal struct {
	staticHandler http.Handler
}

func (p *portal) Index(c *gin.Context) {
	c.Status(http.StatusOK)
	c.Writer.Write(web.MustAsset("index.html"))
}

func (p *portal) GetMainAssets(c *gin.Context) {
	c.Status(http.StatusOK)
	c.Writer.Write(web.MustAsset("index.html"))
}

func (p *portal) GetStaticAssets(c *gin.Context) {
	p.staticHandler.ServeHTTP(c.Writer, c.Request)
}

func (p *portal) HeadStaticAssets(c *gin.Context) {
	p.staticHandler.ServeHTTP(c.Writer, c.Request)
}

// NewPortal return a Portal instance
func NewPortal() api.Portal {
	fs := web.NewFileSystem()
	return &portal{
		staticHandler: http.StripPrefix("/static", http.FileServer(fs)),
	}
}
