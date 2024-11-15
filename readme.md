# CLI-TypeRacer

**CLI-TypeRacer** is a lightweight, terminal-based typing game designed for developers to have fun and compete with friends directly from the command line. Whether you're waiting for your CI pipelines to finish or just looking for a quick break, this project offers a simple and enjoyable way to challenge your typing speed without leaving your terminal.

## 🚀 Features

- **Multiplayer Fun**: Play TypeRacer-style games with friends through a client-server setup, entirely in your terminal.
- **Cross-Platform Compatibility**: Easy to install via `curl`, with pre-compiled binaries available for MacOS and Windows systems.
- **Quick and Easy Setup**: No complicated setup. Just download the appropriate version, and you're good to go.
- **Developed in Go**: Built using Go, both the server and client provide a fast and reliable experience.
- **Integration Tests with Bun**: We use Bun for integration tests to ensure seamless performance across systems.

## 🛠️ Installation and Usage

<-- TODO -->

Main goal is to publish compiled version as a part of pipeline somewhere and provide users with prepared curls to download it for their OS version.

## Developer Guide

**How to run server?**

```
# From main catalog
cd server
go run main.go
```

Make sure port 8080 is free (@TODO: this have to be moved into env variables)

**How to run client?**

```
# From main catalog
cd client
go run main.go
```

**How to run integration tests?**
Before make sure:

1. Server is started somewhere and you can connect to it
2. You have latest version of bun installed

```
# From main catalog
cd test
bun run test
```

## **🛑 Project Status**

- [x] Initial PoC's
- [x] Communication initialisation
- [x] Hosting Game
- [x] Joining Game
- [x] Starting Game (Ready Check)
- [x] Running Game (Live time preview of your opponent progress and stats)
- [ ] Random text to race on
- [x] Game Summary
- [x] Replay

### Additional Features

- [ ] Nick Change
- [x] Server URL Change (we will host single server somewhere but we won't hold you back from hosting your own). To change server URL use environment variable `TR_WS_URL`

Each of these steps have to be integration tests covered and tested before checking out.
