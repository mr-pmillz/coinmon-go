package main

import "testing"

func Test_getJSON(t *testing.T) {
	type args struct {
		url    string
		target interface{}
	}
	coinData := new(CoinData)
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test_getJSON 1", args: args{
			url:    "https://api.coincap.io/v2/assets?limit=10",
			target: coinData,
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := getJSON(tt.args.url, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("getJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

//func Test_prettyFloatString(t *testing.T) {
//	type args struct {
//		num                string
//		percent            bool
//		nearestThousandFMT bool
//		prec4              bool
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    string
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := prettyFloatString(tt.args.num, tt.args.percent, tt.args.nearestThousandFMT, tt.args.prec4)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("prettyFloatString() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("prettyFloatString() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_printTable(t *testing.T) {
//	type args struct {
//		coinData *CoinData
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := printTable(tt.args.coinData); (err != nil) != tt.wantErr {
//				t.Errorf("printTable() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func Test_printTotalMarketCap(t *testing.T) {
//	type args struct {
//		coinData *CoinData
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := printTotalMarketCap(tt.args.coinData); (err != nil) != tt.wantErr {
//				t.Errorf("printTotalMarketCap() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func Test_main(t *testing.T) {
//	tests := []struct {
//		name string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			main()
//		})
//	}
//}
