package ecowitt

import (
	"errors"
	"net/url"
	"strconv"
	"time"
)

type EcowittReading struct {
	BarometricPressureAbsolute   float64
	BarometicPressureRelative    float64
	HumidityIndoors              float64
	TemperatureIndoorsFahrenheit float64
	ReadingTime                  time.Time
}

func BuildReading(readingMap map[string]string) (EcowittReading, error) {
	r := EcowittReading{
		BarometricPressureAbsolute:   0,
		BarometicPressureRelative:    0,
		HumidityIndoors:              0,
		TemperatureIndoorsFahrenheit: 0,
		ReadingTime:                  time.Time{},
	}

	if baromAbs, err := strconv.ParseFloat(readingMap["baromabsin"], 64); err != nil {
		return r, errors.New("cannot parse absolute barometric pressure")
	} else {
		r.BarometricPressureAbsolute = baromAbs
	}
	if baromRel, err := strconv.ParseFloat(readingMap["baromrelin"], 64); err != nil {
		return r, errors.New("cannot parse relative barometric pressure")
	} else {
		r.BarometicPressureRelative = baromRel
	}

	if humidIndoors, err := strconv.ParseFloat(readingMap["humidityin"], 64); err != nil {
		return r, errors.New("cannot parse indoors humidity")
	} else {
		r.HumidityIndoors = humidIndoors
	}

	if tempInF, err := strconv.ParseFloat(readingMap["tempinf"], 64); err != nil {
		return r, errors.New("cannot parse indoors temperature (F)")
	} else {
		r.TemperatureIndoorsFahrenheit = tempInF
	}

	layout := "2006-01-02 15:04:05"
	if parsedTime, err := time.Parse(layout, readingMap["dateutc"]); err != nil {
		return r, errors.New("cannot parse reading dateutc")
	} else {
		r.ReadingTime = parsedTime
	}

	return r, nil
}

func BuildReadingMap(params url.Values) map[string]string {
	readingMap := make(map[string]string)

	for key, values := range params {
		if len(values) > 0 {
			readingMap[key] = values[0]
		}
	}
	return readingMap
}

const EcowittMeasurementHttpPath = "/data/report/"
