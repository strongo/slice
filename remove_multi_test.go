package slice

import (
	"fmt"
	"strings"
	"testing"
)

func TestRemove(t *testing.T) {
	t.Parallel()

	for _, testCase := range []struct {
		source       []string
		remove       []string
		result       []string
		removedCount int
	}{
		{removedCount: 0, source: []string{"v1", "v2", "v3"}, remove: []string{}, result: []string{"v1", "v2", "v3"}},
		{removedCount: 0, source: []string{"v1", "v2", "v3"}, remove: []string{"v0"}, result: []string{"v1", "v2", "v3"}},
		{removedCount: 1, source: []string{"v1", "v2", "v3"}, remove: []string{"v1"}, result: []string{"v2", "v3"}},
		{removedCount: 1, source: []string{"v1", "v2", "v3"}, remove: []string{"v2"}, result: []string{"v1", "v3"}},
		{removedCount: 1, source: []string{"v1", "v2", "v3"}, remove: []string{"v3"}, result: []string{"v1", "v2"}},
		{removedCount: 2, source: []string{"v1", "v2", "v3"}, remove: []string{"v1", "v2"}, result: []string{"v3"}},
		{removedCount: 2, source: []string{"v1", "v2", "v3"}, remove: []string{"v1", "v3"}, result: []string{"v2"}},
		{removedCount: 2, source: []string{"v1", "v2", "v3"}, remove: []string{"v2", "v3"}, result: []string{"v1"}},
		{removedCount: 4, source: []string{"v1", "v2", "v2", "v3", "v2", "v2"}, remove: []string{"v2"}, result: []string{"v1", "v3"}},
	} {
		t.Run(fmt.Sprintf("source=%s_remove=%s", strings.Join(testCase.source, ","), strings.Join(testCase.remove, ",")), func(t *testing.T) {
			result, removedCount := RemoveMulti(testCase.source, testCase.remove)
			if removedCount != testCase.removedCount {
				t.Errorf("RemoveMulti(%v, %v) expected to return removedCount=%v, got %v", testCase.source, testCase.remove, testCase.removedCount, removedCount)
			}
			if !Equal(result, testCase.result) {
				t.Errorf("RemoveMulti(%v, %v) expected to return %v, got %v", testCase.source, testCase.remove, testCase.result, result)
			}
		})
	}
}
