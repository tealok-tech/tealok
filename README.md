# Tealok

A system for managing containers for personal applications.

## Incantations

Start the podman API service:

```
sudo podman system service -t 0
```

Build the Tealok binary

```
nix-shell
go build main.go
```

Run the binary as root so it can talk to the rootful service

```
sudo ./main
```

### Continuous Build-and-test

Install [air]():

```
go install github.com/air-verse/air@latest
```

Run air

```
~/go/bin/air
```

This will run `go run` constantly

### Generate new database migration

```
cd database/migrations
go-bindata -pkg migrations .
```
