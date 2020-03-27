// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package config

import (
	"fmt"
)

// Config application config model
type Config struct {
	Application Application
	Serve       Serve
}

// Application indicate application section config
type Application struct {
	Name        string
	Version     string
	Authors     []string
	Description string
}

// Serve indicate serve section config
type Serve struct {
	Addr string
}

// SetAddr set addr
func (c *Serve) SetAddr(addr string) {
	if addr != "" {
		c.Addr = addr
	}
}

// String return string object
func (c *Config) String() string {
	return fmt.Sprintf("{application:%s, serve:%s", &c.Application, &c.Serve)
}

// String string object
func (c *Application) String() string {
	return fmt.Sprintf("{name:%q, version:%q, authors:%v,description:%q}",
		c.Name, c.Version, c.Authors, c.Description)
}

// String string object
func (c *Serve) String() string {
	return fmt.Sprintf("{addr:%q}", c.Addr)
}
