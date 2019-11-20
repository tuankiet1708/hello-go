package main

import (
	"encoding/json"
	"fmt"
)

type Location struct {
	Name      string  `json:"name"`
	Address   string  `json:"adr"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Type      string  `json:"type"`
}

type Locations struct {
	Locs []Location `json:"loc"`
}

func main() {
	var locs Locations
	jsonStr := `{"loc":[{"name":"ATM Techcombank", "adr":"Chung cư K26, Dương Quảng Hàm, P.7, Gò Vấp","lat":10.832052230834961,"lon":106.68563842773438,"type":"Tiện ích"},{"name":"CGV PVT","adr":"Tầng 5 Vincom Plaza 12 Phan Văn Trị, P7, Gò Vấp","lat":10.827040672302246,"lon":106.68864440917969,"type":"Giải trí"}]}`
	json.Unmarshal([]byte(jsonStr), &locs)
	fmt.Println(locs)
}
