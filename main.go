package main

import (
	"fmt"
	"net/http"

	"github.com/Ezian/gof-art/channel"
	"github.com/Ezian/gof-art/mutex"
	"github.com/Ezian/gof-art/naive"
	"github.com/Ezian/gof-art/utils"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, "Hello, you've requested: %s and that's wrong\n", req.URL.Path)
	})
	http.HandleFunc("/Gophers.jpg", func(rw http.ResponseWriter, req *http.Request) {
		http.ServeFile(rw, req, "./images/Gophers.jpg")
	})
	http.HandleFunc("/naive", utils.CreateHandler(naive.HandlePost, naive.HandleGet))
	http.HandleFunc("/mutex", utils.CreateHandler(mutex.HandlePost, mutex.HandleGet))
	http.HandleFunc("/channel", utils.CreateHandler(channel.HandlePost, channel.HandleGet))

	http.ListenAndServe(":9999", nil)
}
