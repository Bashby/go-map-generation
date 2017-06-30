package websocket

import (
	"log"
	"net/http"
)

func Serve(address string) {
	// Initiate message hub
	hub := newHub()
	go hub.run()

	// Configure the webserver
	http.HandleFunc("/", serveRoot)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebsocketRequest(hub, w, r)
	})

	// Listen
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		log.Printf("[Webserver] Serving on %s...\n", address)
	}
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	log.Printf("[Webserver] Serving %s\n", r.URL.String())
	if r.URL.Path != "/" {
		http.Error(w, "Not Found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	http.ServeFile(w, r, "websocket_test.html")
}
