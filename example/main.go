package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mosn/holmes"
)

func init() {
	http.HandleFunc("/make1gb", make1gbslice)
	go http.ListenAndServe(":10003", nil)
}

var h = holmes.New("5s", "1m", "/tmp", false).
	EnableCPUDump().Config(10, 25, 80).
	EnableGoroutineDump().Config(10, 25, 20000).
	EnableMemDump().Config(3, 25, 80)

func main() {
	h.Start()
	for range time.Tick(time.Second) {
		fmt.Println("time passed like fly")
		//fmt.Printf("%#v\n", h)
	}
}

func make1gbslice(wr http.ResponseWriter, req *http.Request) {
	var a = make([]byte, 1073741824)
	_ = a
}

func deadlock(wr http.ResponseWriter, req *http.Request) {
	// TODO
}

func channelBlock(wr http.ResponseWriter, req *http.Request) {
	// TODO
}

func goroutineleaks(wr http.ResponseWriter, req *http.Request) {
	// TODO
}
