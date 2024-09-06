#!/usr/bin/env bash
pushd database/sqlc
sqlc generate
popd
pushd database/migrations
go-bindata -pkg migrations .
popd
echo "Building daemon"
go build -o ./tmp/tealok-daemon daemon.go
echo "Building client"
go build -o ./tmp/tealok-client client.go
