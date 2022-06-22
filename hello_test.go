package demo

import "testing"

func TestHello(t *testing.T) {
	Hello()
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello()
	}
}
