# ws-reload #

Auto-reload webpage by monitoring files/directories, for web development.

It's a simple *Go* program to

- watch one or more files/directories
- tail (follow) log file
- start a websocket server and serve events

On client side, a *JavaScript* to

- connect to websocket server and listen to events
- reload page once file/directory changed
- display content from log file

## Background ##

I wanted a simple way to do *reload-on-save* when I was working on a *PHP* project.


## Build and Run ##

Update `config.json.example` and save as `config.json`.

(note, *directories* watch are **NOT recursive**)

**Build**

```
$ go get .
$ go build monitor.go
$ go build ws-server
```

**Run**

```
$ ./ws-server && ./monitor -f config.json
```

## Todo ##

- move into package.
- combine server and monitor as one.

