#!/usr/bin/make
include .env

PROJECT_NAME=$(shell basename "$(PWD)")


build:
	go build -o bin/${PROJECT_NAME} cmd/server/server.go

run: build
	bin/$(PROJECT_NAME)

