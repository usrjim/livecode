# Livecode #

Auto-reload webpage by monitoring files/directories, for web development purpose.

It's a simple *Go* program to

- watch one or more files/directories
- tail (follow) log file
- start a websocket server to serve events

On client side, a *JavaScript* to

- connect to websocket server and listen to events
- reloads page once file/directory changed
- `console.log` content from log file

## Background ##

I wanted a simple way to do *reload-on-save* when I was working on a *PHP* project.


## Build and Run ##

Copy `sample_config.json` file to `config.json` and update it.

(or update `c2j.php` then run `$ php c2j.php > config.json`)

(note, *directories* watch are not **recursive**)

```
$ go get .
$ go build livecode.go
$ ./livecode -f config.json
```

Include `livecode.js` in webpage.

## Todo ##

Package it.

