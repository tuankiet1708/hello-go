package main

import (
	"log"
	"math"
	"net"
	"net/rpc"
)

type Point struct {
	Lat, Lon float64
}

type Points struct {
	A Point
	B Point
}

type Distance int

const EARTH_RADIUS = 6378137 // Bán kính trái đất tính theo mét

func (t *Distance) Calculate(args *Points, dis *Distance) error {
	p1 := args.A
	p2 := args.B
	dLat := (p2.Lat - p1.Lat) * (math.Pi / 180.0)
	dLon := (p2.Lon - p1.Lon) * (math.Pi / 180.0)

	lat1 := p1.Lat * (math.Pi / 180.0)
	lat2 := p2.Lat * (math.Pi / 180.0)

	a := (math.Sin(dLat/2) * math.Sin(dLat/2)) + (math.Sin(dLon/2) * math.Sin(dLon/2) * math.Cos(lat1) * math.Cos(lat2))

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	*dis = Distance(EARTH_RADIUS * c)
	return nil
}

func main() {
	dis := new(Distance)
	rpc.Register(dis)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		log.Fatal("Lỗi phân giải địa chỉ" + err.Error())
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Println("Lỗi lắng nghe kết nối" + err.Error())
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}
