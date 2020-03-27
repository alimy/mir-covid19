package v2

import (
	"github.com/alimy/mir/v2"
	"github.com/alimy/mir/v2/engine"
)

func init() {
	engine.AddEntry(new(Site))
}

// Site site v2 interface info
type Site struct {
	Group    mir.Group `mir:"v2"`
	Index    mir.Get   `mir:"/index/"`
	Articles mir.Get   `mir:"/articles/:category/"`
	Category mir.Get   `mir:"/category/"`
}