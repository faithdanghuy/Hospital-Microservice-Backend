SHELL := /bin/bash

export GOROOT := /c/Program Files/Go
export GOPATH := /c/Users/Nguyen Hi Huu/go
export GOMODCACHE := $(GOPATH)/pkg/mod
export GOCACHE:= /c/Users/Nguyen Hi Huu/AppData/Local/go-build
export GOBIN := $(GOPATH)/bin
export PATH := $(GOROOT)/bin:$(GOBIN):$(PATH)
export TEMP := /c/Users/Nguyen Hi Huu/AppData/Local/Temp
export TMP := /c/Users/Nguyen Hi Huu/AppData/Local/Temp

user:
	go run user-service/cmd/app/main.go
appointment:
	go run appointment-service/cmd/app/main.go
prescription:
	go run prescription-service/cmd/app/main.go