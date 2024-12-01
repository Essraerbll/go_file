package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// Sunucu TCP bağlantısını dinlemek için başlatıyoruz
	listen, err := net.Listen("tcp", "0.0.0.0:9000") // Sunucu her IP'den gelen bağlantıları kabul eder
	if err != nil {
		fmt.Println("Error setting up server:", err)
		return
	}
	defer listen.Close()

	fmt.Println("Server is listening on port 9000...")

	// Bağlantı bekliyoruz
	conn, err := listen.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connection established with client.")

	// output.txt dosyasını oluşturuyoruz
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Gelen veriyi alıyoruz ve dosyaya yazıyoruz
	_, err = io.Copy(outputFile, conn)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}

	fmt.Println("File transfer complete. Data saved to output.txt")
}
