package utils

import "testing"

func TestRandString(t *testing.T) {
	type args struct {
		size    int
		set     int
		include string
		exclude string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Should get empty", args{0, GenRandAll, "", "1234"}, "", false},
		{"Should throw error when exclude too high", args{0, GenRandNone, "", "1234"}, "", true},
		{"Should generate 4 zeros string when exclude", args{4, GenRandDigit, "", "123456789"}, "0000", false},
		{"Should generate 4 zeros string when include", args{4, GenRandNone, "0", ""}, "0000", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RandString(tt.args.size, tt.args.set, tt.args.include, tt.args.exclude)
			if (err != nil) != tt.wantErr {
				t.Errorf("RandString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RandString() = %v, want %v", got, tt.want)
			}
		})
	}
}
