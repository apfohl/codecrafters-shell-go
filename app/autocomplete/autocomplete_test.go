package autocomplete

import "testing"

func Test_findLongestPrefix(t *testing.T) {
	type args struct {
		commands []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "two commands with prefix",
			args: args{commands: []string{
				"xyz_foo",
				"xyz_bar",
			}},
			want: "xyz_",
		},
		{
			name: "two commands and first is prefix",
			args: args{commands: []string{
				"clang",
				"clang-format",
			}},
			want: "clang",
		},
		{
			name: "one command",
			args: args{commands: []string{
				"clang",
			}},
			want: "clang",
		},
		{
			name: "zero commands",
			args: args{commands: []string{}},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLongestPrefix(tt.args.commands); got != tt.want {
				t.Errorf("findLongestPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
