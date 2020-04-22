package sos
import "testing"

// sample code
var a []int

// initialize with sample code
func init() {
    for i := 0; i < 1000000; i++ {
        a = append(a, i)
    }
}

func BenchmarkConcurrent(b *testing.B) {
    for n := 0; n < b.N; n++ {
        MergeSortMulti(a)
    }
}


func BenchmarNotConcurrent(b *testing.B) {
    for n := 0; n < b.N; n++ {
        MergeSort(a)
    }
}
