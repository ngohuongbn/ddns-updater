package main

import (
	"testing"
)

func Test_regexIP(t *testing.T) {
	cases := []struct {
		s  string
		ip string
	}{
		{
			"dsa dskd 2 | 32.43 210.125.56.230 dsad",
			"210.125.56.230",
		},
	}
	for _, c := range cases {
		out := regexIP(c.s)
		if out != c.ip {
			t.Errorf("regexIP(%s) == %s want %s", c.s, out, c.ip)
		}
	}
}
