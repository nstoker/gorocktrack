#!/usr/bin/make
include .env

PROJECT_NAME=$(shell basename "$(PWD)")

build:
	go build -o bin/${PROJECT_NAME} main.go

run: build
	bin/$(PROJECT_NAME)

