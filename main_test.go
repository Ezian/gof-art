package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

const url = "http://localhost:9999/ascii"

func postJSON() {

	var jsonStr = []byte(`{"url":"https://2.bp.blogspot.com/-50t8QbXgxwI/WGgpaXNAYWI/AAAAAAAAEPE/SKJ-Bu12qpkrP7kklk1_QmWTehLoBRFcwCLcB/s1600/Gophers.jpg","width":80}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[POST] error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("[POST] response Status:", resp.Status)
}

func getResult() {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("[GET] error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("[GET] response Status:", resp.Status)
}

func runSeriesAsRoutine(f func(), count int, wait time.Duration, group *sync.WaitGroup) {
	group.Add(1)
	go func() {
		defer group.Done()
		for i := 0; i < count; i++ {
			f()
			time.Sleep(wait)
		}
	}()
}

func TestDatarace(t *testing.T) {
	// Run server
	go main()

	var group sync.WaitGroup

	for i := 0; i < 5; i++ {
		runSeriesAsRoutine(postJSON, 3, time.Second, &group)
		time.Sleep(time.Second)
	}

	for i := 0; i < 5; i++ {
		runSeriesAsRoutine(getResult, 5, 200*time.Millisecond, &group)
	}

	group.Wait()
}
