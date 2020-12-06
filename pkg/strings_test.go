package pkg

import "testing"

func TestStringWordSearch(t *testing.T) {
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
			got := StringWordSearch(tc.text, tc.word)
			if got != tc.want {
				t.Errorf("got: %t, want: %t", got, tc.want)
			}
		})
	}
}
