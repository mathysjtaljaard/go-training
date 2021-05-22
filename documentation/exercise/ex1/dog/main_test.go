package dog

import "testing"

func TestYears(t *testing.T) {
	type args struct {
		dogYears int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "happy Dog", args: args{dogYears: 1}, want: 7, wantErr: false},
		{name: "sad Dog", args: args{dogYears: 0}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Years(tt.args.dogYears)
			if (err != nil) != tt.wantErr {
				t.Errorf("Years() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Years() = %v, want %v", got, tt.want)
			}
		})
	}
}
