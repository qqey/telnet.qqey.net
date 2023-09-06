package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var counter int = 0

func handleConnection(conn net.Conn) {
	defer conn.Close()
	counter++

	var welcomeMessage = fmt.Sprintf(`
You must use UTF-8 Encoding to view this page!!

Welcome to the telnet.qqey.net!

You are the %d th access human!
Your IP Address is %s 

Web: https://qqey.net
GitHub Organization: https://github.com/qqey

`, counter, conn.RemoteAddr())

	conn.Write([]byte(welcomeMessage))
}

func main() {
	listener, err := net.Listen("tcp", "localhost:23")
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("TCP server is listening on port 23")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("Shutting down the server...")
		listener.Close()
		os.Exit(0)
	}()

	for {
		conn, err := listener.Accept()

		fmt.Println(conn.RemoteAddr())
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
