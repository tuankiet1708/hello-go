package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// Đầu tiên tạo địa chỉ TCP từ chuỗi URL của server: localhost:8888
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:8888")
	if err != nil {
		log.Fatal("Lỗi phân giải địa chỉ: " + err.Error())
	}

	// Tiếp theo tạo kết nối đến server:
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal("Lỗi tạo socket: " + err.Error())
	}

	// Lúc này đã có thể gửi dữ liệu lên server:
	_, err = conn.Write([]byte("Mấy giờ?"))
	if err != nil {
		log.Fatal("Lỗi gửi dữ liệu: " + err.Error())
	}

	// Nhận phản hồi từ server:
	message := make([]byte, 512)
	n, err := conn.Read(message)
	if err != nil {
		log.Fatal("Lỗi nhận dữ liệu: " + err.Error())
	}
	fmt.Printf("Nhận %d byte từ server: %s", n, string(message[:n]))
}
