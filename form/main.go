package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

var locationStore = make(map[int]Location)
var id int = 0

type Location struct {
	Name      string  `json:"name"`
	Address   string  `json:"adr"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Type      string  `json:"type"`
}

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var locpage string = `
            <html>
                <head>
                    <title>Địa điểm mới</title>
                </head>
                <body>
                    <h2>THÊM ĐỊA ĐIỂM MỚI</h2>
                    <form action="/locations" method="post">
                        Tên:<input type="text" name="name"><br/>
                        Địa chỉ:<input type="text" name="adr"><br/>
                        Vĩ độ:<input type="text" name="lat"><br/>
                        Kinh độ:<input type="text" name="lon"><br/>
                        Loại:<select name="type">
                                <option value="-">-</option>
                                <option value="Ăn uống">Ăn uống</option>
                                <option value="Nghỉ ngơi">Nghỉ ngơi</option>
                                <option value="Đi lại">Đi lại</option>
                                <option value="Giải trí">Giải trí</option>
                                <option value="Tiện ích">Tiện ích</option></select><br/>
                     <input type="submit" value="Thêm">
					</form>
					<form enctype="multipart/form-data" action="/upload" method="post">
						<input type="file" name="upfile" accept=".jpg,.png" />
						<input type="submit" value="Tải lên" />
					</form>
                </body>
            </html>`
		fmt.Fprintf(w, locpage)
	} else {
		var location Location
		// Phân tích dữ liệu gửi từ form
		r.ParseForm()
		location.Name = r.PostForm.Get("name")
		location.Name = template.HTMLEscapeString(r.PostForm.Get("name"))

		if len(strings.TrimSpace(location.Name)) == 0 {
			fmt.Fprintf(w, "Thêm không thành công do tên địa điểm rỗng!\nThông tin các địa điểm: %v", locationStore)
			return
		}
		location.Address = r.PostForm.Get("adr")

		if check, _ := regexp.MatchString("^([-+]?)([0-9]*)(.?)([0-9]+)$", r.PostForm.Get("lat")); !check {
			fmt.Fprintf(w, "Thêm không thành công do sai latitude!\nThông tin các địa điểm: %v", locationStore)
			return
		}

		if check, _ := regexp.MatchString("^([-+]?)([0-9]*)(.?)([0-9]+)$", r.PostForm.Get("lon")); !check {
			fmt.Fprintf(w, "Thêm không thành công do sai longitudes!\nThông tin các địa điểm: %v", locationStore)
			return
		}

		// if check, _ := regexp.MatchString(`^([\w\.\_]+)@(\w+).([a-z]{2,})$`, r.PostForm.Get("email")); !check {
		// 	fmt.Fprintf(w, "Thêm không thành công do sai email!\nThông tin các địa điểm: %v", locationStore)
		// 	return
		// }

		loctype := []string{"Ăn uống", "Nghỉ ngơi", "Đi lại", "Giải trí", "Tiện ích"}
		for _, v := range loctype {
			if v == r.PostForm.Get("type") {
				location.Type = v
				break
			}
		}
		if len(location.Type) == 0 {
			fmt.Fprintf(w, "Thêm không thành công do sai loại địa điểm!\nThông tin các địa điểm: %v", locationStore)
			return
		}

		var err error

		location.Latitude, err = strconv.ParseFloat(r.PostForm.Get("lat"), 32)
		if err != nil {
			panic(err)
		}
		location.Longitude, err = strconv.ParseFloat(r.PostForm.Get("lon"), 32)
		if err != nil {
			panic(err)
		}
		location.Type = r.PostForm.Get("type")
		id++
		locationStore[id] = location
		fmt.Fprintf(w, "Thêm thành công!\nThông tin các địa điểm: %v", locationStore)
	}
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("upfile")
	if err != nil {
		fmt.Fprintf(w, "Tải file lỗi!")
		return
	}
	defer file.Close()

	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		fmt.Fprintf(w, "Đọc file lỗi!")
		return
	}

	filetype := http.DetectContentType(buff)
	if filetype == "image/jpeg" || filetype == "image/jpg" || filetype == "image/png" {
		f, err := os.OpenFile("./public/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Fprintf(w, "Lưu file lỗi!")
			return
		}
		defer f.Close()
		io.Copy(f, file)
		fmt.Fprintf(w, "Tải file thành công: %v", handler.Header)
	} else {
		fmt.Fprintf(w, "Định dạng file không hỗ trợ: %s", filetype)
	}
}

func main() {
	http.HandleFunc("/", LocationHandler)
	http.HandleFunc("/locations", LocationHandler)
	http.HandleFunc("/upload", FileHandler)
	http.ListenAndServe(":8080", nil)
}
