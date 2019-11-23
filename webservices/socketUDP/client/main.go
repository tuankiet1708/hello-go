package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// Đầu tiên tạo địa chỉ UDP từ chuỗi URL của server: localhost:8888
	udpAddr, err := net.ResolveUDPAddr("udp4", "localhost:8888")
	if err != nil {
		log.Fatal("Lỗi phân giải địa chỉ: " + err.Error())
	}
	// Tiếp theo tạo kết nối đến server:
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatal("Lỗi tạo socket: " + err.Error())
	}
	// Lúc này đã có thể gửi dữ liệu lên server:
	_, err = conn.Write([]byte("Mấy giờ rồi?"))
	if err != nil {
		log.Fatal("Lỗi gửi dữ liệu: " + err.Error())
	}
	// Nhận phản hồi từ server:
	message := make([]byte, 512)
	size, err := conn.Read(message)
	if err != nil {
		log.Fatal("Lỗi nhận dữ liệu: " + err.Error())
	}
	fmt.Printf("Nhận %d byte từ server: %s", size, string(message[:size]))

}
