// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

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
