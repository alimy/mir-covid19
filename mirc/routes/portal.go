// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package routes

import (
	"github.com/alimy/mir/v2"
	"github.com/alimy/mir/v2/engine"
)

func init() {
	engine.AddEntry(new(Portal))
}

// Portal web ui interface handler
type Portal struct {
	Index            mir.Get  `mir:"/"`
	GetMainAssets    mir.Get  `mir:"/index.html"`
	GetStaticAssets  mir.Get  `mir:"/static/*filepath"`
	HeadStaticAssets mir.Head `mir:"/static/*filepath"`
}
