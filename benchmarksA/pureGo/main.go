package main

import (
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var reqCount = 0
var mutex sync.Mutex

func main() {
	println("Listening on port http://localhost:8000")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8000", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	reqCount++
	myReqCount := reqCount
	mutex.Unlock()

	PauseMs(1000)
	io.WriteString(w, "Hello from pure go ! "+strconv.Itoa(myReqCount))
}

func PauseMs(timeInMs int) {
	duration := time.Millisecond * time.Duration(timeInMs)
	time.Sleep(duration)
}
