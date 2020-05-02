package main

import (
	"io"
	"strings"
	"testing"
)

func TestMarshal(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		// test empty
		{
			in:   "",
			want: "null",
		},
		// test json
		{
			in:   `{"test":"test"}`,
			want: "test: test",
		},
		// test json
		{
			in:   "test: test",
			want: `{"test":"test"}`,
		},
	}

	for _, test := range tests {
		bytes, err := marshal([]byte(test.in))
		if err != nil {
			t.Error(err)
		}

		got := strings.TrimSpace(string(bytes))
		if test.want != got {
			t.Errorf("want: %#v, got: %#v", test.want, got)
		}
	}
}

func TestRead(t *testing.T) {
	tests := []struct {
		stdin io.Reader
		args  []string
		want  string
		err   error
	}{
		// test empty args and stdin
		{
			stdin: strings.NewReader(""),
			args:  []string{"yj"},
			want:  "",
		},
		// test stdin
		{
			stdin: strings.NewReader(`{"test":"test"}`),
			args:  []string{"yj"},
			want:  `{"test":"test"}`,
		},
		// test file arg
		{
			args: []string{"yj", "testdata/test.json"},
			want: `{"test":"test"}`,
		},
		// test error
		{
			args: []string{"yj", "too", "many", "args"},
			err:  errTooManyArgs,
		},
	}

	for _, test := range tests {
		bytes, err := read(test.stdin, test.args)
		if err != nil && err != test.err {
			t.Error(err)
		}

		got := strings.TrimSpace(string(bytes))
		if test.want != got {
			t.Errorf("want: %#v, got: %#v", test.want, got)
		}
	}
}
