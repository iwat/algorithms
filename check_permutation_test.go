package algorithms

import (
	"testing"
)

func TestCheckPermutation(t *testing.T) {
	assertPermutation("abcdefg", "gfedcba", t)
	refutePermutation("abcdefg", "gfedcb", t)
	assertPermutation("abcdefgabcdefg", "gfedcbagfedcba", t)
	refutePermutation("abcdefgabcdefg", "gfedcb", t)
}

func assertPermutation(sample, source string, t *testing.T) {
	if !CheckPermutation(sample, source) {
		t.Fatal("assert", sample, source)
	}
}

func refutePermutation(sample, source string, t *testing.T) {
	if CheckPermutation(sample, source) {
		t.Fatal("refute", sample, source)
	}
}
