// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package routes

import (
	. "github.com/alimy/mir/v2"
	. "github.com/alimy/mir/v2/engine"
)

func init() {
	AddEntry(new(Pneumonia))
}

// Pneumonia pneumonia data fetch interface
type Pneumonia struct {
	Group           Group `mir:"ncovh5api/THPneumoniaService"`
	GetContents     Post  `mir:"getContents"`
	GetAreaContents Post  `mir:"getAreaContents"`
}
