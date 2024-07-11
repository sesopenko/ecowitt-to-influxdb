package main

import (
	"ecowitt-to-influxdb/internal/ecowitt-to-influxdb/ecowitt"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		log.Printf("Inavlid request method")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		log.Printf("Invalid request method")
		return
	}
	fmt.Fprintf(w, "Received POST request with body: %s\n", string(body))
	log.Printf("Received POST request with body: %s\n", string(body))
	params, err := url.ParseQuery(string(body))
	if err != nil {
		http.Error(w, "Error parsing body", http.StatusBadRequest)
		log.Printf("Invalid request, unable to parse body")
		return
	}
	readingMap := ecowitt.BuildReadingMap(params)

	printReadingMap(readingMap)
	reading, err := ecowitt.BuildReading(readingMap)
	if err != nil {
		http.Error(w, "Error parsing body", http.StatusBadRequest)
		log.Printf("Invalid request, unable to parse body into reading struct")
	}
	log.Printf("parsed results: %+v", reading)
}

func printReadingMap(readingMap map[string]string) {
	for key, value := range readingMap {
		fmt.Printf("%s: %s\n", key, value)
	}
}

func main() {

	http.HandleFunc(ecowitt.EcowittMeasurementHttpPath, postHandler)
	port := 20555
	listenAddr := fmt.Sprintf(":%d", port)
	log.Printf("Starting server on address %s", listenAddr)
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}

}
