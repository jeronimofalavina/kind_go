package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func app(w http.ResponseWriter, r *http.Request) {

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	now := time.Now()
	fmt.Fprintf(w, "date: %s\nhostname: %s\n", now, hostname)
}

func handleRequests() {
	http.HandleFunc("/app", app)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8181", nil))
}

func main() {
	handleRequests()
}
