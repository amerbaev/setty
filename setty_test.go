package setty

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Intersect_StingKey(t *testing.T) {
	type testCase struct {
		name                 string
		set1, set2, expected map[string]struct{}
	}

	testCases := []testCase{
		{
			name:     "empty sets",
			set1:     map[string]struct{}{},
			set2:     map[string]struct{}{},
			expected: map[string]struct{}{},
		},
		{
			name:     "one empty set",
			set1:     map[string]struct{}{},
			set2:     map[string]struct{}{"a": {}},
			expected: map[string]struct{}{},
		},
		{
			name:     "no matches",
			set1:     map[string]struct{}{"a": {}, "b": {}},
			set2:     map[string]struct{}{"c": {}, "d": {}},
			expected: map[string]struct{}{},
		},
		{
			name:     "one match",
			set1:     map[string]struct{}{"a": {}, "b": {}},
			set2:     map[string]struct{}{"b": {}, "c": {}},
			expected: map[string]struct{}{"b": {}},
		},
		{
			name:     "full match",
			set1:     map[string]struct{}{"a": {}, "b": {}},
			set2:     map[string]struct{}{"b": {}, "a": {}},
			expected: map[string]struct{}{"a": {}, "b": {}},
		},
	}

	for _, tc := range testCases {
		t.Run("string"+tc.name, func(t *testing.T) {
			result := Intersect(tc.set1, tc.set2)
			assert.Equal(t, tc.expected, result)
		})
	}
}
