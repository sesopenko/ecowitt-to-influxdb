package main

import (
	"ecowitt-to-influxdb/internal/ecowitt"
	"ecowitt-to-influxdb/internal/influx"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Handler struct {
	InfluxClient influxdb2.Client
	InfluxConfig influx.InfluxConfig
}

func (h Handler) postHandler(w http.ResponseWriter, r *http.Request) {
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
	params, err := url.ParseQuery(string(body))
	if err != nil {
		http.Error(w, "Error parsing body", http.StatusBadRequest)
		log.Printf("Invalid request, unable to parse body")
		return
	}
	readingMap := ecowitt.BuildReadingMap(params)

	reading, err := ecowitt.BuildReading(readingMap)
	if err != nil {
		http.Error(w, "Error parsing body", http.StatusBadRequest)
		log.Printf("Invalid request, unable to parse body into reading struct")
	}
	writeApi := h.InfluxClient.WriteAPI(h.InfluxConfig.Org, h.InfluxConfig.Bucket)
	p := influxdb2.NewPoint("local_weather",
		map[string]string{
			"country_prov_city": h.InfluxConfig.CountryProvCity,
			"source":            "ecowitt",
		},
		map[string]interface{}{
			"internal_temperature_fahrenheight": reading.TemperatureIndoorsFahrenheit,
			"barometric_pressure_absolute_inhg": reading.BarometricPressureAbsolute,
			"barometric_pressure_relative_inhg": reading.BarometicPressureRelative,
			"humidity_indoors":                  reading.HumidityIndoors,
		},
		reading.ReadingTime)
	writeApi.WritePoint(p)
	writeApi.Flush()

	defer writeApi.Flush()
}

func postHandler(w http.ResponseWriter, r *http.Request) {

}

func printReadingMap(readingMap map[string]string) {
	for key, value := range readingMap {
		fmt.Printf("%s: %s\n", key, value)
	}
}

func main() {
	influxConfig := influx.GetInfluxConfig()
	influxClient := influxdb2.NewClientWithOptions(
		influxConfig.Url,
		influxConfig.AuthToken,
		// batch size of 20 is simply from example on api library README.md
		influxdb2.DefaultOptions().SetUseGZip(true).SetBatchSize(20),
	)

	handler := Handler{
		InfluxClient: influxClient,
		InfluxConfig: influxConfig,
	}

	// get non-blocking write client
	http.HandleFunc(ecowitt.EcowittMeasurementHttpPath, handler.postHandler)
	port := 20555
	listenAddr := fmt.Sprintf(":%d", port)
	log.Printf("Starting server on address %s", listenAddr)
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
	defer influxClient.Close()
}
