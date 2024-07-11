# sesopenko/ecowitt-to-influxdb

This project was created so I can stream data from my Ecowitt WP1200 weather gateway to my influxdb

## Requirements

* Influxdb 2 or later (tested with 2.7)
* Go version 22 or later
* Docker if you want to build the docker image yourself.

## Configuration

Requires the following environment variables:

* ETOI_INFLUX_URL: full url to your influx service, including http:// or https://
* ETOI_AUTH_TOKEN: authorization token for your influx service
* ETOI_ORG: influx organization to use
* ETOI_BUCKET: influx bucket to use
* ETOI_COUNTRY_PROV_CITY: country/prov/city string to use. Recommend the following example: CA-AB:Medicine Hat
  * ISO-3166-1 2 letter country code
  * a dash
  * ISO-3166-2 2 letter province/state/region code
  * a colon
  * city full name

### Configuring Ecowitt

1. Download the ecowitt app and add your console/gateway
2. When viewing your sensor, press the triple dot and choose "Others"
3. Press DIY Upload Servers
4. Choose Customized
5. Protocol: `Ecowitt`
6. Server IP/host: the ip address of the server running this app/container
7. path: `/data/report/`
8. port: `20555`
9. Upload Interval: Whatever frequency makes sense for you.  I'm using 16 seconds for now.
10. Press Save

## Influx Data

```
_measurement: local_weather
  country_prov_city: <YOUR COUNTRY PROV CITY>
    internal_temperature_fahrenheight
    barometric_pressure_absolute_inhg
    barometric_pressure_relative_inhg
    humidity_indoors
```

Barrometric pressure relative doesn't seem to work properly for me because I don't think my ecowitt gateway calculates
it when sending it. You'll have to look up how to calculate relative with an influxdb map.

## Equipment I'm using

* Ecowitt GW1200 Wireless Gateway and internal sensor
  * Measures temp, humidity, and pressure on a short lead.
  * Connects to wifi and configured with the app to send measurements "out to the cloud"
* Synolgoy NAS
  * Runs the influxdb docker container
  * Runs this app as a docker container

## Help & Bug Reports

This software comes with zero warranty or guarantees. Use at your own risk.

Once I have this project running I will likely not add features, which is why I'm open sourcing it. If you want it to do
more than what it's currently capable of then fork the repo change it. You may ask questions via the issue reporter, but
I don't have time to respond to feature requests.

## Licensed GNU GPL v3

This software is licensed under the GNU GPL v3 license, which is included in [LICENSE.txt](LICENSE.txt).

## Copyright

Copyright (c) Sean Esopenko 2024 All Rights Reserved