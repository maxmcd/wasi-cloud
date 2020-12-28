# Wasi Cloud

Silly name, but the intention is a to define an interface for a wasi cloud provider. When writing [embly](github.com/embly/embly) I ended up defining a lot of custom syscalls that were really just partial implementations of socket and filesystem syscalls, but with different purposes. It seems like the better idea might be to implement "cloud platform" functionality through a filesystem. That way, instead of having to reimplement new system calls in every programming language you can just provide a virtual filesystem implementation through wasi.


## Runtime

We assume for the rest of this doc that we're talking about programs running webassembly on a server. We also assume that there are two types of execution types:
1. "Functions" run once in response to some kind of input. They can optionally return a response and/or exit.
2. Daemonized processes that run for a set (or unlimited) lifetime. These could exist for the duration of a websocket or tcp connection, or they could just be on all the time handling some task.

## Key Value Store

Here's an easy one. We mount a folder at `/beta/kv` if you write a file to that folder the name of the file is the key and the contents are the value. We mount this as a shared folder and it's shared by other users of this shared filesystem. We could implement [filesystem locks](https://github.com/WebAssembly/wasi-filesystem/issues/2) in order to build transactional KV operations on top of the key/value store.

We could also scope directories

```
├── /beta/kv/
│   ├── local/
│   ├── region/
│   ├── global/
```

Placing files in these folders would share them with a different subset of other executions. In this case, they're based on geographic proximity, but many groupings are possible.

## Request/response

When executing a "function" we want to provide the webassembly execution with request data and allow it to write a response. We could use files for this task. Maybe we have a folder

```
├── /request/
│   ├── txt
│   ├── json
```

Each of these files could be read to read the http request. The `txt` file might contain raw http, the `json` might contain a json serialized version of the request. The txt file could be read as a stream or a static file.

A response could be written to a `/response` folder. Either as a raw http stream, or json, or some other format.

The host implementation would create these files on the fly and listen for file creation events before waiting for the response.

A clear downside of this approach is its lack of compatibility with a stdlib http library. The library would need to support parsing and generation of arbitrary http bytes. One way to get around this might be by allowing the host to listen for http on a filesystem socket. To parse a singleton request

## Daemonized request/response

TODO

## IPC/RPC

TODO

## Other shared filesystem resources

TODO

 - s3/blob store
 - permissions
 - security, secret store
