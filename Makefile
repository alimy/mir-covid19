GOFMT ?= gofmt -s -w
GOFILES := $(shell find . -name "*.go" -type f)
ASSETS_DATA_FILES := $(shell find assets | sed 's/  /\\ /g' | xargs -0)
WEBUI_DATA_FILES := $(shell find web | sed 's/  /\\ /g' | xargs -0)

LDFLAGS += -X "github.com/alimy/mir-covid19/version.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S %Z')"
LDFLAGS += -X "github.com/alimy/mir-covid19/version.GitHash=$(shell git rev-parse HEAD)"

.PHONY: default
default: serve

.PHONY: build
build: fmt
	go build -ldflags '$(LDFLAGS)' -o covid main.go

.PHONY: build
serve: fmt
	go run -ldflags '$(LDFLAGS)' main.go serve

.PHONY: gen-assets
gen-assets: $(ASSETS_DATA_FILES)
	-rm -f internal/assets/assets_gen.go
	go generate internal/assets/assets.go
	@$(GOFMT) internal/assets

.PHONY: gen-webui
gen-webui: $(WEBUI_DATA_FILES)
	-rm -f web/web_gen.go
	go generate web/web.go
	@$(GOFMT) web

.PHONY: generate
generate:
	@cd mirc
	@-rm -rf gen
	@go generate mirc/main.go
	@cd ..
	@$(GOFMT) ./

.PHONY: fmt
fmt:
	$(GOFMT) $(GOFILES)
