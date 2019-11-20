package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Vui lòng thêm URL server cần check!")
		return
	}

	if err := checkServer(os.Args[1]); err == nil {
		fmt.Println("Server đang hoạt động!")
	} else {
		fmt.Println(err)
	}
}

func checkServer(url string) error {
	const timeout = 1 * time.Minute
	fmt.Println("timeout", timeout)
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // thành công
		}
		log.Printf("Server không trả lời do lỗi (%s); thử lại ...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("Server %s không phản hồi sau %s phút!", url, timeout)
}
