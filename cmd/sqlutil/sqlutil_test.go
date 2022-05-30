package sqlutil_test

import (
	"testing"

	"github.com/madhab452/csvtosql/cmd/sqlutil"
)

func TestToColumnName(t *testing.T) {
	tests := []struct {
		description string
		wantError   error
		given       string
		want        string
	}{
		{
			description: "Single space between words",
			wantError:   nil,
			given:       "test col",
			want:        "test_col",
		},
		{
			description: "Double space between words",
			wantError:   nil,
			given:       "test  col",
			want:        "test_col",
		},
		{
			description: "Single char.",
			wantError:   nil,
			given:       "x",
			want:        "x",
		},
		{
			description: "Empty value",
			wantError:   nil,
			given:       "",
			want:        "unknown_1",
		},
		{
			description: "Empty value",
			wantError:   nil,
			given:       "",
			want:        "unknown_2",
		},
		{
			description: "special characters",
			wantError:   nil,
			given:       "test*** ___'''***'''special  col",
			want:        "test_special_col",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := sqlutil.ToColumnName(tt.given)
			if got != tt.want {
				t.Errorf("unexpected, got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestToTableName(t *testing.T) {
	tests := []struct {
		description string
		wantError   error
		given       string
		want        string
	}{
		{
			description: "",
			wantError:   nil,
			given:       "./csvs/BTC-USD-2.csv",
			want:        "btc_usd_2",
		},
		{
			description: "",
			wantError:   nil,
			given:       "/Users/lorem/Downloads/Sample Test (test 101).csv",
			want:        "sample_test_test_101_",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := sqlutil.ToTableName(tt.given)
			if got != tt.want {
				t.Errorf("unexpected, got: %v, want: %v", got, tt.want)
			}
		})
	}
}
