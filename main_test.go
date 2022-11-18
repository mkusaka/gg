package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parse(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "https:// prefix",
			args: args{
				args: []string{"https://github.com/mkusaka/gg"},
			},
			want: []string{"get", "github.com/mkusaka/gg"},
		},
		{
			name: "git prefix with .git suffix",
			args: args{
				args: []string{"https://github.com/mkusaka/gg.git"},
			},
			want: []string{"get", "github.com/mkusaka/gg"},
		},
		// TODO: support git@github.com:mkusaka/trending.git format
		//{
		//	name: "git prefix with .git suffix",
		//	args: args{
		//		args: []string{"git@github.com:mkusaka/trending.git"},
		//	},
		//	want: []string{"get", "github.com/mkusaka/gg"},
		//},
		{
			name: "no prefix and suffix",
			args: args{
				args: []string{"github.com/mkusaka/gg"},
			},
			want: []string{"get", "github.com/mkusaka/gg"},
		},
		{
			name: "with -u flag",
			args: args{
				args: []string{"-u", "https://github.com/mkusaka/gg"},
			},
			want: []string{"get", "-u", "github.com/mkusaka/gg"},
		},
		{
			name: "with -u -x flag",
			args: args{
				args: []string{"-u", "-x", "https://github.com/mkusaka/gg"},
			},
			want: []string{"get", "-u", "-x", "github.com/mkusaka/gg"},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {
			got := parse(tt.args.args)
			assert.Equal(t, got, tt.want)
		})
	}
}
