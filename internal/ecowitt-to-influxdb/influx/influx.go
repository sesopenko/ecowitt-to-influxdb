package influx

import "os"

type InfluxConfig struct {
	Url       string
	AuthToken string
	Org       string
	Bucket    string
	// <ISO 3166-1>-<ISO-3166-2>:<city name> ie: CA-BC:Prince Rupert
	CountryProvCity string
}

type InfluxMeasurement struct {
	MeasurementName string
	Tags            map[string]string
	Fields          map[string]interface{}
	Time            string // Time in RFC3339 format or Unix timestamp (in nanoseconds)
}

func GetInfluxConfig() InfluxConfig {
	c := InfluxConfig{
		Url:             os.Getenv("ETOI_INFLUX_URL"),
		AuthToken:       os.Getenv("ETOI_AUTH_TOKENL"),
		Org:             os.Getenv("ETOI_ORG"),
		Bucket:          os.Getenv("ETOI_BUCKET"),
		CountryProvCity: os.Getenv("ETOI_COUNTRY_PROV_CITY"),
	}
	return c
}
