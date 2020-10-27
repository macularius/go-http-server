package main

import (
	"flag"
	"log" // пакет для логирования
	"net"
	"net/http" // пакет для поддержки HTTP протокола
)

func main() {
	var host, port, dir *string

	host = flag.String("host", "", `HTTP server's host. Default: "localhost"`)
	port = flag.String("port", "8080", `HTTP server's port. Default: "8080"`)
	dir = flag.String("dir", "./public", `HTTP server's dir. Default: "./public"`)

	flag.Parse()

	log.Printf("Server listen %s:%s. Dir: %s\n", *host, *port, *dir)

	// server for './public' dir
	log.Fatal(http.ListenAndServe(net.JoinHostPort(*host, *port), http.FileServer(http.Dir(*dir))))
}
