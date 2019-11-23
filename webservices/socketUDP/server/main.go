package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handleClient(conn *net.UDPConn) {
	message := make([]byte, 512)
	n, addr, err := conn.ReadFromUDP(message)
	if err != nil {
		log.Fatal("Lỗi nhận dữ liệu: " + err.Error())
	}
	fmt.Printf("Nhận %d byte từ client: %s\n", n, string(message[:n]))

	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}

func main() {
	// Đầu tiên cũng tạo địa chỉ như ở client:
	udpAddr, err := net.ResolveUDPAddr("udp4", "localhost:8888")
	if err != nil {
		log.Fatal("Lỗi phân giải địa chỉ: " + err.Error())
	}
	// Tiếp theo là phần đăng ký và lắng nghe kết nối từ client:
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal("Lỗi đăng ký kết nối: " + err.Error())
	}
	for {
		handleClient(conn)
	}
}
