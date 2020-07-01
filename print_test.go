package jtlr

import (
	"testing"
)

func TestPrettyPrint(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "a",
			args: args{
				input: `{"a": [134, 2], "b": {"a":1, "b":2}}`,
			},
		},
		{
			name: "b",
			args: args{
				input: `{"a": [134, {"a": 1}, true, [1, 2, 3], false], "b": {"a":1, "b":{"a":1, "b":2}}, "c": true, "d": null}`,
			},
		},
		{
			name: "c",
			args: args{
				input: `{"code":2,"message":"zz\u672a中"}`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrettyPrint(tt.args.input)
		})
	}
}
