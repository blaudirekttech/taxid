package taxid

import (
	"testing"
)

func TestValidateTaxID(t *testing.T) {
	type args struct {
		country TaxCountry
		taxID   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid German Tax ID",
			args: args{
				country: Germany,
				taxID:   "42 445 871 363",
			},
			wantErr: false,
		},
		{
			name: "Valid Austrian Tax ID",
			args: args{
				country: Austria,
				taxID:   "931736581",
			},
			wantErr: false,
		},
		{
			name: "Invalid German Tax ID",
			args: args{
				country: Germany,
				taxID:   "41 445 871 363",
			},
			wantErr: true,
		},
		{
			name: "Invalid Austrian Tax ID",
			args: args{
				country: Austria,
				taxID:   "41 445 871 363",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Validate(tt.args.country, tt.args.taxID); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
