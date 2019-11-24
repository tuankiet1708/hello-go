package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

var btServer *http.Server

func startServer(port string) {
	btServer = &http.Server{Addr: "127.0.0.1:" + port}

	http.Handle("/", http.FileServer(http.Dir("public")))

	fmt.Printf("Lắng nghe port 127.0.0.1:%s...", port)
	go func() {
		if err := btServer.ListenAndServe(); err != nil {
			fmt.Println("Lỗi Httpserver: ListenAndServe()", err)
		}
	}()
}

func stopServer() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := btServer.Shutdown(ctx)
	if err != nil {
		fmt.Println("Lỗi tắt server", err)
	}
	return err
}

func main() {
	startServer("8000")
	// stopServer()
}

//
// Khi phương thức Shutdown được gọi, chúng sẽ lần lượt dừng việc lắng nghe kết nối mới, đóng các kết nối đang mở hoặc tạm nghỉ. Nếu việc tắt dịch vụ lâu hơn thời gian khai báo ở biến ngữ cảnh ctx, hệ thống sẽ báo lỗi. Như vậy với phương thức Shutdown, chúng ta đã có thể tự mình tắt dịch vụ một cách chủ động. Nếu muốn sử dụng dịch vụ bên thứ ba, chúng ta có thể thử tylerb/graceful hay facebookgo/httpdown
