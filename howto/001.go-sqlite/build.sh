#!/bin/bash
# === builds current step only

gofmt -w main.go
go run main.go
