package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	println("Listening on port http://localhost:8000")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8000", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	//PauseMs(1000)
	io.WriteString(w, "Hello from pure go !")
}

func PauseMs(timeInMs int) {
	duration := time.Millisecond * time.Duration(timeInMs)
	time.Sleep(duration)
}
