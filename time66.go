package time66

import (
	"time"

	"github.com/kelvins/sunrisesunset"
)

var cycle time.Duration

func getSunriseSunset(lat, lon, offset float64, t time.Time) (time.Time, time.Time, error) {
	p := sunrisesunset.Parameters{
		Latitude:  lat,
		Longitude: lon,
		UtcOffset: offset,
		Date:      t,
	}
	sunrise, sunset, err := p.GetSunriseSunset()
	if err != nil {
		return time.Now(), time.Now(), err
	}
	sunrise = time.Date(t.Year(), t.Month(), t.Day(), sunrise.Hour(), sunrise.Minute(), sunrise.Second(), sunrise.Nanosecond(), t.Location())
	sunset = time.Date(t.Year(), t.Month(), t.Day(), sunset.Hour(), sunset.Minute(), sunset.Second(), sunset.Nanosecond(), t.Location())
	return sunrise, sunset, err
}

// GetTime returns 6/6 time for given location and local time
func GetTime(lat, lon, offset float64, t time.Time) (time.Time, error) {
	sunrise, sunset, err := getSunriseSunset(lat, lon, offset, t)
	if err != nil {
		return time.Now(), err
	}
	// println(sunset.String())
	var duration time.Duration
	var start time.Time
	var restart time.Time
	if t.Before(sunrise) {
		_, ysunset, err := getSunriseSunset(lat, lon, offset, t.AddDate(0, 0, -1))
		if err != nil {
			return time.Now(), err
		}
		// night before sunrise
		// duration := prev day sunset - today's sunrise
		duration = sunrise.Sub(ysunset)
		start = ysunset
		restart = time.Date(ysunset.Year(), ysunset.Month(), ysunset.Day(), 18, 0, 0, 0, t.Location())
	} else if t.After(sunset) {
		tsunrise, _, err := getSunriseSunset(lat, lon, offset, t.AddDate(0, 0, +1))
		if err != nil {
			return time.Now(), err
		}
		// night after sunset
		// TODO: today's sunset -  next day sunrise
		duration = tsunrise.Sub(sunset)
		start = sunset
		restart = time.Date(t.Year(), t.Month(), t.Day(), 18, 0, 0, 0, t.Location())
	} else {
		// day
		duration = sunset.Sub(sunrise)
		start = sunrise
		restart = time.Date(t.Year(), t.Month(), t.Day(), 6, 0, 0, 0, t.Location())
	}

	c := float64(duration.Nanoseconds()) / float64(cycle.Nanoseconds())
	nanoseconds := time.Duration(float64(t.Sub(start).Nanoseconds()) / c)
	return restart.Add(time.Nanosecond * nanoseconds), nil
}

func init() {
	cycle, _ = time.ParseDuration("12h")
}
