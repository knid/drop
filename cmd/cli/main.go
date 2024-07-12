package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gliderlabs/ssh"
	"github.com/knid/drop/handler"
	"github.com/knid/drop/models"
)

const (
	VERSION = "0.0.1"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "usage: %s <PORT> <KEYFILE>\n", os.Args[0])
		os.Exit(1)
	}

	listenPort, keyFile := os.Args[1], os.Args[2]

	dropHandler := handler.DropHandler{}
	dropHandler.Drops = make(map[string]*models.Drop)
	ssh.Handle(dropHandler.Default)

	log.Printf("Drop v%s\n", VERSION)
	log.Printf("Listening Port: %s\n", listenPort)
	log.Printf("Using Key File: %s\n", keyFile)

	serverKeyFile := ssh.HostKeyFile(keyFile)
	log.Fatal(ssh.ListenAndServe("0.0.0.0:"+listenPort, nil, serverKeyFile))
}
