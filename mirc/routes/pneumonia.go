package routes

import (
	"github.com/alimy/mir/v2"
	"github.com/alimy/mir/v2/engine"
)

func init() {
	engine.AddEntry(new(Pneumonia))
}

// Pneumonia pneumonia data fetch interface
type Pneumonia struct {
	Group           mir.Group `mir:"ncovh5api/THPneumoniaService"`
	GetContents     mir.Post  `mir:"getContents"`
	GetAreaContents mir.Post  `mir:"getAreaContents"`
}
