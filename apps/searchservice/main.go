package main

import (
	"flag"

	"github.com/moetang/searchservice/server"
)

var (
	listenString   string
	dictionaryPath string
)

func init() {
	flag.StringVar(&listenString, "listen", "tcp://0.0.0.0:20100", "e.g. tcp://0.0.0.0:20100")
	flag.StringVar(&dictionaryPath, "dict", "dictionary.txt", "e.g. /home/user/dictionary.txt")

	flag.Parse()
}

func main() {
	config := &server.ServerConfig{
		DictionaryPath: dictionaryPath,
		ListenString:   listenString,
	}
	server.StartAndListen(config)
}
