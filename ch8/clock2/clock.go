// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 222.

// Clock is a TCP server that periodically writes the time.
package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func handleConn(c net.Conn, loc *time.Location) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().In(loc).Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// exercise 8.1
	// TODO: take in TZ from environment variable and port as command line arg
	// ex:
	// $ TZ=US/Eastern ./clock2 -port 8010 &

	tz := os.Getenv("TZ")
	location, err := time.LoadLocation(tz)
	if err != nil {
		log.Fatal(err)
	}

	port := "8000"
	if len(os.Args) > 2 && os.Args[1] == "-port" {
		port = os.Args[2]
	}

	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, location) // handle connections concurrently
	}
	//!-
}
