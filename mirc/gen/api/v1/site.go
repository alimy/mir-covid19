// Code generated by go-mir. DO NOT EDIT.

package v1

import (
	"github.com/gin-gonic/gin"
)

type Site interface {
	// Chain provide handlers chain for gin
	Chain() gin.HandlersChain

	Index(*gin.Context)
	Articles(*gin.Context)
}

// RegisterSiteServant register Site servant to gin
func RegisterSiteServant(e *gin.Engine, s Site) {
	router := e.Group("v1")
	// use chain for router
	middlewares := s.Chain()
	router.Use(middlewares...)

	// register routes info to router
	router.Handle("GET", "/index/", s.Index)
	router.Handle("GET", "/articles/:category/", s.Articles)
}