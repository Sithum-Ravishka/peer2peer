package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func handleFileTransfer(conn net.Conn) {
	defer conn.Close()

	// Receive the file name
	fileNameBuf := make([]byte, 1024)
	n, err := conn.Read(fileNameBuf)
	if err != nil {
		log.Println("Error reading file name:", err)
		return
	}
	fileName := string(fileNameBuf[:n])

	// Create the file
	file, err := os.Create(fileName)
	if err != nil {
		log.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Receive the file data
	_, err = io.Copy(file, conn)
	if err != nil {
		log.Println("Error receiving file data:", err)
		return
	}

	fmt.Println("File", fileName, "received successfully")
}

func main() {
	// Start a TCP server on port 3000
	ln1, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	defer ln1.Close()

	fmt.Println("Server is listening on port 3000")

	// Start a TCP server on port 4000
	ln2, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatal(err)
	}
	defer ln2.Close()

	fmt.Println("Server is listening on port 4000")

	// Accept incoming connections on port 3000
	go func() {
		for {
			conn, err := ln1.Accept()
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Println("Accepted connection on port 3000")
			go handleFileTransfer(conn)
		}
	}()

	// Accept incoming connections on port 4000
	go func() {
		for {
			conn, err := ln2.Accept()
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Println("Accepted connection on port 4000")
			go handleFileTransfer(conn)
		}
	}()

	select {}
}
