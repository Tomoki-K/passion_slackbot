package main

import "testing"

// table driven test
var testData = []struct {
	in       string
	expected bool
}{
	{"パッション", true},
	{"got passion??", true},
	{"パッションフルーツ", true},
	{"やきそば", false},
	{"ぱっ しょん", false},
	{"PassioN", true},
	{"", false},
	{"   ", false},
}

func TestIncludesPassion(t *testing.T) {
	for _, c := range testData {
		out := IncludesPassion(c.in)
		if out != c.expected {
			t.Errorf("Error: Got %v\n want %v", out, c.expected)
		}
	}
}
