// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package assets

//go:generate go-bindata -nomemcopy -pkg=assets -ignore=README.md -prefix=../../assets -debug=false -o=assets_gen.go ../../assets/config/...
