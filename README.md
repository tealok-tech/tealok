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
