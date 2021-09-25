package ints

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUniqueSlice(t *testing.T) {
	testCases := map[string]struct {
		input []int
		want  []int
	}{
		"empty": {
			input: nil,
			want:  nil,
		},
		"one": {
			input: []int{1},
			want:  []int{1},
		},
		"two": {
			input: []int{1, 1},
			want:  []int{1},
		},
		"three": {
			input: []int{1, 1, 2},
			want:  []int{1, 2},
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
