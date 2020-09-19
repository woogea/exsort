package main

import (
	"strings"
	"testing"
)

func Test_exsort(t *testing.T) {
	type args struct {
		lines     []string
		regex     string
		column    int
		rank      int
		asc       bool
		delimiter string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "versions single column",
			args: args{
				lines:     []string{"1.2.3", "4.5.6", "2.3.4"},
				regex:     "[.]",
				column:    0,
				rank:      1000,
				asc:       false,
				delimiter: " ",
			},
			want: []string{"4.5.6", "2.3.4", "1.2.3"},
		},
		{
			name: "versions single column with extra lines",
			args: args{
				lines:     []string{"1.2.3", "4.5.6", "2.3.4", ""},
				regex:     "[.]",
				column:    0,
				rank:      1000,
				asc:       false,
				delimiter: " ",
			},
			want: []string{"4.5.6", "2.3.4", "1.2.3"},
		},
		{
			name: "versions single column, many",
			args: args{
				lines:     []string{"1.2.3", "4.5.6", "2.3.4", "0.0.1", "10.20.30", "0.9.10"},
				regex:     "[.]",
				column:    0,
				rank:      1000,
				asc:       false,
				delimiter: " ",
			},
			want: []string{"10.20.30", "4.5.6", "2.3.4", "1.2.3", "0.9.10", "0.0.1"},
		},
		{
			name: "versions two columns",
			args: args{
				lines:     []string{"version 1.2.3", "version 4.5.6", "version 2.3.4"},
				regex:     "[.]",
				column:    1,
				rank:      1000,
				asc:       false,
				delimiter: " ",
			},
			want: []string{"version 4.5.6", "version 2.3.4", "version 1.2.3"},
		},
		{
			name: "versions two columns with chara 'v'",
			args: args{
				lines:     []string{"version123 v1.2.3", "version456 v4.5.6", "version234 v2.3.4"},
				regex:     "[.]",
				column:    1,
				rank:      1000,
				asc:       true,
				delimiter: " ",
			},
			want: []string{"version123 v1.2.3", "version234 v2.3.4", "version456 v4.5.6"},
		},
		{
			name: "major version is 0",
			args: args{
				lines:     []string{"123 1.2.3", "999 0.9.8", "120 4.5.6"},
				regex:     "[.]",
				column:    1,
				rank:      1000,
				asc:       true,
				delimiter: " ",
			},
			want: []string{"999 0.9.8", "123 1.2.3", "120 4.5.6"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			delimiter = tt.args.delimiter
			column = tt.args.column
			asc = tt.args.asc
			lines := exsort(tt.args.lines, tt.args.regex, tt.args.column, tt.args.rank, tt.args.asc)
			b := true
			for i := range lines {
				if strings.Compare(lines[i], tt.want[i]) != 0 {
					b = false
					break
				}
			}
			if !b {
				t.Errorf("wrong result: %v, want %v", lines, tt.want)
			}
		})
	}
}
