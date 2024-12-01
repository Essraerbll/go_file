package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Kullanıcıdan dosya adını alıyoruz
	fmt.Print("Enter the filename to send to the server: ")
	var filename string
	fmt.Scanln(&filename)

	// Dosyayı açıyoruz
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Dosyanın tamamını okuyoruz
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Server'a bağlanıyoruz (Server'ın IP adresi ve portu)
	conn, err := net.Dial("tcp", "192.168.1.112:9000") // Server IP adresi
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to the server. Sending file...")

	// Dosyanın tamamını server'a gönderiyoruz
	_, err = conn.Write(buffer)
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}

	fmt.Println("File transfer complete. Closing connection.")
}
