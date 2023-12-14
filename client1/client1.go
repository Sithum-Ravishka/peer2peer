package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func sendFileToPort3000(fileName string) {
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Send the file name
	_, err = conn.Write([]byte(fileName))
	if err != nil {
		log.Println("Error sending file name:", err)
		return
	}

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Send the file data
	_, err = io.Copy(conn, file)
	if err != nil {
		log.Println("Error sending file data:", err)
		return
	}

	fmt.Println("File", fileName, "sent successfully via port 3000")
}

func main() {
	fileName := "example.png"

	// Send the file via port 3000
	sendFileToPort3000(fileName)
}
