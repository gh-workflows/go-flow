package wordgame

import (
	"testing"
)

func TestPalindome(t *testing.T) {
	var words = []string{"kayak", "bob"}
	for _, word := range words {
		if !IsPalindrome(word) {
			t.Errorf("'%s' incorrectly not identifed as palindrome", word)
		}
	}

}

func TestNonPalindome(t *testing.T) {
	var words = []string{"apple", "Orange", "pear"}
	for _, word := range words {
		if IsPalindrome(word) {
			t.Errorf("'%s' incorrectly not identifed as palindrome", word)
		}
	}
}
