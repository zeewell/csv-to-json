package main

import "testing"

func BenchmarkSelfConcatOperator1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		main()
	}
}
