package main

import (
	"math/rand"
	"testing"
)

func TestCheckValidString(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"()", true},
		{"(*)", true},
		{"(*))", true},
		{")(", false},
		{"", true},
		{"*", true},
		{"((*", false},
		{"(((**", false},
		{"(**)", true},
		{"*)(", false},
		{"**))", true},
		{"(", false},
	}
	for _, c := range cases {
		if got := CheckValidString(c.in); got != c.want {
			t.Errorf("CheckValidString(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

// reference is an independent O(n²) DP over (position, open count).
func reference(s string) bool {
	n := len(s)
	reachable := map[int]bool{0: true} // possible open counts
	for i := 0; i < n; i++ {
		next := map[int]bool{}
		for open := range reachable {
			switch s[i] {
			case '(':
				next[open+1] = true
			case ')':
				if open > 0 {
					next[open-1] = true
				}
			default:
				next[open+1] = true
				if open > 0 {
					next[open-1] = true
				}
				next[open] = true
			}
		}
		reachable = next
		if len(reachable) == 0 {
			return false
		}
	}
	return reachable[0]
}

// Property: agree with the reference DP across random strings.
func TestAgainstReference(t *testing.T) {
	rng := rand.New(rand.NewSource(7))
	alphabet := []byte("()*")
	for trial := 0; trial < 500; trial++ {
		n := rng.Intn(24)
		buf := make([]byte, n)
		for i := range buf {
			buf[i] = alphabet[rng.Intn(3)]
		}
		s := string(buf)
		if got, want := CheckValidString(s), reference(s); got != want {
			t.Fatalf("mismatch on %q: fast %v, reference %v", s, got, want)
		}
	}
}
