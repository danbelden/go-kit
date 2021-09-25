package int64s

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUniqueSlice(t *testing.T) {
	testCases := map[string]struct {
		input []int64
		want  []int64
	}{
		"empty": {
			input: nil,
			want:  nil,
		},
		"one": {
			input: []int64{1},
			want:  []int64{1},
		},
		"two": {
			input: []int64{1, 1},
			want:  []int64{1},
		},
		"three": {
			input: []int64{1, 1, 2},
			want:  []int64{1, 2},
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
