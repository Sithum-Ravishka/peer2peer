package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Start a server on one node
	go startServer()

	// Connect to the server from another node
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Open the file to be sent
	file, err := os.Open("example.png")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file and send it over the connection
	buffer := make([]byte, 1024)
	for {
		bytesRead, err := file.Read(buffer)
		if err != nil {
			break
		}
		conn.Write(buffer[:bytesRead])
	}

	fmt.Println("File sent successfully")
}

func startServer() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}
	defer conn.Close()

	// Receive the file sent by the client
	receivedFile, err := os.Create("received_file.png")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer receivedFile.Close()

	buffer := make([]byte, 1024)
	for {
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			break
		}
		receivedFile.Write(buffer[:bytesRead])
	}

	fmt.Println("File received successfully")
}
