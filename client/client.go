package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func sendPNGFromPort3000To4000(pngFileName string) {
	conn, err := net.Dial("tcp", "localhost:4000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Send the file name
	_, err = conn.Write([]byte(pngFileName))
	if err != nil {
		log.Println("Error sending file name:", err)
		return
	}

	// Open the PNG file
	file, err := os.Open(pngFileName)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Send the PNG file data
	_, err = io.Copy(conn, file)
	if err != nil {
		log.Println("Error sending file data:", err)
		return
	}

	fmt.Println("PNG file", pngFileName, "sent successfully from port 3000 to port 4000")
}

func sendZIPFromPort4000To3000(zipFileName string) {
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Send the file name
	_, err = conn.Write([]byte(zipFileName))
	if err != nil {
		log.Println("Error sending file name:", err)
		return
	}

	// Open the ZIP file
	file, err := os.Open(zipFileName)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Send the ZIP file data
	_, err = io.Copy(conn, file)
	if err != nil {
		log.Println("Error sending file data:", err)
		return
	}

	fmt.Println("ZIP file", zipFileName, "sent successfully from port 4000 to port 3000")
}

func main() {
	pngFileName := "example.png"
	zipFileName := "example.zip"

	// Send the PNG file from port 3000 to port 4000
	sendPNGFromPort3000To4000(pngFileName)

	// Send the ZIP file from port 4000 to port 3000
	sendZIPFromPort4000To3000(zipFileName)
}
