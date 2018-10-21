package time66

import (
	"reflect"
	"testing"
	"time"
)

func Test_getTime(t *testing.T) {
	loc := time.FixedZone("UTC+2", 2*60*60)
	type args struct {
		lat    float64
		lon    float64
		offset float64
		t      time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			"Test Stockholm sunrise",
			args{
				lat:    59.3293,
				lon:    18.0686,
				offset: 2.0,
				t:      time.Date(2018, 10, 21, 7, 39, 34, 0, loc),
			},
			time.Date(2018, 10, 21, 6, 0, 0, 0, loc),
			false,
		},
		{
			"Test Stockholm sunset",
			args{
				lat:    59.3293,
				lon:    18.0686,
				offset: 2.0,
				t:      time.Date(2018, 10, 21, 17, 24, 03, 0, loc),
			},
			time.Date(2018, 10, 21, 18, 0, 0, 0, loc),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getTime(tt.args.lat, tt.args.lon, tt.args.offset, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
