// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package routes

import (
	. "github.com/alimy/mir/v2"
	. "github.com/alimy/mir/v2/engine"
)

func init() {
	AddEntry(new(Portal))
}

// Portal web ui interface handler
type Portal struct {
	Index            Get  `mir:"/"`
	GetMainAssets    Get  `mir:"/index.html"`
	GetStaticAssets  Get  `mir:"/static/*filepath"`
	HeadStaticAssets Head `mir:"/static/*filepath"`
}
