package v1

import (
	"github.com/alimy/mir/v2"
	"github.com/alimy/mir/v2/engine"
)

func init() {
	engine.AddEntry(new(Site))
}

// Site site v1 interface info
type Site struct {
	Chain    mir.Chain `mir:"-"`
	Group    mir.Group `mir:"v1"`
	Index    mir.Get   `mir:"/index/"`
	Articles mir.Get   `mir:"/articles/:category/"`
}
