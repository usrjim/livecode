package main

import (
	"encoding/json"
	"flag"
	"github.com/fsnotify/fsnotify"
	"github.com/hpcloud/tail"
	"golang.org/x/net/websocket"
	"io/ioutil"
	"log"
	"time"
)

type Config struct {
	Origin string   `json:"origin"`
	Server string   `json:"server"`
	Target []string `json:"target"`
	Log    string   `json:"log"`
}

type Message struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

var config_json string
var origin string
var url string

func init() {
	flag.StringVar(&config_json, "f", "config.json", "config file in json format.")
}

func main() {
	flag.Parse()
	content, _ := ioutil.ReadFile(config_json)
	var config Config
	err := json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal(err)
	}

	origin = config.Origin
	url = config.Server

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	go keepalive(*ws)
	go tailFile(config.Log, *ws)
	watchFiles(config.Target, *ws)
}

func keepalive(ws websocket.Conn) {
	for {
		message := []byte("{\"type\":\"broadcast\",\"payload\":\"ping\"}")
		_, err := ws.Write(message)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(30 * time.Second)
	}
}

func tailFile(f string, ws websocket.Conn) {
	t, _ := tail.TailFile(f, tail.Config{Follow: true})
	for line := range t.Lines {
		//log.Println(line.Text)
		m := &Message{
			Type:    "broadcast",
			Payload: line.Text,
		}
		jm, _ := json.Marshal(m)
		ws.Write(jm)
	}
}

func watchFiles(watch_paths []string, ws websocket.Conn) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				// log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)

					m := &Message{
						Type:    "broadcast",
						Payload: "reload",
					}
					jm, _ := json.Marshal(m)
					ws.Write(jm)

					//log.Println(string(jm))
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	for _, f := range watch_paths {
		err = watcher.Add(f)
		if err != nil {
			log.Fatal(err)
		}
	}
	<-done
}
