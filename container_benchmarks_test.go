package godi

import "testing"

func BenchmarkContainer_New(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c := Container{}
		c.New()
	}
}
