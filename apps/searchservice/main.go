package main

import (
	"github.com/moetang/searchservice/server"
)

func main() {
	config := &server.ServerConfig{
		DictionaryPath: "dictionary.txt",
		ListenString:   "tcp://0.0.0.0:15394",
	}
	server.StartAndListen(config)
}
