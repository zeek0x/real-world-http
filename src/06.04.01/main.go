package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handleChunkedResponse(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}
	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "Chunk #%d\n", i)
		flusher.Flush()
		time.Sleep(500 * time.Millisecond)
	}
	flusher.Flush()
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/chunked", handleChunkedResponse)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
