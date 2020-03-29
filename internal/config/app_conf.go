// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package config

import (
	"fmt"
	"net/url"
	"strings"
)

// Config application config model
type Config struct {
	Application Application
	Serve       Serve
	Database    Database
	Redis       Redis
	Develop     Develop
}

// Application indicate application section config
type Application struct {
	Name        string
	Version     string
	Authors     []string
	Description string
	RunMode     string `toml:"run_mode"`
}

// Serve indicate serve section config
type Serve struct {
	Addr string
}

// Database indicate database section config
type Database struct {
	Type     string
	Host     string
	Name     string
	User     string
	Password string
	Param    string
	SslMode  string `toml:"ssl_mode"`
	Path     string
}

// Redis indicate redis section config
type Redis struct {
	Addr     string
	Password string
	DB       int `toml:"db"`
}

// Develop indicate develop section config
type Develop struct {
	DatawareFake bool `toml:"dataware_fake"`
	CacheFake    bool `toml:"cache_fake"`
}

// SetAddr set addr
func (c Serve) SetAddr(addr string) {
	if addr != "" {
		c.Addr = addr
	}
}

// String return string object
func (c Config) String() string {
	return fmt.Sprintf("{application:%s, serve:%s, database:%s, develop:%s}", c.Application, c.Serve, c.Database, c.Develop)
}

// String string object
func (c Application) String() string {
	return fmt.Sprintf("{name:%q, version:%q, authors:%v,description:%q, run_mode:%q}",
		c.Name, c.Version, c.Authors, c.Description, c.RunMode)
}

// String string object
func (c Serve) String() string {
	return fmt.Sprintf("{addr:%q}", c.Addr)
}

// String string object
func (c Database) String() string {
	return fmt.Sprintf("{type:%q, host:%q, name:%q, user:%q, password:%q, param: %q, ssl_mode:%q, path:%q}",
		c.Type, c.Host, c.Name, c.User, c.Password, c.Param, c.SslMode, c.Path)
}

// String string object
func (c Redis) String() string {
	return fmt.Sprintf("{addr:%q, password:%q, db:%d}", c.Addr, c.Password, c.DB)
}

// String string object
func (c Develop) String() string {
	return fmt.Sprintf("{dataware_fake:%t, cache_fake:%t}", c.DatawareFake, c.CacheFake)
}

// Dsn return database type and DSN(Database Source Name)
func (c Database) Dsn() (string, string) {
	var dialect, dsn string
	switch c.Type {
	case "mysql":
		dialect, dsn = "mysql", c.mysqlDsn()
	case "postgres":
		dialect, dsn = "postgres", c.postgreDsn()
	case "sqlite3":
		dialect, dsn = "sqlite3", c.Path
	default:
		dialect, dsn = "mysql", c.mysqlDsn()
	}
	return dialect, dsn
}

func (c *Database) mysqlDsn() string {
	var param = "?"
	if c.Param != "" {
		param = param + c.Param + "&"
	}
	var dsn string
	if c.Host[0] == '/' { // use unix socket
		dsn = fmt.Sprintf("%s:%s@unix(%s)/%s%scharset=utf8mb4&parseTime=true",
			c.User, c.Password, c.Host, c.Name, param)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s%scharset=utf8mb4&parseTime=true",
			c.User, c.Password, c.Host, c.Name, param)
	}
	return dsn
}

func (c Database) postgreDsn() string {
	host, port := "127.0.0.1", "5432"
	if strings.Contains(c.Host, ":") && !strings.HasSuffix(c.Host, "]") {
		idx := strings.LastIndex(c.Host, ":")
		host = c.Host[:idx]
		port = c.Host[idx+1:]
	} else if len(c.Host) > 0 {
		host = c.Host
	}
	var param = "?"
	if c.Param != "" {
		param = param + c.Param + "&"
	}
	var dsn string
	if host[0] == '/' { // use unix socket
		dsn = fmt.Sprintf("postgres://%s:%s@:%s/%s%ssslmode=%s&host=%s",
			url.QueryEscape(c.User), url.QueryEscape(c.Password), port, c.Name, param, c.SslMode, host)
	} else {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s%ssslmode=%s",
			url.QueryEscape(c.User), url.QueryEscape(c.Password), host, port, c.Name, param, c.SslMode)
	}
	return dsn
}
