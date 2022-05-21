package qb_test

import (
	"testing"

	"github.com/madhab452/csvtosql/cmd/qb"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		description string
		wantError   error
		want        string
		cols        []string
		vals        [][]string
		tblname     string
	}{
		{
			description: "Create Table Query",
			wantError:   nil,
			want:        `INSERT INTO test (col_1, col_2) VALUES ('val1','val2'), ('val3','val4');`,
			tblname:     "test",
			cols:        []string{"col_1", "col_2"},
			vals:        [][]string{{"val1", "val2"}, {"val3", "val4"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			createTblQuery := qb.Insert(func(qb *qb.InsertBuilder) {
				qb.Table(tt.tblname)
				for _, col := range tt.cols {
					qb.AddCol(col)
				}
				for _, val := range tt.vals {
					qb.AddRow(val)
				}
			}).ToSql()

			if createTblQuery != tt.want {
				t.Errorf("unexpected, got %v, want %v", createTblQuery, tt.want)
			}
		})
	}
}
