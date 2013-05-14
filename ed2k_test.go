package ed2k

import "fmt"

func ExampleNoInput() {
	ed2k := New(false)
	fmt.Println(ed2k)
	// Output: 31d6cfe0d16ae931b73c59d7e0c089c0
}

func ExampleSmall() {
	ed2k := New(false) // mode is irrelevant for small inputs
	ed2k.Write([]byte("small example"))
	fmt.Println(ed2k)
	// Output: 3e01197bc54364cb86a41738b06ae679
}

func nullTest(nullChunk bool, size int) {
	zeros := make([]byte, size)
	ed2k := New(nullChunk)
	ed2k.Write(zeros)
	fmt.Println(ed2k)
}

func ExampleSingleChunk_nullChunk() {
	nullTest(true, 9728000)
	// Output: fc21d9af828f92a8df64beac3357425d
}

func ExampleTwoChunks_nullChunk() {
	nullTest(true, 2*9728000)
	// Output: 114b21c63a74b6ca922291a11177dd5c
}

func ExampleSingleChunk_noNullChunk() {
	nullTest(false, 9728000)
	// Output: d7def262a127cd79096a108e7a9fc138
}

func ExampleTwoChunks_noNullChunk() {
	nullTest(false, 2*9728000)
	// Output: 194ee9e4fa79b2ee9f8829284c466051
}