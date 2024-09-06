#!/usr/bin/env bash
pushd database/sqlc
sqlc generate
popd
pushd database/migrations
go-bindata -pkg migrations .
popd
go build -o ./tmp/tealok-daemon daemon.go
go build -o ./tmp/tealok-client client.go
