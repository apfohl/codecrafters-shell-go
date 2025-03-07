package arguments_test

import (
	"reflect"
	"testing"

	"github.com/codecrafters-io/shell-starter-go/app/arguments"
)

func TestParseArgs(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "single word",
			args: args{input: "foo"},
			want: []string{"foo"},
		},
		{
			name: "two words",
			args: args{input: "foo bar"},
			want: []string{"foo", "bar"},
		},
		{
			name: "single quoted word",
			args: args{input: "'foo'"},
			want: []string{"foo"},
		},
		{
			name: "two words with more spaces",
			args: args{input: "world     example"},
			want: []string{"world", "example"},
		},
		{
			name: "complex example",
			args: args{input: "a  'b c'  g'g''h'h"},
			want: []string{"a", "b c", "gghh"},
		},
		{
			name: "double quotes simple",
			args: args{input: "\"quz  hello\"  \"bar\""},
			want: []string{"quz  hello", "bar"},
		},
		{
			name: "double quotes more complex",
			args: args{input: "\"bar\"  \"shell's\"  \"foo\""},
			want: []string{"bar", "shell's", "foo"},
		},
		{
			name: "single quotes more complex",
			args: args{input: "'bar'  'shell\"s'  'foo'"},
			want: []string{"bar", "shell\"s", "foo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arguments.ParseArgs(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
