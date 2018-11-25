package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/aquilax/time66"
)

func main() {
	var lat = flag.Float64("lat", 0, "Latitude")
	var lon = flag.Float64("lon", 0, "Longitude")
	var offset = flag.Float64("offset", 0, "UTC Offset")
	flag.Parse()

	t, err := time66.GetTime(*lat, *lon, *offset, time.Now())
	if err != nil {
		panic(err)
	}
	fmt.Println(t.Format(time.RFC3339))
}
