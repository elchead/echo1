// $ 6g echo.go && 6l -o echo echo.6
// $ ./echo
//
//  ~ in another terminal ~
//
// $ nc localhost 3540

package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"time"
)

const PORT = 3540

func main() {
	fmt.Println("Listen on 3540")
	now := time.Now()
	go func() {
		server, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
		if server == nil {
			panic(err)
		}
		conns := clientConns(server)
		for {
			go handleConn(<-conns)
		}
	}()
	// wait for connection
	for {
		_, err := net.Dial("tcp", "localhost:3540")
		if err == nil {
			break
		}
	}
	fmt.Println("Time:", time.Since(now))
}

func clientConns(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	i := 0
	go func() {
		for {
			client, err := listener.Accept()
			if client == nil {
				fmt.Println(err)
				continue
			}
			i++
			fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()
	return ch
}

func handleConn(client net.Conn) {
	b := bufio.NewReader(client)
	for {
		line, err := b.ReadBytes('\n')
		if err != nil { // EOF, or worse
			break
		}
		client.Write(line)
	}
}
