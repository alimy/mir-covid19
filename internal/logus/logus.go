// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package logus

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// log level
const (
	LevelTrace = "trace"
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"
	LevelFatal = "fatal"
	LevelPanic = "panic"
)

const (
	defaultLogLevel = LevelDebug
)

// SetLevel set log level
func SetLevel(level string) {
	l, e := logrus.ParseLevel(level)
	if e != nil {
		fmt.Printf("invalid log level so use default: %s", e)
		l, _ = logrus.ParseLevel(defaultLogLevel)
	}
	logrus.SetLevel(l)
}
