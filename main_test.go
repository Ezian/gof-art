package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

// Run server
func init() {
	go main()
}

func postJSON(url string) {

	var jsonStr = []byte(`{"url":"http://localhost:9999/Gophers.jpg","width":80}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[POST] error:", err)
		return
	}
	defer resp.Body.Close()
}

func getResult(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("[GET] error:", err)
		return
	}
	defer resp.Body.Close()
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

func testDataraceURL(url string) {

	var group sync.WaitGroup

	for i := 0; i < 2; i++ {
		runSeriesAsRoutine(func() { postJSON(url) }, 10, 20*time.Millisecond, &group)
		time.Sleep(time.Second)
	}

	for i := 0; i < 2; i++ {
		runSeriesAsRoutine(func() { getResult(url) }, 10, 20*time.Millisecond, &group)
	}

	group.Wait()
}

func TestDataraceNaive(t *testing.T) {
	testDataraceURL("http://localhost:9999/naive")
}

func TestDataraceMutex(t *testing.T) {
	testDataraceURL("http://localhost:9999/mutex")
}
func TestDataraceChannel(t *testing.T) {
	testDataraceURL("http://localhost:9999/channel")
}
