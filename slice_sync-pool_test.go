package syncpooltest

import (
	"sync"
	"testing"
)

func Benchmark_Stack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := make([]byte, 1024)
		_ = data
	}
}

var data []byte

func Benchmark_Heap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data = make([]byte, 1024)
	}
}

func Benchmark_Pool_ReturnNonpointer(b *testing.B) {
	var bytePool = sync.Pool{
		New: func() interface{} {
			b := make([]byte, 1024)
			return b
		},
	}
	for i := 0; i < b.N; i++ {
		data = bytePool.Get().([]byte)
		_ = data
		bytePool.Put(data)
	}
}

var data_pointer *[]byte

func Benchmark_Pool_ReturnPointer(b *testing.B) {
	var bytePool = sync.Pool{
		New: func() interface{} {
			b := make([]byte, 1024)
			return &b
		},
	}
	for i := 0; i < b.N; i++ {
		data_pointer := bytePool.Get().(*[]byte)
		_ = data_pointer
		bytePool.Put(data_pointer)
	}
}
