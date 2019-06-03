package syncpooltest

import "testing"

type DataContainer struct {
	a []byte
}

var d *DataContainer

func Benchmark_Inline_Heap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// heap
		d = &DataContainer{
			//heap
			make([]byte, 1024),
		}
		_ = d
	}
}

func Benchmark_NotInline_Heap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// heap
		d = &DataContainer{}
		// heap
		d.a = make([]byte, 1024)
		_ = d
	}
}

func Benchmark_Inline_Stack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// stack
		d := &DataContainer{
			// stack
			make([]byte, 1024),
		}
		_ = d
	}
}

func Benchmark_NotInline_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// d is in stack
		d := &DataContainer{}
		// d.a is in heap
		d.a = make([]byte, 1024)
		_ = d
	}
}
