#!/usr/bin/env bash
set -e

go test -coverprofile=cover.out -coverpkg=./...
go tool cover -html=cover.out -o cover.html 
browse cover.html