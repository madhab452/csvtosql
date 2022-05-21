package qb_test

import (
	"testing"

	"github.com/madhab452/csvtosql/cmd/qb"
)

func TestCreateTbl(t *testing.T) {
	tests := []struct {
		description string
		wantError   error
		want        string
		cols        []string
		tblname     string
	}{
		{
			description: "Create Table Query",
			wantError:   nil,
			want:        `CREATE TABLE IF NOT EXISTS test (col_1 VARCHAR(255), col_2 VARCHAR(255));`,
			tblname:     "test",
			cols:        []string{"col_1", "col_2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			createTblQuery := qb.CreateTbl(func(qb *qb.CreateTblBuilder) {
				qb.Table(tt.tblname)
				for _, col := range tt.cols {
					qb.AddCol(col)
				}
			}).ToSql()

			if createTblQuery != tt.want {
				t.Errorf("unexpected, got %v, want %v", createTblQuery, tt.want)
			}
		})
	}
}
