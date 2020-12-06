package strings

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUniqueSlice(t *testing.T) {
	testCases := map[string]struct {
		input []string
		want  []string
	}{
		"empty": {
			input: nil,
			want:  nil,
		},
		"one": {
			input: []string{"test"},
			want:  []string{"test"},
		},
		"two": {
			input: []string{"test", "test"},
			want:  []string{"test"},
		},
		"three": {
			input: []string{"test", "test", "Test"},
			want:  []string{"test", "Test"},
		},
	}
	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			got := UniqueSlice(tc.input)
			if diff := cmp.Diff(got, tc.want); diff != "" {
				t.Errorf("-got +want: %s", diff)
			}
		})
	}
}

func TestSearchWord(t *testing.T) {
	testCases := map[string]struct {
		text string
		word string
		want bool
	}{
		"empty": {
			text: "",
			word: "",
			want: false,
		},
		"inside word": {
			text: "the quick brown fox",
			word: "uick",
			want: false,
		},
		"exact match": {
			text: "the quick brown fox",
			word: "quick",
			want: true,
		},
		"case-insensitive match": {
			text: "the quick brown fox",
			word: "Quick",
			want: false,
		},
	}
	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			got := SearchWord(tc.text, tc.word)
			if got != tc.want {
				t.Errorf("got: %t, want: %t", got, tc.want)
			}
		})
	}
}
