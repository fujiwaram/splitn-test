package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplitN_A(t *testing.T) {
	type args struct {
		s   string
		sep string
		n   int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "-1",
			args: args{
				"A,B,C,D",
				",",
				-1,
			},
			want: []string{"A", "B", "C", "D"},
		},
		{
			name: "0",
			args: args{
				"A,B,C,D",
				",",
				0,
			},
			want: nil,
		},
		{
			name: "1",
			args: args{
				"A,B,C,D",
				",",
				1,
			},
			want: []string{"A,B,C,D"},
		},
		{
			name: "2",
			args: args{
				"A,B,C,D",
				",",
				2,
			},
			want: []string{"A", "B,C,D"},
		},
		{
			name: "3",
			args: args{
				"A,B,C,D",
				",",
				3,
			},
			want: []string{"A", "B", "C,D"},
		},
		{
			name: "4",
			args: args{
				"A,B,C,D",
				",",
				4,
			},
			want: []string{"A", "B", "C", "D"},
		},
		{
			name: "empty_string_0",
			args: args{
				"",
				",",
				0,
			},
			want: nil,
		},
		{
			name: "empty_string_1",
			args: args{
				"",
				",",
				1,
			},
			want: []string{""},
		},
		{
			name: "empty_string_2",
			args: args{
				"",
				",",
				2,
			},
			want: []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings.SplitN(tt.args.s, tt.args.sep, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("strings.SplitN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitN_B(t *testing.T) {
	type args struct {
		s   string
		sep string
		n   int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "-1", args: args{"A,B,C,D", ",", -1}, want: []string{"A", "B", "C", "D"}},
		{name: "0", args: args{"A,B,C,D", ",", 0}, want: nil},
		{name: "1", args: args{"A,B,C,D", ",", 1}, want: []string{"A,B,C,D"}},
		{name: "2", args: args{"A,B,C,D", ",", 2}, want: []string{"A", "B,C,D"}},
		{name: "3", args: args{"A,B,C,D", ",", 3}, want: []string{"A", "B", "C,D"}},
		{name: "4", args: args{"A,B,C,D", ",", 4}, want: []string{"A", "B", "C", "D"}},
		{name: "empty_string_0", args: args{"", ",", 0}, want: nil},
		{name: "empty_string_1", args: args{"", ",", 1}, want: []string{""}},
		{name: "empty_string_2", args: args{"", ",", 2}, want: []string{""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings.SplitN(tt.args.s, tt.args.sep, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("strings.SplitN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitN_C(t *testing.T) {
	type args struct {
		s   string
		sep string
		n   int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "-1" /*       */, args: args{"A,B,C,D", ",", -1}, want: []string{"A", "B", "C", "D"}},
		{name: "0" /*        */, args: args{"A,B,C,D", ",", 0}, want: nil},
		{name: "1" /*        */, args: args{"A,B,C,D", ",", 1}, want: []string{"A,B,C,D"}},
		{name: "2" /*        */, args: args{"A,B,C,D", ",", 2}, want: []string{"A", "B,C,D"}},
		{name: "3" /*        */, args: args{"A,B,C,D", ",", 3}, want: []string{"A", "B", "C,D"}},
		{name: "4" /*        */, args: args{"A,B,C,D", ",", 4}, want: []string{"A", "B", "C", "D"}},
		{name: "empty_string_0", args: args{"" /*   */, ",", 0}, want: nil},
		{name: "empty_string_1", args: args{"" /*   */, ",", 1}, want: []string{""}},
		{name: "empty_string_2", args: args{"" /*   */, ",", 2}, want: []string{""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings.SplitN(tt.args.s, tt.args.sep, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("strings.SplitN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitN_D(t *testing.T) {
	type args struct {
		s   string
		sep string
		n   int
	}
	type test struct {
		name string
		args args
		want []string
	}
	makeTest := func(name string, args args, want []string) test {
		return test{name: name, args: args, want: want}
	}
	tests := []test{
		makeTest("-1" /*       */, args{"A,B,C,D", ",", -1}, []string{"A", "B", "C", "D"}),
		makeTest("0" /*        */, args{"A,B,C,D", ",", 0}, nil),
		makeTest("1" /*        */, args{"A,B,C,D", ",", 1}, []string{"A,B,C,D"}),
		makeTest("2" /*        */, args{"A,B,C,D", ",", 2}, []string{"A", "B,C,D"}),
		makeTest("3" /*        */, args{"A,B,C,D", ",", 3}, []string{"A", "B", "C,D"}),
		makeTest("4" /*        */, args{"A,B,C,D", ",", 4}, []string{"A", "B", "C", "D"}),
		makeTest("empty_string_0", args{"" /*   */, ",", 0}, nil),
		makeTest("empty_string_1", args{"" /*   */, ",", 1}, []string{""}),
		makeTest("empty_string_2", args{"" /*   */, ",", 2}, []string{""}),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings.SplitN(tt.args.s, tt.args.sep, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("strings.SplitN() = %v, want %v", got, tt.want)
			}
		})
	}
}
