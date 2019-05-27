#!/usr/bin/env bash
set -e

go test -coverprofile=cover.out -coverpkg=github.com/Ezian/gof-art/channel,github.com/Ezian/gof-art/mutex,github.com/Ezian/gof-art/naive,github.com/Ezian/gof-art/utils
go tool cover -html=cover.out -o cover.html 
browse cover.html