package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handleClient(conn *net.TCPConn) {
	message := make([]byte, 512)
	n, err := conn.Read(message)
	if err != nil {
		log.Fatal("Lỗi nhận dữ liệu: " + err.Error())
	}
	fmt.Printf("Nhận %d byte từ client: %s\n", n, string(message[:n]))

	daytime := time.Now().String()
	conn.Write([]byte(daytime))
}

func main() {
	// Đầu tiên cũng tạo địa chỉ như ở client:
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:8888")
	if err != nil {
		log.Fatal("Lỗi phân giải địa chỉ: " + err.Error())
	}
	// Tiếp theo là phần đăng ký và lắng nghe kết nối từ client:
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal("Lỗi đăng ký kết nối: " + err.Error())
	}

	// for {
	// 	conn, err := listener.AcceptTCP()
	// 	if err != nil {
	// 		log.Println("Lỗi mở kết nối: " + err.Error())
	// 		continue
	// 	}
	// 	message := make([]byte, 512)
	// 	n, err := conn.Read(message)
	// 	if err != nil {
	// 		log.Fatal("Lỗi nhận dữ liệu: " + err.Error())
	// 	}
	// 	fmt.Printf("Nhận %d byte từ client: %s\n", n, string(message[:n]))

	// 	daytime := time.Now().String()
	// 	conn.Write([]byte(daytime))
	// 	conn.Close()
	// }

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Println("Lỗi mở kết nối: " + err.Error())
			continue
		}
		go handleClient(conn)
	}
}
