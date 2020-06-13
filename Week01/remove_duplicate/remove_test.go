package remove_duplicate

import (
	"testing"
)

func BenchmarkRemoveDuplicates(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	}
}

func BenchmarkRemoveDuplicates2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		RemoveDuplicates2([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	}
}

//goos: darwin
//goarch: amd64
//pkg: algorithm010/Week01/remove_duplicate
//BenchmarkRemoveDuplicates-8    	93278048	        12.3 ns/op	       0 B/op	       0 allocs/op
//BenchmarkRemoveDuplicates2-8   	66292068	        15.7 ns/op	       0 B/op	       0 allocs/op
//PASS
//ok  	algorithm010/Week01/remove_duplicate	2.231s
