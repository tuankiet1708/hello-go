package main

import (
	"html/template"
	"log"
	"os"
)

type LocType struct {
	Type string
}

type Location struct {
	Name      string
	Address   string
	Latitude  float64
	Longitude float64
	Type      []*LocType
}

const loctmpl = `
Địa điểm: {{.Name}}
Địa chỉ: {{.Address}}
Tọa độ: {{.Latitude}};{{.Longitude}}
Loại địa điểm:{{with .Type}}{{range .}} {{.Type}}{{end}}{{end}}
Bán cầu: {{with $x := .Latitude}}{{if gt $x 0.0 }} Bắc bán cầu {{ else }} Nam bán cầu {{end}}{{end}}
{{. | printf "%v"}}`

func main() {
	t, err := template.New("location").Parse(loctmpl)
	if err != nil {
		log.Fatal("Lỗi Parse: ", err)
		return
	}
	t0 := LocType{Type: "Ăn uống"}
	t1 := LocType{Type: "Tiện ích"}
	t2 := LocType{Type: "Giải trí"}
	l := Location{
		Name:      "Vincom PVT",
		Address:   "12 Phan Văn Trị, P.7, Gò Vấp, TP.HCM, Việt Nam",
		Latitude:  10.827040672302246,
		Longitude: 106.68864440917969,
		Type:      []*LocType{&t0, &t1, &t2}}
	if err := t.Execute(os.Stdout, l); err != nil {
		log.Fatal("Lỗi Execute: ", err)
		return
	}
}
