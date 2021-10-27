package utils

import "testing"

func TestRoundPrec(t *testing.T) {
	type args struct {
		x    float64
		prec int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "TestRoundPrec 1", args: args{
			x:    0.82345,
			prec: 0,
		}, want: 1},
		{name: "TestRoundPrec 2", args: args{
			x:    3.82345,
			prec: 0,
		}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundPrec(tt.args.x, tt.args.prec); got != tt.want {
				t.Errorf("RoundPrec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumberFormat(t *testing.T) {
	type args struct {
		number       float64
		decimals     int
		decPoint     string
		thousandsSep string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestNumberFormat 1", args: args{
			number:       143.234234,
			decimals:     2,
			decPoint:     ".",
			thousandsSep: ",",
		}, want: "143.23"},
		{name: "TestNumberFormat 1", args: args{
			number:       1234567.234234,
			decimals:     2,
			decPoint:     ".",
			thousandsSep: ",",
		}, want: "1,234,567.23"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberFormat(tt.args.number, tt.args.decimals, tt.args.decPoint, tt.args.thousandsSep); got != tt.want {
				t.Errorf("NumberFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoundInt(t *testing.T) {
	type args struct {
		input float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "TestRountInt 1", args: args{input: 12345.12345}, want: 12345},
		{name: "TestRountInt 2", args: args{input: 12345.9999}, want: 12346},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundInt(tt.args.input); got != tt.want {
				t.Errorf("RoundInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatNumber(t *testing.T) {
	type args struct {
		input float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestFormatNumber", args: args{input: 12345.12345}, want: "12,345.00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatNumber(tt.args.input); got != tt.want {
				t.Errorf("FormatNumber() = %v\n want %v", got, tt.want)
			}
		})
	}
}

func TestNearestThousandFormat(t *testing.T) {
	type args struct {
		num float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestNearestThousandFormat", args: args{num: 1234567.1234567}, want: "1.2M"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NearestThousandFormat(tt.args.num); got != tt.want {
				t.Errorf("NearestThousandFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
