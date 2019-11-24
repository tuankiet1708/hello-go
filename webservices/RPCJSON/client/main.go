package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Point struct {
	Lat, Lon float64
}

type Points struct {
	A Point
	B Point
}

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Lỗi kết nối:", err.Error())
	}
	// Synchronous call
	args := Points{Point{10.832052230834961, 106.68563842773438}, Point{10.827040672302246, 106.68864440917969}}
	var res int
	err = client.Call("Distance.Calculate", args, &res)
	if err != nil {
		log.Fatal("Lỗi gọi rpc Distance.CalculateDistance:", err)
	}
	fmt.Printf("Khoảng cách từ %v đến %v là %dm\n", args.A, args.B, res)
}
