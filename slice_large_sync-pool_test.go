package syncpooltest

import (
	"sync"
	"testing"
)

const size = 64 * 1024 //65536

func Benchmark_LargeSize_Stack_EqualOrLess65535(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// move to heap when  size > 65535
		dataLarge := make([]byte, size-1)
		_ = dataLarge
	}
}
func Benchmark_LargeSize_Heap_LargerThan65535(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// move to heap when  size > 65535
		dataLarge := make([]byte, size)
		_ = dataLarge
	}
}

var dataLarge []byte

func Benchmark_LargeSize_Heap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dataLarge = make([]byte, size)
	}
}

func Benchmark_LargeSize_Pool_ReturnNonpointer(b *testing.B) {
	var bytePool = sync.Pool{
		New: func() interface{} {
			b := make([]byte, size)
			return b
		},
	}
	for i := 0; i < b.N; i++ {
		dataLarge = bytePool.Get().([]byte)
		_ = dataLarge
		bytePool.Put(dataLarge)
	}
}

var dataLargePointer *[]byte

func Benchmark_LargeSize_Pool_ReturnPointer(b *testing.B) {
	var bytePool = sync.Pool{
		New: func() interface{} {
			b := make([]byte, size)
			return &b
		},
	}
	for i := 0; i < b.N; i++ {
		dataLargePointer := bytePool.Get().(*[]byte)
		_ = dataLargePointer
		bytePool.Put(dataLargePointer)
	}
}
