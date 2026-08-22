# go-deadlock

[![Go Reference](https://pkg.go.dev/badge/github.com/paralin/go-deadlock.svg)](https://pkg.go.dev/github.com/paralin/go-deadlock)

## Introduction

**go-deadlock** is a [Deadlock](https://store.steampowered.com/app/1422450/) client plugin for [go-steam](https://github.com/paralin/go-steam).

Deadlock is a Valve Source 2 game with a Game Coordinator (GC). This library
implements the GC protocol for Deadlock, reusing the patterns from
[go-dota2](https://github.com/paralin/go-dota2).

The protobuf definitions come from the
[SteamDatabase/Protobufs](https://github.com/SteamDatabase/Protobufs)
repository, vendored as a git submodule (`deadlock/` tree). Code generation
uses the [aperturerobotics/common](https://github.com/aperturerobotics/common)
WASI protoc toolchain (protobuf-go-lite output only).

## Usage

```go
cl := steam.NewClient(...)
dl := deadlock.New(cl, logrus.StandardLogger())
// connect, then:
dl.SetPlaying(true) // sets app 1422450 as played
dl.SayHello()       // request a GC session; wait for events.ClientWelcomed
```

## Implementation Progress

 - [x] GC session state management (hello / welcome / connection status)
 - [x] Request tracking (`MakeRequest` with response matching)
 - [x] SOCache shared-object cache
 - [x] Generated API surface from `EGCCitadelClientMessages` (64 methods)
 - [x] Generated GC-to-client events
 - [ ] Deadlock-specific SO cache type IDs (await live capture)
 - [ ] Party / lobby convenience helpers

## Regenerating protobufs

```sh
make gen    # update submodule protos, regenerate protocol/, apigen, format
```
