package counter

import (
	"testing"
	"time"
)

var counter = New(time.Second, 10)
func BenchmarkWindowSlidingCounter_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		counter.Add(3)
	}
}
