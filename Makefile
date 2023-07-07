BINARY=waster
VERSION=1.0.0
DATE=`date +%FT%T%z`
.PHONY: build deploy upgrade

default:
	@echo ${BINARY}
	@echo ${VERSION}
	@echo ${DATE}

build:
	@GOOS=linux GOARCH=amd64 go build -o build/${BINARY}
	@echo "[ok] build"

upgrade:
	@go-mod-upgrade