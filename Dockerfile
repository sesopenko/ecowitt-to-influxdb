FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY internal/ecowitt/*.go ./internal/ecowitt/
COPY internal/influx/*.go ./internal/influx/

RUN CGO_ENABLED=0 GOOS=linux go build -o /ecowitt-to-influxdb
EXPOSE  20555
CMD ["/ecowitt-to-influxdb"]