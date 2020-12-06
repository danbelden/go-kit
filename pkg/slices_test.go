package pkg

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStringSliceUnique(t *testing.T) {
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
			got := StringSliceUnique(tc.input)
			if diff := cmp.Diff(got, tc.want); diff != "" {
				t.Errorf("-got +want: %s", diff)
			}
		})
	}
}
