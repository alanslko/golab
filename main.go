package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"encoding/json"

	"github.com/alanslko/golab/internal/ginlab"
)

// nullable(YES), editable(YES), unique(NO), autofill(NO), refreshable(NO), displayed(YES)

type parent struct {
	Name string `json:"parentName"`
	Kid  child  `isNullable:"NO" isKey:"YES"`
}

type child struct {
	Name string
	Age  int
}

func init() {
	//	fmt.Println(database.GetName())
	//	fmt.Println(formatter.GetName())
}

func init() {
	initRouter()
}

func initRouter() {
	http.HandleFunc("/golab", handleGoLabRequest)
}

func handleGoLabRequest(w http.ResponseWriter, r *http.Request) {
	write(w, p)
}

func write(w http.ResponseWriter, data interface{}) {
	w.Header().Set("content-type", "application/json")
	b, err := json.Marshal(data)
	if err != nil {
		log.Printf("[info] write http response (%v) meet error (%v)", data, err)
	}
	w.Write(b)
}

var p parent

func main() {
	p = parent{"cindy", child{"stefan", 7}}
	fmt.Println(p.Name)
	var t = reflect.TypeOf(p)

	var tag = t.Field(1).Tag
	fmt.Printf("reflect: %s %s", tag.Get("isNullable"), tag.Get("isKey"))
	/*
		if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
			log.Printf("%v", err)
		}
	*/

	ginlab.GinLab()
}
