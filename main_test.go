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
		inlcudes  string
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
				inlcudes:  "",
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
				inlcudes:  "",
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
				inlcudes:  "",
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
				inlcudes:  "",
			},
			want: []string{"version 4.5.6", "version 2.3.4", "version 1.2.3"},
		},
		{
			name: "versions two columns with chara 'v'",
			args: args{
				lines:     []string{"version123 v1.2.3", "version456 v4.5.6", "version012 v0.1.2", "version234 v2.3.4"},
				regex:     "[.]",
				column:    1,
				rank:      1000,
				asc:       true,
				delimiter: " ",
				inlcudes:  "",
			},
			want: []string{"version012 v0.1.2", "version123 v1.2.3", "version234 v2.3.4", "version456 v4.5.6"},
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
				inlcudes:  "",
			},
			want: []string{"999 0.9.8", "123 1.2.3", "120 4.5.6"},
		},
		{
			name: "includes test",
			args: args{
				lines:     []string{"123 v1.2.3", "381 vv4.5.6", "999 v0.9.8", "120 v4.5.6", "000 v3.2.1", "321 v3.3.3-1234567", "123 2.2.2"},
				regex:     "[.]",
				column:    1,
				rank:      1000,
				asc:       true,
				delimiter: " ",
				inlcudes:  "^v[0-9]+[.][0-9]+[.][0-9]+$",
			},
			want: []string{"999 v0.9.8", "123 v1.2.3", "000 v3.2.1", "120 v4.5.6"},
		},
		{
			name: "includes test with desc",
			args: args{
				lines:     []string{"123 v1.2.3", "381 vv4.5.6", "999 v0.9.8", "120 v4.5.6", "000 v3.2.1", "321 v3.3.3-1234567", "123 2.2.2"},
				regex:     "[.]",
				column:    1,
				rank:      1000,
				asc:       false,
				delimiter: " ",
				inlcudes:  "^v[0-9]+[.][0-9]+[.][0-9]+$",
			},
			want: []string{"120 v4.5.6", "000 v3.2.1", "123 v1.2.3", "999 v0.9.8"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			delimiter = tt.args.delimiter
			column = tt.args.column
			asc = tt.args.asc
			includes = tt.args.inlcudes
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
