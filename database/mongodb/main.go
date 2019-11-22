package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Location struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	Name      string
	Address   string
	Latitude  float64
	Longitude float64
	Type      string
}

func main() {
	// Kết nối cơ sở dữ liệu:
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		log.Fatal("Lỗi kết nối:", err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("location")

	// Thêm địa điểm:
	err = c.Insert(
		&Location{bson.NewObjectId(), "CGV PVT", "Tầng 5 Vincom 12 Phan Văn Trị, P.7, Gò Vấp", 10.8270, 106.6886, "Giải trí"},
		&Location{bson.NewObjectId(), "ATM Techcombank", "CC K26 Dương Quảng Hàm, P.7, Gò Vấp", 10.830604, 106.6840686, "Tiện ích"})

	if err != nil {
		log.Fatal("Lỗi thêm dữ liệu:", err)
	}

	// Lấy toàn bộ địa điểm, sắp xếp tăng dần theo tên:
	result := Location{}
	iter := c.Find(nil).Sort("name").Iter()
	var ids []bson.ObjectId
	for iter.Next(&result) {
		fmt.Println("Địa điểm:", result)
		ids = append(ids, result.Id)
	}

	fmt.Println("IDs", ids)

	if err = iter.Close(); err != nil {
		log.Fatal(err)
	}
}
