# Tealok

Tealok is sort of like a container orchestration system (Kubernetes, Docker Swarm, Mesos), but probably not like you'd expect if you're familiar with them. It's not meant to scale a large service over many compute nodes. It's meant to be extremely easy to use. Specifically something a random non-technical user could profitably use.

Tealok is a container orchestration system for people who don't want to have to know what a container is, why you would orchestrate it, of why it would require a system.

Tealok ships with an append-only database to ensure the ability to rollback. It has service discovery mechanisms, integrations with DNS provider APIs, built-in backup integration with third-party providers, the ability to start-up and manage a certificate authority, and the ability to run in a distributed mode for fault-tolerance.

It is, essentially, a very ambitious project.

We do crazy things like one-click installs to set up your own self-hosted email service (that runs JMAP, because it's better).

Or one-click installs to use a transparent proxy to filter all web traffic in your family's home.

We're building it because we think that sufficiently advanced technology is indistinguishable from magic and were tired of being wizards that just conjure money out of attention for Corporations of Unusual Size.

## Technical Details

Tealok is written in Go. That's not dogmatic, it just seems like Go was the fastest way to build the thing we want to build.

It uses Podman to actually run the containers. It uses SQLite for the append-only configuration database. Backups are handled by Backblaze B2. Certificate authority is from Step CA.

## Contributing

The project is very early, we're building the prototype now. You can [join the mailing list](https://mailing-list.tealok.tech/subscription/form).

Feel free to create issues and send pull-requests.

### Incantations

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

## FAQ

### Are you serious?

Yes.

### How much funding do you have?

Only what we've put up from out personal savings.

### So you're bootstrapping.

Yes. Also, not a question.

### Are you stupid? Spending other people's money is awesome.

Probably.

Spending other people's money also makes you beholden to them. We fight for the users.

### Can this possible work?

Absolutely, if we work together and keep at it. Archimedes said with a large enough Go codebase we can move the whole world. He was smart.

### How will you make money?

Eventually we want all of us who contribute to this project to make money. We'll make it be selling support and selling development prioritization. The goal is not to build a walled garden around a cathedral, but a fusion-powered space-bazaar than can accommodate everyone.

### How much of it works at this point?

Eh, like 10%. See our roadmap.
