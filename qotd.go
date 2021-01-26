package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
)

var lines [][]byte

func main() {
	if len(os.Args) < 2 {
		println("No quote file path")
		return
	}

	dat, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	lines = bytes.Split(dat, []byte("\n"))

	for i, v := range lines {
		fmt.Println(i, ":->", string(v))
	}

	tcp, err := net.Listen("tcp", ":17")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := tcp.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	conn.Write(lines[rand.Int()%len(lines)])
	conn.Write([]byte("\n"))
	conn.Close()
}
