package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

var (
	//go:embed config.json
	config []byte
	//go:embed all:build/*
	frontend embed.FS
)

func main() {
	var cfg Config
	if err := json.Unmarshal(config, &cfg); err != nil {
		log.Fatalln("Error reading config file:", err)
	}
	frontend, err := fs.Sub(frontend, "build")
	if err != nil {
		log.Fatalln("Error retrieving frontend directory:", err)
	}
	s := NewServer(cfg)
	go s.Loop()
	http.HandleFunc("/ws", s.handleWebsocket)
	http.Handle("/", http.FileServer(http.FS(frontend)))
	var listener net.Listener
	if addr, ok := os.LookupEnv("LISTEN_ADDR"); ok {
		listener, err = net.Listen("tcp", addr)
	} else {
		port := 8080
		for port < 65536 {
			listener, err = net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(port))
			if err == nil {
				break
			}
			port++
		}
		if err == nil {
			log.Println("Presentation: http://127.0.0.1:" + strconv.Itoa(port))
			log.Println("Admin:        http://127.0.0.1:" + strconv.Itoa(port) + "/admin.html")
		}
	}
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	srv := http.Server{}
	if err := srv.Serve(listener); err != nil {
		log.Println("Error listening and serving:", err)
	}
}
