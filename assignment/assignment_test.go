package assignment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculation(t *testing.T) {

	tests := []struct {
		name    string
		args    string
		want    int
		wantErr bool
	}{
		{
			name:    "Calculation_success_1",
			args:    "10 10 3 4 5",
			want:    49,
			wantErr: false,
		},
		{
			name:    "Calculation_success_2",
			args:    "7 5 2 4",
			want:    15,
			wantErr: false,
		},
		{
			name:    "Calculation_fail_convert_error_3",
			args:    "A 5 2 3 4",
			wantErr: true,
		},
		{
			name:    "Calculation_fail_convert_error_4",
			args:    "7 B 2 3 4",
			wantErr: true,
		},
		{
			name:    "Calculation_fail_convert_error_5",
			args:    "7 5 2 x",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Calculation(tt.args)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, result)
			}
		})
	}
}
