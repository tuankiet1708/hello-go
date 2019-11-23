package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/pat"
)

type Location struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Address   string  `json:"adr"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Type      string  `json:"type"`
}

//Lưu các địa điểm, mất khi server đóng
var locationStore = make(map[int]Location)

//Biến tạo khóa cho map dữ liệu, id địa điểm
var id int = 0

//HTTP Post - /api/locations
func PostLocationHandler(w http.ResponseWriter, r *http.Request) {
	var location Location
	// Giải mã thông tin địa điểm từ json
	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		panic(err)
	}

	id++
	location.Id = id
	locationStore[id] = location
	j, err := json.Marshal(location)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

//HTTP Get - /api/locations
func GetLocationHandler(w http.ResponseWriter, r *http.Request) {
	var locations []Location
	for _, v := range locationStore {
		locations = append(locations, v)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(locations)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//HTTP Get - /api/locations/{id}
func GetOneLocationHandler(w http.ResponseWriter, r *http.Request) {
	k, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil {
		panic(err)
	}

	if get, ok := locationStore[k]; ok {
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(get)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(nil)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write(j)
	}
}

//HTTP Put - /api/locations/{id}
func PutLocationHandler(w http.ResponseWriter, r *http.Request) {
	k, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil {
		panic(err)
	}
	var locationToUpd Location
	// Giải mã thông tin địa điểm từ json
	err = json.NewDecoder(r.Body).Decode(&locationToUpd)
	if err != nil {
		panic(err)
	}
	if _, ok := locationStore[k]; ok {
		// Xóa địa điểm hiện có và thêm địa điểm cập nhật
		delete(locationStore, k)
		locationToUpd.Id = k
		locationStore[k] = locationToUpd
	} else {
		log.Printf("Không tìm thấy khóa %s để xóa", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

//HTTP Delete - /api/locations/{id}
func DeleteLocationHandler(w http.ResponseWriter, r *http.Request) {
	k, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil {
		panic(err)
	}
	// Xóa khỏi map lưu trữ
	if _, ok := locationStore[k]; ok {
		// Xóa địa điểm ứng khóa k
		delete(locationStore, k)
	} else {
		log.Printf("Không tìm thấy khóa %s để xóa", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := pat.New()
	r.Get("/api/locations/{id}", GetOneLocationHandler)
	r.Get("/api/locations", GetLocationHandler)
	r.Post("/api/locations", PostLocationHandler)
	r.Put("/api/locations/{id}", PutLocationHandler)
	r.Delete("/api/locations/{id}", DeleteLocationHandler)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
