// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Remote Hostname = %q\n", r.URL.Host)
	fmt.Fprintf(w, "Remote Addr = %q\n", r.RemoteAddr)
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

	// get client IP address
	ip, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Print("userIP: [", r.RemoteAddr, "] is not IP:port")
	}
	userIP := net.ParseIP(ip)
	if userIP == nil {
		log.Print("userIP: [", r.RemoteAddr, "] is not IP:port")
		return
	}
	fmt.Fprintf(w, "IP: %s\n", ip)
	fmt.Fprintf(w, "Port: %s\n", port)
	fmt.Fprintf(w, "userIP: %v\n", userIP)
}
